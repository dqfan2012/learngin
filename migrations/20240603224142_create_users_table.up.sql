CREATE TABLE users (
   id SERIAL PRIMARY KEY,
   first_name VARCHAR(255) NOT NULL,
   last_name VARCHAR(255) NOT NULL,
   email VARCHAR(255) UNIQUE NOT NULL,
   password VARCHAR(255) NOT NULL,
   role VARCHAR(255) CHECK (role IN ('reader', 'publisher')) DEFAULT 'reader',
   remember_token VARCHAR(100),
   created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
   updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
   email_verified_at TIMESTAMPTZ NULL,
   deleted_at TIMESTAMPTZ NULL
);
