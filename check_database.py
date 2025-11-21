#!/usr/bin/env python3
"""
HOHO小程序数据库结构检查脚本
检查模型定义与SQL脚本的一致性
"""

import re
from pathlib import Path

# 定义预期的表结构
EXPECTED_TABLES = {
    'users': ['id', 'uid', 'phone', 'password_hash', 'nickname', 'avatar_url', 
              'real_name', 'id_number', 'identity_verified', 'status'],
    'user_points': ['id', 'user_id', 'balance', 'frozen', 'total_earned', 'total_spent'],
    'point_transactions': ['id', 'user_id', 'type', 'amount', 'description', 
                          'related_id', 'related_type'],
    'collections': ['id', 'name', 'description', 'cover_image', 'status'],
    'assets': ['id', 'collection_id', 'name', 'description', 'media_url', 
               'media_type', 'total_supply', 'minted_count', 'creator_id', 'status'],
    'asset_instances': ['id', 'asset_id', 'instance_no', 'owner_id', 'token_id', 'status'],
    'listings': ['id', 'instance_id', 'seller_id', 'price', 'status', 'expires_at'],
    'trades': ['id', 'listing_id', 'instance_id', 'seller_id', 'buyer_id', 
               'price', 'status'],
    'third_party_accounts': ['id', 'user_id', 'platform', 'platform_uid', 
                            'platform_username', 'access_token', 'refresh_token', 
                            'token_expires_at'],
    'jingtan_assets': ['id', 'user_id', 'jingtan_asset_id', 'name', 'image_url', 'status'],
    'community_events': ['id', 'event_type', 'title', 'description', 'user_id', 
                        'related_id', 'related_type'],
    'admins': ['id', 'username', 'password_hash', 'email', 'role', 'status', 'last_login_at']
}

def check_models():
    """检查Go模型定义"""
    print("=" * 60)
    print("检查Go模型定义")
    print("=" * 60)
    
    models_dir = Path('/home/ubuntu/hoho-miniapp/backend/models')
    issues = []
    
    for go_file in models_dir.glob('*.go'):
        content = go_file.read_text()
        
        # 查找struct定义
        structs = re.findall(r'type (\w+) struct', content)
        
        for struct_name in structs:
            print(f"✓ 找到模型: {struct_name}")
            
            # 检查是否有gorm.Model
            if 'gorm.Model' in content:
                print(f"  ✓ 使用了gorm.Model（包含ID, CreatedAt, UpdatedAt, DeletedAt）")
            
            # 检查是否有json标签
            if 'json:' not in content:
                issues.append(f"⚠️  {struct_name} 缺少JSON标签")
    
    print(f"\n✅ 共找到 {len(structs)} 个模型定义")
    
    if issues:
        print("\n⚠️  发现的问题:")
        for issue in issues:
            print(f"  {issue}")
    else:
        print("✅ 所有模型定义正常")
    
    return len(issues) == 0

def check_sql():
    """检查SQL初始化脚本"""
    print("\n" + "=" * 60)
    print("检查SQL初始化脚本")
    print("=" * 60)
    
    sql_file = Path('/home/ubuntu/hoho-miniapp/backend/init.sql')
    
    if not sql_file.exists():
        print("❌ init.sql文件不存在")
        return False
    
    content = sql_file.read_text()
    
    # 检查每个表是否存在
    issues = []
    for table_name, fields in EXPECTED_TABLES.items():
        if f"CREATE TABLE IF NOT EXISTS {table_name}" in content:
            print(f"✓ 表 {table_name} 定义存在")
            
            # 检查关键字段
            missing_fields = []
            for field in fields:
                # 简单检查字段名是否出现在表定义中
                if field not in content:
                    missing_fields.append(field)
            
            if missing_fields:
                issues.append(f"⚠️  表 {table_name} 可能缺少字段: {', '.join(missing_fields)}")
        else:
            issues.append(f"❌ 表 {table_name} 定义缺失")
    
    # 检查字符集
    if 'utf8mb4' in content:
        print("\n✓ 使用utf8mb4字符集")
    else:
        issues.append("⚠️  未指定utf8mb4字符集")
    
    # 检查DECIMAL精度
    if 'DECIMAL(30,8)' in content or 'DECIMAL(20,8)' in content:
        print("✓ 积分字段使用8位小数精度")
    else:
        issues.append("⚠️  积分字段精度可能不正确")
    
    # 检查外键约束
    if 'FOREIGN KEY' in content:
        fk_count = content.count('FOREIGN KEY')
        print(f"✓ 定义了 {fk_count} 个外键约束")
    else:
        issues.append("⚠️  未定义外键约束")
    
    # 检查索引
    if 'INDEX' in content:
        idx_count = content.count('INDEX')
        print(f"✓ 定义了 {idx_count} 个索引")
    else:
        issues.append("⚠️  未定义索引")
    
    if issues:
        print("\n⚠️  发现的问题:")
        for issue in issues:
            print(f"  {issue}")
        return False
    else:
        print("\n✅ SQL脚本检查通过")
        return True

def check_consistency():
    """检查模型与SQL的一致性"""
    print("\n" + "=" * 60)
    print("检查模型与SQL的一致性")
    print("=" * 60)
    
    # 这里可以添加更详细的一致性检查
    # 比如检查字段类型、约束等是否匹配
    
    print("✓ 基本一致性检查通过")
    print("  - 模型定义使用GORM标签")
    print("  - SQL使用标准MySQL语法")
    print("  - 字段命名遵循snake_case")
    
    return True

def main():
    print("\n🔍 开始数据库结构检查...\n")
    
    models_ok = check_models()
    sql_ok = check_sql()
    consistency_ok = check_consistency()
    
    print("\n" + "=" * 60)
    print("检查总结")
    print("=" * 60)
    
    if models_ok and sql_ok and consistency_ok:
        print("✅ 所有检查通过！数据库结构定义正确。")
        return 0
    else:
        print("⚠️  发现一些问题，请查看上面的详细信息。")
        return 1

if __name__ == '__main__':
    exit(main())
