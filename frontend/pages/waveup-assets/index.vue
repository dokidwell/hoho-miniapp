<template>
  <view class="waveup-assets-page">
    <!-- 自定义导航栏 -->
    <view class="custom-navbar">
      <view class="navbar-content">
        <view class="back-btn" @click="goBack">
          <text class="back-icon">←</text>
        </view>
        <view class="navbar-title">
          <image class="logo" src="/static/waveup-logo.png" mode="aspectFit" />
          <text class="title-text">Waveup作品</text>
        </view>
        <view class="placeholder"></view>
      </view>
    </view>

    <!-- Banner -->
    <view class="banner">
      <view class="banner-content">
        <text class="banner-title">🌊 Waveup联名系列</text>
        <text class="banner-subtitle">限量发行 · 独家合作</text>
      </view>
    </view>

    <!-- 筛选栏 -->
    <view class="filter-bar">
      <view 
        v-for="filter in filters" 
        :key="filter.value"
        class="filter-item"
        :class="{ active: currentFilter === filter.value }"
        @click="selectFilter(filter.value)"
      >
        <text class="filter-text">{{ filter.label }}</text>
      </view>
    </view>

    <!-- 作品列表 -->
    <scroll-view 
      class="assets-scroll"
      scroll-y
      @scrolltolower="loadMore"
      refresher-enabled
      :refresher-triggered="refreshing"
      @refresherrefresh="onRefresh"
    >
      <view class="assets-list">
        <view 
          v-for="asset in assetList" 
          :key="asset.id"
          class="asset-item"
          @click="viewDetail(asset.id)"
        >
          <image 
            class="asset-image" 
            :src="asset.imageUrl"
            mode="aspectFill"
          />
          <view class="asset-content">
            <view class="asset-header">
              <text class="asset-name">{{ asset.name }}</text>
              <view class="asset-badge" v-if="asset.isLimited">
                <text class="badge-text">限量</text>
              </view>
            </view>
            <text class="asset-desc">{{ asset.description }}</text>
            <view class="asset-footer">
              <view class="supply-info">
                <text class="supply-text">发行量：{{ asset.totalSupply }}</text>
                <text class="minted-text">已铸造：{{ asset.mintedCount }}</text>
              </view>
              <view class="price-info" v-if="asset.price">
                <text class="price-label">价格</text>
                <text class="price-value">{{ asset.price }} 积分</text>
              </view>
            </view>
            <view class="progress-bar">
              <view 
                class="progress-fill" 
                :style="{ width: (asset.mintedCount / asset.totalSupply * 100) + '%' }"
              ></view>
            </view>
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
        <text class="empty-icon">🌊</text>
        <text class="empty-text">暂无Waveup作品</text>
        <text class="empty-hint">敬请期待更多联名系列</text>
      </view>
    </scroll-view>

    <!-- 关于Waveup -->
    <view class="about-fab" @click="showAbout">
      <text class="fab-icon">ℹ️</text>
    </view>
  </view>
</template>

<script setup>
import { ref, onMounted } from 'vue'

// 筛选选项
const filters = ref([
  { label: '全部', value: 'all' },
  { label: '可铸造', value: 'available' },
  { label: '已售罄', value: 'soldout' },
  { label: '限量版', value: 'limited' }
])

const currentFilter = ref('all')
const assetList = ref([])
const loading = ref(false)
const refreshing = ref(false)
const noMore = ref(false)
const page = ref(1)
const pageSize = 10

onMounted(() => {
  loadAssets()
})

