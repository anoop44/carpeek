-- 005_subscriptions_up.sql
-- Add subscription support to users table

ALTER TABLE users ADD COLUMN IF NOT EXISTS is_subscriber BOOLEAN DEFAULT FALSE;
ALTER TABLE users ADD COLUMN IF NOT EXISTS subscription_status VARCHAR(50);
ALTER TABLE users ADD COLUMN IF NOT EXISTS subscription_product_id VARCHAR(255);
ALTER TABLE users ADD COLUMN IF NOT EXISTS subscription_expires_at TIMESTAMP;

-- Index for subscriber filtering (e.g. deciding whether to show ads)
CREATE INDEX IF NOT EXISTS idx_users_is_subscriber ON users(is_subscriber) WHERE is_subscriber = TRUE;

-- Subscription event log for audit trail
CREATE TABLE IF NOT EXISTS subscription_events (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    event_type VARCHAR(50) NOT NULL,
    rc_event_id VARCHAR(255),
    details TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_subscription_events_user_id ON subscription_events(user_id);
