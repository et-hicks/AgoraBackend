CREATE TABLE threads (
    id int NOT NULL
    createdTime time NOT NULL
    lastCommentedTime time NOT NULL
    title varchar(255) NOT NULL
    creatorId int NOT NULL
    likes int NOT NULL
    dislikes int NOT NULL
    clicks int NOT NULL
    watchers int NOT NULL
    urlUUID varchar(255) NOT NULL
    imageurl text(65,535) NOT NULL
    isPublic boolean NOT NULL
    isReported boolean NOT NULL
    bfs json NOT NULL
    description varchar(255) NOT NULL
)
