CREATE TABLE MINION
(
    ID              IDENTITY PRIMARY KEY,
    NAME           VARCHAR(255),
);

CREATE TABLE STRING_ID_MINION
(
    ID   VARCHAR(255) PRIMARY KEY,
    NAME VARCHAR(255)
);

CREATE TABLE VERSIONED_MINION
(
    ID   INT PRIMARY KEY,
    NAME VARCHAR(255),
    VERSION INT
);
