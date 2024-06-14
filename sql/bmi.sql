CREATE DATABASE IF NOT EXISTS `dev`
/*!40100 DEFAULT CHARACTER SET utf8 COLLATE utf8_unicode_ci */
;

USE `dev`;

CREATE TABLE `bmi` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `name` varchar(255) NOT NULL,
    `weight` varchar(255) NOT NULL,
    `height` varchar(255) NOT NULL,
    `bmi` varchar(255) NOT NULL,
    `updated_at` datetime DEFAULT NULL,
    `created_at` datetime DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB AUTO_INCREMENT = 7 DEFAULT CHARSET = utf8 COLLATE = utf8_unicode_ci;