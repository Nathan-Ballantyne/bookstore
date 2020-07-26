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
book_id INTEGER NOT NULL
);

-- DROP TABLE list_type;

CREATE TABLE list_type (
id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
user_id INTEGER NOT NULL,
name VARCHAR(255) NOT NULL
);

-- Add an index on the type column.
CREATE INDEX idx_bookstore_type ON list_type(user_id);

-- DROP TABLE book;

-- add values to tables 
-- INSERT INTO users (name, email, hashed_password, active) VALUES
                 --   ('John Smith','john.smith@email.com','hgsbmciwekuysyjizfzivnxiebzpyoryxjvdwvrdogvlmemvszzmeimkjuws',TRUE);
                 
-- ADD list type 
INSERT INTO list_type (user_id, name) VALUES (1, 'WishList');
                 
-- add books 
INSERT INTO book (title, author, release_year, page_count, cover, series, read_status, rating) VALUES 
				('The Name if the Wind', 'Patrick Rothfuss', 2008, 672, '../../assets/images/Name_of_the_Wind_cover.jpg', 'Kingkiller Chronicle', 'read', 10),
                ('The Wise Man\'s Fear', 'Patrick Rothfuss', 2012, 1008, 'N/A', 'Kingkiller Chronicle', 'read', 10),
                ('The Slow Regard of Silent Things', 'Patrick Rothfuss', 2016, 176, 'N/A', 'Kingkiller Chronicle', 'read', 10),
                ('The Way of Kings', 'Brandon Sanderson', 2010, 1258, 'N/A', 'Stormlight Archive', 'read', 8),
                ('Words of Radiance', 'Brandon Sanderson', 2014, 1328, 'N/A', 'Stormlight Archive', 'read', 10),
                ('Oathbringer', 'Brandon Sanderson', 2017, 1243, 'N/A', 'Stormlight Archive', 'read', 9);
        
UPDATE book 
SET 
	cover = '../../assets/images/Name_of_the_Wind_cover.jpg'
WHERE
	id = 1;
        
CREATE USER 'web'@'localhost';
GRANT SELECT, INSERT, UPDATE ON bookstore.* TO 'web'@'localhost';
-- Important: Make sure to swap 'pass' with a password of your own choosing.
ALTER USER 'web'@'localhost' IDENTIFIED BY 'pass';

                