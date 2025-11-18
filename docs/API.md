# HOHO数字藏品集换小程序 - API设计文档

## 概述

本文档定义了HOHO小程序后端API的规范，包括所有端点、请求/响应格式、错误处理等。

## API基本信息

- **基础URL**：`https://api.hoho.app/api/v1`（生产环境）
- **基础URL**：`http://localhost:8080/api/v1`（开发环境）
- **协议**：HTTPS（生产）/ HTTP（开发）
- **数据格式**：JSON
- **字符编码**：UTF-8

## 通用规范

### 请求格式

所有请求都应包含以下headers：

```
Content-Type: application/json
Authorization: Bearer {token}（需要认证的请求）
```

### 响应格式

所有响应都遵循统一的格式：

```json
{
  "code": 0,
  "message": "success",
  "data": {}
}
```

**字段说明**：
- `code`：状态码（0表示成功，非0表示错误）
- `message`：提示信息
- `data`：返回的数据

### 错误码定义

| 错误码 | HTTP状态码 | 说明 |
| :--- | :--- | :--- |
| 0 | 200 | 成功 |
| 400 | 400 | 请求参数错误 |
| 401 | 401 | 未授权（需要登录） |
| 403 | 403 | 禁止访问 |
| 404 | 404 | 资源不存在 |
| 409 | 409 | 冲突（如重复注册） |
| 500 | 500 | 服务器错误 |
| 1001 | 400 | 用户不存在 |
| 1002 | 400 | 密码错误 |
| 1003 | 400 | 用户已被禁用 |
| 2001 | 400 | 积分不足 |
| 2002 | 400 | 藏品不存在 |
| 2003 | 400 | 挂售单不存在 |
| 2004 | 409 | 藏品已被出售 |

## API端点

### 1. 用户相关API

#### 1.1 用户注册

**端点**：`POST /users/register`

**请求体**：
```json
{
  "phone": "13800138000",
  "password": "password123",
  "confirmPassword": "password123"
}
```

**响应**：
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "id": 1,
    "uid": "U1001",
    "phone": "13800138000",
    "nickname": null,
    "token": "eyJhbGciOiJIUzI1NiIs..."
  }
}
```

**错误情况**：
- 409：手机号已被注册
- 400：密码不符合要求

#### 1.2 用户登录

**端点**：`POST /users/login`

**请求体**：
```json
{
  "phone": "13800138000",
  "password": "password123"
}
```

**响应**：
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "id": 1,
    "uid": "U1001",
    "phone": "13800138000",
    "nickname": "用户昵称",
    "avatar_url": "https://...",
    "token": "eyJhbGciOiJIUzI1NiIs..."
  }
}
```

**错误情况**：
- 1001：用户不存在
- 1002：密码错误
- 1003：用户已被禁用

#### 1.3 获取个人资料

**端点**：`GET /users/profile`

**需要认证**：是

**响应**：
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "id": 1,
    "uid": "U1001",
    "phone": "13800138000",
    "nickname": "用户昵称",
    "avatar_url": "https://...",
    "real_name": "张三",
    "identity_verified": true,
    "status": "active",
    "created_at": "2025-01-01T00:00:00Z"
  }
}
```

#### 1.4 更新个人资料

**端点**：`PUT /users/profile`

**需要认证**：是

**请求体**：
```json
{
  "nickname": "新昵称",
  "avatar_url": "https://..."
}
```

**响应**：同1.3

#### 1.5 实名认证

**端点**：`POST /users/verify-identity`

**需要认证**：是

**请求体**：
```json
{
  "real_name": "张三",
  "id_number": "110101199003071234"
}
```

**响应**：
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "id": 1,
    "real_name": "张三",
    "identity_verified": true
  }
}
```

#### 1.6 获取用户积分

**端点**：`GET /users/points`

**需要认证**：是

**响应**：
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "balance": "10000.00000000",
    "frozen": "100.00000000",
    "total_earned": "50000.00000000",
    "total_spent": "40000.00000000"
  }
}
```

#### 1.7 用户登出

**端点**：`POST /users/logout`

**需要认证**：是

**响应**：
```json
{
  "code": 0,
  "message": "success"
}
```

### 2. 藏品相关API

#### 2.1 获取藏品列表

**端点**：`GET /assets`

**需要认证**：否

**查询参数**：
- `page`：页码（默认1）
- `page_size`：每页数量（默认20）
- `collection_id`：集合ID（可选）
- `status`：状态（可选，approved/all）
- `sort`：排序方式（可选，newest/popular）

**响应**：
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "list": [
      {
        "id": 1,
        "collection_id": 1,
        "name": "真我HOHO",
        "description": "...",
        "image_url": "https://...",
        "total_supply": 30,
        "minted_count": 30,
        "creator_id": 1,
        "status": "approved",
        "created_at": "2025-01-01T00:00:00Z"
      }
    ],
    "total": 100,
    "page": 1,
    "page_size": 20
  }
}
```

