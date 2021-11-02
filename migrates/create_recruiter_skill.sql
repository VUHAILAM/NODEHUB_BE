CREATE TABLE `recruiter_skill`
(
    `id` int NOT NULL AUTO_INCREMENT,
    `recruiter_id` bigint,
    `skill_id` int,
    `created_at`  datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`  datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY(`id`),
    FOREIGN KEY (`recruiter_id`) REFERENCES recruiter(`recruiter_id`),
    FOREIGN KEY (`skill_id`) REFERENCES skill(`skill_id`)
)