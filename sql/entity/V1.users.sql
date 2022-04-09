CREATE TABLE user (
    id int NOT NULL
    createdTime time NOT NULL
    lastUpdatedTime time NOT NULL
    firstName text(255) NULL
    lastName text(255) NULL
    userName char(82) NOT NULL
    email varchar(255) NOT NULL
    password char(82) NOT NULL
    phoneNumber int NOT NULL
    phoneCode int NOT NULL
    AccountType text(255) NOT NULL
    fefs json NOT NULL
)