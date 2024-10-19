-- Create the default database
CREATE DATABASE IF NOT EXISTS test_db;

-- Create the user with no password
CREATE USER 'user'@'%' IDENTIFIED BY '';

-- Grant read/write access to all databases
GRANT SELECT, INSERT, UPDATE, DELETE ON *.* TO 'user'@'%';

-- Grant all privileges on the test_db database
GRANT ALL PRIVILEGES ON test_db.* TO 'user'@'%';

-- Flush privileges to apply changes
FLUSH PRIVILEGES;
