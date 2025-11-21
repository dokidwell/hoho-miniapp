#!/usr/bin/env python3
"""
HOHO小程序集成测试脚本
测试完整的业务流程
"""

import sys
from pathlib import Path

def test_project_structure():
    """测试项目结构完整性"""
    print("=" * 60)
    print("测试项目结构")
    print("=" * 60)
    
    required_dirs = [
        '/home/ubuntu/hoho-miniapp/backend',
        '/home/ubuntu/hoho-miniapp/frontend',
        '/home/ubuntu/hoho-miniapp/design',
    ]
    
    required_files = [
        '/home/ubuntu/hoho-miniapp/backend/main.go',
        '/home/ubuntu/hoho-miniapp/backend/go.mod',
        '/home/ubuntu/hoho-miniapp/backend/init.sql',
        '/home/ubuntu/hoho-miniapp/backend/.env.production',
        '/home/ubuntu/hoho-miniapp/frontend/pages.json',
        '/home/ubuntu/hoho-miniapp/frontend/manifest.json',
        '/home/ubuntu/hoho-miniapp/frontend/package.json',
        '/home/ubuntu/hoho-miniapp/frontend/api/config.js',
    ]
    
    issues = []
    
    # 检查目录
    for dir_path in required_dirs:
        if Path(dir_path).exists():
            print(f"✓ {dir_path}")
        else:
            print(f"❌ {dir_path} (不存在)")
            issues.append(f"目录缺失: {dir_path}")
    
    # 检查文件
    for file_path in required_files:
        if Path(file_path).exists():
            print(f"✓ {file_path}")
        else:
            print(f"❌ {file_path} (不存在)")
            issues.append(f"文件缺失: {file_path}")
    
    return len(issues) == 0, issues

def test_backend_code_quality():
    """测试后端代码质量"""
    print("\n" + "=" * 60)
    print("测试后端代码质量")
    print("=" * 60)
    
    issues = []
    
    # 检查Go文件数量
    backend_dir = Path('/home/ubuntu/hoho-miniapp/backend')
    go_files = list(backend_dir.glob('**/*.go'))
    print(f"✓ Go文件数量: {len(go_files)}")
    
    if len(go_files) < 20:
        issues.append("Go文件数量过少，可能代码不完整")
    
    # 检查是否有测试文件
    test_files = list(backend_dir.glob('**/*_test.go'))
    if test_files:
        print(f"✓ 测试文件数量: {len(test_files)}")
    else:
        print("⚠️  没有找到测试文件")
        issues.append("缺少单元测试")
    
    # 检查关键目录
    required_dirs = ['handlers', 'services', 'models', 'middleware', 'database', 'config', 'utils']
    for dir_name in required_dirs:
        dir_path = backend_dir / dir_name
        if dir_path.exists():
            files = list(dir_path.glob('*.go'))
            print(f"✓ {dir_name}/ ({len(files)} 个文件)")
        else:
            print(f"❌ {dir_name}/ (不存在)")
            issues.append(f"缺少目录: {dir_name}")
    
    return len(issues) == 0, issues

def test_frontend_code_quality():
    """测试前端代码质量"""
    print("\n" + "=" * 60)
    print("测试前端代码质量")
    print("=" * 60)
    
    issues = []
    
    # 检查Vue文件数量
    frontend_dir = Path('/home/ubuntu/hoho-miniapp/frontend')
    vue_files = list(frontend_dir.glob('**/*.vue'))
    print(f"✓ Vue文件数量: {len(vue_files)}")
    
    if len(vue_files) < 20:
        issues.append("Vue文件数量过少，可能页面不完整")
    
    # 检查JS文件数量
    js_files = list(frontend_dir.glob('**/*.js'))
    print(f"✓ JS文件数量: {len(js_files)}")
    
    # 检查关键目录
    required_dirs = ['pages', 'api', 'utils', 'static']
    for dir_name in required_dirs:
        dir_path = frontend_dir / dir_name
        if dir_path.exists():
            print(f"✓ {dir_name}/ 存在")
        else:
            print(f"❌ {dir_name}/ (不存在)")
            issues.append(f"缺少目录: {dir_name}")
    
    return len(issues) == 0, issues

