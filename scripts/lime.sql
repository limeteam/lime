/*
 Navicat Premium Data Transfer

 Source Server         : local-dev
 Source Server Type    : MySQL
 Source Server Version : 50628
 Source Host           : 172.72.33.11:3306
 Source Schema         : lime

 Target Server Type    : MySQL
 Target Server Version : 50628
 File Encoding         : 65001

 Date: 10/05/2020 21:37:07
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for novel_book
-- ----------------------------
DROP TABLE IF EXISTS `novel_book`;
CREATE TABLE `novel_book`  (
  `id` int(11) NOT NULL COMMENT '分类ID',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '分类名称',
  `old_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '原名',
  `channel_id` tinyint(2) NULL DEFAULT NULL COMMENT '所属频道',
  `category_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '小说分类',
  `desc` varchar(2555) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '小说描述',
  `cover` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '小说封面',
  `author` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '小说作者',
  `source` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '来源',
  `split_rule` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '拆分规则：0 智能分章 1 固定字数分章 2 标签分章',
  `upload_file` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '上传小说',
  `status` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '小说状态，0连载中，1已完结，2太监',
  `attribute` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '属性，0 新书 1 热门 2 会员 3 热门 4 精选 5 大神',
  `chapter_price` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '每章价格',
  `thousand_characters_price` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '千字价格',
  `score` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '评分',
  `text_num` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '小说字数',
  `chapter_num` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '小说章节数',
  `chapter_updated_at` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '最新章节时间',
  `chapter_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '最新章节id',
  `chapter_title` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '最新章节标题',
  `views` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '浏览次数',
  `collect_num` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '收藏数',
  `online_status` tinyint(2) NOT NULL COMMENT '上架状态：0 已上架 1 已下架',
  `is_sensitivity` tinyint(2) NOT NULL COMMENT '是否敏感： 0 不敏感 1 敏感',
  `created_at` timestamp(0) NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '创建时间',
  `updated_at` timestamp(0) NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '修改时间',
  `deleted_at` timestamp(0) NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for novel_book_links
-- ----------------------------
DROP TABLE IF EXISTS `novel_book_links`;
CREATE TABLE `novel_book_links`  (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
  `novel_id` int(11) UNSIGNED NOT NULL COMMENT '小说ID',
  `link` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '采集来源小说简介URL',
  `source` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '采集站标识',
  `chapter_link` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '采集来源章节页面URL',
  `updated_at` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '更新时间',
  `created_at` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '创建时间',
  `deleted_at` timestamp(0) NOT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `udx_link_source`(`source`, `link`) USING BTREE,
  INDEX `idx_novid`(`novel_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '小说采集站点链接关联表' ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for novel_category
-- ----------------------------
DROP TABLE IF EXISTS `novel_category`;
CREATE TABLE `novel_category`  (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '分类ID',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '分类名称',
  `channel_id` tinyint(2) NULL DEFAULT NULL COMMENT '所属频道',
  `novel_num` int(11) NULL DEFAULT NULL COMMENT '小说数',
  `sort` int(11) NULL DEFAULT NULL COMMENT '排序',
  `created_at` timestamp(0) NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '创建时间',
  `updated_at` timestamp(0) NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '修改时间',
  `deleted_at` timestamp(0) NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 15 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of novel_category
-- ----------------------------
INSERT INTO `novel_category` VALUES (1, '网游竞技', 1, 0, 4, '2020-05-10 19:12:56', '2020-05-10 19:14:42', NULL);
INSERT INTO `novel_category` VALUES (2, '同人耽美', 2, 0, 12, '2020-05-10 19:12:56', '2020-05-10 19:14:47', NULL);
INSERT INTO `novel_category` VALUES (3, '穿越架空', 2, 0, 7, '2020-05-10 19:12:56', '2020-05-10 19:14:53', NULL);
INSERT INTO `novel_category` VALUES (4, '都市现实', 1, 0, 2, '2020-05-10 19:12:56', '2020-05-10 19:15:00', NULL);
INSERT INTO `novel_category` VALUES (5, '科幻未来', 1, 0, 4, '2020-05-10 19:12:56', '2020-05-10 19:15:05', NULL);
INSERT INTO `novel_category` VALUES (6, '武侠仙侠', 1, 0, 3, '2020-05-10 19:12:56', '2020-05-10 19:21:50', NULL);
INSERT INTO `novel_category` VALUES (7, '玄幻奇幻', 1, 0, 6, '2020-05-10 19:12:56', '2020-05-10 19:11:54', NULL);
INSERT INTO `novel_category` VALUES (8, '青春校园', 2, 0, 2, '2020-05-10 19:12:56', '2020-05-10 19:12:10', NULL);
INSERT INTO `novel_category` VALUES (9, '总裁高干', 2, 0, 12, '2020-05-10 19:12:56', '2020-05-10 19:12:24', NULL);
INSERT INTO `novel_category` VALUES (10, '军事历史', 1, 0, 9, '2020-05-10 19:12:56', '2020-05-10 19:12:40', NULL);
INSERT INTO `novel_category` VALUES (11, '悬疑灵异', 1, 0, 8, '2020-05-10 19:12:56', '2020-05-10 19:12:56', NULL);
INSERT INTO `novel_category` VALUES (12, '古代情缘', 2, 0, 11, '2020-05-10 19:12:56', '2020-05-10 19:13:09', NULL);
INSERT INTO `novel_category` VALUES (13, '现代爱情', 0, 0, 13, '2020-05-10 19:12:56', '2020-05-10 19:58:10', NULL);
INSERT INTO `novel_category` VALUES (14, '其他类型', 0, 0, 14, '2020-05-10 21:08:35', '2020-05-10 21:08:35', NULL);

-- ----------------------------
-- Table structure for novel_chapter
-- ----------------------------
DROP TABLE IF EXISTS `novel_chapter`;
CREATE TABLE `novel_chapter`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `book_id` int(11) NULL DEFAULT NULL,
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `content` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `status` int(11) NULL DEFAULT NULL,
  `url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `create_time` timestamp(0) NULL DEFAULT NULL,
  `last_update_time` timestamp(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for novel_chapter_0000
-- ----------------------------
DROP TABLE IF EXISTS `novel_chapter_0000`;
CREATE TABLE `novel_chapter_0000`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `novel_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '小说ID',
  `chapter_no` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '章节编号',
  `title` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '章节标题',
  `desc` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '章节内容',
  `link` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '章节采集链接',
  `is_vip` tinyint(2) NOT NULL COMMENT '是否收费',
  `source` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '章节采集站点源',
  `views` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '浏览次数',
  `text_num` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '章节字数',
  `status` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '章节采集状态0正常，1失败',
  `try_views` tinyint(2) UNSIGNED NOT NULL DEFAULT 0 COMMENT '采集重试次数',
  `created_at` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0),
  `updated_at` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0),
  `deleted_at` timestamp(0) NOT NULL DEFAULT '0000-00-00 00:00:00',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `udx_novid_no`(`novel_id`, `chapter_no`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for novel_comment
-- ----------------------------
DROP TABLE IF EXISTS `novel_comment`;
CREATE TABLE `novel_comment`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `novel_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '小说ID',
  `username` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '章节标题',
  `content` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '内容',
  `likes` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '点赞数',
  `source` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '来源',
  `created_at` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0),
  `updated_at` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0),
  `deleted_at` timestamp(0) NOT NULL DEFAULT '0000-00-00 00:00:00',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Compact;

SET FOREIGN_KEY_CHECKS = 1;
