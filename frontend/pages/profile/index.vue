<template>
  <view class="page">
    <!-- 顶部导航 -->
    <view class="navbar">
      <view class="nav-actions">
        <text class="nav-icon" @click="showMore">⋯</text>
        <text class="nav-icon" @click="scanCode">📷</text>
      </view>
    </view>

    <!-- 用户信息卡 -->
    <view class="user-card">
      <view class="user-header" @click="goToProfile">
        <view class="user-avatar">👤</view>
        <view class="user-info">
          <view class="user-phone">{{ maskPhone(userInfo.phone) }}</view>
          <view class="user-tags">
            <view class="tag">{{ userInfo.is_verified ? '已认证' : '未认证' }}</view>
            <view class="tag">{{ isBoundJingtan ? '已绑定鲸探' : '未绑定鲸探' }}</view>
          </view>
        </view>
        <text class="arrow-icon">→</text>
      </view>

      <!-- 积分卡 -->
      <view class="points-card">
        <view class="points-info">
          <text class="points-label">我的积分</text>
          <text class="points-value number-display">{{ formatPoints(points) }}</text>
        </view>
        <view class="points-action" @click="goToTasks">
          <text class="action-text">任务中心</text>
          <text class="action-icon">→</text>
        </view>
      </view>
    </view>

    <!-- 功能网格（4x2） -->
    <view class="function-grid">
      <view class="grid-item" @click="goTo('/pages/my-assets/index')">
        <text class="grid-icon">🎨</text>
        <text class="grid-label">作品集</text>
      </view>
      <view class="grid-item" @click="goTo('/pages/shop/index')">
        <text class="grid-icon">🛍️</text>
        <text class="grid-label">周边</text>
      </view>
      <view class="grid-item" @click="goTo('/pages/service/index')">
        <text class="grid-icon">💬</text>
        <text class="grid-label">客服</text>
      </view>
      <view class="grid-item" @click="goTo('/pages/settings/index')">
        <text class="grid-icon">⚙️</text>
        <text class="grid-label">设置</text>
      </view>
      <view class="grid-item" @click="goTo('/pages/trade-history/index')">
        <text class="grid-icon">🔄</text>
        <text class="grid-label">集换记录</text>
      </view>
      <view class="grid-item" @click="goTo('/pages/community-assets/index')">
        <text class="grid-icon">🌟</text>
        <text class="grid-label">社区作品</text>
      </view>
      <view class="grid-item" @click="goTo('/pages/jingtan-assets/index')">
        <text class="grid-icon">🐋</text>
        <text class="grid-label">鲸探作品</text>
      </view>
      <view class="grid-item" @click="goTo('/pages/waveup-assets/index')">
        <text class="grid-icon">🌊</text>
        <text class="grid-label">WAVEUP作品</text>
      </view>
    </view>

    <!-- 列表项 -->
    <view class="list-section">
      <view class="list-item" @click="goTo('/pages/identity-verify/index')">
        <text class="list-label">身份认证</text>
        <text class="list-arrow">→</text>
      </view>
      
      <view class="list-item">
        <text class="list-label">UID</text>
        <text class="list-value">{{ userInfo.id || '-' }}</text>
      </view>
      
      <view class="list-item" @click="goTo('/pages/third-party/index')">
        <text class="list-label">第三方关联</text>
        <text class="list-arrow">→</text>
      </view>
    </view>

    <!-- 热门活动 -->
    <view class="activity-section">
      <view class="section-title">热门活动</view>
      <view class="activity-cards">
        <view class="activity-card">
          <text class="activity-placeholder">暂时没有更多内容...</text>
        </view>
        <view class="activity-card">
          <text class="activity-placeholder">暂时没有更多内容...</text>
        </view>
      </view>
    </view>

    <!-- 底部导航栏 -->
    <TabBar :active="4" />
  </view>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import TabBar from '@/components/TabBar/TabBar.vue'
import request from '@/api/request'
import { API_ENDPOINTS } from '@/api/config'
import { maskPhone, formatPoints } from '@/utils/format'

const userInfo = ref({})
const points = ref(0)
const isBoundJingtan = ref(false)

onMounted(() => {
  fetchUserInfo()
  fetchPoints()
  checkJingtanBinding()
})

// 获取用户信息
async function fetchUserInfo() {
  try {
    const res = await request.get(API_ENDPOINTS.USER.PROFILE)
    userInfo.value = res
  } catch (error) {
    console.error('获取用户信息失败:', error)
  }
}

// 获取积分余额
async function fetchPoints() {
  try {
    const res = await request.get(API_ENDPOINTS.USER.GET_POINTS)
    points.value = parseFloat(res.balance || 0)
  } catch (error) {
    console.error('获取积分失败:', error)
  }
}

