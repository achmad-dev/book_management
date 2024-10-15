CREATE TABLE IF NOT EXISTS authors (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid() UNIQUE,
    name text NOT NULL UNIQUE,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now()
);