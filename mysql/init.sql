CREATE DATABASE suscan default charset utf8 COLLATE utf8_general_ci;
use suscan;

-- ----------------------------
-- Table structure for assets
-- ----------------------------
DROP TABLE IF EXISTS `assets`;
CREATE TABLE `assets` (
  `id` int NOT NULL AUTO_INCREMENT,
  `target` varchar(255) DEFAULT NULL,
  `created_time` varchar(255) DEFAULT NULL,
  `updated_time` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of assets
-- ----------------------------

-- ----------------------------
-- Table structure for iplist
-- ----------------------------
DROP TABLE IF EXISTS `iplist`;
CREATE TABLE `iplist` (
  `id` int NOT NULL AUTO_INCREMENT,
  `url` varchar(255) DEFAULT NULL,
  `ip` varchar(255) DEFAULT NULL,
  `port` varchar(255) DEFAULT NULL,
  `state` varchar(255) DEFAULT NULL,
  `protocol` varchar(255) DEFAULT NULL,
  `service` varchar(255) DEFAULT NULL,
  `created_time` varchar(255) DEFAULT NULL,
  `updated_time` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of iplist
-- ----------------------------


-- ----------------------------
-- Table structure for log
-- ----------------------------
DROP TABLE IF EXISTS `log`;
CREATE TABLE `log` (
  `id` int NOT NULL AUTO_INCREMENT,
  `task_name` varchar(255) DEFAULT NULL,
  `task_type` varchar(255) DEFAULT NULL,
  `all_num` int DEFAULT NULL,
  `run_time` varchar(255) DEFAULT NULL,
  `created_time` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


-- ----------------------------
-- Table structure for setting
-- ----------------------------
DROP TABLE IF EXISTS `setting`;
CREATE TABLE `setting` (
  `id` int NOT NULL AUTO_INCREMENT,
  `thread` varchar(255) DEFAULT NULL,
  `port` varchar(255) DEFAULT NULL,
  `cmd` varchar(255) DEFAULT NULL,
  `created_time` varchar(255) DEFAULT NULL,
  `updated_time` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of setting
-- ----------------------------

INSERT INTO `setting` VALUES (1, '2000', '80,443', '3', '20210819163924', '20210819164714');

