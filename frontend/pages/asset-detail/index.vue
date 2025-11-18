<template>
  <view class="page">
    <!-- 顶部导航 -->
    <view class="navbar">
      <view class="back-btn" @click="goBack">
        <text class="back-icon">←</text>
      </view>
      <view class="nav-actions">
        <text class="nav-icon" @click="shareAsset">📤</text>
      </view>
    </view>

    <!-- 加载中 -->
    <view v-if="loading" class="loading-wrapper">
      <text class="loading-text">加载中...</text>
    </view>

    <!-- 藏品内容 -->
    <view v-else-if="assetDetail" class="content">
      <!-- 藏品图片 -->
      <view class="asset-image-section">
        <view class="asset-image-placeholder">
          <text class="placeholder-emoji">🖼️</text>
        </view>
      </view>

      <!-- 藏品信息 -->
      <view class="asset-info-section">
        <view class="asset-header">
          <text class="asset-name">{{ assetDetail.name }}</text>
          <text class="asset-serial" v-if="assetDetail.serial_number">#{{ assetDetail.serial_number }}</text>
        </view>
        
        <view class="asset-creator">
          <text class="creator-emoji">👤</text>
          <text class="creator-name">{{ assetDetail.creator?.nickname || '匿名创作者' }}</text>
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
            <text class="stat-value">{{ assetDetail.holder_count || 0 }}</text>
          </view>
          <view class="stat-item">
            <text class="stat-label">交易量</text>
            <text class="stat-value">{{ assetDetail.trade_count || 0 }}</text>
          </view>
        </view>
      </view>

      <!-- 操作按钮 -->
      <view class="action-section">
        <button class="btn btn-primary" @click="buyNow">立即兑换</button>
        <button v-if="isOwner" class="btn btn-secondary" @click="createListing">挂售</button>
      </view>
    </view>
  </view>
</template>

<script setup>
import { ref, onMounted, computed } from 'vue'
import request from '@/api/request'
import { API_ENDPOINTS } from '@/api/config'
import { formatPoints } from '@/utils/format'

const loading = ref(false)
const assetDetail = ref(null)
const assetId = ref('')

const isOwner = computed(() => {
  // TODO: 判断是否是持有者
  return false
})

onMounted(() => {
  // 获取路由参数
  const pages = getCurrentPages()
  const currentPage = pages[pages.length - 1]
  assetId.value = currentPage.options.id
  
  if (assetId.value) {
    fetchAssetDetail()
  }
})

// 获取藏品详情
async function fetchAssetDetail() {
  loading.value = true
  try {
    const res = await request.get(API_ENDPOINTS.ASSET.GET_DETAIL(assetId.value))
    assetDetail.value = res
  } catch (error) {
    uni.showToast({
      title: error.message || '加载失败',
      icon: 'none'
    })
  } finally {
    loading.value = false
  }
}

// 立即兑换
function buyNow() {
  uni.navigateTo({
    url: `/pages/exchange/index?assetId=${assetId.value}`
  })
}

// 挂售
function createListing() {
  uni.navigateTo({
    url: `/pages/listing-create/index?assetId=${assetId.value}`
  })
}

// 分享
function shareAsset() {
  uni.showShareMenu({
    title: assetDetail.value?.name,
    path: `/pages/asset-detail/index?id=${assetId.value}`
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
  
  .nav-actions {
    display: flex;
    gap: 24rpx;
    
    .nav-icon {
      font-size: 40rpx;
    }
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
  padding-bottom: 200rpx;
}

.asset-image-section {
  width: 100%;
  height: 750rpx;
  background-color: #FFFFFF;
  
  .asset-image-placeholder {
    width: 100%;
    height: 100%;
    background: linear-gradient(135deg, #f5f7fa 0%, #c3cfe2 100%);
    display: flex;
    align-items: center;
    justify-content: center;
    
    .placeholder-emoji {
      font-size: 200rpx;
    }
  }
}

.asset-info-section {
  background-color: #FFFFFF;
  margin-top: 16rpx;
  padding: 32rpx;
  
  .asset-header {
    display: flex;
    align-items: baseline;
    gap: 16rpx;
    margin-bottom: 24rpx;
    
    .asset-name {
      font-size: 40rpx;
      font-weight: 700;
      color: #000000;
    }
    
    .asset-serial {
      font-size: 28rpx;
      color: #999999;
    }
  }
  
  .asset-creator {
    display: flex;
    align-items: center;
    gap: 12rpx;
    margin-bottom: 24rpx;
    
    .creator-emoji {
      font-size: 32rpx;
    }
    
    .creator-name {
      font-size: 26rpx;
      color: #666666;
    }
  }
  
  .asset-description {
    font-size: 28rpx;
    color: #333333;
    line-height: 1.6;
    margin-bottom: 32rpx;
  }
  
  .price-section {
    display: flex;
    gap: 32rpx;
    padding: 32rpx 0;
    border-top: 1rpx solid #F0F0F0;
    border-bottom: 1rpx solid #F0F0F0;
    margin-bottom: 32rpx;
    
    .price-item {
      flex: 1;
      
      .price-label {
        display: block;
        font-size: 24rpx;
        color: #999999;
        margin-bottom: 12rpx;
      }
      
      .price-value {
        display: block;
        font-size: 32rpx;
        font-weight: 700;
        color: #000000;
      }
    }
  }
  
  .stats-section {
    display: flex;
    gap: 32rpx;
    
    .stat-item {
      flex: 1;
      text-align: center;
      
      .stat-label {
        display: block;
        font-size: 24rpx;
        color: #999999;
        margin-bottom: 12rpx;
      }
      
      .stat-value {
        display: block;
        font-size: 32rpx;
        font-weight: 700;
        color: #000000;
      }
    }
  }
}

.action-section {
  position: fixed;
  bottom: 0;
  left: 0;
  right: 0;
  display: flex;
  gap: 24rpx;
  padding: 24rpx 32rpx;
  background-color: #FFFFFF;
  border-top: 1rpx solid #E8E8E8;
  padding-bottom: constant(safe-area-inset-bottom);
  padding-bottom: env(safe-area-inset-bottom);
  
  .btn {
    flex: 1;
    height: 88rpx;
    border-radius: 12rpx;
    font-size: 30rpx;
    font-weight: 600;
    border: none;
    
    &.btn-primary {
      background-color: #000000;
      color: #FFFFFF;
    }
    
    &.btn-secondary {
      background-color: #F5F5F5;
      color: #000000;
    }
  }
}

.number-display {
  font-family: 'Monaco', 'Menlo', 'Consolas', monospace;
}
</style>
