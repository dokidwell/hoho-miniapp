# 🎨 HOHO Park - 数字藏品小程序

<div align="center">

![HOHO Park](https://img.shields.io/badge/HOHO-Park-667eea)
![Go Version](https://img.shields.io/badge/Go-1.21-00ADD8)
![Vue Version](https://img.shields.io/badge/Vue-3.3-4FC08D)
![License](https://img.shields.io/badge/license-MIT-green)

**一个完整的数字藏品小程序平台**

[功能特性](#-功能特性) • [快速开始](#-快速开始) • [文档](#-文档) • [部署](#-部署)

</div>

---

## 📖 项目简介

HOHO Park是一个功能完整的数字藏品（NFT）小程序平台，支持藏品铸造、交易、空投、社区活动等核心功能。采用前后端分离架构，技术栈成熟稳定。

### 技术栈

**后端**：Go + Gin + GORM + MySQL + Redis  
**前端**：uni-app (Vue 3) + 微信小程序  
**部署**：Ubuntu + Nginx + SSL

---

## ✨ 功能特性

### 核心功能

- 🎯 **用户系统** - 注册、登录、实名认证
- 💰 **积分系统** - 高精度积分（8位小数）
- 🖼️ **藏品管理** - 铸造、转让、销毁
- 🛒 **集换市场** - 挂单、交易、手续费
- 🎁 **空投活动** - 社区奖励、活动参与
- 🌐 **第三方对接** - 鲸探API集成
- 👨‍💼 **管理后台** - 用户、藏品、统计

### 技术亮点

- ✅ 高精度decimal积分系统
- ✅ 完善的事务处理
- ✅ 30+单元测试
- ✅ 统一错误处理
- ✅ 性能优化（缓存、索引）
- ✅ 安全加固（HTTPS、JWT）

---

## 🚀 快速开始

### 环境要求

- Go 1.21+
- MySQL 8.0+
- Redis 6.0+
- Node.js 16+
- HBuilderX（前端编译）

### 后端启动

```bash
# 克隆项目
git clone https://github.com/dokidwell/hoho-miniapp.git
cd hoho-miniapp/backend

# 安装依赖
export GOPROXY=https://goproxy.cn,direct
go mod tidy

# 配置环境变量
cp .env.example .env
# 编辑.env文件，填入数据库等配置

# 初始化数据库
./setup_database.sh

# 运行
go run main.go
```

### 前端启动

```bash
# 进入前端目录
cd frontend

# 使用HBuilderX打开项目
# 运行 → 运行到小程序模拟器 → 微信开发者工具
```

---

## 📚 文档

| 文档 | 说明 |
|-----|------|
| [最终交付报告](./最终交付报告.md) | 完整的项目交付报告 |
| [部署完成报告](./部署完成报告.md) | 服务器部署详情 |
| [快速参考指南](./快速参考指南.md) | 常用命令和故障排查 |
| [用户体验分析](./用户体验全面分析报告.md) | 用户旅程和优化建议 |
| [数据库初始化指南](./数据库初始化操作指南.md) | 数据库初始化步骤 |
| [前端编译说明](./前端编译说明.md) | 前端编译和发布 |
| [测试报告](./测试报告.md) | 测试结果和问题 |

---

## 🌐 部署

### 生产环境

**API地址**：https://api.hohopark.com  
**服务器**：腾讯云 Ubuntu 22.04  
**SSL证书**：Let's Encrypt（自动续期）

### 部署步骤

详见 [部署完成报告](./部署完成报告.md)

---

## 📊 项目统计

- **代码行数**：~10,000行
- **后端文件**：25个Go文件
- **前端页面**：30个Vue页面
- **数据表**：12个MySQL表
- **API端点**：50+个RESTful API
- **单元测试**：30+个测试用例
- **完成度**：98%

---

## 🧪 测试

```bash
# 运行后端单元测试
cd backend
./run_tests.sh

# 运行用户旅程测试
cd ..
python3 user_journey_test.py
```

---

## 📝 许可证

MIT License

---

## 🤝 贡献

欢迎提交Issue和Pull Request！

---

## 📞 联系方式

- GitHub：https://github.com/dokidwell/hoho-miniapp
- 问题反馈：https://github.com/dokidwell/hoho-miniapp/issues

---

<div align="center">

**Made with ❤️ by HOHO Team**

</div>
