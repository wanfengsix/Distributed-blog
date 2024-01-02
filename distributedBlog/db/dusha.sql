/*
 Navicat Premium Data Transfer

 Source Server         : DistributedBlog
 Source Server Type    : MySQL
 Source Server Version : 80100
 Source Host           : localhost:3306
 Source Schema         : dusha

 Target Server Type    : MySQL
 Target Server Version : 80100
 File Encoding         : 65001

 Date: 02/01/2024 22:36:36
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for article
-- ----------------------------
DROP TABLE IF EXISTS `article`;
CREATE TABLE `article`  (
  `Article_ID` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `head` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `date` datetime NULL DEFAULT NULL,
  `UID` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `likes_nums` int NULL DEFAULT NULL,
  `comment_nums` int NULL DEFAULT NULL,
  `article_url` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `abstract` varchar(80) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `is_visible` int NULL DEFAULT NULL,
  PRIMARY KEY (`Article_ID`) USING BTREE,
  INDEX `UID`(`UID` ASC) USING BTREE,
  INDEX `head`(`head` ASC) USING BTREE,
  CONSTRAINT `article_ibfk_1` FOREIGN KEY (`UID`) REFERENCES `user` (`UID`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of article
-- ----------------------------
INSERT INTO `article` VALUES ('0', '前端出什么事了！', '2023-12-05 00:00:00', '0', 99, 17, '1.txt', NULL, NULL);
INSERT INTO `article` VALUES ('026', '243', '2024-01-02 00:00:00', '0', 0, 0, '026.txt', NULL, NULL);
INSERT INTO `article` VALUES ('04', '', '2024-01-02 00:00:00', '0', 0, 0, '04.txt', NULL, NULL);
INSERT INTO `article` VALUES ('047', '123', '2024-01-02 00:00:00', '0', 0, 0, '047.txt', NULL, NULL);
INSERT INTO `article` VALUES ('1', '百度', '2023-12-12 20:12:06', '0', 435, 0, '1.txt', NULL, NULL);
INSERT INTO `article` VALUES ('2', '阿里', '2023-12-12 20:13:03', '0', 546, 0, '1.txt', NULL, NULL);
INSERT INTO `article` VALUES ('3', '腾讯', '2023-12-12 20:13:22', '0', 54, 0, '1.txt', NULL, NULL);
INSERT INTO `article` VALUES ('4', '阿阿阿里', '2024-01-01 21:29:17', '0', 3245, 0, '1.txt', NULL, NULL);
INSERT INTO `article` VALUES ('5', '亚马逊', '2023-12-12 20:14:00', '0', 55, 0, '1.txt', NULL, NULL);
INSERT INTO `article` VALUES ('6', '马斯克', '2023-12-12 20:14:19', '0', 355, 0, '1.txt', NULL, NULL);
INSERT INTO `article` VALUES ('7', '百百百度！', '2024-01-01 21:15:23', '0', 342, 3423, '1.txt', NULL, NULL);
INSERT INTO `article` VALUES ('8', 'spacex', '2023-12-12 20:14:38', '0', 5677, 0, '1.txt', NULL, NULL);
INSERT INTO `article` VALUES ('9', '谷歌', '2023-12-12 20:13:39', '0', 75, 0, '1.txt', NULL, NULL);

-- ----------------------------
-- Table structure for comment
-- ----------------------------
DROP TABLE IF EXISTS `comment`;
CREATE TABLE `comment`  (
  `Comment_ID` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `Comment_content` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `Article_ID` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `UID` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  PRIMARY KEY (`Comment_ID`) USING BTREE,
  INDEX `UID`(`UID` ASC) USING BTREE,
  INDEX `Article_ID`(`Article_ID` ASC) USING BTREE,
  CONSTRAINT `comment_ibfk_1` FOREIGN KEY (`Article_ID`) REFERENCES `article` (`Article_ID`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `comment_ibfk_2` FOREIGN KEY (`UID`) REFERENCES `user` (`UID`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of comment
-- ----------------------------
INSERT INTO `comment` VALUES ('02023-12-17 14:38:15', '', '0', '0');
INSERT INTO `comment` VALUES ('02023-12-17 14:38:58', '324324', '0', '0');
INSERT INTO `comment` VALUES ('02023-12-17 19:12:52', '', '0', '0');

-- ----------------------------
-- Table structure for draft
-- ----------------------------
DROP TABLE IF EXISTS `draft`;
CREATE TABLE `draft`  (
  `UID` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `draft_date` datetime NULL DEFAULT NULL,
  `draft_head` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `draft_url` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  PRIMARY KEY (`UID`) USING BTREE,
  CONSTRAINT `draft_ibfk_1` FOREIGN KEY (`UID`) REFERENCES `user` (`UID`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of draft
-- ----------------------------
INSERT INTO `draft` VALUES ('0', '2024-01-02 00:00:00', '214', '074.txt');

-- ----------------------------
-- Table structure for follow
-- ----------------------------
DROP TABLE IF EXISTS `follow`;
CREATE TABLE `follow`  (
  `followed_id` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `following_id` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  PRIMARY KEY (`followed_id`, `following_id`) USING BTREE,
  INDEX `following_id`(`following_id` ASC) USING BTREE,
  CONSTRAINT `follow_ibfk_1` FOREIGN KEY (`followed_id`) REFERENCES `user` (`UID`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `follow_ibfk_2` FOREIGN KEY (`following_id`) REFERENCES `user` (`UID`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of follow
-- ----------------------------

-- ----------------------------
-- Table structure for likes
-- ----------------------------
DROP TABLE IF EXISTS `likes`;
CREATE TABLE `likes`  (
  `UID` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `Article_ID` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  PRIMARY KEY (`UID`, `Article_ID`) USING BTREE,
  INDEX `Article_ID`(`Article_ID` ASC) USING BTREE,
  CONSTRAINT `likes_ibfk_1` FOREIGN KEY (`UID`) REFERENCES `user` (`UID`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `likes_ibfk_2` FOREIGN KEY (`Article_ID`) REFERENCES `article` (`Article_ID`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of likes
-- ----------------------------

-- ----------------------------
-- Table structure for notice
-- ----------------------------
DROP TABLE IF EXISTS `notice`;
CREATE TABLE `notice`  (
  `UID` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `u_name` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `author_ID` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `author_name` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `Article_ID` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `head` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `type` varchar(6) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `is_read` int NULL DEFAULT NULL,
  PRIMARY KEY (`UID`, `author_ID`, `Article_ID`, `type`) USING BTREE,
  INDEX `u_name`(`u_name` ASC) USING BTREE,
  INDEX `author_ID`(`author_ID` ASC) USING BTREE,
  INDEX `author_name`(`author_name` ASC) USING BTREE,
  INDEX `Article_ID`(`Article_ID` ASC) USING BTREE,
  INDEX `head`(`head` ASC) USING BTREE,
  CONSTRAINT `notice_ibfk_1` FOREIGN KEY (`UID`) REFERENCES `user` (`UID`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `notice_ibfk_2` FOREIGN KEY (`u_name`) REFERENCES `user` (`u_name`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `notice_ibfk_3` FOREIGN KEY (`author_ID`) REFERENCES `user` (`UID`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `notice_ibfk_4` FOREIGN KEY (`author_name`) REFERENCES `user` (`u_name`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `notice_ibfk_5` FOREIGN KEY (`Article_ID`) REFERENCES `article` (`Article_ID`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `notice_ibfk_6` FOREIGN KEY (`head`) REFERENCES `article` (`head`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of notice
-- ----------------------------

-- ----------------------------
-- Table structure for readed
-- ----------------------------
DROP TABLE IF EXISTS `readed`;
CREATE TABLE `readed`  (
  `UID` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `Article_ID` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  PRIMARY KEY (`UID`, `Article_ID`) USING BTREE,
  INDEX `Article_ID`(`Article_ID` ASC) USING BTREE,
  CONSTRAINT `readed_ibfk_1` FOREIGN KEY (`UID`) REFERENCES `user` (`UID`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `readed_ibfk_2` FOREIGN KEY (`Article_ID`) REFERENCES `article` (`Article_ID`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of readed
-- ----------------------------

-- ----------------------------
-- Table structure for register
-- ----------------------------
DROP TABLE IF EXISTS `register`;
CREATE TABLE `register`  (
  `u_name` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `password` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `secret_protection1` varchar(80) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `secret_protection2` varchar(80) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `secret_protection3` varchar(80) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `is_Admin` int NOT NULL,
  PRIMARY KEY (`u_name`) USING BTREE,
  CONSTRAINT `register_ibfk_1` FOREIGN KEY (`u_name`) REFERENCES `user` (`u_name`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of register
-- ----------------------------
INSERT INTO `register` VALUES ('admin', 'admin', '1', '2', '3', 1);

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `UID` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `u_name` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `following` int NULL DEFAULT NULL,
  `followed` int NULL DEFAULT NULL,
  `article_nums` int NULL DEFAULT NULL,
  `read_nums` int NULL DEFAULT NULL,
  `comment_nums` int NULL DEFAULT NULL,
  `likes_nums` int NULL DEFAULT NULL,
  `level` int NULL DEFAULT NULL,
  PRIMARY KEY (`UID`) USING BTREE,
  UNIQUE INDEX `u_name`(`u_name` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES ('0', 'admin', 0, 0, 0, 0, 17, 0, 0);
INSERT INTO `user` VALUES ('324', '小李', 44, 44, 44, 44, 444, 4, 4);
INSERT INTO `user` VALUES ('343', '小明', 222, 222, 2, 90, 0, 7, 7);
INSERT INTO `user` VALUES ('366', '李盼', 78, 777, 2, 90, 0, 7, 7);
INSERT INTO `user` VALUES ('455', '马应龙', 333, 243, 546, 90, 7, 7, 7);

-- ----------------------------
-- Table structure for user_information
-- ----------------------------
DROP TABLE IF EXISTS `user_information`;
CREATE TABLE `user_information`  (
  `UID` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `signature` varchar(80) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `avatar_url` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  PRIMARY KEY (`UID`) USING BTREE,
  CONSTRAINT `user_information_ibfk_1` FOREIGN KEY (`UID`) REFERENCES `user` (`UID`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of user_information
-- ----------------------------
INSERT INTO `user_information` VALUES ('0', 'sdada', '管理员头像.jpg');

SET FOREIGN_KEY_CHECKS = 1;
