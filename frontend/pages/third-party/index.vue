<template>
  <view class="page">
    <!-- 顶部导航 -->
    <view class="navbar">
      <view class="back-btn" @click="goBack">
        <text class="back-icon">←</text>
      </view>
      <view class="navbar-title">第三方关联</view>
      <view class="placeholder"></view>
    </view>

    <!-- 关联列表 -->
    <view class="bind-list">
      <!-- 鲸探 -->
      <view class="bind-item">
        <view class="bind-info">
          <view class="bind-logo">🐋</view>
          <view class="bind-text">
            <text class="bind-name">鲸探</text>
            <text class="bind-status" :class="{ 'bind-status-active': bindings.jingtan }">
              {{ bindings.jingtan ? '已绑定' : '未绑定' }}
            </text>
          </view>
        </view>
        <button 
          v-if="!bindings.jingtan" 
          class="btn btn-primary btn-small" 
          @click="bindJingtan"
        >
          绑定
        </button>
        <button 
          v-else 
          class="btn btn-secondary btn-small" 
          @click="unbindJingtan"
        >
          解绑
        </button>
      </view>

      <!-- Waveup -->
      <view class="bind-item">
        <view class="bind-info">
          <view class="bind-logo">🌊</view>
          <view class="bind-text">
            <text class="bind-name">Waveup</text>
            <text class="bind-status">未绑定</text>
          </view>
        </view>
        <button class="btn btn-primary btn-small" @click="showComingSoon">绑定</button>
      </view>

      <!-- XMeta -->
      <view class="bind-item">
        <view class="bind-info">
          <view class="bind-logo">🎮</view>
          <view class="bind-text">
            <text class="bind-name">XMeta</text>
            <text class="bind-status">未绑定</text>
          </view>
        </view>
        <button class="btn btn-primary btn-small" @click="showComingSoon">绑定</button>
      </view>
    </view>

    <!-- 绑定说明 -->
    <view class="bind-notice">
      <view class="notice-title">📌 绑定说明</view>
      <view class="notice-content">
        <text class="notice-text">1. 绑定第三方账户后，可以查看您在该平台的藏品</text>
        <text class="notice-text">2. 第三方藏品仅供展示，不可在本平台交易</text>
        <text class="notice-text">3. 绑定的手机号必须与小程序登录手机号一致</text>
      </view>
    </view>
  </view>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import request from '@/api/request'
import { API_ENDPOINTS } from '@/api/config'

const bindings = ref({
  jingtan: false,
  waveup: false,
  xmeta: false
})

onMounted(() => {
  fetchBindings()
})

// 获取绑定状态
async function fetchBindings() {
  try {
    const res = await request.get(API_ENDPOINTS.THIRD_PARTY.LIST)
    bindings.value = {
      jingtan: res.some(item => item.platform === 'jingtan'),
      waveup: res.some(item => item.platform === 'waveup'),
      xmeta: res.some(item => item.platform === 'xmeta')
    }
  } catch (error) {
    console.error('获取绑定状态失败:', error)
  }
}

// 绑定鲸探账户
function bindJingtan() {
  uni.showModal({
    title: '绑定鲸探账户',
    content: '请确保您的鲸探账户手机号与小程序登录手机号一致',
    editable: true,
    placeholderText: '请输入鲸探账户ID',
    success: async (res) => {
      if (res.confirm && res.content) {
        try {
          await request.post(API_ENDPOINTS.THIRD_PARTY.BIND, {
            platform: 'jingtan',
            account_id: res.content
          })
          
          uni.showToast({
            title: '绑定成功',
            icon: 'success'
          })
          
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
function unbindJingtan() {
  uni.showModal({
    title: '确认解绑',
    content: '解绑后将无法查看鲸探作品',
    success: async (res) => {
      if (res.confirm) {
        try {
          await request.delete(API_ENDPOINTS.THIRD_PARTY.UNBIND('jingtan'))
          
          uni.showToast({
            title: '解绑成功',
            icon: 'success'
          })
          
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

// 敬请期待
function showComingSoon() {
  uni.showToast({
    title: '该平台即将开放',
    icon: 'none'
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

.bind-list {
  padding: 32rpx;
  
  .bind-item {
    display: flex;
    align-items: center;
    justify-content: space-between;
    background-color: #FFFFFF;
    border-radius: 16rpx;
    padding: 32rpx;
    margin-bottom: 24rpx;
    
    .bind-info {
      display: flex;
      align-items: center;
      flex: 1;
      
      .bind-logo {
        width: 96rpx;
        height: 96rpx;
        border-radius: 50%;
        background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
        display: flex;
        align-items: center;
        justify-content: center;
        font-size: 56rpx;
        margin-right: 24rpx;
      }
      
      .bind-text {
        display: flex;
        flex-direction: column;
        gap: 8rpx;
        
        .bind-name {
          font-size: 30rpx;
          font-weight: 600;
          color: #000000;
        }
        
        .bind-status {
          font-size: 24rpx;
          color: #999999;
          
          &.bind-status-active {
            color: #52C41A;
          }
        }
      }
    }
    
    .btn {
      padding: 16rpx 32rpx;
      border-radius: 48rpx;
      font-size: 26rpx;
      font-weight: 600;
      border: none;
      
      &.btn-primary {
        background-color: #000000;
        color: #FFFFFF;
      }
      
      &.btn-secondary {
        background-color: #F5F5F5;
        color: #666666;
      }
      
      &.btn-small {
        min-width: 120rpx;
      }
    }
  }
}

.bind-notice {
  margin: 32rpx;
  background-color: #FFF9E6;
  border-radius: 16rpx;
  padding: 32rpx;
  
  .notice-title {
    font-size: 28rpx;
    font-weight: 600;
    color: #FF9800;
    margin-bottom: 16rpx;
  }
  
  .notice-content {
    display: flex;
    flex-direction: column;
    gap: 12rpx;
    
    .notice-text {
      font-size: 24rpx;
      color: #FF9800;
      line-height: 1.6;
    }
  }
}
</style>
