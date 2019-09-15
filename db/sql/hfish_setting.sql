-- ----------------------------
--  Table structure for `hfish_setting`
-- ----------------------------
DROP TABLE IF EXISTS `hfish_setting`;
CREATE TABLE `hfish_setting` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `type` varchar(50) NOT NULL DEFAULT '',
  `info` varchar(50) NOT NULL DEFAULT '',
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
INSERT INTO `hfish_setting` VALUES ('1', 'mail', '', '2019-09-02 20:15:00', '0', 'E-mail 群发', '群发邮件SMTP服务器配置'), ('2', 'alertMail', '', '2019-09-02 18:58:12', '0', 'E-mail 通知', '蜜罐告警会通过邮件告知信息'), ('3', 'webHook', '', '2019-09-03 11:49:00', '0', 'WebHook 通知', '蜜罐告警会请求指定API告知信息'), ('4', 'whiteIp', '', '2019-09-02 20:15:00', '0', 'IP 白名单', '蜜罐上钩会过滤掉白名单IP');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
