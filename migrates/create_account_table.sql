CREATE TABLE `account`
(
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    `information_id` bigint(20) NOT NULL,
    `email` varchar(50) NOT NULL,
    `password` varchar(255) NOT NULL,
    `token_hash` varchar(255) NOT NULL,
    `phone` varchar(50) NOT NULL,
    `type` int NOT NULL,
    `status` tinyint(1),
    `created_at`  datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`  datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY(`id`),
    UNIQUE KEY `email` (`email`)
)