-- HOHO Park 数据库设计 V2.0
-- 统一术语：藏品(asset) → 作品(artwork)
-- 新增功能：创作、任务、公告、出价、阳光账户

-- ============================================
-- 1. 用户相关表
-- ============================================

-- 用户表（保持不变）
CREATE TABLE IF NOT EXISTS `users` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '用户ID',
  `uid` varchar(20) NOT NULL COMMENT '用户唯一标识',
  `phone` varchar(20) DEFAULT NULL COMMENT '手机号',
  `password_hash` varchar(255) DEFAULT NULL COMMENT '密码哈希',
  `nickname` varchar(50) DEFAULT NULL COMMENT '昵称',
  `avatar` varchar(255) DEFAULT NULL COMMENT '头像URL',
  `is_verified` tinyint(1) DEFAULT '0' COMMENT '是否实名认证',
  `verification_name` varchar(50) DEFAULT NULL COMMENT '实名姓名',
  `verification_id_card` varchar(18) DEFAULT NULL COMMENT '身份证号',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_uid` (`uid`),
  UNIQUE KEY `idx_phone` (`phone`),
  KEY `idx_created_at` (`created_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户表';

-- 用户积分表（保持不变）
CREATE TABLE IF NOT EXISTS `user_points` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `user_id` bigint unsigned NOT NULL COMMENT '用户ID',
  `available_points` decimal(20,8) NOT NULL DEFAULT '0.00000000' COMMENT '可用积分',
  `frozen_points` decimal(20,8) NOT NULL DEFAULT '0.00000000' COMMENT '冻结积分',
  `total_earned` decimal(20,8) NOT NULL DEFAULT '0.00000000' COMMENT '累计获得',
  `total_spent` decimal(20,8) NOT NULL DEFAULT '0.00000000' COMMENT '累计消费',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_user_id` (`user_id`),
  CONSTRAINT `fk_user_points_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户积分表';

-- 积分交易记录表（保持不变）
CREATE TABLE IF NOT EXISTS `point_transactions` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '交易ID',
  `user_id` bigint unsigned NOT NULL COMMENT '用户ID',
  `type` enum('earn','spend','freeze','unfreeze','transfer_in','transfer_out') NOT NULL COMMENT '交易类型',
  `amount` decimal(20,8) NOT NULL COMMENT '金额',
  `balance_after` decimal(20,8) NOT NULL COMMENT '交易后余额',
  `source` varchar(50) NOT NULL COMMENT '来源',
  `description` varchar(255) DEFAULT NULL COMMENT '描述',
  `related_id` bigint unsigned DEFAULT NULL COMMENT '关联ID',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_type` (`type`),
  KEY `idx_created_at` (`created_at`),
  CONSTRAINT `fk_point_transactions_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='积分交易记录表';

-- ============================================
-- 2. 作品相关表（术语统一：asset → artwork）
-- ============================================

