-- Your SQL goes here
CREATE TABLE users (
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    username varchar(500) NOT NULL,
    token varchar(100) NOT NULL,
    wins int NOT NULL,
    losses int NOT NULL
);
