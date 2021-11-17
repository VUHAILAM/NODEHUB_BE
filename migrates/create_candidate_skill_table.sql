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
)