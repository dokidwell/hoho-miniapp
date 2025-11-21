#!/bin/bash

# HOHO小程序数据库初始化脚本
# 用途：初始化数据库表结构并创建管理员账户

set -e  # 遇到错误立即退出

echo "=========================================="
echo "HOHO小程序数据库初始化"
echo "=========================================="
echo ""

# 数据库配置
DB_NAME="hoho_miniapp"
DB_USER="hoho"
DB_PASSWORD="HoHo@2024!Secure"

# 管理员配置
ADMIN_USERNAME="admin"
ADMIN_PASSWORD="Admin@123456"
ADMIN_EMAIL="admin@hohopark.com"

# 颜色输出
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m' # No Color

# 检查MySQL是否运行
echo -e "${YELLOW}1. 检查MySQL服务状态...${NC}"
if sudo systemctl is-active --quiet mysql; then
    echo -e "${GREEN}✓ MySQL服务正在运行${NC}"
else
    echo -e "${RED}✗ MySQL服务未运行${NC}"
    echo "正在启动MySQL服务..."
    sudo systemctl start mysql
    sleep 3
    if sudo systemctl is-active --quiet mysql; then
        echo -e "${GREEN}✓ MySQL服务已启动${NC}"
    else
        echo -e "${RED}✗ MySQL服务启动失败${NC}"
        exit 1
    fi
fi
echo ""

# 检查数据库是否存在
echo -e "${YELLOW}2. 检查数据库...${NC}"
if sudo mysql -e "USE ${DB_NAME};" 2>/dev/null; then
    echo -e "${GREEN}✓ 数据库 ${DB_NAME} 已存在${NC}"
else
    echo -e "${RED}✗ 数据库 ${DB_NAME} 不存在${NC}"
    echo "正在创建数据库..."
    sudo mysql -e "CREATE DATABASE ${DB_NAME} CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;"
    echo -e "${GREEN}✓ 数据库创建成功${NC}"
fi
echo ""

# 检查表是否已存在
echo -e "${YELLOW}3. 检查数据库表...${NC}"
TABLE_COUNT=$(sudo mysql ${DB_NAME} -sN -e "SELECT COUNT(*) FROM information_schema.tables WHERE table_schema = '${DB_NAME}';")
if [ "$TABLE_COUNT" -gt 0 ]; then
    echo -e "${YELLOW}⚠ 发现 ${TABLE_COUNT} 个已存在的表${NC}"
    echo "是否要删除所有表并重新创建？(y/N)"
    read -r response
    if [[ "$response" =~ ^([yY][eE][sS]|[yY])$ ]]; then
        echo "正在删除所有表..."
        sudo mysql ${DB_NAME} -e "SET FOREIGN_KEY_CHECKS = 0;"
        sudo mysql ${DB_NAME} -e "
            SELECT CONCAT('DROP TABLE IF EXISTS \`', table_name, '\`;')
            FROM information_schema.tables
            WHERE table_schema = '${DB_NAME}';
        " | grep DROP | sudo mysql ${DB_NAME}
        sudo mysql ${DB_NAME} -e "SET FOREIGN_KEY_CHECKS = 1;"
        echo -e "${GREEN}✓ 所有表已删除${NC}"
    else
        echo "跳过表删除"
    fi
else
    echo -e "${GREEN}✓ 数据库为空，准备创建表${NC}"
fi
echo ""

# 执行SQL初始化脚本
echo -e "${YELLOW}4. 执行SQL初始化脚本...${NC}"
if [ -f "init.sql" ]; then
    echo "正在执行 init.sql..."
    sudo mysql ${DB_NAME} < init.sql
    echo -e "${GREEN}✓ SQL脚本执行成功${NC}"
else
    echo -e "${RED}✗ 找不到 init.sql 文件${NC}"
    exit 1
fi
echo ""

# 验证表创建
echo -e "${YELLOW}5. 验证表创建...${NC}"
CREATED_TABLES=$(sudo mysql ${DB_NAME} -sN -e "SHOW TABLES;")
TABLE_COUNT=$(echo "$CREATED_TABLES" | wc -l)
echo -e "${GREEN}✓ 成功创建 ${TABLE_COUNT} 个表：${NC}"
echo "$CREATED_TABLES" | while read table; do
    echo "  - $table"
done
echo ""

# 生成管理员密码哈希
echo -e "${YELLOW}6. 创建管理员账户...${NC}"

