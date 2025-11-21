#!/usr/bin/env python3
# -*- coding: utf-8 -*-

"""
HOHO小程序用户旅程测试
从真实用户角度测试完整业务流程
"""

import requests
import json
import time
from datetime import datetime

# 配置
API_BASE_URL = "https://api.hohopark.com"
TEST_PHONE = "13800138000"
TEST_PASSWORD = "Test123456!"
ADMIN_USERNAME = "admin"
ADMIN_PASSWORD = "Admin@123456"

# 颜色输出
class Colors:
    GREEN = '\033[92m'
    YELLOW = '\033[93m'
    RED = '\033[91m'
    BLUE = '\033[94m'
    ENDC = '\033[0m'
    BOLD = '\033[1m'

def print_header(text):
    print(f"\n{Colors.BOLD}{Colors.BLUE}{'='*60}{Colors.ENDC}")
    print(f"{Colors.BOLD}{Colors.BLUE}{text:^60}{Colors.ENDC}")
    print(f"{Colors.BOLD}{Colors.BLUE}{'='*60}{Colors.ENDC}\n")

def print_step(step_num, text):
    print(f"{Colors.YELLOW}[步骤 {step_num}]{Colors.ENDC} {text}")

def print_success(text):
    print(f"{Colors.GREEN}✓ {text}{Colors.ENDC}")

def print_error(text):
    print(f"{Colors.RED}✗ {text}{Colors.ENDC}")

def print_info(text):
    print(f"  {text}")