#### 2.2 获取藏品详情

**端点**：`GET /assets/{id}`

**需要认证**：否

**响应**：
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "id": 1,
    "collection_id": 1,
    "name": "真我HOHO",
    "description": "...",
    "image_url": "https://...",
    "video_url": "https://...",
    "audio_url": "https://...",
    "total_supply": 30,
    "minted_count": 30,
    "creator": {
      "id": 1,
      "uid": "U1001",
      "nickname": "创作者昵称"
    },
    "status": "approved",
    "created_at": "2025-01-01T00:00:00Z"
  }
}
```

#### 2.3 创建藏品

**端点**：`POST /assets`

**需要认证**：是

**请求体**：
```json
{
  "collection_id": 1,
  "name": "真我HOHO",
  "description": "...",
  "image_url": "https://...",
  "video_url": "https://...",
  "audio_url": "https://...",
  "total_supply": 30
}
```

**响应**：
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "id": 1,
    "collection_id": 1,
    "name": "真我HOHO",
    "status": "pending_review"
  }
}
```

#### 2.4 获取藏品实例列表

**端点**：`GET /assets/{id}/instances`

**需要认证**：否

**查询参数**：
- `owner_id`：所有者ID（可选）
- `status`：状态（可选）

**响应**：
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "list": [
      {
        "id": 1,
        "asset_id": 1,
        "serial_number": 1,
        "owner_id": 2,
        "status": "owned",
        "created_at": "2025-01-01T00:00:00Z"
      }
    ],
    "total": 30
  }
}
```

#### 2.5 获取鲸探资产列表

**端点**：`GET /assets/jingtan`

**需要认证**：是

**响应**：
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "list": [
      {
        "id": 1,
        "jingtan_asset_id": "jt_123456",
        "asset_name": "鲸探资产名称",
        "asset_image_url": "https://...",
        "collection_name": "系列名称",
        "quantity": 1,
        "synced_at": "2025-01-01T00:00:00Z"
      }
    ]
  }
}
```

#### 2.6 同步鲸探资产

**端点**：`POST /assets/jingtan/sync`

**需要认证**：是

**响应**：
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "synced_count": 5,
    "list": [...]
  }
}
```

### 3. 交易相关API

#### 3.1 创建挂售单

**端点**：`POST /listings`

**需要认证**：是

**请求体**：
```json
{
  "asset_instance_id": 1,
  "price": "18500.00000000"
}
```

**响应**：
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "id": 1,
    "seller_id": 1,
    "asset_instance_id": 1,
    "price": "18500.00000000",
    "status": "active",
    "created_at": "2025-01-01T00:00:00Z"
  }
}
```

**错误情况**：
- 2001：积分不足（用于冻结）
- 2002：藏品不存在

#### 3.2 获取挂售单列表

**端点**：`GET /listings`

**需要认证**：否

**查询参数**：
- `asset_id`：藏品ID（可选）
- `seller_id`：卖家ID（可选）
- `status`：状态（可选）
- `sort`：排序方式（可选，price_asc/price_desc/newest）

**响应**：
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "list": [
      {
        "id": 1,
        "seller_id": 1,
        "asset_instance_id": 1,
        "asset": {
          "id": 1,
          "name": "真我HOHO",
          "image_url": "https://..."
        },
        "serial_number": 1,
        "price": "18500.00000000",
        "seller": {
          "id": 1,
          "uid": "U1001",
          "nickname": "卖家昵称"
        },
        "status": "active",
        "created_at": "2025-01-01T00:00:00Z"
      }
    ],
    "total": 100,
    "page": 1
  }
}
```

#### 3.3 取消挂售

**端点**：`POST /listings/{id}/cancel`

**需要认证**：是

**响应**：
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "id": 1,
    "status": "canceled"
  }
}
```

#### 3.4 执行交易

**端点**：`POST /trades/execute`

**需要认证**：是

**请求体**：
```json
{
  "listing_id": 1
}
```

**响应**：
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "id": 1,
    "listing_id": 1,
    "buyer_id": 2,
    "seller_id": 1,
    "price": "18500.00000000",
    "trade_fee": "462.50000000",
    "creator_royalty": "462.50000000",
    "seller_receive": "17575.00000000",
    "status": "completed",
    "completed_at": "2025-01-01T00:00:00Z"
  }
}
```

**错误情况**：
- 2001：积分不足
- 2003：挂售单不存在
- 2004：藏品已被出售

#### 3.5 获取交易历史

**端点**：`GET /trades/history`

**需要认证**：是

**查询参数**：
- `role`：角色（buyer/seller/all）
- `status`：状态（可选）

**响应**：
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "list": [
      {
        "id": 1,
        "buyer_id": 2,
        "seller_id": 1,
        "asset": {
          "id": 1,
          "name": "真我HOHO",
          "image_url": "https://..."
        },
        "serial_number": 1,
        "price": "18500.00000000",
        "status": "completed",
        "completed_at": "2025-01-01T00:00:00Z"
      }
    ],
    "total": 50
  }
}
```

