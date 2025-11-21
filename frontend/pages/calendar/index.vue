<template>
  <view class="calendar-page">
    <!-- 顶部标题 -->
    <view class="header">
      <text class="title">发售日历</text>
      <text class="subtitle">即将发售的精彩作品</text>
    </view>

    <!-- 日历视图 -->
    <view class="calendar-view">
      <view class="month-selector">
        <text class="prev" @click="prevMonth">‹</text>
        <text class="current-month">{{ currentYear }}年{{ currentMonth }}月</text>
        <text class="next" @click="nextMonth">›</text>
      </view>
      
      <view class="calendar-grid">
        <view v-for="day in weekDays" :key="day" class="week-day">{{ day }}</view>
        <view 
          v-for="(date, index) in calendarDates" 
          :key="index"
          class="date-cell"
          :class="{ 
            'other-month': date.otherMonth,
            'today': date.isToday,
            'has-event': date.hasEvent
          }"
          @click="selectDate(date)"
        >
          <text class="date-number">{{ date.day }}</text>
          <view v-if="date.hasEvent" class="event-dot"></view>
        </view>
      </view>
    </view>

    <!-- 发售列表 -->
    <view class="release-list">
      <view class="list-header">
        <text class="date-text">{{ selectedDateText }}</text>
        <text class="count-text">{{ filteredReleases.length }}个作品</text>
      </view>
      
      <view v-if="filteredReleases.length === 0" class="empty">
        <image src="/static/images/empty.png" mode="aspectFit"></image>
        <text>该日期暂无发售计划</text>
      </view>
      
      <view v-for="release in filteredReleases" :key="release.id" class="release-card" @click="goToDetail(release)">
        <image :src="release.image_url" mode="aspectFill" class="cover"></image>
        
        <view class="content">
          <view class="top">
            <text class="name">{{ release.name }}</text>
            <view v-if="release.is_limited" class="limited-badge">限量</view>
          </view>
          
          <text class="creator">创作者：{{ release.creator_name }}</text>
          
          <view class="info-row">
            <view class="info-item">
              <text class="label">发行数量</text>
              <text class="value">{{ release.total_supply }}份</text>
            </view>
            <view class="info-item">
              <text class="label">价格</text>
              <text class="value price">{{ release.price }}积分</text>
            </view>
          </view>
          
          <view class="time-info">
            <view class="countdown" v-if="release.countdown > 0">
              <text class="label">距离发售还有</text>
              <text class="time">{{ formatCountdown(release.countdown) }}</text>
            </view>
            <view v-else class="on-sale">
              <text>正在发售中</text>
            </view>
          </view>
          
          <view class="actions">
            <button v-if="release.can_subscribe" class="subscribe-btn" @click.stop="subscribe(release.id)">
              {{ release.subscribed ? '已预约' : '预约提醒' }}
            </button>
            <button v-else class="buy-btn" @click.stop="goToDetail(release)">
              立即购买
            </button>
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
      currentYear: new Date().getFullYear(),
      currentMonth: new Date().getMonth() + 1,
      selectedDate: new Date(),
      weekDays: ['日', '一', '二', '三', '四', '五', '六'],
      calendarDates: [],
      releases: [],
      loading: false
    }
  },
  
  computed: {
    selectedDateText() {
      const date = this.selectedDate
      return `${date.getFullYear()}年${date.getMonth() + 1}月${date.getDate()}日`
    },
    
    filteredReleases() {
      const selectedDateStr = this.formatDate(this.selectedDate)
      return this.releases.filter(r => {
        const releaseDate = this.formatDate(new Date(r.publish_time))
        return releaseDate === selectedDateStr
      })
    }
  },
  
  onLoad() {
    this.generateCalendar()
    this.loadReleases()
  },
  
  methods: {
    generateCalendar() {
      const year = this.currentYear
      const month = this.currentMonth
      const firstDay = new Date(year, month - 1, 1).getDay()
      const daysInMonth = new Date(year, month, 0).getDate()
      const daysInPrevMonth = new Date(year, month - 1, 0).getDate()
      
      const dates = []
      
      // 上个月的日期
      for (let i = firstDay - 1; i >= 0; i--) {
        dates.push({
          day: daysInPrevMonth - i,
          otherMonth: true,
          hasEvent: false
        })
      }
      
      // 当前月的日期
      for (let i = 1; i <= daysInMonth; i++) {
        const date = new Date(year, month - 1, i)
        dates.push({
          day: i,
          otherMonth: false,
          isToday: this.isToday(date),
          hasEvent: this.hasEvent(date),
          fullDate: date
        })
      }
      
      // 下个月的日期
      const remainingDays = 42 - dates.length
      for (let i = 1; i <= remainingDays; i++) {
        dates.push({
          day: i,
          otherMonth: true,
          hasEvent: false
        })
      }
      
      this.calendarDates = dates
    },
    
    isToday(date) {
      const today = new Date()
      return date.getFullYear() === today.getFullYear() &&
             date.getMonth() === today.getMonth() &&
             date.getDate() === today.getDate()
    },
    
    hasEvent(date) {
      const dateStr = this.formatDate(date)
      return this.releases.some(r => {
        const releaseDate = this.formatDate(new Date(r.publish_time))
        return releaseDate === dateStr
      })
    },
    
    formatDate(date) {
      return `${date.getFullYear()}-${String(date.getMonth() + 1).padStart(2, '0')}-${String(date.getDate()).padStart(2, '0')}`
    },
    
    selectDate(date) {
      if (date.otherMonth) return
      this.selectedDate = date.fullDate
    },
    
    prevMonth() {
      if (this.currentMonth === 1) {
        this.currentYear--
        this.currentMonth = 12
      } else {
        this.currentMonth--
      }
      this.generateCalendar()
    },
    
    nextMonth() {
      if (this.currentMonth === 12) {
        this.currentYear++
        this.currentMonth = 1
      } else {
        this.currentMonth++
      }
      this.generateCalendar()
    },
    
    async loadReleases() {
      this.loading = true
      try {
        // 这里调用API获取发售计划
        // const res = await getUpcomingReleases()
        // this.releases = res.releases || []
        
        // 模拟数据
        this.releases = [
          {
            id: 1,
            name: '大鱼HOHO·限量版',
            creator_name: 'HOHO官方',
            image_url: '/static/images/artwork1.jpg',
            total_supply: 200,
            price: '99.99',
            publish_time: new Date().toISOString(),
            is_limited: true,
            can_subscribe: false,
            subscribed: false,
            countdown: 0
          }
        ]
        
        this.generateCalendar()
      } catch (error) {
        uni.showToast({
          title: error.message || '加载失败',
          icon: 'none'
        })
      } finally {
        this.loading = false
      }
    },
    
    formatCountdown(seconds) {
      const days = Math.floor(seconds / 86400)
      const hours = Math.floor((seconds % 86400) / 3600)
      const minutes = Math.floor((seconds % 3600) / 60)
      
      if (days > 0) {
        return `${days}天${hours}小时`
      } else if (hours > 0) {
        return `${hours}小时${minutes}分钟`
      } else {
        return `${minutes}分钟`
      }
    },
    
    async subscribe(id) {
      try {
        // await subscribeRelease(id)
        uni.showToast({
          title: '预约成功',
          icon: 'success'
        })
        this.loadReleases()
      } catch (error) {
        uni.showToast({
          title: error.message || '预约失败',
          icon: 'none'
        })
      }
    },
    
    goToDetail(release) {
      uni.navigateTo({
        url: `/pages/artwork-detail/index?id=${release.id}`
      })
    }
  }
}
</script>

