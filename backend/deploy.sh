#!/bin/bash

# HOHO小程序后端一键部署脚本
# 使用方法: bash deploy.sh

set -e

echo "================================"
echo "HOHO后端部署脚本"
echo "================================"
echo ""

# 颜色定义
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# 检查是否为root用户
if [ "$EUID" -ne 0 ]; then 
    echo -e "${RED}请使用root用户运行此脚本${NC}"
    echo "使用命令: sudo bash deploy.sh"
    exit 1
fi

echo -e "${GREEN}步骤1: 检查系统环境${NC}"
echo "当前系统: $(lsb_release -d | cut -f2)"
echo ""

# 安装必要软件
echo -e "${GREEN}步骤2: 安装必要软件${NC}"
echo "更新软件包列表..."
apt update -qq

echo "安装Go语言..."
if ! command -v go &> /dev/null; then
    wget -q https://go.dev/dl/go1.21.0.linux-amd64.tar.gz
    tar -C /usr/local -xzf go1.21.0.linux-amd64.tar.gz
    echo 'export PATH=$PATH:/usr/local/go/bin' >> /etc/profile
    source /etc/profile
    rm go1.21.0.linux-amd64.tar.gz
    echo -e "${GREEN}✓ Go已安装${NC}"
else
    echo -e "${GREEN}✓ Go已存在${NC}"
fi

echo "安装MySQL..."
if ! command -v mysql &> /dev/null; then
    DEBIAN_FRONTEND=noninteractive apt install -y mysql-server
    systemctl start mysql
    systemctl enable mysql
    echo -e "${GREEN}✓ MySQL已安装${NC}"
else
    echo -e "${GREEN}✓ MySQL已存在${NC}"
fi

echo "安装Redis..."
if ! command -v redis-server &> /dev/null; then
    apt install -y redis-server
    systemctl start redis
    systemctl enable redis
    echo -e "${GREEN}✓ Redis已安装${NC}"
else
    echo -e "${GREEN}✓ Redis已存在${NC}"
fi

echo "安装Nginx..."
if ! command -v nginx &> /dev/null; then
    apt install -y nginx
    systemctl start nginx
    systemctl enable nginx
    echo -e "${GREEN}✓ Nginx已安装${NC}"
else
    echo -e "${GREEN}✓ Nginx已存在${NC}"
fi

echo ""

# 配置数据库
echo -e "${GREEN}步骤3: 配置数据库${NC}"
read -p "请输入MySQL root密码（如果是首次安装，直接回车）: " MYSQL_ROOT_PASS

if [ -z "$MYSQL_ROOT_PASS" ]; then
    MYSQL_CMD="mysql"
else
    MYSQL_CMD="mysql -uroot -p${MYSQL_ROOT_PASS}"
fi

read -p "请输入要创建的数据库名称 [hoho_park]: " DB_NAME
DB_NAME=${DB_NAME:-hoho_park}

read -p "请输入数据库用户名 [hoho]: " DB_USER
DB_USER=${DB_USER:-hoho}

read -sp "请输入数据库密码: " DB_PASS
echo ""

echo "创建数据库和用户..."
$MYSQL_CMD <<EOF
CREATE DATABASE IF NOT EXISTS ${DB_NAME} CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
CREATE USER IF NOT EXISTS '${DB_USER}'@'localhost' IDENTIFIED BY '${DB_PASS}';
GRANT ALL PRIVILEGES ON ${DB_NAME}.* TO '${DB_USER}'@'localhost';
FLUSH PRIVILEGES;
EOF

echo -e "${GREEN}✓ 数据库配置完成${NC}"
echo ""

# 配置环境变量
echo -e "${GREEN}步骤4: 配置环境变量${NC}"

if [ ! -f .env ]; then
    cp .env.example .env
fi

read -p "请输入API服务端口 [8080]: " API_PORT
API_PORT=${API_PORT:-8080}

read -p "请输入JWT密钥（随机字符串，至少32位）: " JWT_SECRET
if [ -z "$JWT_SECRET" ]; then
    JWT_SECRET=$(openssl rand -base64 32)
    echo "已自动生成JWT密钥: $JWT_SECRET"
fi

# 写入.env文件
cat > .env <<EOF
# 数据库配置
DB_HOST=localhost
DB_PORT=3306
DB_USER=${DB_USER}
DB_PASSWORD=${DB_PASS}
DB_NAME=${DB_NAME}

# Redis配置
REDIS_HOST=localhost
REDIS_PORT=6379
REDIS_PASSWORD=

