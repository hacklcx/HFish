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
  `last_update_time` datetime NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `un_agent` (`agent_name`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;

SET FOREIGN_KEY_CHECKS = 1;
