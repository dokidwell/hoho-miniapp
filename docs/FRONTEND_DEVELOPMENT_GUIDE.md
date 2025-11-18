# HOHO小程序前端开发指南

## 📋 项目概述

本文档为HOHO数字藏品集换小程序的前端开发提供完整指导。后端已100%完成，前端需要基于Uni-app框架开发，使用白色主题的极简风格，精准复刻设计稿。

## 🎨 设计规范

### 颜色规范
```scss
// 主色调
--color-primary: #000000;           // 主色（黑色）
--color-accent: #0066FF;            // 强调色（蓝色）

// 文本色
--color-text-primary: #000000;      // 主文本
--color-text-secondary: #666666;    // 次要文本
--color-text-tertiary: #999999;     // 三级文本
--color-text-white: #FFFFFF;        // 白色文本

// 背景色
--color-bg-primary: #FFFFFF;        // 主背景
--color-bg-secondary: #F5F5F5;      // 次要背景
--color-bg-tertiary: #FAFAFA;       // 三级背景

// 边框色
--color-border: #E8E8E8;            // 边框
--color-border-light: #F0F0F0;      // 浅色边框
```

### 字体规范
```scss
--font-size-xs: 20rpx;    // 12px
--font-size-sm: 24rpx;    // 14px
--font-size-base: 28rpx;  // 16px
--font-size-lg: 32rpx;    // 18px
--font-size-xl: 36rpx;    // 20px
--font-size-2xl: 40rpx;   // 22px
--font-size-3xl: 48rpx;   // 26px
```

### 间距规范
```scss
--spacing-xs: 8rpx;       // 4px
--spacing-sm: 16rpx;      // 8px
--spacing-md: 24rpx;      // 12px
--spacing-lg: 32rpx;      // 16px
--spacing-xl: 48rpx;      // 24px
--spacing-2xl: 64rpx;     // 32px
```

### 圆角规范
```scss
--radius-sm: 8rpx;        // 小圆角
--radius-md: 12rpx;       // 中圆角
--radius-lg: 16rpx;       // 大圆角
--radius-xl: 24rpx;       // 超大圆角
--radius-full: 9999rpx;   // 完全圆角
```

## 📱 页面清单

### 核心页面（5个）
1. **首页（作品列表）** - `/pages/index/index.vue`
2. **集换中心** - `/pages/jijhuan/index.vue`
3. **我的（个人中心）** - `/pages/profile/index.vue`
4. **生态页面** - `/pages/ecology/index.vue`
5. **透明公示** - `/pages/transparent-ledger/index.vue`

### 功能页面（8个）
6. **创作（铸造）** - `/pages/create/index.vue`
7. **藏品详情** - `/pages/asset-detail/index.vue`
8. **鲸探作品** - `/pages/jingtan-assets/index.vue`
9. **第三方关联** - `/pages/third-party/index.vue`
10. **挂售页面** - `/pages/listing-create/index.vue`
11. **兑换页面** - `/pages/exchange/index.vue`
12. **登录页面** - `/pages/login/index.vue`
13. **注册页面** - `/pages/register/index.vue`

## 🔌 后端API端点

### 基础配置
```javascript
// frontend/api/config.js
export const API_BASE_URL = 'http://your-server-ip:8080'  // 替换为实际服务器地址

export const API_ENDPOINTS = {
  // 用户相关
  USER: {
    REGISTER: '/api/v1/users/register',
    LOGIN: '/api/v1/users/login',
    PROFILE: '/api/v1/users/profile',
    UPDATE_PROFILE: '/api/v1/users/profile',
    VERIFY_IDENTITY: '/api/v1/users/verify-identity',
    GET_POINTS: '/api/v1/users/points',
  },
  
  // 藏品相关
  ASSET: {
    LIST: '/api/v1/assets',
    DETAIL: '/api/v1/assets/:id',
    CREATE: '/api/v1/assets',  // 铸造
  },
  
  // 交易相关
  TRADE: {
    GET_LISTINGS: '/api/v1/listings',
    CREATE_LISTING: '/api/v1/listings',
    CANCEL_LISTING: '/api/v1/listings/:id/cancel',
    EXECUTE_TRADE: '/api/v1/trades',
  },
  
  // 社区事件
  EVENT: {
    LIST: '/api/v1/events',
    DETAIL: '/api/v1/events/:id',
  },
  
  // 鲸探API
  JINGTAN: {
    BIND: '/api/v1/jingtan/bind',
    UNBIND: '/api/v1/jingtan/unbind',
    SYNC: '/api/v1/jingtan/sync',
    ASSETS: '/api/v1/jingtan/assets',
  },
}
```

## 📄 页面详细开发说明

---

## 1. 首页（作品列表）

### 设计稿参考
- 文件：`作品.png`
- 特点：顶部Banner + 卡片式藏品展示

### 页面结构
```vue
<template>
  <view class="page">
    <!-- Banner区域 -->
    <view class="banner-section">
      <image class="banner-image" src="/static/images/banner-welcome.png" mode="aspectFill" />
      <view class="banner-text">欢迎来到HOHO Park!</view>
    </view>

    <!-- 藏品卡片列表 -->
    <view class="cards-section">
      <view v-for="asset in assets" :key="asset.id" class="asset-card" @click="goToDetail(asset.id)">
        <image class="card-image" :src="asset.image_url" mode="aspectFill" />
        <view class="card-action">
          <view class="action-btn">立即查看</view>
        </view>
      </view>
    </view>

    <!-- 底部导航栏 -->
    <TabBar :active="0" />
  </view>
</template>
```

### API调用
```javascript
// 获取藏品列表
async function fetchAssets() {
  const res = await request.get(API_ENDPOINTS.ASSET.LIST, {
    page: 1,
    page_size: 20,
    status: 'approved'  // 只显示已审核通过的藏品
  })
  assets.value = res.data.list
}
```

### 样式要点
- Banner高度：600rpx
- 卡片间距：32rpx
- 卡片圆角：16rpx
- "立即查看"按钮：黑色半透明背景，白色文字

---

## 2. 集换中心

### 设计稿参考
- 文件：`集换.png`
- 特点：搜索框 + 分类Tab + 2列网格布局