# JWT配置
JWT_SECRET=${JWT_SECRET}

# 服务器配置
SERVER_PORT=${API_PORT}
SERVER_MODE=release

# 文件上传配置
UPLOAD_PATH=./uploads
MAX_UPLOAD_SIZE=10485760

# 鲸探API配置（如果有）
JINGTAN_API_URL=
JINGTAN_API_KEY=
EOF

echo -e "${GREEN}✓ 环境变量配置完成${NC}"
echo ""

# 编译项目
echo -e "${GREEN}步骤5: 编译项目${NC}"
echo "下载依赖..."
go mod download

echo "编译..."
go build -o hoho-backend main.go

if [ -f hoho-backend ]; then
    echo -e "${GREEN}✓ 编译成功${NC}"
else
    echo -e "${RED}✗ 编译失败${NC}"
    exit 1
fi
echo ""

# 创建systemd服务
echo -e "${GREEN}步骤6: 创建系统服务${NC}"

CURRENT_DIR=$(pwd)

cat > /etc/systemd/system/hoho-backend.service <<EOF
[Unit]
Description=HOHO Backend Service
After=network.target mysql.service redis.service

[Service]
Type=simple
User=ubuntu
WorkingDirectory=${CURRENT_DIR}
ExecStart=${CURRENT_DIR}/hoho-backend
Restart=always
RestartSec=5
StandardOutput=append:/var/log/hoho-backend.log
StandardError=append:/var/log/hoho-backend-error.log

[Install]
WantedBy=multi-user.target
EOF

systemctl daemon-reload
systemctl enable hoho-backend
echo -e "${GREEN}✓ 系统服务已创建${NC}"
echo ""

# 配置Nginx
echo -e "${GREEN}步骤7: 配置Nginx反向代理${NC}"
read -p "请输入你的API域名（如 api.hohopark.com）: " API_DOMAIN

if [ -z "$API_DOMAIN" ]; then
    echo -e "${YELLOW}未配置域名，跳过Nginx配置${NC}"
else
    cat > /etc/nginx/sites-available/hoho-api <<EOF
server {
    listen 80;
    server_name ${API_DOMAIN};

    client_max_body_size 10M;

    location / {
        proxy_pass http://localhost:${API_PORT};
        proxy_set_header Host \$host;
        proxy_set_header X-Real-IP \$remote_addr;
        proxy_set_header X-Forwarded-For \$proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto \$scheme;
    }
}
EOF

    ln -sf /etc/nginx/sites-available/hoho-api /etc/nginx/sites-enabled/
    nginx -t && systemctl reload nginx
    echo -e "${GREEN}✓ Nginx配置完成${NC}"
    
    # 提示配置HTTPS
    echo ""
    echo -e "${YELLOW}重要提示：微信小程序必须使用HTTPS${NC}"
    echo "请运行以下命令配置SSL证书："
    echo ""
    echo "  apt install -y certbot python3-certbot-nginx"
    echo "  certbot --nginx -d ${API_DOMAIN}"
    echo ""
fi
echo ""

# 启动服务
echo -e "${GREEN}步骤8: 启动服务${NC}"
systemctl start hoho-backend

sleep 2

if systemctl is-active --quiet hoho-backend; then
    echo -e "${GREEN}✓ 服务启动成功${NC}"
else
    echo -e "${RED}✗ 服务启动失败${NC}"
    echo "查看日志: journalctl -u hoho-backend -f"
    exit 1
fi

echo ""
echo "================================"
echo -e "${GREEN}部署完成！${NC}"
echo "================================"
echo ""
echo "服务状态: systemctl status hoho-backend"
echo "查看日志: tail -f /var/log/hoho-backend.log"
echo "停止服务: systemctl stop hoho-backend"
echo "重启服务: systemctl restart hoho-backend"
echo ""

if [ -n "$API_DOMAIN" ]; then
    echo "API地址: http://${API_DOMAIN}"
    echo ""
    echo -e "${YELLOW}下一步：配置HTTPS证书${NC}"
    echo "运行命令: certbot --nginx -d ${API_DOMAIN}"
else
    echo "API地址: http://你的服务器IP:${API_PORT}"
fi

echo ""
echo -e "${GREEN}数据库信息：${NC}"
echo "数据库名: ${DB_NAME}"
echo "用户名: ${DB_USER}"
echo "密码: ${DB_PASS}"
echo ""
echo "请妥善保管以上信息！"
