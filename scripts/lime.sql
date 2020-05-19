/*
 Navicat Premium Data Transfer

 Source Server         : windows-dev
 Source Server Type    : MySQL
 Source Server Version : 50729
 Source Host           : localhost:3306
 Source Schema         : lime

 Target Server Type    : MySQL
 Target Server Version : 50729
 File Encoding         : 65001

 Date: 19/05/2020 23:26:17
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for comic_category
-- ----------------------------
DROP TABLE IF EXISTS `comic_category`;
CREATE TABLE `comic_category`  (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '分类ID',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '分类名称',
  `comic_num` int(11) NULL DEFAULT NULL COMMENT '小说数',
  `sort` int(11) NULL DEFAULT NULL COMMENT '排序',
  `created_at` timestamp(0) NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '创建时间',
  `updated_at` timestamp(0) NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '修改时间',
  `deleted_at` timestamp(0) NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 17 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of comic_category
-- ----------------------------
INSERT INTO `comic_category` VALUES (1, '	日常', 1, 2, '2020-05-19 22:34:02', '2020-05-19 22:34:30', NULL);
INSERT INTO `comic_category` VALUES (2, '搞笑', 0, 1, '2020-05-19 22:34:45', '2020-05-19 22:34:45', NULL);
INSERT INTO `comic_category` VALUES (3, '热血', 0, 2, '2020-05-19 22:34:55', '2020-05-19 22:34:55', NULL);
INSERT INTO `comic_category` VALUES (4, '古风', 0, 0, '2020-05-19 22:35:03', '2020-05-19 22:35:03', NULL);
INSERT INTO `comic_category` VALUES (5, '猎奇', 0, 0, '2020-05-19 22:35:11', '2020-05-19 22:35:11', NULL);
INSERT INTO `comic_category` VALUES (6, '同人', 0, 0, '2020-05-19 22:35:18', '2020-05-19 22:35:18', NULL);
INSERT INTO `comic_category` VALUES (7, '都市', 0, 0, '2020-05-19 22:35:24', '2020-05-19 22:35:24', NULL);
INSERT INTO `comic_category` VALUES (8, '冒险', 0, 0, '2020-05-19 22:35:32', '2020-05-19 22:35:32', NULL);
INSERT INTO `comic_category` VALUES (9, '恋爱', 0, 0, '2020-05-19 22:35:38', '2020-05-19 22:35:38', NULL);
INSERT INTO `comic_category` VALUES (10, '校园', 0, 0, '2020-05-19 22:35:45', '2020-05-19 22:35:45', NULL);
INSERT INTO `comic_category` VALUES (11, '日漫', 0, 0, '2020-05-19 22:35:51', '2020-05-19 22:35:51', NULL);
INSERT INTO `comic_category` VALUES (12, '治愈', 0, 0, '2020-05-19 22:35:56', '2020-05-19 22:35:56', NULL);
INSERT INTO `comic_category` VALUES (13, '玄幻', 0, 0, '2020-05-19 22:36:03', '2020-05-19 22:36:03', NULL);
INSERT INTO `comic_category` VALUES (14, '奇幻', 0, 0, '2020-05-19 22:36:08', '2020-05-19 22:36:08', NULL);
INSERT INTO `comic_category` VALUES (15, '恐怖', 0, 0, '2020-05-19 22:36:15', '2020-05-19 22:36:15', NULL);
INSERT INTO `comic_category` VALUES (16, '悬疑', 0, 0, '2020-05-19 22:36:20', '2020-05-19 22:36:20', NULL);

-- ----------------------------
-- Table structure for comic_comment
-- ----------------------------
DROP TABLE IF EXISTS `comic_comment`;
CREATE TABLE `comic_comment`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `comic_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '漫画ID',
  `username` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户',
  `content` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '内容',
  `source` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '来源',
  `created_at` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0),
  `updated_at` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0),
  `deleted_at` timestamp(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of comic_comment
-- ----------------------------
INSERT INTO `comic_comment` VALUES (3, 1, '春天的女孩', '发发发', '网络', '2020-05-19 22:57:32', '2020-05-19 22:57:32', NULL);

-- ----------------------------
-- Table structure for comics
-- ----------------------------
DROP TABLE IF EXISTS `comics`;
CREATE TABLE `comics`  (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '漫画名',
  `old_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '原名',
  `channel_id` tinyint(2) NULL DEFAULT NULL COMMENT '所属频道',
  `category_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '漫画分类',
  `desc` varchar(2555) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '漫画简介',
  `horizontal_cover` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '横版封面',
  `vertical_cover` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '竖版封面',
  `author` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '漫画作者',
  `source` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '来源',
  `status` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '小说状态，0连载中，1已完结',
  `flag` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '属性，0 免费 1 热门 2 会员 3 推荐',
  `chapter_price` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '章节定价',
  `chapter_updated_at` timestamp(0) NULL DEFAULT NULL COMMENT '最新章节时间',
  `views` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '浏览次数',
  `collect_num` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '收藏数',
  `online_status` tinyint(2) NOT NULL COMMENT '上架状态：0 已上架 1 已下架',
  `is_sensitivity` tinyint(2) NOT NULL COMMENT '是否敏感： 0 不敏感 1 敏感',
  `created_at` timestamp(0) NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '创建时间',
  `updated_at` timestamp(0) NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '修改时间',
  `deleted_at` timestamp(0) NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for novel_book
-- ----------------------------
DROP TABLE IF EXISTS `novel_book`;
CREATE TABLE `novel_book`  (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '小说名称',
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
  `flag` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '属性，0 新书 1 热门 2 会员 3 热门 4 精选 5 大神',
  `chapter_price` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '每章价格',
  `thousand_characters_price` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '千字价格',
  `score` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '评分',
  `text_num` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '小说字数',
  `chapter_num` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '小说章节数',
  `chapter_updated_at` timestamp(0) NULL DEFAULT NULL COMMENT '最新章节时间',
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
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of novel_book
-- ----------------------------
INSERT INTO `novel_book` VALUES (1, '生生不息 ', '生生不息2 ', 2, 2, 'werwerwer', '/upload/images/1589430526.jpg', '张真人', '网络', 0, '', 0, '新书,会员,热门', 2, 4, 2, 0, 0, '2020-05-16 10:13:55', 0, '', 0, 0, 0, 0, '2020-05-12 17:20:59', '2020-05-16 19:10:25', NULL);
INSERT INTO `novel_book` VALUES (2, '生生不息2', '', 1, 1, 'werwerwe', '/upload/images/1589549130.jpg', '张真人', '网络', 0, '', 1, '0,大神,精选,爽文,会员', 0, 0, 0, 0, 0, '2020-05-16 10:14:07', 0, '', 0, 0, 0, 0, '2020-05-12 17:24:13', '2020-05-16 19:01:40', NULL);
INSERT INTO `novel_book` VALUES (3, '高冷夫君：傲娇娘子不伺候', '高冷夫君：傲娇娘子不伺候', 2, 3, '一年前，二人初遇，她是他眼中那个纠缠不休的花痴女子。\n一年后，不经意间，她成了他心中那个求而不得的心仪之人。\n且看高冷的他，和傲娇的她，如何相知相爱，相爱相杀。', '/upload/images/1589641476.jpg', '执笔依旧', '网络', 0, '', 0, '', 2, 2, 2, 0, 0, NULL, 0, '', 2, 2, 0, 0, '2020-05-16 23:08:46', '2020-05-16 23:08:46', NULL);

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
  `deleted_at` timestamp(0) NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `udx_link_source`(`source`, `link`) USING BTREE,
  INDEX `idx_novid`(`novel_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '小说采集站点链接关联表' ROW_FORMAT = Dynamic;

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
  `deleted_at` timestamp(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `udx_novid_no`(`novel_id`, `chapter_no`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of novel_chapter_0000
-- ----------------------------
INSERT INTO `novel_chapter_0000` VALUES (1, 1, 1, '第一章 自强不息', '<p>第一章 自强不息</p>', '', 1, '', 0, 0, 0, 0, '2020-05-14 23:08:35', '2020-05-14 23:19:49', NULL);
INSERT INTO `novel_chapter_0000` VALUES (2, 1, 1, '第二章 星星之火', '<p>第二章 星星之火1</p>', '', 1, '', 0, 0, 0, 0, '2020-05-14 23:17:06', '2020-05-16 16:16:31', NULL);

-- ----------------------------
-- Table structure for novel_comment
-- ----------------------------
DROP TABLE IF EXISTS `novel_comment`;
CREATE TABLE `novel_comment`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `novel_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '小说ID',
  `username` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户',
  `content` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '内容',
  `likes` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '点赞数',
  `source` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '来源',
  `created_at` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0),
  `updated_at` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0),
  `deleted_at` timestamp(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of novel_comment
-- ----------------------------
INSERT INTO `novel_comment` VALUES (1, 1, '梧桐', '你好啊', 1, '2', '2020-05-16 20:58:39', '2020-05-16 20:58:39', NULL);
INSERT INTO `novel_comment` VALUES (2, 2, 'coso', '真好看', 34, '2', '2020-05-16 21:00:40', '2020-05-16 21:00:40', NULL);

SET FOREIGN_KEY_CHECKS = 1;
