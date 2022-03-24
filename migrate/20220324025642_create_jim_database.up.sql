CREATE DATABASE IF NOT EXISTS `jim`;
USE `jim`;

CREATE TABLE IF NOT EXISTS `device` (
`id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '自增主键',
`device_id` bigint NOT NULL COMMENT '设备id',
`app_id` bigint unsigned NOT NULL COMMENT 'app_id',
`user_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '账户id',
`type` tinyint NOT NULL COMMENT '设备类型,1:Android；2：IOS；3：Windows; 4：MacOS；5：Web',
`brand` varchar(20)  NOT NULL COMMENT '手机厂商',
`model` varchar(20)  NOT NULL COMMENT '机型',
`system_version` varchar(10)  NOT NULL COMMENT '系统版本',
`sdk_version` varchar(10)  NOT NULL COMMENT 'app版本',
`status` tinyint NOT NULL DEFAULT '0' COMMENT '在线状态，0：离线；1：在线',
`conn_addr` varchar(25)  NOT NULL COMMENT '连接层服务器地址',
`conn_fd` bigint NOT NULL COMMENT 'TCP连接对应的文件描述符',
`create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
`update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
PRIMARY KEY (`id`),
UNIQUE KEY `uk_device_id` (`device_id`) USING BTREE,
KEY `idx_app_id_user_id` (`app_id`,`user_id`) USING BTREE
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4  COMMENT='设备';


REPLACE INTO `device` (`id`, `device_id`, `app_id`, `user_id`, `type`, `brand`, `model`, `system_version`, `sdk_version`, `status`, `conn_addr`, `conn_fd`, `create_time`, `update_time`) VALUES
(1, 1, 1, 1, 2, 'apple', ' ', ' ', ' ', 0, '127.0.0.1', 1, '2021-11-09 12:04:18', '2021-11-09 12:04:21');

CREATE TABLE IF NOT EXISTS `device_ack` (
`id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '自增主键',
`device_id` bigint unsigned NOT NULL COMMENT '设备id',
`ack` bigint unsigned NOT NULL DEFAULT '0' COMMENT '收到消息确认号',
`create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
`update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
PRIMARY KEY (`id`) USING BTREE,
UNIQUE KEY `uk_device_id` (`device_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4  COMMENT='设备消息同步序列号';

CREATE TABLE IF NOT EXISTS `feed` (
`id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '动态 ID',
`user_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '用户ID',
`content` text NOT NULL COMMENT '动态内容',
`type` tinyint NOT NULL DEFAULT '0' COMMENT '动态类型',
`like_count` int unsigned NOT NULL DEFAULT '0' COMMENT '动态点赞数',
`view_count` int unsigned NOT NULL DEFAULT '0' COMMENT '动态阅读数',
`comment_count` int unsigned NOT NULL DEFAULT '0' COMMENT '动态评论数',
`operator` bigint unsigned NOT NULL DEFAULT '0' COMMENT '审核人',
`remark` varchar(500)NOT NULL DEFAULT '' COMMENT '备注',
`hot` bigint unsigned NOT NULL DEFAULT '0' COMMENT '热门排序值',
`is_enable` tinyint unsigned NOT NULL DEFAULT '1' COMMENT '是否启用',
`review_status` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '审核状态 0-未审核 1-已审核 2-未通过',
`created_at` datetime DEFAULT NULL,
`updated_at` datetime DEFAULT NULL,
`deleted_at` datetime DEFAULT NULL,
PRIMARY KEY (`id`),
KEY `company_id_is_enable_review_status` (`user_id`,`is_enable`,`review_status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4  COMMENT='动态';


CREATE TABLE IF NOT EXISTS `feed_image` (
`id` bigint unsigned NOT NULL DEFAULT '0' COMMENT 'ID',
`feed_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '动态ID',
`img_url` varchar(2000)NOT NULL DEFAULT '' COMMENT '图片URL',
`created_at` datetime DEFAULT NULL,
`updated_at` datetime DEFAULT NULL,
`deleted_at` datetime DEFAULT NULL,
PRIMARY KEY (`id`),
KEY `feeds_user_id_index` (`feed_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4  COMMENT='动态图片';

CREATE TABLE IF NOT EXISTS `feed_like` (
`id` bigint unsigned NOT NULL DEFAULT '0' COMMENT 'ID',
`feed_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '动态ID',
`user_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '用户ID',
`created_at` datetime DEFAULT NULL,
`updated_at` datetime DEFAULT NULL,
`deleted_at` datetime DEFAULT NULL,
PRIMARY KEY (`id`),
KEY `feeds_user_id_index` (`feed_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4  COMMENT='动态点赞';


CREATE TABLE IF NOT EXISTS `feed_video` (
`id` int unsigned NOT NULL AUTO_INCREMENT,
`feed_id` int unsigned NOT NULL DEFAULT '0' COMMENT '所属动态id',
`video_url` varchar(2048)NOT NULL DEFAULT '' COMMENT '视频URL',
`cover_url` varchar(2048)NOT NULL DEFAULT '' COMMENT '视频封面URL',
`height` int unsigned NOT NULL DEFAULT '0' COMMENT '视频高度',
`width` int unsigned NOT NULL DEFAULT '0' COMMENT '视频宽度',
`duration` float unsigned NOT NULL DEFAULT '0' COMMENT '视频时长',
`created_at` datetime DEFAULT NULL,
`updated_at` datetime DEFAULT NULL,
`deleted_at` datetime DEFAULT NULL,
PRIMARY KEY (`id`),
KEY `feed_videos_feed_id_index` (`feed_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4  COMMENT='动态视频';


CREATE TABLE IF NOT EXISTS `gid` (
 `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '自增主键',
 `business_id` varchar(128)  NOT NULL COMMENT '业务id',
`max_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '最大id',
`step` int unsigned NOT NULL DEFAULT '1000' COMMENT '步长',
`description` varchar(255)  NOT NULL COMMENT '描述',
`create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
`update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
PRIMARY KEY (`id`),
UNIQUE KEY `uk_business_id` (`business_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4  COMMENT='分布式自增主键';


CREATE TABLE IF NOT EXISTS `goadmin_menu` (
`id` int unsigned NOT NULL AUTO_INCREMENT,
`parent_id` int unsigned NOT NULL DEFAULT '0',
`type` tinyint unsigned NOT NULL DEFAULT '0',
`order` int unsigned NOT NULL DEFAULT '0',
`title` varchar(50)NOT NULL,
`icon` varchar(50)NOT NULL,
`uri` varchar(3000)NOT NULL DEFAULT '',
`header` varchar(150)DEFAULT NULL,
`plugin_name` varchar(150)NOT NULL DEFAULT '',
`uuid` varchar(150)DEFAULT NULL,
`created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
`updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
PRIMARY KEY (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 ;

-- Dumping data for table jim.goadmin_menu: ~7 rows (大约)
/*!40000 ALTER TABLE `goadmin_menu` DISABLE KEYS */;
REPLACE INTO `goadmin_menu` (`id`, `parent_id`, `type`, `order`, `title`, `icon`, `uri`, `header`, `plugin_name`, `uuid`, `created_at`, `updated_at`) VALUES
(1, 0, 1, 2, 'Admin', 'fa-tasks', '', NULL, '', NULL, '2019-09-10 00:00:00', '2019-09-10 00:00:00'),
(2, 1, 1, 2, 'Users', 'fa-users', '/info/manager', NULL, '', NULL, '2019-09-10 00:00:00', '2019-09-10 00:00:00'),
(3, 1, 1, 3, 'Roles', 'fa-user', '/info/roles', NULL, '', NULL, '2019-09-10 00:00:00', '2019-09-10 00:00:00'),
(4, 1, 1, 4, 'Permission', 'fa-ban', '/info/permission', NULL, '', NULL, '2019-09-10 00:00:00', '2019-09-10 00:00:00'),
(5, 1, 1, 5, 'Menu', 'fa-bars', '/menu', NULL, '', NULL, '2019-09-10 00:00:00', '2019-09-10 00:00:00'),
(6, 1, 1, 6, 'Operation log', 'fa-history', '/info/op', NULL, '', NULL, '2019-09-10 00:00:00', '2019-09-10 00:00:00'),
(7, 0, 1, 1, 'Dashboard', 'fa-bar-chart', '/', NULL, '', NULL, '2019-09-10 00:00:00', '2019-09-10 00:00:00'),
(8, 0, 0, 2, '用户', 'fa-blind', '/user', '', '', NULL, '2022-01-26 09:24:29', '2022-01-26 09:24:29');

CREATE TABLE IF NOT EXISTS `goadmin_operation_log` (
`id` int unsigned NOT NULL AUTO_INCREMENT,
`user_id` int unsigned NOT NULL,
`path` varchar(255)NOT NULL,
`method` varchar(10)NOT NULL,
`ip` varchar(15)NOT NULL,
`input` text NOT NULL,
`created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
`updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
PRIMARY KEY (`id`),
KEY `admin_operation_log_user_id_index` (`user_id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 ;

REPLACE INTO `goadmin_operation_log` (`id`, `user_id`, `path`, `method`, `ip`, `input`, `created_at`, `updated_at`) VALUES
(1, 1, '/admin', 'GET', '127.0.0.1', '', '2022-01-26 09:23:10', '2022-01-26 09:23:10'),
(2, 1, '/admin/info/roles', 'GET', '127.0.0.1', '', '2022-01-26 09:23:18', '2022-01-26 09:23:18'),
(3, 1, '/admin/menu', 'GET', '127.0.0.1', '', '2022-01-26 09:23:32', '2022-01-26 09:23:32'),
(4, 1, '/admin/menu/new', 'POST', '127.0.0.1', '{"__go_admin_previous_":["/admin/menu"],"__go_admin_t_":["e3e75eee-f679-4c22-9365-641d6c280396"],"header":[""],"icon":["fa-blind"],"parent_id":["0"],"plugin_name":[""],"title":["用户"],"uri":["/user"]}', '2022-01-26 09:24:29', '2022-01-26 09:24:29'),
(5, 1, '/admin', 'GET', '127.0.0.1', '', '2022-01-26 09:26:52', '2022-01-26 09:26:52'),
(6, 1, '/admin/info/manager', 'GET', '127.0.0.1', '', '2022-01-26 09:27:03', '2022-01-26 09:27:03'),
(7, 1, '/admin/info/roles', 'GET', '127.0.0.1', '', '2022-01-26 09:28:50', '2022-01-26 09:28:50'),
(8, 1, '/admin', 'GET', '127.0.0.1', '', '2022-01-26 09:34:01', '2022-01-26 09:34:01'),
(9, 1, '/admin/application/info', 'GET', '127.0.0.1', '', '2022-01-26 09:34:15', '2022-01-26 09:34:15'),
(10, 1, '/admin/application/info', 'GET', '127.0.0.1', '', '2022-01-26 09:34:21', '2022-01-26 09:34:21'),
(11, 1, '/admin/plugins', 'GET', '127.0.0.1', '', '2022-01-26 09:34:30', '2022-01-26 09:34:30'),
(12, 1, '/admin/plugins/store', 'GET', '127.0.0.1', '', '2022-01-26 09:34:39', '2022-01-26 09:34:39'),
(13, 1, '/admin/plugins', 'GET', '127.0.0.1', '', '2022-01-26 09:34:55', '2022-01-26 09:34:55'),
(14, 1, '/admin', 'GET', '127.0.0.1', '', '2022-01-26 09:37:12', '2022-01-26 09:37:12'),
(15, 1, '/admin/plugins', 'GET', '127.0.0.1', '', '2022-01-26 09:37:41', '2022-01-26 09:37:41'),
(16, 1, '/admin/plugins/store', 'GET', '127.0.0.1', '', '2022-01-26 09:37:43', '2022-01-26 09:37:43'),
(17, 1, '/admin/info/site/edit', 'GET', '127.0.0.1', '', '2022-01-26 09:37:46', '2022-01-26 09:37:46'),
(18, 1, '/admin/info/generate/new', 'GET', '127.0.0.1', '', '2022-01-26 09:37:54', '2022-01-26 09:37:54'),
(19, 1, '/admin/application/info', 'GET', '127.0.0.1', '', '2022-01-26 09:37:56', '2022-01-26 09:37:56'),
(20, 1, '/admin/application/info', 'GET', '127.0.0.1', '', '2022-01-26 09:37:59', '2022-01-26 09:37:59'),
(21, 1, '/admin', 'GET', '127.0.0.1', '', '2022-01-26 09:38:02', '2022-01-26 09:38:02'),
(22, 1, '/admin', 'GET', '127.0.0.1', '', '2022-01-26 09:38:04', '2022-01-26 09:38:04'),
(23, 1, '/admin/info/op', 'GET', '127.0.0.1', '', '2022-01-26 09:38:08', '2022-01-26 09:38:08'),
(24, 1, '/admin/menu', 'GET', '127.0.0.1', '', '2022-01-26 09:38:09', '2022-01-26 09:38:09'),
(25, 1, '/admin/info/permission', 'GET', '127.0.0.1', '', '2022-01-26 09:38:10', '2022-01-26 09:38:10'),
(26, 1, '/admin/info/roles', 'GET', '127.0.0.1', '', '2022-01-26 09:38:11', '2022-01-26 09:38:11'),
(27, 1, '/admin/info/op', 'GET', '127.0.0.1', '', '2022-01-26 09:38:13', '2022-01-26 09:38:13'),
(28, 1, '/admin', 'GET', '127.0.0.1', '', '2022-01-26 09:38:15', '2022-01-26 09:38:15'),
(29, 1, '/admin', 'GET', '127.0.0.1', '', '2022-01-26 09:52:51', '2022-01-26 09:52:51'),
(30, 1, '/admin/info/user', 'GET', '127.0.0.1', '', '2022-01-26 09:53:06', '2022-01-26 09:53:06'),
(31, 1, '/admin/info/user', 'GET', '127.0.0.1', '', '2022-01-26 09:53:33', '2022-01-26 09:53:33'),
(32, 1, '/admin/info/user', 'GET', '127.0.0.1', '', '2022-01-26 09:53:35', '2022-01-26 09:53:35'),
(33, 1, '/admin/info/group_user', 'GET', '127.0.0.1', '', '2022-01-26 09:53:56', '2022-01-26 09:53:56'),
(34, 1, '/admin', 'GET', '::1', '', '2022-01-26 09:55:44', '2022-01-26 09:55:44'),
(35, 1, '/admin/info/group', 'GET', '::1', '', '2022-01-26 09:55:55', '2022-01-26 09:55:55'),
(36, 1, '/admin/info/group', 'GET', '::1', '', '2022-01-26 09:59:18', '2022-01-26 09:59:18'),
(37, 1, '/admin/info/group', 'GET', '::1', '', '2022-01-26 09:59:25', '2022-01-26 09:59:25'),
(38, 1, '/admin/info/user', 'GET', '::1', '', '2022-01-26 09:59:29', '2022-01-26 09:59:29'),
(39, 1, '/admin/info/user/edit', 'GET', '::1', '', '2022-01-26 09:59:33', '2022-01-26 09:59:33'),
(40, 1, '/admin/info/user', 'GET', '::1', '', '2022-01-26 09:59:41', '2022-01-26 09:59:41'),
(41, 1, '/admin/info/goup', 'GET', '::1', '', '2022-01-26 09:59:45', '2022-01-26 09:59:45'),
(42, 1, '/admin/info/goup_a', 'GET', '::1', '', '2022-01-26 09:59:52', '2022-01-26 09:59:52'),
(43, 1, '/admin/info/user', 'GET', '::1', '', '2022-01-26 10:00:05', '2022-01-26 10:00:05'),
(44, 1, '/admin/info/message', 'GET', '::1', '', '2022-01-26 10:00:08', '2022-01-26 10:00:08'),
(45, 1, '/admin/info/group', 'GET', '::1', '', '2022-01-26 10:00:11', '2022-01-26 10:00:11'),
(46, 1, '/admin/info/group', 'GET', '::1', '', '2022-01-26 10:00:32', '2022-01-26 10:00:32'),
(47, 1, '/admin/info/group', 'GET', '::1', '', '2022-01-26 10:00:53', '2022-01-26 10:00:53'),
(48, 1, '/admin/info/group', 'GET', '::1', '', '2022-01-26 10:02:42', '2022-01-26 10:02:42'),
(49, 1, '/admin/info/group', 'GET', '::1', '', '2022-01-26 10:03:46', '2022-01-26 10:03:46'),
(50, 1, '/admin/info/user', 'GET', '::1', '', '2022-01-26 10:04:39', '2022-01-26 10:04:39'),
(51, 1, '/admin/info/user/edit', 'GET', '::1', '', '2022-01-26 10:04:45', '2022-01-26 10:04:45'),
(52, 1, '/admin/info/user', 'GET', '::1', '', '2022-01-26 10:06:01', '2022-01-26 10:06:01'),
(53, 1, '/admin/info/user', 'GET', '::1', '', '2022-01-26 10:08:29', '2022-01-26 10:08:29'),
(54, 1, '/admin/info/user/edit', 'GET', '::1', '', '2022-01-26 10:08:32', '2022-01-26 10:08:32'),
(55, 1, '/admin/info/group', 'GET', '::1', '', '2022-01-26 10:10:35', '2022-01-26 10:10:35'),
(56, 1, '/admin/info/user', 'GET', '::1', '', '2022-01-26 10:11:19', '2022-01-26 10:11:19'),
(57, 1, '/admin/info/user/detail', 'GET', '::1', '', '2022-01-26 10:33:23', '2022-01-26 10:33:23'),
(58, 1, '/admin/info/user', 'GET', '::1', '', '2022-01-26 10:33:27', '2022-01-26 10:33:27'),
(59, 1, '/admin/info/user', 'GET', '::1', '', '2022-01-26 10:35:45', '2022-01-26 10:35:45'),
(60, 1, '/admin/info/user', 'GET', '::1', '', '2022-01-26 10:37:14', '2022-01-26 10:37:14'),
(61, 1, '/admin/export/user', 'POST', '::1', '', '2022-01-26 10:37:33', '2022-01-26 10:37:33'),
(62, 1, '/admin/info/user/detail', 'GET', '::1', '', '2022-01-26 10:38:24', '2022-01-26 10:38:24'),
(63, 1, '/admin/info/user/edit', 'GET', '::1', '', '2022-01-26 10:38:28', '2022-01-26 10:38:28'),
(64, 1, '/admin/info/user/detail', 'GET', '::1', '', '2022-01-26 10:38:29', '2022-01-26 10:38:29'),
(65, 1, '/admin/info/user', 'GET', '::1', '', '2022-01-26 10:39:34', '2022-01-26 10:39:34'),
(66, 1, '/admin/info/user', 'GET', '::1', '', '2022-01-26 10:43:19', '2022-01-26 10:43:19'),
(67, 1, '/admin/export/user', 'POST', '::1', '', '2022-01-26 10:43:23', '2022-01-26 10:43:23'),
(68, 1, '/admin/info/user', 'GET', '::1', '', '2022-01-26 10:51:02', '2022-01-26 10:51:02'),
(69, 1, '/admin/plugins', 'GET', '::1', '', '2022-01-26 10:51:06', '2022-01-26 10:51:06'),
(70, 1, '/admin/plugins/store', 'GET', '::1', '', '2022-01-26 10:51:09', '2022-01-26 10:51:09'),
(71, 1, '/admin/plugin/download', 'POST', '::1', '', '2022-01-26 10:51:14', '2022-01-26 10:51:14'),
(72, 1, '/admin/plugin/download', 'POST', '::1', '', '2022-01-26 10:51:21', '2022-01-26 10:51:21'),
(73, 1, '/admin', 'GET', '::1', '', '2022-01-26 10:58:48', '2022-01-26 10:58:48'),
(74, 1, '/admin/plugins/store', 'GET', '::1', '', '2022-01-26 10:58:51', '2022-01-26 10:58:51'),
(75, 1, '/admin/plugins/store', 'GET', '::1', '', '2022-01-26 10:58:53', '2022-01-26 10:58:53'),
(76, 1, '/admin/info/plugin_filemanager/new', 'GET', '::1', '', '2022-01-26 10:58:54', '2022-01-26 10:58:54'),
(77, 1, '/admin/info/plugin_filemanager/new', 'GET', '::1', '', '2022-01-26 10:58:55', '2022-01-26 10:58:55'),
(78, 1, '/admin/plugins/store', 'GET', '::1', '', '2022-01-26 10:59:02', '2022-01-26 10:59:02'),
(79, 1, '/admin/info/plugin_librarian/new', 'GET', '::1', '', '2022-01-26 10:59:03', '2022-01-26 10:59:03'),
(80, 1, '/admin/info/plugin_librarian/new', 'GET', '::1', '', '2022-01-26 10:59:03', '2022-01-26 10:59:03'),
(81, 1, '/admin', 'GET', '127.0.0.1', '', '2022-01-27 07:24:16', '2022-01-27 07:24:16'),
(82, 1, '/admin/info/user', 'GET', '127.0.0.1', '', '2022-01-27 07:24:51', '2022-01-27 07:24:51'),
(83, 1, '/admin/info/group', 'GET', '127.0.0.1', '', '2022-01-27 07:24:57', '2022-01-27 07:24:57'),
(84, 1, '/admin/info/group', 'GET', '127.0.0.1', '', '2022-01-27 07:27:31', '2022-01-27 07:27:31'),
(85, 1, '/admin/info/group', 'GET', '127.0.0.1', '', '2022-01-27 07:29:42', '2022-01-27 07:29:42'),
(86, 1, '/admin/info/group', 'GET', '127.0.0.1', '', '2022-01-27 07:35:07', '2022-01-27 07:35:07'),
(87, 1, '/admin/info/group', 'GET', '127.0.0.1', '', '2022-01-27 07:35:07', '2022-01-27 07:35:07'),
(88, 1, '/admin/info/group', 'GET', '127.0.0.1', '', '2022-01-27 07:39:00', '2022-01-27 07:39:00'),
(89, 1, '/admin/info/group', 'GET', '127.0.0.1', '', '2022-01-27 07:41:52', '2022-01-27 07:41:52'),
(90, 1, '/admin/info/group', 'GET', '127.0.0.1', '', '2022-01-27 07:44:45', '2022-01-27 07:44:45'),
(91, 1, '/admin/info/group', 'GET', '127.0.0.1', '', '2022-01-27 07:45:26', '2022-01-27 07:45:26'),
(92, 1, '/admin/info/group', 'GET', '127.0.0.1', '', '2022-01-27 07:47:32', '2022-01-27 07:47:32'),
(93, 1, '/admin/info/group', 'GET', '127.0.0.1', '', '2022-01-27 07:57:45', '2022-01-27 07:57:45'),
(94, 1, '/admin/info/group', 'GET', '127.0.0.1', '', '2022-01-27 08:00:24', '2022-01-27 08:00:24'),
(95, 1, '/admin/info/group', 'GET', '127.0.0.1', '', '2022-01-27 08:15:46', '2022-01-27 08:15:46'),
(96, 1, '/admin/info/group', 'GET', '127.0.0.1', '', '2022-01-27 08:15:48', '2022-01-27 08:15:48');

CREATE TABLE IF NOT EXISTS `goadmin_permissions` (
 `id` int unsigned NOT NULL AUTO_INCREMENT,
 `name` varchar(50)NOT NULL,
`slug` varchar(50)NOT NULL,
`http_method` varchar(255)DEFAULT NULL,
`http_path` text NOT NULL,
`created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
`updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
PRIMARY KEY (`id`),
UNIQUE KEY `admin_permissions_name_unique` (`name`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 ;


REPLACE INTO `goadmin_permissions` (`id`, `name`, `slug`, `http_method`, `http_path`, `created_at`, `updated_at`) VALUES
(1, 'All permission', '*', '', '*', '2019-09-10 00:00:00', '2019-09-10 00:00:00'),
(2, 'Dashboard', 'dashboard', 'GET,PUT,POST,DELETE', '/', '2019-09-10 00:00:00', '2019-09-10 00:00:00'),
(3, 'device 查询', 'device_query', 'GET', '/info/device', '2022-01-26 09:41:58', '2022-01-26 09:41:58'),
(4, 'device 编辑页显示', 'device_show_edit', 'GET', '/info/device/edit', '2022-01-26 09:41:58', '2022-01-26 09:41:58'),
(5, 'device 新建记录页显示', 'device_show_create', 'GET', '/info/device/new', '2022-01-26 09:41:58', '2022-01-26 09:41:58'),
(6, 'device 编辑', 'device_edit', 'POST', '/edit/device', '2022-01-26 09:41:58', '2022-01-26 09:41:58'),
(7, 'device 新建', 'device_create', 'POST', '/new/device', '2022-01-26 09:41:58', '2022-01-26 09:41:58'),
(8, 'device 删除', 'device_delete', 'POST', '/delete/device', '2022-01-26 09:41:58', '2022-01-26 09:41:58'),
(9, 'device 导出', 'device_export', 'POST', '/export/device', '2022-01-26 09:41:59', '2022-01-26 09:41:59'),
(10, 'device_ack 查询', 'device_ack_query', 'GET', '/info/device_ack', '2022-01-26 09:41:59', '2022-01-26 09:41:59'),
(11, 'device_ack 编辑页显示', 'device_ack_show_edit', 'GET', '/info/device_ack/edit', '2022-01-26 09:41:59', '2022-01-26 09:41:59'),
(12, 'device_ack 新建记录页显示', 'device_ack_show_create', 'GET', '/info/device_ack/new', '2022-01-26 09:41:59', '2022-01-26 09:41:59'),
(13, 'device_ack 编辑', 'device_ack_edit', 'POST', '/edit/device_ack', '2022-01-26 09:41:59', '2022-01-26 09:41:59'),
(14, 'device_ack 新建', 'device_ack_create', 'POST', '/new/device_ack', '2022-01-26 09:41:59', '2022-01-26 09:41:59'),
(15, 'device_ack 删除', 'device_ack_delete', 'POST', '/delete/device_ack', '2022-01-26 09:42:00', '2022-01-26 09:42:00'),
(16, 'device_ack 导出', 'device_ack_export', 'POST', '/export/device_ack', '2022-01-26 09:42:00', '2022-01-26 09:42:00'),
(17, 'gid 查询', 'gid_query', 'GET', '/info/gid', '2022-01-26 09:42:00', '2022-01-26 09:42:00'),
(18, 'gid 编辑页显示', 'gid_show_edit', 'GET', '/info/gid/edit', '2022-01-26 09:42:00', '2022-01-26 09:42:00'),
(19, 'gid 新建记录页显示', 'gid_show_create', 'GET', '/info/gid/new', '2022-01-26 09:42:00', '2022-01-26 09:42:00'),
(20, 'gid 编辑', 'gid_edit', 'POST', '/edit/gid', '2022-01-26 09:42:01', '2022-01-26 09:42:01'),
(21, 'gid 新建', 'gid_create', 'POST', '/new/gid', '2022-01-26 09:42:01', '2022-01-26 09:42:01'),
(22, 'gid 删除', 'gid_delete', 'POST', '/delete/gid', '2022-01-26 09:42:01', '2022-01-26 09:42:01'),
(23, 'gid 导出', 'gid_export', 'POST', '/export/gid', '2022-01-26 09:42:01', '2022-01-26 09:42:01'),
(24, 'group 查询', 'group_query', 'GET', '/info/group', '2022-01-26 09:42:01', '2022-01-26 09:42:01'),
(25, 'group 编辑页显示', 'group_show_edit', 'GET', '/info/group/edit', '2022-01-26 09:42:02', '2022-01-26 09:42:02'),
(26, 'group 新建记录页显示', 'group_show_create', 'GET', '/info/group/new', '2022-01-26 09:42:02', '2022-01-26 09:42:02'),
(27, 'group 编辑', 'group_edit', 'POST', '/edit/group', '2022-01-26 09:42:02', '2022-01-26 09:42:02'),
(28, 'group 新建', 'group_create', 'POST', '/new/group', '2022-01-26 09:42:02', '2022-01-26 09:42:02'),
(29, 'group 删除', 'group_delete', 'POST', '/delete/group', '2022-01-26 09:42:02', '2022-01-26 09:42:02'),
(30, 'group 导出', 'group_export', 'POST', '/export/group', '2022-01-26 09:42:02', '2022-01-26 09:42:02'),
(31, 'group_user 查询', 'group_user_query', 'GET', '/info/group_user', '2022-01-26 09:42:03', '2022-01-26 09:42:03'),
(32, 'group_user 编辑页显示', 'group_user_show_edit', 'GET', '/info/group_user/edit', '2022-01-26 09:42:03', '2022-01-26 09:42:03'),
(33, 'group_user 新建记录页显示', 'group_user_show_create', 'GET', '/info/group_user/new', '2022-01-26 09:42:03', '2022-01-26 09:42:03'),
(34, 'group_user 编辑', 'group_user_edit', 'POST', '/edit/group_user', '2022-01-26 09:42:03', '2022-01-26 09:42:03'),
(35, 'group_user 新建', 'group_user_create', 'POST', '/new/group_user', '2022-01-26 09:42:03', '2022-01-26 09:42:03'),
(36, 'group_user 删除', 'group_user_delete', 'POST', '/delete/group_user', '2022-01-26 09:42:04', '2022-01-26 09:42:04'),
(37, 'group_user 导出', 'group_user_export', 'POST', '/export/group_user', '2022-01-26 09:42:04', '2022-01-26 09:42:04'),
(38, 'message 查询', 'message_query', 'GET', '/info/message', '2022-01-26 09:42:04', '2022-01-26 09:42:04'),
(39, 'message 编辑页显示', 'message_show_edit', 'GET', '/info/message/edit', '2022-01-26 09:42:04', '2022-01-26 09:42:04'),
(40, 'message 新建记录页显示', 'message_show_create', 'GET', '/info/message/new', '2022-01-26 09:42:04', '2022-01-26 09:42:04'),
(41, 'message 编辑', 'message_edit', 'POST', '/edit/message', '2022-01-26 09:42:05', '2022-01-26 09:42:05'),
(42, 'message 新建', 'message_create', 'POST', '/new/message', '2022-01-26 09:42:05', '2022-01-26 09:42:05'),
(43, 'message 删除', 'message_delete', 'POST', '/delete/message', '2022-01-26 09:42:05', '2022-01-26 09:42:05'),
(44, 'message 导出', 'message_export', 'POST', '/export/message', '2022-01-26 09:42:05', '2022-01-26 09:42:05'),
(45, 'user 查询', 'user_query', 'GET', '/info/user', '2022-01-26 09:42:05', '2022-01-26 09:42:05'),
(46, 'user 编辑页显示', 'user_show_edit', 'GET', '/info/user/edit', '2022-01-26 09:42:05', '2022-01-26 09:42:05'),
(47, 'user 新建记录页显示', 'user_show_create', 'GET', '/info/user/new', '2022-01-26 09:42:05', '2022-01-26 09:42:05'),
(48, 'user 编辑', 'user_edit', 'POST', '/edit/user', '2022-01-26 09:42:06', '2022-01-26 09:42:06'),
(49, 'user 新建', 'user_create', 'POST', '/new/user', '2022-01-26 09:42:06', '2022-01-26 09:42:06'),
(50, 'user 删除', 'user_delete', 'POST', '/delete/user', '2022-01-26 09:42:06', '2022-01-26 09:42:06'),
(51, 'user 导出', 'user_export', 'POST', '/export/user', '2022-01-26 09:42:06', '2022-01-26 09:42:06'),
(52, 'group_a 查询', 'group_a_query', 'GET', '/info/group_a', '2022-01-26 09:50:56', '2022-01-26 09:50:56'),
(53, 'group_a 编辑页显示', 'group_a_show_edit', 'GET', '/info/group_a/edit', '2022-01-26 09:50:56', '2022-01-26 09:50:56'),
(54, 'group_a 新建记录页显示', 'group_a_show_create', 'GET', '/info/group_a/new', '2022-01-26 09:50:57', '2022-01-26 09:50:57'),
(55, 'group_a 编辑', 'group_a_edit', 'POST', '/edit/group_a', '2022-01-26 09:50:57', '2022-01-26 09:50:57'),
(56, 'group_a 新建', 'group_a_create', 'POST', '/new/group_a', '2022-01-26 09:50:57', '2022-01-26 09:50:57'),
(57, 'group_a 删除', 'group_a_delete', 'POST', '/delete/group_a', '2022-01-26 09:50:57', '2022-01-26 09:50:57'),
(58, 'group_a 导出', 'group_a_export', 'POST', '/export/group_a', '2022-01-26 09:50:57', '2022-01-26 09:50:57');


CREATE TABLE IF NOT EXISTS `goadmin_roles` (
`id` int unsigned NOT NULL AUTO_INCREMENT,
`name` varchar(50)NOT NULL,
`slug` varchar(50)NOT NULL,
`created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
`updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
PRIMARY KEY (`id`),
UNIQUE KEY `admin_roles_name_unique` (`name`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 ;

REPLACE INTO `goadmin_roles` (`id`, `name`, `slug`, `created_at`, `updated_at`) VALUES
(1, 'Administrator', 'administrator', '2019-09-10 00:00:00', '2019-09-10 00:00:00'),
(2, 'Operator', 'operator', '2019-09-10 00:00:00', '2019-09-10 00:00:00');

CREATE TABLE IF NOT EXISTS `goadmin_role_menu` (
`role_id` int unsigned NOT NULL,
`menu_id` int unsigned NOT NULL,
`created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
`updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
KEY `admin_role_menu_role_id_menu_id_index` (`role_id`,`menu_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ;


REPLACE INTO `goadmin_role_menu` (`role_id`, `menu_id`, `created_at`, `updated_at`) VALUES
(1, 1, '2019-09-10 00:00:00', '2019-09-10 00:00:00'),
(1, 7, '2019-09-10 00:00:00', '2019-09-10 00:00:00'),
(2, 7, '2019-09-10 00:00:00', '2019-09-10 00:00:00'),
(1, 8, '2019-09-11 10:20:55', '2019-09-11 10:20:55'),
(2, 8, '2019-09-11 10:20:55', '2019-09-11 10:20:55');

CREATE TABLE IF NOT EXISTS `goadmin_role_permissions` (
`role_id` int unsigned NOT NULL,
`permission_id` int unsigned NOT NULL,
`created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
`updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  UNIQUE KEY `admin_role_permissions` (`role_id`,`permission_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ;


REPLACE INTO `goadmin_role_permissions` (`role_id`, `permission_id`, `created_at`, `updated_at`) VALUES
(1, 1, '2019-09-10 00:00:00', '2019-09-10 00:00:00'),
(1, 2, '2019-09-10 00:00:00', '2019-09-10 00:00:00'),
(2, 2, '2019-09-10 00:00:00', '2019-09-10 00:00:00');

CREATE TABLE IF NOT EXISTS `goadmin_role_users` (
`role_id` int unsigned NOT NULL,
`user_id` int unsigned NOT NULL,
`created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
`updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
UNIQUE KEY `admin_user_roles` (`role_id`,`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ;


REPLACE INTO `goadmin_role_users` (`role_id`, `user_id`, `created_at`, `updated_at`) VALUES
(1, 1, '2019-09-10 00:00:00', '2019-09-10 00:00:00'),
(2, 2, '2019-09-10 00:00:00', '2019-09-10 00:00:00');

CREATE TABLE IF NOT EXISTS `goadmin_session` (
 `id` int unsigned NOT NULL AUTO_INCREMENT,
 `sid` varchar(50)NOT NULL DEFAULT '',
`values` varchar(3000)NOT NULL DEFAULT '',
`created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
`updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
PRIMARY KEY (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 ;


REPLACE INTO `goadmin_session` (`id`, `sid`, `values`, `created_at`, `updated_at`) VALUES
(16, 'c033b5a8-0795-461e-ae75-bfccd2f5e37a', '{"user_id":1}', '2022-01-27 07:24:16', '2022-01-27 07:24:16');

CREATE TABLE IF NOT EXISTS `goadmin_site` (
`id` int unsigned NOT NULL AUTO_INCREMENT,
`key` varchar(100)DEFAULT NULL,
`value` longtext  ,
`description` varchar(3000)DEFAULT NULL,
`state` tinyint unsigned NOT NULL DEFAULT '0',
`created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
`updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
PRIMARY KEY (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 ;


REPLACE INTO `goadmin_site` (`id`, `key`, `value`, `description`, `state`, `created_at`, `updated_at`) VALUES
(1, 'app_id', 'dwJRmCKHDriverVu7', NULL, 1, '2022-01-26 09:22:19', '2022-01-26 09:22:19'),
(2, 'env', 'local', NULL, 1, '2022-01-26 09:22:19', '2022-01-26 09:22:19'),
(3, 'error_log_path', 'E:/go_project/jim_admin/logs/error.log', NULL, 1, '2022-01-26 09:22:19', '2022-01-26 09:22:19'),
(4, 'info_log_off', 'false', NULL, 1, '2022-01-26 09:22:20', '2022-01-26 09:22:20'),
(5, 'file_upload_engine', '{"name":"local"}', NULL, 1, '2022-01-26 09:22:20', '2022-01-26 09:22:20'),
(6, 'prohibit_config_modification', 'false', NULL, 1, '2022-01-26 09:22:20', '2022-01-26 09:22:20'),
(7, 'sql_log', 'false', NULL, 1, '2022-01-26 09:22:20', '2022-01-26 09:22:20'),
(8, 'logger_encoder_duration', '', NULL, 1, '2022-01-26 09:22:21', '2022-01-26 09:22:21'),
(9, 'animation_type', '', NULL, 1, '2022-01-26 09:22:21', '2022-01-26 09:22:21'),
(10, 'language', 'cn', NULL, 1, '2022-01-26 09:22:21', '2022-01-26 09:22:21'),
(11, 'logger_level', '0', NULL, 1, '2022-01-26 09:22:21', '2022-01-26 09:22:21'),
(12, 'custom_head_html', '', NULL, 1, '2022-01-26 09:22:22', '2022-01-26 09:22:22'),
(13, 'animation_delay', '0.00', NULL, 1, '2022-01-26 09:22:22', '2022-01-26 09:22:22'),
(14, 'allow_del_operation_log', 'false', NULL, 1, '2022-01-26 09:22:22', '2022-01-26 09:22:22'),
(15, 'debug', 'true', NULL, 1, '2022-01-26 09:22:22', '2022-01-26 09:22:22'),
(16, 'bootstrap_file_path', 'E:/go_project/jim_admin/bootstrap.go', NULL, 1, '2022-01-26 09:22:22', '2022-01-26 09:22:22'),
(17, 'mini_logo', 'JA', NULL, 1, '2022-01-26 09:22:23', '2022-01-26 09:22:23'),
(18, 'logger_encoder_time_key', '', NULL, 1, '2022-01-26 09:22:23', '2022-01-26 09:22:23'),
(19, 'logger_encoder_time', '', NULL, 1, '2022-01-26 09:22:23', '2022-01-26 09:22:23'),
(20, 'login_title', 'GoAdmin', NULL, 1, '2022-01-26 09:22:23', '2022-01-26 09:22:23'),
(21, 'operation_log_off', 'false', NULL, 1, '2022-01-26 09:22:24', '2022-01-26 09:22:24'),
(22, 'index_url', '/', NULL, 1, '2022-01-26 09:22:24', '2022-01-26 09:22:24'),
(23, 'access_log_off', 'false', NULL, 1, '2022-01-26 09:22:24', '2022-01-26 09:22:24'),
(24, 'logger_encoder_message_key', '', NULL, 1, '2022-01-26 09:22:24', '2022-01-26 09:22:24'),
(25, 'logger_encoder_level', '', NULL, 1, '2022-01-26 09:22:24', '2022-01-26 09:22:24'),
(26, 'exclude_theme_components', 'null', NULL, 1, '2022-01-26 09:22:24', '2022-01-26 09:22:24'),
(27, 'login_url', '/login', NULL, 1, '2022-01-26 09:22:25', '2022-01-26 09:22:25'),
(28, 'access_log_path', 'E:/go_project/jim_admin/logs/access.log', NULL, 1, '2022-01-26 09:22:25', '2022-01-26 09:22:25'),
(29, 'hide_plugin_entrance', 'false', NULL, 1, '2022-01-26 09:22:25', '2022-01-26 09:22:25'),
(30, 'asset_root_path', './public/', NULL, 1, '2022-01-26 09:22:25', '2022-01-26 09:22:25'),
(31, 'custom_404_html', '', NULL, 1, '2022-01-26 09:22:25', '2022-01-26 09:22:25'),
(32, 'logger_encoder_name_key', '', NULL, 1, '2022-01-26 09:22:26', '2022-01-26 09:22:26'),
(33, 'logger_encoder_stacktrace_key', '', NULL, 1, '2022-01-26 09:22:26', '2022-01-26 09:22:26'),
(34, 'color_scheme', '', NULL, 1, '2022-01-26 09:22:26', '2022-01-26 09:22:26'),
(35, 'asset_url', '', NULL, 1, '2022-01-26 09:22:26', '2022-01-26 09:22:26'),
(36, 'login_logo', 'JimAdmin', NULL, 1, '2022-01-26 09:22:26', '2022-01-26 09:22:26'),
(37, 'extra', '', NULL, 1, '2022-01-26 09:22:26', '2022-01-26 09:22:26'),
(38, 'hide_tool_entrance', 'false', NULL, 1, '2022-01-26 09:22:26', '2022-01-26 09:22:26'),
(39, 'custom_500_html', '', NULL, 1, '2022-01-26 09:22:27', '2022-01-26 09:22:27'),
(40, 'logo', 'JimAdmin', NULL, 1, '2022-01-26 09:22:27', '2022-01-26 09:22:27'),
(41, 'animation_duration', '0.00', NULL, 1, '2022-01-26 09:22:27', '2022-01-26 09:22:27'),
(42, 'site_off', 'false', NULL, 1, '2022-01-26 09:22:27', '2022-01-26 09:22:27'),
(43, 'go_mod_file_path', 'E:/go_project/jim_admin/go.mod', NULL, 1, '2022-01-26 09:22:28', '2022-01-26 09:22:28'),
(44, 'logger_rotate_max_age', '0', NULL, 1, '2022-01-26 09:22:28', '2022-01-26 09:22:28'),
(45, 'logger_rotate_compress', 'false', NULL, 1, '2022-01-26 09:22:28', '2022-01-26 09:22:28'),
(46, 'logger_encoder_encoding', '', NULL, 1, '2022-01-26 09:22:28', '2022-01-26 09:22:28'),
(47, 'session_life_time', '7200', NULL, 1, '2022-01-26 09:22:28', '2022-01-26 09:22:28'),
(48, 'auth_user_table', 'goadmin_users', NULL, 1, '2022-01-26 09:22:28', '2022-01-26 09:22:28'),
(49, 'open_admin_api', 'false', NULL, 1, '2022-01-26 09:22:29', '2022-01-26 09:22:29'),
(50, 'logger_rotate_max_backups', '0', NULL, 1, '2022-01-26 09:22:29', '2022-01-26 09:22:29'),
(51, 'logger_rotate_max_size', '0', NULL, 1, '2022-01-26 09:22:29', '2022-01-26 09:22:29'),
(52, 'no_limit_login_ip', 'false', NULL, 1, '2022-01-26 09:22:29', '2022-01-26 09:22:29'),
(53, 'theme', 'sword', NULL, 1, '2022-01-26 09:22:29', '2022-01-26 09:22:29'),
(54, 'access_assets_log_off', 'false', NULL, 1, '2022-01-26 09:22:30', '2022-01-26 09:22:30'),
(55, 'domain', '', NULL, 1, '2022-01-26 09:22:30', '2022-01-26 09:22:30'),
(56, 'hide_app_info_entrance', 'false', NULL, 1, '2022-01-26 09:22:30', '2022-01-26 09:22:30'),
(57, 'hide_visitor_user_center_entrance', 'false', NULL, 1, '2022-01-26 09:22:30', '2022-01-26 09:22:30'),
(58, 'error_log_off', 'false', NULL, 1, '2022-01-26 09:22:31', '2022-01-26 09:22:31'),
(59, 'logger_encoder_level_key', '', NULL, 1, '2022-01-26 09:22:31', '2022-01-26 09:22:31'),
(60, 'logger_encoder_caller_key', '', NULL, 1, '2022-01-26 09:22:31', '2022-01-26 09:22:31'),
(61, 'logger_encoder_caller', '', NULL, 1, '2022-01-26 09:22:31', '2022-01-26 09:22:31'),
(62, 'custom_foot_html', '', NULL, 1, '2022-01-26 09:22:31', '2022-01-26 09:22:31'),
(63, 'footer_info', '', NULL, 1, '2022-01-26 09:22:31', '2022-01-26 09:22:31'),
(64, 'custom_403_html', '', NULL, 1, '2022-01-26 09:22:32', '2022-01-26 09:22:32'),
(65, 'url_prefix', 'admin', NULL, 1, '2022-01-26 09:22:32', '2022-01-26 09:22:32'),
(66, 'title', 'JimAdmin', NULL, 1, '2022-01-26 09:22:32', '2022-01-26 09:22:32'),
(67, 'info_log_path', 'E:/go_project/jim_admin/logs/info.log', NULL, 1, '2022-01-26 09:22:32', '2022-01-26 09:22:32'),
(68, 'hide_config_center_entrance', 'false', NULL, 1, '2022-01-26 09:22:32', '2022-01-26 09:22:32');

CREATE TABLE IF NOT EXISTS `goadmin_users` (
`id` int unsigned NOT NULL AUTO_INCREMENT,
`username` varchar(100)NOT NULL,
`password` varchar(100)NOT NULL DEFAULT '',
`name` varchar(100)NOT NULL,
`avatar` varchar(255)DEFAULT NULL,
`remember_token` varchar(100)DEFAULT NULL,
`created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
`updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
PRIMARY KEY (`id`),
UNIQUE KEY `admin_users_username_unique` (`username`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 ;


REPLACE INTO `goadmin_users` (`id`, `username`, `password`, `name`, `avatar`, `remember_token`, `created_at`, `updated_at`) VALUES
(1, 'admin', '$2a$10$YK04HUCz/9ZrFceqCNpFR.wkD0Ew1Z9uVUK.ARmgqnZhyftElLAH2', 'admin', '', 'tlNcBVK9AvfYH7WEnwB1RKvocJu8FfRy4um3DJtwdHuJy0dwFsLOgAc0xUfh', '2019-09-10 00:00:00', '2019-09-10 00:00:00'),
(2, 'operator', '$2a$10$rVqkOzHjN2MdlEprRflb1eGP0oZXuSrbJLOmJagFsCd81YZm0bsh.', 'Operator', '', NULL, '2019-09-10 00:00:00', '2019-09-10 00:00:00');

CREATE TABLE IF NOT EXISTS `goadmin_user_permissions` (
`user_id` int unsigned NOT NULL,
`permission_id` int unsigned NOT NULL,
`created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
`updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  UNIQUE KEY `admin_user_permissions` (`user_id`,`permission_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ;


REPLACE INTO `goadmin_user_permissions` (`user_id`, `permission_id`, `created_at`, `updated_at`) VALUES
(1, 1, '2019-09-10 00:00:00', '2019-09-10 00:00:00'),
(2, 2, '2019-09-10 00:00:00', '2019-09-10 00:00:00');

CREATE TABLE IF NOT EXISTS `group` (
`id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '自增主键',
`name` varchar(50)NOT NULL DEFAULT '' COMMENT '群组名称',
`introduction` varchar(255)NOT NULL DEFAULT '' COMMENT '群组简介',
`extra` varchar(1024)NOT NULL DEFAULT '' COMMENT '附加属性',
`atavar_url` varchar(1024)NOT NULL DEFAULT '' COMMENT '头像',
`user_id` bigint NOT NULL DEFAULT '0' COMMENT '群主ID',
`created_at` datetime DEFAULT NULL COMMENT '创建时间',
`updated_at` datetime DEFAULT NULL COMMENT '更新时间',
`deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
PRIMARY KEY (`id`),
KEY `user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4  COMMENT='群组';

CREATE TABLE IF NOT EXISTS `group_user` (
`id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '自增主键',
`group_id` bigint unsigned NOT NULL COMMENT '组id',
`user_id` bigint unsigned NOT NULL COMMENT '用户id',
`user_show_name` varchar(20)    NOT NULL COMMENT '用户在群组的昵称',
`extra` varchar(1024)  NOT NULL COMMENT '附加属性',
`created_at` datetime DEFAULT NULL COMMENT '创建时间',
`updated_at` datetime DEFAULT NULL COMMENT '更新时间',
`deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
PRIMARY KEY (`id`),
UNIQUE KEY `uk_group_id_user_id` (`group_id`,`user_id`) USING BTREE,
KEY `deleted_at` (`deleted_at`),
KEY `idx_user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4  COMMENT='群组成员关系';


CREATE TABLE IF NOT EXISTS `message_content` (
 `id` bigint unsigned NOT NULL AUTO_INCREMENT,
 `type` tinyint NOT NULL DEFAULT '0' COMMENT '消息类型',
 `body` text  NOT NULL COMMENT '内容',
 `extra` text NOT NULL COMMENT '扩展',
 `created_at` datetime DEFAULT NULL COMMENT '创建时间',
 `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
 `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
 PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COMMENT='消息内容表';


CREATE TABLE IF NOT EXISTS `message_index` (
`id` bigint unsigned NOT NULL AUTO_INCREMENT,
`sender_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '发送人ID',
`receiver_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '接收人ID',
`message_content_id` bigint NOT NULL DEFAULT '0' COMMENT '消息ID',
`group_id` bigint NOT NULL DEFAULT '0' COMMENT '群ID',
`status` tinyint NOT NULL DEFAULT '1' COMMENT '状态：1-待送达，2-已送到，3-已确认，0-已删除',
`created_at` datetime DEFAULT NULL COMMENT '创建时间',
`updated_at` datetime DEFAULT NULL COMMENT '更新时间',
`deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COMMENT='消息发送表';


CREATE TABLE IF NOT EXISTS `user` (
`id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '自增主键',
`nickname` varchar(50)    NOT NULL DEFAULT '' COMMENT '昵称',
`password` varchar(100)    NOT NULL DEFAULT '' COMMENT '密码',
`gender` enum('1','2')  NOT NULL COMMENT '性别，0:未知；1:男；2:女',
`avatar_url` varchar(1024)    NOT NULL DEFAULT '' COMMENT '用户头像链接',
`extra` varchar(1024)    NOT NULL DEFAULT '' COMMENT '附加属性',
`created_at` datetime DEFAULT NULL COMMENT '创建时间',
`updated_at` datetime DEFAULT NULL COMMENT '更新时间',
`deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
PRIMARY KEY (`id`),
KEY `deleted_at` (`deleted_at`),
KEY `created_at` (`created_at`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4  COMMENT='用户';


REPLACE INTO `user` (`id`, `nickname`, `password`, `gender`, `avatar_url`, `extra`, `created_at`, `updated_at`, `deleted_at`) VALUES
(9, 'jidan1640931088', '$2a$14$4Qsk0SeGwfvSYvDloyy69OSZKDJyhXXFv9YD//AegV3lRwTXqrmMW', '1', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2021-12-31 14:11:30', '2021-12-31 14:11:30', NULL),
(10, 'jidan1640931100', '$2a$14$WYYW4q7CMPCHpt7bfwMKoug5y69dz9XwSAPZRy/iN7c2VcTXc96qu', '1', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2021-12-31 14:11:42', '2021-12-31 14:11:42', NULL),
(11, 'jidan1640931240', '$2a$14$ZHTW13054XICkr685rtatuYI8XThsiWrMSLWpZDUL9WnlEZLs0N7i', '1', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2021-12-31 14:14:02', '2021-12-31 14:14:02', NULL),
(12, 'jidan1640931242', '$2a$14$5KtxVMSU/YeJ2cD9gAPJL..IVFY7UzLJxLJNNxbC5/p9IRmTD6gf.', '1', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2021-12-31 14:14:03', '2021-12-31 14:14:03', NULL),
(13, 'jidan1640931243', '$2a$14$rVj14NILNrKkt6gNswBP7e0LV6npgESYGl4P46bOB50iWaFAau5xK', '1', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2021-12-31 14:14:05', '2021-12-31 14:14:05', NULL),
(14, 'jidan1640931245', '$2a$14$bZjWNIAhlBlhcmF3UY/Keuzh7Nk71pxP5i2iQn7tiIxnoGjgsGVrm', '1', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2021-12-31 14:14:06', '2021-12-31 14:14:06', NULL),
(15, 'jidan1640931246', '$2a$14$geYwPnphkOZOM0tuO58BouO6brHStzDgLeDuNrOHe0CN1qL2aGrKa', '1', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2021-12-31 14:14:08', '2021-12-31 14:14:08', NULL),
(16, 'jidan1640931247', '$2a$14$tFRbpwpi99OMq8F.doNqEOlhZoY65MSkKQb5mVhXJEfmNfVys4JOC', '1', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2021-12-31 14:14:09', '2021-12-31 14:14:09', NULL),
(17, 'jidan1640931249', '$2a$14$eyAjMxZYzOJvU38eYI/j3uCe8zUGL9LgtjYYikxrfK43uFUeLxCKe', '1', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2021-12-31 14:14:11', '2021-12-31 14:14:11', NULL),
(18, 'jidan1640931250', '$2a$14$oAZSmRdX6oREptUynn/V3OWx9uerx01JwmYNFkr7DqxXPmGzgPjym', '1', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2021-12-31 14:14:12', '2021-12-31 14:14:12', NULL),
(19, 'jidan1640931252', '$2a$14$U6RLxtGxAJHRipV3S.eaFuH9CwSmHZ/ted9G68gqsu5uWyaRUXEky', '1', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2021-12-31 14:14:13', '2021-12-31 14:14:13', NULL),
(20, 'jidan1640931253', '$2a$14$iIzhucZcRYO2LFMGj83fb..DnPhAwaZFQhvGaQCWlyAFGGNQEPvoe', '1', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2021-12-31 14:14:15', '2021-12-31 14:14:15', NULL),
(21, 'jidan1640931254', '$2a$14$LFtSonIqXmhxNiva1bAHrOEOoXgJE.Y0IqxbiYGFbVbDpLBjFsKDC', '1', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2021-12-31 14:14:16', '2021-12-31 14:14:16', NULL),
(22, 'jidan1640931256', '$2a$14$FN60wzuYzvrgcQooHzJTlO3vggyyEyHyaoLrYQmpe7o1jHP2SI81G', '1', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2021-12-31 14:14:18', '2021-12-31 14:14:18', NULL),
(23, 'jidan1640931257', '$2a$14$CSnsgxsvG2AjC0HTqujmSusWq7EDdUN8gxg2ueb9k3E5R07F715jS', '1', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2021-12-31 14:14:19', '2021-12-31 14:14:19', NULL),
(24, 'jidan1640931259', '$2a$14$vj705XxDqGxNWP25iAi5g.oFh8UoBeBMD/n5UxHf1E/cIKZqAhHNa', '1', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2021-12-31 14:14:21', '2021-12-31 14:14:21', NULL),
(25, 'jidan1640931260', '$2a$14$tpNSWmKPrpf/pI91Q1zwM.cCwNQFMckbAOXxXJKFSRtEa.p/M/Nzm', '1', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2021-12-31 14:14:22', '2021-12-31 14:14:22', NULL),
(26, 'jidan1640931262', '$2a$14$M608.5CQS.fbuQ2yn7jz4eRPjTMo3R4FdwR5SrFoxmVDMleZayB/q', '1', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2021-12-31 14:14:24', '2021-12-31 14:14:24', NULL),
(27, 'jidan1640931263', '$2a$14$/H5Ggi.5QIGKjUGcBb8KpO5QUKF888vHucWg3jld1flB1v9PL68d6', '1', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2021-12-31 14:14:25', '2021-12-31 14:14:25', NULL),
(28, 'jidan1640931265', '$2a$14$obTaQUQpdbp33V59r1/Sv.mKx.dg2C9rU1hNE2uODSTy5LQqMIvUW', '1', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2021-12-31 14:14:26', '2021-12-31 14:14:26', NULL),
(29, 'jidan1640931266', '$2a$14$MtaFK8fXJvQrqjHgmhL1Du5b2gAm2dMh2ttNsHc6MC0dyAT7ODagq', '1', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2021-12-31 14:14:28', '2021-12-31 14:14:28', NULL),
(30, 'jidan1640931267', '$2a$14$RxzHsGmRu77glzEJpJ6dWefsWa7jm8KCdnkFoREcjfb9E74rATmPC', '1', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2021-12-31 14:14:29', '2021-12-31 14:14:29', NULL),
(31, 'jidan1640931269', '$2a$14$V6VWaLX0Tzmwd9yLOVfg1e61.qGIc36GqJ0MKHYdpLTc9p/aN/.Oe', '1', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2021-12-31 14:14:31', '2021-12-31 14:14:31', NULL),
(32, 'jidan1640931270', '$2a$14$FxtPVbCfHFN9RUECWRZuWumMQkqWxFc4mDscWfBZ0AVBkumFQ2Iae', '1', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2021-12-31 14:14:32', '2021-12-31 14:14:32', NULL),
(33, 'jidan1640931272', '$2a$14$sZRdprXGlqZX/WeJ5FGETuAztT/1/L1PP42flNccWgWnmsuM2O32m', '1', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2021-12-31 14:14:34', '2021-12-31 14:14:34', NULL),
(34, 'jidan1640931273', '$2a$14$cTbKXzdx8noDX8VGxkFAfu8rIi6rgF6Q6F/sCgPePaaSYVw/mh2Ya', '1', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2021-12-31 14:14:35', '2021-12-31 14:14:35', NULL),
(35, 'jidan1640931275', '$2a$14$RlS4x/bnUftHFjCtwGk4wOR7UXv3XdyA9aHng1aIpMfDzjHThJNfW', '1', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2021-12-31 14:14:36', '2021-12-31 14:14:36', NULL),
(36, 'jidan1640931276', '$2a$14$ZSzYOgJdz4pKS/a3.JvyO.LMARy7OzEI68K0CPy6aFgMgDd3IBPH2', '1', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2021-12-31 14:14:38', '2021-12-31 14:14:38', NULL),
(37, 'jidan1640931278', '$2a$14$GrDig3GW1my1C9fpbVWdzuKcH1hdD7C2HZQVoivUjP3.O6JTGAQaq', '1', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2021-12-31 14:14:39', '2021-12-31 14:14:39', NULL),
(38, 'jidan1640931279', '$2a$14$lm099H8fPoB0kFLyshyAjuwS4bzKsX6t7u4h3qKHI0BHgatx.o3jO', '1', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2021-12-31 14:14:41', '2021-12-31 14:14:41', NULL),
(39, 'jidan1640931281', '$2a$14$t6h6S5mCUbQ9Jh0DqoC/D.6Bf0E.ZYyNvdbTlLlcqveNlb5sDZQXO', '1', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2021-12-31 14:14:42', '2021-12-31 14:14:42', NULL),
(40, 'jidan1640931282', '$2a$14$oIpaPzM09fO7u0DNBsxx5.ih/Yp5/EZlaJIozJS39rjYQnWDRK7Aa', '1', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2021-12-31 14:14:44', '2021-12-31 14:14:44', NULL),
(41, 'jidan1640931284', '$2a$14$vcD9GIbHfcKORH6g2XB9M.25avszQiWomFp7S3.FiSk2/q12esoKm', '1', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2021-12-31 14:14:45', '2021-12-31 14:14:45', NULL),
(42, 'jidan1640931285', '$2a$14$VWyCyy3TDVoaCbDdzfGMs.fcjTDuzo2qRlMhyPqaiJ10sZmqncuw2', '1', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2021-12-31 14:14:47', '2021-12-31 14:14:47', NULL),
(43, 'jidan1640931286', '$2a$14$yBp0NG3ixkLEG1WEoBEKduAX4kp1hb9aJpFrOWeBuJkk1KiEF0RtC', '1', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2021-12-31 14:14:48', '2021-12-31 14:14:48', NULL),
(44, 'jidan1640931288', '$2a$14$UZ3zspF1pttSvqytbeuvXOJWh.HVLj3LAnTslg4f5EiT7ShXdbPL2', '1', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2021-12-31 14:14:50', '2021-12-31 14:14:50', NULL),
(45, 'jidan1640931289', '$2a$14$mTNCy2Cu.g3dVUDf8DvBYOE/jk32eWrDywqSCKtn/X1.pYYAwXroC', '1', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2021-12-31 14:14:51', '2021-12-31 14:14:51', NULL),
(46, 'jidan1640931291', '$2a$14$pBoHfQfPs1zAypafHhM6C.z71CVGGV9RJRSsF8GULHixdobArAZ0i', '1', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2021-12-31 14:14:53', '2021-12-31 14:14:53', NULL),
(47, 'jidan1640931292', '$2a$14$eHDC17R.Ca3H9ICeLSZeyuP29LY9aNI7JnfWHkE84whcwp6Ekd97.', '1', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2021-12-31 14:14:54', '2021-12-31 14:14:54', NULL),
(48, 'jidan1640931294', '$2a$14$NqPM3CQG9OoOPFFjF4sFROAzO3ehQWqijD7kGbLTRxqOCrkR8M7Fe', '1', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2021-12-31 14:14:55', '2021-12-31 14:14:55', NULL),
(49, 'jidan1640931295', '$2a$14$k4W.HGr5tiFvEtuvJ.XV/eo.DrCyXneF3qWMsmMMwNwGT1cqJcN9a', '1', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2021-12-31 14:14:57', '2021-12-31 14:14:57', NULL),
(50, 'jidan1640931297', '$2a$14$kPEHn3JkC8jDZ4bWYPdiO.UbuInJCxco/ACqJaFWoAnaQzF3HmJV6', '1', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2021-12-31 14:14:58', '2021-12-31 14:14:58', NULL),
(51, 'jidan1640931298', '$2a$14$MzsSNpdn7kkjVWL5NROMQueD4giFIkGUm5vLrRkLz9YEKgGBv1wrG', '1', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2021-12-31 14:15:00', '2021-12-31 14:15:00', NULL),
(52, 'jidan1640931300', '$2a$14$DYoCD344gHswipetKhOFYO1BJ69n8B2cHLDYbvjQ7FdhgR.Yc9RL.', '1', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2021-12-31 14:15:01', '2021-12-31 14:15:01', NULL),
(53, 'jidan1640931301', '$2a$14$F1kxTRF.YthnTqRrNvCKjulCkEd0xe4RKN40mFCw6wnJhIno4kIT6', '1', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2021-12-31 14:15:03', '2021-12-31 14:15:03', NULL),
(54, 'jidan1640931303', '$2a$14$0VPl.eVAj54CIp6W8DwHDeX66xymzNUkuHp5lSKn6dUJQoSzSsise', '1', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2021-12-31 14:15:04', '2021-12-31 14:15:04', NULL),
(55, 'jidan1640931304', '$2a$14$7EjuY5sPcfTD8Ly6Q6wt7ePqYQQwky1/9IwdvYrH14VPFWG1WnYlO', '1', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2021-12-31 14:15:06', '2021-12-31 14:15:06', NULL),
(56, 'jidan1640931305', '$2a$14$dhn4LeRjuBXUXwXIk3.bv.CiRiyV04g9oYvIqi/5EYrsaPQ9TFLa6', '1', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2021-12-31 14:15:07', '2021-12-31 14:15:07', NULL),
(57, 'jidan1640931307', '$2a$14$S8tbfsmvJABt7fw1qAjj7eOKXlkDd/iPUHH5OdD2PJ24LYk9Xb.bm', '1', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2021-12-31 14:15:09', '2021-12-31 14:15:09', NULL),
(58, 'jidan1640931308', '$2a$14$rh/noM8TbW2UImPLmL3HDeT3KKYTL7.FahjMnK5H7RfImAvK4Yl7C', '1', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2021-12-31 14:15:10', '2021-12-31 14:15:10', NULL),
(59, 'jidan1640931310', '$2a$14$IAys.kxTJ016yCGS3IxvTeqLUQM0dxWOikv9vNTMiifT2acw/N9ES', '1', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2021-12-31 14:15:11', '2021-12-31 14:15:11', NULL),
(60, 'jidan1640931311', '$2a$14$T0vbMlDLQN6JF2vnxZ7cTuOO26BZTWxPkaYFQnTkS.ADbjYvRBzeG', '1', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2021-12-31 14:15:13', '2021-12-31 14:15:13', NULL),
(61, 'jidan1640931313', '$2a$14$HTeBMtvxHDNcikEaGmgTKOmc1/OQMU6NYpQf1voXILFufFf1b/OXm', '1', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2021-12-31 14:15:14', '2021-12-31 14:15:14', NULL),
(62, 'jidan1640931314', '$2a$14$p6UarFKPJVk7gLDSi0jUDeXiq5H.MKbhG/JXAHxulLMJ7FzlmcL3a', '1', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2021-12-31 14:15:16', '2021-12-31 14:15:16', NULL),
(63, 'jidan1640931316', '$2a$14$NKZ5O5KPsq01AYyKcppRfep2tzFlVIkRCLf2P8LqRuZBUScnaQWqa', '1', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2021-12-31 14:15:17', '2021-12-31 14:15:17', NULL),
(64, 'jidan1640931317', '$2a$14$7f8yvC.xhv8qlO2eENgjBuNG87rzKBnuIgDidrmcw8ayQIwDdGKaq', '1', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2021-12-31 14:15:19', '2021-12-31 14:15:19', NULL),
(65, 'jidan1640931318', '$2a$14$tGwv.ZxXCd9Cg9U77Sy7luhcTJUKafzDkxgDWmwGyH5nJTzegHBB2', '1', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2021-12-31 14:15:20', '2021-12-31 14:15:20', NULL),
(66, 'jidan1640931320', '$2a$14$52KBwHY4fzsMRFpI6G1QUOTNCkjn9HZfYUuOsLIaLjRvNjxZ.Dq2G', '1', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2021-12-31 14:15:22', '2021-12-31 14:15:22', NULL),
(67, 'jidan1640931321', '$2a$14$Axja9Z.hkcZ28SoOQyYc1u7fULSO0pt5rEuaV3z4h2KxUhIQeqPze', '1', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2021-12-31 14:15:23', '2021-12-31 14:15:23', NULL),
(68, 'jidan1640931323', '$2a$14$1eSRerx.dKvacF5YqjqWcOvjz3IEyJEvxlgyXq6yc4wDKqJi7cRZ.', '1', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2021-12-31 14:15:24', '2021-12-31 14:15:24', NULL),
(69, 'jidan1640931324', '$2a$14$i6N56EAeJ.GUpIcueSdMEeoH38ePGAgTwDz.Jn0jATHtVs5Z6iuNm', '1', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2021-12-31 14:15:26', '2021-12-31 14:15:26', NULL),
(70, 'jidan1640931325', '$2a$14$hnzcmfGJnmItiL7fNeTdxeV4hStCagNXdMiJL.71L7mgPUxuk/a32', '1', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2021-12-31 14:15:27', '2021-12-31 14:15:27', NULL),
(71, 'jidan1640931327', '$2a$14$8VJxG5Y9TAkc2p9.I9ZMROXnRsBKe9wBM5QxuuwJ5rk1/HNfA81pa', '1', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2021-12-31 14:15:29', '2021-12-31 14:15:29', NULL),
(72, 'jidan1640931328', '$2a$14$RG9t9kHUhen4HqfQXQZgE.XrL1WVu8FMJBU.ntK5mhQlxs5o7Ausu', '1', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2021-12-31 14:15:30', '2021-12-31 14:15:30', NULL),
(73, 'jidan1640931330', '$2a$14$VytO4IKfWbJHv4ISBeNHsOMB2UBdWTASelppcHrDNgA6ixwJSexTG', '1', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2021-12-31 14:15:31', '2021-12-31 14:15:31', NULL),
(74, 'jidan1640931331', '$2a$14$sCkQMbgF.dadQazSf5TCXOJiHXKEj9oQEdirTqXyNzhKBx7iuRUrS', '1', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2021-12-31 14:15:33', '2021-12-31 14:15:33', NULL),
(75, 'jidan1640931333', '$2a$14$/wCaeLMu3wCMAnl1QwviiuOxu9RtigpebgtnurhhJeJnUOmoFvz5S', '1', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2021-12-31 14:15:34', '2021-12-31 14:15:34', NULL),
(76, 'jidan1640931334', '$2a$14$KH/26I1HzdVdZfa7N45GQuGau7zggkMf9pILeePQVndsj49K6HDIa', '1', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2021-12-31 14:15:36', '2021-12-31 14:15:36', NULL),
(77, 'jidan1640931335', '$2a$14$iIaH0NPmWaoNPnZj40WLseFcWaKmCZXRDlisd6.EKTPfDgMq948VC', '1', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2021-12-31 14:15:37', '2021-12-31 14:15:37', NULL),
(78, 'jidan1640931337', '$2a$14$Ralp0J4lSyu0ZDfVldfRAOTGBh1T/zdqofC.SZSnqGBUwZnhmHe6u', '1', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2021-12-31 14:15:39', '2021-12-31 14:15:39', NULL),
(79, 'jidan1640931338', '$2a$14$E8Wd0kO8.h0txvRRNg1ZjeZhWolT26mVR.xsbmaWwiSags3Pp6at2', '1', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2021-12-31 14:15:40', '2021-12-31 14:15:40', NULL),
(80, 'jidan1640931340', '$2a$14$OVFH5G0TVyrmc7pPVl686.dG6pj3kH8VMdR8mu/ndoqbciPce1Dfa', '1', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2021-12-31 14:15:41', '2021-12-31 14:15:41', NULL),
(81, 'jidan1640931341', '$2a$14$sVXwhMdnoUq33mEKOv.4c.ENVYkWMNfHxWQnkqqVSvmDWc9V8Ft/W', '1', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2021-12-31 14:15:43', '2021-12-31 14:15:43', NULL),
(82, 'jidan1640931343', '$2a$14$0HJ4bMTMpyz9Qri15QHb7upzkQptzDj80j961oVATV/l2nKOSJk42', '1', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2021-12-31 14:15:44', '2021-12-31 14:15:44', NULL),
(83, 'jidan1640931344', '$2a$14$dxhZOI2XA.fmwDiKB70ZQurGAKeqyUsfUXeN/zCSG.cySukj9wAVa', '1', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2021-12-31 14:15:46', '2021-12-31 14:15:46', NULL),
(84, 'jidan1640931346', '$2a$14$V1mlytdZdsK/0mtuOBSwe.yRSaKsfWoLPdvJO1pSWAjWQjIzi2r/u', '1', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2021-12-31 14:15:47', '2021-12-31 14:15:47', NULL),
(85, 'jidan1640931347', '$2a$14$nwXfGwYqnHDgLY0Igz6AUOeaM23L2TgTEm5o3SR0S1fv/dQm.iyzq', '1', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2021-12-31 14:15:49', '2021-12-31 14:15:49', NULL),
(86, 'jidan1640931348', '$2a$14$LHOhHkKvxjd3MBSGeyX97uE2Gs2DGKelAB5iypwtIZyuyyq/pJDOi', '1', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2021-12-31 14:15:50', '2021-12-31 14:15:50', NULL),
(87, 'jidan1640931350', '$2a$14$aLgcsaW0UDvIiYZQRQ0UK.1uG5H8UJi1PpcwDn6sqmVvn1Y0fXtdi', '1', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2021-12-31 14:15:52', '2021-12-31 14:15:52', NULL),
(88, 'jidan1640931351', '$2a$14$xXcXV/FpXtLpRLg7oNWU/.JS6pr5ro9NjkQTiS2YNfzpHs5AxelkS', '1', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2021-12-31 14:15:53', '2021-12-31 14:15:53', NULL),
(89, 'jidan1640931353', '$2a$14$Pr3CAHX8hGeYIEf8JCD7U.96cOINASJrmWqvF1kN9i3QPiT1ctJb2', '1', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2021-12-31 14:15:55', '2021-12-31 14:15:55', NULL),
(90, 'jidan1640931354', '$2a$14$0fT3b3ovPO1JXFTZpeCuy.QYDi65xBR/8AH.fqfVTo8qoUbJCiz/e', '1', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2021-12-31 14:15:56', '2021-12-31 14:15:56', NULL),
(91, 'jidan1640931356', '$2a$14$0mWxZAtJamEzRUAY2EWybOWO/Mwl.wlwIMEvbtkX2/z8..iAfdhUi', '1', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2021-12-31 14:15:57', '2021-12-31 14:15:57', NULL),
(92, 'jidan1640931357', '$2a$14$jgOTcqx6enbiTZQdcmoTHe.PQDfUAuM1Bx22e2WfJBKidLrcplZOm', '1', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2021-12-31 14:15:59', '2021-12-31 14:15:59', NULL),
(93, 'jidan1640931359', '$2a$14$vQTcP6EY3Hi8thFNN1NOb.ynJsw1mW7z8zMm7s8RCBE4uegnl3xrW', '1', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2021-12-31 14:16:00', '2021-12-31 14:16:00', NULL),
(94, 'jidan1640931360', '$2a$14$5sUcX.zeqZZYUlxiGkInPuAU32F5WB2ykmfThdqB3zdveHcvW9fzq', '1', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2021-12-31 14:16:02', '2021-12-31 14:16:02', NULL),
(95, 'jidan1640931361', '$2a$14$X4swn4Viv9cuIEX9trn.U.HWzB0uUg/.QlKwmDWagLDlwVeey17ye', '1', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2021-12-31 14:16:03', '2021-12-31 14:16:03', NULL),
(96, 'jidan1640931363', '$2a$14$Yx07ghyqg71TwFeZHVMJ7OGYLWum/amKP5pfuhueZA/y4zQBlHwTe', '1', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2021-12-31 14:16:05', '2021-12-31 14:16:05', NULL),
(97, 'jidan1640931364', '$2a$14$1SqUue44rBdz7BHpESncY.m.MdAiPpM7ZMDsilYXNZF4biMlrfg0W', '1', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2021-12-31 14:16:06', '2021-12-31 14:16:06', NULL),
(98, 'jidan1640931366', '$2a$14$ECWK2U9V31PjjwrigF3UsOt9fZYcr4WAURvbkmW2QvBegR/ibCsEe', '1', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2021-12-31 14:16:07', '2021-12-31 14:16:07', NULL),
(99, 'jidan1640931367', '$2a$14$ty4S44kJTJegwG2uWEEtmu044bjpSgX4KgHb.mRro25/UZ4yGnVki', '1', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2021-12-31 14:16:09', '2021-12-31 14:16:09', NULL),
(100, 'jidan1640931369', '$2a$14$Zp/F9/U9Qr6y3Fg21JoF0ugSrteWS5Z4ZdMOULYFWQUHR3o4q6cLm', '1', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2021-12-31 14:16:10', '2021-12-31 14:16:10', NULL),
(101, 'jidan1640931370', '$2a$14$RQZenWWeKXA8SmSpe7wYROwgtXIJrZb5zkNx6pzowFavpjlLswC9G', '1', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2021-12-31 14:16:12', '2021-12-31 14:16:12', NULL),
(102, 'jidan1640931372', '$2a$14$VTpuAVd81VvJG./0tgnUX.U.nQjjLkUWdE0j4.1wUFSWKvV7PM1sK', '1', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2021-12-31 14:16:13', '2021-12-31 14:16:13', NULL),
(103, 'jidan1640931373', '$2a$14$PqwHxk7zsDXUpl74WEm/Xu7WhuY4J9fUTvmDlbsZRa3BtkRwV3FqG', '1', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2021-12-31 14:16:15', '2021-12-31 14:16:15', NULL),
(104, 'jidan1640931374', '$2a$14$n/RoOYtWGGVNuf.TTVOp5uW.ECYxClOd/62lkVEkM5fgnnxz/uFKq', '1', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2021-12-31 14:16:16', '2021-12-31 14:16:16', NULL),
(105, 'jidan1640931376', '$2a$14$G.iNKI9BmudWUWL0xVWmnOtlAlofM.y0.70Plr0MASKTnY0ZDKNqC', '1', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2021-12-31 14:16:18', '2021-12-31 14:16:18', NULL),
(106, 'jidan1640931377', '$2a$14$C600Lrto1yEPABhBBG6ZAuz4MKKDoLz0Uh0rdBTillfx1UpEyG8t6', '1', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2021-12-31 14:16:19', '2021-12-31 14:16:19', NULL),
(107, 'jidan1640931379', '$2a$14$YrGHIlSFSsJSH2MUxzImFONgY2WMgsGTuHOQa2CyeXO91raMPRxCW', '1', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2021-12-31 14:16:21', '2021-12-31 14:16:21', NULL),
(108, 'jidan1640931380', '$2a$14$.lm8o1yHc7CpbXkmOi7z.OWkUZ89FrsgoCBohvI3ssGfBcanexo0C', '1', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2021-12-31 14:16:22', '2021-12-31 14:16:22', NULL),
(109, 'jidan1640931382', '$2a$14$j6x6oeh3P30AkfTZql.pQOdhDEXejmgM1TPoRxIaug5FCO95t2JLy', '1', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2021-12-31 14:16:23', '2021-12-31 14:16:23', NULL),
(110, 'jidan1640931383', '$2a$14$l8Vcfakw/Yfl5AauSFKzfep.9K2MKEiUsnX0mARP6Krc0Upk5faaa', '1', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2021-12-31 14:16:25', '2021-12-31 14:16:25', NULL),
(111, 'Ms. May Konopelski', '$2a$14$4zfgCiw99wvK8joT61W12OuAyoWBeX78nYKwZJQ7VOZt8HBIahv1u', '2', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2022-01-28 18:13:40', '2022-01-28 18:13:40', NULL),
(112, 'Prince Izaiah Price', '$2a$14$mu0rNBNKkH32DP14ZuWw7eiispspE2H88RJ7OI.YENwFLbQQKrDp2', '2', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2022-01-28 18:18:06', '2022-01-28 18:18:06', NULL),
(113, 'Princess Nya Keeling', '$2a$14$se4BcobBhbn7i9Ay51W.qebQFf2DakARmNp26qM6FkJmlhXqpWfVe', '2', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2022-01-28 18:21:48', '2022-01-28 18:21:48', NULL),
(114, 'Mr. Jacey McLaughlin', '$2a$14$6a2Ag3tDHY3XHS/zS9KlB.feHLzOulROrnh6fAuYhpZAFn..KqBz.', '2', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2022-01-28 18:27:15', '2022-01-28 18:27:15', NULL),
(115, 'Mrs. Enola Marquardt', '$2a$14$u./915teFHnVCLO.rVapU.i38tPoRBikJhDPZtmFDdU.K9NefE0HG', '2', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2022-01-28 18:43:22', '2022-01-28 18:43:22', NULL),
(116, 'Prof. Greg Kub', '$2a$14$wiL4JNKPHcD5DUwkRIABzO..JhJzcUNH1h1DtZqrVujk4oI0qfPNK', '2', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2022-01-29 09:55:29', '2022-01-29 09:55:29', NULL),
(117, 'Miss Zola Wilderman', '$2a$14$AzmjsABuRHAtoG85lplD5OO3hsOFrqyTbvR9OqZ4G54z1kqE8aYli', '2', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2022-01-29 11:41:15', '2022-01-29 11:41:15', NULL),
(118, 'Princess Alverta McGlynn', '$2a$14$ks2fKUqbH399U1bNncrmO.xQd8rO2fQA4GREYDBHH.ayEg8.i6h9u', '2', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2022-01-29 17:54:58', '2022-01-29 17:54:58', NULL),
(119, 'Miss Tyra Aufderhar', '$2a$14$td1MLGAa/FMlTQtoUgCTZ.gEzcVY6XHTwXTxrjkzZ6ZzGJGcbdDKG', '2', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2022-01-29 17:59:38', '2022-01-29 17:59:38', NULL),
(120, 'Lady Bettye Orn', '$2a$14$bOnFdD5m.Nyng3Gm/uVjBuvRV9qU2ogkoaVcvAu2rWwGSy.xbwpZi', '2', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2022-02-14 16:58:23', '2022-02-14 16:58:23', NULL),
(121, 'Lady Ella Schumm', '$2a$14$tESTq/BA4wiFiX8Ry0tIzeC14o6t1IQ/MdLMzrraxEitycDs/6X/a', '2', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2022-02-14 16:58:37', '2022-02-14 16:58:37', NULL),
(122, 'Queen Jaunita Mitchell', '$2a$14$5ujuMrFMII7XFuazcFecPujVBbF0qnXD6.Fz0S7a9q75maumALzJ6', '2', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2022-02-14 16:59:06', '2022-02-14 16:59:06', NULL),
(123, 'Ms. Leanne Gutkowski', '$2a$14$U6ensjnKDfUntkX38Q8kGO40n6aUHdKh/mQfyR128xeFlN13rj.E6', '2', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2022-02-14 16:59:19', '2022-02-14 16:59:19', NULL),
(124, 'Prof. Dock Wolf', '$2a$14$emvCMRHS9TEUDxtBlmmr7u1cBpa1Qcw27DquFDy3HSpSRnrRgDWtO', '2', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2022-02-14 16:59:59', '2022-02-14 16:59:59', NULL),
(125, 'Princess Lessie Herman', '$2a$14$/uoglVF0wR5VX6oonCSHUuisLWXlnyqtYs.A5YoXL2VWTlHsnh43m', '2', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2022-02-14 17:01:49', '2022-02-14 17:01:49', NULL),
(126, 'Prof. Chyna Schroeder', '$2a$14$8MRZDwEUIOYlgPTzUpECDOKM8.g8dq7S4LIdbcaxxrG8rT7R/rxxW', '2', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2022-02-14 17:02:26', '2022-02-14 17:02:26', NULL),
(127, 'Lord Conrad Grant', '$2a$14$l3as.6KFqDOuNO.BJGnSy.sM2IVK7sz9yOggABN9t/gArkL.u1mcS', '2', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2022-02-14 17:03:27', '2022-02-14 17:03:27', NULL),
(128, 'Lady Lottie Hyatt', '$2a$14$1DtKtx/rGm/LjeluqUpAMOwd17wx5GMybZYvPsJhn4Yjm5AcitrTG', '2', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2022-02-14 17:04:24', '2022-02-14 17:04:24', NULL),
(129, 'Lady Elvie Runolfsdottir', '$2a$14$3QzJSabFesQQfFkeyQiPbemgl6RNcCdPQG40fzsEpYzfEwTSG5Oli', '2', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2022-02-14 17:05:09', '2022-02-14 17:05:09', NULL),
(130, 'Prince Presley Harris', '$2a$14$1JAe1ETNZHbKKSdPrvCOwOWztYn8hSLFLclHhfVeuYIHM77yGyYUe', '2', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2022-02-14 17:06:11', '2022-02-14 17:06:11', NULL),
(131, 'Prof. Alf Collins', '$2a$14$LaJVHpobvouCNMMYToErZuQswxDhx.xY0Kv32ddGLJ1GwZMWDutGq', '2', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2022-02-14 17:06:33', '2022-02-14 17:06:33', NULL),
(132, 'Lady Mallie Purdy', '$2a$14$ZMtZOepljDXLAF9MAN/zWuaO8TMo6RpjJ9bmU6olpWluhbAIx3bcq', '2', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2022-02-14 17:07:51', '2022-02-14 17:07:51', NULL),
(133, 'Queen Elinore Cronin', '$2a$14$Gf1ISYW8fnI5JItSldCoWeBBxsIM7JpJsOxXYywV.F8vghmS8uA8i', '2', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2022-02-14 17:18:57', '2022-02-14 17:18:57', NULL),
(134, 'Miss Katelynn Pacocha', '$2a$14$dIOquNA9gWUxE7fPjMRYRObufo..F2QdcU92tEelEq/iHC7t2AZDC', '2', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2022-02-15 11:39:06', '2022-02-15 11:39:06', NULL),
(135, 'Queen Estrella Green', '$2a$14$qu14F6XYnt6iG0ksGhixYOlbZ67sThewUwajh36Hkbne.PJFNBqf.', '2', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2022-02-15 12:13:06', '2022-02-15 12:13:06', NULL),
(136, 'Ms. Elisa Wiegand', '$2a$14$uJh/2iQy5sJgD5W6PVmOE.MXfuDMpQJiDSOy0hxU8NFrjcYFYATK6', '2', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2022-02-16 12:41:14', '2022-02-16 12:41:14', NULL),
(137, 'Prof. Alfredo Hamill', '$2a$14$XkUn0gtSwWwfnEY5JPLVouyQy8iPCuf2lJhL0SPwf0cdFhKt.hTxC', '2', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2022-02-16 17:16:32', '2022-02-16 17:16:32', NULL),
(138, 'Prince Maximillian Wilkinson', '$2a$14$mXbFVT5zrWMJLZrQ7C6en.sLlEHODaX1sJOYHh8.8/QwcfQP847Zu', '2', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2022-02-16 17:22:32', '2022-02-16 17:22:32', NULL),
(139, 'Lord Charles Schowalter', '$2a$14$XwNtqgBgvmzP9svyT17gOOhNUf3grBskV3kOlH.J/pzHsCnmcqI6W', '2', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2022-02-16 17:23:28', '2022-02-16 17:23:28', NULL),
(140, 'Mrs. Cecelia Parker', '$2a$14$.DkvjkIKvvsV6wZ/LrLH2OIOPkErl9Z5JvGvyjT2kIWifmQFf5TFW', '2', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2022-02-16 17:25:05', '2022-02-16 17:25:05', NULL),
(141, 'Miss Leilani Reichel', '$2a$14$L84hL7csv90Dk.TU/QF8euVPwQus7HEvRZkRUbtisAb0Otj/xgncO', '2', 'https://cdn.libravatar.org/static/img/nobody/80.png', '', '2022-02-16 17:41:02', '2022-02-16 17:41:02', NULL);
