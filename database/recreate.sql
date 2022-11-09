DROP DATABASE IF EXISTS hippo_db;
CREATE DATABASE hippo_db;
USE hippo_db;

DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
    `id`           INTEGER      PRIMARY KEY,
    `email`        VARCHAR(128) NOT NULL UNIQUE,
    `display_name` VARCHAR(128) NOT NULL UNIQUE
);