<template>
  <view class="artwork-detail">
    <view v-if="loading" class="loading">
      <text>加载中...</text>
    </view>
    
    <view v-else-if="artwork" class="content">
      <!-- 作品图片 -->
      <view class="artwork-image">
        <image :src="artwork.image_url" mode="aspectFill"></image>
      </view>
      
      <!-- 作品信息 -->
      <view class="artwork-info">
        <view class="header">
          <text class="name">{{ artwork.name }}</text>
          <view v-if="artwork.is_limited" class="limited-badge">限量</view>
        </view>
        
        <text class="description">{{ artwork.description }}</text>
        
        <!-- 创作者信息 -->
        <view class="creator-section">
          <text class="section-title">创作者</text>
          <view class="creator-info">
            <image :src="artwork.creator_avatar || '/static/images/default-avatar.png'" class="avatar"></image>
            <view class="creator-detail">
              <text class="creator-name">{{ artwork.creator_name }}</text>
              <text class="creator-desc">{{ artwork.creator_bio || '暂无简介' }}</text>
            </view>
          </view>
        </view>
        
        <!-- 发行信息 -->
        <view class="issue-section">
          <text class="section-title">发行信息</text>
          <view class="info-grid">
            <view class="info-item">
              <text class="label">发行数量</text>
              <text class="value">{{ artwork.total_supply }}份</text>
            </view>
            <view class="info-item">
              <text class="label">发行价格</text>
              <text class="value price">{{ artwork.price }}积分</text>
            </view>
            <view class="info-item">
              <text class="label">发行时间</text>
              <text class="value">{{ formatTime(artwork.publish_time) }}</text>
            </view>
            <view class="info-item">
              <text class="label">已售出</text>
              <text class="value">{{ artwork.sold_count || 0 }}份</text>
            </view>
          </view>
        </view>
        
        <!-- 作品属性 -->
        <view v-if="artwork.attributes && artwork.attributes.length > 0" class="attributes-section">
          <text class="section-title">作品属性</text>
          <view class="attributes-grid">
            <view v-for="attr in artwork.attributes" :key="attr.trait_type" class="attribute-item">
              <text class="attr-type">{{ attr.trait_type }}</text>
              <text class="attr-value">{{ attr.value }}</text>
            </view>
          </view>
        </view>
        
        <!-- 版权说明 -->
        <view class="rights-section">
          <text class="section-title">版权说明</text>
          <text class="rights-text">
            本作品的版权归创作者所有。购买者获得该作品的数字收藏权，可以在平台内进行展示和交易，但不得用于商业用途或进行二次创作。
          </text>
        </view>
        
        <!-- 交易历史 -->
        <view v-if="transactions && transactions.length > 0" class="transactions-section">
          <text class="section-title">交易历史</text>
          <view v-for="tx in transactions" :key="tx.id" class="transaction-item">
            <view class="tx-info">
              <text class="tx-type">{{ tx.type === 'primary' ? '首发' : '转售' }}</text>
              <text class="tx-time">{{ formatTime(tx.created_at) }}</text>
            </view>
            <view class="tx-detail">
              <text class="tx-from">{{ tx.from_user }}</text>
              <text class="tx-arrow">→</text>
              <text class="tx-to">{{ tx.to_user }}</text>
            </view>
            <text class="tx-price">{{ tx.price }}积分</text>
          </view>
        </view>
      </view>
      
      <!-- 底部操作栏 -->
      <view class="action-bar">
        <button v-if="artwork.can_buy" class="buy-btn" @click="buyArtwork">
          立即购买 {{ artwork.price }}积分
        </button>
        <button v-else class="sold-out-btn" disabled>
          已售罄
        </button>
      </view>
    </view>
    
    <view v-else class="error">
      <text>作品不存在或已下架</text>
    </view>
  </view>
</template>

<script>
export default {
  data() {
    return {
      id: null,
      artwork: null,
      transactions: [],
      loading: true
    }
  },
  
  onLoad(options) {
    this.id = options.id
    this.loadDetail()
  },
  
  onShareAppMessage() {
    return {
      title: this.artwork?.name || '作品详情',
      path: `/pages/whitepaper-artwork/index?id=${this.id}`,
      imageUrl: this.artwork?.image_url
    }
  },
  
  methods: {
    async loadDetail() {
      this.loading = true
      try {
        // TODO: 调用API获取作品详情
        // const res = await getArtworkDetail(this.id)
        // this.artwork = res.artwork
        // this.transactions = res.transactions || []
        
        // 模拟数据
        this.artwork = {
          id: this.id,
          name: '大鱼HOHO·限量版',
          description: '这是HOHO Park平台的首发限量作品，由官方团队精心创作。作品以大鱼为主题，寓意着在数字艺术的海洋中自由遨游。',
          image_url: '/static/images/artwork1.jpg',
          creator_name: 'HOHO官方',
          creator_avatar: '/static/images/hoho-avatar.png',
          creator_bio: 'HOHO Park官方创作团队',
          total_supply: 200,
          sold_count: 150,
          price: '99.99',
          publish_time: '2024-11-01T10:00:00Z',
          is_limited: true,
          can_buy: true,
          attributes: [
            { trait_type: '稀有度', value: '传说' },
            { trait_type: '系列', value: '创世系列' },
            { trait_type: '编号', value: '#001' }
          ]
        }
        
        this.transactions = [
          {
            id: 1,
            type: 'primary',
            from_user: '平台',
            to_user: '用户A',
            price: '99.99',
            created_at: '2024-11-01T10:05:00Z'
          }
        ]
      } catch (error) {
        uni.showToast({
          title: error.message || '加载失败',
          icon: 'none'
        })
      } finally {
        this.loading = false
      }
    },
    
    formatTime(time) {
      if (!time) return ''
      const date = new Date(time)
      return `${date.getFullYear()}-${String(date.getMonth() + 1).padStart(2, '0')}-${String(date.getDate()).padStart(2, '0')}`
    },
    
    async buyArtwork() {
      try {
        // TODO: 调用购买接口
        uni.showModal({
          title: '确认购买',
          content: `确定花费${this.artwork.price}积分购买该作品吗？`,
          success: (res) => {
            if (res.confirm) {
              // 执行购买逻辑
              uni.showToast({
                title: '购买成功',
                icon: 'success'
              })
            }
          }
        })
      } catch (error) {
        uni.showToast({
          title: error.message || '购买失败',
          icon: 'none'
        })
      }
    }
  }
}
</script>

