<template>
	<view class="announcements-container">
		<!-- 分类筛选 -->
		<view class="filter-tabs">
			<view class="tab" :class="{active: activeType === ''}" @click="switchType('')">
				<text>全部</text>
			</view>
			<view class="tab" :class="{active: activeType === 'system'}" @click="switchType('system')">
				<text>系统</text>
			</view>
			<view class="tab" :class="{active: activeType === 'rule'}" @click="switchType('rule')">
				<text>规则</text>
			</view>
			<view class="tab" :class="{active: activeType === 'event'}" @click="switchType('event')">
				<text>活动</text>
			</view>
		</view>
		
		<!-- 公告列表 -->
		<view class="announcement-list">
			<view class="announcement-item" v-for="item in announcements" :key="item.id" @click="viewDetail(item.id)">
				<view class="item-header">
					<view class="title-row">
						<view class="pinned-badge" v-if="item.is_pinned">📌 置顶</view>
						<view class="priority-badge" :class="'priority-' + item.priority">
							{{getPriorityText(item.priority)}}
						</view>
					</view>
					<text class="title">{{item.title}}</text>
				</view>
				<view class="item-content">
					<text class="content-preview">{{item.content}}</text>
				</view>
				<view class="item-footer">
					<text class="time">{{formatTime(item.published_at)}}</text>
					<text class="type-tag">{{getTypeText(item.type)}}</text>
				</view>
			</view>
		</view>
		
		<!-- 加载更多 -->
		<view class="load-more" v-if="hasMore" @click="loadMore">
			<text>加载更多</text>
		</view>
		
		<!-- 空状态 -->
		<view class="empty-state" v-if="announcements.length === 0 && !loading">
			<text class="empty-icon">📢</text>
			<text class="empty-text">暂无公告</text>
		</view>
	</view>
</template>

<script>
import { getAnnouncements } from '@/api/announcement'

export default {
	data() {
		return {
			activeType: '',
			announcements: [],
			page: 1,
			pageSize: 20,
			hasMore: true,
			loading: false
		}
	},
	onLoad() {
		this.loadAnnouncements()
	},
	methods: {
		async loadAnnouncements(reset = false) {
			if (this.loading) return
			
			if (reset) {
				this.page = 1
				this.announcements = []
				this.hasMore = true
			}
			
			this.loading = true
			try {
				const res = await getAnnouncements({
					type: this.activeType,
					page: this.page,
					page_size: this.pageSize
				})
				
				if (reset) {
					this.announcements = res.announcements
				} else {
					this.announcements.push(...res.announcements)
				}
				
				this.hasMore = this.announcements.length < res.total
			} catch (error) {
				uni.showToast({ title: error.message || '加载失败', icon: 'none' })
			} finally {
				this.loading = false
			}
		},
		switchType(type) {
			this.activeType = type
			this.loadAnnouncements(true)
		},
		loadMore() {
			if (!this.hasMore || this.loading) return
			this.page++
			this.loadAnnouncements()
		},
		viewDetail(id) {
			uni.navigateTo({
				url: `/pages/announcement-detail/index?id=${id}`
			})
		},
		getPriorityText(priority) {
			const texts = {
				'urgent': '紧急',
				'high': '重要',
				'normal': '普通',
				'low': '一般'
			}
			return texts[priority] || '普通'
		},
		getTypeText(type) {
			const texts = {
				'system': '系统公告',
				'rule': '规则变动',
				'event': '活动公告',
				'maintenance': '维护通知'
			}
			return texts[type] || '公告'
		},
		formatTime(time) {
			const date = new Date(time)
			const now = new Date()
			const diff = now - date
			
			if (diff < 60000) return '刚刚'
			if (diff < 3600000) return Math.floor(diff / 60000) + '分钟前'
			if (diff < 86400000) return Math.floor(diff / 3600000) + '小时前'
			if (diff < 604800000) return Math.floor(diff / 86400000) + '天前'
			
			return date.toLocaleDateString()
		}
	}
}
</script>

<style scoped>
.announcements-container {
	min-height: 100vh;
	background: #f5f5f5;
	padding-bottom: 20rpx;
}

.filter-tabs {
	display: flex;
	background: white;
	padding: 20rpx 30rpx;
	gap: 20rpx;
	position: sticky;
	top: 0;
	z-index: 10;
}

.tab {
	flex: 1;
	text-align: center;
	padding: 15rpx 0;
	border-radius: 30rpx;
	font-size: 28rpx;
	color: #666;
	background: #f5f5f5;
}

.tab.active {
	background: #000;
	color: white;
	font-weight: bold;
}

.announcement-list {
	padding: 20rpx 30rpx;
}

.announcement-item {
	background: white;
	border-radius: 16rpx;
	padding: 30rpx;
	margin-bottom: 20rpx;
}

.item-header {
	margin-bottom: 20rpx;
}

.title-row {
	display: flex;
	gap: 15rpx;
	margin-bottom: 15rpx;
}

.pinned-badge {
	font-size: 22rpx;
	color: #FF6B00;
	background: #FFF3E0;
	padding: 5rpx 15rpx;
	border-radius: 20rpx;
}

.priority-badge {
	font-size: 22rpx;
	padding: 5rpx 15rpx;
	border-radius: 20rpx;
}

.priority-urgent {
	background: #FFEBEE;
	color: #F44336;
}

.priority-high {
	background: #FFF3E0;
	color: #FF9800;
}

.priority-normal {
	background: #E3F2FD;
	color: #2196F3;
}

.priority-low {
	background: #F5F5F5;
	color: #999;
}

.title {
	font-size: 32rpx;
	font-weight: bold;
	color: #333;
	line-height: 1.5;
}

.item-content {
	margin-bottom: 20rpx;
}

.content-preview {
	font-size: 28rpx;
	color: #666;
	line-height: 1.6;
	display: -webkit-box;
	-webkit-box-orient: vertical;
	-webkit-line-clamp: 2;
	overflow: hidden;
}

.item-footer {
	display: flex;
	justify-content: space-between;
	align-items: center;
}

.time {
	font-size: 24rpx;
	color: #999;
}

.type-tag {
	font-size: 22rpx;
	color: #666;
	background: #f5f5f5;
	padding: 5rpx 15rpx;
	border-radius: 20rpx;
}

.load-more {
	text-align: center;
	padding: 30rpx;
	font-size: 28rpx;
	color: #666;
}

.empty-state {
	text-align: center;
	padding: 100rpx 0;
}

.empty-icon {
	font-size: 120rpx;
	display: block;
	margin-bottom: 20rpx;
}

.empty-text {
	font-size: 28rpx;
	color: #999;
}
</style>
