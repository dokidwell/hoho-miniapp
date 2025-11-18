<template>
  <view class="page-container">
    <!-- 用户信息卡片 -->
    <view class="user-card">
      <view class="user-header">
        <image :src="user.avatar_url" class="user-avatar"></image>
        <view class="user-info">
          <view class="user-name">{{ user.phone }}</view>
          <view class="user-status">
            <view class="status-badge" v-if="!user.identity_verified">未认证</view>
            <view class="status-badge verified" v-else>已认证</view>
            <view class="status-badge" v-if="user.identity_verified">已绑定身份</view>
          </view>
        </view>
        <uni-icons type="right" size="20" color="#999"></uni-icons>
      </view>

      <!-- 积分余额 -->
      <view class="points-section">
        <view class="points-item">
          <text class="points-label">积分不足？</text>
          <text class="points-link">来...</text>
        </view>
        <view class="task-center">
          <text>任务中心</text>
        </view>
      </view>

      <!-- 用户等级 -->
      <view class="level-section">
        <view class="level-badge">
          <text class="level-text">野生HOHO</text>
          <text class="level-rank">Lv.1</text>
        </view>
      </view>
    </view>

    <!-- 功能菜单 -->
    <view class="menu-section">
      <view class="menu-grid">
        <view class="menu-item" @click="goToPage('collection')">
          <view class="menu-icon">
            <uni-icons type="folder" size="24" color="#333"></uni-icons>
          </view>
          <text class="menu-text">作品集</text>
        </view>

        <view class="menu-item" @click="goToPage('periphery')">
          <view class="menu-icon">
            <uni-icons type="compose" size="24" color="#333"></uni-icons>
          </view>
          <text class="menu-text">周边</text>
        </view>

        <view class="menu-item" @click="goToPage('service')">
          <view class="menu-icon">
            <uni-icons type="headphones" size="24" color="#333"></uni-icons>
          </view>
          <text class="menu-text">客服</text>
        </view>

        <view class="menu-item" @click="goToPage('settings')">
          <view class="menu-icon">
            <uni-icons type="gear" size="24" color="#333"></uni-icons>
          </view>
          <text class="menu-text">设置</text>
        </view>

        <view class="menu-item" @click="goToPage('exchange-record')">
          <view class="menu-icon">
            <uni-icons type="circle" size="24" color="#333"></uni-icons>
          </view>
          <text class="menu-text">集换记录</text>
        </view>

        <view class="menu-item" @click="goToPage('community-works')">
          <view class="menu-icon">
            <uni-icons type="star" size="24" color="#333"></uni-icons>
          </view>
          <text class="menu-text">社区作品</text>
        </view>

        <view class="menu-item" @click="goToPage('jingtan-works')">
          <view class="menu-icon">
            <uni-icons type="diamond" size="24" color="#333"></uni-icons>
          </view>
          <text class="menu-text">鲸探作品</text>
        </view>

        <view class="menu-item" @click="goToPage('waveup-works')">
          <view class="menu-icon">
            <image src="../../static/icons/waveup.png" class="waveup-icon"></image>
          </view>
          <text class="menu-text">WAVEUP作品</text>
        </view>
      </view>
    </view>

    <!-- 账户信息 -->
    <view class="account-section">
      <view class="account-item">
        <text class="account-label">身份认证</text>
        <uni-icons type="right" size="20" color="#999"></uni-icons>
      </view>

      <view class="account-item">
        <view class="account-left">
          <text class="account-label">UID</text>
        </view>
        <text class="account-value">{{ user.uid }}</text>
      </view>

      <view class="account-item">
        <text class="account-label">第三方关联</text>
        <uni-icons type="right" size="20" color="#999"></uni-icons>
      </view>
    </view>

    <!-- 热门活动 -->
    <view class="activity-section">
      <text class="section-title">热门活动</text>
      <view class="activity-grid">
        <view class="activity-card">
          <text class="activity-placeholder">暂时没有更多内容...</text>
        </view>
        <view class="activity-card">
          <text class="activity-placeholder">暂时没有更多内容...</text>
        </view>
      </view>
    </view>
  </view>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useUserStore } from '../../stores/user'

const userStore = useUserStore()

const user = ref({
  phone: '199****9999',
  avatar_url: 'https://via.placeholder.com/80x80?text=User',
  identity_verified: false,
  uid: '1000000010',
  nickname: '野生HOHO'
})

