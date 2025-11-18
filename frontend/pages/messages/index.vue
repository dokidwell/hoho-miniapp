<template>
  <view class="page">
    <view class="navbar">
      <text class="navbar-title">消息中心</text>
    </view>

    <view class="messages-list">
      <view 
        v-for="message in messages" 
        :key="message.id"
        class="message-item"
        :class="{ unread: !message.is_read }"
        @click="handleMessageClick(message)"
      >
        <view class="message-icon">{{ message.icon }}</view>
        <view class="message-content">
          <view class="message-header">
            <text class="message-title">{{ message.title }}</text>
            <text class="message-time">{{ formatTime(message.created_at) }}</text>
          </view>
          <text class="message-text">{{ message.content }}</text>
        </view>
        <view v-if="!message.is_read" class="unread-dot"></view>
      </view>
    </view>

    <view v-if="loading" class="loading-wrapper">
      <text class="loading-text">加载中...</text>
    </view>

    <view v-else-if="messages.length === 0" class="empty-wrapper">
      <text class="empty-emoji">📬</text>
      <text class="empty-text">暂无消息</text>
    </view>
  </view>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import request from '@/api/request'
import { API_ENDPOINTS } from '@/api/config'
import { formatTime } from '@/utils/format'

const messages = ref([])
const loading = ref(false)

onMounted(() => {
  fetchMessages()
})

async function fetchMessages() {
  loading.value = true
  try {
    const res = await request.get(API_ENDPOINTS.USER.MESSAGES)
    messages.value = res.list || []
  } catch (error) {
    console.error('获取消息失败:', error)
  } finally {
    loading.value = false
  }
}

async function handleMessageClick(message) {
  if (!message.is_read) {
    try {
      await request.post(`${API_ENDPOINTS.USER.MESSAGES}/${message.id}/read`)
      message.is_read = true
    } catch (error) {
      console.error('标记已读失败:', error)
    }
  }
  
  // 根据消息类型跳转
  if (message.link_type === 'asset') {
    uni.navigateTo({
      url: `/pages/asset-detail/index?id=${message.link_id}`
    })
  }
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

.messages-list {
  padding: 24rpx 32rpx;
}

.message-item {
  background-color: #FFFFFF;
  border-radius: 16rpx;
  padding: 24rpx;
  margin-bottom: 16rpx;
  display: flex;
  gap: 24rpx;
  position: relative;
  
  &.unread {
    background-color: #F0F7FF;
  }
  
  .message-icon {
    font-size: 64rpx;
    line-height: 1;
  }
  
  .message-content {
    flex: 1;
    
    .message-header {
      display: flex;
      justify-content: space-between;
      align-items: center;
      margin-bottom: 8rpx;
      
      .message-title {
        font-size: 28rpx;
        font-weight: 600;
        color: #000000;
      }
      
      .message-time {
        font-size: 22rpx;
        color: #999999;
      }
    }
    
    .message-text {
      font-size: 26rpx;
      color: #666666;
      line-height: 1.6;
      display: -webkit-box;
      -webkit-line-clamp: 2;
      -webkit-box-orient: vertical;
      overflow: hidden;
    }
  }
  
  .unread-dot {
    position: absolute;
    top: 24rpx;
    right: 24rpx;
    width: 16rpx;
    height: 16rpx;
    background-color: #F44336;
    border-radius: 50%;
  }
}

.loading-wrapper, .empty-wrapper {
  padding: 200rpx 0;
  text-align: center;
  
  .loading-text, .empty-text {
    font-size: 28rpx;
    color: #999999;
    display: block;
  }
  
  .empty-emoji {
    font-size: 120rpx;
    display: block;
    margin-bottom: 24rpx;
  }
}
</style>
