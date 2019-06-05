/*
Navicat MySQL Data Transfer

Source Server         : lc
Source Server Version : 50562
Source Host           : localhost:3306
Source Database       : ljx

Target Server Type    : MYSQL
Target Server Version : 50562
File Encoding         : 65001

Date: 2019-01-11 16:48:58
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for `wauth`
-- ----------------------------
DROP TABLE IF EXISTS `wauth`;
CREATE TABLE `wauth` (
  `auth_id` int(11) NOT NULL AUTO_INCREMENT,
  `auth_name` varchar(256) DEFAULT NULL,
  `auth_val` text,
  `remark` varchar(256) DEFAULT NULL,
  `pid` int(11) DEFAULT NULL,
  `ordered` int(11) DEFAULT NULL,
  PRIMARY KEY (`auth_id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of wauth
-- ----------------------------
INSERT INTO `wauth` VALUES ('1', '系统管理', '', null, '0', '0');
INSERT INTO `wauth` VALUES ('2', '用户管理', '/userlist', null, '1', '1');

-- ----------------------------
-- Table structure for `wrole`
-- ----------------------------
DROP TABLE IF EXISTS `wrole`;
CREATE TABLE `wrole` (
  `role_id` int(11) NOT NULL AUTO_INCREMENT,
  `role_name` varchar(256) DEFAULT NULL,
  `auths` text,
  `remark` varchar(256) DEFAULT NULL,
  PRIMARY KEY (`role_id`)
) ENGINE=InnoDB AUTO_INCREMENT=102 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of wrole
-- ----------------------------
INSERT INTO `wrole` VALUES ('100', '管理员', '1,2,3,4,5', null);
INSERT INTO `wrole` VALUES ('101', '普通用户', null, null);

-- ----------------------------
-- Table structure for `wuser`
-- ----------------------------
DROP TABLE IF EXISTS `wuser`;
CREATE TABLE `wuser` (
  `username` varchar(64) NOT NULL,
  `pwd` varchar(256) DEFAULT NULL,
  `nick_name` varchar(256) DEFAULT NULL,
  `real_name` varchar(256) DEFAULT NULL,
  `role_id` int(11) DEFAULT NULL,
  `remark` varchar(256) DEFAULT NULL,
  PRIMARY KEY (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of wuser
-- ----------------------------
INSERT INTO `wuser` VALUES ('admin', '0DPiKuNIrrVmD8IUCuw1hQxNqZc=', 'cxx', 'cxx', '100', null);
