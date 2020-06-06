USE RussianDoll;

CREATE TABLE IF NOT EXISTS `russian_doll`
(
    `id`             VARCHAR(36) NOT NULL,
    `name`           VARCHAR(30) NOT NULL,

    PRIMARY KEY (`id`)

) ENGINE = InnoDB
  DEFAULT CHARACTER SET = UTF8MB4
  COLLATE = utf8mb4_unicode_520_ci;

/*CREATE TABLE `russian_doll_parent` (
 `russian_doll_id` varchar(36) NOT NULL UNIQUE,
 `parent_id` varchar(36) NOT NULL,
 `date_created` datetime NOT NULL,
 PRIMARY KEY (`russian_doll_id`,`parent_id`),
 CONSTRAINT `fk_rp_russian_doll_id` FOREIGN KEY (`russian_doll_id`) REFERENCES `russial_doll` (`id`) ON DELETE CASCADE
) ENGINE = InnoDB
DEFAULT CHARACTER SET = UTF8MB4
COLLATE = utf8mb4_unicode_520_ci;*/