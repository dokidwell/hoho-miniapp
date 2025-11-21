<template>
  <view class="community-assets-page">
    <!-- 自定义导航栏 -->
    <view class="custom-navbar">
      <view class="navbar-content">
        <view class="back-btn" @click="goBack">
          <text class="back-icon">←</text>
        </view>
        <view class="navbar-title">社区作品</view>
        <view class="placeholder"></view>
      </view>
    </view>

    <!-- 分类标签 -->
    <scroll-view class="category-tabs" scroll-x>
      <view class="tabs-content">
        <view 
          v-for="(category, index) in categories" 
          :key="index"
          class="tab-item"
          :class="{ active: currentCategory === category.value }"
          @click="selectCategory(category.value)"
        >
          <text class="tab-text">{{ category.label }}</text>
        </view>
      </view>
    </scroll-view>

    <!-- 作品列表 -->
    <scroll-view 
      class="assets-scroll"
      scroll-y
      @scrolltolower="loadMore"
      refresher-enabled
      :refresher-triggered="refreshing"
      @refresherrefresh="onRefresh"
    >
      <view class="assets-grid">
        <view 
          v-for="asset in assetList" 
          :key="asset.id"
          class="asset-card"
          @click="viewDetail(asset.id)"
        >
          <image 
            class="asset-image" 
            :src="asset.imageUrl"
            mode="aspectFill"
          />
          <view class="asset-info">
            <text class="asset-name">{{ asset.name }}</text>
            <view class="asset-meta">
              <view class="creator">
                <image class="creator-avatar" :src="asset.creatorAvatar" />
                <text class="creator-name">{{ asset.creatorName }}</text>
              </view>
              <view class="stats">
                <text class="stat-item">❤️ {{ asset.likes }}</text>
              </view>
            </view>
          </view>
          <view class="asset-badge" v-if="asset.isNew">
            <text class="badge-text">NEW</text>
          </view>
        </view>
      </view>

      <!-- 加载状态 -->
      <view class="loading-status">
        <text v-if="loading" class="loading-text">加载中...</text>
        <text v-else-if="noMore" class="loading-text">没有更多了</text>
      </view>

      <!-- 空状态 -->
      <view v-if="assetList.length === 0 && !loading" class="empty-state">
        <text class="empty-icon">📦</text>
        <text class="empty-text">暂无社区作品</text>
        <text class="empty-hint">快来创作第一件作品吧！</text>
      </view>
    </scroll-view>

    <!-- 创作按钮 -->
    <view class="create-fab" @click="goCreate">
      <text class="fab-icon">+</text>
    </view>
  </view>
</template>

<script setup>
import { ref, onMounted } from 'vue'

// 分类列表
const categories = ref([
  { label: '全部', value: 'all' },
  { label: '艺术', value: 'art' },
  { label: '摄影', value: 'photo' },
  { label: '插画', value: 'illustration' },
  { label: '3D', value: '3d' },
  { label: '音乐', value: 'music' },
  { label: '视频', value: 'video' }
])

const currentCategory = ref('all')
const assetList = ref([])
const loading = ref(false)
const refreshing = ref(false)
const noMore = ref(false)
const page = ref(1)
const pageSize = 20

onMounted(() => {
  loadAssets()
})

// 选择分类
const selectCategory = (value) => {
  currentCategory.value = value
  page.value = 1
  assetList.value = []
  noMore.value = false
  loadAssets()
}

// 加载作品列表
const loadAssets = async () => {
  if (loading.value || noMore.value) return
  
  loading.value = true
  
  try {
    // TODO: 调用API获取作品列表
    // const res = await uni.$api.assets.getCommunityList({
    //   category: currentCategory.value,
    //   page: page.value,
    //   pageSize: pageSize
    // })
    
    // 临时模拟数据
    await new Promise(resolve => setTimeout(resolve, 1000))
    
    const mockData = generateMockAssets(page.value)
    
    if (mockData.length < pageSize) {
      noMore.value = true
    }
    
    assetList.value = [...assetList.value, ...mockData]
    page.value++
    
  } catch (error) {
    console.error('加载作品列表失败:', error)
    uni.showToast({
      title: '加载失败',
      icon: 'none'
    })
  } finally {
    loading.value = false
    refreshing.value = false
  }
}

// 生成模拟数据
const generateMockAssets = (pageNum) => {
  const assets = []
  const count = pageNum === 1 ? pageSize : Math.min(pageSize, 5)
  
  for (let i = 0; i < count; i++) {
    const id = (pageNum - 1) * pageSize + i + 1
    assets.push({
      id: id,
      name: `社区作品 #${id}`,
      imageUrl: '/static/placeholder.png',
      creatorName: `创作者${id}`,
      creatorAvatar: '/static/avatar.png',
      likes: Math.floor(Math.random() * 1000),
      isNew: Math.random() > 0.7
    })
  }
  
  return assets
}