-- 作品表（原assets表）
CREATE TABLE IF NOT EXISTS `artworks` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '作品ID',
  `title` varchar(100) NOT NULL COMMENT '作品标题',
  `description` text COMMENT '作品描述',
  `creator_name` varchar(50) DEFAULT NULL COMMENT '创作者名称',
  `media_type` enum('image','video','audio','3d') NOT NULL DEFAULT 'image' COMMENT '媒体类型',
  `media_url` varchar(255) NOT NULL COMMENT '媒体URL',
  `thumbnail_url` varchar(255) DEFAULT NULL COMMENT '缩略图URL',
  `total_supply` int unsigned NOT NULL COMMENT '总发行量',
  `minted_count` int unsigned NOT NULL DEFAULT '0' COMMENT '已铸造数量',
  `price` decimal(20,8) NOT NULL COMMENT '价格（积分）',
  `source` enum('jingtan','platform','community','waveup') NOT NULL DEFAULT 'platform' COMMENT '来源',
  `series` varchar(50) DEFAULT NULL COMMENT '系列名称',
  `release_date` timestamp NULL DEFAULT NULL COMMENT '发售时间',
  `status` enum('draft','pending','active','sold_out','archived') NOT NULL DEFAULT 'draft' COMMENT '状态',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_source` (`source`),
  KEY `idx_status` (`status`),
  KEY `idx_release_date` (`release_date`),
  KEY `idx_series` (`series`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='作品表';

-- 作品实例表（原asset_instances表）
CREATE TABLE IF NOT EXISTS `artwork_instances` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '实例ID',
  `artwork_id` bigint unsigned NOT NULL COMMENT '作品ID',
  `owner_id` bigint unsigned NOT NULL COMMENT '所有者ID',
  `serial_number` varchar(50) NOT NULL COMMENT '序列号',
  `minted_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '铸造时间',
  `status` enum('owned','listed','locked') NOT NULL DEFAULT 'owned' COMMENT '状态',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_serial_number` (`serial_number`),
  KEY `idx_artwork_id` (`artwork_id`),
  KEY `idx_owner_id` (`owner_id`),
  KEY `idx_status` (`status`),
  CONSTRAINT `fk_artwork_instances_artwork` FOREIGN KEY (`artwork_id`) REFERENCES `artworks` (`id`) ON DELETE CASCADE,
  CONSTRAINT `fk_artwork_instances_owner` FOREIGN KEY (`owner_id`) REFERENCES `users` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='作品实例表';

-- ============================================
-- 3. 创作相关表（新增）
-- ============================================

-- 创作表
CREATE TABLE IF NOT EXISTS `creations` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '创作ID',
  `user_id` bigint unsigned NOT NULL COMMENT '创作者ID',
  `title` varchar(100) NOT NULL COMMENT '作品标题',
  `description` text COMMENT '作品描述',
  `media_url` varchar(255) NOT NULL COMMENT '作品图片URL',
  `thumbnail_url` varchar(255) DEFAULT NULL COMMENT '缩略图URL',
  `total_supply` int unsigned NOT NULL COMMENT '发行数量',
  `price` decimal(20,8) NOT NULL COMMENT '售价（积分）',
  `commission_rate` decimal(5,2) NOT NULL DEFAULT '40.00' COMMENT '平台分成比例（%）',
  `status` enum('pending','approved','rejected','published') NOT NULL DEFAULT 'pending' COMMENT '状态',
  `reject_reason` varchar(255) DEFAULT NULL COMMENT '拒绝原因',
  `审核员_id` bigint unsigned DEFAULT NULL COMMENT '审核员ID',
  `审核_at` timestamp NULL DEFAULT NULL COMMENT '审核时间',
  `published_at` timestamp NULL DEFAULT NULL COMMENT '发布时间',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_status` (`status`),
  KEY `idx_created_at` (`created_at`),
  CONSTRAINT `fk_creations_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='创作表';

-- ============================================
-- 4. 交易相关表
-- ============================================

-- 挂单表（保持不变，但关联改为artworks）
CREATE TABLE IF NOT EXISTS `listings` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '挂单ID',
  `seller_id` bigint unsigned NOT NULL COMMENT '卖家ID',
  `artwork_instance_id` bigint unsigned NOT NULL COMMENT '作品实例ID',
  `price` decimal(20,8) NOT NULL COMMENT '价格',
  `status` enum('active','sold','cancelled') NOT NULL DEFAULT 'active' COMMENT '状态',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_seller_id` (`seller_id`),
  KEY `idx_artwork_instance_id` (`artwork_instance_id`),
  KEY `idx_status` (`status`),
  CONSTRAINT `fk_listings_seller` FOREIGN KEY (`seller_id`) REFERENCES `users` (`id`) ON DELETE CASCADE,
  CONSTRAINT `fk_listings_artwork_instance` FOREIGN KEY (`artwork_instance_id`) REFERENCES `artwork_instances` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='挂单表';

