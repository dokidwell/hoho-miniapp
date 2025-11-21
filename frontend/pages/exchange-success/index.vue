<template>
  <view class="exchange-success-page">
    <!-- 成功动画 -->
    <view class="success-animation">
      <view class="success-circle">
        <text class="success-icon">✓</text>
      </view>
      <text class="success-title">兑换成功！</text>
      <text class="success-subtitle">恭喜你获得新的数字藏品</text>
    </view>

    <!-- 藏品信息 -->
    <view class="asset-card">
      <image 
        class="asset-image" 
        :src="assetInfo.imageUrl || '/static/placeholder.png'"
        mode="aspectFill"
      />
      <view class="asset-info">
        <text class="asset-name">{{ assetInfo.name || '数字藏品' }}</text>
        <text class="asset-number">#{{ assetInfo.instanceNo || '0001' }}</text>
        <view class="asset-meta">
          <text class="meta-item">发行量：{{ assetInfo.totalSupply || '1000' }}</text>
          <text class="meta-item">编号：{{ assetInfo.serialNumber || '1' }}/{{ assetInfo.totalSupply || '1000' }}</text>
        </view>
      </view>
    </view>

    <!-- 兑换详情 -->
    <view class="exchange-details">
      <text class="details-title">兑换详情</text>
      <view class="detail-row">
        <text class="detail-label">消耗积分</text>
        <text class="detail-value points">-{{ exchangeInfo.points || '0' }}</text>
      </view>
      <view class="detail-row">
        <text class="detail-label">兑换时间</text>
        <text class="detail-value">{{ exchangeInfo.time || getCurrentTime() }}</text>
      </view>
      <view class="detail-row">
        <text class="detail-label">订单号</text>
        <text class="detail-value small">{{ exchangeInfo.orderId || generateOrderId() }}</text>
      </view>
    </view>

    <!-- 操作按钮 -->
    <view class="actions">
      <button class="action-btn primary" @click="viewAsset">
        <text class="btn-text">查看藏品</text>
      </button>
      <button class="action-btn secondary" @click="shareAsset">
        <text class="btn-text">分享给好友</text>
      </button>
      <button class="action-btn tertiary" @click="goHome">
        <text class="btn-text">返回首页</text>
      </button>
    </view>

    <!-- 温馨提示 -->
    <view class="tips">
      <text class="tips-title">💡 温馨提示</text>
      <text class="tips-text">• 藏品已自动添加到"我的作品"</text>
      <text class="tips-text">• 可在"透明公示"中查看所有权记录</text>
      <text class="tips-text">• 持有满7天后可在集换市场挂售</text>
    </view>
  </view>
</template>

<script setup>
import { ref, onMounted } from 'vue'

// 藏品信息
const assetInfo = ref({
  name: '',
  imageUrl: '',
  instanceNo: '',
  totalSupply: '',
  serialNumber: ''
})

// 兑换信息
const exchangeInfo = ref({
  points: '',
  time: '',
  orderId: ''
})

onMounted(() => {
  // 从页面参数获取信息
  const pages = getCurrentPages()
  const currentPage = pages[pages.length - 1]
  const options = currentPage.options
  
  if (options.assetId) {
    loadAssetInfo(options.assetId)
  }
  
  if (options.points) {
    exchangeInfo.value.points = options.points
  }
  
  if (options.orderId) {
    exchangeInfo.value.orderId = options.orderId
  }
  
  exchangeInfo.value.time = getCurrentTime()
})

// 加载藏品信息
const loadAssetInfo = async (assetId) => {
  try {
    // TODO: 调用API获取藏品信息
    // const res = await uni.$api.assets.getDetail(assetId)
    // assetInfo.value = res.data
    
    // 临时模拟数据
    assetInfo.value = {
      name: '限量数字藏品',
      imageUrl: '/static/placeholder.png',
      instanceNo: '0001',
      totalSupply: '1000',
      serialNumber: '1'
    }
  } catch (error) {
    console.error('加载藏品信息失败:', error)
  }
}

// 获取当前时间
const getCurrentTime = () => {
  const now = new Date()
  const year = now.getFullYear()
  const month = String(now.getMonth() + 1).padStart(2, '0')
  const day = String(now.getDate()).padStart(2, '0')
  const hours = String(now.getHours()).padStart(2, '0')
  const minutes = String(now.getMinutes()).padStart(2, '0')
  const seconds = String(now.getSeconds()).padStart(2, '0')
  
  return `${year}-${month}-${day} ${hours}:${minutes}:${seconds}`
}

// 生成订单号
const generateOrderId = () => {
  const timestamp = Date.now()
  const random = Math.floor(Math.random() * 10000).toString().padStart(4, '0')
  return `EX${timestamp}${random}`
}

// 查看藏品
const viewAsset = () => {
  uni.navigateTo({
    url: '/pages/asset-detail/index?id=' + (assetInfo.value.id || '1')
  })
}

// 分享藏品
const shareAsset = () => {
  uni.showShareMenu({
    withShareTicket: true,
    menus: ['shareAppMessage', 'shareTimeline']
  })
  
  uni.showToast({
    title: '点击右上角分享',
    icon: 'none'
  })
}

// 返回首页
const goHome = () => {
  uni.switchTab({
    url: '/pages/index/index'
  })
}