### 页面结构
```vue
<template>
  <view class="page">
    <!-- 顶部搜索栏 -->
    <view class="header">
      <view class="search-box">
        <image class="search-icon" src="/static/icons/search.png" />
        <input class="search-input" placeholder="搜索..." v-model="keyword" />
      </view>
      <view class="filter-btn" @click="showFilter">
        <image class="filter-icon" src="/static/icons/filter.png" />
      </view>
    </view>

    <!-- 集换专区标题 -->
    <view class="section-title">集换专区</view>

    <!-- 分类Tab -->
    <scroll-view class="category-tabs" scroll-x>
      <view class="tab-item" 
        v-for="cat in categories" 
        :key="cat.value"
        :class="{ 'tab-active': currentCategory === cat.value }"
        @click="selectCategory(cat.value)">
        {{ cat.label }}
      </view>
    </scroll-view>

    <!-- 藏品网格（2列） -->
    <scroll-view class="asset-grid" scroll-y @scrolltolower="loadMore">
      <view class="grid-wrapper">
        <view v-for="item in listings" :key="item.id" class="asset-card" @click="goToDetail(item.id)">
          <!-- 藏品图片 -->
          <view class="asset-image-wrapper">
            <image class="asset-image" :src="item.asset.image_url" mode="aspectFill" />
            
            <!-- 左上角标签 -->
            <view v-if="item.is_official" class="asset-tag">
              <image class="tag-icon" src="/static/icons/official.png" />
              <text class="tag-text">合集作品</text>
            </view>
            
            <!-- 左下角可兑数量 -->
            <view class="asset-available">
              <image class="available-icon" src="/static/icons/clock.png" />
              <text class="available-text">{{ item.available_count }}份可兑</text>
            </view>
          </view>

          <!-- 藏品信息 -->
          <view class="asset-info">
            <view class="asset-name">{{ item.asset.name }}</view>
            <view class="asset-price-row">
              <text class="price-value number-display">{{ formatPrice(item.price) }}</text>
              <text class="price-unit">起</text>
            </view>
            <view class="asset-supply">
              <image class="supply-icon" src="/static/icons/total.png" />
              <text class="supply-text">{{ item.total_supply }}份</text>
            </view>
          </view>
        </view>
      </view>
    </scroll-view>

    <!-- 底部导航栏 -->
    <TabBar :active="2" />
  </view>
</template>
```

### API调用
```javascript
// 获取挂售列表
async function fetchListings() {
  const res = await request.get(API_ENDPOINTS.TRADE.GET_LISTINGS, {
    page: page.value,
    page_size: 20,
    category: currentCategory.value === 'all' ? '' : currentCategory.value,
    sort_by: 'price',  // price, created_at
    sort_order: 'asc'  // asc, desc
  })
  listings.value = [...listings.value, ...res.data.list]
}
```

### 数据格式
```javascript
// 挂售列表返回格式
{
  "code": 0,
  "message": "success",
  "data": {
    "list": [
      {
        "id": 1,
        "asset": {
          "id": 1,
          "name": "HOHO的1000天",
          "image_url": "https://...",
          "total_supply": 5000
        },
        "price": "18500.00000000",
        "available_count": 61,
        "is_official": true
      }
    ],
    "total": 100,
    "page": 1,
    "page_size": 20
  }
}
```

### 样式要点
- 网格布局：2列，间距24rpx
- 卡片圆角：16rpx
- 标签背景：黑色半透明（rgba(0, 0, 0, 0.7)）
- 价格字体：等宽字体（monospace）

---

## 3. 我的（个人中心）

### 设计稿参考
- 文件：`我的.png`
- 特点：用户信息卡 + 功能网格 + 列表项

### 页面结构
```vue
<template>
  <view class="page">
    <!-- 顶部导航 -->
    <view class="navbar">
      <view class="back-btn" @click="goBack">
        <image src="/static/icons/back.png" />
      </view>
      <view class="nav-actions">
        <image src="/static/icons/more.png" @click="showMore" />
        <image src="/static/icons/scan.png" @click="scanCode" />
      </view>
    </view>

    <!-- 用户信息卡 -->
    <view class="user-card">
      <view class="user-header">
        <image class="user-avatar" :src="userInfo.avatar_url || '/static/images/default-avatar.png'" />
        <view class="user-info">
          <view class="user-phone">{{ maskPhone(userInfo.phone) }}</view>
          <view class="user-tags">
            <view class="tag">未认证</view>
            <view class="tag">未绑定鲸探</view>
          </view>
        </view>
        <image class="arrow-icon" src="/static/icons/arrow-right.png" />
      </view>

      <!-- 野生HOHO等级卡 -->
      <view class="level-card">
        <view class="level-info">
          <text class="level-title">野生HOHO</text>
          <text class="level-value">Lv.1</text>
        </view>
      </view>

      <!-- 任务中心 -->
      <view class="task-center">
        <image class="task-icon" src="/static/icons/diamond.png" />
        <text class="task-text">积分不够？来...</text>
        <view class="task-btn">任务中心</view>
      </view>
    </view>

    <!-- 功能网格（4x2） -->
    <view class="function-grid">
      <view class="grid-item" @click="goTo('/pages/my-assets/index')">
        <image class="grid-icon" src="/static/icons/collection.png" />
        <text class="grid-label">作品集</text>
      </view>
      <view class="grid-item" @click="goTo('/pages/shop/index')">
        <image class="grid-icon" src="/static/icons/shop.png" />
        <text class="grid-label">周边</text>
      </view>
      <view class="grid-item" @click="goTo('/pages/service/index')">
        <image class="grid-icon" src="/static/icons/service.png" />
        <text class="grid-label">客服</text>
      </view>
      <view class="grid-item" @click="goTo('/pages/settings/index')">
        <image class="grid-icon" src="/static/icons/settings.png" />
        <text class="grid-label">设置</text>
      </view>
      <view class="grid-item" @click="goTo('/pages/trade-history/index')">
        <image class="grid-icon" src="/static/icons/exchange-history.png" />
        <text class="grid-label">集换记录</text>
      </view>
      <view class="grid-item" @click="goTo('/pages/community-assets/index')">
        <image class="grid-icon" src="/static/icons/community.png" />
        <text class="grid-label">社区作品</text>
      </view>
      <view class="grid-item" @click="goTo('/pages/jingtan-assets/index')">
        <image class="grid-icon" src="/static/icons/jingtan.png" />
        <text class="grid-label">鲸探作品</text>
      </view>
      <view class="grid-item" @click="goTo('/pages/waveup-assets/index')">
        <image class="grid-icon" src="/static/icons/waveup.png" />
        <text class="grid-label">WAVEUP作品</text>
      </view>
    </view>

    <!-- 列表项 -->
    <view class="list-section">
      <view class="list-item" @click="goTo('/pages/identity-verify/index')">
        <text class="list-label">身份认证</text>
        <image class="list-arrow" src="/static/icons/arrow-right.png" />
      </view>
      
      <view class="list-item">
        <text class="list-label">UID</text>
        <text class="list-value">{{ userInfo.uid }}</text>
      </view>
      
      <view class="list-item" @click="goTo('/pages/third-party/index')">
        <text class="list-label">第三方关联</text>
        <image class="list-arrow" src="/static/icons/arrow-right.png" />
      </view>
    </view>

    <!-- 热门活动 -->
    <view class="activity-section">
      <view class="section-title">热门活动</view>
      <view class="activity-cards">
        <view class="activity-card">
          <text class="activity-placeholder">暂时没有更多内容...</text>
        </view>
        <view class="activity-card">
          <text class="activity-placeholder">暂时没有更多内容...</text>
        </view>
      </view>
    </view>

    <!-- 底部导航栏 -->
    <TabBar :active="4" />
  </view>
</template>
```

### API调用
```javascript
// 获取用户信息
async function fetchUserInfo() {
  const res = await request.get(API_ENDPOINTS.USER.PROFILE)
  userInfo.value = res.data
}

// 获取积分余额
async function fetchPoints() {
  const res = await request.get(API_ENDPOINTS.USER.GET_POINTS)
  points.value = res.data.balance
}
```

