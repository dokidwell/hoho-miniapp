<template>
  <view class="page">
    <view class="navbar">
      <text class="navbar-title">选择地区</text>
    </view>
    
    <view class="content">
      <text class="emoji">🗺️</text>
      <text class="title">选择地区</text>
      <text class="subtitle">功能开发中...</text>
    </view>
  </view>
</template>

<script setup>
// 选择地区页面
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

.content {
  padding: 200rpx 32rpx;
  text-align: center;
  
  .emoji {
    font-size: 120rpx;
    display: block;
    margin-bottom: 24rpx;
  }
  
  .title {
    font-size: 32rpx;
    font-weight: 600;
    color: #000000;
    display: block;
    margin-bottom: 12rpx;
  }
  
  .subtitle {
    font-size: 26rpx;
    color: #999999;
    display: block;
  }
}
</style>
