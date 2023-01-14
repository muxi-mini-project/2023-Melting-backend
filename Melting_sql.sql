##用户信息
CREATE TABLE `Melting_User_info`(
    `id` INT(9) AUTO_INCREMENT COMMENT '序号'
    `nick_name` VARCHAR(14) DEFAULT '昵称' COMMENT '最多七个汉字'
    `position` VARCHAR (30) DEFAULT '职位' COMMENT '职位'
    PRIMARY KEY (`id`)
)

##活动策划基础信息
CREATE TABLE `My_Melting_proposal_fundermentalinfo`(
    `user_id` INT(9) AUTO_INCREMENT
    `proposal_id` INT(5) AUTO_INCREMENT COMMENT '活动序号'
    `place` VARCHAR(60) DEFAULT '活动位置' COMMENT '活动位置'
    `time` DATETIME DEFAULT NULL COMMENT '活动时间'
    `flock` INT DEFAULT 0 COMMENT '参与人数'
    `department_name` VARCHAR(30) DEFAULT '部门名称' COMMENT '部门'
    `proposal_name` VARCHAR(40) DEFAULT '活动名称' COMMENT '活动'
    `proposal_aim` VARCHAR(300) DEFAULT '活动目的' COMMENT '活动目的'
    `inform_txt` VARCHAR(600) DEFAULT '活动通知' COMMENT '活动通知文本'
    `target` VARCHAR DEFAULT '00000000' COMMENT '活动预算'
    `rules_expression` VARCHAR(400) DEFAULT '规则说明' COMMENT '规则说明' 
    `advice&basis` VARCHAR(400) DEFAULT '建议及其根据' COMMENT '建议及其根据' 
    `optional_tables` VARCHAR(20) DEFAULT '可选标签' COMMENT '可选标签'
    PRIMARY KEY (`proposal_id`)
    CONSTRAINT user_proposal_foreignkey FOREIGN KEY (user_id) REFERENCES Melting_User_info (id)
)

##活动安排信息
CREATE TABLE `proposal_arrangement`(
    `arrangement_id` INT(9) AUTO_INCREMENT
    `hintword` VARCHAR(300) DEFAULT '提示语' COMMENT '提示语'
    `game` VARCHAR(400) DEFAULT '游戏项目' COMMENT '游戏项目'
    `proposal_nodes` VARCHAR(400) DEFAULT '项目环节' COMMENT '项目环节'
PRIMARY KEY (`arrangement_id`)
CONSTRAINT proposal_arrangement_foreignkey FOREIGN KEY (arrangement_id) REFERENCES My_Melting_proposal_fundermentalinfo (proposal_id)
)

##游戏对应题库
CREATE TABLE `question_warehouse` (
    `warehouse_id` INT(9) AUTO_INCREMENT
    `questions` VARCHAR 
PRIMARY KEY (`warehouse_id`)
CONSTRAINT game_question_foreignkey FOREIGN KEY (warehouse_id) REFERENCES proposal_arrangement (game)
)