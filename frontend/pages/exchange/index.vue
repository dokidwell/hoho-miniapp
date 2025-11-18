<template>
  <view class="page">
    <!-- 顶部导航 -->
    <view class="navbar">
      <view class="back-btn" @click="goBack">
        <text class="back-icon">←</text>
      </view>
      <view class="navbar-title">兑换</view>
      <view class="placeholder"></view>
    </view>

    <!-- 加载中 -->
    <view v-if="loading" class="loading-wrapper">
      <text class="loading-text">加载中...</text>
    </view>

    <!-- 兑换内容 -->
    <view v-else-if="listingInfo" class="content">
      <!-- 藏品预览 -->
      <view class="asset-preview">
        <view class="asset-image">
          <text class="asset-emoji">🖼️</text>
        </view>
        <view class="asset-info">
          <text class="asset-name">{{ listingInfo.asset?.name }}</text>
          <text class="asset-serial" v-if="listingInfo.serial_number">#{{ listingInfo.serial_number }}</text>
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
          <text class="price-value number-display" :class="{ 'insufficient': myPoints < parseFloat(listingInfo.price) }">
            {{ formatPoints(myPoints) }}
          </text>
        </view>
      </view>

      <!-- 卖家信息 -->
      <view class="seller-info">
        <text class="info-label">卖家</text>
        <view class="seller-detail">
          <text class="seller-emoji">👤</text>
          <text class="seller-name">{{ listingInfo.seller?.nickname || '匿名用户' }}</text>
        </view>
      </view>

      <!-- 提交按钮 -->
      <view class="action-section">
        <button 
          class="btn btn-primary" 
          @click="confirmExchange" 
          :loading="exchanging"
          :disabled="exchanging || myPoints < parseFloat(listingInfo.price)"
        >
          {{ myPoints < parseFloat(listingInfo.price) ? '积分不足' : (exchanging ? '兑换中...' : '确认兑换') }}
        </button>
      </view>
    </view>
  </view>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import request from '@/api/request'
import { API_ENDPOINTS } from '@/api/config'
import { formatPoints } from '@/utils/format'

const loading = ref(false)
const exchanging = ref(false)
const listingInfo = ref(null)
const myPoints = ref(0)
const listingId = ref('')

onMounted(() => {
  // 获取路由参数
  const pages = getCurrentPages()
  const currentPage = pages[pages.length - 1]
  listingId.value = currentPage.options.listingId
  
  if (listingId.value) {
    fetchListingInfo()
    fetchMyPoints()
  }
})

// 获取挂售信息
async function fetchListingInfo() {
  loading.value = true
  try {
    const res = await request.get(API_ENDPOINTS.TRADE.GET_LISTING_DETAIL(listingId.value))
    listingInfo.value = res
  } catch (error) {
    uni.showToast({
      title: error.message || '加载失败',
      icon: 'none'
    })
  } finally {
    loading.value = false
  }
}

// 获取我的积分
async function fetchMyPoints() {
  try {
    const res = await request.get(API_ENDPOINTS.USER.GET_POINTS)
    myPoints.value = parseFloat(res.balance || 0)
  } catch (error) {
    console.error('获取积分失败:', error)
  }
}

// 确认兑换
function confirmExchange() {
  uni.showModal({
    title: '确认兑换',
    content: `确认用 ${listingInfo.value.price} 积分兑换这个藏品吗？`,
    success: async (res) => {
      if (res.confirm) {
        await executeExchange()
      }
    }
  })
}

// 执行兑换
async function executeExchange() {
  exchanging.value = true
  try {
    const res = await request.post(API_ENDPOINTS.TRADE.EXECUTE_TRADE, {
      listing_id: listingId.value
    })
    
    uni.showModal({
      title: '兑换成功',
      content: '恭喜您成功兑换藏品！',
      showCancel: false,
      success: () => {
        uni.navigateBack()
      }
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

.loading-wrapper {
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 200rpx 0;
  
  .loading-text {
    font-size: 28rpx;
    color: #999999;
  }
}

.content {
  padding: 32rpx;
}

.asset-preview {
  display: flex;
  align-items: center;
  background-color: #FFFFFF;
  padding: 32rpx;
  border-radius: 16rpx;
  margin-bottom: 24rpx;
  
  .asset-image {
    width: 120rpx;
    height: 120rpx;
    background: linear-gradient(135deg, #f5f7fa 0%, #c3cfe2 100%);
    border-radius: 12rpx;
    display: flex;
    align-items: center;
    justify-content: center;
    margin-right: 24rpx;
    
    .asset-emoji {
      font-size: 64rpx;
    }
  }
  
  .asset-info {
    flex: 1;
    display: flex;
    flex-direction: column;
    gap: 8rpx;
    
    .asset-name {
      font-size: 30rpx;
      font-weight: 600;
      color: #000000;
    }
    
    .asset-serial {
      font-size: 24rpx;
      color: #999999;
    }
  }
}

.price-info {
  background-color: #FFFFFF;
  border-radius: 16rpx;
  padding: 32rpx;
  margin-bottom: 24rpx;
  
  .price-row {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 20rpx 0;
    
    &:not(:last-child) {
      border-bottom: 1rpx solid #F0F0F0;
    }
    
    .price-label {
      font-size: 28rpx;
      color: #666666;
    }
    
    .price-value {
      font-size: 32rpx;
      font-weight: 700;
      color: #000000;
      
      &.insufficient {
        color: #FF4D4F;
      }
    }
  }
}

.seller-info {
  background-color: #FFFFFF;
  border-radius: 16rpx;
  padding: 32rpx;
  margin-bottom: 24rpx;
  
  .info-label {
    display: block;
    font-size: 28rpx;
    color: #666666;
    margin-bottom: 16rpx;
  }
  
  .seller-detail {
    display: flex;
    align-items: center;
    gap: 12rpx;
    
    .seller-emoji {
      font-size: 40rpx;
    }
    
    .seller-name {
      font-size: 28rpx;
      color: #000000;
      font-weight: 500;
    }
  }
}

.action-section {
  padding-top: 32rpx;
  
  .btn {
    width: 100%;
    height: 96rpx;
    border-radius: 12rpx;
    font-size: 32rpx;
    font-weight: 600;
    border: none;
    
    &.btn-primary {
      background-color: #000000;
      color: #FFFFFF;
      
      &:disabled {
        opacity: 0.4;
        background-color: #CCCCCC;
      }
    }
  }
}

.number-display {
  font-family: 'Monaco', 'Menlo', 'Consolas', monospace;
}
</style>