### 工具函数
```javascript
// 手机号脱敏
function maskPhone(phone) {
  if (!phone) return '未绑定'
  return phone.replace(/(\d{3})\d{4}(\d{4})/, '$1****$2')
}
```

### 样式要点
- 用户卡片：白色背景，16rpx圆角，阴影
- 等级卡：深色背景，渐变效果
- 功能网格：4列2行，图标48rpx
- 列表项：白色背景，上下padding 32rpx

---

## 4. 生态页面

### 设计稿参考
- 文件：`生态.png`
- 特点：搜索框 + Banner + 功能入口 + 合作伙伴

### 页面结构
```vue
<template>
  <view class="page">
    <!-- 顶部导航 -->
    <view class="navbar">
      <view class="back-btn" @click="goBack">
        <image src="/static/icons/back.png" />
      </view>
      <view class="nav-actions">
        <image src="/static/icons/more.png" />
        <image src="/static/icons/scan.png" />
      </view>
    </view>

    <!-- 搜索框 -->
    <view class="search-section">
      <view class="search-box">
        <image class="search-icon" src="/static/icons/search.png" />
        <input class="search-input" placeholder="请输入内容" />
      </view>
    </view>

    <!-- Banner -->
    <view class="banner-section">
      <image class="banner-image" src="/static/images/ecology-banner.png" mode="aspectFill" />
    </view>

    <!-- 功能入口（左图右文） -->
    <view class="function-section">
      <image class="function-image" src="/static/images/ecology-function.png" mode="aspectFill" />
      
      <view class="function-list">
        <view class="function-item" @click="goTo('/pages/whitepaper/index')">
          <view class="function-info">
            <text class="function-title">白皮书</text>
            <text class="function-desc">完成了解HOHO IP进程</text>
          </view>
          <image class="function-icon" src="/static/icons/whitepaper.png" />
        </view>
        
        <view class="function-item" @click="goTo('/pages/governance/index')">
          <view class="function-info">
            <text class="function-title">社区治理</text>
            <text class="function-desc">每一位社区成员都可以参加</text>
          </view>
          <image class="function-icon" src="/static/icons/governance.png" />
        </view>
        
        <view class="function-item" @click="goTo('/pages/transparent-ledger/index')">
          <view class="function-info">
            <text class="function-title">透明公示</text>
            <text class="function-desc">不可篡改的透明制度</text>
          </view>
          <image class="function-icon" src="/static/icons/transparent.png" />
        </view>
      </view>
    </view>

    <!-- 流动广告 -->
    <view class="ad-section">
      <image class="ad-icon" src="/static/icons/ad.png" />
      <text class="ad-text">流动广告流动广告流动广告流动广告流动广告</text>
    </view>

    <!-- HOHO和他的朋友们 -->
    <view class="partners-section">
      <view class="section-title">HOHO和他的朋友们</view>
      <view class="partners-grid">
        <view class="partner-item" v-for="partner in partners" :key="partner.id">
          <image class="partner-logo" :src="partner.logo" mode="aspectFit" />
        </view>
      </view>
    </view>

    <!-- 底部导航栏 -->
    <TabBar :active="3" />
  </view>
</template>
```

### 数据
```javascript
const partners = ref([
  { id: 1, name: '鲸探', logo: '/static/images/partner-jingtan.png' },
  { id: 2, name: 'Waveup', logo: '/static/images/partner-waveup.png' },
  { id: 3, name: 'XMeta', logo: '/static/images/partner-xmeta.png' },
  { id: 4, name: '品拍', logo: '/static/images/partner-pinpai.png' },
  { id: 5, name: '敬请期待', logo: '/static/images/partner-coming.png' },
])
```

### 样式要点
- Banner高度：400rpx
- 功能项：白色背景，左右padding 32rpx
- 合作伙伴：圆形logo，直径120rpx
- 流动广告：滚动效果

---

## 5. 透明公示

### 设计稿参考
- 文件：`透明公示.png`、`透明公示（1）.png`、`透明公示（2）.png`
- 特点：事件列表 + 详情页

### 页面结构
```vue
<template>
  <view class="page">
    <!-- 顶部导航 -->
    <view class="navbar">
      <view class="navbar-title">透明公示</view>
    </view>

    <!-- 事件类型筛选 -->
    <scroll-view class="filter-tabs" scroll-x>
      <view class="tab-item" 
        v-for="type in eventTypes" 
        :key="type.value"
        :class="{ 'tab-active': currentType === type.value }"
        @click="selectType(type.value)">
        {{ type.label }}
      </view>
    </scroll-view>

    <!-- 事件列表 -->
    <scroll-view class="event-list" scroll-y @scrolltolower="loadMore">
      <view v-for="event in events" :key="event.id" class="event-item" @click="goToDetail(event.id)">
        <view class="event-header">
          <view class="event-type-tag" :class="`type-${event.event_type}`">
            {{ getEventTypeName(event.event_type) }}
          </view>
          <view class="event-time">{{ formatTime(event.created_at) }}</view>
        </view>
        
        <view class="event-content">
          <text class="event-desc">{{ event.description }}</text>
        </view>
        
        <view class="event-footer">
          <view class="event-id">ID: {{ event.id }}</view>
          <view class="event-hash">Hash: {{ shortHash(event.hash) }}</view>
        </view>
      </view>
    </scroll-view>

    <!-- 底部导航栏 -->
    <TabBar :active="2" />
  </view>
</template>
```

### API调用
```javascript
// 获取社区事件列表
async function fetchEvents() {
  const res = await request.get(API_ENDPOINTS.EVENT.LIST, {
    page: page.value,
    page_size: 20,
    event_type: currentType.value === 'all' ? '' : currentType.value
  })
  events.value = [...events.value, ...res.data.list]
}

// 获取事件详情
async function fetchEventDetail(eventId) {
  const res = await request.get(API_ENDPOINTS.EVENT.DETAIL.replace(':id', eventId))
  eventDetail.value = res.data
}
```

### 事件类型
```javascript
const eventTypes = [
  { label: '全部', value: 'all' },
  { label: '铸造', value: 'mint' },
  { label: '交易', value: 'trade' },
  { label: '空投', value: 'airdrop' },
  { label: '挂售', value: 'listing' },
  { label: '取消挂售', value: 'cancel_listing' },
]
```

### 工具函数
```javascript
// 格式化时间
function formatTime(timestamp) {
  const date = new Date(timestamp)
  return `${date.getFullYear()}-${String(date.getMonth() + 1).padStart(2, '0')}-${String(date.getDate()).padStart(2, '0')} ${String(date.getHours()).padStart(2, '0')}:${String(date.getMinutes()).padStart(2, '0')}`
}

// 短Hash
function shortHash(hash) {
  if (!hash) return '-'
  return `${hash.slice(0, 6)}...${hash.slice(-4)}`
}

// 事件类型名称
function getEventTypeName(type) {
  const map = {
    'mint': '铸造',
    'trade': '交易',
    'airdrop': '空投',
    'listing': '挂售',
    'cancel_listing': '取消挂售',
  }
  return map[type] || type
}
```

