#!/usr/bin/env python3
"""
HOHO小程序前端代码检查脚本
检查前端页面、配置和API调用
"""

import json
from pathlib import Path
import re

def check_pages_config():
    """检查pages.json配置"""
    print("=" * 60)
    print("检查pages.json配置")
    print("=" * 60)
    
    pages_json = Path('/home/ubuntu/hoho-miniapp/frontend/pages.json')
    
    if not pages_json.exists():
        print("❌ pages.json不存在")
        return False
    
    with open(pages_json, 'r', encoding='utf-8') as f:
        config = json.load(f)
    
    # 检查页面配置
    pages = config.get('pages', [])
    print(f"✓ 配置了 {len(pages)} 个页面")
    
    # 检查TabBar配置
    tabbar = config.get('tabBar', {})
    if tabbar:
        print(f"✓ TabBar配置存在")
        tabbar_pages = tabbar.get('list', [])
        print(f"  - TabBar页面数: {len(tabbar_pages)}")
        print(f"  - 使用自定义TabBar: {tabbar.get('custom', False)}")
    else:
        print("⚠️  未配置TabBar")
    
    # 检查全局样式
    global_style = config.get('globalStyle', {})
    if global_style:
        print(f"✓ 全局样式配置存在")
        print(f"  - 导航栏标题: {global_style.get('navigationBarTitleText', 'N/A')}")
    
    # 检查权限配置
    permissions = config.get('permission', {})
    if permissions:
        print(f"✓ 权限配置存在 ({len(permissions)} 项)")
    
    return True

def check_pages_files():
    """检查页面文件是否存在"""
    print("\n" + "=" * 60)
    print("检查页面文件")
    print("=" * 60)
    
    pages_json = Path('/home/ubuntu/hoho-miniapp/frontend/pages.json')
    
    with open(pages_json, 'r', encoding='utf-8') as f:
        config = json.load(f)
    
    pages = config.get('pages', [])
    missing_pages = []
    
    for page in pages:
        page_path = page.get('path', '')
        vue_file = Path(f'/home/ubuntu/hoho-miniapp/frontend/{page_path}.vue')
        
        if vue_file.exists():
            print(f"✓ {page_path}")
        else:
            print(f"❌ {page_path} (文件不存在)")
            missing_pages.append(page_path)
    
    if missing_pages:
        print(f"\n⚠️  缺少 {len(missing_pages)} 个页面文件")
        return False
    else:
        print(f"\n✅ 所有 {len(pages)} 个页面文件都存在")
        return True

def check_api_config():
    """检查API配置"""
    print("\n" + "=" * 60)
    print("检查API配置")
    print("=" * 60)
    
    config_file = Path('/home/ubuntu/hoho-miniapp/frontend/api/config.js')
    
    if not config_file.exists():
        print("❌ api/config.js不存在")
        return False
    
    with open(config_file, 'r', encoding='utf-8') as f:
        content = f.read()
    
    # 检查API_BASE_URL
    if 'API_BASE_URL' in content:
        print("✓ API_BASE_URL已定义")
        
        # 检查生产环境URL
        if 'https://api.hohopark.com' in content:
            print("  ✓ 生产环境URL: https://api.hohopark.com")
        else:
            print("  ⚠️  生产环境URL可能不正确")
        
        # 检查开发环境URL
        if 'localhost:8080' in content:
            print("  ✓ 开发环境URL: http://localhost:8080")
    else:
        print("❌ API_BASE_URL未定义")
        return False
    
    # 检查API端点定义
    if 'API_ENDPOINTS' in content:
        print("✓ API_ENDPOINTS已定义")
        
        # 统计端点数量
        endpoints = re.findall(r'(\w+):\s*[\'"]', content)
        print(f"  - 定义了约 {len(endpoints)} 个API端点")
    else:
        print("⚠️  API_ENDPOINTS未定义")
    
    return True

def check_components():
    """检查组件"""
    print("\n" + "=" * 60)
    print("检查组件")
    print("=" * 60)
    
    components_dir = Path('/home/ubuntu/hoho-miniapp/frontend/components')
    
    if not components_dir.exists():
        print("⚠️  components目录不存在")
        return False
    
    components = list(components_dir.glob('**/'))
    component_files = list(components_dir.glob('**/*.vue'))
    
    print(f"✓ 找到 {len(components)-1} 个组件目录")
    print(f"✓ 找到 {len(component_files)} 个组件文件")
    
    # 检查TabBar组件
    tabbar_comp = components_dir / 'TabBar' / 'index.vue'
    if tabbar_comp.exists():
        print("  ✓ TabBar组件存在")
    else:
        print("  ⚠️  TabBar组件不存在")
    
    return True

