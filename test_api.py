#!/usr/bin/env python3
"""
HOHO小程序API接口测试脚本
测试所有API端点的可用性和功能
"""

import requests
import json
from typing import Dict, Any, Optional
import time

# API基础URL
BASE_URL = "https://api.hohopark.com/api/v1"

# 测试结果统计
test_results = {
    'total': 0,
    'passed': 0,
    'failed': 0,
    'skipped': 0
}

class APITester:
    def __init__(self, base_url: str):
        self.base_url = base_url
        self.token = None
        self.admin_token = None
    
    def test_endpoint(self, name: str, method: str, endpoint: str, 
                     data: Optional[Dict] = None, 
                     headers: Optional[Dict] = None,
                     expected_status: int = 200,
                     auth_required: bool = False) -> bool:
        """测试单个API端点"""
        url = f"{self.base_url}{endpoint}"
        
        print(f"\n{'='*60}")
        print(f"测试: {name}")
        print(f"方法: {method}")
        print(f"URL: {url}")
        
        test_results['total'] += 1
        
        # 准备请求头
        req_headers = headers or {}
        if auth_required and self.token:
            req_headers['Authorization'] = f'Bearer {self.token}'
        
        try:
            if method == 'GET':
                response = requests.get(url, headers=req_headers, timeout=10)
            elif method == 'POST':
                response = requests.post(url, json=data, headers=req_headers, timeout=10)
            elif method == 'PUT':
                response = requests.put(url, json=data, headers=req_headers, timeout=10)
            elif method == 'DELETE':
                response = requests.delete(url, headers=req_headers, timeout=10)
            else:
                print(f"❌ 不支持的HTTP方法: {method}")
                test_results['failed'] += 1
                return False
            
            print(f"状态码: {response.status_code}")
            
            # 尝试解析JSON响应
            try:
                resp_data = response.json()
                print(f"响应: {json.dumps(resp_data, ensure_ascii=False, indent=2)[:200]}...")
            except:
                print(f"响应: {response.text[:200]}...")
            
            # 检查状态码
            if response.status_code == expected_status:
                print(f"✅ 测试通过")
                test_results['passed'] += 1
                return True
            else:
                print(f"⚠️  状态码不匹配 (期望: {expected_status}, 实际: {response.status_code})")
                test_results['failed'] += 1
                return False
                
        except requests.exceptions.Timeout:
            print(f"❌ 请求超时")
            test_results['failed'] += 1
            return False
        except requests.exceptions.ConnectionError:
            print(f"❌ 连接失败")
            test_results['failed'] += 1
            return False
        except Exception as e:
            print(f"❌ 错误: {str(e)}")
            test_results['failed'] += 1
            return False
    
    def test_health(self):
        """测试健康检查"""
        return self.test_endpoint(
            "健康检查",
            "GET",
            "/health",
            expected_status=200
        )
    
    def test_user_register(self):
        """测试用户注册"""
        # 注意：这个测试会失败，因为需要真实的短信验证码
        return self.test_endpoint(
            "用户注册",
            "POST",
            "/users/register",
            data={
                "phone": "13800138000",
                "password": "Test123456",
                "code": "123456"  # 假的验证码
            },
            expected_status=400  # 预期失败
        )
    
    def test_user_login(self):
        """测试用户登录"""
        # 注意：这个测试会失败，因为用户不存在
        return self.test_endpoint(
            "用户登录",
            "POST",
            "/users/login",
            data={
                "phone": "13800138000",
                "password": "Test123456"
            },
            expected_status=401  # 预期失败（用户不存在）
        )
    
    def test_assets_list(self):
        """测试藏品列表"""
        return self.test_endpoint(
            "藏品列表",
            "GET",
            "/assets",
            expected_status=200
        )
    
    def test_listings_list(self):
        """测试交易挂单列表"""
        return self.test_endpoint(
            "交易挂单列表",
            "GET",
            "/listings",
            expected_status=200
        )
    
    def test_events_list(self):
        """测试社区事件列表"""
        return self.test_endpoint(
            "社区事件列表",
            "GET",
            "/events",
            expected_status=200
        )
    
    def test_admin_login(self):
        """测试管理员登录"""
        return self.test_endpoint(
            "管理员登录",
            "POST",
            "/admin/login",
            data={
                "username": "admin",
                "password": "Admin@123456"
            },
            expected_status=401  # 预期失败（管理员不存在）
        )

def test_api_structure():
    """测试API路由结构"""
    print("\n" + "="*60)
    print("检查API路由定义")
    print("="*60)
    
    main_file = '/home/ubuntu/hoho-miniapp/backend/main.go'
    
    try:
        with open(main_file, 'r') as f:
            content = f.read()
        
        # 查找路由定义
        routes = []
        
        # 用户路由
        if 'users := v1.Group("/users")' in content:
            print("✓ 用户路由组定义")
            routes.append("users")
        
        # 资产路由
        if 'assets := v1.Group("/assets")' in content:
            print("✓ 资产路由组定义")
            routes.append("assets")
        
        # 交易路由
        if 'listings := v1.Group("/listings")' in content:
            print("✓ 交易路由组定义")
            routes.append("listings")
        
        # 事件路由
        if 'events := v1.Group("/events")' in content:
            print("✓ 事件路由组定义")
            routes.append("events")
        
        # 管理员路由
        if 'admin := v1.Group("/admin")' in content:
            print("✓ 管理员路由组定义")
            routes.append("admin")
        
        print(f"\n✅ 共定义了 {len(routes)} 个路由组")
        return True
        
    except Exception as e:
        print(f"❌ 错误: {str(e)}")
        return False

def main():
    print("\n🔍 开始API接口测试...\n")
    print("⚠️  注意: 某些测试预期会失败（如需要验证码、认证等）")
    
    # 检查API路由结构
    test_api_structure()
    
    # 创建测试器
    tester = APITester(BASE_URL)
    
    # 测试公开端点
    print("\n" + "="*60)
    print("测试公开API端点")
    print("="*60)
    
    # 健康检查（应该成功）
    tester.test_health()
    
    # 藏品列表（应该成功）
    tester.test_assets_list()
    
    # 交易列表（应该成功）
    tester.test_listings_list()
    
    # 事件列表（应该成功）
    tester.test_events_list()
    
    # 测试认证端点
    print("\n" + "="*60)
    print("测试认证API端点")
    print("="*60)
    
    # 用户注册（预期失败 - 需要验证码）
    tester.test_user_register()
    
    # 用户登录（预期失败 - 用户不存在）
    tester.test_user_login()
    
    # 管理员登录（预期失败 - 管理员不存在）
    tester.test_admin_login()
    
    # 打印测试总结
    print("\n" + "="*60)
    print("测试总结")
    print("="*60)
    print(f"总测试数: {test_results['total']}")
    print(f"✅ 通过: {test_results['passed']}")
    print(f"❌ 失败: {test_results['failed']}")
    print(f"⏭️  跳过: {test_results['skipped']}")
    
    success_rate = (test_results['passed'] / test_results['total'] * 100) if test_results['total'] > 0 else 0
    print(f"\n成功率: {success_rate:.1f}%")
    
    if test_results['failed'] == 0:
        print("\n✅ 所有测试通过！")
        return 0
    else:
        print(f"\n⚠️  有 {test_results['failed']} 个测试失败")
        print("注意: 某些失败是预期的（如需要验证码、用户不存在等）")
        return 1

if __name__ == '__main__':
    exit(main())
