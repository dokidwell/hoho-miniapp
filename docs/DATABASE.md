# HOHO数字藏品集换小程序 - 数据库设计文档

## 概述

本文档详细描述了HOHO小程序的数据库设计，包括所有表的结构、字段定义、关系和索引。

## 数据库选择

- **数据库系统**：MySQL 8.0+
- **字符集**：utf8mb4
- **排序规则**：utf8mb4_unicode_ci
- **引擎**：InnoDB（支持事务）

## 核心设计原则

1. **积分精度**：所有积分相关字段使用 `DECIMAL(30, 8)` 以支持8位小数精度
2. **原子性**：交易相关操作使用数据库事务确保原子性
3. **可追溯性**：所有关键操作都记录在流水表中
4. **性能优化**：为常用查询字段添加索引

## 表结构设计

### 1. 用户相关表

#### 1.1 users（用户表）

存储用户的基本信息。

| 字段名 | 类型 | 约束 | 说明 |
| :--- | :--- | :--- | :--- |
| id | BIGINT | PK, AI | 用户ID |
| uid | VARCHAR(20) | UQ | 用户UID（展示给用户的ID） |
| phone | VARCHAR(20) | UQ | 手机号 |
| password_hash | VARCHAR(255) | NN | 密码哈希 |
| nickname | VARCHAR(100) | | 昵称 |
| avatar_url | VARCHAR(500) | | 头像URL |
| real_name | VARCHAR(100) | | 真实姓名 |
| id_number | VARCHAR(50) | | 身份证号 |
| identity_verified | BOOLEAN | DF:FALSE | 是否实名认证 |
| status | ENUM | DF:active | 账户状态（active/suspended/banned） |
| created_at | TIMESTAMP | DF:NOW | 创建时间 |
| updated_at | TIMESTAMP | DF:NOW | 更新时间 |

**索引**：
- `idx_phone`：phone字段
- `idx_uid`：uid字段
- `idx_created_at`：created_at字段

#### 1.2 user_points（用户积分表）

存储用户的积分余额和统计信息。

| 字段名 | 类型 | 约束 | 说明 |
| :--- | :--- | :--- | :--- |
| id | BIGINT | PK, AI | 积分ID |
| user_id | BIGINT | UQ, FK | 用户ID |
| balance | DECIMAL(30,8) | DF:0 | 积分余额（8位小数） |
| frozen | DECIMAL(30,8) | DF:0 | 冻结积分 |
| total_earned | DECIMAL(30,8) | DF:0 | 总获得积分 |
| total_spent | DECIMAL(30,8) | DF:0 | 总消费积分 |
| updated_at | TIMESTAMP | DF:NOW | 更新时间 |

**关键设计**：
- `balance`：用户可用积分
- `frozen`：待交易确认的积分（交易时冻结，完成后转移）
- 使用DECIMAL确保精度，避免浮点数精度问题

#### 1.3 point_transactions（积分流水表）

记录所有积分变动，实现完全可追溯。

| 字段名 | 类型 | 约束 | 说明 |
| :--- | :--- | :--- | :--- |
| id | BIGINT | PK, AI | 流水ID |
| user_id | BIGINT | FK | 用户ID |
| type | ENUM | | 流水类型（earn/spend/adjust/freeze/unfreeze） |
| amount | DECIMAL(30,8) | NN | 金额（8位小数） |
| description | VARCHAR(255) | | 描述 |
| related_id | BIGINT | | 关联ID（如交易ID） |
| related_type | VARCHAR(50) | | 关联类型（如trade、mint） |
| created_at | TIMESTAMP | DF:NOW | 创建时间 |

**索引**：
- `idx_user_id`：user_id字段
- `idx_created_at`：created_at字段
- `idx_related`：related_type和related_id的组合索引

#### 1.4 third_party_accounts（第三方账户关联表）

存储用户与第三方平台（鲸探、Waveup等）的关联信息。

