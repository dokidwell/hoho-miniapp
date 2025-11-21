-- HOHO小程序数据库初始化脚本
-- 创建日期: 2025-11-19

-- 使用数据库
USE hoho_miniapp;

-- 1. 用户表
CREATE TABLE IF NOT EXISTS users (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    uid VARCHAR(20) UNIQUE NOT NULL COMMENT '用户唯一标识',
    phone VARCHAR(20) UNIQUE NOT NULL COMMENT '手机号',
    password_hash VARCHAR(255) NOT NULL COMMENT '密码哈希',
    nickname VARCHAR(100) COMMENT '昵称',
    avatar_url VARCHAR(500) COMMENT '头像URL',
    real_name VARCHAR(100) COMMENT '真实姓名',
    id_number VARCHAR(50) COMMENT '身份证号',
    identity_verified BOOLEAN DEFAULT FALSE COMMENT '是否实名认证',
    status ENUM('active', 'suspended', 'banned') DEFAULT 'active' COMMENT '账户状态',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    INDEX idx_phone (phone),
    INDEX idx_uid (uid),
    INDEX idx_status (status)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户表';

-- 2. 用户积分表
CREATE TABLE IF NOT EXISTS user_points (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    user_id BIGINT UNSIGNED UNIQUE NOT NULL COMMENT '用户ID',
    balance DECIMAL(30,8) DEFAULT 0.00000000 COMMENT '可用积分余额',
    frozen DECIMAL(30,8) DEFAULT 0.00000000 COMMENT '冻结积分',
    total_earned DECIMAL(30,8) DEFAULT 0.00000000 COMMENT '累计获得',
    total_spent DECIMAL(30,8) DEFAULT 0.00000000 COMMENT '累计消费',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    INDEX idx_user_id (user_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户积分表';

-- 3. 积分交易记录表
CREATE TABLE IF NOT EXISTS point_transactions (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    user_id BIGINT UNSIGNED NOT NULL COMMENT '用户ID',
    type ENUM('earn', 'spend', 'adjust', 'freeze', 'unfreeze') NOT NULL COMMENT '交易类型',
    amount DECIMAL(30,8) NOT NULL COMMENT '交易金额',
    description VARCHAR(255) COMMENT '交易描述',
    related_id BIGINT UNSIGNED COMMENT '关联ID',
    related_type VARCHAR(50) COMMENT '关联类型',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    INDEX idx_user_id (user_id),
    INDEX idx_type (type),
    INDEX idx_related (related_id, related_type),
    INDEX idx_created_at (created_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='积分交易记录表';

-- 4. 藏品集合表
CREATE TABLE IF NOT EXISTS collections (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL COMMENT '集合名称',
    description VARCHAR(500) COMMENT '集合描述',
    cover_image VARCHAR(500) COMMENT '封面图片',
    status ENUM('active', 'inactive') DEFAULT 'active' COMMENT '状态',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    INDEX idx_status (status)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='藏品集合表';

-- 5. 藏品表 (SKU)
CREATE TABLE IF NOT EXISTS assets (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    collection_id BIGINT UNSIGNED NOT NULL COMMENT '所属集合ID',
    name VARCHAR(100) NOT NULL COMMENT '藏品名称',
    description VARCHAR(500) COMMENT '藏品描述',
    media_url VARCHAR(500) NOT NULL COMMENT '媒体URL',
    media_type ENUM('image', 'video', 'audio') NOT NULL COMMENT '媒体类型',
    total_supply INT NOT NULL COMMENT '总发行量',
    minted_count INT DEFAULT 0 COMMENT '已铸造数量',
    creator_id BIGINT UNSIGNED NOT NULL COMMENT '创作者ID',
    status ENUM('pending_review', 'approved', 'rejected', 'active', 'inactive') DEFAULT 'pending_review' COMMENT '状态',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    FOREIGN KEY (collection_id) REFERENCES collections(id),
    FOREIGN KEY (creator_id) REFERENCES users(id),
    INDEX idx_collection (collection_id),
    INDEX idx_creator (creator_id),
    INDEX idx_status (status)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='藏品表';

-- 6. 藏品实例表
CREATE TABLE IF NOT EXISTS asset_instances (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    asset_id BIGINT UNSIGNED NOT NULL COMMENT '藏品ID',
    instance_no INT NOT NULL COMMENT '实例编号',
    owner_id BIGINT UNSIGNED NOT NULL COMMENT '持有者ID',
    token_id VARCHAR(255) UNIQUE NOT NULL COMMENT '唯一TokenID',
    status ENUM('in_wallet', 'on_sale', 'pending_trade', 'burned') DEFAULT 'in_wallet' COMMENT '状态',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    FOREIGN KEY (asset_id) REFERENCES assets(id),
    FOREIGN KEY (owner_id) REFERENCES users(id),
    INDEX idx_asset (asset_id),
    INDEX idx_owner (owner_id),
    INDEX idx_token (token_id),
    INDEX idx_status (status)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='藏品实例表';

-- 7. 交易挂单表
CREATE TABLE IF NOT EXISTS listings (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    instance_id BIGINT UNSIGNED NOT NULL COMMENT '藏品实例ID',
    seller_id BIGINT UNSIGNED NOT NULL COMMENT '卖家ID',
    price DECIMAL(30,8) NOT NULL COMMENT '挂单价格',
    status ENUM('active', 'sold', 'cancelled', 'expired') DEFAULT 'active' COMMENT '状态',
    expires_at TIMESTAMP NULL COMMENT '过期时间',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    FOREIGN KEY (instance_id) REFERENCES asset_instances(id),
    FOREIGN KEY (seller_id) REFERENCES users(id),
    INDEX idx_instance (instance_id),
    INDEX idx_seller (seller_id),
    INDEX idx_status (status),
    INDEX idx_price (price)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='交易挂单表';

-- 8. 交易记录表
CREATE TABLE IF NOT EXISTS trades (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    listing_id BIGINT UNSIGNED NOT NULL COMMENT '挂单ID',
    instance_id BIGINT UNSIGNED NOT NULL COMMENT '藏品实例ID',
    seller_id BIGINT UNSIGNED NOT NULL COMMENT '卖家ID',
    buyer_id BIGINT UNSIGNED NOT NULL COMMENT '买家ID',
    price DECIMAL(30,8) NOT NULL COMMENT '成交价格',
    status ENUM('pending', 'completed', 'failed', 'cancelled') DEFAULT 'pending' COMMENT '状态',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    FOREIGN KEY (listing_id) REFERENCES listings(id),
    FOREIGN KEY (instance_id) REFERENCES asset_instances(id),
    FOREIGN KEY (seller_id) REFERENCES users(id),
    FOREIGN KEY (buyer_id) REFERENCES users(id),
    INDEX idx_listing (listing_id),
    INDEX idx_seller (seller_id),
    INDEX idx_buyer (buyer_id),
    INDEX idx_status (status),
    INDEX idx_created_at (created_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='交易记录表';

-- 9. 第三方账户绑定表
CREATE TABLE IF NOT EXISTS third_party_accounts (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    user_id BIGINT UNSIGNED NOT NULL COMMENT '用户ID',
    platform VARCHAR(50) NOT NULL COMMENT '平台名称',
    platform_uid VARCHAR(255) NOT NULL COMMENT '平台用户ID',
    platform_username VARCHAR(255) COMMENT '平台用户名',
    access_token VARCHAR(500) COMMENT '访问令牌',
    refresh_token VARCHAR(500) COMMENT '刷新令牌',
    token_expires_at TIMESTAMP NULL COMMENT '令牌过期时间',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    UNIQUE KEY uk_user_platform (user_id, platform),
    INDEX idx_platform (platform),
    INDEX idx_platform_uid (platform_uid)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='第三方账户绑定表';

-- 10. 鲸探资产映射表
CREATE TABLE IF NOT EXISTS jingtan_assets (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    user_id BIGINT UNSIGNED NOT NULL COMMENT '用户ID',
    jingtan_asset_id VARCHAR(255) UNIQUE NOT NULL COMMENT '鲸探资产ID',
    name VARCHAR(100) COMMENT '资产名称',
    image_url VARCHAR(500) COMMENT '资产图片',
    status ENUM('active', 'inactive') DEFAULT 'active' COMMENT '状态',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    INDEX idx_user (user_id),
    INDEX idx_jingtan_asset (jingtan_asset_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='鲸探资产映射表';

-- 11. 社区事件表
CREATE TABLE IF NOT EXISTS community_events (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    event_type VARCHAR(50) NOT NULL COMMENT '事件类型',
    title VARCHAR(200) NOT NULL COMMENT '事件标题',
    description TEXT COMMENT '事件描述',
    user_id BIGINT UNSIGNED COMMENT '关联用户ID',
    related_id BIGINT UNSIGNED COMMENT '关联ID',
    related_type VARCHAR(50) COMMENT '关联类型',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    INDEX idx_event_type (event_type),
    INDEX idx_user (user_id),
    INDEX idx_related (related_id, related_type),
    INDEX idx_created_at (created_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='社区事件表';

-- 12. 管理员表
CREATE TABLE IF NOT EXISTS admins (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL COMMENT '用户名',
    password_hash VARCHAR(255) NOT NULL COMMENT '密码哈希',
    email VARCHAR(100) COMMENT '邮箱',
    role ENUM('super_admin', 'admin', 'operator') DEFAULT 'operator' COMMENT '角色',
    status ENUM('active', 'suspended') DEFAULT 'active' COMMENT '状态',
    last_login_at TIMESTAMP NULL COMMENT '最后登录时间',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    INDEX idx_username (username),
    INDEX idx_status (status)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='管理员表';

-- 插入默认管理员账户 (密码: Admin@123456)
-- 密码哈希使用bcrypt生成，需要在应用层生成
-- INSERT INTO admins (username, password_hash, email, role, status) 
-- VALUES ('admin', '$2a$10$...', 'admin@hohopark.com', 'super_admin', 'active');

-- 创建完成
SELECT '✅ 数据库表创建完成！' AS message;
