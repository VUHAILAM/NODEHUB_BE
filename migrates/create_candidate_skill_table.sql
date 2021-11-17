<<<<<<< HEAD
CREATE TABLE `candidate_skill`
(
    `id` int NOT NULL AUTO_INCREMENT,
    `candidate_id` bigint,
    `skill_id` int,
    `media` varchar(255) NOT NULL,
    `created_at`  datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`  datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY(`id`),
    FOREIGN KEY (`candidate_id`) REFERENCES candidate(`candidate_id`),
    FOREIGN KEY (`skill_id`) REFERENCES skill(`skill_id`)
=======
CREATE TABLE `candidate`
(
    `candidate_id` BIGINT NOT NULL,
    `first_name` varchar(255) CHARSET utf8 NOT NULL,
    `last_name` varchar(255) CHARSET utf8 NOT NULL,
    `birth_day` varchar(50) NOT NULL,
    `address` varchar(255) CHARSET utf8 NOT NULL,
    `avatar` varchar(255) NOT NULL,
    `banner` varchar(255) NOT NULL,
    `phone` varchar(50) NOT NULL,
    `find_job` tinyint(1),
    `nodehub_review` LONGTEXT CHARSET utf8 NOT NULL,
    `cv_manage` LONGTEXT CHARSET utf8 NOT NULL,
    `experience_manage` LONGTEXT CHARSET utf8 NOT NULL,
    `social_manage` LONGTEXT CHARSET utf8 NOT NULL,
    `project_manage` LONGTEXT CHARSET utf8 NOT NULL,
    `certificate_manage` LONGTEXT CHARSET utf8 NOT NULL,
    `prize_manage` LONGTEXT CHARSET utf8 NOT NULL,
    `created_at`  datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`  datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY(`candidate_id`),
    FOREIGN KEY (`candidate_id`) REFERENCES account(`id`)
>>>>>>> 0565716 ([17-11] 1)
)