CREATE TABLE `notification`
(
    `notification_id` BIGINT NOT NULL AUTO_INCREMENT,
    `candidate_id` BIGINT DEFAULT 0,
    `recruiter_id` BIGINT DEFAULT 0,
    `title` varchar(255) NOT NULL,
    `content`  varchar(255) NOT NULL,
    `key`  varchar(255) NOT NULL,
    `check_read`  bool NOT NULL,
    `created_at`  datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`  datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY(`notification_id`)
)