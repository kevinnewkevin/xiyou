/*
Navicat MySQL Data Transfer

Source Server         : 1
Source Server Version : 50631
Source Host           : 10.10.10.254:3306
Source Database       : xygame

Target Server Type    : MYSQL
Target Server Version : 50631
File Encoding         : 65001

Date: 2017-11-24 18:30:29
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for `Player`
-- ----------------------------
DROP TABLE IF EXISTS `Player`;
CREATE TABLE `Player` (
  `PlayerId` int(64) NOT NULL AUTO_INCREMENT,
  `Username` varchar(255) COLLATE utf8_bin NOT NULL,
  `BinData` blob NOT NULL,
  `InstId` int(64) NOT NULL,
  PRIMARY KEY (`PlayerId`)
) ENGINE=InnoDB AUTO_INCREMENT=253 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

DROP TABLE IF EXISTS `Unit`;
CREATE TABLE `Unit` (
  `UnitId` int(11) NOT NULL,
  `OwnerId` int(11) NOT NULL,
  `BinData` blob NOT NULL,
  PRIMARY KEY (`UnitId`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

DROP TABLE IF EXISTS `TopList`;
CREATE TABLE `TopList` (
  `InstId` int(64) NOT NULL,
  `BinData` blob NOT NULL,
  PRIMARY KEY (`InstId`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin ROW_FORMAT=COMPACT;
