-- 006_subscriptions_hybrid_up.sql
-- Add hybrid provider support to users table

ALTER TABLE users ADD COLUMN IF NOT EXISTS subscription_provider VARCHAR(50);
-- 'razorpay' or 'revenuecat'

-- Update subscription_events to be more generic 
ALTER TABLE subscription_events RENAME COLUMN rc_event_id TO external_event_id;
ALTER TABLE subscription_events ADD COLUMN IF NOT EXISTS provider VARCHAR(50);

-- Update existing records if any (assuming they were RC based on current code)
UPDATE users SET subscription_provider = 'revenuecat' WHERE is_subscriber = TRUE AND subscription_provider IS NULL;
UPDATE subscription_events SET provider = 'revenuecat' WHERE provider IS NULL;