---

## 6. 创作（铸造）

### 设计稿参考
- 文件：`创作1.png`、`创作2.png`、`创作3.png`、`创作4.png`
- 特点：多步骤表单

### 页面结构
```vue
<template>
  <view class="page">
    <!-- 顶部导航 -->
    <view class="navbar">
      <view class="back-btn" @click="goBack">
        <image src="/static/icons/back.png" />
      </view>
      <view class="navbar-title">创作</view>
    </view>

    <!-- 步骤指示器 -->
    <view class="steps">
      <view class="step-item" :class="{ 'step-active': currentStep >= 1 }">
        <view class="step-number">1</view>
        <text class="step-label">上传作品</text>
      </view>
      <view class="step-line" :class="{ 'step-active': currentStep >= 2 }"></view>
      <view class="step-item" :class="{ 'step-active': currentStep >= 2 }">
        <view class="step-number">2</view>
        <text class="step-label">填写信息</text>
      </view>
      <view class="step-line" :class="{ 'step-active': currentStep >= 3 }"></view>
      <view class="step-item" :class="{ 'step-active': currentStep >= 3 }">
        <view class="step-number">3</view>
        <text class="step-label">确认提交</text>
      </view>
    </view>

    <!-- 步骤1：上传作品 -->
    <view v-if="currentStep === 1" class="step-content">
      <view class="upload-section">
        <view class="upload-box" @click="chooseImage">
          <image v-if="formData.image_url" class="preview-image" :src="formData.image_url" mode="aspectFill" />
          <view v-else class="upload-placeholder">
            <image class="upload-icon" src="/static/icons/upload.png" />
            <text class="upload-text">点击上传作品</text>
            <text class="upload-hint">支持图片、视频、音频</text>
          </view>
        </view>
      </view>
      
      <view class="btn-group">
        <button class="btn btn-primary" @click="nextStep">下一步</button>
      </view>
    </view>

    <!-- 步骤2：填写信息 -->
    <view v-if="currentStep === 2" class="step-content">
      <view class="form-section">
        <view class="form-item">
          <text class="form-label">作品名称</text>
          <input class="form-input" v-model="formData.name" placeholder="请输入作品名称" />
        </view>
        
        <view class="form-item">
          <text class="form-label">作品描述</text>
          <textarea class="form-textarea" v-model="formData.description" placeholder="请输入作品描述" />
        </view>
        
        <view class="form-item">
          <text class="form-label">制作数量</text>
          <input class="form-input" type="number" v-model="formData.total_supply" placeholder="请输入制作数量" />
        </view>
        
        <view class="form-item">
          <text class="form-label">分类</text>
          <picker mode="selector" :range="categories" range-key="label" @change="onCategoryChange">
            <view class="picker-value">
              {{ selectedCategory ? selectedCategory.label : '请选择分类' }}
            </view>
          </picker>
        </view>
      </view>
      
      <view class="btn-group">
        <button class="btn btn-secondary" @click="prevStep">上一步</button>
        <button class="btn btn-primary" @click="nextStep">下一步</button>
      </view>
    </view>

    <!-- 步骤3：确认提交 -->
    <view v-if="currentStep === 3" class="step-content">
      <view class="preview-section">
        <view class="preview-title">作品预览</view>
        <image class="preview-image-large" :src="formData.image_url" mode="aspectFill" />
        
        <view class="preview-info">
          <view class="info-row">
            <text class="info-label">作品名称</text>
            <text class="info-value">{{ formData.name }}</text>
          </view>
          <view class="info-row">
            <text class="info-label">作品描述</text>
            <text class="info-value">{{ formData.description }}</text>
          </view>
          <view class="info-row">
            <text class="info-label">制作数量</text>
            <text class="info-value">{{ formData.total_supply }}份</text>
          </view>
          <view class="info-row">
            <text class="info-label">分类</text>
            <text class="info-value">{{ selectedCategory.label }}</text>
          </view>
        </view>
      </view>
      
      <view class="btn-group">
        <button class="btn btn-secondary" @click="prevStep">上一步</button>
        <button class="btn btn-primary" @click="submitCreate" :loading="submitting">提交审核</button>
      </view>
    </view>
  </view>
</template>
```

### API调用
```javascript
// 上传图片
async function uploadImage(filePath) {
  const res = await uni.uploadFile({
    url: `${API_BASE_URL}/api/v1/upload`,
    filePath: filePath,
    name: 'file',
    header: {
      'Authorization': `Bearer ${getToken()}`
    }
  })
  const data = JSON.parse(res.data)
  return data.data.url
}

// 提交铸造请求
async function submitCreate() {
  submitting.value = true
  try {
    const res = await request.post(API_ENDPOINTS.ASSET.CREATE, {
      name: formData.name,
      description: formData.description,
      image_url: formData.image_url,
      total_supply: parseInt(formData.total_supply),
      category: selectedCategory.value
    })
    
    uni.showToast({
      title: '提交成功，等待审核',
      icon: 'success'
    })
    
    setTimeout(() => {
      uni.navigateBack()
    }, 1500)
  } catch (error) {
    uni.showToast({
      title: error.message || '提交失败',
      icon: 'none'
    })
  } finally {
    submitting.value = false
  }
}
```

---

## 7. 藏品详情

### 页面结构
```vue
<template>
  <view class="page">
    <!-- 顶部导航 -->
    <view class="navbar">
      <view class="back-btn" @click="goBack">
        <image src="/static/icons/back.png" />
      </view>
      <view class="nav-actions">
        <image src="/static/icons/share.png" @click="shareAsset" />
      </view>
    </view>

    <!-- 藏品图片 -->
    <view class="asset-image-section">
      <image class="asset-image" :src="assetDetail.image_url" mode="aspectFill" />
    </view>

    <!-- 藏品信息 -->
    <view class="asset-info-section">
      <view class="asset-header">
        <text class="asset-name">{{ assetDetail.name }}</text>
        <text class="asset-serial">#{{ assetDetail.serial_number }}</text>
      </view>
      
      <view class="asset-creator">
        <image class="creator-avatar" :src="assetDetail.creator.avatar_url" />
        <text class="creator-name">{{ assetDetail.creator.nickname }}</text>
      </view>
      
      <view class="asset-description">
        <text>{{ assetDetail.description }}</text>
      </view>
      
      <!-- 价格信息 -->
      <view class="price-section">
        <view class="price-item">
          <text class="price-label">当前价格</text>
          <text class="price-value number-display">{{ formatPoints(assetDetail.current_price) }}</text>
        </view>
        <view class="price-item">
          <text class="price-label">底价</text>
          <text class="price-value number-display">{{ formatPoints(assetDetail.floor_price) }}</text>
        </view>
      </view>
      
      <!-- 统计信息 -->
      <view class="stats-section">
        <view class="stat-item">
          <text class="stat-label">总量</text>
          <text class="stat-value">{{ assetDetail.total_supply }}</text>
        </view>
        <view class="stat-item">
          <text class="stat-label">持有人</text>
          <text class="stat-value">{{ assetDetail.holder_count }}</text>
        </view>
        <view class="stat-item">
          <text class="stat-label">交易量</text>
          <text class="stat-value">{{ assetDetail.trade_count }}</text>
        </view>
      </view>
    </view>

    <!-- 操作按钮 -->
    <view class="action-section">
      <button class="btn btn-primary" @click="buyNow">立即兑换</button>
      <button v-if="isOwner" class="btn btn-secondary" @click="createListing">挂售</button>
    </view>
  </view>
</template>
```

