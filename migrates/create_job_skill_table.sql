CREATE TABLE `job_skill`
(
    `id` BIGINT NOT NULL AUTO_INCREMENT,
    `job_id` bigint(20) NOT NULL,
    `skill_id` bigint(20) NOT NULL,
    `created_at`  datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`  datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY(`id`)
)
