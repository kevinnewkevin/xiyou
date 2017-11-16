/*
Navicat MySQL Data Transfer

Source Server         : 254
Source Server Version : 1
Source Host           : 10.10.10.254:3306
Source Database       : GameXY

Target Server Type    : MYSQL
Target Server Version : 1
File Encoding         : 1

Date: 2017-11-14 16:10:28
*/

SET FOREIGN_KEY_CHECKS=0;

/** 
 * @brief 帮派表
 * */
DROP TABLE IF EXISTS `Guild`;	
CREATE TABLE `Guild`
(
	`GuildId`	    INT 		NOT NULL,/*** 帮派ID*/
	`GuildName`   	VARCHAR(60)	NOT NULL,/*** 帮派名*/
	`Master`		BIGINT 		NOT NULL,/*** 帮主*/
	`MasterName`  	VARCHAR(60)	NOT NULL,/*** 帮主名(未更新)*/
	`GuildVal`  	INT			NOT NULL,/*** 帮派积分*/
	`CreatTime` 	BIGINT 		NOT NULL,/*** 创建时间*/
	`RequestList` 	BLOB 		NOT NULL,
	`RequestFlag`	INT			NOT NULL,/*** 是否需要申请*/
	`Require`		INT			NOT NULL,/*** 申请条件*/	
	PRIMARY KEY(guildId),
	UNIQUE KEY guildNameIdx   (guildName)
)ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
/**
 * @brief 帮派成员表
 * */
 DROP TABLE IF EXISTS `GuildMember`;
CREATE TABLE `GuildMember`
(
	`GuildId`		INT 		NOT NULL,/*** 帮派ID(外键 帮派表)*/
	`RoleId`		BIGINT 		NOT NULL,/*** 成员人物ID*/
	`RoleName`   	VARCHAR(60)	NOT NULL,/*** 成员名*/
	`Rolelevel`     TINYINT 	NOT NULL,/*** 成员等级*/
	`Job`			TINYINT 	NOT NULL,/*** 职位*/
	`Contribution`	INT 		NOT NULL,/*** 捐献*/
	`TianTiVal`		INT 		NOT NULL,/*** 天梯积分*/
	PRIMARY KEY(roleId)
)ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
/**
 * @brief 捐赠表
 * */
 DROP TABLE IF EXISTS `GuildAssistant`;
CREATE TABLE `GuildAssistant`
(
	`AssistantId`		INT 		NOT NULL,
	`RoleName`   		VARCHAR(60)	NOT NULL,/*** 成员名*/
	`GuildId`			INT 		NOT NULL,/*** 帮派ID(外键 帮派表)*/
	`AssistantItem`		INT 		NOT NULL,/*** 捐献道具*/
	`CrtCount`			INT 		NOT NULL,/*** 当前捐献数量*/
	`MaxCount`			INT 		NOT NULL,/*** 当前捐献数量*/
	`CatchNum`			INT 		NOT NULL,/*** 角色不在线收到的捐赠数量*/
	`Donator` 			BLOB 		NOT NULL,/*** 捐赠者列表*/
	PRIMARY KEY(AssistantId)
)ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;


