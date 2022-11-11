CREATE TABLE `units`
(
    `id`                  bigint(20) unsigned  NOT NULL AUTO_INCREMENT,
    `gov_id`              varchar(191)         NOT NULL,
    `tower`               varchar(191)         NOT NULL,
    `floor`               varchar(191)         NOT NULL,
    `unit_no`             varchar(191)         NOT NULL,
    `ajb`                 mediumblob               NULL,
    `akte`                mediumblob               NULL,
    `created_at`          datetime             NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`          datetime             NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `units_on_gov_id` (`gov_id`) USING BTREE,
    UNIQUE KEY `units_on_identifiers` (`tower`, `floor`, `unit_no`) USING BTREE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

CREATE TABLE `tenants`
(
    `id`                  bigint(20) unsigned  NOT NULL AUTO_INCREMENT,
    `unit_id`             bigint(20) unsigned  NOT NULL,
    `user_id`             bigint(20) unsigned  NOT NULL,
    `role`                smallint             NOT NULL,
    `start_at`            datetime                 NULL,
    `end_at`              datetime                 NULL,
    `created_at`          datetime             NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`          datetime             NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    KEY `tenants_on_user_unit` (`user_id`, `unit_id`, `end_at`) USING BTREE,
    KEY `tenants_on_unit` (`unit_id`, `end_at`) USING BTREE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;
