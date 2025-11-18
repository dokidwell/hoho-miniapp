<template>
  <view class="page">
    <!-- 顶部导航 -->
    <view class="navbar">
      <view class="navbar-title">生态</view>
      <view class="nav-actions">
        <text class="nav-icon" @click="showMore">⋯</text>
        <text class="nav-icon" @click="scanCode">📷</text>
      </view>
    </view>

    <!-- 搜索框 -->
    <view class="search-section">
      <view class="search-box">
        <text class="search-icon">🔍</text>
        <input class="search-input" placeholder="请输入内容" v-model="keyword" />
      </view>
    </view>

    <!-- Banner -->
    <view class="banner-section">
      <view class="banner-card">
        <text class="banner-text">🌟 探索HOHO生态系统</text>
      </view>
    </view>

    <!-- 功能入口 -->
    <view class="function-section">
      <view class="section-title">核心功能</view>
      
      <view class="function-list">
        <view class="function-item" @click="goTo('/pages/whitepaper/index')">
          <view class="function-content">
            <view class="function-icon">📄</view>
            <view class="function-info">
              <text class="function-title">白皮书</text>
              <text class="function-desc">完整了解HOHO IP进程</text>
            </view>
          </view>
          <text class="arrow-icon">→</text>
        </view>
        
        <view class="function-item" @click="goTo('/pages/governance/index')">
          <view class="function-content">
            <view class="function-icon">🗳️</view>
            <view class="function-info">
              <text class="function-title">社区治理</text>
              <text class="function-desc">每一位社区成员都可以参加</text>
            </view>
          </view>
          <text class="arrow-icon">→</text>
        </view>
        
        <view class="function-item" @click="goTo('/pages/transparent-ledger/index')">
          <view class="function-content">
            <view class="function-icon">🔒</view>
            <view class="function-info">
              <text class="function-title">透明公示</text>
              <text class="function-desc">不可篡改的透明制度</text>
            </view>
          </view>
          <text class="arrow-icon">→</text>
        </view>
      </view>
    </view>

    <!-- 流动广告 -->
    <view class="ad-section">
      <view class="ad-content">
        <text class="ad-icon">📢</text>
        <text class="ad-text">欢迎来到HOHO Park数字作品集换平台</text>
      </view>
    </view>

    <!-- HOHO和他的朋友们 -->
    <view class="partners-section">
      <view class="section-title">HOHO和他的朋友们</view>
      <view class="partners-grid">
        <view class="partner-item" v-for="partner in partners" :key="partner.id">
          <view class="partner-logo">{{ partner.emoji }}</view>
          <text class="partner-name">{{ partner.name }}</text>
        </view>
      </view>
    </view>

    <!-- 底部导航栏 -->
    <TabBar :active="3" />
  </view>
</template>

<script setup>
import { ref } from 'vue'
import TabBar from '@/components/TabBar/TabBar.vue'

const keyword = ref('')

const partners = ref([
  { id: 1, name: '鲸探', emoji: '🐋' },
  { id: 2, name: 'Waveup', emoji: '🌊' },
  { id: 3, name: 'XMeta', emoji: '🎮' },
  { id: 4, name: '品拍', emoji: '🎨' },
  { id: 5, name: '敬请期待', emoji: '✨' },
])

function goTo(url) {
  uni.navigateTo({ url })
}

function showMore() {
  uni.showToast({
    title: '更多功能开发中',
    icon: 'none'
  })
}

function scanCode() {
  uni.scanCode({
    success: (res) => {
      console.log('扫码结果:', res)
    }
  })
}
</script>

<style lang="scss" scoped>
.page {
  min-height: 100vh;
  background-color: #F5F5F5;
  padding-bottom: 140rpx;
}

.navbar {
  display: flex;
  align-items: center;
  justify-content: space-between;
  height: 88rpx;
  padding: 0 32rpx;
  background-color: #FFFFFF;
  border-bottom: 1px solid #E8E8E8;
  
  .navbar-title {
    font-size: 32rpx;
    font-weight: 600;
    color: #000000;
  }
  
  .nav-actions {
    display: flex;
    gap: 24rpx;
    
    .nav-icon {
      font-size: 40rpx;
    }
  }
}

.search-section {
  padding: 24rpx 32rpx;
  background-color: #FFFFFF;
  
  .search-box {
    display: flex;
    align-items: center;
    background-color: #F5F5F5;
    border-radius: 48rpx;
    padding: 0 32rpx;
    height: 72rpx;
    
    .search-icon {
      font-size: 32rpx;
      margin-right: 16rpx;
    }
    
    .search-input {
      flex: 1;
      font-size: 28rpx;
    }
  }
}

.banner-section {
  padding: 32rpx;
  
  .banner-card {
    background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
    border-radius: 24rpx;
    padding: 80rpx 48rpx;
    display: flex;
    align-items: center;
    justify-content: center;
    
    .banner-text {
      font-size: 36rpx;
      font-weight: 600;
      color: #FFFFFF;
      letter-spacing: 2rpx;
    }
  }
}

.function-section {
  padding: 0 32rpx 32rpx;
  
  .section-title {
    font-size: 32rpx;
    font-weight: 600;
    color: #000000;
    margin-bottom: 24rpx;
  }
  
  .function-list {
    background-color: #FFFFFF;
    border-radius: 16rpx;
    overflow: hidden;
    
    .function-item {
      display: flex;
      align-items: center;
      justify-content: space-between;
      padding: 32rpx;
      border-bottom: 1px solid #F0F0F0;
      
      &:last-child {
        border-bottom: none;
      }
      
      .function-content {
        display: flex;
        align-items: center;
        flex: 1;
        
        .function-icon {
          font-size: 64rpx;
          margin-right: 24rpx;
        }
        
        .function-info {
          display: flex;
          flex-direction: column;
          gap: 8rpx;
          
          .function-title {
            font-size: 30rpx;
            font-weight: 600;
            color: #000000;
          }
          
          .function-desc {
            font-size: 24rpx;
            color: #999999;
          }
        }
      }
      
      .arrow-icon {
        font-size: 32rpx;
        color: #CCCCCC;
      }
    }
  }
}

.ad-section {
  margin: 32rpx;
  background-color: #FFF9E6;
  border-radius: 16rpx;
  padding: 24rpx 32rpx;
  
  .ad-content {
    display: flex;
    align-items: center;
    
    .ad-icon {
      font-size: 32rpx;
      margin-right: 16rpx;
    }
    
    .ad-text {
      flex: 1;
      font-size: 26rpx;
      color: #FF9800;
      font-weight: 500;
    }
  }
}

.partners-section {
  padding: 0 32rpx 32rpx;
  
  .section-title {
    font-size: 32rpx;
    font-weight: 600;
    color: #000000;
    margin-bottom: 24rpx;
  }
  
  .partners-grid {
    display: grid;
    grid-template-columns: repeat(4, 1fr);
    gap: 24rpx;
    
    .partner-item {
      display: flex;
      flex-direction: column;
      align-items: center;
      background-color: #FFFFFF;
      border-radius: 16rpx;
      padding: 32rpx 16rpx;
      
      .partner-logo {
        font-size: 64rpx;
        margin-bottom: 16rpx;
      }
      
      .partner-name {
        font-size: 22rpx;
        color: #666666;
        text-align: center;
      }
    }
  }
}
</style>
