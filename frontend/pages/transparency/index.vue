<template>
  <view class="transparency-page">
    <!-- 顶部标题 -->
    <view class="header">
      <text class="title">透明公示</text>
      <text class="subtitle">所有数据公开透明，接受社区监督</text>
    </view>

    <!-- Tab切换 -->
    <scroll-view class="tabs" scroll-x>
      <view 
        v-for="tab in tabs" 
        :key="tab.key"
        class="tab-item"
        :class="{ active: currentTab === tab.key }"
        @click="switchTab(tab.key)"
      >
        {{ tab.label }}
      </view>
    </scroll-view>

    <!-- 内容区域 -->
    <view class="content">
      <!-- Tab 1: 交易记录 -->
      <view v-if="currentTab === 'transactions'" class="tab-content">
        <view class="filter-bar">
          <picker mode="date" @change="onDateChange">
            <view class="filter-btn">
              <text>{{ filterDate || '选择日期' }}</text>
            </view>
          </picker>
        </view>
        
        <view v-if="transactions.length === 0" class="empty">
          <image src="/static/images/empty.png" mode="aspectFit"></image>
          <text>暂无交易记录</text>
        </view>
        
        <view v-for="item in transactions" :key="item.id" class="list-card">
          <view class="card-header">
            <text class="card-title">{{ item.artwork_name }}</text>
            <text class="card-amount">{{ item.price }}积分</text>
          </view>
          <view class="card-row">
            <text class="label">买家：</text>
            <text class="value">{{ item.buyer_name }}</text>
          </view>
          <view class="card-row">
            <text class="label">卖家：</text>
            <text class="value">{{ item.seller_name }}</text>
          </view>
          <view class="card-row">
            <text class="label">交易时间：</text>
            <text class="value">{{ formatTime(item.created_at) }}</text>
          </view>
          <view class="card-row">
            <text class="label">平台手续费：</text>
            <text class="value highlight">{{ item.platform_fee }}积分</text>
          </view>
        </view>
      </view>

      <!-- Tab 2: 积分流水 -->
      <view v-if="currentTab === 'points'" class="tab-content">
        <view class="summary-card">
          <view class="summary-item">
            <text class="summary-label">总发行量</text>
            <text class="summary-value">{{ pointsSummary.total_issued || '0.00' }}</text>
          </view>
          <view class="summary-item">
            <text class="summary-label">总流通量</text>
            <text class="summary-value">{{ pointsSummary.total_circulation || '0.00' }}</text>
          </view>
          <view class="summary-item">
            <text class="summary-label">总销毁量</text>
            <text class="summary-value">{{ pointsSummary.total_burned || '0.00' }}</text>
          </view>
        </view>
        
        <view v-for="item in pointRecords" :key="item.id" class="list-card">
          <view class="card-header">
            <text class="card-title">{{ getPointTypeText(item.type) }}</text>
            <text class="card-amount" :class="item.amount > 0 ? 'positive' : 'negative'">
              {{ item.amount > 0 ? '+' : '' }}{{ item.amount }}
            </text>
          </view>
          <view class="card-row">
            <text class="label">用户：</text>
            <text class="value">{{ item.user_name }}</text>
          </view>
          <view class="card-row">
            <text class="label">时间：</text>
            <text class="value">{{ formatTime(item.created_at) }}</text>
          </view>
          <view v-if="item.description" class="card-row">
            <text class="label">说明：</text>
            <text class="value">{{ item.description }}</text>
          </view>
        </view>
      </view>

      <!-- Tab 3: 转赠记录 -->
      <view v-if="currentTab === 'transfers'" class="tab-content">
        <view v-if="transfers.length === 0" class="empty">
          <image src="/static/images/empty.png" mode="aspectFit"></image>
          <text>暂无转赠记录</text>
        </view>
        
        <view v-for="item in transfers" :key="item.id" class="list-card">
          <view class="card-header">
            <text class="card-title">{{ item.artwork_name }}</text>
            <text class="card-badge">转赠</text>
          </view>
          <view class="card-row">
            <text class="label">转出方：</text>
            <text class="value">{{ item.from_user_name }}</text>
          </view>
          <view class="card-row">
            <text class="label">接收方：</text>
            <text class="value">{{ item.to_user_name }}</text>
          </view>
          <view class="card-row">
            <text class="label">转赠时间：</text>
            <text class="value">{{ formatTime(item.created_at) }}</text>
          </view>
        </view>
      </view>

      <!-- Tab 4: 规则变更 -->
      <view v-if="currentTab === 'rules'" class="tab-content">
        <view v-if="ruleChanges.length === 0" class="empty">
          <image src="/static/images/empty.png" mode="aspectFit"></image>
          <text>暂无规则变更记录</text>
        </view>
        
        <view v-for="item in ruleChanges" :key="item.id" class="list-card">
          <view class="card-header">
            <text class="card-title">{{ item.title }}</text>
            <text class="card-badge important">重要</text>
          </view>
          <view class="card-row">
            <text class="label">变更类型：</text>
            <text class="value">{{ getRuleTypeText(item.type) }}</text>
          </view>
          <view class="card-row">
            <text class="label">生效时间：</text>
            <text class="value">{{ formatTime(item.effective_at) }}</text>
          </view>
          <view class="card-content">
            <text class="content-title">变更内容：</text>
            <rich-text :nodes="item.content"></rich-text>
          </view>
        </view>
      </view>

      <!-- Tab 5: 审核记录 -->
      <view v-if="currentTab === 'reviews'" class="tab-content">
        <view v-if="reviews.length === 0" class="empty">
          <image src="/static/images/empty.png" mode="aspectFit"></image>
          <text>暂无审核记录</text>
        </view>
        
        <view v-for="item in reviews" :key="item.id" class="list-card">
          <view class="card-header">
            <text class="card-title">{{ item.creation_name }}</text>
            <view class="status-badge" :class="item.status">
              {{ getReviewStatusText(item.status) }}
            </view>
          </view>
          <view class="card-row">
            <text class="label">创作者：</text>
            <text class="value">{{ item.creator_name }}</text>
          </view>
          <view class="card-row">
            <text class="label">提交时间：</text>
            <text class="value">{{ formatTime(item.submitted_at) }}</text>
          </view>
          <view class="card-row">
            <text class="label">审核时间：</text>
            <text class="value">{{ formatTime(item.reviewed_at) }}</text>
          </view>
          <view v-if="item.review_note" class="card-content">
            <text class="content-title">审核意见：</text>
            <text>{{ item.review_note }}</text>
          </view>
        </view>
      </view>

      <!-- Tab 6: 公告记录 -->
      <view v-if="currentTab === 'announcements'" class="tab-content">
        <view v-if="announcements.length === 0" class="empty">
          <image src="/static/images/empty.png" mode="aspectFit"></image>
          <text>暂无公告</text>
        </view>
        
        <view 
          v-for="item in announcements" 
          :key="item.id" 
          class="announcement-card"
          @click="goToAnnouncement(item.id)"
        >
          <view v-if="item.pinned" class="pinned-badge">置顶</view>
          <text class="announcement-title">{{ item.title }}</text>
          <text class="announcement-summary">{{ item.summary }}</text>
          <view class="announcement-footer">
            <text class="time">{{ formatTime(item.published_at) }}</text>
            <text class="arrow">›</text>
          </view>
        </view>
      </view>
    </view>
  </view>
