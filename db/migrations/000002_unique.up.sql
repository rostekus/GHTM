-- Add unique constraints on email and username
ALTER TABLE users
ADD CONSTRAINT unique_email UNIQUE (email),
ADD CONSTRAINT unique_username UNIQUE (username);

-- Add an index on email
CREATE INDEX idx_email ON users(email);
