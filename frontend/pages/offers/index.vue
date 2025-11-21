<template>
  <view class="offers-page">
    <!-- 顶部Tab切换 -->
    <view class="tabs">
      <view 
        v-for="tab in tabs" 
        :key="tab.key"
        class="tab-item"
        :class="{ active: currentTab === tab.key }"
        @click="switchTab(tab.key)"
      >
        {{ tab.label }}
        <view v-if="tab.count > 0" class="badge">{{ tab.count }}</view>
      </view>
    </view>

    <!-- 我的出价列表 -->
    <view v-if="currentTab === 'my'" class="offer-list">
      <view v-if="myOffers.length === 0" class="empty">
        <image src="/static/images/empty.png" mode="aspectFit"></image>
        <text>暂无出价记录</text>
      </view>
      
      <view v-for="offer in myOffers" :key="offer.id" class="offer-card">
        <view class="artwork-info">
          <image :src="offer.artwork_instance.artwork.image_url" mode="aspectFill"></image>
          <view class="info">
            <text class="name">{{ offer.artwork_instance.artwork.name }}</text>
            <text class="serial">#{{ offer.artwork_instance.serial_number }}</text>
          </view>
        </view>
        
        <view class="offer-info">
          <view class="price-row">
            <text class="label">出价金额</text>
            <text class="price">{{ offer.price }} 积分</text>
          </view>
          <view class="time-row">
            <text class="label">出价时间</text>
            <text class="time">{{ formatTime(offer.created_at) }}</text>
          </view>
          <view class="status-row">
            <text class="label">状态</text>
            <text class="status" :class="offer.status">{{ getStatusText(offer.status) }}</text>
          </view>
          <view v-if="offer.status === 'pending'" class="expire-row">
            <text class="label">有效期至</text>
            <text class="expire">{{ formatTime(offer.expires_at) }}</text>
          </view>
        </view>
        
        <view v-if="offer.status === 'pending'" class="actions">
          <button class="cancel-btn" @click="cancelOffer(offer.id)">取消出价</button>
        </view>
      </view>
    </view>

    <!-- 收到的出价列表 -->
    <view v-if="currentTab === 'received'" class="offer-list">
      <view v-if="receivedOffers.length === 0" class="empty">
        <image src="/static/images/empty.png" mode="aspectFit"></image>
        <text>暂无收到的出价</text>
      </view>
      
      <view v-for="offer in receivedOffers" :key="offer.id" class="offer-card">
        <view class="buyer-info">
          <image :src="offer.buyer.avatar" mode="aspectFill"></image>
          <view class="info">
            <text class="name">{{ offer.buyer.nickname }}</text>
            <text class="time">{{ formatTime(offer.created_at) }}</text>
          </view>
        </view>
        
        <view class="artwork-info">
          <image :src="offer.artwork_instance.artwork.image_url" mode="aspectFill"></image>
          <view class="info">
            <text class="name">{{ offer.artwork_instance.artwork.name }}</text>
            <text class="serial">#{{ offer.artwork_instance.serial_number }}</text>
          </view>
        </view>
        
        <view class="offer-info">
          <view class="price-row">
            <text class="label">出价金额</text>
            <text class="price">{{ offer.price }} 积分</text>
          </view>
          <view v-if="offer.message" class="message-row">
            <text class="label">留言</text>
            <text class="message">{{ offer.message }}</text>
          </view>
          <view class="status-row">
            <text class="label">状态</text>
            <text class="status" :class="offer.status">{{ getStatusText(offer.status) }}</text>
          </view>
        </view>
        
        <view v-if="offer.status === 'pending'" class="actions">
          <button class="reject-btn" @click="rejectOffer(offer.id)">拒绝</button>
          <button class="accept-btn" @click="acceptOffer(offer.id)">接受</button>
        </view>
      </view>
    </view>
  </view>
</template>

<script>
import { getMyOffers, getReceivedOffers, acceptOffer, rejectOffer, cancelOffer } from '@/api/offer.js'