| 字段名 | 类型 | 约束 | 说明 |
| :--- | :--- | :--- | :--- |
| id | BIGINT | PK, AI | 关联ID |
| user_id | BIGINT | FK | 用户ID |
| platform | VARCHAR(50) | NN | 平台名称（jingtan/waveup） |
| platform_uid | VARCHAR(255) | NN | 平台用户ID |
| platform_username | VARCHAR(255) | | 平台用户名 |
| access_token | VARCHAR(500) | | 访问令牌 |
| refresh_token | VARCHAR(500) | | 刷新令牌 |
| token_expires_at | TIMESTAMP | | 令牌过期时间 |
| created_at | TIMESTAMP | DF:NOW | 创建时间 |
| updated_at | TIMESTAMP | DF:NOW | 更新时间 |

**索引**：
- `unique_platform`：user_id和platform的唯一组合索引
- `idx_user_id`：user_id字段
- `idx_platform`：platform字段

### 2. 藏品相关表

#### 2.1 collections（集合/系列表）

存储藏品集合信息。

| 字段名 | 类型 | 约束 | 说明 |
| :--- | :--- | :--- | :--- |
| id | BIGINT | PK, AI | 集合ID |
| name | VARCHAR(255) | NN | 集合名称 |
| description | TEXT | | 集合描述 |
| creator_id | BIGINT | FK | 创作者ID |
| cover_image_url | VARCHAR(500) | | 封面图URL |
| status | ENUM | DF:draft | 状态（draft/published/archived） |
| created_at | TIMESTAMP | DF:NOW | 创建时间 |
| updated_at | TIMESTAMP | DF:NOW | 更新时间 |

#### 2.2 assets（藏品模板表）

存储藏品的模板信息。一个藏品可以有多个编号的实例。

| 字段名 | 类型 | 约束 | 说明 |
| :--- | :--- | :--- | :--- |
| id | BIGINT | PK, AI | 藏品ID |
| collection_id | BIGINT | FK | 所属集合ID |
| name | VARCHAR(255) | NN | 藏品名称 |
| description | TEXT | | 藏品描述 |
| image_url | VARCHAR(500) | | 藏品图片URL |
| video_url | VARCHAR(500) | | 藏品视频URL |
| audio_url | VARCHAR(500) | | 藏品音频URL |
| total_supply | INT | NN, DF:1 | 总供应量 |
| minted_count | INT | DF:0 | 已铸造数量 |
| creator_id | BIGINT | FK | 创作者ID |
| status | ENUM | DF:draft | 状态（draft/pending_review/approved/rejected/archived） |
| review_notes | TEXT | | 审核意见 |
| created_at | TIMESTAMP | DF:NOW | 创建时间 |
| updated_at | TIMESTAMP | DF:NOW | 更新时间 |

**设计说明**：
- `total_supply`：设计时指定的总数量
- `minted_count`：已经铸造出来的数量
- `status`：用于审核流程管理

#### 2.3 asset_instances（藏品实例表）

存储具体的藏品编号。

| 字段名 | 类型 | 约束 | 说明 |
| :--- | :--- | :--- | :--- |
| id | BIGINT | PK, AI | 实例ID |
| asset_id | BIGINT | FK | 所属藏品ID |
| serial_number | INT | NN | 编号 |
| owner_id | BIGINT | FK | 当前所有者ID |
| status | ENUM | DF:owned | 状态（owned/on_sale/burned） |
| created_at | TIMESTAMP | DF:NOW | 创建时间 |
| updated_at | TIMESTAMP | DF:NOW | 更新时间 |

**索引**：
- `unique_serial`：asset_id和serial_number的唯一组合索引

**设计说明**：
- 每个藏品的每个编号都是一个独立的实例
- `owner_id`：当前所有者，交易时更新此字段
- `status`：用于标记藏品状态

