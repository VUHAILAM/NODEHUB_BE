CREATE TABLE `job_candidate`
(
    `id` BIGINT NOT NULL AUTO_INCREMENT,
    `job_id` bigint(20) NOT NULL,
    `candidate_id` bigint(20) MOT NULL,
    `status` varchar(255) CHARSET utf8 NOT NULL,
    `created_at`  datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`  datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY(`id`)
)