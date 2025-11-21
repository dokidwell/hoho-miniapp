#!/bin/bash

# HOHO小程序单元测试运行脚本

set -e

echo "=========================================="
echo "HOHO小程序单元测试"
echo "=========================================="
echo ""

# 颜色输出
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m' # No Color

# 设置Go代理
export GOPROXY=https://goproxy.cn,direct

# 安装测试依赖
echo -e "${YELLOW}1. 安装测试依赖...${NC}"
go get github.com/stretchr/testify/assert
go get github.com/stretchr/testify/mock
echo -e "${GREEN}✓ 依赖安装完成${NC}"
echo ""

# 运行所有测试
echo -e "${YELLOW}2. 运行单元测试...${NC}"
echo ""

# 运行测试并显示详细输出
go test -v ./services/... 2>&1 | while IFS= read -r line; do
    if [[ $line == *"PASS"* ]]; then
        echo -e "${GREEN}$line${NC}"
    elif [[ $line == *"FAIL"* ]]; then
        echo -e "${RED}$line${NC}"
    elif [[ $line == *"RUN"* ]]; then
        echo -e "${YELLOW}$line${NC}"
    else
        echo "$line"
    fi
done

echo ""

# 运行测试覆盖率
echo -e "${YELLOW}3. 生成测试覆盖率报告...${NC}"
go test -coverprofile=coverage.out ./services/...
go tool cover -html=coverage.out -o coverage.html

COVERAGE=$(go tool cover -func=coverage.out | grep total | awk '{print $3}')
echo -e "${GREEN}✓ 总体覆盖率: $COVERAGE${NC}"
echo "  详细报告: coverage.html"
echo ""

# 运行基准测试
echo -e "${YELLOW}4. 运行基准测试...${NC}"
go test -bench=. -benchmem ./services/... | while IFS= read -r line; do
    if [[ $line == *"Benchmark"* ]]; then
        echo -e "${GREEN}$line${NC}"
    else
        echo "$line"
    fi
done
echo ""

# 检查代码质量
echo -e "${YELLOW}5. 检查代码质量...${NC}"

# 安装golint（如果未安装）
if ! command -v golint &> /dev/null; then
    echo "安装golint..."
    go install golang.org/x/lint/golint@latest
fi

# 运行golint
echo "运行golint..."
golint ./... | while IFS= read -r line; do
    if [ -n "$line" ]; then
        echo -e "${YELLOW}  $line${NC}"
    fi
done

# 运行go vet
echo "运行go vet..."
go vet ./... 2>&1 | while IFS= read -r line; do
    if [ -n "$line" ]; then
        echo -e "${YELLOW}  $line${NC}"
    fi
done

echo -e "${GREEN}✓ 代码质量检查完成${NC}"
echo ""

# 总结
echo "=========================================="
echo -e "${GREEN}测试完成！${NC}"
echo "=========================================="
echo ""
echo "查看详细报告："
echo "  - 覆盖率报告: coverage.html"
echo "  - 覆盖率数据: coverage.out"
echo ""
echo "运行特定测试："
echo "  go test -v ./services/user_test.go"
echo "  go test -v ./services/point_test.go"
echo "  go test -v ./services/asset_test.go"
echo ""
echo "运行特定测试用例："
echo "  go test -v -run TestRegisterUser ./services/"
echo ""