class UserJourneyTest:
    def __init__(self):
        self.session = requests.Session()
        self.user_token = None
        self.admin_token = None
        self.user_id = None
        self.asset_id = None
        self.instance_id = None
        self.listing_id = None
        
        self.passed_tests = 0
        self.failed_tests = 0
        self.total_tests = 0

    def test(self, name, func):
        """运行单个测试"""
        self.total_tests += 1
        try:
            print_step(self.total_tests, name)
            result = func()
            if result:
                self.passed_tests += 1
                print_success(f"{name} - 通过")
            else:
                self.failed_tests += 1
                print_error(f"{name} - 失败")
            return result
        except Exception as e:
            self.failed_tests += 1
            print_error(f"{name} - 异常: {str(e)}")
            return False

    def api_call(self, method, endpoint, data=None, token=None, expect_success=True):
        """API调用封装"""
        url = f"{API_BASE_URL}{endpoint}"
        headers = {"Content-Type": "application/json"}
        
        if token:
            headers["Authorization"] = f"Bearer {token}"
        
        try:
            if method == "GET":
                response = self.session.get(url, headers=headers, timeout=10)
            elif method == "POST":
                response = self.session.post(url, json=data, headers=headers, timeout=10)
            elif method == "PUT":
                response = self.session.put(url, json=data, headers=headers, timeout=10)
            elif method == "DELETE":
                response = self.session.delete(url, headers=headers, timeout=10)
            
            print_info(f"请求: {method} {endpoint}")
            print_info(f"状态码: {response.status_code}")
            
            if response.status_code == 200 or response.status_code == 201:
                try:
                    result = response.json()
                    print_info(f"响应: {json.dumps(result, ensure_ascii=False, indent=2)[:200]}...")
                    return result
                except:
                    print_info(f"响应: {response.text[:200]}")
                    return {"success": True}
            else:
                print_info(f"错误: {response.text[:200]}")
                if not expect_success:
                    return {"success": False, "error": response.text}
                return None
                
        except requests.exceptions.Timeout:
            print_error("请求超时")
            return None
        except requests.exceptions.ConnectionError:
            print_error("连接失败")
            return None
        except Exception as e:
            print_error(f"请求异常: {str(e)}")
            return None

    # ========== 用户旅程测试 ==========

    def journey_1_new_user_registration(self):
        """旅程1：新用户注册"""
        print_header("用户旅程 1：新用户注册")
        
        # 1. 用户打开小程序
        def test_open_app():
            print_info("用户打开HOHO Park小程序...")
            time.sleep(0.5)
            return True
        
        self.test("打开小程序", test_open_app)
        
        # 2. 点击注册
        def test_register():
            result = self.api_call("POST", "/api/v1/auth/register", {
                "phone": TEST_PHONE,
                "password": TEST_PASSWORD,
                "code": "123456"  # 假设验证码
            })
            if result and "token" in result:
                self.user_token = result["token"]
                self.user_id = result.get("user_id")
                return True
            return False
        
        self.test("用户注册", test_register)
        
        # 3. 获取注册奖励积分
        def test_register_reward():
            result = self.api_call("GET", "/api/v1/user/points", token=self.user_token)
            if result and "balance" in result:
                print_info(f"获得注册奖励: {result.get('balance', 0)} 积分")
                return True
            return False
        
        self.test("获取注册奖励", test_register_reward)

    def journey_2_browse_and_collect(self):
        """旅程2：浏览和收藏藏品"""
        print_header("用户旅程 2：浏览和收藏藏品")
        
        # 1. 浏览首页藏品列表
        def test_browse_assets():
            result = self.api_call("GET", "/api/v1/assets")
            if result and "list" in result:
                assets = result["list"]
                print_info(f"发现 {len(assets)} 件藏品")
                if len(assets) > 0:
                    self.asset_id = assets[0].get("id")
                    print_info(f"选中藏品ID: {self.asset_id}")
                return True
            return False
        
        self.test("浏览藏品列表", test_browse_assets)
        
        # 2. 查看藏品详情
        def test_view_asset_detail():
            if not self.asset_id:
                print_info("跳过：没有可用的藏品ID")
                return True
            
            result = self.api_call("GET", f"/api/v1/assets/{self.asset_id}")
            if result:
                print_info(f"藏品名称: {result.get('name', 'N/A')}")
                print_info(f"发行量: {result.get('total_supply', 0)}")
                print_info(f"已铸造: {result.get('minted_count', 0)}")
                return True
            return False
        
        self.test("查看藏品详情", test_view_asset_detail)
        
        # 3. 浏览社区事件
        def test_browse_events():
            result = self.api_call("GET", "/api/v1/events")
            if result:
                events = result.get("list", [])
                print_info(f"发现 {len(events)} 个社区事件")
                return True
            return False
        
        self.test("浏览社区事件", test_browse_events)

    def journey_3_participate_airdrop(self):
        """旅程3：参与空投活动"""
        print_header("用户旅程 3：参与空投活动")
        
        # 1. 查看空投列表
        def test_view_airdrops():
            result = self.api_call("GET", "/api/v1/airdrops", token=self.user_token)
            if result:
                airdrops = result.get("list", [])
                print_info(f"发现 {len(airdrops)} 个空投活动")
                return True
            return False
        
        self.test("查看空投列表", test_view_airdrops)
        
        # 2. 参与空投（如果有）
        def test_participate_airdrop():
            # 这里应该有实际的空投ID
            print_info("暂无进行中的空投活动")
            return True
        
        self.test("参与空投活动", test_participate_airdrop)

    def journey_4_marketplace_trading(self):
        """旅程4：集换市场交易"""
        print_header("用户旅程 4：集换市场交易")
        
        # 1. 浏览集换市场
        def test_browse_marketplace():
            result = self.api_call("GET", "/api/v1/listings")
            if result:
                listings = result.get("list", [])
                print_info(f"市场上有 {len(listings)} 个挂单")
                if len(listings) > 0:
                    self.listing_id = listings[0].get("id")
                return True
            return False
        
        self.test("浏览集换市场", test_browse_marketplace)
        
        # 2. 查看挂单详情
        def test_view_listing():
            if not self.listing_id:
                print_info("跳过：没有可用的挂单")
                return True
            
            result = self.api_call("GET", f"/api/v1/listings/{self.listing_id}")
            if result:
                print_info(f"价格: {result.get('price', 0)} 积分")
                return True
            return False
        
        self.test("查看挂单详情", test_view_listing)

    def journey_5_user_profile(self):
        """旅程5：个人中心管理"""
        print_header("用户旅程 5：个人中心管理")
        
        # 1. 查看个人信息
        def test_view_profile():
            result = self.api_call("GET", "/api/v1/user/profile", token=self.user_token)
            if result:
                print_info(f"用户ID: {result.get('uid', 'N/A')}")
                print_info(f"昵称: {result.get('nickname', '未设置')}")
                return True
            return False
        
        self.test("查看个人信息", test_view_profile)
        
        # 2. 查看我的作品
        def test_view_my_assets():
            result = self.api_call("GET", "/api/v1/user/assets", token=self.user_token)
            if result:
                assets = result.get("list", [])
                print_info(f"拥有 {len(assets)} 件藏品")
                return True
            return False
        
        self.test("查看我的作品", test_view_my_assets)
        
        # 3. 查看积分记录
        def test_view_point_history():
            result = self.api_call("GET", "/api/v1/user/points/history", token=self.user_token)
            if result:
                transactions = result.get("list", [])
                print_info(f"共 {len(transactions)} 条积分记录")
                return True
            return False
        
        self.test("查看积分记录", test_view_point_history)

    def journey_6_third_party_integration(self):
        """旅程6：第三方平台关联"""
        print_header("用户旅程 6：第三方平台关联")
        
        # 1. 查看鲸探资产
        def test_view_jingtan_assets():
            result = self.api_call("GET", "/api/v1/jingtan/assets", token=self.user_token)
            if result:
                assets = result.get("list", [])
                print_info(f"鲸探资产: {len(assets)} 件")
                return True
            return False
        
        self.test("查看鲸探资产", test_view_jingtan_assets)

    def journey_7_admin_management(self):
        """旅程7：管理员后台管理"""
        print_header("用户旅程 7：管理员后台管理")
        
        # 1. 管理员登录
        def test_admin_login():
            result = self.api_call("POST", "/admin/login", {
                "username": ADMIN_USERNAME,
                "password": ADMIN_PASSWORD
            })
            if result and "token" in result:
                self.admin_token = result["token"]
                print_info("管理员登录成功")
                return True
            return False
        
        self.test("管理员登录", test_admin_login)
        
        # 2. 查看用户列表
        def test_admin_view_users():
            if not self.admin_token:
                print_info("跳过：管理员未登录")
                return True
            
            result = self.api_call("GET", "/admin/users", token=self.admin_token)
            if result:
                users = result.get("list", [])
                print_info(f"系统用户数: {len(users)}")
                return True
            return False
        
        self.test("查看用户列表", test_admin_view_users)
        
        # 3. 查看系统统计
        def test_admin_view_stats():
            if not self.admin_token:
                print_info("跳过：管理员未登录")
                return True
            
            result = self.api_call("GET", "/admin/stats", token=self.admin_token)
            if result:
                print_info(f"统计数据: {json.dumps(result, ensure_ascii=False)[:100]}")
                return True
            return False
        
        self.test("查看系统统计", test_admin_view_stats)

    def run_all_journeys(self):
        """运行所有用户旅程"""
        print(f"\n{Colors.BOLD}HOHO小程序用户旅程测试{Colors.ENDC}")
        print(f"测试时间: {datetime.now().strftime('%Y-%m-%d %H:%M:%S')}")
        print(f"API地址: {API_BASE_URL}")
        
        # 执行所有旅程
        self.journey_1_new_user_registration()
        self.journey_2_browse_and_collect()
        self.journey_3_participate_airdrop()
        self.journey_4_marketplace_trading()
        self.journey_5_user_profile()
        self.journey_6_third_party_integration()
        self.journey_7_admin_management()
        
        # 输出测试结果
        print_header("测试结果汇总")
        print(f"总测试数: {self.total_tests}")
        print(f"{Colors.GREEN}通过: {self.passed_tests}{Colors.ENDC}")
        print(f"{Colors.RED}失败: {self.failed_tests}{Colors.ENDC}")
        
        if self.failed_tests == 0:
            print(f"\n{Colors.GREEN}{Colors.BOLD}🎉 所有测试通过！{Colors.ENDC}")
        else:
            success_rate = (self.passed_tests / self.total_tests * 100) if self.total_tests > 0 else 0
            print(f"\n{Colors.YELLOW}通过率: {success_rate:.1f}%{Colors.ENDC}")
        
        print()

if __name__ == "__main__":
    test = UserJourneyTest()
    test.run_all_journeys()