-- 交易表（保持不变）
CREATE TABLE IF NOT EXISTS `trades` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '交易ID',
  `seller_id` bigint unsigned NOT NULL COMMENT '卖家ID',
  `buyer_id` bigint unsigned NOT NULL COMMENT '买家ID',
  `artwork_instance_id` bigint unsigned NOT NULL COMMENT '作品实例ID',
  `price` decimal(20,8) NOT NULL COMMENT '成交价格',
  `fee` decimal(20,8) NOT NULL DEFAULT '0.00000000' COMMENT '手续费',
  `type` enum('direct','listing','offer') NOT NULL DEFAULT 'listing' COMMENT '交易类型',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '交易时间',
  PRIMARY KEY (`id`),
  KEY `idx_seller_id` (`seller_id`),
  KEY `idx_buyer_id` (`buyer_id`),
  KEY `idx_artwork_instance_id` (`artwork_instance_id`),
  KEY `idx_created_at` (`created_at`),
  CONSTRAINT `fk_trades_seller` FOREIGN KEY (`seller_id`) REFERENCES `users` (`id`) ON DELETE CASCADE,
  CONSTRAINT `fk_trades_buyer` FOREIGN KEY (`buyer_id`) REFERENCES `users` (`id`) ON DELETE CASCADE,
  CONSTRAINT `fk_trades_artwork_instance` FOREIGN KEY (`artwork_instance_id`) REFERENCES `artwork_instances` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='交易表';

