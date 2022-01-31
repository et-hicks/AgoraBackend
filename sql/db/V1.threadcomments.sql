CREATE TABLE threadcomments (
    id int NOT NULL
    createdTime time NOT NULL
    lastUpdatedTime time NOT NULL
    threadId int NOT NULL
    creatorId int NOT NULL
    parentcommentId int NOT NULL
    comment text(65,535)
    isedited boolean NOT NULL
    likes int NOT NULL
    dislikes int NOT NULL
    isReported boolean NOT NULL
    fefs jsons NOT NULL
)
