CREATE TABLE `recruiter`
(
    `recruiter_id` bigint(20) NOT NULL AUTO_INCREMENT,
    `account_id` bigint(20) NOT NULL,
    `company_name` varchar(255) CHARSET utf8 NOT NULL,
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
    `created_at`  datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`  datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY(`recruiter_id`)
)