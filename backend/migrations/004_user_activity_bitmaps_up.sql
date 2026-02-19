-- Create user_activity_bitmaps table for efficient streak tracking
CREATE TABLE user_activity_bitmaps (
    user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    month_date DATE NOT NULL, -- First day of the month (e.g., '2026-02-01')
    participation_bitmap INTEGER DEFAULT 0, -- 32 bits, each bit represents a day of the month
    submission_bitmap INTEGER DEFAULT 0, -- 32 bits, each bit represents a day of the month
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (user_id, month_date)
);

CREATE INDEX idx_user_activity_bitmaps_month ON user_activity_bitmaps(month_date);
