/*
SQLyog Ultimate
MySQL - 5.0.96-community-log : Database - fasp
*********************************************************************
*/

CREATE DATABASE `service_config` ;

/*Table structure for table `userinfo` */

DROP TABLE IF EXISTS `userinfo`;

CREATE TABLE `userinfo` (
  `uid` int(10) NOT NULL auto_increment,
  `username` varchar(64) default NULL,
  `department` varchar(64) default NULL,
  `created` date default NULL,
  PRIMARY KEY  (`uid`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

