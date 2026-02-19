-- Create makes table
CREATE TABLE makes (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create models table
CREATE TABLE models (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    make_id INTEGER REFERENCES makes(id) ON DELETE CASCADE,
    year_range VARCHAR(50), 
    generation VARCHAR(100),
    location VARCHAR(50) DEFAULT 'Global',
    codename VARCHAR(100),
    image_url TEXT,
    known_for TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT uq_models_make_name_gen_year UNIQUE(make_id, name, year_range, generation)
);

-- Create challenges table
CREATE TABLE challenges (
    id SERIAL PRIMARY KEY,
    date DATE UNIQUE NOT NULL,
    image_url TEXT NOT NULL,
    solution_make_id INTEGER REFERENCES makes(id),
    solution_model_id INTEGER REFERENCES models(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create users table
-- Users can be anonymous (just anonymous_id) or linked to Google account
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    anonymous_id UUID UNIQUE NOT NULL,
    google_id VARCHAR(255) UNIQUE,  -- Google's unique user ID
    email VARCHAR(255) UNIQUE,
    display_name VARCHAR(255),
    profile_picture_url TEXT,
    is_linked BOOLEAN DEFAULT FALSE,  -- true if linked to Google account
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create submissions table
CREATE TABLE submissions (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    challenge_id INTEGER REFERENCES challenges(id) ON DELETE CASCADE,
    make_id INTEGER,
    model_id INTEGER,
    is_correct BOOLEAN DEFAULT FALSE,
    is_make_correct BOOLEAN DEFAULT FALSE,
    is_model_correct BOOLEAN DEFAULT FALSE,
    attempt_number INTEGER NOT NULL DEFAULT 1,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create user_challenge_scores table to store final calculated scores
-- Points system:
-- 5 points: Both correct on 1st attempt
-- 3 points: Both correct on 2nd attempt
-- 1 point: Both correct on 3rd attempt
-- 0.5 points: Make correct in latest submission (only if consistently correct)
-- Bonus points: 1 point each for correct year_range, generation, codename
CREATE TABLE user_challenge_scores (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    challenge_id INTEGER REFERENCES challenges(id) ON DELETE CASCADE,
    attempt_number INTEGER NOT NULL,  -- Which attempt they solved it on (1, 2, or 3)
    full_solve_points DECIMAL(3,1) DEFAULT 0,  -- 5, 3, or 1 for full solve
    make_bonus_points DECIMAL(3,1) DEFAULT 0,  -- 0.5 if make was consistently correct
    bonus_round_points DECIMAL(3,1) DEFAULT 0,  -- up to 3 points for year_range, generation, codename
    total_points DECIMAL(3,1) DEFAULT 0,  -- full_solve_points + make_bonus_points + bonus_round_points
    is_fully_solved BOOLEAN DEFAULT FALSE,  -- true if both make and model are correct
    make_ever_wrong BOOLEAN DEFAULT FALSE,  -- true if make was wrong in any attempt (edge case)
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT unique_user_challenge_score UNIQUE(user_id, challenge_id)
);

-- Create bonus_submissions table to track bonus round attempts
-- Each bonus type (year_range, generation, codename) can only be attempted once
-- Each correct answer gives 1 point
CREATE TABLE bonus_submissions (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    challenge_id INTEGER REFERENCES challenges(id) ON DELETE CASCADE,
    bonus_type VARCHAR(20) NOT NULL CHECK (bonus_type IN ('year_range', 'generation', 'codename')),
    submitted_value VARCHAR(100) NOT NULL,
    is_correct BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT unique_user_challenge_bonus UNIQUE(user_id, challenge_id, bonus_type)
);

-- Create indexes
CREATE INDEX idx_models_make_id ON models(make_id);
CREATE INDEX idx_challenges_date ON challenges(date);
CREATE INDEX idx_submissions_user_challenge ON submissions(user_id, challenge_id);
CREATE INDEX idx_user_challenge_scores_user ON user_challenge_scores(user_id);
CREATE INDEX idx_user_challenge_scores_challenge ON user_challenge_scores(challenge_id);
CREATE INDEX idx_user_challenge_scores_points ON user_challenge_scores(total_points DESC);
CREATE INDEX idx_bonus_submissions_user_challenge ON bonus_submissions(user_id, challenge_id);
CREATE INDEX idx_users_google_id ON users(google_id);
CREATE INDEX idx_users_email ON users(email);

