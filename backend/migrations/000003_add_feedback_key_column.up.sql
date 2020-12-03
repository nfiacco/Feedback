ALTER TABLE users ADD COLUMN IF NOT EXISTS feedback_key VARCHAR(8);

CREATE INDEX users_feedback_key_idx ON users(feedback_key);
