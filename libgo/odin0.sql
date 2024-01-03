/*
CREATE TABLE IF NOT EXISTS  `tablename`(
  `created_at` timestamp NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NULL DEFAULT NULL
)ENGINE=InnoDB DEFAULT CHARSET=latin1;
*/
DROP DATABASE `odin`;
CREATE DATABASE `odin`;
USE `odin`;
CREATE TABLE IF NOT EXISTS `user` (
  `userid` varchar(255) NOT NULL,
  `ownerid` varchar(255) NOT NULL,
  `username` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `active` BOOLEAN NOT NULL,
  `anonymous` BOOLEAN NOT NULL,
  `verified` BOOLEAN NOT NULL,
  `admin` BOOLEAN NOT NULL,
  `created_at` timestamp NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`userid`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE IF NOT EXISTS `apikey` (
  `apikey` varchar(255) NOT NULL,
  `ownerid` varchar(255) NOT NULL,
  `active` BOOLEAN NOT NULL,
  `created_at` timestamp NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`apikey`),
  FOREIGN KEY (`ownerid`) REFERENCES `user` (`userid`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE IF NOT EXISTS `apt` (
  `aptname` varchar(255) NOT NULL,
  `code` int(255) NOT NULL,
  `id` varchar(255) NOT NULL,
  `description` TEXT,
  `active` BOOLEAN NOT NULL,
  `created_at` timestamp NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE IF NOT EXISTS `appointments` (
  `userid` varchar(255) NOT NULL,
  `appointmentid` varchar(25) NOT NULL,
  `title` varchar(25) NOT NULL,
  `description` varchar(255) NOT NULL,
  `done` BOOLEAN NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`appointmentid`),
  FOREIGN KEY (`userid`) REFERENCES `user` (`userid`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE IF NOT EXISTS `virus` (
  `virusid` varchar(255) NOT NULL,
  `hash` JSON,
  `aptid` varchar(255) NOT NULL,
  `virustype` varchar(255) NOT NULL,
  `filetype` varchar(255) NOT NULL,
  `communicationmode` varchar(255) NOT NULL,
  `ostype` varchar(255) NOT NULL,
  `description` TEXT NOT NULL,
  `created_at` timestamp NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`virusid`),
  FOREIGN KEY (`aptid`) REFERENCES `apt` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE IF NOT EXISTS `threats` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `threat_ip` JSON,
  `virusid` varchar(255) NOT NULL,
  `details` varchar(255) NOT NULL,
  `active` BOOLEAN NOT NULL,
  `created_at` timestamp NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NULL DEFAULT NULL,
   PRIMARY KEY (`id`),
  FOREIGN KEY (`virusid`) REFERENCES `virus` (`virusid`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE IF NOT EXISTS `threatactivity` (
  `name` varchar(255) NOT NULL,
  `activity_id` varchar(255) NOT NULL,
  `creatorid` varchar(255) NOT NULL,
  `description` varchar(255) NOT NULL,
  `validated` BOOLEAN NOT NULL,
  `created_at` timestamp NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`activity_id`),
  FOREIGN KEY (`creatorid`) REFERENCES `user` (`userid`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE IF NOT EXISTS `motherships` (
  `ownerid` varchar(255) NOT NULL,
  `name` varchar(255) NOT NULL,
  `msid` varchar(255) NOT NULL,
  `ipaddr` INT UNSIGNED NOT NULL,
  `port` INT(255) NOT NULL,
  `tunnel_address` varchar(255) NOT NULL,
  `description` TEXT NOT NULL,
  `active` BOOLEAN NOT NULL,
  `created_at` timestamp NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`msid`),
  FOREIGN KEY (`ownerid`) REFERENCES `user` (`userid`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE IF NOT EXISTS `minion` (
  `minionid` varchar(255) NOT NULL,
  `name` varchar(255) NOT NULL,
  `uname` varchar(255) NOT NULL,
  `userid` varchar(255) NOT NULL,
  `groupid` varchar(255) NOT NULL,
  `homedir` varchar(255) NOT NULL,
  `ostype` varchar(255) NOT NULL,
  `description` TEXT NOT NULL,
  `installed` BOOLEAN NOT NULL,
  `mothershipid` varchar(255) NOT NULL,
  `minionip` INT UNSIGNED NOT NULL,
  `ownerid` varchar(255) NOT NULL,
  `lastseen` varchar(255) NOT NULL,
  `created_at` timestamp NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`minionid`),
  FOREIGN KEY (`userid`) REFERENCES `user` (`userid`),
  FOREIGN KEY (`mothershipid`) REFERENCES `motherships` (`msid`),
  FOREIGN KEY (`ownerid`) REFERENCES `user` (`userid`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE IF NOT EXISTS `hashes` (
  `userid` varchar(255) NOT NULL,
  `hash` varchar(255) NOT NULL,
  `created_at` timestamp NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`userid`),
  FOREIGN KEY (`userid`) REFERENCES `user` (`userid`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE IF NOT EXISTS `assets` (
  `asset_id` VARCHAR(255) NOT NULL,
  `name` VARCHAR(255) NOT NULL,
  `agentid` VARCHAR(255) NOT NULL,
  `description` TEXT NOT NULL,
  `active` BOOLEAN NOT NULL,
  `created_at` timestamp NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`asset_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE IF NOT EXISTS `plugins` (
  `owner` varchar(255) NOT NULL,
  `name` varchar(255) NOT NULL,
  `hash` varchar(255) NOT NULL,
  `plugin_type` INT(255) NOT NULL,
  `description` TEXT NOT NULL,
  `active` BOOLEAN NOT NULL,
  `signed` BOOLEAN NOT NULL,
  `created_at` timestamp NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`hash`),
  FOREIGN KEY (`owner`) REFERENCES `user` (`userid`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE IF NOT EXISTS `ioc` (
  `ioc_id` INT UNSIGNED NOT NULL,
  `virusid`  varchar(255) NOT NULL,
  `type` VARCHAR(255) NOT NULL,
  `value` TEXT NOT NULL,
  `source` VARCHAR(255),
  `created_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP NULL DEFAULT NULL,
  PRIMARY KEY (`ioc_id`),
  CONSTRAINT `fk_virus`
    FOREIGN KEY (`virusid`) REFERENCES `virus` (`virusid`)
    ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=latin1;


CREATE TABLE IF NOT EXISTS `yara_rule` (
  `yr_id` INT UNSIGNED NOT NULL,
  `virus_id` VARCHAR(255) NOT NULL,
  `ioc_id` INT UNSIGNED,
  `name` VARCHAR(255) NOT NULL,
  `meta` JSON,
  `condition` TEXT NOT NULL,
  `actions` JSON,
  `created_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP NULL DEFAULT NULL,
  PRIMARY KEY (`yr_id`),
  FOREIGN KEY (`virus_id`) REFERENCES `virus` (`virusid`),
  FOREIGN KEY (`ioc_id`) REFERENCES `ioc` (`ioc_id`)
    ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=latin1;


CREATE TABLE IF NOT EXISTS events (
  id INT AUTO_INCREMENT PRIMARY KEY,
  EId VARCHAR(255) NOT NULL,
  OS INT NOT NULL,
  Handled BOOLEAN NOT NULL,
  Level INT NOT NULL,
  `created_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP NULL DEFAULT NULL,
);
