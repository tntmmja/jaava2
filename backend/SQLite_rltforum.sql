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


/*
create table user (
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

-- kustutab tabeli posts
/*drop table user
*/

ALTER TABLE user RENAME COLUMN username TO nickname;

-- Create posts table
CREATE TABLE posts (
    id INTEGER PRIMARY KEY,
    title TEXT,
    content TEXT,
    user_id INTEGER,
    FOREIGN KEY (user_id) REFERENCES users (id)
);

-- Create messages table
CREATE TABLE messages (
    id INTEGER PRIMARY KEY,
    sender_id INTEGER,
    receiver_id INTEGER,
    message TEXT,
    timestamp DATETIME,
    FOREIGN KEY (sender_id) REFERENCES users (id),
    FOREIGN KEY (receiver_id) REFERENCES users (id)
);


