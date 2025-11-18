<template>
  <view class="page-container">
    <!-- 自定义导航栏 -->
    <view class="navbar">
      <view class="navbar-left">
        <uni-icons type="left" size="24" color="#333" @click="goBack"></uni-icons>
      </view>
      <view class="navbar-title">透明公示</view>
      <view class="navbar-right">
        <uni-icons type="more-filled" size="24" color="#333"></uni-icons>
        <view class="filter-toggle">
          <uni-icons type="circle" size="20" color="#3a8fff"></uni-icons>
        </view>
      </view>
    </view>

    <!-- 页面描述 -->
    <view class="description">
      <text>你可在此查询所有事件、流转、积分使用情况。</text>
    </view>

    <!-- 标签页 -->
    <view class="tabs">
      <view
        class="tab-item"
        :class="{ active: activeTab === 'events' }"
        @click="activeTab = 'events'"
      >
        事件列表
      </view>
      <view
        class="tab-item"
        :class="{ active: activeTab === 'points' }"
        @click="activeTab = 'points'"
      >
        积分
      </view>
      <view
        class="tab-item"
        :class="{ active: activeTab === 'transfer' }"
        @click="activeTab = 'transfer'"
      >
        流转情况
      </view>
    </view>

    <!-- 事件列表 -->
    <view class="events-list" v-if="activeTab === 'events'">
      <view class="event-item" v-for="event in events" :key="event.id">
        <view class="event-marker">•</view>
        <view class="event-content">
          <view class="event-title">{{ event.title }}</view>
          <view class="event-time">{{ event.created_at }}</view>
        </view>
      </view>
    </view>

    <!-- 积分流转 -->
    <view class="points-list" v-if="activeTab === 'points'">
      <view class="points-item" v-for="point in pointsHistory" :key="point.id">
        <view class="points-type">{{ point.type }}</view>
        <view class="points-amount" :class="{ positive: point.type === 'earn' }">
          {{ point.amount }}
        </view>
        <view class="points-time">{{ point.created_at }}</view>
      </view>
    </view>

    <!-- 流转情况 -->
    <view class="transfer-list" v-if="activeTab === 'transfer'">
      <view class="transfer-item" v-for="transfer in transfers" :key="transfer.id">
        <view class="transfer-from">{{ transfer.from }}</view>
        <view class="transfer-arrow">→</view>
        <view class="transfer-to">{{ transfer.to }}</view>
        <view class="transfer-amount">{{ transfer.amount }}</view>
      </view>
    </view>

    <!-- 加载更多 -->
    <view class="load-more" v-if="!loading && hasMore">
      <text @click="loadMore">加载更多</text>
    </view>
  </view>
</template>

<script setup>
import { ref, onMounted } from 'vue'

const activeTab = ref('events')
const events = ref([
  {
    id: 1,
    title: '用户uid678910 兑换了 作品野生HOHO',
    created_at: '2025-10-14 15:00:23'
  },
  {
    id: 2,
    title: '用户uid678910 发起了 提案：最大制作卡片份数',
    created_at: '2025-10-14 15:00:23'
  },
  {
    id: 3,
    title: '用户uid678910 登录了 HOHOPark',
    created_at: '2025-10-14 15:00:23'
  }
])

const pointsHistory = ref([
  {
    id: 1,
    type: '获得',
    amount: '+100.00000000',
    created_at: '2025-10-14 15:00:23'
  },
  {
    id: 2,
    type: '消费',
    amount: '-50.00000000',
    created_at: '2025-10-14 14:30:00'
  }
])

const transfers = ref([
  {
    id: 1,
    from: '平台',
    to: '用户A',
    amount: '100.00000000'
  },
  {
    id: 2,
    from: '用户A',
    to: '用户B',
    amount: '50.00000000'
  }
])

const loading = ref(false)
const hasMore = ref(true)

// 返回
const goBack = () => {
  uni.navigateBack()
}

// 加载更多
const loadMore = () => {
  if (!loading.value && hasMore.value) {
    loading.value = true
    // TODO: 加载更多数据
    setTimeout(() => {
      loading.value = false
    }, 1000)
  }
}

// 页面加载
onMounted(() => {
  // 初始化数据
})
</script>

<style lang="scss" scoped>
.page-container {
  width: 100%;
  min-height: 100vh;
  background-color: #f5f5f5;
}

// 导航栏
.navbar {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 12px 16px;
  background-color: #ffffff;
  border-bottom: 1px solid #f0f0f0;
}

.navbar-left,
.navbar-right {
  display: flex;
  align-items: center;
  gap: 12px;
}

.navbar-title {
  font-size: 18px;
  font-weight: 600;
  color: #000;
  flex: 1;
  text-align: center;
}

.filter-toggle {
  width: 40px;
  height: 40px;
  display: flex;
  align-items: center;
  justify-content: center;
  background-color: #f0f0f0;
  border-radius: 50%;
}

// 描述
.description {
  padding: 16px;
  background-color: #ffffff;
  margin: 12px;
  border-radius: 12px;
  font-size: 14px;
  color: #666;
  text-align: center;
}

// 标签页
.tabs {
  display: flex;
  padding: 0 12px;
  background-color: #ffffff;
  margin: 12px;
  border-radius: 12px;
  gap: 0;
}

.tab-item {
  flex: 1;
  padding: 12px;
  text-align: center;
  font-size: 14px;
  color: #666;
  border-bottom: 2px solid transparent;
  cursor: pointer;
  transition: all 0.3s ease;

  &.active {
    color: #000;
    border-bottom-color: #000;
    font-weight: 600;
  }
}

// 事件列表
.events-list {
  padding: 12px;
}

.event-item {
  display: flex;
  gap: 12px;
  padding: 12px;
  background-color: #ffffff;
  border-radius: 8px;
  margin-bottom: 8px;
}

.event-marker {
  font-size: 16px;
  color: #3a8fff;
  min-width: 20px;
}

.event-content {
  flex: 1;
}

.event-title {
  font-size: 14px;
  color: #333;
  margin-bottom: 4px;
}

.event-time {
  font-size: 12px;
  color: #999;
}

// 积分列表
.points-list {
  padding: 12px;
}

.points-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 12px;
  background-color: #ffffff;
  border-radius: 8px;
  margin-bottom: 8px;
}

.points-type {
  font-size: 14px;
  color: #333;
  flex: 1;
}

.points-amount {
  font-size: 14px;
  color: #d32f2f;
  font-weight: 600;
  min-width: 120px;
  text-align: right;

  &.positive {
    color: #52c41a;
  }
}

.points-time {
  font-size: 12px;
  color: #999;
  margin-left: 12px;
  min-width: 150px;
  text-align: right;
}

// 流转列表
.transfer-list {
  padding: 12px;
}

.transfer-item {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 12px;
  background-color: #ffffff;
  border-radius: 8px;
  margin-bottom: 8px;
  font-size: 14px;
}

.transfer-from,
.transfer-to {
  flex: 1;
  color: #333;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.transfer-arrow {
  color: #999;
  margin: 0 4px;
}

.transfer-amount {
  color: #3a8fff;
  font-weight: 600;
  min-width: 120px;
  text-align: right;
}

// 加载更多
.load-more {
  text-align: center;
  padding: 20px;

  text {
    color: #3a8fff;
    font-size: 14px;
    cursor: pointer;
  }
}
</style>
