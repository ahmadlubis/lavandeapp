CREATE TABLE `users`
(
    `id`                  bigint(20) unsigned  NOT NULL AUTO_INCREMENT,
    `name`                varchar(255)             NULL,
    `nik`                 char(16)                 NULL,
    `religion`            varchar(255)             NULL,
    `email`               varchar(255)         NOT NULL,
    `phone_no`            varchar(32)              NULL,
    `role`                smallint    unsigned NOT NULL,
    `status`              smallint    unsigned NOT NULL,
    `password`            binary(60)           NOT NULL,
    `created_at`          datetime             NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`          datetime             NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `users_on_nik` (`nik`) USING BTREE,
    UNIQUE KEY `users_on_email` (`email`) USING BTREE,
    UNIQUE KEY `users_on_phone_no` (`phone_no`) USING BTREE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;
