SET NAMES utf8;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
--  Table structure for `hfish_colony`
-- ----------------------------
DROP TABLE IF EXISTS `hfish_colony`;
CREATE TABLE `hfish_colony` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `agent_name` varchar(20) NOT NULL DEFAULT '',
  `agent_ip` varchar(20) NOT NULL DEFAULT '',
  `web_status` int(2) NOT NULL DEFAULT '0',
  `deep_status` int(2) NOT NULL DEFAULT '0',
  `ssh_status` int(2) NOT NULL DEFAULT '0',
  `redis_status` int(2) NOT NULL DEFAULT '0',
  `mysql_status` int(2) NOT NULL DEFAULT '0',
  `http_status` int(2) NOT NULL DEFAULT '0',
  `telnet_status` int(2) NOT NULL DEFAULT '0',
  `ftp_status` int(2) NOT NULL DEFAULT '0',
  `mem_cache_status` int(2) NOT NULL DEFAULT '0',
  `plug_status` int(2) NOT NULL DEFAULT '0',
  `es_status` int(2) NOT NULL DEFAULT '0',
  `tftp_status` int(2) NOT NULL DEFAULT '0',
  `vnc_status` int(2) NOT NULL DEFAULT '0',
  `custom_status` int(2) NOT NULL DEFAULT '0',
  `last_update_time` datetime NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `un_agent` (`agent_name`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
--  Table structure for `hfish_info`
-- ----------------------------
DROP TABLE IF EXISTS `hfish_info`;
CREATE TABLE `hfish_info` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `type` varchar(20) NOT NULL DEFAULT '',
  `project_name` varchar(20) NOT NULL DEFAULT '',
  `agent` varchar(20) NOT NULL DEFAULT '',
  `ip` varchar(20) NOT NULL DEFAULT '',
  `country` varchar(10) NOT NULL DEFAULT '',
  `region` varchar(10) NOT NULL DEFAULT '',
  `city` varchar(10) NOT NULL,
  `info` text NOT NULL,
  `create_time` datetime NOT NULL,
  PRIMARY KEY (`id`),
  KEY `info_index_1` (`type`) USING BTREE,
  KEY `info_index_2` (`country`) USING BTREE,
  KEY `info_index_3` (`type`,`create_time`) USING BTREE,
  KEY `info_index_4` (`ip`) USING BTREE,
  KEY `info_index_5` (`agent`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
--  Table structure for `hfish_passwd`
-- ----------------------------
DROP TABLE IF EXISTS `hfish_passwd`;
CREATE TABLE `hfish_passwd` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `type` varchar(50) NOT NULL DEFAULT '',
  `account` varchar(50) NOT NULL DEFAULT '',
  `password` varchar(50) NOT NULL DEFAULT '',
  `create_time` datetime NOT NULL,
  PRIMARY KEY (`id`),
  KEY `index_key_1` (`account`) USING BTREE,
  KEY `index_key_2` (`password`) USING BTREE,
  KEY `index_key_3` (`type`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
--  Table structure for `hfish_setting`
-- ----------------------------
DROP TABLE IF EXISTS `hfish_setting`;
CREATE TABLE `hfish_setting` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `type` varchar(50) NOT NULL DEFAULT '',
  `info` text,
  `update_time` datetime NOT NULL,
  `status` int(2) NOT NULL DEFAULT '0',
  `setting_name` varchar(50) NOT NULL DEFAULT '',
  `setting_dis` varchar(50) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  UNIQUE KEY `index_key` (`type`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
--  Records of `hfish_setting`
-- ----------------------------
BEGIN;
INSERT INTO `hfish_setting` VALUES ('1', 'mail', '', '2019-09-02 20:15:00', '0', 'E-mail 群发', '群发邮件SMTP服务器配置'), ('2', 'alertMail', '', '2019-09-02 18:58:12', '0', 'E-mail 通知', '蜜罐告警会通过邮件告知信息'), ('3', 'webHook', '', '2019-09-03 11:49:00', '0', 'WebHook 通知', '蜜罐告警会请求指定API告知信息'), ('4', 'whiteIp', '', '2019-09-02 20:15:00', '0', 'IP 白名单', '蜜罐上钩会过滤掉白名单IP'), ('5', 'passwdTM', '', '2020-02-24 12:04', '0', '密码脱敏', '发送邮件如果有密码内容自动加星');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
