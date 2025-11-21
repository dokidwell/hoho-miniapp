<template>
  <view class="announcement-detail">
    <view v-if="loading" class="loading">
      <text>加载中...</text>
    </view>
    
    <view v-else-if="announcement" class="content">
      <!-- 标题 -->
      <view class="header">
        <view v-if="announcement.pinned" class="pinned-badge">置顶</view>
        <text class="title">{{ announcement.title }}</text>
        <view class="meta">
          <text class="time">{{ formatTime(announcement.published_at) }}</text>
          <text class="views">{{ announcement.views || 0 }}次浏览</text>
        </view>
      </view>
      
      <!-- 内容 -->
      <view class="body">
        <rich-text :nodes="announcement.content"></rich-text>
      </view>
      
      <!-- 附件 -->
      <view v-if="announcement.attachments && announcement.attachments.length > 0" class="attachments">
        <text class="section-title">附件</text>
        <view 
          v-for="(attachment, index) in announcement.attachments" 
          :key="index"
          class="attachment-item"
          @click="downloadAttachment(attachment)"
        >
          <image src="/static/images/file-icon.png" mode="aspectFit"></image>
          <text>{{ attachment.name }}</text>
        </view>
      </view>
      
      <!-- 相关链接 -->
      <view v-if="announcement.links && announcement.links.length > 0" class="links">
        <text class="section-title">相关链接</text>
        <view 
          v-for="(link, index) in announcement.links" 
          :key="index"
          class="link-item"
          @click="openLink(link.url)"
        >
          <text>{{ link.title }}</text>
          <text class="arrow">›</text>
        </view>
      </view>
    </view>
    
    <view v-else class="error">
      <text>公告不存在或已删除</text>
    </view>
  </view>
</template>

<script>
import { getAnnouncementDetail, markAnnouncementAsRead } from '@/api/announcement.js'

export default {
  data() {
    return {
      id: null,
      announcement: null,
      loading: true
    }
  },
  
  onLoad(options) {
    this.id = options.id
    this.loadDetail()
  },
  
  onShareAppMessage() {
    return {
      title: this.announcement?.title || '公告详情',
      path: `/pages/announcement-detail/index?id=${this.id}`
    }
  },
  
  methods: {
    async loadDetail() {
      this.loading = true
      try {
        const res = await getAnnouncementDetail(this.id)
        this.announcement = res.announcement
        
        // 标记为已读
        await markAnnouncementAsRead(this.id)
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
      return `${date.getFullYear()}-${String(date.getMonth() + 1).padStart(2, '0')}-${String(date.getDate()).padStart(2, '0')} ${String(date.getHours()).padStart(2, '0')}:${String(date.getMinutes()).padStart(2, '0')}`
    },
    
    downloadAttachment(attachment) {
      uni.downloadFile({
        url: attachment.url,
        success: (res) => {
          if (res.statusCode === 200) {
            uni.showToast({
              title: '下载成功',
              icon: 'success'
            })
          }
        },
        fail: () => {
          uni.showToast({
            title: '下载失败',
            icon: 'none'
          })
        }
      })
    },
    
    openLink(url) {
      // 如果是小程序内部页面
      if (url.startsWith('/pages')) {
        uni.navigateTo({ url })
      } else {
        // 外部链接，复制到剪贴板
        uni.setClipboardData({
          data: url,
          success: () => {
            uni.showToast({
              title: '链接已复制',
              icon: 'success'
            })
          }
        })
      }
    }
  }
}
</script>

<style scoped>
.announcement-detail {
  min-height: 100vh;
  background: #fff;
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

.content {
  padding: 32rpx;
}

.header {
  padding-bottom: 32rpx;
  border-bottom: 1px solid #f0f0f0;
  margin-bottom: 32rpx;
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

.title {
  display: block;
  font-size: 40rpx;
  font-weight: 600;
  line-height: 1.5;
  margin-bottom: 16rpx;
}

.meta {
  display: flex;
  gap: 32rpx;
  font-size: 24rpx;
  color: #999;
}

.body {
  font-size: 28rpx;
  line-height: 1.8;
  color: #333;
  margin-bottom: 48rpx;
}

.section-title {
  display: block;
  font-size: 32rpx;
  font-weight: 600;
  margin-bottom: 24rpx;
}

.attachments {
  margin-bottom: 48rpx;
}

.attachment-item {
  display: flex;
  align-items: center;
  padding: 24rpx;
  background: #f5f5f5;
  border-radius: 12rpx;
  margin-bottom: 16rpx;
}

.attachment-item image {
  width: 48rpx;
  height: 48rpx;
  margin-right: 16rpx;
}

.attachment-item text {
  flex: 1;
  font-size: 28rpx;
  color: #333;
}

.links {
  margin-bottom: 48rpx;
}

.link-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 24rpx 0;
  border-bottom: 1px solid #f0f0f0;
}

.link-item text {
  font-size: 28rpx;
  color: #1890ff;
}

.arrow {
  font-size: 32rpx;
  color: #999;
}
</style>