def check_static_resources():
    """检查静态资源"""
    print("\n" + "=" * 60)
    print("检查静态资源")
    print("=" * 60)
    
    static_dir = Path('/home/ubuntu/hoho-miniapp/frontend/static')
    
    if not static_dir.exists():
        print("⚠️  static目录不存在")
        return False
    
    # 检查图标
    icons_dir = static_dir / 'icons'
    if icons_dir.exists():
        icons = list(icons_dir.glob('*.png'))
        print(f"✓ 找到 {len(icons)} 个图标文件")
    else:
        print("⚠️  icons目录不存在")
    
    # 检查图片
    images = list(static_dir.glob('**/*.png')) + list(static_dir.glob('**/*.jpg'))
    print(f"✓ 共找到 {len(images)} 个图片文件")
    
    return True

def check_manifest():
    """检查manifest.json"""
    print("\n" + "=" * 60)
    print("检查manifest.json")
    print("=" * 60)
    
    manifest_file = Path('/home/ubuntu/hoho-miniapp/frontend/manifest.json')
    
    if not manifest_file.exists():
        print("⚠️  manifest.json不存在")
        return False
    
    with open(manifest_file, 'r', encoding='utf-8') as f:
        manifest = json.load(f)
    
    # 检查基本信息
    print(f"✓ 应用名称: {manifest.get('name', 'N/A')}")
    print(f"✓ 应用ID: {manifest.get('appid', 'N/A')}")
    print(f"✓ 版本: {manifest.get('versionName', 'N/A')}")
    
    # 检查微信小程序配置
    mp_weixin = manifest.get('mp-weixin', {})
    if mp_weixin:
        print("✓ 微信小程序配置存在")
        print(f"  - AppID: {mp_weixin.get('appid', 'N/A')}")
    else:
        print("⚠️  微信小程序配置不存在")
    
    return True

def check_package_json():
    """检查package.json"""
    print("\n" + "=" * 60)
    print("检查package.json")
    print("=" * 60)
    
    package_file = Path('/home/ubuntu/hoho-miniapp/frontend/package.json')
    
    if not package_file.exists():
        print("❌ package.json不存在")
        return False
    
    with open(package_file, 'r', encoding='utf-8') as f:
        package = json.load(f)
    
    # 检查基本信息
    print(f"✓ 项目名称: {package.get('name', 'N/A')}")
    print(f"✓ 版本: {package.get('version', 'N/A')}")
    
    # 检查脚本
    scripts = package.get('scripts', {})
    if scripts:
        print(f"✓ 定义了 {len(scripts)} 个脚本命令")
        if 'build:mp-weixin' in scripts:
            print("  ✓ 包含微信小程序编译命令")
    
    # 检查依赖
    dependencies = package.get('dependencies', {})
    dev_dependencies = package.get('devDependencies', {})
    
    print(f"✓ 生产依赖: {len(dependencies)} 个")
    print(f"✓ 开发依赖: {len(dev_dependencies)} 个")
    
    # 检查关键依赖
    if 'vue' in dependencies:
        print(f"  ✓ Vue版本: {dependencies['vue']}")
    
    return True

def main():
    print("\n🔍 开始前端代码检查...\n")
    
    results = []
    
    results.append(('pages.json配置', check_pages_config()))
    results.append(('页面文件', check_pages_files()))
    results.append(('API配置', check_api_config()))
    results.append(('组件', check_components()))
    results.append(('静态资源', check_static_resources()))
    results.append(('manifest.json', check_manifest()))
    results.append(('package.json', check_package_json()))
    
    # 打印总结
    print("\n" + "=" * 60)
    print("检查总结")
    print("=" * 60)
    
    passed = sum(1 for _, result in results if result)
    total = len(results)
    
    for name, result in results:
        status = "✅" if result else "❌"
        print(f"{status} {name}")
    
    print(f"\n通过率: {passed}/{total} ({passed/total*100:.1f}%)")
    
    if passed == total:
        print("\n✅ 所有检查通过！前端代码结构正常。")
        return 0
    else:
        print(f"\n⚠️  有 {total-passed} 项检查未通过，请查看详细信息。")
        return 1

if __name__ == '__main__':
    exit(main())
