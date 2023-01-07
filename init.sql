## modified 1/7/2022
##用户信息
##由于开发主力过于迅猛，被迫作为监工的我在此留下标记。
## ovo

CREATE TABLE if not exists`Users`(
                                     `uid` INT AUTO_INCREMENT COMMENT '序号',
                                     `photo` VARCHAR(200) COMMENT '头像',
                                     `nick_name` VARCHAR(14) DEFAULT '昵称' COMMENT '最多七个汉字',
                                     `position` VARCHAR (30) DEFAULT '职位' COMMENT '职位',
                                     `tag` VARCHAR(100) COMMENT '标签',
                                     `auth` varchar(100),
                                     `muxipass` int ,
                                     `studentid` int,
                                     `qq` int,
                                     `description` varchar(100) not null,
                                     PRIMARY KEY (`uid`)
);

##活动策划基础信息
CREATE TABLE if not exists `proposal_info` (
                                               `info_id` INT AUTO_INCREMENT COMMENT '活动序号',
                                               `place` VARCHAR(60) DEFAULT '活动位置' COMMENT '活动位置',
                                               `time` DATETIME DEFAULT NULL COMMENT '活动时间',
                                               `department` VARCHAR(30) DEFAULT '部门名称' COMMENT '部门',
                                               `name` VARCHAR(40) DEFAULT '活动名称' COMMENT '活动',
                                               `aim` VARCHAR(300) DEFAULT '活动目的' COMMENT '活动目的',
                                               `budget` json COMMENT '活动预算',## 以json list存储
                                               `optional_tables` VARCHAR(20) DEFAULT '可选标签' COMMENT '可选标签',
                                                `uid` int,
                                                `game` VARCHAR(400) DEFAULT '游戏项目' COMMENT '游戏项目',
                                                `nodes` json COMMENT '项目环节',
                                                PRIMARY KEY (`info_id`)

);

##游戏对应题库
CREATE TABLE if not exists`questions` (
                                          `qid` INT AUTO_INCREMENT,
                                          `questions` VARCHAR(100) ,
                                          `type` varchar(100),
                                          PRIMARY KEY (`qid`)
);

##标签
CREATE TABLE if not exists`tags`(
                                    `tid` int auto_increment,
                                    `name` varchar(10) not null,
                                    `description` varchar(100) not null,
                                    primary key(`tid`)
);

##模板
CREATE TABLE if not exists`template`(
                                    `temid`int auto_increment,
                                    `name`varchar(60) not null,
                                    `context`varchar(400),
                                    primary key(`temid`)
);

