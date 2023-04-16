CREATE TABLE IF NOT EXISTS User (
    userId VARCHAR(255) NOT NULL PRIMARY KEY,
    firstName VARCHAR(255) NOT NULL,
    lastName VARCHAR(255) NOT NULL,
    nickName VARCHAR(255),
    email VARCHAR(255) NOT NULL,
    DOB TIMESTAMP NOT NULL,
    avatar VARCHAR(255),
    aboutMe VARCHAR(255),
    createdAt TIMESTAMP NOT NULL,
    isLoggedIn BOOLEAN DEFAULT 0,
    isPublic BOOLEAN DEFAULT 1,
    numFollowers INTEGER DEFAULT 0,
    numFollowing INTEGER DEFAULT 0,
    numPosts INTEGER DEFAULT 0,
    password VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS UserSessions (
    sessionId VARCHAR(255) NOT NULL PRIMARY KEY,
    userId VARCHAR(255) NOT NULL,
    createdAt TIMESTAMP NOT NULL,
    FOREIGN KEY (userId) REFERENCES User(userId)
);