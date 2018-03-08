CREATE TABLE users  (
    userid INTEGER PRIMARY KEY,
    fullname TEXT,
    nickname TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL
);

CREATE TABLE compliments (
    complimentid INTEGER PRIMARY KEY,
    sender TEXT NOT NULL,
    FOREIGN KEY (sender) REFERENCES users(nickname),
    receiver TEXT NOT NULL,
    FOREIGN KEY (sender) REFERENCES users(nickname),
    public INTEGER NOT NULL
);

