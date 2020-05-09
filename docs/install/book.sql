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

 Date: 09/05/2020 16:14:46
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for book
-- ----------------------------
DROP TABLE IF EXISTS `book`;
CREATE TABLE `book`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `author` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `image` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `status` int(11) NULL DEFAULT NULL,
  `url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `desc` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `create_time` datetime(0) NULL DEFAULT NULL,
  `last_update_time` datetime(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for chapter
-- ----------------------------
DROP TABLE IF EXISTS `chapter`;
CREATE TABLE `chapter`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `book_id` int(11) NULL DEFAULT NULL,
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `content` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `status` int(11) NULL DEFAULT NULL,
  `url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `create_time` datetime(0) NULL DEFAULT NULL,
  `last_update_time` datetime(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Compact;

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
  `created_at` int(10) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` int(10) NULL DEFAULT NULL COMMENT '修改时间',
  `deleted_at` int(10) NULL DEFAULT NULL COMMENT '删除时间',
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
  `updated_at` int(11) UNSIGNED NOT NULL COMMENT '更新时间',
  `created_at` int(11) UNSIGNED NOT NULL COMMENT '创建时间',
  `deleted_at` int(11) UNSIGNED NOT NULL COMMENT '删除时间',
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
  `created_at` int(10) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` int(10) NULL DEFAULT NULL COMMENT '修改时间',
  `deleted_at` int(10) NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 14 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of novel_category
-- ----------------------------
INSERT INTO `novel_category` VALUES (1, '网游竞技	', 1, NULL, 13, NULL, NULL, 0);
INSERT INTO `novel_category` VALUES (2, '同人耽美	', 2, NULL, 12, NULL, NULL, 0);
INSERT INTO `novel_category` VALUES (3, '穿越架空	', 3, NULL, 11, NULL, NULL, 0);
INSERT INTO `novel_category` VALUES (4, '都市现实	', 4, NULL, 10, NULL, NULL, 0);
INSERT INTO `novel_category` VALUES (5, '科幻未来	', 5, NULL, 9, NULL, NULL, 0);
INSERT INTO `novel_category` VALUES (6, '武侠仙侠	', 6, NULL, 8, NULL, NULL, 0);
INSERT INTO `novel_category` VALUES (7, '玄幻奇幻	', 7, NULL, 7, NULL, NULL, 0);
INSERT INTO `novel_category` VALUES (8, '青春校园	', 8, NULL, 6, NULL, NULL, 0);
INSERT INTO `novel_category` VALUES (9, '总裁高干	', 9, NULL, 5, NULL, NULL, 0);
INSERT INTO `novel_category` VALUES (10, '军事历史	', 10, NULL, 4, NULL, NULL, 0);
INSERT INTO `novel_category` VALUES (11, '悬疑灵异	', 11, NULL, 3, NULL, NULL, 0);
INSERT INTO `novel_category` VALUES (12, '古代情缘	', 12, NULL, 2, NULL, NULL, 0);
INSERT INTO `novel_category` VALUES (13, '现代爱情	', 13, NULL, 1, NULL, NULL, 0);

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
  `created_at` int(10) UNSIGNED NOT NULL DEFAULT 0,
  `updated_at` int(10) UNSIGNED NOT NULL DEFAULT 0,
  `deleted_at` int(10) UNSIGNED NOT NULL DEFAULT 0,
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
  `created_at` int(10) UNSIGNED NOT NULL DEFAULT 0,
  `updated_at` int(10) UNSIGNED NOT NULL DEFAULT 0,
  `deleted_at` int(10) UNSIGNED NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for task
-- ----------------------------
DROP TABLE IF EXISTS `task`;
CREATE TABLE `task`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `task_name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `task_rule_name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `task_desc` varchar(512) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `status` tinyint(4) NOT NULL DEFAULT 0,
  `counts` int(11) NOT NULL DEFAULT 0,
  `cron_spec` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `output_type` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `opt_user_agent` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `opt_max_depth` int(11) NOT NULL DEFAULT 0,
  `opt_allowed_domains` varchar(512) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `opt_url_filters` varchar(512) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `opt_max_body_size` int(11) NOT NULL DEFAULT 0,
  `opt_request_timeout` int(11) NOT NULL DEFAULT 10,
  `auto_migrate` tinyint(4) NOT NULL DEFAULT 0,
  `limit_enable` tinyint(4) NOT NULL DEFAULT 0,
  `limit_domain_regexp` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `limit_domain_glob` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `limit_delay` int(11) NOT NULL DEFAULT 0,
  `limit_random_delay` int(11) NOT NULL DEFAULT 0,
  `limit_parallelism` int(11) NOT NULL DEFAULT 0,
  `proxy_urls` varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `created_at` datetime(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0),
  `updated_at` datetime(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) ON UPDATE CURRENT_TIMESTAMP(0),
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uk_task_name`(`task_name`) USING BTREE,
  INDEX `idx_created_at`(`created_at`) USING BTREE,
  INDEX `idx_updated_at`(`updated_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Compact;

SET FOREIGN_KEY_CHECKS = 1;