// 下拉刷新
const onRefresh = () => {
  refreshing.value = true
  page.value = 1
  assetList.value = []
  noMore.value = false
  loadAssets()
}

// 加载更多
const loadMore = () => {
  loadAssets()
}

// 查看详情
const viewDetail = (id) => {
  uni.navigateTo({
    url: `/pages/asset-detail/index?id=${id}`
  })
}

// 去创作
const goCreate = () => {
  uni.switchTab({
    url: '/pages/create/index'
  })
}

// 返回
const goBack = () => {
  uni.navigateBack()
}
</script>

<style scoped>
.community-assets-page {
  width: 100%;
  height: 100vh;
  background: #f5f5f5;
}

/* 自定义导航栏 */
.custom-navbar {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  background: white;
  z-index: 1000;
  padding-top: env(safe-area-inset-top);
  box-shadow: 0 2rpx 10rpx rgba(0, 0, 0, 0.05);
}

.navbar-content {
  display: flex;
  align-items: center;
  justify-content: space-between;
  height: 88rpx;
  padding: 0 30rpx;
}

.back-btn {
  width: 60rpx;
  height: 60rpx;
  display: flex;
  align-items: center;
  justify-content: center;
}

.back-icon {
  font-size: 40rpx;
  color: #333;
}

.navbar-title {
  font-size: 32rpx;
  font-weight: bold;
  color: #333;
}

.placeholder {
  width: 60rpx;
}

/* 分类标签 */
.category-tabs {
  position: fixed;
  top: calc(88rpx + env(safe-area-inset-top));
  left: 0;
  right: 0;
  background: white;
  z-index: 999;
  white-space: nowrap;
  border-bottom: 2rpx solid #f0f0f0;
}

.tabs-content {
  display: inline-flex;
  padding: 20rpx 30rpx;
  gap: 30rpx;
}

.tab-item {
  padding: 12rpx 30rpx;
  border-radius: 30rpx;
  background: #f5f5f5;
  transition: all 0.3s;
}

.tab-item.active {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
}

.tab-text {
  font-size: 28rpx;
  color: #666;
  white-space: nowrap;
}

.tab-item.active .tab-text {
  color: white;
  font-weight: bold;
}

/* 作品列表 */
.assets-scroll {
  height: 100vh;
  padding-top: calc(148rpx + env(safe-area-inset-top));
  padding-bottom: 120rpx;
}

.assets-grid {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 20rpx;
  padding: 20rpx;
}

.asset-card {
  background: white;
  border-radius: 15rpx;
  overflow: hidden;
  box-shadow: 0 4rpx 20rpx rgba(0, 0, 0, 0.08);
  position: relative;
}

.asset-image {
  width: 100%;
  height: 340rpx;
  background: #f5f5f5;
}

.asset-info {
  padding: 20rpx;
}

.asset-name {
  display: block;
  font-size: 28rpx;
  font-weight: bold;
  color: #333;
  margin-bottom: 15rpx;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.asset-meta {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.creator {
  display: flex;
  align-items: center;
  gap: 10rpx;
  flex: 1;
  overflow: hidden;
}

.creator-avatar {
  width: 40rpx;
  height: 40rpx;
  border-radius: 50%;
  background: #f5f5f5;
  flex-shrink: 0;
}

.creator-name {
  font-size: 24rpx;
  color: #999;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.stats {
  flex-shrink: 0;
}

.stat-item {
  font-size: 24rpx;
  color: #999;
}

.asset-badge {
  position: absolute;
  top: 20rpx;
  right: 20rpx;
  background: linear-gradient(135deg, #ff6b6b 0%, #ee5a6f 100%);
  padding: 8rpx 20rpx;
  border-radius: 20rpx;
  box-shadow: 0 4rpx 15rpx rgba(255, 107, 107, 0.4);
}

.badge-text {
  font-size: 20rpx;
  color: white;
  font-weight: bold;
}

/* 加载状态 */
.loading-status {
  text-align: center;
  padding: 40rpx 0;
}

.loading-text {
  font-size: 24rpx;
  color: #999;
}

/* 空状态 */
.empty-state {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 200rpx 60rpx;
}

.empty-icon {
  font-size: 120rpx;
  margin-bottom: 30rpx;
}

.empty-text {
  font-size: 32rpx;
  color: #999;
  margin-bottom: 15rpx;
}

.empty-hint {
  font-size: 26rpx;
  color: #ccc;
}

/* 创作按钮 */
.create-fab {
  position: fixed;
  right: 30rpx;
  bottom: calc(120rpx + env(safe-area-inset-bottom));
  width: 100rpx;
  height: 100rpx;
  border-radius: 50%;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  display: flex;
  align-items: center;
  justify-content: center;
  box-shadow: 0 10rpx 40rpx rgba(102, 126, 234, 0.4);
  z-index: 998;
}

.fab-icon {
  font-size: 60rpx;
  color: white;
  font-weight: 300;
}
</style>
