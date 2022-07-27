DROP TABLE IF EXISTS users
CREATE TABLE users (
	id text PRIMARY KEY,
	full_name text,
	email text,
	phone text,
    created_at TIMESTAMP with time zone,
    updated_at TIMESTAMP with time zone
)

DROP TABLE IF EXISTS news
CREATE TABLE news (
	id text PRIMARY KEY,
	title text,
	summary text,
	author text,
	content text,
	language text,
	image text
)
