CREATE DATABASE
IF NOT EXISTS game;

USE game;

CREATE TABLE
IF NOT EXISTS action
(
    id int NOT NULL AUTO_INCREMENT,
    userId int NOT NULL,
    gameId int NOT NULL,
    action TEXT NOT NULL,
    PRIMARY KEY
(id)
);