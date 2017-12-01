/*
Navicat MySQL Data Transfer

Source Server         : 254
Source Server Version : 50631
Source Host           : 10.10.10.254:3306
Source Database       : xygame

Target Server Type    : MYSQL
Target Server Version : 50631
File Encoding         : 65001

Date: 2017-12-01 16:10:59
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for `CheckPointBattleRecord`
-- ----------------------------
DROP TABLE IF EXISTS `CheckPointBattleRecord`;
CREATE TABLE `CheckPointBattleRecord` (
  `CheckPointId` int(32) NOT NULL,
  `Data` blob NOT NULL,
  PRIMARY KEY (`CheckPointId`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ----------------------------
-- Records of CheckPointBattleRecord
-- ----------------------------
