CREATE DATABASE bookstore CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

USE bookstore;

CREATE TABLE users (
id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
name VARCHAR(255) NOT NULL,
email VARCHAR(255) NOT NULL,
hashed_password CHAR(60) NOT NULL,
active BOOLEAN NOT NULL DEFAULT TRUE
);
ALTER TABLE users ADD CONSTRAINT users_uc_email UNIQUE (email);


CREATE TABLE book (
id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
author VARCHAR(255) NOT NULL,
release_year INTEGER NOT NULL,
page_count INTEGER NOT NULL,
cover VARCHAR(255) NOT NULL,
series VARCHAR(255) NOT NULL,
read_status VARCHAR(255) NOT NULL,
rating INTEGER NOT NULL
);

CREATE TABLE list (
id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
book_id INTEGER NOT NULL,
user_id INTEGER NOT NULL,
type VARCHAR(255) NOT NULL
);

-- Add an index on the type column.
CREATE INDEX idx_bookstore_type ON list(type);



--add values to tables 
INSERT INTO users (name, email, hashed_password, active) VALUES
                    ('John Smith','john.smith@email.com','hgsbmciwekuysyjizfzivnxiebzpyoryxjvdwvrdogvlmemvszzmeimkjuws',TRUE);