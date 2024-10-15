CREATE TABLE IF NOT EXISTS category (
    id VARCHAR(255) PRIMARY KEY DEFAULT gen_random_uuid() UNIQUE,
    name VARCHAR(255) NOT NULL UNIQUE,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now()
);