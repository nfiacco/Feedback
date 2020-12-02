package database

import (
	"context"
	"fmt"
	"log"
	"os"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const DB_PWD_SECRET = "projects/621422061156/secrets/cloud-sql-feedback-password/versions/latest"

func InitDatabase() (*gorm.DB, error) {
	if isProd() {
		return initDatabaseProd()
	} else {
		return initDatabaseDev()
	}
}

func initDatabaseDev() (*gorm.DB, error) {
	var dbURI string
	dbURI = fmt.Sprintf("user=postgres database=feedback host=localhost")

	db, err := gorm.Open(postgres.Open(dbURI), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("sql.Open: %v", err)
	}

	return db, nil
}

func initDatabaseProd() (*gorm.DB, error) {
	var (
		dbUser                 = mustGetenv("DB_USER")
		instanceConnectionName = mustGetenv("INSTANCE_CONNECTION_NAME")
		dbName                 = mustGetenv("DB_NAME")
	)

	dbPwd, err := fetchSecret(DB_PWD_SECRET)
	if err != nil {
		return nil, err
	}

	socketDir, isSet := os.LookupEnv("DB_SOCKET_DIR")
	if !isSet {
		socketDir = "/cloudsql"
	}

	var dbURI string
	dbURI = fmt.Sprintf("user=%s password=%s database=%s host=%s/%s", dbUser, *dbPwd, dbName, socketDir, instanceConnectionName)

	db, err := gorm.Open(postgres.Open(dbURI), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("sql.Open: %v", err)
	}

	return db, nil
}

func mustGetenv(k string) string {
	v := os.Getenv(k)
	if v == "" {
		log.Fatalf("Warning: %s environment variable not set.\n", k)
	}
	return v
}

func fetchSecret(name string) (*string, error) {
	ctx := context.Background()
	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create secretmanager client: %v", err)
	}

	req := &secretmanagerpb.AccessSecretVersionRequest{
		Name: name,
	}

	result, err := client.AccessSecretVersion(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to access secret version: %v", err)
	}

	secret := string(result.Payload.Data)
	return &secret, nil
}

func isProd() bool {
	_, isSet := os.LookupEnv("IS_PROD")
	return isSet
}
