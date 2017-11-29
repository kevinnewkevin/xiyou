/*
Navicat MySQL Data Transfer

Source Server         : 254
Source Server Version : 50631
Source Host           : 10.10.10.254:3306
Source Database       : xygame

Target Server Type    : MYSQL
Target Server Version : 50631
File Encoding         : 65001

Date: 2017-11-29 12:07:29
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for `BattleReport`
-- ----------------------------
DROP TABLE IF EXISTS `BattleReport`;
CREATE TABLE `BattleReport` (
  `BattleID` int(64) NOT NULL,
  `Report` blob NOT NULL,
  PRIMARY KEY (`BattleID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