// 获取用户信息
const fetchUserProfile = async () => {
  try {
    const profile = await userStore.fetchProfile()
    user.value = profile
  } catch (error) {
    console.error('Failed to fetch profile:', error)
  }
}

// 跳转到页面
const goToPage = (page) => {
  const pageMap = {
    collection: '/pages/profile/collection',
    periphery: '/pages/profile/periphery',
    service: '/pages/profile/service',
    settings: '/pages/profile/settings',
    'exchange-record': '/pages/profile/exchange-record',
    'community-works': '/pages/profile/community-works',
    'jingtan-works': '/pages/profile/jingtan-works',
    'waveup-works': '/pages/profile/waveup-works'
  }

  if (pageMap[page]) {
    uni.navigateTo({
      url: pageMap[page]
    })
  }
}

// 页面加载
onMounted(() => {
  // 检查登录状态
  if (!userStore.isLoggedIn) {
    uni.navigateTo({
      url: '/pages/auth/login'
    })
    return
  }

  fetchUserProfile()
})
</script>

<style lang="scss" scoped>
.page-container {
  width: 100%;
  min-height: 100vh;
  background-color: #f5f5f5;
  padding-bottom: 20px;
}

// 用户卡片
.user-card {
  background-color: #ffffff;
  padding: 16px;
  margin: 12px;
  border-radius: 12px;
}

.user-header {
  display: flex;
  align-items: center;
  gap: 12px;
  margin-bottom: 16px;
  cursor: pointer;
}

.user-avatar {
  width: 60px;
  height: 60px;
  border-radius: 50%;
  background-color: #f0f0f0;
}

.user-info {
  flex: 1;
}

.user-name {
  font-size: 16px;
  font-weight: 600;
  color: #000;
  margin-bottom: 4px;
}

.user-status {
  display: flex;
  gap: 8px;
  align-items: center;
}

.status-badge {
  font-size: 12px;
  color: #fff;
  background-color: #999;
  padding: 2px 8px;
  border-radius: 4px;

  &.verified {
    background-color: #52c41a;
  }
}

// 积分部分
.points-section {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 12px;
  background-color: #f5f5f5;
  border-radius: 8px;
  margin-bottom: 12px;
}

.points-item {
  flex: 1;
  display: flex;
  align-items: center;
  gap: 4px;
}

.points-label {
  font-size: 14px;
  color: #666;
}

.points-link {
  font-size: 14px;
  color: #3a8fff;
  cursor: pointer;
}

.task-center {
  padding: 6px 12px;
  background-color: #000;
  color: #fff;
  border-radius: 20px;
  font-size: 12px;
  cursor: pointer;
}

// 等级部分
.level-section {
  background-color: #f5f5f5;
  padding: 12px;
  border-radius: 8px;
}

.level-badge {
  display: flex;
  align-items: center;
  gap: 8px;
}

.level-text {
  font-size: 14px;
  color: #333;
}

.level-rank {
  font-size: 12px;
  color: #999;
}

// 菜单部分
.menu-section {
  padding: 12px;
}

.menu-grid {
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  gap: 12px;
  background-color: #ffffff;
  padding: 12px;
  border-radius: 12px;
}

.menu-item {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 8px;
  cursor: pointer;
  transition: transform 0.3s ease;

  &:active {
    transform: scale(0.95);
  }
}

.menu-icon {
  width: 40px;
  height: 40px;
  display: flex;
  align-items: center;
  justify-content: center;
  background-color: #f5f5f5;
  border-radius: 8px;
}

.waveup-icon {
  width: 24px;
  height: 24px;
}

.menu-text {
  font-size: 12px;
  color: #333;
  text-align: center;
}

// 账户信息
.account-section {
  background-color: #ffffff;
  margin: 12px;
  border-radius: 12px;
  overflow: hidden;
}

.account-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 12px 16px;
  border-bottom: 1px solid #f0f0f0;
  cursor: pointer;

  &:last-child {
    border-bottom: none;
  }
}

.account-label {
  font-size: 14px;
  color: #333;
}

.account-left {
  flex: 1;
}

.account-value {
  font-size: 14px;
  color: #999;
}

// 活动部分
.activity-section {
  padding: 12px;
}

.section-title {
  font-size: 14px;
  font-weight: 600;
  color: #333;
  display: block;
  margin-bottom: 12px;
}

.activity-grid {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 12px;
}

.activity-card {
  background-color: #ffffff;
  padding: 20px;
  border-radius: 12px;
  text-align: center;
  min-height: 100px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.activity-placeholder {
  font-size: 12px;
  color: #999;
}
</style>
