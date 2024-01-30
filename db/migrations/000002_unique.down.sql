-- Remove the unique constraints on email and username
ALTER TABLE users
DROP CONSTRAINT IF EXISTS unique_email,
DROP CONSTRAINT IF EXISTS unique_username;

-- Remove the index on email
DROP INDEX IF EXISTS idx_email;
