/*
 Navicat Premium Data Transfer

 Source Server         : .
 Source Server Type    : MySQL
 Source Server Version : 80017
 Source Host           : 127.0.0.1:3306
 Source Schema         : gomicro

 Target Server Type    : MySQL
 Target Server Version : 80017
 File Encoding         : 65001

 Date: 06/12/2019 19:41:48
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `address` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `phone` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES (1, 'wt2rfjyr1c', 'beijingwt2rfjyr1c', '134wt2rfjyr');
INSERT INTO `user` VALUES (2, '7up7guhtsq', 'beijing7up7guhtsq', '1347up7guht');
INSERT INTO `user` VALUES (3, 'hp45hspno8', 'beijinghp45hspno8', '134hp45hspn');
INSERT INTO `user` VALUES (4, 'ij53vhutk9', 'beijingij53vhutk9', '134ij53vhut');
INSERT INTO `user` VALUES (5, '234', 'address', '12343234535');
INSERT INTO `user` VALUES (6, '123', '123', '123');
INSERT INTO `user` VALUES (7, '123', '123', '123');
INSERT INTO `user` VALUES (8, '7esz53uaoa', 'beijing7esz53uaoa', '1347esz53ua');

SET FOREIGN_KEY_CHECKS = 1;
