<template>
  <view class="page">
    <!-- 顶部导航 -->
    <view class="navbar">
      <text class="navbar-title">集换记录</text>
    </view>

    <!-- 记录列表 -->
    <view class="history-list">
      <view 
        v-for="record in records" 
        :key="record.id"
        class="history-item"
        @click="goToDetail(record.asset_id)"
      >
        <view class="item-header">
          <view class="type-badge" :class="record.type">
            {{ getTypeName(record.type) }}
          </view>
          <text class="time">{{ formatTime(record.created_at) }}</text>
        </view>
        
        <view class="item-content">
          <image 
            class="asset-image" 
            :src="record.asset_image || '/static/images/placeholder.png'" 
            mode="aspectFill"
          />
          <view class="asset-info">
            <text class="asset-name">{{ record.asset_name }}</text>
            <text class="asset-number">#{{ record.serial_number }}</text>
            <view class="price-info">
              <text class="price-label">积分：</text>
              <text class="price-value number-display">{{ formatPoints(record.points) }}</text>
            </view>
          </view>
        </view>
        
        <view class="item-footer">
          <text class="status" :class="record.status">
            {{ getStatusName(record.status) }}
          </text>
        </view>
      </view>
    </view>

    <!-- 加载状态 -->
    <view v-if="loading" class="loading-wrapper">
      <text class="loading-text">加载中...</text>
    </view>

    <!-- 空状态 -->
    <view v-else-if="records.length === 0" class="empty-wrapper">
      <text class="empty-emoji">📋</text>
      <text class="empty-text">暂无记录</text>
    </view>
  </view>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import request from '@/api/request'
import { API_ENDPOINTS } from '@/api/config'
import { formatPoints, formatTime } from '@/utils/format'

const records = ref([])
const loading = ref(false)

onMounted(() => {
  fetchTradeHistory()
})

// 获取交易记录
async function fetchTradeHistory() {
  loading.value = true
  try {
    const res = await request.get(API_ENDPOINTS.TRADE.HISTORY)
    records.value = res.list || []
  } catch (error) {
    console.error('获取记录失败:', error)
    uni.showToast({
      title: '获取记录失败',
      icon: 'none'
    })
  } finally {
    loading.value = false
  }
}

// 获取类型名称
function getTypeName(type) {
  const typeMap = {
    'exchange': '兑换',
    'listing': '挂售',
    'cancel_listing': '取消挂售',
    'purchase': '购买'
  }
  return typeMap[type] || type
}

// 获取状态名称
function getStatusName(status) {
  const statusMap = {
    'pending': '处理中',
    'completed': '已完成',
    'cancelled': '已取消',
    'failed': '失败'
  }
  return statusMap[status] || status
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

.history-list {
  padding: 24rpx 32rpx;
}

.history-item {
  background-color: #FFFFFF;
  border-radius: 16rpx;
  padding: 24rpx;
  margin-bottom: 24rpx;
  box-shadow: 0 4rpx 12rpx rgba(0, 0, 0, 0.05);
  
  .item-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 16rpx;
    
    .type-badge {
      padding: 6rpx 16rpx;
      border-radius: 6rpx;
      font-size: 22rpx;
      font-weight: 600;
      
      &.exchange {
        background-color: #E3F2FD;
        color: #2196F3;
      }
      
      &.listing {
        background-color: #FFF3E0;
        color: #FF9800;
      }
      
      &.cancel_listing {
        background-color: #FFEBEE;
        color: #F44336;
      }
      
      &.purchase {
        background-color: #E8F5E9;
        color: #4CAF50;
      }
    }
    
    .time {
      font-size: 24rpx;
      color: #999999;
    }
  }
  
  .item-content {
    display: flex;
    gap: 24rpx;
    margin-bottom: 16rpx;
    
    .asset-image {
      width: 160rpx;
      height: 160rpx;
      border-radius: 12rpx;
      background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
    }
    
    .asset-info {
      flex: 1;
      display: flex;
      flex-direction: column;
      justify-content: center;
      
      .asset-name {
        font-size: 28rpx;
        font-weight: 600;
        color: #000000;
        margin-bottom: 8rpx;
        overflow: hidden;
        text-overflow: ellipsis;
        white-space: nowrap;
      }
      
      .asset-number {
        font-size: 24rpx;
        color: #999999;
        margin-bottom: 12rpx;
      }
      
      .price-info {
        display: flex;
        align-items: baseline;
        
        .price-label {
          font-size: 24rpx;
          color: #666666;
        }
        
        .price-value {
          font-size: 32rpx;
          font-weight: 600;
          color: #667eea;
          font-family: 'Courier New', Courier, monospace;
        }
      }
    }
  }
  
  .item-footer {
    text-align: right;
    
    .status {
      font-size: 24rpx;
      padding: 6rpx 16rpx;
      border-radius: 6rpx;
      
      &.completed {
        background-color: #E8F5E9;
        color: #4CAF50;
      }
      
      &.pending {
        background-color: #FFF3E0;
        color: #FF9800;
      }
      
      &.cancelled,
      &.failed {
        background-color: #FFEBEE;
        color: #F44336;
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
  }
}

.number-display {
  font-family: 'Courier New', Courier, monospace;
}
</style>
