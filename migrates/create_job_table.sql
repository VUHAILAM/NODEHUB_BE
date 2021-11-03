CREATE TABLE `job`
(
    `job_id` bigint(20) NOT NULL AUTO_INCREMENT,
    `recruiter_id` bigint(20) NOT NULL,
    `title` varchar(255) CHARSET utf8 NOT NULL,
    `description` LONGTEXT CHARSET utf8 NOT NULL,
    `salary_range` varchar(255) CHARSET utf8 NOT NULL,
    `quantity` int NOT NULL ,
    `role` varchar(255) CHARSET utf8 NOT NULL,
    `experience` varchar(255) CHARSET utf8 NOT NULL,
    `location` varchar(255) CHARSET utf8 NOT NULL,
    `hire_date` datetime,
    `status` tinyint(1),
    `created_at`  datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`  datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY(`job_id`)
)