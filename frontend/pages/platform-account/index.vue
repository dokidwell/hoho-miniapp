<template>
  <view class="platform-account">
    <!-- 顶部总览 -->
    <view class="overview">
      <text class="title">阳光账户</text>
      <text class="subtitle">平台收入透明公示</text>
      
      <view class="balance-card">
        <text class="label">当前余额</text>
        <text class="balance">{{ account.balance || '0.00' }}</text>
        <text class="unit">积分</text>
      </view>
      
      <view class="stats">
        <view class="stat-item">
          <text class="value">{{ account.total_income || '0.00' }}</text>
          <text class="label">累计收入</text>
        </view>
        <view class="stat-item">
          <text class="value">{{ account.total_expense || '0.00' }}</text>
          <text class="label">累计支出</text>
        </view>
        <view class="stat-item">
          <text class="value">{{ transactions.length }}</text>
          <text class="label">交易笔数</text>
        </view>
      </view>
    </view>

    <!-- Tab切换 -->
    <view class="tabs">
      <view 
        v-for="tab in tabs" 
        :key="tab.key"
        class="tab-item"
        :class="{ active: currentTab === tab.key }"
        @click="switchTab(tab.key)"
      >
        {{ tab.label }}
      </view>
    </view>

    <!-- 交易记录 -->
    <view class="transaction-list">
      <view v-if="filteredTransactions.length === 0" class="empty">
        <image src="/static/images/empty.png" mode="aspectFit"></image>
        <text>暂无交易记录</text>
      </view>
      
      <view v-for="transaction in filteredTransactions" :key="transaction.id" class="transaction-card">
        <view class="left">
          <text class="type">{{ getTypeText(transaction.type) }}</text>
          <text class="time">{{ formatTime(transaction.created_at) }}</text>
          <text v-if="transaction.description" class="desc">{{ transaction.description }}</text>
        </view>
        
        <view class="right">
          <text 
            class="amount" 
            :class="transaction.type === 'income' ? 'income' : 'expense'"
          >
            {{ transaction.type === 'income' ? '+' : '-' }}{{ transaction.amount }}
          </text>
        </view>
      </view>
    </view>
  </view>
</template>

<script>
import { getPlatformAccount, getPlatformTransactions } from '@/api/platform.js'

export default {
  data() {
    return {
      account: {},
      transactions: [],
      currentTab: 'all',
      tabs: [
        { key: 'all', label: '全部' },
        { key: 'income', label: '收入' },
        { key: 'expense', label: '支出' }
      ],
      loading: false
    }
  },
  
  computed: {
    filteredTransactions() {
      if (this.currentTab === 'all') {
        return this.transactions
      }
      return this.transactions.filter(t => t.type === this.currentTab)
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
    async loadData() {
      this.loading = true
      try {
        const [accountRes, transactionsRes] = await Promise.all([
          getPlatformAccount(),
          getPlatformTransactions({ page: 1, page_size: 100 })
        ])
        
        this.account = accountRes.account || {}
        this.transactions = transactionsRes.transactions || []
      } catch (error) {
        uni.showToast({
          title: error.message || '加载失败',
          icon: 'none'
        })
      } finally {
        this.loading = false
      }
    },
    
    switchTab(key) {
      this.currentTab = key
    },
    
    getTypeText(type) {
      const typeMap = {
        creation_commission: '创作分成',
        transaction_fee: '交易手续费',
        airdrop: '空投支出',
        task_reward: '任务奖励支出',
        refund: '退款支出',
        other_income: '其他收入',
        other_expense: '其他支出'
      }
      return typeMap[type] || type
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
.platform-account {
  min-height: 100vh;
  background: #f5f5f5;
}

.overview {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  padding: 48rpx 32rpx;
  color: #fff;
}

.title {
  display: block;
  font-size: 40rpx;
  font-weight: 600;
  margin-bottom: 8rpx;
}

.subtitle {
  display: block;
  font-size: 24rpx;
  opacity: 0.8;
  margin-bottom: 48rpx;
}

.balance-card {
  background: rgba(255, 255, 255, 0.2);
  border-radius: 16rpx;
  padding: 32rpx;
  text-align: center;
  margin-bottom: 32rpx;
}

.balance-card .label {
  display: block;
  font-size: 24rpx;
  opacity: 0.8;
  margin-bottom: 16rpx;
}

.balance {
  display: block;
  font-size: 72rpx;
  font-weight: 600;
  margin-bottom: 8rpx;
}

.unit {
  display: block;
  font-size: 24rpx;
  opacity: 0.8;
}

.stats {
  display: flex;
  justify-content: space-around;
}

.stat-item {
  text-align: center;
}

.stat-item .value {
  display: block;
  font-size: 32rpx;
  font-weight: 600;
  margin-bottom: 8rpx;
}

.stat-item .label {
  display: block;
  font-size: 24rpx;
  opacity: 0.8;
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

.transaction-list {
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

.transaction-card {
  display: flex;
  justify-content: space-between;
  align-items: center;
  background: #fff;
  border-radius: 16rpx;
  padding: 32rpx;
  margin-bottom: 16rpx;
}

.left {
  flex: 1;
}

.type {
  display: block;
  font-size: 32rpx;
  font-weight: 600;
  margin-bottom: 8rpx;
}

.time {
  display: block;
  font-size: 24rpx;
  color: #999;
  margin-bottom: 8rpx;
}

.desc {
  display: block;
  font-size: 24rpx;
  color: #666;
}

.amount {
  font-size: 36rpx;
  font-weight: 600;
}

.amount.income {
  color: #52c41a;
}

.amount.expense {
  color: #ff4d4f;
}
</style>