#### 2.4 jingtan_assets（鲸探映射资产表）

存储从鲸探同步过来的资产信息（只读展示）。

| 字段名 | 类型 | 约束 | 说明 |
| :--- | :--- | :--- | :--- |
| id | BIGINT | PK, AI | 映射ID |
| user_id | BIGINT | FK | 用户ID |
| jingtan_asset_id | VARCHAR(255) | NN | 鲸探资产ID |
| asset_name | VARCHAR(255) | | 资产名称 |
| asset_image_url | VARCHAR(500) | | 资产图片URL |
| collection_name | VARCHAR(255) | | 系列名称 |
| quantity | INT | DF:1 | 数量 |
| synced_at | TIMESTAMP | | 同步时间 |
| created_at | TIMESTAMP | DF:NOW | 创建时间 |

**索引**：
- `unique_jingtan`：user_id和jingtan_asset_id的唯一组合索引

### 3. 交易相关表

#### 3.1 listings（挂售单表）

存储用户的挂售信息。

| 字段名 | 类型 | 约束 | 说明 |
| :--- | :--- | :--- | :--- |
| id | BIGINT | PK, AI | 挂售单ID |
| seller_id | BIGINT | FK | 卖家ID |
| asset_instance_id | BIGINT | FK | 藏品实例ID |
| price | DECIMAL(30,8) | NN | 价格（8位小数） |
| quantity | INT | DF:1 | 数量 |
| status | ENUM | DF:active | 状态（active/sold/canceled） |
| created_at | TIMESTAMP | DF:NOW | 创建时间 |
| updated_at | TIMESTAMP | DF:NOW | 更新时间 |

#### 3.2 trades（交易表）

记录所有交易信息，包括手续费和版税计算。

| 字段名 | 类型 | 约束 | 说明 |
| :--- | :--- | :--- | :--- |
| id | BIGINT | PK, AI | 交易ID |
| listing_id | BIGINT | FK | 挂售单ID |
| buyer_id | BIGINT | FK | 买家ID |
| seller_id | BIGINT | FK | 卖家ID |
| asset_instance_id | BIGINT | FK | 藏品实例ID |
| price | DECIMAL(30,8) | NN | 成交价格（8位小数） |
| trade_fee | DECIMAL(30,8) | NN | 交易手续费（8位小数） |
| creator_royalty | DECIMAL(30,8) | NN | 创作者版税（8位小数） |
| seller_receive | DECIMAL(30,8) | NN | 卖家实际收到（8位小数） |
| status | ENUM | DF:pending | 状态（pending/completed/failed/canceled） |
| created_at | TIMESTAMP | DF:NOW | 创建时间 |
| completed_at | TIMESTAMP | | 完成时间 |

**设计说明**：
- `price`：成交价格
- `trade_fee`：平台手续费（2.5%）
- `creator_royalty`：创作者版税（2.5%）
- `seller_receive = price - trade_fee - creator_royalty`
- 所有字段都使用DECIMAL(30,8)确保精度

### 4. 社区事件表

#### 4.1 community_events（社区事件公示表）

记录所有关键事件，实现透明公示功能。

| 字段名 | 类型 | 约束 | 说明 |
| :--- | :--- | :--- | :--- |
| id | BIGINT | PK, AI | 事件ID |
| user_id | BIGINT | | 相关用户ID |
| event_type | VARCHAR(50) | NN | 事件类型（mint/trade/airdrop/login等） |
| title | VARCHAR(255) | NN | 事件标题 |
| description | TEXT | | 事件描述 |
| related_id | BIGINT | | 关联ID |
| related_type | VARCHAR(50) | | 关联类型 |
| data | JSON | | 事件数据（JSON格式） |
| created_at | TIMESTAMP | DF:NOW | 创建时间 |

**事件类型示例**：
- `mint`：藏品铸造
- `trade`：藏品交易
- `airdrop`：空投
- `login`：用户登录
- `verify`：实名认证

