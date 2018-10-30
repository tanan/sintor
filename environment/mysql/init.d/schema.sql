CREATE DATABASE `sintor`;

CREATE TABLE `sintor`.`schedule` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(64) NOT NULL,
  `interval` INT NOT NULL,
  `monitor_id` INT NOT NULL,
  PRIMARY KEY (`id`));

CREATE TABLE `sintor`.`monitor` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(64) NOT NULL,
  `status` VARCHAR(5) NOT NULL,
  `url` VARCHAR(1024) NOT NULL,
  `method` VARCHAR(6) NOT NULL DEFAULT 'GET',
  `body` VARCHAR(2048) NULL,
  PRIMARY KEY (`id`);
)