<style scoped>
.calendar-page {
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

.calendar-view {
  background: #fff;
  margin: 24rpx;
  border-radius: 16rpx;
  padding: 32rpx;
}

.month-selector {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 32rpx;
}

.prev,
.next {
  font-size: 48rpx;
  color: #666;
  padding: 0 24rpx;
}

.current-month {
  font-size: 32rpx;
  font-weight: 600;
}

.calendar-grid {
  display: grid;
  grid-template-columns: repeat(7, 1fr);
  gap: 16rpx;
}

.week-day {
  text-align: center;
  font-size: 24rpx;
  color: #999;
  padding: 16rpx 0;
}

.date-cell {
  aspect-ratio: 1;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  position: relative;
  border-radius: 12rpx;
}

.date-cell.other-month .date-number {
  color: #ccc;
}

.date-cell.today {
  background: #f0f0f0;
}

.date-cell.has-event {
  background: #e6f7ff;
}

.date-number {
  font-size: 28rpx;
}

.event-dot {
  width: 8rpx;
  height: 8rpx;
  background: #1890ff;
  border-radius: 50%;
  position: absolute;
  bottom: 8rpx;
}

.release-list {
  padding: 24rpx;
}

.list-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 24rpx;
}

.date-text {
  font-size: 32rpx;
  font-weight: 600;
}

.count-text {
  font-size: 24rpx;
  color: #999;
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

.release-card {
  background: #fff;
  border-radius: 16rpx;
  overflow: hidden;
  margin-bottom: 24rpx;
}

.cover {
  width: 100%;
  height: 400rpx;
}

.content {
  padding: 32rpx;
}

.top {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-bottom: 16rpx;
}

.name {
  font-size: 32rpx;
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

.creator {
  display: block;
  font-size: 24rpx;
  color: #999;
  margin-bottom: 24rpx;
}

.info-row {
  display: flex;
  gap: 48rpx;
  margin-bottom: 24rpx;
}

.info-item {
  flex: 1;
}

.label {
  display: block;
  font-size: 24rpx;
  color: #999;
  margin-bottom: 8rpx;
}

.value {
  display: block;
  font-size: 28rpx;
  font-weight: 600;
}

.value.price {
  color: #ff4d4f;
}

.time-info {
  margin-bottom: 24rpx;
}

.countdown {
  display: flex;
  align-items: center;
  gap: 16rpx;
}

.countdown .label {
  font-size: 24rpx;
  color: #666;
}

.countdown .time {
  font-size: 28rpx;
  font-weight: 600;
  color: #1890ff;
}

.on-sale text {
  font-size: 28rpx;
  color: #52c41a;
  font-weight: 600;
}

.actions {
  display: flex;
  gap: 24rpx;
}

.actions button {
  flex: 1;
  height: 80rpx;
  line-height: 80rpx;
  border-radius: 40rpx;
  font-size: 28rpx;
  border: none;
}

.subscribe-btn {
  background: #f5f5f5;
  color: #666;
}

.buy-btn {
  background: #000;
  color: #fff;
}
</style>
