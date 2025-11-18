<template>
  <view class="page">
    <!-- 顶部导航 -->
    <view class="navbar">
      <view class="back-btn" @click="goBack">
        <text class="back-icon">←</text>
      </view>
      <view class="navbar-title">鲸探作品</view>
      <view class="placeholder"></view>
    </view>

    <!-- 绑定提示 -->
    <view v-if="!isBound" class="bind-tip">
      <text class="tip-emoji">🔗</text>
      <text class="tip-text">您还未绑定鲸探账户</text>
      <button class="btn btn-primary btn-small" @click="goToBind">去绑定</button>
    </view>

    <!-- 作品列表 -->
    <view v-else class="asset-list-wrapper">
      <!-- 加载中 -->
      <view v-if="loading" class="loading-wrapper">
        <text class="loading-text">加载中...</text>
      </view>

      <!-- 作品列表 -->
      <view v-else-if="jingtanAssets.length > 0" class="asset-list">
        <view v-for="asset in jingtanAssets" :key="asset.id" class="asset-card">
          <view class="asset-image">
            <text class="asset-emoji">🖼️</text>
            <view class="asset-tag">不可交易</view>
          </view>
          <view class="asset-info">
            <text class="asset-name">{{ asset.name }}</text>
            <text class="asset-source">来自鲸探</text>
          </view>
        </view>
      </view>

      <!-- 空状态 -->
      <view v-else class="empty-wrapper">
        <text class="empty-emoji">📭</text>
        <text class="empty-text">暂无鲸探作品</text>
      </view>
    </view>
  </view>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import request from '@/api/request'
import { API_ENDPOINTS } from '@/api/config'

const loading = ref(false)
const isBound = ref(false)
const jingtanAssets = ref([])
const page = ref(1)
const pageSize = ref(20)

onMounted(() => {
  checkBinding()
})

// 检查绑定状态
async function checkBinding() {
  try {
    const res = await request.get(API_ENDPOINTS.THIRD_PARTY.LIST)
    isBound.value = res.some(item => item.platform === 'jingtan')
    
    if (isBound.value) {
      fetchJingtanAssets()
    }
  } catch (error) {
    console.error('检查绑定状态失败:', error)
  }
}

// 获取鲸探作品列表
async function fetchJingtanAssets() {
  loading.value = true
  try {
    const res = await request.get(API_ENDPOINTS.THIRD_PARTY.GET_JINGTAN_ASSETS, {
      page: page.value,
      page_size: pageSize.value
    })
    jingtanAssets.value = res.list || []
  } catch (error) {
    uni.showToast({
      title: error.message || '加载失败',
      icon: 'none'
    })
  } finally {
    loading.value = false
  }
}

// 去绑定
function goToBind() {
  uni.navigateTo({
    url: '/pages/third-party/index'
  })
}

// 返回
function goBack() {
  uni.navigateBack()
}
</script>

<style lang="scss" scoped>
.page {
  min-height: 100vh;
  background-color: #F5F5F5;
}

.navbar {
  display: flex;
  align-items: center;
  justify-content: space-between;
  height: 88rpx;
  padding: 0 32rpx;
  background-color: #FFFFFF;
  border-bottom: 1px solid #E8E8E8;
  
  .back-btn {
    width: 64rpx;
    height: 64rpx;
    display: flex;
    align-items: center;
    justify-content: center;
    
    .back-icon {
      font-size: 48rpx;
      color: #000000;
    }
  }
  
  .navbar-title {
    font-size: 32rpx;
    font-weight: 600;
    color: #000000;
  }
  
  .placeholder {
    width: 64rpx;
  }
}

.bind-tip {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 120rpx 64rpx;
  
  .tip-emoji {
    font-size: 120rpx;
    margin-bottom: 24rpx;
  }
  
  .tip-text {
    font-size: 28rpx;
    color: #999999;
    margin-bottom: 48rpx;
  }
  
  .btn {
    padding: 20rpx 64rpx;
    border-radius: 48rpx;
    font-size: 28rpx;
    font-weight: 600;
    border: none;
    
    &.btn-primary {
      background-color: #000000;
      color: #FFFFFF;
    }
    
    &.btn-small {
      min-width: 200rpx;
    }
  }
}

.asset-list-wrapper {
  padding: 32rpx;
}

.loading-wrapper,
.empty-wrapper {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 120rpx 0;
  
  .loading-text,
  .empty-text {
    font-size: 28rpx;
    color: #999999;
    margin-top: 24rpx;
  }
  
  .empty-emoji {
    font-size: 120rpx;
  }
}

.asset-list {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 24rpx;
}

.asset-card {
  background-color: #FFFFFF;
  border-radius: 16rpx;
  overflow: hidden;
  
  .asset-image {
    position: relative;
    width: 100%;
    height: 300rpx;
    background: linear-gradient(135deg, #f5f7fa 0%, #c3cfe2 100%);
    display: flex;
    align-items: center;
    justify-content: center;
    
    .asset-emoji {
      font-size: 80rpx;
    }
    
    .asset-tag {
      position: absolute;
      top: 12rpx;
      left: 12rpx;
      background-color: rgba(0, 0, 0, 0.7);
      color: #FFFFFF;
      font-size: 20rpx;
      padding: 6rpx 12rpx;
      border-radius: 6rpx;
    }
  }
  
  .asset-info {
    padding: 20rpx;
    
    .asset-name {
      display: block;
      font-size: 26rpx;
      font-weight: 600;
      color: #000000;
      margin-bottom: 8rpx;
      overflow: hidden;
      text-overflow: ellipsis;
      white-space: nowrap;
    }
    
    .asset-source {
      display: block;
      font-size: 22rpx;
      color: #999999;
    }
  }
}
</style>