### API调用
```javascript
// 获取藏品详情
async function fetchAssetDetail(assetId) {
  const res = await request.get(API_ENDPOINTS.ASSET.DETAIL.replace(':id', assetId))
  assetDetail.value = res.data
}
```

---

## 8. 鲸探作品

### 页面结构
```vue
<template>
  <view class="page">
    <!-- 顶部导航 -->
    <view class="navbar">
      <view class="back-btn" @click="goBack">
        <image src="/static/icons/back.png" />
      </view>
      <view class="navbar-title">鲸探作品</view>
    </view>

    <!-- 绑定提示 -->
    <view v-if="!isBound" class="bind-tip">
      <text>您还未绑定鲸探账户</text>
      <button class="btn btn-primary btn-small" @click="goToBind">去绑定</button>
    </view>

    <!-- 作品列表 -->
    <scroll-view v-else class="asset-list" scroll-y @scrolltolower="loadMore">
      <view class="list-wrapper">
        <view v-for="asset in jingtanAssets" :key="asset.id" class="asset-card">
          <image class="asset-image" :src="asset.image_url" mode="aspectFill" />
          <view class="asset-info">
            <text class="asset-name">{{ asset.name }}</text>
            <text class="asset-source">来自鲸探</text>
          </view>
          <view class="asset-tag">不可交易</view>
        </view>
      </view>
    </scroll-view>
  </view>
</template>
```

### API调用
```javascript
// 获取鲸探作品列表
async function fetchJingtanAssets() {
  const res = await request.get(API_ENDPOINTS.JINGTAN.ASSETS, {
    page: page.value,
    page_size: 20
  })
  jingtanAssets.value = [...jingtanAssets.value, ...res.data.list]
}
```

---

## 9. 第三方关联

### 设计稿参考
- 文件：`第三方关联.png`

### 页面结构
```vue
<template>
  <view class="page">
    <!-- 顶部导航 -->
    <view class="navbar">
      <view class="back-btn" @click="goBack">
        <image src="/static/icons/back.png" />
      </view>
      <view class="navbar-title">第三方关联</view>
    </view>

    <!-- 关联列表 -->
    <view class="bind-list">
      <!-- 鲸探 -->
      <view class="bind-item">
        <view class="bind-info">
          <image class="bind-logo" src="/static/images/jingtan-logo.png" />
          <view class="bind-text">
            <text class="bind-name">鲸探</text>
            <text class="bind-status" v-if="bindings.jingtan">已绑定</text>
            <text class="bind-status unbind" v-else>未绑定</text>
          </view>
        </view>
        <button v-if="!bindings.jingtan" class="btn btn-primary btn-small" @click="bindJingtan">绑定</button>
        <button v-else class="btn btn-secondary btn-small" @click="unbindJingtan">解绑</button>
      </view>

      <!-- Waveup -->
      <view class="bind-item">
        <view class="bind-info">
          <image class="bind-logo" src="/static/images/waveup-logo.png" />
          <view class="bind-text">
            <text class="bind-name">Waveup</text>
            <text class="bind-status unbind">未绑定</text>
          </view>
        </view>
        <button class="btn btn-primary btn-small" @click="showComingSoon">绑定</button>
      </view>
    </view>

    <!-- 绑定说明 -->
    <view class="bind-notice">
      <text class="notice-title">绑定说明</text>
      <text class="notice-text">1. 绑定第三方账户后，可以查看您在该平台的藏品</text>
      <text class="notice-text">2. 第三方藏品仅供展示，不可在本平台交易</text>
      <text class="notice-text">3. 绑定的手机号必须与小程序登录手机号一致</text>
    </view>
  </view>
</template>
```

### API调用
```javascript
// 绑定鲸探账户
async function bindJingtan() {
  // 弹出输入框，获取鲸探账户ID和手机号
  uni.showModal({
    title: '绑定鲸探账户',
    editable: true,
    placeholderText: '请输入鲸探账户ID',
    success: async (res) => {
      if (res.confirm) {
        try {
          await request.post(API_ENDPOINTS.JINGTAN.BIND, {
            jingtan_account_id: res.content,
            jingtan_phone: userInfo.value.phone  // 使用当前登录手机号
          })
          
          uni.showToast({
            title: '绑定成功',
            icon: 'success'
          })
          
          // 刷新绑定状态
          fetchBindings()
        } catch (error) {
          uni.showToast({
            title: error.message || '绑定失败',
            icon: 'none'
          })
        }
      }
    }
  })
}

// 解绑鲸探账户
async function unbindJingtan() {
  uni.showModal({
    title: '确认解绑',
    content: '解绑后将无法查看鲸探作品',
    success: async (res) => {
      if (res.confirm) {
        try {
          await request.delete(API_ENDPOINTS.JINGTAN.UNBIND)
          
          uni.showToast({
            title: '解绑成功',
            icon: 'success'
          })
          
          // 刷新绑定状态
          fetchBindings()
        } catch (error) {
          uni.showToast({
            title: error.message || '解绑失败',
            icon: 'none'
          })
        }
      }
    }
  })
}
```

---

## 10. 挂售页面

### 设计稿参考
- 文件：`挂售.png`

### 页面结构
```vue
<template>
  <view class="page">
    <!-- 顶部导航 -->
    <view class="navbar">
      <view class="back-btn" @click="goBack">
        <image src="/static/icons/back.png" />
      </view>
      <view class="navbar-title">挂售</view>
    </view>

    <!-- 藏品信息 -->
    <view class="asset-preview">
      <image class="asset-image" :src="assetInfo.image_url" mode="aspectFill" />
      <view class="asset-info">
        <text class="asset-name">{{ assetInfo.name }}</text>
        <text class="asset-serial">#{{ assetInfo.serial_number }}</text>
      </view>
    </view>

    <!-- 价格设置 -->
    <view class="price-setting">
      <view class="form-item">
        <text class="form-label">挂售价格</text>
        <view class="price-input-wrapper">
          <input 
            class="price-input number-display" 
            type="digit" 
            v-model="listingPrice" 
            placeholder="请输入价格"
          />
          <text class="price-unit">积分</text>
        </view>
      </view>
      
      <!-- 费用说明 -->
      <view class="fee-section">
        <view class="fee-item">
          <text class="fee-label">平台手续费（2.5%）</text>
          <text class="fee-value number-display">{{ calculateFee(listingPrice, 0.025) }}</text>
        </view>
        <view class="fee-item">
          <text class="fee-label">创作者版税（2.5%）</text>
          <text class="fee-value number-display">{{ calculateFee(listingPrice, 0.025) }}</text>
        </view>
        <view class="divider"></view>
        <view class="fee-item total">
          <text class="fee-label">您将获得</text>
          <text class="fee-value number-display">{{ calculateReceived(listingPrice) }}</text>
        </view>
      </view>
    </view>

    <!-- 提交按钮 -->
    <view class="action-section">
      <button class="btn btn-primary" @click="submitListing" :loading="submitting">确认挂售</button>
    </view>
  </view>
</template>
```