# 检查是否已安装Python
if command -v python3 &> /dev/null; then
    echo "正在生成密码哈希..."
    
    # 使用Python生成bcrypt哈希
    PASSWORD_HASH=$(python3 << EOF
import bcrypt
password = "${ADMIN_PASSWORD}".encode('utf-8')
hashed = bcrypt.hashpw(password, bcrypt.gensalt())
print(hashed.decode('utf-8'))
EOF
)
    
    if [ -n "$PASSWORD_HASH" ]; then
        echo -e "${GREEN}✓ 密码哈希生成成功${NC}"
        
        # 检查管理员是否已存在
        ADMIN_EXISTS=$(sudo mysql ${DB_NAME} -sN -e "SELECT COUNT(*) FROM admins WHERE username = '${ADMIN_USERNAME}';")
        
        if [ "$ADMIN_EXISTS" -gt 0 ]; then
            echo -e "${YELLOW}⚠ 管理员账户已存在${NC}"
            echo "是否要更新管理员密码？(y/N)"
            read -r response
            if [[ "$response" =~ ^([yY][eE][sS]|[yY])$ ]]; then
                sudo mysql ${DB_NAME} -e "
                    UPDATE admins 
                    SET password_hash = '${PASSWORD_HASH}',
                        updated_at = NOW()
                    WHERE username = '${ADMIN_USERNAME}';
                "
                echo -e "${GREEN}✓ 管理员密码已更新${NC}"
            fi
        else
            # 插入管理员账户
            sudo mysql ${DB_NAME} -e "
                INSERT INTO admins (username, password_hash, email, role, status, created_at, updated_at) 
                VALUES (
                    '${ADMIN_USERNAME}', 
                    '${PASSWORD_HASH}', 
                    '${ADMIN_EMAIL}', 
                    'super_admin', 
                    'active',
                    NOW(),
                    NOW()
                );
            "
            echo -e "${GREEN}✓ 管理员账户创建成功${NC}"
        fi
    else
        echo -e "${RED}✗ 密码哈希生成失败${NC}"
        echo "请手动创建管理员账户"
    fi
else
    echo -e "${YELLOW}⚠ 未安装Python，跳过管理员账户创建${NC}"
    echo "请手动创建管理员账户，或安装Python后重新运行此脚本"
fi
echo ""

# 显示管理员信息
echo -e "${YELLOW}7. 管理员账户信息${NC}"
ADMIN_INFO=$(sudo mysql ${DB_NAME} -e "
    SELECT id, username, email, role, status, created_at 
    FROM admins 
    WHERE username = '${ADMIN_USERNAME}';
" 2>/dev/null)

if [ -n "$ADMIN_INFO" ]; then
    echo "$ADMIN_INFO"
    echo ""
    echo -e "${GREEN}管理员登录信息：${NC}"
    echo "  用户名: ${ADMIN_USERNAME}"
    echo "  密码: ${ADMIN_PASSWORD}"
    echo "  登录地址: https://api.hohopark.com/admin/login"
else
    echo -e "${YELLOW}⚠ 管理员账户未创建${NC}"
fi
echo ""

# 重启后端服务
echo -e "${YELLOW}8. 重启后端服务...${NC}"
if sudo systemctl is-active --quiet hoho-api; then
    sudo systemctl restart hoho-api
    sleep 2
    if sudo systemctl is-active --quiet hoho-api; then
        echo -e "${GREEN}✓ 后端服务已重启${NC}"
    else
        echo -e "${RED}✗ 后端服务重启失败${NC}"
        echo "查看日志："
        sudo journalctl -u hoho-api -n 20 --no-pager
    fi
else
    echo -e "${YELLOW}⚠ 后端服务未运行，正在启动...${NC}"
    sudo systemctl start hoho-api
    sleep 2
    if sudo systemctl is-active --quiet hoho-api; then
        echo -e "${GREEN}✓ 后端服务已启动${NC}"
    else
        echo -e "${RED}✗ 后端服务启动失败${NC}"
    fi
fi
echo ""

# 测试API
echo -e "${YELLOW}9. 测试API连接...${NC}"
sleep 3
API_RESPONSE=$(curl -s -o /dev/null -w "%{http_code}" https://api.hohopark.com/health 2>/dev/null || echo "000")

if [ "$API_RESPONSE" = "200" ]; then
    echo -e "${GREEN}✓ API连接正常${NC}"
    echo "  健康检查: https://api.hohopark.com/health"
elif [ "$API_RESPONSE" = "404" ]; then
    echo -e "${YELLOW}⚠ API返回404，但服务可能正常运行${NC}"
    echo "  尝试访问: https://api.hohopark.com/api/v1/assets"
else
    echo -e "${RED}✗ API连接失败 (HTTP ${API_RESPONSE})${NC}"
    echo "  请检查Nginx和后端服务状态"
fi
echo ""

# 完成
echo "=========================================="
echo -e "${GREEN}数据库初始化完成！${NC}"
echo "=========================================="
echo ""
echo "下一步："
echo "1. 测试管理员登录: https://api.hohopark.com/admin/login"
echo "2. 测试API接口: https://api.hohopark.com/api/v1/assets"
echo "3. 查看服务日志: sudo journalctl -u hoho-api -f"
echo ""
echo "如有问题，请查看："
echo "- 后端日志: sudo journalctl -u hoho-api -n 50"
echo "- Nginx日志: sudo tail -f /var/log/nginx/hoho-api-error.log"
echo "- 数据库状态: sudo mysql ${DB_NAME} -e 'SHOW TABLES;'"
echo ""