// 分享配置
onShareAppMessage(() => {
  return {
    title: `我刚兑换了${assetInfo.value.name}！`,
    path: '/pages/asset-detail/index?id=' + (assetInfo.value.id || '1'),
    imageUrl: assetInfo.value.imageUrl
  }
})

onShareTimeline(() => {
  return {
    title: `我刚兑换了${assetInfo.value.name}！`,
    query: 'id=' + (assetInfo.value.id || '1'),
    imageUrl: assetInfo.value.imageUrl
  }
})
</script>

<style scoped>
.exchange-success-page {
  min-height: 100vh;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  padding: 60rpx 30rpx;
}

/* 成功动画 */
.success-animation {
  text-align: center;
  margin-bottom: 60rpx;
  animation: fadeInUp 0.6s ease-out;
}

@keyframes fadeInUp {
  from {
    opacity: 0;
    transform: translateY(30rpx);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

.success-circle {
  width: 160rpx;
  height: 160rpx;
  background: white;
  border-radius: 50%;
  margin: 0 auto 40rpx;
  display: flex;
  align-items: center;
  justify-content: center;
  box-shadow: 0 20rpx 60rpx rgba(0, 0, 0, 0.2);
  animation: scaleIn 0.5s cubic-bezier(0.68, -0.55, 0.265, 1.55);
}

@keyframes scaleIn {
  from {
    transform: scale(0);
  }
  to {
    transform: scale(1);
  }
}

.success-icon {
  font-size: 100rpx;
  color: #52c41a;
  font-weight: bold;
}

.success-title {
  display: block;
  font-size: 48rpx;
  font-weight: bold;
  color: white;
  margin-bottom: 20rpx;
}

.success-subtitle {
  display: block;
  font-size: 28rpx;
  color: rgba(255, 255, 255, 0.8);
}

/* 藏品卡片 */
.asset-card {
  background: white;
  border-radius: 20rpx;
  overflow: hidden;
  margin-bottom: 30rpx;
  box-shadow: 0 20rpx 60rpx rgba(0, 0, 0, 0.15);
  animation: fadeInUp 0.6s ease-out 0.2s both;
}

.asset-image {
  width: 100%;
  height: 400rpx;
  background: #f5f5f5;
}

.asset-info {
  padding: 30rpx;
}

.asset-name {
  display: block;
  font-size: 36rpx;
  font-weight: bold;
  color: #333;
  margin-bottom: 10rpx;
}

.asset-number {
  display: block;
  font-size: 28rpx;
  color: #999;
  margin-bottom: 20rpx;
}

.asset-meta {
  display: flex;
  gap: 30rpx;
}

.meta-item {
  font-size: 24rpx;
  color: #666;
  padding: 10rpx 20rpx;
  background: #f5f5f5;
  border-radius: 8rpx;
}

/* 兑换详情 */
.exchange-details {
  background: white;
  border-radius: 20rpx;
  padding: 30rpx;
  margin-bottom: 30rpx;
  box-shadow: 0 20rpx 60rpx rgba(0, 0, 0, 0.15);
  animation: fadeInUp 0.6s ease-out 0.3s both;
}

.details-title {
  display: block;
  font-size: 32rpx;
  font-weight: bold;
  color: #333;
  margin-bottom: 30rpx;
  padding-bottom: 20rpx;
  border-bottom: 2rpx solid #f0f0f0;
}

.detail-row {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 20rpx;
}

.detail-row:last-child {
  margin-bottom: 0;
}

.detail-label {
  font-size: 28rpx;
  color: #666;
}

.detail-value {
  font-size: 28rpx;
  color: #333;
  font-weight: 500;
}

.detail-value.points {
  color: #ff4d4f;
  font-weight: bold;
}

.detail-value.small {
  font-size: 24rpx;
  color: #999;
}

/* 操作按钮 */
.actions {
  display: flex;
  flex-direction: column;
  gap: 20rpx;
  margin-bottom: 40rpx;
  animation: fadeInUp 0.6s ease-out 0.4s both;
}

.action-btn {
  height: 88rpx;
  border-radius: 44rpx;
  display: flex;
  align-items: center;
  justify-content: center;
  border: none;
  box-shadow: 0 10rpx 30rpx rgba(0, 0, 0, 0.1);
}

.action-btn.primary {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
}

.action-btn.secondary {
  background: white;
}

.action-btn.tertiary {
  background: rgba(255, 255, 255, 0.3);
  backdrop-filter: blur(10px);
}

.action-btn.primary .btn-text {
  color: white;
}

.action-btn.secondary .btn-text {
  color: #667eea;
}

.action-btn.tertiary .btn-text {
  color: white;
}

.btn-text {
  font-size: 32rpx;
  font-weight: 500;
}

/* 温馨提示 */
.tips {
  background: rgba(255, 255, 255, 0.2);
  backdrop-filter: blur(10px);
  border-radius: 20rpx;
  padding: 30rpx;
  animation: fadeInUp 0.6s ease-out 0.5s both;
}

.tips-title {
  display: block;
  font-size: 28rpx;
  font-weight: bold;
  color: white;
  margin-bottom: 20rpx;
}

.tips-text {
  display: block;
  font-size: 24rpx;
  color: rgba(255, 255, 255, 0.9);
  line-height: 40rpx;
  margin-bottom: 10rpx;
}

.tips-text:last-child {
  margin-bottom: 0;
}
</style>