### API调用
```javascript
// 创建挂售单
async function submitListing() {
  if (!listingPrice.value || parseFloat(listingPrice.value) <= 0) {
    uni.showToast({
      title: '请输入有效价格',
      icon: 'none'
    })
    return
  }
  
  submitting.value = true
  try {
    const res = await request.post(API_ENDPOINTS.TRADE.CREATE_LISTING, {
      asset_instance_id: assetInfo.value.id,
      price: listingPrice.value
    })
    
    uni.showToast({
      title: '挂售成功',
      icon: 'success'
    })
    
    setTimeout(() => {
      uni.navigateBack()
    }, 1500)
  } catch (error) {
    uni.showToast({
      title: error.message || '挂售失败',
      icon: 'none'
    })
  } finally {
    submitting.value = false
  }
}
```

### 工具函数
```javascript
// 计算手续费
function calculateFee(price, rate) {
  if (!price) return '0.00000000'
  return (parseFloat(price) * rate).toFixed(8)
}

// 计算实际收入
function calculateReceived(price) {
  if (!price) return '0.00000000'
  const platformFee = parseFloat(price) * 0.025
  const royalty = parseFloat(price) * 0.025
  return (parseFloat(price) - platformFee - royalty).toFixed(8)
}
```

---

## 11. 兑换页面

### 设计稿参考
- 文件：`兑换页.png`、`兑换成功.png`

### 页面结构
```vue
<template>
  <view class="page">
    <!-- 顶部导航 -->
    <view class="navbar">
      <view class="back-btn" @click="goBack">
        <image src="/static/icons/back.png" />
      </view>
      <view class="navbar-title">兑换</view>
    </view>

    <!-- 藏品信息 -->
    <view class="asset-preview">
      <image class="asset-image" :src="listingInfo.asset.image_url" mode="aspectFill" />
      <view class="asset-info">
        <text class="asset-name">{{ listingInfo.asset.name }}</text>
        <text class="asset-serial">#{{ listingInfo.serial_number }}</text>
      </view>
    </view>

    <!-- 价格信息 -->
    <view class="price-info">
      <view class="price-row">
        <text class="price-label">兑换价格</text>
        <text class="price-value number-display">{{ formatPoints(listingInfo.price) }}</text>
      </view>
      <view class="price-row">
        <text class="price-label">我的积分</text>
        <text class="price-value number-display">{{ formatPoints(myPoints) }}</text>
      </view>
    </view>

    <!-- 卖家信息 -->
    <view class="seller-info">
      <text class="info-label">卖家</text>
      <view class="seller-detail">
        <image class="seller-avatar" :src="listingInfo.seller.avatar_url" />
        <text class="seller-name">{{ listingInfo.seller.nickname }}</text>
      </view>
    </view>

    <!-- 提交按钮 -->
    <view class="action-section">
      <button 
        class="btn btn-primary" 
        @click="confirmExchange" 
        :loading="exchanging"
        :disabled="myPoints < parseFloat(listingInfo.price)"
      >
        {{ myPoints < parseFloat(listingInfo.price) ? '积分不足' : '确认兑换' }}
      </button>
    </view>
  </view>
</template>
```

### API调用
```javascript
// 执行交易
async function confirmExchange() {
  uni.showModal({
    title: '确认兑换',
    content: `确认用 ${listingInfo.value.price} 积分兑换这个藏品吗？`,
    success: async (res) => {
      if (res.confirm) {
        exchanging.value = true
        try {
          const result = await request.post(API_ENDPOINTS.TRADE.EXECUTE_TRADE, {
            listing_id: listingInfo.value.id
          })
          
          // 显示成功页面
          uni.redirectTo({
            url: `/pages/exchange-success/index?tradeId=${result.data.trade_id}`
          })
        } catch (error) {
          uni.showToast({
            title: error.message || '兑换失败',
            icon: 'none'
          })
        } finally {
          exchanging.value = false
        }
      }
    }
  })
}
```

---

## 12. 登录页面

### 页面结构
```vue
<template>
  <view class="page">
    <!-- Logo -->
    <view class="logo-section">
      <image class="logo" src="/static/images/logo.png" mode="aspectFit" />
      <text class="app-name">HOHO</text>
    </view>

    <!-- 登录表单 -->
    <view class="form-section">
      <view class="form-item">
        <input 
          class="form-input" 
          type="number" 
          v-model="phone" 
          placeholder="请输入手机号"
          maxlength="11"
        />
      </view>
      
      <view class="form-item">
        <input 
          class="form-input" 
          type="text" 
          v-model="password" 
          password 
          placeholder="请输入密码"
        />
      </view>
      
      <button class="btn btn-primary" @click="handleLogin" :loading="logging">登录</button>
      
      <view class="form-footer">
        <text class="link-text" @click="goToRegister">还没有账号？去注册</text>
      </view>
    </view>
  </view>
</template>
```

### API调用
```javascript
// 登录
async function handleLogin() {
  if (!phone.value || phone.value.length !== 11) {
    uni.showToast({
      title: '请输入正确的手机号',
      icon: 'none'
    })
    return
  }
  
  if (!password.value) {
    uni.showToast({
      title: '请输入密码',
      icon: 'none'
    })
    return
  }
  
  logging.value = true
  try {
    const res = await request.post(API_ENDPOINTS.USER.LOGIN, {
      phone: phone.value,
      password: password.value
    })
    
    // 保存token
    uni.setStorageSync('token', res.data.token)
    uni.setStorageSync('userInfo', res.data.user)
    
    uni.showToast({
      title: '登录成功',
      icon: 'success'
    })
    
    setTimeout(() => {
      uni.switchTab({
        url: '/pages/index/index'
      })
    }, 1500)
  } catch (error) {
    uni.showToast({
      title: error.message || '登录失败',
      icon: 'none'
    })
  } finally {
    logging.value = false
  }
}
```

---

## 13. 注册页面

### 页面结构
```vue
<template>
  <view class="page">
    <!-- Logo -->
    <view class="logo-section">
      <image class="logo" src="/static/images/logo.png" mode="aspectFit" />
      <text class="app-name">HOHO</text>
    </view>

    <!-- 注册表单 -->
    <view class="form-section">
      <view class="form-item">
        <input 
          class="form-input" 
          type="number" 
          v-model="phone" 
          placeholder="请输入手机号"
          maxlength="11"
        />
      </view>
      
      <view class="form-item">
        <input 
          class="form-input" 
          type="text" 
          v-model="password" 
          password 
          placeholder="请设置密码（6-20位）"
        />
      </view>
      
      <view class="form-item">
        <input 
          class="form-input" 
          type="text" 
          v-model="confirmPassword" 
          password 
          placeholder="请再次输入密码"
        />
      </view>
      
      <button class="btn btn-primary" @click="handleRegister" :loading="registering">注册</button>
      
      <view class="form-footer">
        <text class="link-text" @click="goToLogin">已有账号？去登录</text>
      </view>
    </view>
  </view>
</template>
```

