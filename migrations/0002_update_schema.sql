-- Update database schema to match current models
-- Version: 2.0
-- Created: 2024-06-15

BEGIN;

-- Drop existing tables if they exist
DROP TABLE IF EXISTS swipes;
DROP TABLE IF EXISTS matches;
DROP TABLE IF EXISTS profiles;
DROP TABLE IF EXISTS users;

-- Create updated users table
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email TEXT UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    last_login_at TIMESTAMP,
    is_active BOOLEAN DEFAULT TRUE NOT NULL,
    is_verified BOOLEAN DEFAULT FALSE NOT NULL
);

-- Create updated profiles table
CREATE TABLE profiles (
    id SERIAL PRIMARY KEY,
    user_id INTEGER UNIQUE NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    display_name TEXT NOT NULL,
    bio TEXT,
    gender TEXT,
    birth_date DATE,
    location TEXT,
    photos TEXT[],
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);

-- Create updated matches table with status
CREATE TABLE matches (
    id SERIAL PRIMARY KEY,
    user1_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    user2_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    status TEXT NOT NULL DEFAULT 'pending',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    UNIQUE(user1_id, user2_id)
);

-- Create messages table
CREATE TABLE messages (
    id SERIAL PRIMARY KEY,
    sender_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    receiver_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    content TEXT NOT NULL,
    is_read BOOLEAN DEFAULT FALSE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);

-- Create preferences table
CREATE TABLE preferences (
    id SERIAL PRIMARY KEY,
    user_id INTEGER UNIQUE NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    match_distance INTEGER DEFAULT 50 NOT NULL,
    min_age INTEGER DEFAULT 18 NOT NULL,
    max_age INTEGER DEFAULT 99 NOT NULL,
    notify_new_matches BOOLEAN DEFAULT TRUE NOT NULL,
    notify_messages BOOLEAN DEFAULT TRUE NOT NULL,
    show_online_status BOOLEAN DEFAULT TRUE NOT NULL,
    show_last_active BOOLEAN DEFAULT TRUE NOT NULL,
    show_distance BOOLEAN DEFAULT TRUE NOT NULL
);

-- Create indexes for better query performance
CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_profiles_user_id ON profiles(user_id);
CREATE INDEX idx_matches_user1_id ON matches(user1_id);
CREATE INDEX idx_matches_user2_id ON matches(user2_id);
CREATE INDEX idx_messages_sender_id ON messages(sender_id);
CREATE INDEX idx_messages_receiver_id ON messages(receiver_id);
CREATE INDEX idx_preferences_user_id ON preferences(user_id);

COMMIT;
