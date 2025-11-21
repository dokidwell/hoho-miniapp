<template>
	<view class="tasks-container">
		<!-- 顶部积分余额 -->
		<view class="points-header">
			<view class="points-info">
				<text class="points-label">我的积分</text>
				<text class="points-value">{{userPoints}}</text>
			</view>
			<view class="signin-btn" @click="dailySignIn" v-if="!todaySignedIn">
				<text>📅 每日签到</text>
			</view>
			<view class="signed-badge" v-else>
				<text>✅ 今日已签到</text>
			</view>
		</view>
		
		<!-- 任务分类 -->
		<view class="task-tabs">
			<view class="tab" :class="{active: activeTab === 'daily'}" @click="switchTab('daily')">
				<text>每日任务</text>
			</view>
			<view class="tab" :class="{active: activeTab === 'once'}" @click="switchTab('once')">
				<text>新手任务</text>
			</view>
			<view class="tab" :class="{active: activeTab === 'achievement'}" @click="switchTab('achievement')">
				<text>成就任务</text>
			</view>
		</view>
		
		<!-- 任务列表 -->
		<view class="task-list">
			<view class="task-item" v-for="task in filteredTasks" :key="task.id">
				<view class="task-icon">{{getTaskIcon(task.condition_type)}}</view>
				<view class="task-content">
					<view class="task-name">{{task.name}}</view>
					<view class="task-desc">{{task.description}}</view>
					<view class="task-reward">
						<text class="reward-icon">💰</text>
						<text class="reward-text">{{task.reward_points}} 积分</text>
					</view>
				</view>
				<view class="task-action">
					<button v-if="!task.completed" class="action-btn" @click="goToTask(task)">去完成</button>
					<button v-else-if="!task.claimed" class="claim-btn" @click="claimReward(task)">领取</button>
					<view v-else class="completed-badge">✅ 已完成</view>
				</view>
			</view>
		</view>
		
		<!-- 空状态 -->
		<view class="empty-state" v-if="filteredTasks.length === 0">
			<text class="empty-icon">📋</text>
			<text class="empty-text">暂无任务</text>
		</view>
	</view>
</template>

<script>
import { getTasks, claimTaskReward, dailySignIn } from '@/api/task'
import { getPointsBalance } from '@/api/points'

export default {
	data() {
		return {
			activeTab: 'daily',
			tasks: [],
			userPoints: '0.00',
			todaySignedIn: false
		}
	},
	computed: {
		filteredTasks() {
			return this.tasks.filter(task => task.type === this.activeTab)
		}
	},
	onLoad() {
		this.loadTasks()
		this.loadPoints()
	},
	methods: {
		async loadTasks() {
			try {
				const res = await getTasks()
				this.tasks = res.tasks
				
				// 检查今日是否已签到
				const signInTask = this.tasks.find(t => t.condition_type === 'daily_signin')
				if (signInTask) {
					this.todaySignedIn = signInTask.completed
				}
			} catch (error) {
				uni.showToast({ title: error.message || '加载失败', icon: 'none' })
			}
		},
		async loadPoints() {
			try {
				const res = await getPointsBalance()
				this.userPoints = res.available_points
			} catch (error) {
				console.error('加载积分失败', error)
			}
		},
		switchTab(tab) {
			this.activeTab = tab
		},
		async dailySignIn() {
			try {
				await dailySignIn()
				uni.showToast({ title: '签到成功', icon: 'success' })
				this.loadTasks()
				this.loadPoints()
			} catch (error) {
				uni.showToast({ title: error.message || '签到失败', icon: 'none' })
			}
		},
		async claimReward(task) {
			try {
				await claimTaskReward(task.completion_id)
				uni.showToast({ title: '领取成功', icon: 'success' })
				this.loadTasks()
				this.loadPoints()
			} catch (error) {
				uni.showToast({ title: error.message || '领取失败', icon: 'none' })
			}
		},
		goToTask(task) {
			// 根据任务类型跳转到相应页面
			const routes = {
				'daily_signin': '/pages/tasks/index',
				'first_creation': '/pages/create/index',
				'first_purchase': '/pages/exchange/index',
				'first_sale': '/pages/my-assets/index',
				'complete_profile': '/pages/profile-edit/index'
			}
			
			const route = routes[task.condition_type]
			if (route) {
				uni.navigateTo({ url: route })
			}
		},
		getTaskIcon(conditionType) {
			const icons = {
				'daily_signin': '📅',
				'first_creation': '🎨',
				'first_purchase': '🛍️',
				'first_sale': '💰',
				'complete_profile': '👤',
				'invite_friend': '👥'
			}
			return icons[conditionType] || '✨'
		}
	}
}
</script>

<style scoped>
.tasks-container {
	min-height: 100vh;
	background: #f5f5f5;
	padding-bottom: 20rpx;
}

.points-header {
	background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
	padding: 40rpx 30rpx;
	display: flex;
	justify-content: space-between;
	align-items: center;
}

.points-info {
	display: flex;
	flex-direction: column;
	gap: 10rpx;
}

.points-label {
	font-size: 24rpx;
	color: rgba(255, 255, 255, 0.8);
}

.points-value {
	font-size: 48rpx;
	font-weight: bold;
	color: white;
}

.signin-btn {
	background: white;
	color: #667eea;
	padding: 15rpx 30rpx;
	border-radius: 30rpx;
	font-size: 28rpx;
	font-weight: bold;
}

.signed-badge {
	background: rgba(255, 255, 255, 0.2);
	color: white;
	padding: 15rpx 30rpx;
	border-radius: 30rpx;
	font-size: 24rpx;
}

.task-tabs {
	display: flex;
	background: white;
	padding: 20rpx 30rpx;
	gap: 20rpx;
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
	background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
	color: white;
	font-weight: bold;
}

.task-list {
	padding: 20rpx 30rpx;
}

.task-item {
	background: white;
	border-radius: 16rpx;
	padding: 30rpx;
	margin-bottom: 20rpx;
	display: flex;
	align-items: center;
	gap: 20rpx;
}

.task-icon {
	font-size: 60rpx;
	width: 100rpx;
	height: 100rpx;
	display: flex;
	align-items: center;
	justify-content: center;
	background: #f5f5f5;
	border-radius: 50%;
}

.task-content {
	flex: 1;
}

.task-name {
	font-size: 32rpx;
	font-weight: bold;
	margin-bottom: 10rpx;
}

.task-desc {
	font-size: 24rpx;
	color: #999;
	margin-bottom: 15rpx;
}

.task-reward {
	display: flex;
	align-items: center;
	gap: 10rpx;
}

.reward-icon {
	font-size: 28rpx;
}

.reward-text {
	font-size: 28rpx;
	color: #FF6B00;
	font-weight: bold;
}

.task-action {
	min-width: 150rpx;
}

.action-btn, .claim-btn {
	padding: 15rpx 30rpx;
	border-radius: 30rpx;
	font-size: 26rpx;
	border: none;
}

.action-btn {
	background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
	color: white;
}

.claim-btn {
	background: #FF6B00;
	color: white;
}

.completed-badge {
	font-size: 24rpx;
	color: #4CD964;
	text-align: center;
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
