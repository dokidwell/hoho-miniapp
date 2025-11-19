# 🔑 API密钥申请指南

> 详细的API申请步骤，跟着做就行

## 📋 需要申请的API

你需要申请2个必须的API：
1. **短信服务**（用于注册登录验证码）
2. **对象存储**（用于图片上传）

可选：
3. **鲸探API**（用于第三方作品同步，可以后续再配置）

---

## 🎯 方案选择

### 推荐方案：腾讯云（一站式）

**优点**：
- ✅ 一个账号搞定所有服务
- ✅ 界面友好，容易操作
- ✅ 文档完善
- ✅ 价格透明

**缺点**：
- ❌ 需要实名认证

### 备选方案：阿里云

**优点**：
- ✅ 服务稳定
- ✅ 价格便宜一点点

**缺点**：
- ❌ 界面相对复杂

**建议**：如果你已经有腾讯云或阿里云账号，就用现有的。如果都没有，推荐腾讯云。

---

## 📱 一、申请短信服务

### 腾讯云短信（推荐）

#### 1. 注册/登录账号
- 访问：https://cloud.tencent.com
- 注册账号（需要实名认证）
- 登录控制台

#### 2. 开通短信服务
- 搜索"短信"或访问：https://console.cloud.tencent.com/smsv2
- 点击"开通"
- 同意协议

#### 3. 创建应用
- 左侧菜单：应用管理 → 创建应用
- 应用名称：`HOHO小程序`
- 应用简介：`数字作品集换平台`
- 点击"创建应用"
- **记录下 SDK AppID**（类似：1400123456）

#### 4. 配置签名
- 左侧菜单：国内短信 → 签名管理 → 创建签名
- 签名类型：小程序
- 签名用途：自用
- 签名内容：`HOHO作品`（2-8个字）
- 上传小程序截图（微信小程序后台截图）
- 提交审核（通常2小时内通过）
- **记录下签名内容**

#### 5. 配置模板
- 左侧菜单：国内短信 → 正文模板管理 → 创建正文模板
- 模板名称：`登录验证码`
- 短信类型：验证码
- 短信内容：
  ```
  您的验证码是{1}，{2}分钟内有效，请勿泄露。
  ```
- 提交审核（通常2小时内通过）
- **记录下模板ID**（类似：1234567）

#### 6. 获取密钥
- 右上角头像 → 访问管理 → 访问密钥 → API密钥管理
- 点击"新建密钥"
- **记录下 SecretId 和 SecretKey**（只显示一次，务必保存！）

#### 7. 充值
- 左侧菜单：套餐包管理
- 购买国内短信套餐包
- 建议：100元（约2000-3000条）

#### 8. 整理信息
```
【腾讯云短信】
SMS_APP_ID=1400123456
SMS_APP_KEY=你的SecretKey
SMS_SIGN_NAME=HOHO作品
SMS_TEMPLATE_ID=1234567
```

---

### 阿里云短信（备选）

#### 1. 注册/登录
- 访问：https://www.aliyun.com
- 注册并实名认证

#### 2. 开通短信服务
- 搜索"短信服务"或访问：https://dysms.console.aliyun.com
- 开通服务

#### 3. 配置签名
- 左侧菜单：国内消息 → 签名管理 → 添加签名
- 签名名称：`HOHO作品`
- 适用场景：验证码
- 签名来源：小程序
- 上传证明文件
- 提交审核

#### 4. 配置模板
- 左侧菜单：国内消息 → 模板管理 → 添加模板
- 模板名称：`登录验证码`
- 模板内容：
  ```
  您的验证码是${code}，${time}分钟内有效。
  ```
- 提交审核

#### 5. 获取密钥
- 右上角头像 → AccessKey管理
- 创建AccessKey
- **记录 AccessKeyId 和 AccessKeySecret**

#### 6. 充值
- 套餐包管理 → 购买套餐包
- 建议：100元

#### 7. 整理信息
```
【阿里云短信】
SMS_ACCESS_KEY_ID=你的AccessKeyId
SMS_ACCESS_KEY_SECRET=你的AccessKeySecret
SMS_SIGN_NAME=HOHO作品
SMS_TEMPLATE_CODE=SMS_123456789
```

---

## 📦 二、申请对象存储

### 腾讯云COS（推荐）

#### 1. 开通服务
- 搜索"对象存储"或访问：https://console.cloud.tencent.com/cos
- 开通服务

#### 2. 创建存储桶
- 点击"创建存储桶"
- 名称：`hoho-assets-你的数字`（全局唯一，小写字母和数字）
- 所属地域：选择离你服务器近的（如：广州）
- 访问权限：公有读私有写
- 点击"创建"

#### 3. 配置跨域
- 进入存储桶 → 安全管理 → 跨域访问CORS设置
- 添加规则：
  - 来源Origin：`*`
  - 操作Methods：`GET, POST, PUT, DELETE, HEAD`
  - Allow-Headers：`*`
  - Expose-Headers：`ETag`
  - 超时Max-Age：`600`
- 保存

