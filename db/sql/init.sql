-- データベースの作成
CREATE DATABASE IF NOT EXISTS `go_next_todo` CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
USE `go_next_todo`;

-- Users テーブルの作成
CREATE TABLE IF NOT EXISTS `users` (
    `id` CHAR(36) NOT NULL PRIMARY KEY,
    `username` VARCHAR(255) NOT NULL,
    `email` VARCHAR(255) NOT NULL UNIQUE,
    `password` VARCHAR(255) NOT NULL,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at` TIMESTAMP NULL DEFAULT NULL
);

-- Tasks テーブルの作成
CREATE TABLE IF NOT EXISTS `tasks` (
    `id` CHAR(36) NOT NULL PRIMARY KEY,
    `title` VARCHAR(255) NOT NULL,
    `description` TEXT,
    `completed` BOOLEAN DEFAULT FALSE,
    `user_id` CHAR(36) NOT NULL REFERENCES `users`(`id`),
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at` TIMESTAMP NULL DEFAULT NULL
);
