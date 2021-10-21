CREATE TABLE `skill`
(
    `skill_id` int NOT NULL AUTO_INCREMENT,
    `name` varchar(255) NOT NULL,
    `description` longtext NOT NULL,
    `questions` longtext NOT NULL,
    `icon` varchar(255) NOT NULL,
    `status` bool NOT NULL,
    `created_at`  datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`  datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY(`skill_id`)
)