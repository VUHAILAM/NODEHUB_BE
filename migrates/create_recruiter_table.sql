CREATE TABLE `recruiter`
(
    `recruiter_id` bigint(20) NOT NULL,
    `name` varchar(255) CHARSET utf8 NOT NULL,
    `address` varchar(255) CHARSET utf8 NOT NULL,
    `avartar` varchar(255) NOT NULL,
    `banner` varchar(255) NOT NULL,
    `phone` varchar(50) NOT NULL,
    `website` varchar(255) NOT NULL,
    `description` MEDIUMTEXT CHARSET utf8 NOT NULL,
    `employee_quantity` varchar(255) NOT NULL,
    `contacter_name` varchar(255) CHARSET utf8 NOT NULL,
    `contacter_phone` varchar(50) NOT NULL,
    `media` varchar(255) NOT NULL,
    `activeVIP` bool NOT NULL, 
    `created_at`  datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`  datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY(`recruiter_id`),
    FOREIGN KEY (`recruiter_id`) REFERENCES account(`id`)
)

ALTER TABLE `recruiter`
ADD COLUMN nodehub_review varchar(255) CHARSET utf8 AFTER activeVIP;