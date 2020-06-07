USE Graph;

CREATE TABLE IF NOT EXISTS `vertex`
(
    `id`             VARCHAR(36) NOT NULL,
    `type`           VARCHAR(30) NOT NULL,

    PRIMARY KEY (`id`)

) ENGINE = InnoDB
  DEFAULT CHARACTER SET = UTF8MB4
  COLLATE = utf8mb4_unicode_520_ci;

CREATE TABLE `edge` (
 `vertex_id` varchar(36) NOT NULL UNIQUE,
 `parent_id` varchar(36) NOT NULL,
 PRIMARY KEY (`vertex_id`,`parent_id`),
 CONSTRAINT `fk_rp_vertex_id` FOREIGN KEY (`vertex_id`) REFERENCES `vertex` (`id`) ON DELETE CASCADE,
 CONSTRAINT `fk_rp_parent_id` FOREIGN KEY (`parent_id`) REFERENCES `vertex` (`id`) ON DELETE CASCADE
) ENGINE = InnoDB
DEFAULT CHARACTER SET = UTF8MB4
COLLATE = utf8mb4_unicode_520_ci;