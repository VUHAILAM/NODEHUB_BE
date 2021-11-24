CREATE TABLE `follow`
(
    `id` BIGINT NOT NULL AUTO_INCREMENT,
    `candidate_id` bigint(20) NOT NULL,
    `recruiter_id` bigint(20) NOT NULL,
    `created_at`  datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`  datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY(`id`)
)