### 5. 后台管理表

#### 5.1 admins（管理员表）

存储后台管理员信息。

| 字段名 | 类型 | 约束 | 说明 |
| :--- | :--- | :--- | :--- |
| id | BIGINT | PK, AI | 管理员ID |
| username | VARCHAR(100) | UQ, NN | 用户名 |
| password_hash | VARCHAR(255) | NN | 密码哈希 |
| email | VARCHAR(100) | UQ | 邮箱 |
| role | ENUM | DF:admin | 角色（super_admin/admin/reviewer/customer_service） |
| status | ENUM | DF:active | 状态（active/inactive） |
| last_login_at | TIMESTAMP | | 最后登录时间 |
| created_at | TIMESTAMP | DF:NOW | 创建时间 |
| updated_at | TIMESTAMP | DF:NOW | 更新时间 |

#### 5.2 system_config（系统配置表）

存储平台的配置参数。

| 字段名 | 类型 | 约束 | 说明 |
| :--- | :--- | :--- | :--- |
| id | BIGINT | PK, AI | 配置ID |
| key_name | VARCHAR(100) | UQ, NN | 配置键 |
| key_value | TEXT | NN | 配置值 |
| description | VARCHAR(255) | | 描述 |
| updated_at | TIMESTAMP | DF:NOW | 更新时间 |

**常见配置**：
- `trade_fee_rate`：交易手续费率
- `creator_royalty_rate`：创作者版税率
- `points_decimal_places`：积分小数位数
- `jingtan_api_enabled`：鲸探API是否启用

#### 5.3 operation_logs（操作日志表）

记录所有管理员操作。

| 字段名 | 类型 | 约束 | 说明 |
| :--- | :--- | :--- | :--- |
| id | BIGINT | PK, AI | 日志ID |
| admin_id | BIGINT | FK | 管理员ID |
| operation | VARCHAR(100) | NN | 操作类型 |
| target_type | VARCHAR(50) | | 目标类型 |
| target_id | BIGINT | | 目标ID |
| details | TEXT | | 操作详情 |
| ip_address | VARCHAR(50) | | IP地址 |
| created_at | TIMESTAMP | DF:NOW | 创建时间 |

## 关键业务流程的数据操作

### 交易流程

1. **用户挂售**：
   - 在`listings`表插入一条记录
   - 在`asset_instances`表更新`status`为`on_sale`
   - 在`point_transactions`表记录积分冻结

2. **用户购买**：
   - 在`trades`表插入交易记录
   - 计算`trade_fee`和`creator_royalty`
   - 在`point_transactions`表记录买家积分扣减、卖家积分增加、创作者版税
   - 在`asset_instances`表更新`owner_id`和`status`
   - 在`listings`表更新`status`为`sold`
   - 在`community_events`表记录交易事件

### 藏品铸造流程

1. **创建藏品**：
   - 在`assets`表插入藏品模板
   - `status`设为`pending_review`

2. **审核通过**：
   - 在`assets`表更新`status`为`approved`
   - 根据`total_supply`在`asset_instances`表批量插入实例记录
   - 在`community_events`表记录铸造事件

## 性能优化建议

1. **索引优化**：
   - 为常用查询字段添加索引
   - 对于复合查询，使用组合索引
   - 定期分析和优化索引

2. **查询优化**：
   - 避免全表扫描
   - 使用LIMIT分页
   - 对于大数据量查询，考虑使用分区表

3. **缓存策略**：
   - 使用Redis缓存用户余额
   - 缓存热门藏品信息
   - 缓存交易手续费率等配置

4. **事务管理**：
   - 交易操作使用事务确保原子性
   - 设置合理的隔离级别
   - 避免长事务

## 数据备份和恢复

- 定期备份数据库（建议每天至少一次）
- 保留至少7天的备份
- 定期测试恢复流程
- 建立灾难恢复计划
