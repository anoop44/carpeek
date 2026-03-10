-- 006_subscriptions_hybrid_down.sql
-- Remove hybrid provider support

ALTER TABLE users DROP COLUMN IF EXISTS subscription_provider;

-- Revert subscription_events
ALTER TABLE subscription_events RENAME COLUMN external_event_id TO rc_event_id;
ALTER TABLE subscription_events DROP COLUMN IF EXISTS provider;
