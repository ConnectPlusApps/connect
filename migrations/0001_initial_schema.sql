-- Initial database schema for Connect+
-- Version: 1.0
-- Created: 2023-10-15

BEGIN;

-- Users table stores core user information
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username TEXT UNIQUE NOT NULL,
    email TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL,
    first_name TEXT,
    last_name TEXT,
    bio TEXT,
    gender_identity TEXT,
    sexual_orientation TEXT,
    profile_picture_url TEXT,
    date_of_birth DATE,
    location TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);

-- Profiles table stores detailed profile information
CREATE TABLE profiles (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    interests TEXT[],
    relationship_goals TEXT,
    about_me TEXT,
    photos TEXT[],
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);

-- Matches table records mutual matches between users
CREATE TABLE matches (
    id SERIAL PRIMARY KEY,
    user1_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    user2_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    matched_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    UNIQUE(user1_id, user2_id)
);

-- Swipes table records user swipe actions
CREATE TABLE swipes (
    id SERIAL PRIMARY KEY,
    swiper_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    swiped_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    direction BOOLEAN NOT NULL, -- true for right swipe, false for left
    swiped_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    UNIQUE(swiper_id, swiped_id)
);

-- Indexes for better query performance
CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_profiles_user_id ON profiles(user_id);
CREATE INDEX idx_matches_user1_id ON matches(user1_id);
CREATE INDEX idx_matches_user2_id ON matches(user2_id);
CREATE INDEX idx_swipes_swiper_id ON swipes(swiper_id);
CREATE INDEX idx_swipes_swiped_id ON swipes(swiped_id);

COMMIT;
