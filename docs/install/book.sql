/*
SQLyog Enterprise v13.1.1 (64 bit)
MySQL - 5.6.28-log : Database - book
*********************************************************************
*/

/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/`book` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci */;

USE `book`;

/*Table structure for table `book` */

DROP TABLE IF EXISTS `book`;

CREATE TABLE `book` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `name` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT 'NULL' COMMENT '书本名称',
  `author` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT 'NULL' COMMENT '作者',
  `image` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '图片',
  `status` varchar(2) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '小说状态',
  `url` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '采集地址',
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `last_update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='书库';

/*Data for the table `book` */

insert  into `book`(`id`,`name`,`author`,`image`,`status`,`url`,`create_time`,`last_update_time`) values 
(1,'NULL','NULL',NULL,'0','https://www.qu.la/book/3987/','2018-10-17 09:23:04','2018-10-17 09:23:04'),
(2,'NULL','NULL',NULL,'0','http://www.800xs.net/book_72996/','2018-11-13 12:22:58','2018-11-13 12:22:58'),
(3,'NULL','NULL',NULL,'0','https://www.xbaquge.com/files/article/html/40/40670/','2019-08-02 23:49:35','2019-08-02 23:49:35');

/*Table structure for table `chapter` */

DROP TABLE IF EXISTS `chapter`;

CREATE TABLE `chapter` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `book_id` int(11) DEFAULT '0' COMMENT '书本id',
  `title` varchar(100) DEFAULT 'NULL' COMMENT '章节名称',
  `content` text COMMENT '内容',
  `status` tinyint(2) DEFAULT '0' COMMENT '状态',
  `volume` varchar(100) DEFAULT 'NULL' COMMENT '卷',
  `sort` int(11) DEFAULT '0' COMMENT '章节',
  `pre` int(11) DEFAULT '0' COMMENT '上一章',
  `next` int(11) DEFAULT '0' COMMENT '下一章',
  `url` varchar(100) DEFAULT 'NULL' COMMENT '章节url',
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `last_update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8 COMMENT='采集章节';

/*Table structure for table `task` */

DROP TABLE IF EXISTS `task`;

CREATE TABLE `task` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `task_name` varchar(64) NOT NULL COMMENT '任务名称',
  `task_rule_name` varchar(64) NOT NULL COMMENT '任务规则名',
  `task_desc` varchar(512) NOT NULL DEFAULT '' COMMENT '任务描述',
  `status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '状态 0 未开始 1 开始',
  `counts` int(11) NOT NULL DEFAULT '0' COMMENT '次数',
  `cron_spec` varchar(64) NOT NULL DEFAULT '' COMMENT '定时任务',
  `output_type` varchar(64) NOT NULL COMMENT '导出类型',
  `output_exportdb_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '导出数据',
  `opt_user_agent` varchar(128) NOT NULL DEFAULT '' COMMENT '用户代理',
  `opt_max_depth` int(11) NOT NULL DEFAULT '0' COMMENT '爬虫最大深度',
  `opt_allowed_domains` varchar(512) NOT NULL DEFAULT '' COMMENT '允许访问的域名',
  `opt_url_filters` varchar(512) NOT NULL DEFAULT '' COMMENT 'URL过滤',
  `opt_max_body_size` int(11) NOT NULL DEFAULT '0' COMMENT '最大body值',
  `opt_request_timeout` int(11) NOT NULL DEFAULT '10' COMMENT '请求超时时间',
  `auto_migrate` tinyint(4) NOT NULL DEFAULT '0',
  `limit_enable` tinyint(4) NOT NULL DEFAULT '0' COMMENT '频率限制',
  `limit_domain_regexp` varchar(128) NOT NULL DEFAULT '' COMMENT '域名glob匹配regexp',
  `limit_domain_glob` varchar(128) NOT NULL DEFAULT '' COMMENT '域名glob匹配',
  `limit_delay` int(11) NOT NULL DEFAULT '0' COMMENT '延迟',
  `limit_random_delay` int(11) NOT NULL DEFAULT '0' COMMENT '随机延迟',
  `limit_parallelism` int(11) NOT NULL DEFAULT '0' COMMENT '请求并发度',
  `proxy_urls` varchar(2048) NOT NULL DEFAULT '' COMMENT '代理列表',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_task_name` (`task_name`),
  KEY `idx_created_at` (`created_at`),
  KEY `idx_updated_at` (`updated_at`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8mb4 COMMENT='任务表';

/*Data for the table `task` */

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