### API调用
```javascript
// 注册
async function handleRegister() {
  // 验证手机号
  if (!phone.value || phone.value.length !== 11) {
    uni.showToast({
      title: '请输入正确的手机号',
      icon: 'none'
    })
    return
  }
  
  // 验证密码
  if (!password.value || password.value.length < 6 || password.value.length > 20) {
    uni.showToast({
      title: '密码长度为6-20位',
      icon: 'none'
    })
    return
  }
  
  // 验证确认密码
  if (password.value !== confirmPassword.value) {
    uni.showToast({
      title: '两次密码不一致',
      icon: 'none'
    })
    return
  }
  
  registering.value = true
  try {
    const res = await request.post(API_ENDPOINTS.USER.REGISTER, {
      phone: phone.value,
      password: password.value
    })
    
    uni.showToast({
      title: '注册成功',
      icon: 'success'
    })
    
    setTimeout(() => {
      goToLogin()
    }, 1500)
  } catch (error) {
    uni.showToast({
      title: error.message || '注册失败',
      icon: 'none'
    })
  } finally {
    registering.value = false
  }
}
```

---

## 🧩 通用组件

### TabBar组件

创建文件：`/components/TabBar/TabBar.vue`

```vue
<template>
  <view class="tabbar safe-area-bottom">
    <view 
      class="tabbar-item" 
      :class="{ 'tabbar-active': active === 0 }"
      @click="switchTab('/pages/index/index')"
    >
      <image 
        class="tabbar-icon" 
        :src="active === 0 ? '/static/icons/tab-works-active.png' : '/static/icons/tab-works.png'" 
        mode="aspectFit" 
      />
      <text class="tabbar-label">作品</text>
    </view>
    
    <view 
      class="tabbar-item" 
      :class="{ 'tabbar-active': active === 1 }"
      @click="switchTab('/pages/create/index')"
    >
      <image 
        class="tabbar-icon" 
        :src="active === 1 ? '/static/icons/tab-create-active.png' : '/static/icons/tab-create.png'" 
        mode="aspectFit" 
      />
      <text class="tabbar-label">创作</text>
    </view>
    
    <view 
      class="tabbar-item" 
      :class="{ 'tabbar-active': active === 2 }"
      @click="switchTab('/pages/jijhuan/index')"
    >
      <image 
        class="tabbar-icon" 
        :src="active === 2 ? '/static/icons/tab-exchange-active.png' : '/static/icons/tab-exchange.png'" 
        mode="aspectFit" 
      />
      <text class="tabbar-label">集换</text>
    </view>
    
    <view 
      class="tabbar-item" 
      :class="{ 'tabbar-active': active === 3 }"
      @click="switchTab('/pages/ecology/index')"
    >
      <image 
        class="tabbar-icon" 
        :src="active === 3 ? '/static/icons/tab-ecology-active.png' : '/static/icons/tab-ecology.png'" 
        mode="aspectFit" 
      />
      <text class="tabbar-label">生态</text>
    </view>
    
    <view 
      class="tabbar-item" 
      :class="{ 'tabbar-active': active === 4 }"
      @click="switchTab('/pages/profile/index')"
    >
      <image 
        class="tabbar-icon" 
        :src="active === 4 ? '/static/icons/tab-profile-active.png' : '/static/icons/tab-profile.png'" 
        mode="aspectFit" 
      />
      <text class="tabbar-label">我的</text>
    </view>
  </view>
</template>

<script setup>
const props = defineProps({
  active: {
    type: Number,
    default: 0
  }
})

const switchTab = (url) => {
  uni.switchTab({ url })
}
</script>

<style lang="scss" scoped>
.tabbar {
  position: fixed;
  bottom: 0;
  left: 0;
  right: 0;
  display: flex;
  align-items: center;
  justify-content: space-around;
  height: 120rpx;
  background-color: var(--color-bg-primary);
  border-top: 1px solid var(--color-border);
  z-index: 1000;
  
  .tabbar-item {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    gap: 8rpx;
    flex: 1;
    height: 100%;
    color: var(--color-text-tertiary);
    
    .tabbar-icon {
      width: 48rpx;
      height: 48rpx;
    }
    
    .tabbar-label {
      font-size: var(--font-size-xs);
    }
    
    &.tabbar-active {
      color: var(--color-text-primary);
      font-weight: 500;
    }
  }
}
</style>
```

---

## 🔧 工具函数

### 积分格式化

创建文件：`/utils/format.js`

```javascript
/**
 * 格式化积分显示（8位小数）
 * @param {string|number} points - 积分值
 * @param {number} decimals - 小数位数，默认8位
 * @returns {string} 格式化后的积分
 */
export function formatPoints(points, decimals = 8) {
  if (points === null || points === undefined) return '0.00000000'
  return parseFloat(points).toFixed(decimals)
}

/**
 * 手机号脱敏
 * @param {string} phone - 手机号
 * @returns {string} 脱敏后的手机号
 */
export function maskPhone(phone) {
  if (!phone) return '未绑定'
  return phone.replace(/(\d{3})\d{4}(\d{4})/, '$1****$2')
}

/**
 * 时间格式化
 * @param {string|number} timestamp - 时间戳
 * @returns {string} 格式化后的时间
 */
export function formatTime(timestamp) {
  const date = new Date(timestamp)
  const year = date.getFullYear()
  const month = String(date.getMonth() + 1).padStart(2, '0')
  const day = String(date.getDate()).padStart(2, '0')
  const hour = String(date.getHours()).padStart(2, '0')
  const minute = String(date.getMinutes()).padStart(2, '0')
  return `${year}-${month}-${day} ${hour}:${minute}`
}

/**
 * 短Hash
 * @param {string} hash - 完整Hash
 * @returns {string} 短Hash
 */
export function shortHash(hash) {
  if (!hash) return '-'
  return `${hash.slice(0, 6)}...${hash.slice(-4)}`
}
```

---

## 🔐 认证拦截

### HTTP请求拦截器

更新文件：`/api/request.js`

```javascript
import { API_BASE_URL } from './config'

// 获取token
function getToken() {
  return uni.getStorageSync('token') || ''
}

// 请求拦截
function request(options) {
  return new Promise((resolve, reject) => {
    uni.request({
      url: `${API_BASE_URL}${options.url}`,
      method: options.method || 'GET',
      data: options.data || {},
      header: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${getToken()}`,
        ...options.header
      },
      success: (res) => {
        if (res.statusCode === 200) {
          if (res.data.code === 0) {
            resolve(res.data)
          } else {
            reject(new Error(res.data.message || '请求失败'))
          }
        } else if (res.statusCode === 401) {
          // 未登录，跳转到登录页
          uni.removeStorageSync('token')
          uni.removeStorageSync('userInfo')
          uni.reLaunch({
            url: '/pages/login/index'
          })
          reject(new Error('请先登录'))
        } else {
          reject(new Error(res.data.message || '请求失败'))
        }
      },
      fail: (error) => {
        reject(error)
      }
    })
  })
}

