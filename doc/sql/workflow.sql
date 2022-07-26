/*
Navicat MySQL Data Transfer

Source Server         : gomysql
Source Server Version : 80029
Source Host           : localhost:3306
Source Database       : workflow

Target Server Type    : MYSQL
Target Server Version : 80029
File Encoding         : 65001

Date: 2022-07-26 21:55:30
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for act_execution
-- ----------------------------
DROP TABLE IF EXISTS `act_execution`;
CREATE TABLE `act_execution` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `rev` int DEFAULT NULL,
  `proc_inst_id` bigint DEFAULT NULL COMMENT '流程实例ID',
  `proc_def_id` bigint DEFAULT NULL COMMENT '流程实例ID',
  `node_infos` varchar(5000) CHARACTER SET utf8mb3 COLLATE utf8_general_ci DEFAULT NULL COMMENT '节点审批顺序',
  `is_active` tinyint DEFAULT NULL COMMENT '是否审批中：0:审批中，1：审批完成',
  `start_time` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8_general_ci DEFAULT NULL COMMENT '发起时间',
  PRIMARY KEY (`id`),
  KEY `execution_proc_inst_id_proc_inst_id_foreign` (`proc_inst_id`) USING BTREE,
  KEY `idx_id` (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of act_execution
-- ----------------------------

-- ----------------------------
-- Table structure for act_identity_link
-- ----------------------------
DROP TABLE IF EXISTS `act_identity_link`;
CREATE TABLE `act_identity_link` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `type` varchar(20) CHARACTER SET utf8mb3 COLLATE utf8_general_ci DEFAULT NULL COMMENT '人员类型',
  `user_id` bigint DEFAULT NULL COMMENT '审批人ID',
  `user_name` varchar(20) CHARACTER SET utf8mb3 COLLATE utf8_general_ci DEFAULT NULL COMMENT '审批人姓名',
  `step` int DEFAULT NULL COMMENT '审批步数',
  `proc_inst_id` bigint DEFAULT NULL COMMENT '流程实例ID',
  `target_id` bigint DEFAULT NULL COMMENT '岗位ID',
  `comment` varchar(500) CHARACTER SET utf8mb3 COLLATE utf8_general_ci DEFAULT NULL COMMENT '评论',
  `result` tinyint DEFAULT NULL COMMENT '审批结果:3驳回至发起人、4驳回到上一级、6未通过、7已通过',
  PRIMARY KEY (`id`),
  KEY `identitylink_proc_inst_id_proc_inst_id_foreign` (`proc_inst_id`) USING BTREE,
  KEY `idx_id` (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of act_identity_link
-- ----------------------------

-- ----------------------------
-- Table structure for act_proc_def
-- ----------------------------
DROP TABLE IF EXISTS `act_proc_def`;
CREATE TABLE `act_proc_def` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `name` varchar(200) CHARACTER SET utf8mb3 COLLATE utf8_general_ci NOT NULL COMMENT '流程名称',
  `code` varchar(200) CHARACTER SET utf8mb3 COLLATE utf8_general_ci DEFAULT NULL COMMENT '流程编码',
  `version` int DEFAULT NULL COMMENT '版本',
  `resource` varchar(10000) CHARACTER SET utf8mb3 COLLATE utf8_general_ci DEFAULT NULL COMMENT '流程图数据',
  `create_user_id` bigint DEFAULT NULL COMMENT '创建人ID',
  `create_user_name` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8_general_ci DEFAULT NULL COMMENT '创建人名称',
  `create_time` datetime DEFAULT NULL,
  `target_id` bigint DEFAULT NULL COMMENT '组织ID',
  `yewu_form_id` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8_general_ci DEFAULT NULL COMMENT '业务表单ID',
  `yewu_name` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8_general_ci DEFAULT NULL COMMENT '业务表单名称',
  `remain_date` int DEFAULT NULL COMMENT '审批限定时间',
  `remain_unit` varchar(10) CHARACTER SET utf8mb3 COLLATE utf8_general_ci DEFAULT NULL COMMENT '审批限定时间单位:下拉选：时/天',
  `is_del` tinyint DEFAULT NULL COMMENT '是否删除:0:未删除，1：已删除',
  PRIMARY KEY (`id`),
  KEY `idx_id` (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=8754705275782332417 DEFAULT CHARSET=utf8mb3 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of act_proc_def
-- ----------------------------

-- ----------------------------
-- Table structure for act_proc_inst
-- ----------------------------
DROP TABLE IF EXISTS `act_proc_inst`;
CREATE TABLE `act_proc_inst` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `proc_def_id` bigint NOT NULL COMMENT '流程定义ID',
  `proc_def_name` varchar(200) CHARACTER SET utf8mb3 COLLATE utf8_general_ci DEFAULT NULL COMMENT '流程定义名称',
  `title` varchar(200) CHARACTER SET utf8mb3 COLLATE utf8_general_ci DEFAULT NULL COMMENT '发起流程标题',
  `target_id` bigint DEFAULT NULL COMMENT '组织ID',
  `node_id` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8_general_ci DEFAULT NULL COMMENT '节点ID',
  `task_id` bigint DEFAULT NULL COMMENT '当前审批任务ID',
  `start_time` datetime(6) DEFAULT NULL COMMENT '发起时间',
  `end_time` datetime(6) DEFAULT NULL COMMENT '结束时间',
  `start_user_id` varchar(0) CHARACTER SET utf8mb3 COLLATE utf8_general_ci DEFAULT NULL COMMENT '发起人ID',
  `start_user_name` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8_general_ci DEFAULT NULL COMMENT '发起人姓名',
  `is_finished` tinyint(1) DEFAULT '0' COMMENT '流程是否结束:0:未结束，1：已结束',
  `state` tinyint(1) DEFAULT NULL COMMENT '流程状态:"类型为:1待处理、2处理中、3驳回至发起人、4驳回到上一级、\n5已撤回、6未通过、7已通过、8废弃"',
  `data_id` bigint DEFAULT NULL COMMENT '流程绑定数据ID',
  `is_del` tinyint(1) DEFAULT NULL COMMENT '是否删除:0:未删除，1：已删除',
  PRIMARY KEY (`id`),
  KEY `idx_id` (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of act_proc_inst
-- ----------------------------

-- ----------------------------
-- Table structure for act_task
-- ----------------------------
DROP TABLE IF EXISTS `act_task`;
CREATE TABLE `act_task` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `node_id` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8_general_ci NOT NULL COMMENT '节点ID',
  `step` int NOT NULL COMMENT '流程步数:发起人step=0',
  `proc_inst_id` bigint NOT NULL COMMENT '流程实例ID',
  `create_time` datetime(6) DEFAULT NULL COMMENT '任务创建时间',
  `claim_time` datetime(6) DEFAULT NULL COMMENT '节点最新审批时间',
  `member_count` int DEFAULT '1' COMMENT '需审批人数:表示当前任务需要多少人审批之后才能结束，默认是 1',
  `un_complete_num` int DEFAULT '1' COMMENT '未审批人数',
  `agree_num` int DEFAULT NULL COMMENT '已通过人数',
  `is_finished` tinyint(1) DEFAULT '0' COMMENT '任务是否完成:0：未结束 1：已完成',
  PRIMARY KEY (`id`),
  KEY `task_proc_inst_id_proc_inst_id_foreign` (`proc_inst_id`) USING BTREE,
  KEY `idx_id` (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of act_task
-- ----------------------------

-- ----------------------------
-- Table structure for act_web_hook
-- ----------------------------
DROP TABLE IF EXISTS `act_web_hook`;
CREATE TABLE `act_web_hook` (
  `id` bigint NOT NULL,
  `flow_id` bigint NOT NULL COMMENT '流程实例ID',
  `flow_state` tinyint DEFAULT NULL COMMENT '流程实例状态',
  `hook_addr` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8_general_ci DEFAULT NULL COMMENT 'hook地址',
  `rm` varchar(10) CHARACTER SET utf8mb3 COLLATE utf8_general_ci DEFAULT NULL COMMENT '请求类型:GET、POST、PUT、DELETE请求',
  `create_time` datetime(6) DEFAULT NULL,
  `update_time` datetime(6) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of act_web_hook
-- ----------------------------
