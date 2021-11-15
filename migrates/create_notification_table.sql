CREATE TABLE `notification`
(
    `account_id` bigint(20) NOT NULL,
    `title` varchar(255) NOT NULL,
    `content`  varchar(255) NOT NULL,
    `key`  varchar(255) NOT NULL,
    `check_read`  bool NOT NULL,
    `created_at`  datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`  datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY(`account_id`),
    FOREIGN KEY (`account_id`) REFERENCES account(`id`)
)