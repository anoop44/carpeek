-- 005_subscriptions_down.sql
-- Remove subscription support

DROP TABLE IF EXISTS subscription_events;

DROP INDEX IF EXISTS idx_users_is_subscriber;

ALTER TABLE users DROP COLUMN IF EXISTS subscription_expires_at;
ALTER TABLE users DROP COLUMN IF EXISTS subscription_product_id;
ALTER TABLE users DROP COLUMN IF EXISTS subscription_status;
ALTER TABLE users DROP COLUMN IF EXISTS is_subscriber;
