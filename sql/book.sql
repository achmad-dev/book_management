CREATE TABLE IF NOT EXISTS books (
    id VARCHAR(255) PRIMARY KEY DEFAULT gen_random_uuid() UNIQUE,
    author_id VARCHAR(255) NOT NULL,
    category_id VARCHAR(255) NOT NULL,
    title VARCHAR(255) NOT NULL UNIQUE,
    author VARCHAR(255) NOT NULL,
    category VARCHAR(255) NOT NULL,
    stock INT NOT NULL DEFAULT 0,
    borrowed INT NOT NULL DEFAULT 0,
    is_popular BOOLEAN NOT NULL DEFAULT FALSE,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now()
);