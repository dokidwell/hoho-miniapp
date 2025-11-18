-- HOHO数字藏品集换小程序 - 数据库初始化脚本

-- ==================== 用户相关表 ====================

-- 用户表
CREATE TABLE IF NOT EXISTS users (
  id BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT '用户ID',
  uid VARCHAR(20) UNIQUE NOT NULL COMMENT '用户UID',
  phone VARCHAR(20) UNIQUE NOT NULL COMMENT '手机号',
  password_hash VARCHAR(255) NOT NULL COMMENT '密码哈希',
  nickname VARCHAR(100) COMMENT '昵称',
  avatar_url VARCHAR(500) COMMENT '头像URL',
  real_name VARCHAR(100) COMMENT '真实姓名',
  id_number VARCHAR(50) COMMENT '身份证号',
  identity_verified BOOLEAN DEFAULT FALSE COMMENT '是否实名认证',
  status ENUM('active', 'suspended', 'banned') DEFAULT 'active' COMMENT '账户状态',
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  INDEX idx_phone (phone),
  INDEX idx_uid (uid),
  INDEX idx_created_at (created_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户表';

-- 用户积分表
CREATE TABLE IF NOT EXISTS user_points (
  id BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT '积分ID',
  user_id BIGINT NOT NULL UNIQUE COMMENT '用户ID',
  balance DECIMAL(30, 8) DEFAULT 0 COMMENT '积分余额（8位小数）',
  frozen DECIMAL(30, 8) DEFAULT 0 COMMENT '冻结积分（待交易确认）',
  total_earned DECIMAL(30, 8) DEFAULT 0 COMMENT '总获得积分',
  total_spent DECIMAL(30, 8) DEFAULT 0 COMMENT '总消费积分',
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
  INDEX idx_user_id (user_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户积分表';

-- 积分流水表
CREATE TABLE IF NOT EXISTS point_transactions (
  id BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT '流水ID',
  user_id BIGINT NOT NULL COMMENT '用户ID',
  type ENUM('earn', 'spend', 'adjust', 'freeze', 'unfreeze') COMMENT '流水类型',
  amount DECIMAL(30, 8) NOT NULL COMMENT '金额（8位小数）',
  description VARCHAR(255) COMMENT '描述',
  related_id BIGINT COMMENT '关联ID（如交易ID、铸造ID等）',
  related_type VARCHAR(50) COMMENT '关联类型（如trade、mint等）',
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
  INDEX idx_user_id (user_id),
  INDEX idx_created_at (created_at),
  INDEX idx_related (related_type, related_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='积分流水表';

-- 第三方账户关联表
CREATE TABLE IF NOT EXISTS third_party_accounts (
  id BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT '关联ID',
  user_id BIGINT NOT NULL COMMENT '用户ID',
  platform VARCHAR(50) NOT NULL COMMENT '平台（jingtan, waveup等）',
  platform_uid VARCHAR(255) NOT NULL COMMENT '平台用户ID',
  platform_username VARCHAR(255) COMMENT '平台用户名',
  access_token VARCHAR(500) COMMENT '访问令牌',
  refresh_token VARCHAR(500) COMMENT '刷新令牌',
  token_expires_at TIMESTAMP COMMENT '令牌过期时间',
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
  UNIQUE KEY unique_platform (user_id, platform),
  INDEX idx_user_id (user_id),
  INDEX idx_platform (platform)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='第三方账户关联表';

-- ==================== 藏品相关表 ====================

-- 集合/系列表
CREATE TABLE IF NOT EXISTS collections (
  id BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT '集合ID',
  name VARCHAR(255) NOT NULL COMMENT '集合名称',
  description TEXT COMMENT '集合描述',
  creator_id BIGINT NOT NULL COMMENT '创作者ID',
  cover_image_url VARCHAR(500) COMMENT '封面图URL',
  status ENUM('draft', 'published', 'archived') DEFAULT 'draft' COMMENT '状态',
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  FOREIGN KEY (creator_id) REFERENCES users(id) ON DELETE CASCADE,
  INDEX idx_creator_id (creator_id),
  INDEX idx_status (status),
  INDEX idx_created_at (created_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='集合/系列表';

-- 藏品模板表
CREATE TABLE IF NOT EXISTS assets (
  id BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT '藏品ID',
  collection_id BIGINT NOT NULL COMMENT '所属集合ID',
  name VARCHAR(255) NOT NULL COMMENT '藏品名称',
  description TEXT COMMENT '藏品描述',
  image_url VARCHAR(500) COMMENT '藏品图片URL',
  video_url VARCHAR(500) COMMENT '藏品视频URL',
  audio_url VARCHAR(500) COMMENT '藏品音频URL',
  total_supply INT NOT NULL DEFAULT 1 COMMENT '总供应量',
  minted_count INT DEFAULT 0 COMMENT '已铸造数量',
  creator_id BIGINT NOT NULL COMMENT '创作者ID',
  status ENUM('draft', 'pending_review', 'approved', 'rejected', 'archived') DEFAULT 'draft' COMMENT '状态',
  review_notes TEXT COMMENT '审核意见',
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  FOREIGN KEY (collection_id) REFERENCES collections(id) ON DELETE CASCADE,
  FOREIGN KEY (creator_id) REFERENCES users(id) ON DELETE CASCADE,
  INDEX idx_collection_id (collection_id),
  INDEX idx_creator_id (creator_id),
  INDEX idx_status (status),
  INDEX idx_created_at (created_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='藏品模板表';

-- 藏品实例表（具体的编号）
CREATE TABLE IF NOT EXISTS asset_instances (
  id BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT '实例ID',
  asset_id BIGINT NOT NULL COMMENT '所属藏品ID',
  serial_number INT NOT NULL COMMENT '编号',
  owner_id BIGINT NOT NULL COMMENT '当前所有者ID',
  status ENUM('owned', 'on_sale', 'burned') DEFAULT 'owned' COMMENT '状态',
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  FOREIGN KEY (asset_id) REFERENCES assets(id) ON DELETE CASCADE,
  FOREIGN KEY (owner_id) REFERENCES users(id) ON DELETE CASCADE,
  UNIQUE KEY unique_serial (asset_id, serial_number),
  INDEX idx_asset_id (asset_id),
  INDEX idx_owner_id (owner_id),
  INDEX idx_status (status),
  INDEX idx_created_at (created_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='藏品实例表';

-- 鲸探映射资产表（只读，用于展示）
CREATE TABLE IF NOT EXISTS jingtan_assets (
  id BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT '映射ID',
  user_id BIGINT NOT NULL COMMENT '用户ID',
  jingtan_asset_id VARCHAR(255) NOT NULL COMMENT '鲸探资产ID',
  asset_name VARCHAR(255) COMMENT '资产名称',
  asset_image_url VARCHAR(500) COMMENT '资产图片URL',
  collection_name VARCHAR(255) COMMENT '系列名称',
  quantity INT DEFAULT 1 COMMENT '数量',
  synced_at TIMESTAMP COMMENT '同步时间',
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
  UNIQUE KEY unique_jingtan (user_id, jingtan_asset_id),
  INDEX idx_user_id (user_id),
  INDEX idx_synced_at (synced_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='鲸探映射资产表';

-- ==================== 交易相关表 ====================

-- 挂售单表
CREATE TABLE IF NOT EXISTS listings (
  id BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT '挂售单ID',
  seller_id BIGINT NOT NULL COMMENT '卖家ID',
  asset_instance_id BIGINT NOT NULL COMMENT '藏品实例ID',
  price DECIMAL(30, 8) NOT NULL COMMENT '价格（8位小数）',
  quantity INT DEFAULT 1 COMMENT '数量',
  status ENUM('active', 'sold', 'canceled') DEFAULT 'active' COMMENT '状态',
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  FOREIGN KEY (seller_id) REFERENCES users(id) ON DELETE CASCADE,
  FOREIGN KEY (asset_instance_id) REFERENCES asset_instances(id) ON DELETE CASCADE,
  INDEX idx_seller_id (seller_id),
  INDEX idx_status (status),
  INDEX idx_created_at (created_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='挂售单表';

-- 交易表
CREATE TABLE IF NOT EXISTS trades (
  id BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT '交易ID',
  listing_id BIGINT NOT NULL COMMENT '挂售单ID',
  buyer_id BIGINT NOT NULL COMMENT '买家ID',
  seller_id BIGINT NOT NULL COMMENT '卖家ID',
  asset_instance_id BIGINT NOT NULL COMMENT '藏品实例ID',
  price DECIMAL(30, 8) NOT NULL COMMENT '成交价格（8位小数）',
  trade_fee DECIMAL(30, 8) NOT NULL COMMENT '交易手续费（8位小数）',
  creator_royalty DECIMAL(30, 8) NOT NULL COMMENT '创作者版税（8位小数）',
  seller_receive DECIMAL(30, 8) NOT NULL COMMENT '卖家实际收到（8位小数）',
  status ENUM('pending', 'completed', 'failed', 'canceled') DEFAULT 'pending' COMMENT '状态',
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  completed_at TIMESTAMP COMMENT '完成时间',
  FOREIGN KEY (listing_id) REFERENCES listings(id) ON DELETE CASCADE,
  FOREIGN KEY (buyer_id) REFERENCES users(id) ON DELETE CASCADE,
  FOREIGN KEY (seller_id) REFERENCES users(id) ON DELETE CASCADE,
  FOREIGN KEY (asset_instance_id) REFERENCES asset_instances(id) ON DELETE CASCADE,
  INDEX idx_buyer_id (buyer_id),
  INDEX idx_seller_id (seller_id),
  INDEX idx_status (status),
  INDEX idx_created_at (created_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='交易表';

-- ==================== 社区事件表 ====================

-- 社区事件公示表
CREATE TABLE IF NOT EXISTS community_events (
  id BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT '事件ID',
  user_id BIGINT COMMENT '相关用户ID',
  event_type VARCHAR(50) NOT NULL COMMENT '事件类型（mint, trade, airdrop, login等）',
  title VARCHAR(255) NOT NULL COMMENT '事件标题',
  description TEXT COMMENT '事件描述',
  related_id BIGINT COMMENT '关联ID',
  related_type VARCHAR(50) COMMENT '关联类型',
  data JSON COMMENT '事件数据（JSON格式）',
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  INDEX idx_user_id (user_id),
  INDEX idx_event_type (event_type),
  INDEX idx_created_at (created_at),
  INDEX idx_related (related_type, related_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='社区事件公示表';

-- ==================== 后台管理相关表 ====================

-- 管理员表
CREATE TABLE IF NOT EXISTS admins (
  id BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT '管理员ID',
  username VARCHAR(100) UNIQUE NOT NULL COMMENT '用户名',
  password_hash VARCHAR(255) NOT NULL COMMENT '密码哈希',
  email VARCHAR(100) UNIQUE COMMENT '邮箱',
  role ENUM('super_admin', 'admin', 'reviewer', 'customer_service') DEFAULT 'admin' COMMENT '角色',
  status ENUM('active', 'inactive') DEFAULT 'active' COMMENT '状态',
  last_login_at TIMESTAMP COMMENT '最后登录时间',
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  INDEX idx_username (username),
  INDEX idx_role (role)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='管理员表';

-- 系统配置表
CREATE TABLE IF NOT EXISTS system_config (
  id BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT '配置ID',
  key_name VARCHAR(100) UNIQUE NOT NULL COMMENT '配置键',
  key_value TEXT NOT NULL COMMENT '配置值',
  description VARCHAR(255) COMMENT '描述',
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  INDEX idx_key_name (key_name)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='系统配置表';

-- 操作日志表
CREATE TABLE IF NOT EXISTS operation_logs (
  id BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT '日志ID',
  admin_id BIGINT COMMENT '管理员ID',
  operation VARCHAR(100) NOT NULL COMMENT '操作类型',
  target_type VARCHAR(50) COMMENT '目标类型',
  target_id BIGINT COMMENT '目标ID',
  details TEXT COMMENT '操作详情',
  ip_address VARCHAR(50) COMMENT 'IP地址',
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  FOREIGN KEY (admin_id) REFERENCES admins(id) ON DELETE SET NULL,
  INDEX idx_admin_id (admin_id),
  INDEX idx_created_at (created_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='操作日志表';

-- ==================== 初始化数据 ====================

-- 插入系统配置
INSERT INTO system_config (key_name, key_value, description) VALUES
('trade_fee_rate', '0.025', '交易手续费率（2.5%）'),
('creator_royalty_rate', '0.025', '创作者版税率（2.5%）'),
('points_decimal_places', '8', '积分小数位数'),
('platform_name', 'HOHO', '平台名称'),
('jingtan_api_enabled', 'false', '鲸探API是否启用'),
('max_upload_size', '52428800', '最大上传文件大小（50MB）'),
('allowed_image_formats', 'jpg,jpeg,png,gif,webp', '允许的图片格式'),
('allowed_video_formats', 'mp4,avi,mov,webm', '允许的视频格式'),
('allowed_audio_formats', 'mp3,wav,m4a', '允许的音频格式');
