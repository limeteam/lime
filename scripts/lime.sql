/*
 Navicat Premium Data Transfer

 Source Server         : windows-dev
 Source Server Type    : MySQL
 Source Server Version : 50726
 Source Host           : localhost:3306
 Source Schema         : lime

 Target Server Type    : MySQL
 Target Server Version : 50726
 File Encoding         : 65001

 Date: 29/05/2020 23:08:44
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
  `cover` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '小说封面',
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
INSERT INTO `novel_book` VALUES (1, '剑焚九霄', '剑焚九霄', 2, 2, '这是一个充满元素气息的世界—鸿界。各种元素游荡在世界的各个角落，在这环境下孕育而生的生命经过漫长时间的演化，生活，学习，争夺，繁衍下对于各种元素的掌握到达了一个巅峰，迎来一个大时代—三极天下。鸿界分为...', 'http://beiwo-new.oss-cn-beijing.aliyuncs.com/cover/84/7757bb60adec76a2166f8af2cde68117.jpeg?x-oss-process=image%2Fresize%2Cw_300%2Ch_400%2Cm_lfit', '孤梦夕阳', '网络', 0, '', 0, '新书,会员,热门', 2, 4, 2, 2542452, 100, '2020-05-16 10:13:55', 0, '', 100000, 0, 0, 0, '2020-05-12 17:20:59', '2020-05-16 19:10:25', NULL);
INSERT INTO `novel_book` VALUES (2, '生生不息2', '', 1, 1, '十年前，我拿剪刀戳伤了他的眉心，我被关了半个月的少改所。十年后，他指着我说我下贱，勾引了他的弟弟。我和他相看两厌，他厌恶我，我痛恨他。本以为就这样一辈子在无交集。一天，他却装醉欺辱了我...我恨他，致死！', '/upload/images/1589549130.jpg', '张真人', '网络', 0, '', 1, '0,大神,精选,爽文,会员', 0, 0, 3, 0, 34, '2020-05-16 10:14:07', 0, '', 0, 0, 0, 0, '2020-05-12 17:24:13', '2020-05-16 19:01:40', NULL);
INSERT INTO `novel_book` VALUES (3, '高冷夫君：傲娇娘子不伺候', '高冷夫君：傲娇娘子不伺候', 2, 3, '一年前，二人初遇，她是他眼中那个纠缠不休的花痴女子。\n一年后，不经意间，她成了他心中那个求而不得的心仪之人。\n且看高冷的他，和傲娇的她，如何相知相爱，相爱相杀。', '/upload/images/1589641476.jpg', '执笔依旧', '网络', 0, '', 0, '', 2, 2, 2, 0, 245, '2020-05-26 23:49:32', 0, '', 2, 2, 0, 0, '2020-05-16 23:08:46', '2020-05-16 23:08:46', NULL);

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
INSERT INTO `novel_chapter_0000` VALUES (1, 1, 1, '第一章 自强不息', '\\t    明天就是男朋友秦牧扬的生日，我漂洋过海来给他过生日，本想给他一个惊喜，却未曾想惊喜倒没有，倒是给自己一个大大的惊吓。n\\n\\tn\\n\\t    下了飞机，坐上一辆出租车，人若倒霉喝凉水都塞牙，出租车走半道翻车了，擦伤了手臂，被警察送往医院的时候，我反应过来正想给秦牧扬打电话时，迎面走来一个男人搀扶着一个孕妇。n\\n\\tn\\n\\t    大着肚子挺的很高一看就像是至少有五个多月的身子。n\\n\\tn\\n\\t    我现在关注的不是那个孕妇。n\\n\\tn\\n\\t    而是孕妇身边的那个男人，那个男人不是别人，正是我的男朋友秦牧扬。n\\n\\tn\\n\\t    秦牧扬对那个孕妇小心翼翼的模样，女人的第六感告诉我秦牧扬和那个女人，关系绝不简单。n\\n\\tn\\n\\t    我走到秦牧扬的面前，尽量平复自己的心情，声音还算正常“你这是在干嘛！”n\\n\\tn\\n\\t    秦牧扬一脸意外的模样，他很意外我怎么会出现在这里。n\\n\\tn\\n\\t    震惊失措躲闪，再看看他身边的这个女人，挑衅的看着我，我的心里隐隐的像是明白了什么。n\\n\\tn\\n\\t    “秦牧扬，我在问你话呢？”n\\n\\tn\\n\\t    我的口气有些咄咄逼人。n\\n\\tn\\n\\t    秦牧扬没有立马就回答我的话，他看了看身边的女人一眼，就拉着我的手走，走了大概几十米米远我甩开他的的手，面部狰狞像个疯子一样愤怒的吼道“那个女人跟你是什么关系，他肚子里的孩子是你的吗？”n\\n\\tn\\n\\t    他眉头紧皱，眼神里闪烁着愧疚。n\\n\\tn\\n\\t    我和他相识18年，他的一个眼神，我就知道他心里的想法是什么。n\\n\\tn\\n\\t    “木子，对不起！”n\\n\\tn\\n\\t    秦牧扬跟我说了五个字，五个字对于我来说却像是天打雷劈一般。n\\n\\tn\\n\\t    他觉得我伤的不够深似的，顿了几秒又道“木子，我们分手吧！”n\\n\\tn\\n\\t    我踉跄的后退了几步，似乎不敢相信自己耳朵里听见的话。n\\n\\tn\\n\\t    “啪！”我重重的一巴掌甩在他的脸上，他没有丝毫躲闪的意思，接了这巴掌。n\\n\\tn\\n\\t    “你……你……欺人太甚！”我再也受不了眼泪像是决了堤的洪水似的一涌而出。n\\n\\tn\\n\\t    我们之间的恋爱关系，建立时间也不过三个月的时间而已，而那个女人的肚子却有五个多月大，也就是说他跟别的女人都有孩子了，三个月前他却跟我表白，让我做他的女朋友。n\\n\\tn\\n\\t    原来我就像是一个傻子一样被他耍的团团转。n\\n\\tn\\n\\t    他是我最爱是人啊，他是这十八年来里，对我最好的人啊，可如今却欺骗了我，玩弄了我。我此时此刻已经无法去形容我心中的愤怒，手不受控制的抬起重重的一巴掌再次的甩在他脸上。n\\n\\tn\\n\\t    同样，他依然没有躲开。n\\n\\tn\\n\\t    倒是那个女人过来，上来就狠狠的给了我一巴掌。n\\n\\tn\\n\\t    “我和他是在你之前，你有什么资格打他，要是论三儿，那个小三儿也是你。”n\\n\\tn\\n\\t    那个女人说我是她和秦牧扬之间的小三儿吗？n\\n\\tn\\n\\t    我捂着自己的脸，秦牧扬也震惊这个女人打了我。n\\n\\tn\\n\\t    “魏冉你在干什么，我们之间的事情不需要你来管，你先回去。”他吼着那个女人。n\\n\\tn\\n\\t    那个叫魏冉的女人见秦牧扬对她发了火，也不打算跟我纠缠什么就要走，我拽着她的手肯定不让她走“事情没有说清楚，谁都别想走。”n\\n\\tn\\n\\t    “好啊！那我不介意跟李小姐说个清楚道个明白，让你清楚谁才是小三儿。”n\\n\\tn\\n\\t    魏冉也不是什么善类。n\\n\\tn\\n\\t    “够了，魏冉你闭嘴。”秦牧扬再次朝那个女人发火。n\\n\\tn\\n\\t    “木子，对不起，这事儿我会慢慢跟你解释清楚，你现在住哪家酒店，我送你回去。”n\\n\\tn\\n\\t    秦牧扬脸上疲倦的神色，他累了，显然是两个女人折腾的他身心疲倦。n\\n\\tn\\n\\t    “我要你的解释，现在！”n\\n\\tn\\n\\t    此刻，我愤怒得像是失去了理智，像是一个动物一般只想着发泄，我一把甩开魏冉的手，双手双脚就要往他身上发泄去。n\\n\\tn\\n\\t    他不躲，那个女人上来护着他，我愤怒到没有丝毫理智，一把推开她。n\\n\\tn\\n\\t    “啊！”n\\n\\tn\\n\\t    女人一声儿尖叫。n\\n\\tn\\n\\t    “魏冉魏冉，你怎么了？”n\\n\\tn\\n\\t    “牧扬，我肚子好疼好疼，我可能要流产了，牧扬我们的孩子。”n\\n\\tn\\n\\t    他拽着我的手，就将我甩在地上跪着。n\\n\\tn\\n\\t    我看着他着急的神色，抱起那个女人，他手上都是血，他疯狂的往医院里冲。n\\n\\tn\\n\\t    我的膝盖被碎石子扎破了，在疼也不及心疼。n\\n\\tn\\n\\t    那个女人趟过的地方大片的血，我反应过来，看着自己的双手才意识到自己刚才都做了什么。n\\n\\tn\\n\\t    医院的走廊里，秦牧扬依靠在墙壁上，一根又一根的香烟抽着，浑身散发着颓废的气息。n\\n\\tn\\n\\t    他在国外留学这几年，每年我都要去看他几次，竟然不知道他什么时候染上了抽烟的恶习。n\\n\\tn\\n\\t    通过他简单的叙述我知道了他和那个魏冉是怎么回事，他和魏冉之前参加过一次聚会，不小心发生了一夜情，魏冉怀孕了却没有告诉他，等肚子都三四个月了才跟他说。n\\n\\tn\\n\\t    也就是说，他跟我表白的时候，并不知道那个女人怀了他的孩子。n\\n\\tn\\n\\t    他没有欺骗我，也没有玩弄我。n\\n\\tn\\n\\t    “木子，你回去吧！是我对不住你。”n\\n\\tn\\n\\t    他抽完又一根烟，抬眸对我说。n\\n\\tn\\n\\t    他还是要跟我分手，可是我不同意，我怎么会同意呢，我是那么的爱他啊！n\\n\\tn\\n\\t    我爱他已经爱到骨子里了，我不能没有他。n\\n\\tn\\n\\t    我听见自己的声音都在颤抖“你……可以让她把这个孩子生下来，也可以和她一起抚养这个孩子，我不介意的，我也会帮着你一起照顾，甚至我可以不生自己的孩子，只要你娶我别娶她！好吗？”我蹲在他面前，将脑袋埋在他的腿上，卑微的祈求。n\\n\\tn\\n\\t    我从来没有这般卑微过。n\\n\\tn\\n\\t    他知道，我是个怎样的性子。n\\n\\tn\\n\\t    这样的卑微我从未有过。n\\n\\tn\\n\\t    秦牧扬抬起我的头，震惊的看着我，他的手指轻柔的捻去我眼角的泪“木子，对不起，我不值得你这样，不值得！”n\\n\\tn\\n\\t    “值得，值得的，你可知道我有多爱你，你比谁都知道，我等了你这么多年，终于等到你的表白，我不想放弃你，求你了，不要，别不要我好吗！”n\\n\\tn\\n\\t    我也也没有了刚才打人的那种张牙舞爪的气势，现在有的只是脆弱。n\\n\\tn\\n\\t    我怕这个男人不要我。n\\n\\tn\\n\\t    他轻轻将我推开。n\\n\\tn\\n\\t    我不甘心，从后面抱住他的腰，将脸埋在他的后背不停的蹭着“我不介意，我真的不介意！”n\\n\\tn\\n\\t    可是他还是掰开了我的手“就算没有魏然肚子里的孩子，我们也走不到最后，别忘了我们是兄妹！”n\\n\\t', '', 1, '', 0, 0, 0, 0, '2020-05-14 23:08:35', '2020-05-14 23:19:49', NULL);
INSERT INTO `novel_chapter_0000` VALUES (2, 1, 2, '第二章 星星之火', '<p>第二章 星星之火1</p>', '', 1, '', 0, 0, 0, 0, '2020-05-14 23:17:06', '2020-05-16 16:16:31', NULL);

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

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '用户名称',
  `mobile` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '手机号',
  `sex` int(11) NOT NULL DEFAULT 0 COMMENT '性别',
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '密码',
  `salt` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '加密盐',
  `faceicon` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '头像地址',
  `wechat` json NULL COMMENT '微信绑定信息',
  `email` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '邮箱',
  `amount` decimal(10, 0) NULL DEFAULT NULL COMMENT '现金金额',
  `coin` decimal(10, 0) NULL DEFAULT NULL COMMENT '金币',
  `exempt_login` tinyint(2) NULL DEFAULT NULL COMMENT '是否免登： 0 否 1免登',
  `source` tinyint(2) NULL DEFAULT NULL COMMENT '来源：0 PC 1 WAP 2 Android 3 IOS 4 公众号 5小程序 ',
  `is_vip` tinyint(2) NULL DEFAULT NULL COMMENT '是否VIP  0 否 1是',
  `channel_id` int(11) NULL DEFAULT NULL COMMENT '渠道',
  `status` int(11) NOT NULL DEFAULT 1 COMMENT '状态 1=正常 2=禁用',
  `create_time` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0),
  `last_login_time` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0),
  `deleted_at` timestamp(0) NULL DEFAULT NULL COMMENT '删除时间',
  `created_at` datetime(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0),
  `updated_at` datetime(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0),
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '用户表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of users
-- ----------------------------
INSERT INTO `users` VALUES (1, 'coso', '1862071111', 1, 'werwerwrw', '223', '', NULL, 'wutongci@hotmail.com', 50, 100, NULL, 2, NULL, 1, 1, '2020-05-22 23:36:33', '2020-05-22 23:36:33', NULL, '2020-05-22 23:58:22', '2020-05-22 23:58:22');

SET FOREIGN_KEY_CHECKS = 1;
