package config

import (
	"context"
	"fmt"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"

	"github.com/spf13/viper"
)

// TODO: move this configuration to YAML or something
type secretConfigType struct {
	secretKey string
	configKey string
}

var databasePassword = secretConfigType{
	secretKey: "projects/621422061156/secrets/cloud-sql-feedback-password/versions/latest",
	configKey: "database-password",
}

var mailgunApiKey = secretConfigType{
	secretKey: "projects/621422061156/secrets/mailgun-api-key/versions/latest",
	configKey: "mailgun-api-key",
}

var mailgunValidationKey = secretConfigType{
	secretKey: "projects/621422061156/secrets/mailgun-validation-key",
	configKey: "mailgun-validation-key",
}

func InitConfig() error {
	if err := addSecretToConfig(databasePassword); err != nil {
		return err
	}

	if err := addSecretToConfig(mailgunApiKey); err != nil {
		return err
	}

	if err := addSecretToConfig(mailgunValidationKey); err != nil {
		return err
	}

	return nil
}

func GetMailgunValidationKey() string {
	return viper.GetString(mailgunValidationKey.configKey)
}

func GetMailgunApiKey() string {
	return viper.GetString(mailgunApiKey.configKey)
}

func GetDbPassword() string {
	return viper.GetString(databasePassword.configKey)
}

func addSecretToConfig(secretConfig secretConfigType) error {
	secret, err := fetchSecret(secretConfig.secretKey)
	if err != nil {
		return err
	}

	viper.Set(secretConfig.configKey, secret)

	return nil
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
