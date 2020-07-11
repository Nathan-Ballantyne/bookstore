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
title VARCHAR(255) NOT NULL,
author VARCHAR(255) NOT NULL,
release_year INTEGER NOT NULL,
page_count INTEGER NOT NULL,
cover VARCHAR(255) NOT NULL,
series VARCHAR(255) NOT NULL,
read_status VARCHAR(255) NOT NULL,
rating INTEGER NOT NULL
);

CREATE TABLE list_content (
type_id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
book_id INTEGER NOT NULL,
user_id INTEGER NOT NULL
);

-- DROP TABLE list;

CREATE TABLE list_type (
id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
name VARCHAR(255) NOT NULL
);

-- Add an index on the type column.
CREATE INDEX idx_bookstore_type ON list(type);

-- DROP TABLE book;

-- add values to tables 
-- INSERT INTO users (name, email, hashed_password, active) VALUES
                 --   ('John Smith','john.smith@email.com','hgsbmciwekuysyjizfzivnxiebzpyoryxjvdwvrdogvlmemvszzmeimkjuws',TRUE);
                 
-- add books 
INSERT INTO book (title, author, release_year, page_count, cover, series, read_status, rating) VALUES 
				('Patrick Rothfuss', 'The Name if the Wind', 2008, 672, 'N/A', 'Kingkiller Chronicle', 'read', 10),
                ('Patrick Rothfuss', 'The Wise Man\'s Fear', 2012, 1008, 'N/A', 'Kingkiller Chronicle', 'read', 10),
                ('Patrick Rothfuss', 'The Slow Regard of Silent Things', 2016, 176, 'N/A', 'Kingkiller Chronicle', 'read', 10),
                ('Brandon Sanderson', 'The Way of Kings', 2010, 1258, 'N/A', 'Stormlight Archive', 'read', 8),
                ('Brandon Sanderson', 'Words of Radiance', 2014, 1328, 'N/A', 'Stormlight Archive', 'read', 10),
                ('Brandon Sanderson', 'Oathbringer', 2017, 1243, 'N/A', 'Stormlight Archive', 'read', 9);
                
CREATE USER 'web'@'localhost';
GRANT SELECT, INSERT, UPDATE ON bookstore.* TO 'web'@'localhost';
-- Important: Make sure to swap 'pass' with a password of your own choosing.
ALTER USER 'web'@'localhost' IDENTIFIED BY '1045';

                