// 选择筛选
const selectFilter = (value) => {
  currentFilter.value = value
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
    // TODO: 调用API获取Waveup作品列表
    // const res = await uni.$api.assets.getWaveupList({
    //   filter: currentFilter.value,
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
  const count = pageNum === 1 ? pageSize : Math.min(pageSize, 3)
  
  for (let i = 0; i < count; i++) {
    const id = (pageNum - 1) * pageSize + i + 1
    const totalSupply = [100, 500, 1000, 2000][Math.floor(Math.random() * 4)]
    const mintedCount = Math.floor(totalSupply * Math.random())
    
    assets.push({
      id: id,
      name: `Waveup联名作品 #${id}`,
      description: '与Waveup独家合作的限量数字藏品，每一件都是独一无二的艺术品',
      imageUrl: '/static/placeholder.png',
      totalSupply: totalSupply,
      mintedCount: mintedCount,
      price: [100, 200, 500, 1000][Math.floor(Math.random() * 4)],
      isLimited: totalSupply <= 500
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

// 显示关于
const showAbout = () => {
  uni.showModal({
    title: '关于Waveup',
    content: 'Waveup是一个专注于数字艺术的创作平台，与HOHO Park深度合作，为用户带来更多优质的数字藏品。',
    confirmText: '了解更多',
    success: (res) => {
      if (res.confirm) {
        // TODO: 跳转到Waveup介绍页面
      }
    }
  })
}

// 返回
const goBack = () => {
  uni.navigateBack()
}
</script>

<style scoped>
.waveup-assets-page {
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
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  z-index: 1000;
  padding-top: env(safe-area-inset-top);
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
  color: white;
}

.navbar-title {
  display: flex;
  align-items: center;
  gap: 15rpx;
}

.logo {
  width: 50rpx;
  height: 50rpx;
  border-radius: 10rpx;
  background: white;
}

.title-text {
  font-size: 32rpx;
  font-weight: bold;
  color: white;
}

.placeholder {
  width: 60rpx;
}

/* Banner */
.banner {
  position: fixed;
  top: calc(88rpx + env(safe-area-inset-top));
  left: 0;
  right: 0;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  z-index: 999;
}

.banner-content {
  padding: 30rpx;
  text-align: center;
}

.banner-title {
  display: block;
  font-size: 36rpx;
  font-weight: bold;
  color: white;
  margin-bottom: 10rpx;
}

.banner-subtitle {
  display: block;
  font-size: 24rpx;
  color: rgba(255, 255, 255, 0.8);
}

/* 筛选栏 */
.filter-bar {
  position: fixed;
  top: calc(178rpx + env(safe-area-inset-top));
  left: 0;
  right: 0;
  background: white;
  z-index: 998;
  display: flex;
  padding: 20rpx 30rpx;
  gap: 20rpx;
  box-shadow: 0 2rpx 10rpx rgba(0, 0, 0, 0.05);
}

.filter-item {
  padding: 12rpx 30rpx;
  border-radius: 30rpx;
  background: #f5f5f5;
  transition: all 0.3s;
}

.filter-item.active {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
}

.filter-text {
  font-size: 26rpx;
  color: #666;
}

.filter-item.active .filter-text {
  color: white;
  font-weight: bold;
}

/* 作品列表 */
.assets-scroll {
  height: 100vh;
  padding-top: calc(258rpx + env(safe-area-inset-top));
  padding-bottom: 120rpx;
}

.assets-list {
  padding: 20rpx;
}

.asset-item {
  background: white;
  border-radius: 20rpx;
  overflow: hidden;
  margin-bottom: 20rpx;
  box-shadow: 0 4rpx 20rpx rgba(0, 0, 0, 0.08);
  display: flex;
}

.asset-image {
  width: 240rpx;
  height: 240rpx;
  background: #f5f5f5;
  flex-shrink: 0;
}

.asset-content {
  flex: 1;
  padding: 20rpx;
  display: flex;
  flex-direction: column;
}

.asset-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-bottom: 10rpx;
}

.asset-name {
  font-size: 30rpx;
  font-weight: bold;
  color: #333;
  flex: 1;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.asset-badge {
  background: linear-gradient(135deg, #ff6b6b 0%, #ee5a6f 100%);
  padding: 6rpx 15rpx;
  border-radius: 15rpx;
  flex-shrink: 0;
}

.badge-text {
  font-size: 20rpx;
  color: white;
  font-weight: bold;
}

.asset-desc {
  font-size: 24rpx;
  color: #999;
  line-height: 36rpx;
  margin-bottom: 15rpx;
  overflow: hidden;
  text-overflow: ellipsis;
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
}

.asset-footer {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 15rpx;
}

.supply-info {
  display: flex;
  gap: 20rpx;
}

.supply-text,
.minted-text {
  font-size: 22rpx;
  color: #999;
}

.price-info {
  display: flex;
  align-items: baseline;
  gap: 10rpx;
}

.price-label {
  font-size: 22rpx;
  color: #999;
}

.price-value {
  font-size: 28rpx;
  color: #667eea;
  font-weight: bold;
}

.progress-bar {
  height: 8rpx;
  background: #f0f0f0;
  border-radius: 4rpx;
  overflow: hidden;
}

.progress-fill {
  height: 100%;
  background: linear-gradient(90deg, #667eea 0%, #764ba2 100%);
  border-radius: 4rpx;
  transition: width 0.3s;
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

/* 关于按钮 */
.about-fab {
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
  z-index: 997;
}

.fab-icon {
  font-size: 50rpx;
}
</style>