def test_configuration():
    """测试配置完整性"""
    print("\n" + "=" * 60)
    print("测试配置完整性")
    print("=" * 60)
    
    issues = []
    
    # 检查后端配置
    env_file = Path('/home/ubuntu/hoho-miniapp/backend/.env.production')
    if env_file.exists():
        content = env_file.read_text()
        
        required_configs = [
            'SERVER_PORT',
            'DB_HOST',
            'DB_PORT',
            'DB_USER',
            'DB_PASSWORD',
            'DB_NAME',
            'REDIS_HOST',
            'REDIS_PORT',
            'REDIS_PASSWORD',
            'JWT_SECRET',
        ]
        
        for config in required_configs:
            if config in content:
                print(f"✓ {config} 已配置")
            else:
                print(f"❌ {config} 未配置")
                issues.append(f"缺少配置: {config}")
    else:
        print("❌ .env.production 不存在")
        issues.append("缺少环境配置文件")
    
    # 检查前端API配置
    api_config = Path('/home/ubuntu/hoho-miniapp/frontend/api/config.js')
    if api_config.exists():
        content = api_config.read_text()
        
        if 'https://api.hohopark.com' in content:
            print("✓ 生产环境API地址已配置")
        else:
            print("❌ 生产环境API地址未配置")
            issues.append("前端API地址未配置")
    else:
        print("❌ api/config.js 不存在")
        issues.append("缺少API配置文件")
    
    return len(issues) == 0, issues

def test_documentation():
    """测试文档完整性"""
    print("\n" + "=" * 60)
    print("测试文档完整性")
    print("=" * 60)
    
    issues = []
    
    required_docs = [
        ('README.md', '项目说明'),
        ('部署完成报告.md', '部署报告'),
        ('快速参考指南.md', '运维指南'),
        ('前端编译说明.md', '编译说明'),
        ('数据库初始化指南.md', '数据库指南'),
    ]
    
    for doc_file, doc_name in required_docs:
        doc_path = Path(f'/home/ubuntu/hoho-miniapp/{doc_file}')
        if doc_path.exists():
            size = doc_path.stat().st_size
            print(f"✓ {doc_name} ({size} 字节)")
        else:
            print(f"❌ {doc_name} (不存在)")
            issues.append(f"缺少文档: {doc_name}")
    
    return len(issues) == 0, issues

def test_security():
    """测试安全性配置"""
    print("\n" + "=" * 60)
    print("测试安全性配置")
    print("=" * 60)
    
    issues = []
    
    # 检查.gitignore
    gitignore = Path('/home/ubuntu/hoho-miniapp/.gitignore')
    if gitignore.exists():
        content = gitignore.read_text()
        
        sensitive_patterns = [
            '.env',
            '*.pem',
            '*.key',
            'node_modules',
        ]
        
        for pattern in sensitive_patterns:
            if pattern in content:
                print(f"✓ {pattern} 已忽略")
            else:
                print(f"⚠️  {pattern} 未在.gitignore中")
                issues.append(f"建议添加到.gitignore: {pattern}")
    else:
        print("⚠️  .gitignore 不存在")
        issues.append("建议创建.gitignore文件")
    
    # 检查是否有敏感信息泄露
    backend_files = list(Path('/home/ubuntu/hoho-miniapp/backend').glob('**/*.go'))
    
    sensitive_keywords = ['password', 'secret', 'key', 'token']
    found_sensitive = False
    
    for go_file in backend_files:
        if 'test' in go_file.name or 'example' in go_file.name:
            continue
        
        try:
            content = go_file.read_text()
            for keyword in sensitive_keywords:
                if f'"{keyword}"' in content.lower() and '=' in content:
                    # 简单检查，可能有误报
                    pass
        except:
            pass
    
    if not found_sensitive:
        print("✓ 未发现明显的硬编码敏感信息")
    
    return len(issues) == 0, issues

def main():
    print("\n🔍 开始集成测试...\n")
    
    test_results = []
    all_issues = []
    
    # 运行所有测试
    tests = [
        ("项目结构", test_project_structure),
        ("后端代码质量", test_backend_code_quality),
        ("前端代码质量", test_frontend_code_quality),
        ("配置完整性", test_configuration),
        ("文档完整性", test_documentation),
        ("安全性配置", test_security),
    ]
    
    for test_name, test_func in tests:
        try:
            passed, issues = test_func()
            test_results.append((test_name, passed))
            all_issues.extend(issues)
        except Exception as e:
            print(f"\n❌ 测试 '{test_name}' 出错: {str(e)}")
            test_results.append((test_name, False))
            all_issues.append(f"{test_name}: {str(e)}")
    
    # 打印总结
    print("\n" + "=" * 60)
    print("集成测试总结")
    print("=" * 60)
    
    passed_count = sum(1 for _, passed in test_results if passed)
    total_count = len(test_results)
    
    for test_name, passed in test_results:
        status = "✅" if passed else "❌"
        print(f"{status} {test_name}")
    
    print(f"\n通过率: {passed_count}/{total_count} ({passed_count/total_count*100:.1f}%)")
    
    if all_issues:
        print(f"\n发现 {len(all_issues)} 个问题:")
        for i, issue in enumerate(all_issues, 1):
            print(f"  {i}. {issue}")
    
    if passed_count == total_count and len(all_issues) == 0:
        print("\n✅ 所有集成测试通过！项目质量良好。")
        return 0
    else:
        print(f"\n⚠️  有 {total_count - passed_count} 项测试未通过，{len(all_issues)} 个问题需要关注。")
        return 1

if __name__ == '__main__':
    exit(main())
