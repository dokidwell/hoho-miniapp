<template>
  <view class="page">
    <!-- 顶部导航 -->
    <view class="navbar">
      <view class="back-btn" @click="goBack">
        <text class="back-icon">←</text>
      </view>
      <view class="navbar-title">挂售</view>
      <view class="placeholder"></view>
    </view>

    <!-- 作品预览 -->
    <view class="asset-preview">
      <view class="asset-image">
        <text class="asset-emoji">🖼️</text>
      </view>
      <view class="asset-info">
        <text class="asset-name">{{ assetInfo.name || '作品名称' }}</text>
        <text class="asset-serial" v-if="assetInfo.serial_number">#{{ assetInfo.serial_number }}</text>
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
          <text class="fee-value number-display highlight">{{ calculateReceived(listingPrice) }}</text>
        </view>
      </view>
    </view>

    <!-- 提交按钮 -->
    <view class="action-section">
      <button 
        class="btn btn-primary" 
        @click="submitListing" 
        :loading="submitting" 
        :disabled="submitting || !listingPrice"
      >
        {{ submitting ? '提交中...' : '确认挂售' }}
      </button>
    </view>
  </view>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import request from '@/api/request'
import { API_ENDPOINTS } from '@/api/config'
import { calculateFee, calculateReceived } from '@/utils/format'

const listingPrice = ref('')
const submitting = ref(false)
const assetInfo = ref({})
const assetId = ref('')

onMounted(() => {
  // 获取路由参数
  const pages = getCurrentPages()
  const currentPage = pages[pages.length - 1]
  assetId.value = currentPage.options.assetId
  
  if (assetId.value) {
    fetchAssetInfo()
  }
})

// 获取作品信息
async function fetchAssetInfo() {
  try {
    const res = await request.get(API_ENDPOINTS.ASSET.GET_DETAIL(assetId.value))
    assetInfo.value = res
  } catch (error) {
    console.error('获取作品信息失败:', error)
  }
}

// 提交挂售
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
    await request.post(API_ENDPOINTS.TRADE.CREATE_LISTING, {
      asset_instance_id: assetId.value,
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

.asset-preview {
  display: flex;
  align-items: center;
  background-color: #FFFFFF;
  margin: 24rpx 32rpx;
  padding: 32rpx;
  border-radius: 16rpx;
  
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

.price-setting {
  background-color: #FFFFFF;
  margin: 0 32rpx;
  padding: 32rpx;
  border-radius: 16rpx;
  
  .form-item {
    margin-bottom: 32rpx;
    
    .form-label {
      display: block;
      font-size: 28rpx;
      color: #000000;
      font-weight: 600;
      margin-bottom: 16rpx;
    }
    
    .price-input-wrapper {
      display: flex;
      align-items: center;
      background-color: #F5F5F5;
      border-radius: 12rpx;
      padding: 0 24rpx;
      height: 96rpx;
      
      .price-input {
        flex: 1;
        font-size: 36rpx;
        font-weight: 600;
        color: #000000;
      }
      
      .price-unit {
        font-size: 28rpx;
        color: #999999;
        margin-left: 16rpx;
      }
    }
  }
  
  .fee-section {
    background-color: #F5F5F5;
    border-radius: 12rpx;
    padding: 24rpx;
    
    .fee-item {
      display: flex;
      justify-content: space-between;
      align-items: center;
      padding: 16rpx 0;
      
      .fee-label {
        font-size: 26rpx;
        color: #666666;
      }
      
      .fee-value {
        font-size: 26rpx;
        color: #000000;
        
        &.highlight {
          font-size: 32rpx;
          font-weight: 700;
          color: #52C41A;
        }
      }
      
      &.total {
        padding-top: 24rpx;
        
        .fee-label {
          font-size: 28rpx;
          font-weight: 600;
          color: #000000;
        }
      }
    }
    
    .divider {
      height: 1rpx;
      background-color: #E8E8E8;
      margin: 16rpx 0;
    }
  }
}

.action-section {
  padding: 32rpx;
  
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
      }
    }
  }
}

.number-display {
  font-family: 'Monaco', 'Menlo', 'Consolas', monospace;
}
</style>
