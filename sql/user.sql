CREATE TABLE IF NOT EXISTS users (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid() UNIQUE,
    username text NOT NULL UNIQUE,
    password text NOT NULL,
    role text NOT NULL,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now()
);

CREATE TABLE IF NOT EXISTS user_borrowed_books (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid() UNIQUE,
    user_id uuid NOT NULL,
    book_title text NOT NULL,
    quantity integer NOT NULL,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    FOREIGN KEY (user_id) REFERENCES users(id)
);