export default {
  data() {
    return {
      currentTab: 'my',
      tabs: [
        { key: 'my', label: '我的出价', count: 0 },
        { key: 'received', label: '收到的出价', count: 0 }
      ],
      myOffers: [],
      receivedOffers: [],
      loading: false
    }
  },
  
  onLoad() {
    this.loadData()
  },
  
  onPullDownRefresh() {
    this.loadData().then(() => {
      uni.stopPullDownRefresh()
    })
  },
  
  methods: {
    switchTab(key) {
      this.currentTab = key
      this.loadData()
    },
    
    async loadData() {
      this.loading = true
      try {
        if (this.currentTab === 'my') {
          const res = await getMyOffers({ status: 'all' })
          this.myOffers = res.offers || []
          this.tabs[0].count = this.myOffers.filter(o => o.status === 'pending').length
        } else {
          const res = await getReceivedOffers({ status: 'all' })
          this.receivedOffers = res.offers || []
          this.tabs[1].count = this.receivedOffers.filter(o => o.status === 'pending').length
        }
      } catch (error) {
        uni.showToast({
          title: error.message || '加载失败',
          icon: 'none'
        })
      } finally {
        this.loading = false
      }
    },
    
    async acceptOffer(id) {
      uni.showModal({
        title: '确认接受',
        content: '接受出价后将自动完成交易，确定要接受吗？',
        success: async (res) => {
          if (res.confirm) {
            try {
              await acceptOffer(id)
              uni.showToast({
                title: '已接受出价',
                icon: 'success'
              })
              this.loadData()
            } catch (error) {
              uni.showToast({
                title: error.message || '操作失败',
                icon: 'none'
              })
            }
          }
        }
      })
    },
    
    async rejectOffer(id) {
      uni.showModal({
        title: '确认拒绝',
        content: '确定要拒绝这个出价吗？',
        success: async (res) => {
          if (res.confirm) {
            try {
              await rejectOffer(id, { reason: '不符合预期' })
              uni.showToast({
                title: '已拒绝出价',
                icon: 'success'
              })
              this.loadData()
            } catch (error) {
              uni.showToast({
                title: error.message || '操作失败',
                icon: 'none'
              })
            }
          }
        }
      })
    },
    
    async cancelOffer(id) {
      uni.showModal({
        title: '确认取消',
        content: '取消出价后将解冻积分，确定要取消吗？',
        success: async (res) => {
          if (res.confirm) {
            try {
              await cancelOffer(id)
              uni.showToast({
                title: '已取消出价',
                icon: 'success'
              })
              this.loadData()
            } catch (error) {
              uni.showToast({
                title: error.message || '操作失败',
                icon: 'none'
              })
            }
          }
        }
      })
    },
    
    getStatusText(status) {
      const statusMap = {
        pending: '待处理',
        accepted: '已接受',
        rejected: '已拒绝',
        cancelled: '已取消',
        expired: '已过期'
      }
      return statusMap[status] || status
    },
    
    formatTime(time) {
      if (!time) return ''
      const date = new Date(time)
      return `${date.getFullYear()}-${String(date.getMonth() + 1).padStart(2, '0')}-${String(date.getDate()).padStart(2, '0')} ${String(date.getHours()).padStart(2, '0')}:${String(date.getMinutes()).padStart(2, '0')}`
    }
  }
}
</script>

<style scoped>
.offers-page {
  min-height: 100vh;
  background: #f5f5f5;
}

.tabs {
  display: flex;
  background: #fff;
  padding: 0 32rpx;
  position: sticky;
  top: 0;
  z-index: 10;
}

.tab-item {
  flex: 1;
  text-align: center;
  padding: 32rpx 0;
  font-size: 28rpx;
  color: #666;
  position: relative;
}

.tab-item.active {
  color: #000;
  font-weight: 600;
}

.tab-item.active::after {
  content: '';
  position: absolute;
  bottom: 0;
  left: 50%;
  transform: translateX(-50%);
  width: 60rpx;
  height: 6rpx;
  background: #000;
  border-radius: 3rpx;
}

.badge {
  position: absolute;
  top: 20rpx;
  right: 20%;
  background: #ff4d4f;
  color: #fff;
  font-size: 20rpx;
  padding: 2rpx 8rpx;
  border-radius: 20rpx;
  min-width: 32rpx;
  text-align: center;
}

.offer-list {
  padding: 24rpx;
}

.empty {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 120rpx 0;
}

.empty image {
  width: 200rpx;
  height: 200rpx;
  margin-bottom: 24rpx;
}

.empty text {
  font-size: 28rpx;
  color: #999;
}

.offer-card {
  background: #fff;
  border-radius: 16rpx;
  padding: 32rpx;
  margin-bottom: 24rpx;
}

.buyer-info,
.artwork-info {
  display: flex;
  align-items: center;
  margin-bottom: 24rpx;
}

.buyer-info image {
  width: 80rpx;
  height: 80rpx;
  border-radius: 40rpx;
  margin-right: 24rpx;
}

.artwork-info image {
  width: 120rpx;
  height: 120rpx;
  border-radius: 12rpx;
  margin-right: 24rpx;
}

.info {
  flex: 1;
}

.info .name {
  display: block;
  font-size: 32rpx;
  font-weight: 600;
  margin-bottom: 8rpx;
}

.info .serial,
.info .time {
  display: block;
  font-size: 24rpx;
  color: #999;
}

.offer-info {
  padding: 24rpx 0;
  border-top: 1px solid #f0f0f0;
}

.price-row,
.time-row,
.status-row,
.expire-row,
.message-row {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 16rpx;
}

.label {
  font-size: 28rpx;
  color: #666;
}

.price {
  font-size: 32rpx;
  font-weight: 600;
  color: #ff4d4f;
}

.time,
.expire {
  font-size: 24rpx;
  color: #999;
}

.status {
  font-size: 28rpx;
  padding: 4rpx 16rpx;
  border-radius: 8rpx;
}

.status.pending {
  color: #1890ff;
  background: #e6f7ff;
}

.status.accepted {
  color: #52c41a;
  background: #f6ffed;
}

.status.rejected,
.status.cancelled,
.status.expired {
  color: #999;
  background: #f5f5f5;
}

.message {
  font-size: 28rpx;
  color: #333;
  text-align: right;
  max-width: 400rpx;
}

.actions {
  display: flex;
  gap: 24rpx;
  margin-top: 24rpx;
}

.actions button {
  flex: 1;
  height: 80rpx;
  line-height: 80rpx;
  border-radius: 40rpx;
  font-size: 28rpx;
  border: none;
}

.cancel-btn,
.reject-btn {
  background: #f5f5f5;
  color: #666;
}

.accept-btn {
  background: #000;
  color: #fff;
}
</style>
