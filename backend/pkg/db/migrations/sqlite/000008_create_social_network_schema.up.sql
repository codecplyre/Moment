CREATE TABLE IF NOT EXISTS "FollowNotif"(
    "userId" TEXT NOT NULL,
    "followingID" TEXT NOT NULL,
    "createdAt" DATETIME NOT NULL,
    "status" TEXT NOT NULL,
    "unread" BOOLEAN NOT NULL,
    FOREIGN KEY("userId") 
        REFERENCES "User"("id"),
    FOREIGN KEY("followingID") 
        REFERENCES "User"("id")
);