CREATE TABLE `setting`
(
    `setting_id` int NOT NULL AUTO_INCREMENT,
    `type` varchar(255) NOT NULL,
    `name` varchar(255) NOT NULL,
    `created_at`  datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`  datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY(`setting_id`)
)