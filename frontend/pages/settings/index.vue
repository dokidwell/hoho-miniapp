<template>
  <view class="page">
    <view class="navbar">
      <text class="navbar-title">设置</text>
    </view>

    <view class="settings-list">
      <!-- 账号安全 -->
      <view class="settings-section">
        <text class="section-title">账号安全</text>
        <view class="settings-item" @click="goToPage('/pages/profile-edit/index')">
          <text class="item-label">个人资料</text>
          <view class="item-right">
            <text class="item-arrow">›</text>
          </view>
        </view>
        <view class="settings-item" @click="changePassword">
          <text class="item-label">修改密码</text>
          <view class="item-right">
            <text class="item-arrow">›</text>
          </view>
        </view>
      </view>

      <!-- 通知设置 -->
      <view class="settings-section">
        <text class="section-title">通知设置</text>
        <view class="settings-item">
          <text class="item-label">交易通知</text>
          <switch :checked="notificationSettings.trade" @change="toggleNotification('trade')" />
        </view>
        <view class="settings-item">
          <text class="item-label">系统通知</text>
          <switch :checked="notificationSettings.system" @change="toggleNotification('system')" />
        </view>
      </view>

      <!-- 关于 -->
      <view class="settings-section">
        <text class="section-title">关于</text>
        <view class="settings-item" @click="goToPage('/pages/agreement/index?type=service')">
          <text class="item-label">用户协议</text>
          <view class="item-right">
            <text class="item-arrow">›</text>
          </view>
        </view>
        <view class="settings-item" @click="goToPage('/pages/agreement/index?type=privacy')">
          <text class="item-label">隐私政策</text>
          <view class="item-right">
            <text class="item-arrow">›</text>
          </view>
        </view>
        <view class="settings-item" @click="goToPage('/pages/whitepaper/index')">
          <text class="item-label">白皮书</text>
          <view class="item-right">
            <text class="item-arrow">›</text>
          </view>
        </view>
        <view class="settings-item">
          <text class="item-label">当前版本</text>
          <view class="item-right">
            <text class="item-value">v1.0.0</text>
          </view>
        </view>
      </view>

      <!-- 退出登录 -->
      <view class="logout-btn" @click="handleLogout">
        <text class="logout-text">退出登录</text>
      </view>
    </view>
  </view>
</template>

<script setup>
import { ref } from 'vue'
import request from '@/api/request'
import { API_ENDPOINTS } from '@/api/config'

const notificationSettings = ref({
  trade: true,
  system: true
})

function goToPage(url) {
  uni.navigateTo({ url })
}

function changePassword() {
  uni.showModal({
    title: '修改密码',
    content: '请前往个人中心修改密码',
    showCancel: false
  })
}

function toggleNotification(type) {
  notificationSettings.value[type] = !notificationSettings.value[type]
  // TODO: 调用API保存设置
}

function handleLogout() {
  uni.showModal({
    title: '提示',
    content: '确定要退出登录吗？',
    success: (res) => {
      if (res.confirm) {
        uni.removeStorageSync('token')
        uni.reLaunch({
          url: '/pages/login/index'
        })
      }
    }
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

.settings-list {
  padding: 24rpx 32rpx;
}

.settings-section {
  margin-bottom: 32rpx;
  
  .section-title {
    font-size: 26rpx;
    color: #999999;
    display: block;
    margin-bottom: 16rpx;
    padding-left: 8rpx;
  }
}

.settings-item {
  background-color: #FFFFFF;
  border-radius: 16rpx;
  padding: 32rpx 24rpx;
  margin-bottom: 2rpx;
  display: flex;
  justify-content: space-between;
  align-items: center;
  
  &:first-of-type {
    border-radius: 16rpx 16rpx 0 0;
  }
  
  &:last-of-type {
    border-radius: 0 0 16rpx 16rpx;
    margin-bottom: 0;
  }
  
  &:only-of-type {
    border-radius: 16rpx;
  }
  
  .item-label {
    font-size: 28rpx;
    color: #000000;
  }
  
  .item-right {
    display: flex;
    align-items: center;
    gap: 16rpx;
    
    .item-value {
      font-size: 26rpx;
      color: #999999;
    }
    
    .item-arrow {
      font-size: 40rpx;
      color: #CCCCCC;
      line-height: 1;
    }
  }
}

.logout-btn {
  background-color: #FFFFFF;
  border-radius: 16rpx;
  padding: 32rpx;
  text-align: center;
  margin-top: 48rpx;
  
  .logout-text {
    font-size: 28rpx;
    color: #F44336;
    font-weight: 600;
  }
}
</style>