-- 出价表（新增 - 心愿单功能）
CREATE TABLE IF NOT EXISTS `offers` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '出价ID',
  `buyer_id` bigint unsigned NOT NULL COMMENT '买家ID',
  `artwork_instance_id` bigint unsigned NOT NULL COMMENT '作品实例ID',
  `price` decimal(20,8) NOT NULL COMMENT '出价',
  `status` enum('pending','accepted','rejected','cancelled','expired') NOT NULL DEFAULT 'pending' COMMENT '状态',
  `expires_at` timestamp NULL DEFAULT NULL COMMENT '过期时间',
  `responded_at` timestamp NULL DEFAULT NULL COMMENT '响应时间',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_buyer_id` (`buyer_id`),
  KEY `idx_artwork_instance_id` (`artwork_instance_id`),
  KEY `idx_status` (`status`),
  KEY `idx_expires_at` (`expires_at`),
  CONSTRAINT `fk_offers_buyer` FOREIGN KEY (`buyer_id`) REFERENCES `users` (`id`) ON DELETE CASCADE,
  CONSTRAINT `fk_offers_artwork_instance` FOREIGN KEY (`artwork_instance_id`) REFERENCES `artwork_instances` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='出价表';

-- ============================================
-- 5. 任务系统表（新增）
-- ============================================

-- 任务表
CREATE TABLE IF NOT EXISTS `tasks` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '任务ID',
  `name` varchar(50) NOT NULL COMMENT '任务名称',
  `description` varchar(255) DEFAULT NULL COMMENT '任务描述',
  `type` enum('daily','once','achievement') NOT NULL COMMENT '任务类型',
  `reward_points` decimal(20,8) NOT NULL COMMENT '奖励积分',
  `condition_type` varchar(50) NOT NULL COMMENT '完成条件类型',
  `condition_value` varchar(255) DEFAULT NULL COMMENT '完成条件值',
  `is_enabled` tinyint(1) NOT NULL DEFAULT '1' COMMENT '是否启用',
  `sort_order` int NOT NULL DEFAULT '0' COMMENT '排序',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_type` (`type`),
  KEY `idx_is_enabled` (`is_enabled`),
  KEY `idx_sort_order` (`sort_order`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='任务表';

-- 用户任务完成记录表
CREATE TABLE IF NOT EXISTS `user_task_completions` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `user_id` bigint unsigned NOT NULL COMMENT '用户ID',
  `task_id` bigint unsigned NOT NULL COMMENT '任务ID',
  `completed_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '完成时间',
  `claimed_at` timestamp NULL DEFAULT NULL COMMENT '领取时间',
  `reward_points` decimal(20,8) NOT NULL COMMENT '奖励积分',
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_task_id` (`task_id`),
  KEY `idx_completed_at` (`completed_at`),
  CONSTRAINT `fk_user_task_completions_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE,
  CONSTRAINT `fk_user_task_completions_task` FOREIGN KEY (`task_id`) REFERENCES `tasks` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户任务完成记录表';

-- ============================================
-- 6. 公告系统表（新增）
-- ============================================

-- 公告表
CREATE TABLE IF NOT EXISTS `announcements` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '公告ID',
  `title` varchar(100) NOT NULL COMMENT '公告标题',
  `content` text NOT NULL COMMENT '公告内容',
  `type` enum('system','rule','event','maintenance') NOT NULL DEFAULT 'system' COMMENT '公告类型',
  `priority` enum('low','normal','high','urgent') NOT NULL DEFAULT 'normal' COMMENT '优先级',
  `is_pinned` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否置顶',
  `is_published` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否发布',
  `published_at` timestamp NULL DEFAULT NULL COMMENT '发布时间',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_type` (`type`),
  KEY `idx_is_published` (`is_published`),
  KEY `idx_published_at` (`published_at`),
  KEY `idx_is_pinned` (`is_pinned`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='公告表';

-- ============================================
-- 7. 通知系统表（新增）
-- ============================================

-- 通知表
CREATE TABLE IF NOT EXISTS `notifications` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '通知ID',
  `user_id` bigint unsigned NOT NULL COMMENT '用户ID',
  `type` enum('offer','trade','task','announcement','system') NOT NULL COMMENT '通知类型',
  `title` varchar(100) NOT NULL COMMENT '通知标题',
  `content` varchar(255) NOT NULL COMMENT '通知内容',
  `related_id` bigint unsigned DEFAULT NULL COMMENT '关联ID',
  `is_read` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否已读',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_type` (`type`),
  KEY `idx_is_read` (`is_read`),
  KEY `idx_created_at` (`created_at`),
  CONSTRAINT `fk_notifications_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='通知表';

-- ============================================
-- 8. 阳光账户表（新增）
-- ============================================

-- 平台账户表
CREATE TABLE IF NOT EXISTS `platform_account` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `total_balance` decimal(20,8) NOT NULL DEFAULT '0.00000000' COMMENT '总余额',
  `commission_income` decimal(20,8) NOT NULL DEFAULT '0.00000000' COMMENT '分成收入',
  `fee_income` decimal(20,8) NOT NULL DEFAULT '0.00000000' COMMENT '手续费收入',
  `total_expense` decimal(20,8) NOT NULL DEFAULT '0.00000000' COMMENT '总支出',
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='平台账户表';

-- 平台账户交易记录表
CREATE TABLE IF NOT EXISTS `platform_transactions` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '交易ID',
  `type` enum('commission','fee','expense','adjustment') NOT NULL COMMENT '交易类型',
  `amount` decimal(20,8) NOT NULL COMMENT '金额',
  `balance_after` decimal(20,8) NOT NULL COMMENT '交易后余额',
  `description` varchar(255) NOT NULL COMMENT '描述',
  `related_id` bigint unsigned DEFAULT NULL COMMENT '关联ID',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_type` (`type`),
  KEY `idx_created_at` (`created_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='平台账户交易记录表';

-- ============================================
-- 9. 系统配置表（新增）
-- ============================================

-- 系统配置表
CREATE TABLE IF NOT EXISTS `system_configs` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `key` varchar(50) NOT NULL COMMENT '配置键',
  `value` text NOT NULL COMMENT '配置值',
  `description` varchar(255) DEFAULT NULL COMMENT '描述',
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_key` (`key`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='系统配置表';

-- ============================================
-- 10. 积分转账表（预留）
-- ============================================

-- 积分转账表
CREATE TABLE IF NOT EXISTS `point_transfers` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '转账ID',
  `from_user_id` bigint unsigned NOT NULL COMMENT '转出用户ID',
  `to_user_id` bigint unsigned NOT NULL COMMENT '转入用户ID',
  `amount` decimal(20,8) NOT NULL COMMENT '转账金额',
  `fee` decimal(20,8) NOT NULL DEFAULT '0.00000000' COMMENT '手续费',
  `message` varchar(255) DEFAULT NULL COMMENT '留言',
  `status` enum('pending','completed','failed') NOT NULL DEFAULT 'pending' COMMENT '状态',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_from_user_id` (`from_user_id`),
  KEY `idx_to_user_id` (`to_user_id`),
  KEY `idx_status` (`status`),
  KEY `idx_created_at` (`created_at`),
  CONSTRAINT `fk_point_transfers_from_user` FOREIGN KEY (`from_user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE,
  CONSTRAINT `fk_point_transfers_to_user` FOREIGN KEY (`to_user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='积分转账表（预留功能）';

-- ============================================
-- 11. 第三方插件相关表（预留）
-- ============================================

-- 第三方账户表（保持不变）
CREATE TABLE IF NOT EXISTS `third_party_accounts` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `user_id` bigint unsigned NOT NULL COMMENT '用户ID',
  `platform` enum('jingtan','waveup') NOT NULL COMMENT '平台',
  `platform_user_id` varchar(100) NOT NULL COMMENT '平台用户ID',
  `access_token` varchar(255) DEFAULT NULL COMMENT '访问令牌',
  `refresh_token` varchar(255) DEFAULT NULL COMMENT '刷新令牌',
  `expires_at` timestamp NULL DEFAULT NULL COMMENT '过期时间',
  `is_enabled` tinyint(1) NOT NULL DEFAULT '1' COMMENT '是否启用',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_user_platform` (`user_id`,`platform`),
  CONSTRAINT `fk_third_party_accounts_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='第三方账户表';

-- ============================================
-- 12. 初始化数据
-- ============================================

-- 初始化平台账户
INSERT INTO `platform_account` (`id`, `total_balance`, `commission_income`, `fee_income`, `total_expense`) 
VALUES (1, 0.00000000, 0.00000000, 0.00000000, 0.00000000)
ON DUPLICATE KEY UPDATE `id` = 1;

-- 初始化系统配置
INSERT INTO `system_configs` (`key`, `value`, `description`) VALUES
('default_commission_rate', '40.00', '默认平台分成比例（%）'),
('trade_fee_rate', '2.00', '交易手续费比例（%）'),
('offer_expire_days', '7', '出价有效期（天）'),
('daily_signin_points', '0.00001000', '每日签到积分'),
('first_creation_points', '10.00000000', '首次创作奖励积分'),
('first_purchase_points', '5.00000000', '首次购买奖励积分'),
('point_transfer_enabled', '0', '积分转账功能开关（0=关闭，1=开启）'),
('jingtan_plugin_enabled', '0', '鲸探插件开关（0=关闭，1=开启）'),
('waveup_plugin_enabled', '0', 'Waveup插件开关（0=关闭，1=开启）')
ON DUPLICATE KEY UPDATE `key` = VALUES(`key`);

-- 初始化默认任务
INSERT INTO `tasks` (`name`, `description`, `type`, `reward_points`, `condition_type`, `condition_value`, `is_enabled`, `sort_order`) VALUES
('每日签到', '每天签到一次，领取积分奖励', 'daily', 0.00001000, 'daily_signin', NULL, 1, 1),
('首次创作', '首次创作作品并通过审核', 'once', 10.00000000, 'first_creation', NULL, 1, 2),
('首次购买', '首次购买作品', 'once', 5.00000000, 'first_purchase', NULL, 1, 3),
('首次出售', '首次出售作品', 'once', 5.00000000, 'first_sale', NULL, 1, 4),
('完善资料', '完善个人资料（头像+昵称）', 'once', 1.00000000, 'complete_profile', NULL, 1, 5)
ON DUPLICATE KEY UPDATE `name` = VALUES(`name`);