// GET请求
export function get(url, params) {
  return request({
    url: params ? `${url}?${new URLSearchParams(params).toString()}` : url,
    method: 'GET'
  })
}

// POST请求
export function post(url, data) {
  return request({
    url,
    method: 'POST',
    data
  })
}

// PUT请求
export function put(url, data) {
  return request({
    url,
    method: 'PUT',
    data
  })
}

// DELETE请求
export function del(url) {
  return request({
    url,
    method: 'DELETE'
  })
}

export default {
  get,
  post,
  put,
  delete: del
}
```

---

## 📦 pages.json配置

更新文件：`/pages.json`

```json
{
  "pages": [
    {
      "path": "pages/index/index",
      "style": {
        "navigationBarTitleText": "HOHO",
        "navigationStyle": "custom"
      }
    },
    {
      "path": "pages/jijhuan/index",
      "style": {
        "navigationBarTitleText": "集换中心",
        "navigationStyle": "custom"
      }
    },
    {
      "path": "pages/profile/index",
      "style": {
        "navigationBarTitleText": "我的",
        "navigationStyle": "custom"
      }
    },
    {
      "path": "pages/ecology/index",
      "style": {
        "navigationBarTitleText": "生态",
        "navigationStyle": "custom"
      }
    },
    {
      "path": "pages/transparent-ledger/index",
      "style": {
        "navigationBarTitleText": "透明公示"
      }
    },
    {
      "path": "pages/create/index",
      "style": {
        "navigationBarTitleText": "创作"
      }
    },
    {
      "path": "pages/asset-detail/index",
      "style": {
        "navigationBarTitleText": "藏品详情"
      }
    },
    {
      "path": "pages/jingtan-assets/index",
      "style": {
        "navigationBarTitleText": "鲸探作品"
      }
    },
    {
      "path": "pages/third-party/index",
      "style": {
        "navigationBarTitleText": "第三方关联"
      }
    },
    {
      "path": "pages/listing-create/index",
      "style": {
        "navigationBarTitleText": "挂售"
      }
    },
    {
      "path": "pages/exchange/index",
      "style": {
        "navigationBarTitleText": "兑换"
      }
    },
    {
      "path": "pages/login/index",
      "style": {
        "navigationBarTitleText": "登录",
        "navigationStyle": "custom"
      }
    },
    {
      "path": "pages/register/index",
      "style": {
        "navigationBarTitleText": "注册",
        "navigationStyle": "custom"
      }
    }
  ],
  "tabBar": {
    "custom": true,
    "list": [
      {
        "pagePath": "pages/index/index",
        "text": "作品"
      },
      {
        "pagePath": "pages/create/index",
        "text": "创作"
      },
      {
        "pagePath": "pages/jijhuan/index",
        "text": "集换"
      },
      {
        "pagePath": "pages/ecology/index",
        "text": "生态"
      },
      {
        "pagePath": "pages/profile/index",
        "text": "我的"
      }
    ]
  },
  "globalStyle": {
    "navigationBarTextStyle": "black",
    "navigationBarTitleText": "HOHO",
    "navigationBarBackgroundColor": "#FFFFFF",
    "backgroundColor": "#F5F5F5"
  }
}
```

---

## 🎨 静态资源

### 需要准备的图标和图片

将以下文件放置在 `/static/` 目录下：

#### 图标（/static/icons/）
- `search.png` - 搜索图标
- `filter.png` - 筛选图标
- `back.png` - 返回图标
- `more.png` - 更多图标
- `scan.png` - 扫码图标
- `arrow-right.png` - 右箭头
- `official.png` - 官方标签图标
- `clock.png` - 时钟图标
- `total.png` - 总量图标
- `upload.png` - 上传图标
- `share.png` - 分享图标
- `diamond.png` - 钻石图标（任务中心）
- `collection.png` - 作品集图标
- `shop.png` - 周边图标
- `service.png` - 客服图标
- `settings.png` - 设置图标
- `exchange-history.png` - 集换记录图标
- `community.png` - 社区作品图标
- `jingtan.png` - 鲸探图标
- `waveup.png` - Waveup图标
- `whitepaper.png` - 白皮书图标
- `governance.png` - 社区治理图标
- `transparent.png` - 透明公示图标
- `ad.png` - 广告图标

#### TabBar图标（/static/icons/）
- `tab-works.png` / `tab-works-active.png` - 作品Tab
- `tab-create.png` / `tab-create-active.png` - 创作Tab
- `tab-exchange.png` / `tab-exchange-active.png` - 集换Tab
- `tab-ecology.png` / `tab-ecology-active.png` - 生态Tab
- `tab-profile.png` / `tab-profile-active.png` - 我的Tab

#### 图片（/static/images/）
- `logo.png` - App Logo
- `banner-welcome.png` - 首页Banner（欢迎来到HOHO Park!）
- `ecology-banner.png` - 生态页面Banner
- `ecology-function.png` - 生态页面功能区左侧图片
- `default-avatar.png` - 默认头像
- `placeholder.png` - 占位图
- `partner-jingtan.png` - 鲸探Logo
- `partner-waveup.png` - Waveup Logo
- `partner-xmeta.png` - XMeta Logo
- `partner-pinpai.png` - 品拍Logo
- `partner-coming.png` - 敬请期待Logo
- `jingtan-logo.png` - 鲸探大Logo（第三方关联页面）
- `waveup-logo.png` - Waveup大Logo（第三方关联页面）

---

## 🚀 部署说明

### 1. 修改API地址

在 `/api/config.js` 中，将 `API_BASE_URL` 修改为你的实际服务器地址：

```javascript
export const API_BASE_URL = 'http://your-server-ip:8080'  // 替换为实际地址
```

### 2. 编译小程序

使用HBuilderX或命令行编译：

```bash
# 微信小程序
npm run build:mp-weixin

# 支付宝小程序
npm run build:mp-alipay

# 抖音小程序
npm run build:mp-toutiao
```

### 3. 上传代码

将编译后的代码上传到对应的小程序开发者工具，然后提交审核。

---

## ✅ 开发检查清单

- [ ] 所有页面已创建
- [ ] TabBar组件已实现
- [ ] API对接已完成
- [ ] 样式符合设计稿
- [ ] 登录/注册功能正常
- [ ] 积分显示精确到8位小数
- [ ] 图片上传功能正常
- [ ] 鲸探API对接正常
- [ ] 交易流程完整
- [ ] 错误处理完善
- [ ] 加载状态显示
- [ ] 空状态处理
- [ ] 静态资源准备完毕
- [ ] API地址已配置
- [ ] 编译测试通过

---

## 📞 联系方式

如有问题，请查看GitHub仓库：https://github.com/dokidwell/hoho-miniapp

---

**祝开发顺利！🎉**
