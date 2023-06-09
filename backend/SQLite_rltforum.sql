-- SQLite
-- kasu kaivitamiseks parem klops - run selected query
/*
CREATE TABLE IF NOT EXISTS user (
id INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE NOT NULL,
firstname TEXT NOT NULL,
lastname TEXT NOT NULL,
age INTEGER NOT NULL,
gender VARCHAR NOT NULL,
username VARCHAR NOT NULL,
email TEXT NOT NULL,
password TEXT NOT NULL,
createdDate REAL,
sessionID TEXT);
*/



-- CREATE TABLE user (
--     id INTEGER PRIMARY KEY AUTOINCREMENT,
--     nickname TEXT,
--     age INTEGER,
--     gender TEXT,
--     firstName TEXT,
--     lastName TEXT,
--     email TEXT,
--     password TEXT,
--     sessionID INTEGER
-- );



-- kustutab tabeli posts
-- DROP TABLE IF EXISTS user;

DELETE FROM user WHERE id BETWEEN 2 AND 4;


-- ALTER TABLE user RENAME COLUMN username TO nickname;

-- Create posts table
-- CREATE TABLE IF NOT EXISTS posts (
-- 	id INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE NOT NULL,
-- 	user_id INTEGER NOT NULL,
-- 	title VARCHAR NOT NULL,
-- 	text VARCHAR NOT NULL,
-- 	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
-- 	like VARCHAR,
-- 	dislike VARCHAR,
--     FOREIGN KEY (user_id) REFERENCES user(id)
--     );

-- Create messages table
-- CREATE TABLE messages (
--     id INTEGER PRIMARY KEY,
--     sender_id INTEGER,
--     receiver_id INTEGER,
--     message TEXT,
--     timestamp DATETIME,
--     FOREIGN KEY (sender_id) REFERENCES users (id),
--     FOREIGN KEY (receiver_id) REFERENCES users (id)
-- );



-- in case i need to change type of likes and dislikes
-- ALTER TABLE posts MODIFY likes INT;
-- ALTER TABLE posts MODIFY dislikes INT;