<style scoped>
.artwork-detail {
  min-height: 100vh;
  background: #f5f5f5;
  padding-bottom: 120rpx;
}

.loading,
.error {
  display: flex;
  align-items: center;
  justify-content: center;
  height: 100vh;
  font-size: 28rpx;
  color: #999;
}

.artwork-image {
  width: 100%;
  height: 750rpx;
  background: #fff;
}

.artwork-image image {
  width: 100%;
  height: 100%;
}

.artwork-info {
  padding: 32rpx;
}

.header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-bottom: 24rpx;
}

.name {
  font-size: 40rpx;
  font-weight: 600;
  flex: 1;
}

.limited-badge {
  background: #ff4d4f;
  color: #fff;
  font-size: 20rpx;
  padding: 4rpx 12rpx;
  border-radius: 8rpx;
  margin-left: 16rpx;
}

.description {
  display: block;
  font-size: 28rpx;
  line-height: 1.6;
  color: #666;
  margin-bottom: 48rpx;
}

.section-title {
  display: block;
  font-size: 32rpx;
  font-weight: 600;
  margin-bottom: 24rpx;
}

.creator-section,
.issue-section,
.attributes-section,
.rights-section,
.transactions-section {
  background: #fff;
  border-radius: 16rpx;
  padding: 32rpx;
  margin-bottom: 24rpx;
}

.creator-info {
  display: flex;
  align-items: center;
  gap: 24rpx;
}

.avatar {
  width: 96rpx;
  height: 96rpx;
  border-radius: 50%;
}

.creator-detail {
  flex: 1;
}

.creator-name {
  display: block;
  font-size: 32rpx;
  font-weight: 600;
  margin-bottom: 8rpx;
}

.creator-desc {
  display: block;
  font-size: 24rpx;
  color: #999;
}

.info-grid {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 32rpx;
}

.info-item {
  display: flex;
  flex-direction: column;
}

.label {
  font-size: 24rpx;
  color: #999;
  margin-bottom: 8rpx;
}

.value {
  font-size: 28rpx;
  font-weight: 600;
}

.value.price {
  color: #ff4d4f;
}

.attributes-grid {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 16rpx;
}

.attribute-item {
  background: #f5f5f5;
  border-radius: 12rpx;
  padding: 24rpx;
  text-align: center;
}

.attr-type {
  display: block;
  font-size: 24rpx;
  color: #999;
  margin-bottom: 8rpx;
}

.attr-value {
  display: block;
  font-size: 28rpx;
  font-weight: 600;
}

.rights-text {
  display: block;
  font-size: 24rpx;
  line-height: 1.6;
  color: #666;
}

.transaction-item {
  padding: 24rpx 0;
  border-bottom: 1px solid #f0f0f0;
}

.transaction-item:last-child {
  border-bottom: none;
}

.tx-info {
  display: flex;
  justify-content: space-between;
  margin-bottom: 16rpx;
}

.tx-type {
  font-size: 24rpx;
  color: #1890ff;
  background: #e6f7ff;
  padding: 4rpx 12rpx;
  border-radius: 8rpx;
}

.tx-time {
  font-size: 24rpx;
  color: #999;
}

.tx-detail {
  display: flex;
  align-items: center;
  gap: 16rpx;
  margin-bottom: 8rpx;
}

.tx-from,
.tx-to {
  font-size: 28rpx;
  color: #333;
}

.tx-arrow {
  font-size: 24rpx;
  color: #999;
}

.tx-price {
  display: block;
  font-size: 32rpx;
  font-weight: 600;
  color: #ff4d4f;
}

.action-bar {
  position: fixed;
  bottom: 0;
  left: 0;
  right: 0;
  background: #fff;
  padding: 24rpx 32rpx;
  padding-bottom: calc(24rpx + env(safe-area-inset-bottom));
  box-shadow: 0 -2rpx 16rpx rgba(0, 0, 0, 0.05);
}

.buy-btn,
.sold-out-btn {
  width: 100%;
  height: 96rpx;
  line-height: 96rpx;
  border-radius: 48rpx;
  font-size: 32rpx;
  font-weight: 600;
  border: none;
}

.buy-btn {
  background: #000;
  color: #fff;
}

.sold-out-btn {
  background: #f5f5f5;
  color: #999;
}
</style>
