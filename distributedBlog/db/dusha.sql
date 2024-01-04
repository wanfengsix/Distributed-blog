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

 Date: 05/01/2024 01:51:16
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
INSERT INTO `article` VALUES ('030', 'spacex', '2024-01-04 00:00:00', '0', 0, 0, '030.txt', '<p>埃隆·马斯克领导的spacex获得一份巨', 0);
INSERT INTO `article` VALUES ('037', '被裁后的一天', '2024-01-04 00:00:00', '0', 0, 0, '037.txt', '<p>距离被裁已经过去了34天</p>', 1);
INSERT INTO `article` VALUES ('050', '前端出什么事了', '2024-01-04 00:00:00', '0', 1, 2, '050.txt', '<p>最近前端出了一个信框架</p>', 1);
INSERT INTO `article` VALUES ('089', '生存or毁灭', '2024-01-04 00:00:00', '0', 0, 0, '089.txt', '<p>今年是qq空间诞生18周年</p>', 1);
INSERT INTO `article` VALUES ('091', '亚马逊', '2024-01-04 00:00:00', '0', 0, 0, '091.txt', '<p>亚马逊是世界最大网上购物商城</p>', 1);
INSERT INTO `article` VALUES ('095', '百度', '2024-01-04 00:00:00', '0', 0, 0, '095.txt', '<p>三大巨头之一</p>', 1);
INSERT INTO `article` VALUES ('11', 'hello', '2024-01-04 00:00:00', '789', 9, 0, '78911.txt', '华人和派遣无法一概而论', 0);
INSERT INTO `article` VALUES ('12336', '听我一句劝，业务代码中', '2024-01-05 00:00:00', '123', 0, 0, '12336.txt', '<p>工作三年多，没用过多线程</p>', 1);
INSERT INTO `article` VALUES ('12344', '大环境越不好，人就越玄学', '2024-01-05 00:00:00', '123', 0, 0, '12344.txt', '<p>2002年，大环境还没这么拉跨的时候<', 1);
INSERT INTO `article` VALUES ('12369', '官宣！掘金年度技术演讲', '2024-01-05 00:00:00', '123', 0, 0, '12369.txt', '<p>2023年即将过去，马上迎来新的一年<', 1);
INSERT INTO `article` VALUES ('12379', '为什么我总学不好TS', '2024-01-05 00:00:00', '123', 0, 0, '12379.txt', '<p>有没有同学，学TS的步骤和我一样</p', 1);
INSERT INTO `article` VALUES ('26', '你好', '2024-01-02 00:00:00', '0', 0, 0, '26.txt', '大家好，我小j，很久不见，这几个月摸鱼摸得过分了，', 0);
INSERT INTO `article` VALUES ('78961', '社会现实告诉我，00后整顿职场就是个笑话', '2024-01-04 00:00:00', '789', 0, 0, '78961.txt', '<p>00后整顿职场，也算是我之前的关键', 0);
INSERT INTO `article` VALUES ('78997', 'hellooo', '2024-01-04 00:00:00', '789', 0, 0, '78997.txt', '事实上，日本之前确实很卷，毕竟是社畜这个词的产生地', 0);

-- ----------------------------
-- Table structure for comment
-- ----------------------------
DROP TABLE IF EXISTS `comment`;
CREATE TABLE `comment`  (
  `Comment_ID` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `Comment_content` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `Article_ID` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `UID` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `date` date NULL DEFAULT NULL,
  PRIMARY KEY (`Comment_ID`) USING BTREE,
  INDEX `UID`(`UID` ASC) USING BTREE,
  INDEX `Article_ID`(`Article_ID` ASC) USING BTREE,
  CONSTRAINT `comment_ibfk_1` FOREIGN KEY (`Article_ID`) REFERENCES `article` (`Article_ID`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `comment_ibfk_2` FOREIGN KEY (`UID`) REFERENCES `user` (`UID`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of comment
-- ----------------------------

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
INSERT INTO `draft` VALUES ('123', '2024-01-04 00:00:00', '3242', '12312.txt');
INSERT INTO `draft` VALUES ('789', '2024-01-04 00:00:00', '213', '78982.txt');

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
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of follow
-- ----------------------------
INSERT INTO `follow` VALUES ('0', '0');
INSERT INTO `follow` VALUES ('0', '123');
INSERT INTO `follow` VALUES ('0', '789');

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
INSERT INTO `likes` VALUES ('0', '050');

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
  `type` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `is_read` int NULL DEFAULT NULL,
  `notice_ID` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  PRIMARY KEY (`notice_ID`) USING BTREE,
  INDEX `UID`(`UID` ASC) USING BTREE,
  INDEX `author_ID`(`author_ID` ASC) USING BTREE,
  INDEX `Article_ID`(`Article_ID` ASC) USING BTREE,
  CONSTRAINT `notice_ibfk_1` FOREIGN KEY (`UID`) REFERENCES `user` (`UID`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `notice_ibfk_2` FOREIGN KEY (`author_ID`) REFERENCES `user` (`UID`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `notice_ibfk_3` FOREIGN KEY (`Article_ID`) REFERENCES `article` (`Article_ID`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of notice
-- ----------------------------
INSERT INTO `notice` VALUES ('0', 'admin', '0', 'admin', '050', '前端出什么事了', '点赞', 0, '02024-01-05187');
INSERT INTO `notice` VALUES ('0', 'admin', '0', 'admin', '050', '前端出什么事了', '评论', 0, '02024-01-05501');
INSERT INTO `notice` VALUES ('0', 'admin', '0', 'admin', '050', '前端出什么事了', '评论', 0, '02024-01-05778');

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
INSERT INTO `register` VALUES ('123', '123', '123', '123', '123', 0);
INSERT INTO `register` VALUES ('456', '456', '456', '456', '456', 0);
INSERT INTO `register` VALUES ('789', '789', '789', '789', '789', 0);
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
INSERT INTO `user` VALUES ('0', 'admin', 1, 3, 0, 0, 42, 2, 0);
INSERT INTO `user` VALUES ('123', '123', 1, 0, 0, 0, 0, 0, 0);
INSERT INTO `user` VALUES ('324', '小李', 44, 44, 44, 44, 444, 4, 4);
INSERT INTO `user` VALUES ('343', '小明', 222, 222, 2, 90, 0, 7, 7);
INSERT INTO `user` VALUES ('366', '李盼', 78, 777, 2, 90, 0, 7, 7);
INSERT INTO `user` VALUES ('455', '马应龙', 333, 243, 546, 90, 7, 7, 7);
INSERT INTO `user` VALUES ('456', '456', 0, 0, 0, 0, 0, 0, 0);
INSERT INTO `user` VALUES ('789', '789', 1, 0, 0, 0, 0, 0, 0);

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
INSERT INTO `user_information` VALUES ('123', '', '123的头像.jpg');
INSERT INTO `user_information` VALUES ('456', '', '456的头像.jpg');
INSERT INTO `user_information` VALUES ('789', '', '789的头像.jpg');

SET FOREIGN_KEY_CHECKS = 1;