#### 4. 获取密钥
- 使用短信服务的密钥（SecretId 和 SecretKey）

#### 5. 整理信息
```
【腾讯云COS】
COS_SECRET_ID=你的SecretId（和短信服务一样）
COS_SECRET_KEY=你的SecretKey（和短信服务一样）
COS_BUCKET=hoho-assets-1234567890
COS_REGION=ap-guangzhou
```

---

### 阿里云OSS（备选）

#### 1. 开通服务
- 搜索"对象存储OSS"或访问：https://oss.console.aliyun.com
- 开通服务

#### 2. 创建Bucket
- 点击"创建Bucket"
- Bucket名称：`hoho-assets`（全局唯一）
- 地域：选择离服务器近的（如：华南1）
- 读写权限：公共读
- 点击"确定"

#### 3. 配置跨域
- 进入Bucket → 权限管理 → 跨域设置
- 创建规则：
  - 来源：`*`
  - 允许Methods：`GET, POST, PUT, DELETE, HEAD`
  - 允许Headers：`*`
  - 暴露Headers：`ETag`
- 保存

#### 4. 获取密钥
- 使用短信服务的密钥（AccessKeyId 和 AccessKeySecret）

#### 5. 整理信息
```
【阿里云OSS】
OSS_ACCESS_KEY_ID=你的AccessKeyId（和短信服务一样）
OSS_ACCESS_KEY_SECRET=你的AccessKeySecret（和短信服务一样）
OSS_BUCKET=hoho-assets
OSS_ENDPOINT=oss-cn-shenzhen.aliyuncs.com
```

---

## 🐋 三、申请鲸探API（可选）

**说明**：鲸探API需要和鲸探官方申请，流程较复杂。

**建议**：先跳过，等小程序上线后再申请。

如果你已经有鲸探API：
```
【鲸探API】
JINGTAN_API_KEY=你的API密钥
JINGTAN_API_SECRET=你的API密钥
```

---

## ✅ 最终整理

### 如果你选择腾讯云：

```
【服务器信息】
服务器IP：你的服务器IP
SSH端口：22
SSH密码：你的root密码

【域名信息】
API域名：api.你的域名.com
小程序AppID：wx...（微信小程序后台获取）

【腾讯云短信】
SMS_APP_ID=1400123456
SMS_APP_KEY=你的SecretKey
SMS_SIGN_NAME=HOHO作品
SMS_TEMPLATE_ID=1234567

【腾讯云COS】
COS_SECRET_ID=你的SecretId（和短信一样）
COS_SECRET_KEY=你的SecretKey（和短信一样）
COS_BUCKET=hoho-assets-1234567890
COS_REGION=ap-guangzhou

【鲸探API】（可选）
JINGTAN_API_KEY=
JINGTAN_API_SECRET=
```

### 如果你选择阿里云：

```
【服务器信息】
服务器IP：你的服务器IP
SSH端口：22
SSH密码：你的root密码

【域名信息】
API域名：api.你的域名.com
小程序AppID：wx...

【阿里云短信】
SMS_ACCESS_KEY_ID=你的AccessKeyId
SMS_ACCESS_KEY_SECRET=你的AccessKeySecret
SMS_SIGN_NAME=HOHO作品
SMS_TEMPLATE_CODE=SMS_123456789

【阿里云OSS】
OSS_ACCESS_KEY_ID=你的AccessKeyId（和短信一样）
OSS_ACCESS_KEY_SECRET=你的AccessKeySecret（和短信一样）
OSS_BUCKET=hoho-assets
OSS_ENDPOINT=oss-cn-shenzhen.aliyuncs.com

【鲸探API】（可选）
JINGTAN_API_KEY=
JINGTAN_API_SECRET=
```

---

## 💰 成本说明

### 腾讯云
- **短信**：100元充值，约0.045元/条，可发2200条
- **COS**：首月免费50GB存储+10GB流量，之后约0.12元/GB/月

### 阿里云
- **短信**：100元充值，约0.04元/条，可发2500条
- **OSS**：首月免费40GB存储+10GB流量，之后约0.12元/GB/月

**总计**：约100-150元可以用很久

---

## 📞 常见问题

**Q: 签名和模板审核要多久？**  
A: 通常2小时内，最多1个工作日。

**Q: 密钥只显示一次，忘记保存怎么办？**  
A: 删除旧密钥，重新创建新密钥。

**Q: 腾讯云和阿里云可以混用吗？**  
A: 可以，比如短信用腾讯云，存储用阿里云。但不推荐，管理麻烦。

**Q: 鲸探API一定要申请吗？**  
A: 不一定，可以先跳过，不影响其他功能。

**Q: 我不会操作怎么办？**  
A: 把账号密码发给我，我帮你配置（不推荐，安全考虑）。或者按照步骤截图给我，我指导你。

---

## 🎯 下一步

**申请好后**：
1. 把信息整理成上面的格式
2. 发给我
3. 我会配置环境变量
4. 然后开始部署后端

---

**加油！这是最关键的一步！** 💪
