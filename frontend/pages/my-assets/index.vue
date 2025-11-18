<template>
  <view class="page">
    <!-- 顶部导航 -->
    <view class="navbar">
      <text class="navbar-title">我的作品</text>
    </view>

    <!-- 筛选标签 -->
    <view class="filter-tabs">
      <view 
        v-for="tab in tabs" 
        :key="tab.id"
        class="tab-item"
        :class="{ active: currentTab === tab.id }"
        @click="switchTab(tab.id)"
      >
        {{ tab.name }}
      </view>
    </view>

    <!-- 作品列表 -->
    <view class="assets-list">
      <view 
        v-for="asset in filteredAssets" 
        :key="asset.id"
        class="asset-item"
        @click="goToDetail(asset.id)"
      >
        <image 
          class="asset-image" 
          :src="asset.image_url || '/static/images/placeholder.png'" 
          mode="aspectFill"
        />
        <view class="asset-info">
          <text class="asset-name">{{ asset.name }}</text>
          <text class="asset-number">#{{ asset.serial_number }}</text>
          <view class="asset-status">
            <text v-if="asset.is_listed" class="status-tag listed">已挂售</text>
            <text v-else class="status-tag">可转赠</text>
          </view>
        </view>
      </view>
    </view>

    <!-- 加载状态 -->
    <view v-if="loading" class="loading-wrapper">
      <text class="loading-text">加载中...</text>
    </view>

    <!-- 空状态 -->
    <view v-else-if="filteredAssets.length === 0" class="empty-wrapper">
      <text class="empty-emoji">📦</text>
      <text class="empty-text">暂无作品</text>
      <text class="empty-hint">去创作或兑换作品吧</text>
    </view>
  </view>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import request from '@/api/request'
import { API_ENDPOINTS } from '@/api/config'

const tabs = [
  { id: 'all', name: '全部' },
  { id: 'available', name: '可转赠' },
  { id: 'listed', name: '已挂售' }
]

const currentTab = ref('all')
const assets = ref([])
const loading = ref(false)

// 筛选后的作品列表
const filteredAssets = computed(() => {
  if (currentTab.value === 'all') {
    return assets.value
  } else if (currentTab.value === 'available') {
    return assets.value.filter(item => !item.is_listed)
  } else if (currentTab.value === 'listed') {
    return assets.value.filter(item => item.is_listed)
  }
  return assets.value
})

onMounted(() => {
  fetchMyAssets()
})

// 获取我的作品
async function fetchMyAssets() {
  loading.value = true
  try {
    const res = await request.get(API_ENDPOINTS.ASSET.MY_ASSETS)
    assets.value = res.list || []
  } catch (error) {
    console.error('获取作品失败:', error)
    uni.showToast({
      title: '获取作品失败',
      icon: 'none'
    })
  } finally {
    loading.value = false
  }
}

// 切换标签
function switchTab(tabId) {
  currentTab.value = tabId
}

// 跳转到作品详情
function goToDetail(assetId) {
  uni.navigateTo({
    url: `/pages/asset-detail/index?id=${assetId}`
  })
}
</script>

<style lang="scss" scoped>
.page {
  min-height: 100vh;
  background-color: #F5F5F5;
}

.navbar {
  height: 88rpx;
  background-color: #FFFFFF;
  display: flex;
  align-items: center;
  justify-content: center;
  border-bottom: 1rpx solid #F0F0F0;
  
  .navbar-title {
    font-size: 32rpx;
    font-weight: 600;
    color: #000000;
  }
}

.filter-tabs {
  display: flex;
  background-color: #FFFFFF;
  padding: 24rpx 32rpx;
  gap: 24rpx;
  
  .tab-item {
    flex: 1;
    height: 64rpx;
    display: flex;
    align-items: center;
    justify-content: center;
    background-color: #F5F5F5;
    border-radius: 32rpx;
    font-size: 28rpx;
    color: #666666;
    transition: all 0.3s;
    
    &.active {
      background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
      color: #FFFFFF;
      font-weight: 600;
    }
  }
}

.assets-list {
  padding: 24rpx 32rpx;
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 24rpx;
}

.asset-item {
  background-color: #FFFFFF;
  border-radius: 16rpx;
  overflow: hidden;
  box-shadow: 0 4rpx 12rpx rgba(0, 0, 0, 0.05);
  
  .asset-image {
    width: 100%;
    height: 320rpx;
    background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  }
  
  .asset-info {
    padding: 24rpx;
    
    .asset-name {
      font-size: 28rpx;
      font-weight: 600;
      color: #000000;
      display: block;
      margin-bottom: 8rpx;
      overflow: hidden;
      text-overflow: ellipsis;
      white-space: nowrap;
    }
    
    .asset-number {
      font-size: 24rpx;
      color: #999999;
      display: block;
      margin-bottom: 12rpx;
    }
    
    .asset-status {
      .status-tag {
        display: inline-block;
        padding: 6rpx 16rpx;
        background-color: #E8F5E9;
        color: #4CAF50;
        font-size: 22rpx;
        border-radius: 6rpx;
        
        &.listed {
          background-color: #FFF3E0;
          color: #FF9800;
        }
      }
    }
  }
}

.loading-wrapper {
  padding: 120rpx 0;
  text-align: center;
  
  .loading-text {
    font-size: 28rpx;
    color: #999999;
  }
}

.empty-wrapper {
  padding: 200rpx 0;
  text-align: center;
  
  .empty-emoji {
    font-size: 120rpx;
    display: block;
    margin-bottom: 24rpx;
  }
  
  .empty-text {
    font-size: 32rpx;
    color: #666666;
    display: block;
    margin-bottom: 12rpx;
  }
  
  .empty-hint {
    font-size: 24rpx;
    color: #999999;
    display: block;
  }
}
</style>
