/*
 Navicat Premium Data Transfer

 Source Server         : localhost_3306
 Source Server Type    : MySQL
 Source Server Version : 50528
 Source Host           : localhost:3306
 Source Schema         : gproject

 Target Server Type    : MySQL
 Target Server Version : 50528
 File Encoding         : 65001

 Date: 02/05/2020 10:28:23
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for details
-- ----------------------------
DROP TABLE IF EXISTS `details`;
CREATE TABLE `details`  (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '逻辑id',
  `invoice_id` int(11) NULL DEFAULT NULL COMMENT '发票id',
  `goods_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '货物名称',
  `specifications` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '规格型号',
  `unit` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '单位',
  `number` float NULL DEFAULT NULL COMMENT '数量',
  `price` decimal(10, 2) NULL DEFAULT NULL COMMENT '单价',
  `amount_money` decimal(10, 0) NULL DEFAULT NULL COMMENT '金额',
  `tax_amount` decimal(10, 0) NULL DEFAULT NULL COMMENT '税额',
  `product_code` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '商品编码',
  `create_time` datetime NULL DEFAULT NULL COMMENT '录入系统时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 701 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of details
-- ----------------------------
INSERT INTO `details` VALUES (1, 22580, '201702', 'HORI PS4 遥控游玩用 拓展握把 for psv 2000 索尼plays', 'null', 0, 1.00, 152, 152, '152.0', '2020-04-30 18:20:45');
INSERT INTO `details` VALUES (2, 22580, '201702', '好孩子 Goodbaby 学步车XB109E-J231BG', 'null', 0, 11.00, 145, 1592, '1591.9', '2020-04-30 18:20:45');
INSERT INTO `details` VALUES (3, 22580, '201702', 'BroadLink RM mini3 黑豆智能WiFi 空调、电视遥控器', 'null', 0, 3.00, 56, 167, '166.69', '2020-04-30 18:20:45');

-- ----------------------------
-- Table structure for invoice
-- ----------------------------
DROP TABLE IF EXISTS `invoice`;
CREATE TABLE `invoice`  (
  `id` int(11) NOT NULL COMMENT '发票id',
  `sale_id` int(11) NULL DEFAULT NULL COMMENT '销售方企业Id',
  `purchase_id` int(11) NULL DEFAULT NULL COMMENT '购买方企业Id',
  `amount_money` float NULL DEFAULT NULL COMMENT '金额',
  `tax_money` float NULL DEFAULT NULL COMMENT '税额',
  `total_money` float NULL DEFAULT NULL COMMENT '价税合计',
  `billing_time` datetime NULL DEFAULT NULL COMMENT '开票日期',
  `void_mark` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '作废标志',
  `create_time` datetime NULL DEFAULT NULL COMMENT '录入系统时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of invoice
-- ----------------------------
INSERT INTO `invoice` VALUES (72536, 251849, 227192, 43529.9, 7400.09, 50930, '2017-05-08 00:00:00', 'N', '2020-04-30 18:17:58');
INSERT INTO `invoice` VALUES (81663, 839641, 819037, 153.63, 26.12, 179.75, '2017-10-16 00:00:00', 'N', '2020-04-30 18:17:58');
INSERT INTO `invoice` VALUES (83187, 839641, 912831, 762.39, 129.61, 892, '2016-12-14 00:00:00', 'N', '2020-04-30 18:17:58');

-- ----------------------------
-- Table structure for operation_log
-- ----------------------------
DROP TABLE IF EXISTS `operation_log`;
CREATE TABLE `operation_log`  (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'Id',
  `module_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '操作模块',
  `operation_type` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '操作类型',
  `method` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '请求方法',
  `ip` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT 'ip地址',
  `uri` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '接口路径',
  `status` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '操作状态',
  `user_id` int(11) NULL DEFAULT NULL COMMENT '用户Id',
  `create_time` datetime NULL DEFAULT NULL COMMENT '操作时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for permission
-- ----------------------------
DROP TABLE IF EXISTS `permission`;
CREATE TABLE `permission`  (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键，权限ID',
  `perm_name` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '权限名称',
  `value` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '权限值',
  `type` tinyint(4) NULL DEFAULT 1 COMMENT '权限类型：1.菜单/模块，2.待定',
  `create_time` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime NULL DEFAULT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 28 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '权限表' ROW_FORMAT = Compact;

-- ----------------------------
-- Records of permission
-- ----------------------------
INSERT INTO `permission` VALUES (0, '所有权限', '*', 1, '2020-04-20 20:25:55', NULL);
INSERT INTO `permission` VALUES (1, '首页', 'm', 1, '2020-04-12 18:22:41', NULL);
INSERT INTO `permission` VALUES (2, '用户管理', 'm:user', 1, '2020-04-12 18:24:42', NULL);
INSERT INTO `permission` VALUES (3, '信息管理', 'm:user:info', 1, '2020-04-12 18:26:30', NULL);
INSERT INTO `permission` VALUES (5, '发行管理', 'm:release', 1, '2020-04-20 16:42:07', NULL);
INSERT INTO `permission` VALUES (6, '系统发行', 'm:release:sys', 1, '2020-04-20 16:45:22', NULL);
INSERT INTO `permission` VALUES (8, '数据统计', 'm:statistics', 1, '2020-04-20 16:47:51', NULL);
INSERT INTO `permission` VALUES (9, '政务管理', 'm:politics', 1, '2020-04-20 16:51:47', NULL);
INSERT INTO `permission` VALUES (10, '税收政策下发', 'm:politics:release', 1, '2020-04-20 16:51:47', NULL);
INSERT INTO `permission` VALUES (11, '业务指导协调', 'm:politics:guide', 1, '2020-04-20 16:51:47', NULL);
INSERT INTO `permission` VALUES (12, '日志管理', 'm:log', 1, '2020-04-20 16:54:46', NULL);
INSERT INTO `permission` VALUES (13, '登录日志', 'm:log:loginlog', 1, '2020-04-20 16:54:46', NULL);
INSERT INTO `permission` VALUES (14, '操作日志', 'm:log:operlog', 1, '2020-04-20 16:54:46', NULL);
INSERT INTO `permission` VALUES (15, '工作汇报', 'm:politics:report', 1, '2020-04-22 17:52:08', NULL);
INSERT INTO `permission` VALUES (16, '发布通知', 'm:politics:notify', 1, '2020-04-22 17:54:14', NULL);
INSERT INTO `permission` VALUES (17, '上级政策文件', 'm:politics:upRelease', 1, '2020-04-28 19:00:23', NULL);
INSERT INTO `permission` VALUES (18, '数据管理', 'm:data', 1, '2020-04-29 09:59:23', NULL);
INSERT INTO `permission` VALUES (19, '发票数据', 'm:data:invoice', 1, '2020-04-29 10:00:02', NULL);
INSERT INTO `permission` VALUES (20, '纳税人档案', 'm:data:archives', 1, '2020-04-29 10:00:32', NULL);
INSERT INTO `permission` VALUES (21, '发票明细', 'm:data:details', 1, '2020-04-30 14:08:31', NULL);
INSERT INTO `permission` VALUES (22, '各省税收数据', 'm:statistics:provincialtax', 1, '2020-04-30 21:17:45', NULL);
INSERT INTO `permission` VALUES (23, '各地区税收数据', 'm:statistics:areataxdata', 1, '2020-04-30 21:25:50', NULL);
INSERT INTO `permission` VALUES (24, '企业数据', 'm:statistics:enterprisedata', 1, '2020-04-30 21:29:16', NULL);
INSERT INTO `permission` VALUES (25, '预测分析', 'm:predicte_analysis', 1, '2020-04-30 21:38:07', NULL);
INSERT INTO `permission` VALUES (26, '异常企业预测', 'm:predicte_analysis:predicte', 1, '2020-04-30 21:38:07', NULL);
INSERT INTO `permission` VALUES (27, '异常结果分析', 'm:predicte_analysis:analysis', 1, '2020-04-30 21:38:07', NULL);

-- ----------------------------
-- Table structure for politics
-- ----------------------------
DROP TABLE IF EXISTS `politics`;
CREATE TABLE `politics`  (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'id',
  `title` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '标题',
  `abstract` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '摘要',
  `content` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '正文',
  `enclosure_path` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '附件存储地址',
  `create_time` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `politics_type` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '发送方向',
  `reading_signs` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '阅读标识',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 16 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of politics
-- ----------------------------
INSERT INTO `politics` VALUES (12, '工作汇报', '工作汇报', '工作汇报', '', '2020-04-28 20:29:14', 'report', ',[1]');
INSERT INTO `politics` VALUES (14, '发布通知', '发布通知', '发布通知', '', '2020-04-28 20:30:59', 'notify', ',[3],[4]');
INSERT INTO `politics` VALUES (15, '1', '1', '1', '', '2020-05-02 10:19:47', 'guide', ',[2]');

-- ----------------------------
-- Table structure for province
-- ----------------------------
DROP TABLE IF EXISTS `province`;
CREATE TABLE `province`  (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `province_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '省份名称',
  `ban` tinyint(11) NULL DEFAULT NULL COMMENT '是否禁用',
  `create_time` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `province_name`(`province_name`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of province
-- ----------------------------
INSERT INTO `province` VALUES (1, '国家总局', 0, '2020-04-26 14:32:45', '2020-04-26 14:32:45');
INSERT INTO `province` VALUES (2, '吉林省', 0, '2020-04-23 12:20:05', '2020-04-23 12:20:08');

-- ----------------------------
-- Table structure for role
-- ----------------------------
DROP TABLE IF EXISTS `role`;
CREATE TABLE `role`  (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '角色id',
  `role_name` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '角色名称',
  `desc` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '角色描述',
  `priority` int(11) NULL DEFAULT NULL COMMENT '级别，越大级别越高',
  `create_time` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime NULL DEFAULT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 7 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '角色表' ROW_FORMAT = Compact;

-- ----------------------------
-- Records of role
-- ----------------------------
INSERT INTO `role` VALUES (1, '总局系统管理员', '老大', 4, '2020-04-12 18:22:00', '2020-04-22 17:12:45');
INSERT INTO `role` VALUES (2, '省级系统管理员', '二老板', 3, '2020-04-21 18:42:59', '2020-04-21 18:42:59');
INSERT INTO `role` VALUES (3, '数据管理员', '三老板', 2, '2020-04-21 18:42:59', '2020-04-21 18:42:59');
INSERT INTO `role` VALUES (4, '稽查管理员', '四老板', 1, '2020-04-21 18:42:59', '2020-04-21 18:42:59');

-- ----------------------------
-- Table structure for role_perm
-- ----------------------------
DROP TABLE IF EXISTS `role_perm`;
CREATE TABLE `role_perm`  (
  `role_id` int(11) NOT NULL COMMENT '角色ID',
  `permission_id` int(11) NOT NULL COMMENT '权限ID',
  PRIMARY KEY (`role_id`, `permission_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '角色-权限 关联表' ROW_FORMAT = Compact;

-- ----------------------------
-- Records of role_perm
-- ----------------------------
INSERT INTO `role_perm` VALUES (1, 2);
INSERT INTO `role_perm` VALUES (1, 3);
INSERT INTO `role_perm` VALUES (1, 4);

-- ----------------------------
-- Table structure for taxpayer
-- ----------------------------
DROP TABLE IF EXISTS `taxpayer`;
CREATE TABLE `taxpayer`  (
  `id` int(11) NOT NULL COMMENT '纳税人Id',
  `industry_code` int(11) NULL DEFAULT NULL COMMENT '行业代码',
  `registration_type` int(11) NULL DEFAULT NULL COMMENT '登记注册类型',
  `open_time` datetime NULL DEFAULT NULL COMMENT '开业登记时间',
  `create_time` datetime NULL DEFAULT NULL COMMENT '录入系统时间',
  `update_time` datetime NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of taxpayer
-- ----------------------------
INSERT INTO `taxpayer` VALUES (1130, 5139, 173, '2004-11-05 00:00:00', '2020-04-30 18:17:34', '2020-04-30 18:17:34');
INSERT INTO `taxpayer` VALUES (1163, 3599, 159, '2011-12-20 00:00:00', '2020-04-30 18:17:34', '2020-04-30 18:17:34');
INSERT INTO `taxpayer` VALUES (1367, 5132, 173, '2009-04-15 00:00:00', '2020-04-30 18:17:34', '2020-04-30 18:17:34');

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键，用户ID',
  `username` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '用户名',
  `nick` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '昵称',
  `password` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '明文-密码',
  `role_id` int(11) NOT NULL COMMENT '角色ID',
  `province_id` int(11) NULL DEFAULT NULL COMMENT '所属省份',
  `create_time` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime NULL DEFAULT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `username`(`username`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 7 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '用户表' ROW_FORMAT = Compact;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES (1, 'admin', '总局老总', '123456', 1, 1, '2020-04-12 18:20:16', '2020-04-22 17:01:37');
INSERT INTO `user` VALUES (2, 'pro-admin', '省级大哥', '123456', 2, 2, '2020-04-22 20:14:13', '2020-04-22 20:14:13');
INSERT INTO `user` VALUES (3, 'data-admin', '数据小弟', '123456', 3, 2, '2020-04-26 10:55:56', '2020-04-26 10:55:56');
INSERT INTO `user` VALUES (4, 'audit-admin', '稽查先生', '123456', 4, 2, '2020-04-26 10:58:07', '2020-04-26 10:58:07');

SET FOREIGN_KEY_CHECKS = 1;