// 检查鲸探绑定状态
async function checkJingtanBinding() {
  try {
    const res = await request.get(API_ENDPOINTS.THIRD_PARTY.LIST)
    isBoundJingtan.value = res.some(item => item.platform === 'jingtan')
  } catch (error) {
    console.error('检查绑定状态失败:', error)
  }
}

// 跳转
function goTo(url) {
  uni.navigateTo({ url })
}

// 跳转到个人资料
function goToProfile() {
  uni.navigateTo({
    url: '/pages/profile-edit/index'
  })
}

// 跳转到任务中心
function goToTasks() {
  uni.showToast({
    title: '任务中心开发中',
    icon: 'none'
  })
}

// 显示更多
function showMore() {
  uni.showToast({
    title: '更多功能开发中',
    icon: 'none'
  })
}

// 扫码
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
  justify-content: flex-end;
  height: 88rpx;
  padding: 0 32rpx;
  background-color: #FFFFFF;
  
  .nav-actions {
    display: flex;
    gap: 24rpx;
    
    .nav-icon {
      font-size: 40rpx;
    }
  }
}

.user-card {
  background-color: #FFFFFF;
  margin: 24rpx 32rpx;
  padding: 32rpx;
  border-radius: 16rpx;
  
  .user-header {
    display: flex;
    align-items: center;
    margin-bottom: 32rpx;
    
    .user-avatar {
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
    
    .user-info {
      flex: 1;
      display: flex;
      flex-direction: column;
      gap: 12rpx;
      
      .user-phone {
        font-size: 32rpx;
        font-weight: 600;
        color: #000000;
      }
      
      .user-tags {
        display: flex;
        gap: 12rpx;
        
        .tag {
          font-size: 22rpx;
          color: #999999;
          background-color: #F5F5F5;
          padding: 6rpx 12rpx;
          border-radius: 6rpx;
        }
      }
    }
    
    .arrow-icon {
      font-size: 32rpx;
      color: #CCCCCC;
    }
  }
  
  .points-card {
    background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
    border-radius: 12rpx;
    padding: 32rpx;
    display: flex;
    justify-content: space-between;
    align-items: center;
    
    .points-info {
      display: flex;
      flex-direction: column;
      gap: 12rpx;
      
      .points-label {
        font-size: 24rpx;
        color: rgba(255, 255, 255, 0.8);
      }
      
      .points-value {
        font-size: 40rpx;
        font-weight: 700;
        color: #FFFFFF;
      }
    }
    
    .points-action {
      display: flex;
      align-items: center;
      gap: 8rpx;
      background-color: rgba(255, 255, 255, 0.2);
      padding: 12rpx 24rpx;
      border-radius: 48rpx;
      
      .action-text {
        font-size: 24rpx;
        color: #FFFFFF;
        font-weight: 500;
      }
      
      .action-icon {
        font-size: 24rpx;
        color: #FFFFFF;
      }
    }
  }
}

.function-grid {
  background-color: #FFFFFF;
  margin: 0 32rpx 24rpx;
  padding: 32rpx;
  border-radius: 16rpx;
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  gap: 32rpx 24rpx;
  
  .grid-item {
    display: flex;
    flex-direction: column;
    align-items: center;
    gap: 12rpx;
    
    .grid-icon {
      font-size: 56rpx;
    }
    
    .grid-label {
      font-size: 24rpx;
      color: #666666;
    }
  }
}

.list-section {
  background-color: #FFFFFF;
  margin: 0 32rpx 24rpx;
  border-radius: 16rpx;
  overflow: hidden;
  
  .list-item {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 32rpx;
    border-bottom: 1rpx solid #F0F0F0;
    
    &:last-child {
      border-bottom: none;
    }
    
    .list-label {
      font-size: 28rpx;
      color: #000000;
    }
    
    .list-value {
      font-size: 28rpx;
      color: #999999;
    }
    
    .list-arrow {
      font-size: 32rpx;
      color: #CCCCCC;
    }
  }
}

.activity-section {
  padding: 0 32rpx 32rpx;
  
  .section-title {
    font-size: 32rpx;
    font-weight: 600;
    color: #000000;
    margin-bottom: 24rpx;
  }
  
  .activity-cards {
    display: flex;
    gap: 24rpx;
    
    .activity-card {
      flex: 1;
      background-color: #FFFFFF;
      border-radius: 16rpx;
      padding: 64rpx 32rpx;
      display: flex;
      align-items: center;
      justify-content: center;
      
      .activity-placeholder {
        font-size: 24rpx;
        color: #CCCCCC;
      }
    }
  }
}

.number-display {
  font-family: 'Monaco', 'Menlo', 'Consolas', monospace;
}
</style>
