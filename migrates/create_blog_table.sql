CREATE TABLE `blog`
(
    `blog_id` int NOT NULL AUTO_INCREMENT,
    `category_id` int,
    `title` varchar(255) NOT NULL,
    `icon` varchar(255) NOT NULL,
    `description` longtext NOT NULL,
    `status` boolean NOT NULL DEFAULT true,
    `created_at`  datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`  datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY(`blog_id`)
    FOREIGN KEY (`category_id`) REFERENCES setting(`setting_id`)
)