-- SQLite
-- kasu kaivitamiseks parem klops - run selected query

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

-- DELETE FROM user WHERE id BETWEEN 2 AND 4;

-- INSERT INTO posts (user_id, title, text, created_at) VALUES (1, 'My last Post', 'back again!', '2023-06-03 12:35:56');
-- INSERT INTO posts (user_id, title, text, created_at) VALUES (1, 'My last Post', 'back again!', '2023-06-03 12:35:56');
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
--     id INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE NOT NULL,
--     sender_id INTEGER NOT NULL,
--     receiver_id INTEGER NOT NULL,
--     message VARCHAR NOT NULL,
--     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
--     FOREIGN KEY (sender_id) REFERENCES user (id),
--     FOREIGN KEY (receiver_id) REFERENCES user (id)
-- );

-- Create comments table
CREATE TABLE IF NOT EXISTS comments (
	id INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE NOT NULL,
	user_id INTEGER NOT NULL,
    post_id INTEGER NOT NULL,
	text VARCHAR NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	    FOREIGN KEY (user_id) REFERENCES user(id)
    );


-- deleted column dislike
-- ALTER TABLE posts
-- DROP COLUMN dislike;


-- INSERT INTO messages (sender_id, receiver_id, message, created_at) VALUES (5, 1, 'five here', '2023-06-13 12:36:56');

-- in case i need to change type of likes and dislikes
-- ALTER TABLE posts MODIFY likes INT;
-- ALTER TABLE posts MODIFY dislikes INT;