### 4. 积分相关API

#### 4.1 获取积分余额

**端点**：`GET /points/balance`

**需要认证**：是

**响应**：
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "balance": "10000.00000000",
    "frozen": "100.00000000"
  }
}
```

#### 4.2 获取积分流水

**端点**：`GET /points/history`

**需要认证**：是

**查询参数**：
- `type`：类型（earn/spend/all）
- `page`：页码
- `page_size`：每页数量

**响应**：
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "list": [
      {
        "id": 1,
        "type": "earn",
        "amount": "100.00000000",
        "description": "签到奖励",
        "related_type": "checkin",
        "created_at": "2025-01-01T00:00:00Z"
      }
    ],
    "total": 100,
    "page": 1
  }
}
```

### 5. 社区事件API

#### 5.1 获取事件列表

**端点**：`GET /events`

**需要认证**：否

**查询参数**：
- `event_type`：事件类型（可选）
- `page`：页码
- `page_size`：每页数量

**响应**：
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "list": [
      {
        "id": 1,
        "user_id": 1,
        "event_type": "trade",
        "title": "用户uid678910 兑换了 作品野生HOHO",
        "description": "...",
        "related_type": "trade",
        "related_id": 1,
        "created_at": "2025-01-01T00:00:00Z"
      }
    ],
    "total": 1000,
    "page": 1
  }
}
```

### 6. 第三方账户API

#### 6.1 获取第三方账户列表

**端点**：`GET /third-party`

**需要认证**：是

**响应**：
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "list": [
      {
        "id": 1,
        "platform": "jingtan",
        "platform_username": "鲸探用户名",
        "created_at": "2025-01-01T00:00:00Z"
      }
    ]
  }
}
```

#### 6.2 绑定第三方账户

**端点**：`POST /third-party/bind`

**需要认证**：是

**请求体**：
```json
{
  "platform": "jingtan",
  "authorization_code": "..."
}
```

**响应**：
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "id": 1,
    "platform": "jingtan",
    "platform_username": "鲸探用户名"
  }
}
```

#### 6.3 解绑第三方账户

**端点**：`POST /third-party/{platform}/unbind`

**需要认证**：是

**响应**：
```json
{
  "code": 0,
  "message": "success"
}
```

## 认证机制

### Token获取

登录成功后，服务器返回JWT token，客户端应将其保存在本地存储中。

### Token使用

在需要认证的请求中，将token放在Authorization header中：

```
Authorization: Bearer {token}
```

### Token过期

如果服务器返回401错误，表示token已过期，客户端应：
1. 清除本地保存的token
2. 跳转到登录页面
3. 提示用户重新登录

## 速率限制

- 每个IP地址每分钟最多1000个请求
- 超过限制会返回429错误

## 数据验证规则

### 手机号

- 必须是11位数字
- 必须以1开头

### 密码

- 长度至少8个字符
- 必须包含字母和数字

### 积分

- 所有积分值都支持8位小数精度
- 最大值：999999999.99999999

## 分页规范

所有列表接口都支持分页，使用以下参数：

- `page`：页码（从1开始）
- `page_size`：每页数量（默认20，最大100）

响应中包含：

- `list`：数据列表
- `total`：总数
- `page`：当前页码
- `page_size`：每页数量

## 排序规范

支持排序的接口会在查询参数中提供`sort`参数，常见的排序方式：

- `newest`：最新优先
- `oldest`：最旧优先
- `price_asc`：价格从低到高
- `price_desc`：价格从高到低
- `popular`：热度优先

## 文件上传

### 上传端点

**端点**：`POST /upload`

**需要认证**：是

**请求格式**：multipart/form-data

**参数**：
- `file`：文件内容
- `type`：文件类型（image/video/audio）

**响应**：
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "url": "https://...",
    "filename": "xxx.jpg"
  }
}
```

### 文件限制

- 图片：最大50MB，格式jpg/jpeg/png/gif/webp
- 视频：最大100MB，格式mp4/avi/mov/webm
- 音频：最大100MB，格式mp3/wav/m4a

## 版本控制

API版本通过URL路径指定：`/api/v1/...`

未来如果有不兼容的更新，会创建新的版本如`/api/v2/...`，旧版本会保留一段时间用于过渡。