</template>

<script>
export default {
  data() {
    return {
      currentTab: 'transactions',
      tabs: [
        { key: 'transactions', label: '交易记录' },
        { key: 'points', label: '积分流水' },
        { key: 'transfers', label: '转赠记录' },
        { key: 'rules', label: '规则变更' },
        { key: 'reviews', label: '审核记录' },
        { key: 'announcements', label: '公告记录' }
      ],
      filterDate: '',
      transactions: [],
      pointRecords: [],
      pointsSummary: {},
      transfers: [],
      ruleChanges: [],
      reviews: [],
      announcements: [],
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
        // 根据当前tab加载对应数据
        switch (this.currentTab) {
          case 'transactions':
            await this.loadTransactions()
            break
          case 'points':
            await this.loadPointRecords()
            break
          case 'transfers':
            await this.loadTransfers()
            break
          case 'rules':
            await this.loadRuleChanges()
            break
          case 'reviews':
            await this.loadReviews()
            break
          case 'announcements':
            await this.loadAnnouncements()
            break
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
    
    async loadTransactions() {
      // TODO: 调用API
      this.transactions = []
    },
    
    async loadPointRecords() {
      // TODO: 调用API
      this.pointRecords = []
      this.pointsSummary = {
        total_issued: '1000000.00',
        total_circulation: '800000.00',
        total_burned: '50000.00'
      }
    },
    
    async loadTransfers() {
      // TODO: 调用API
      this.transfers = []
    },
    
    async loadRuleChanges() {
      // TODO: 调用API
      this.ruleChanges = []
    },
    
    async loadReviews() {
      // TODO: 调用API
      this.reviews = []
    },
    
    async loadAnnouncements() {
      // TODO: 调用API
      this.announcements = []
    },
    
    onDateChange(e) {
      this.filterDate = e.detail.value
      this.loadData()
    },
    
    getPointTypeText(type) {
      const typeMap = {
        airdrop: '空投',
        task_reward: '任务奖励',
        purchase: '购买消耗',
        sale: '出售获得',
        refund: '退款',
        transfer: '转账'
      }
      return typeMap[type] || type
    },
    
    getRuleTypeText(type) {
      const typeMap = {
        commission: '分成比例调整',
        fee: '手续费调整',
        task: '任务规则调整',
        creation: '创作规则调整',
        other: '其他'
      }
      return typeMap[type] || type
    },
    
    getReviewStatusText(status) {
      const statusMap = {
        pending: '待审核',
        approved: '已通过',
        rejected: '已拒绝'
      }
      return statusMap[status] || status
    },
    
    formatTime(time) {
      if (!time) return ''
      const date = new Date(time)
      return `${date.getFullYear()}-${String(date.getMonth() + 1).padStart(2, '0')}-${String(date.getDate()).padStart(2, '0')} ${String(date.getHours()).padStart(2, '0')}:${String(date.getMinutes()).padStart(2, '0')}`
    },
    
    goToAnnouncement(id) {
      uni.navigateTo({
        url: `/pages/announcement-detail/index?id=${id}`
      })
    }
  }
}
</script>

<style scoped>
.transparency-page {
  min-height: 100vh;
  background: #f5f5f5;
}

.header {
  background: #fff;
  padding: 48rpx 32rpx 32rpx;
  text-align: center;
}

.title {
  display: block;
  font-size: 40rpx;
  font-weight: 600;
  margin-bottom: 16rpx;
}

.subtitle {
  display: block;
  font-size: 24rpx;
  color: #999;
}

.tabs {
  display: flex;
  background: #fff;
  padding: 0 32rpx;
  white-space: nowrap;
  border-bottom: 1px solid #f0f0f0;
}

.tab-item {
  display: inline-block;
  padding: 32rpx 24rpx;
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

.content {
  padding: 24rpx;
}

.filter-bar {
  margin-bottom: 24rpx;
}

.filter-btn {
  background: #fff;
  padding: 24rpx 32rpx;
  border-radius: 12rpx;
  font-size: 28rpx;
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

.summary-card {
  background: #fff;
  border-radius: 16rpx;
  padding: 32rpx;
  margin-bottom: 24rpx;
  display: flex;
  justify-content: space-around;
}

.summary-item {
  text-align: center;
}

.summary-label {
  display: block;
  font-size: 24rpx;
  color: #999;
  margin-bottom: 16rpx;
}

.summary-value {
  display: block;
  font-size: 32rpx;
  font-weight: 600;
}

.list-card {
  background: #fff;
  border-radius: 16rpx;
  padding: 32rpx;
  margin-bottom: 16rpx;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 24rpx;
}

.card-title {
  font-size: 32rpx;
  font-weight: 600;
  flex: 1;
}

.card-amount {
  font-size: 32rpx;
  font-weight: 600;
  color: #ff4d4f;
}

.card-amount.positive {
  color: #52c41a;
}

.card-amount.negative {
  color: #ff4d4f;
}

.card-badge {
  background: #f0f0f0;
  color: #666;
  font-size: 20rpx;
  padding: 4rpx 12rpx;
  border-radius: 8rpx;
}

.card-badge.important {
  background: #ff4d4f;
  color: #fff;
}

.status-badge {
  font-size: 24rpx;
  padding: 8rpx 16rpx;
  border-radius: 8rpx;
}

.status-badge.pending {
  background: #fff7e6;
  color: #fa8c16;
}

.status-badge.approved {
  background: #f6ffed;
  color: #52c41a;
}

.status-badge.rejected {
  background: #fff1f0;
  color: #ff4d4f;
}

.card-row {
  display: flex;
  margin-bottom: 16rpx;
  font-size: 28rpx;
}

.card-row:last-child {
  margin-bottom: 0;
}

.label {
  color: #999;
  width: 160rpx;
}

.value {
  flex: 1;
  color: #333;
}

.value.highlight {
  color: #ff4d4f;
  font-weight: 600;
}

.card-content {
  margin-top: 24rpx;
  padding-top: 24rpx;
  border-top: 1px solid #f0f0f0;
  font-size: 28rpx;
  line-height: 1.6;
}

.content-title {
  display: block;
  color: #999;
  margin-bottom: 16rpx;
}

.announcement-card {
  background: #fff;
  border-radius: 16rpx;
  padding: 32rpx;
  margin-bottom: 16rpx;
}

.pinned-badge {
  display: inline-block;
  background: #ff4d4f;
  color: #fff;
  font-size: 20rpx;
  padding: 4rpx 12rpx;
  border-radius: 8rpx;
  margin-bottom: 16rpx;
}

.announcement-title {
  display: block;
  font-size: 32rpx;
  font-weight: 600;
  margin-bottom: 16rpx;
}

.announcement-summary {
  display: block;
  font-size: 28rpx;
  color: #666;
  line-height: 1.6;
  margin-bottom: 24rpx;
  overflow: hidden;
  text-overflow: ellipsis;
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
}

.announcement-footer {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.time {
  font-size: 24rpx;
  color: #999;
}

.arrow {
  font-size: 32rpx;
  color: #999;
}
</style>
