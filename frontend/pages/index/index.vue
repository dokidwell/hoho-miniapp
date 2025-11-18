<template>
  <view class="page">
    <!-- Banner区域 -->
    <view class="banner-section">
      <image 
        class="banner-image" 
        src="/static/images/banner-welcome.png" 
        mode="aspectFill"
      />
      <view class="banner-text">欢迎来到HOHO Park!</view>
    </view>

    <!-- 作品卡片列表 -->
    <view class="cards-section">
      <view 
        v-for="(asset, index) in assets" 
        :key="index"
        class="asset-card"
        @click="handleCardClick(asset)"
      >
        <image 
          class="card-image" 
          :src="asset.image_url || '/static/images/placeholder.png'" 
          mode="aspectFill"
        />
        <view class="card-action">
          <view class="action-btn">立即查看</view>
        </view>
      </view>
    </view>

    <!-- 加载状态 -->
    <view v-if="loading" class="loading-wrapper">
      <text class="loading-text">加载中...</text>
    </view>

    <!-- 空状态 -->
    <view v-else-if="assets.length === 0" class="empty-wrapper">
      <text class="empty-emoji">📭</text>
      <text class="empty-text">暂无作品</text>
    </view>

    <!-- 底部导航栏 -->
    <TabBar :active="0" />
  </view>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import TabBar from '@/components/TabBar/TabBar.vue'
import request from '@/api/request'
import { API_ENDPOINTS } from '@/api/config'

const loading = ref(false)
const assets = ref([])
const page = ref(1)
const pageSize = ref(20)

onMounted(() => {
  fetchAssets()
})

// 获取作品列表
async function fetchAssets() {
  loading.value = true
  try {
    const res = await request.get(API_ENDPOINTS.ASSET.LIST, {
      page: page.value,
      page_size: pageSize.value,
      status: 'approved'
    })
    assets.value = res.list || []
  } catch (error) {
    uni.showToast({
      title: error.message || '加载失败',
      icon: 'none'
    })
  } finally {
    loading.value = false
  }
}

const handleCardClick = (asset) => {
  uni.navigateTo({
    url: `/pages/asset-detail/index?id=${asset.id}`
  })
}
</script>

<style lang="scss" scoped>
.page {
  min-height: 100vh;
  background-color: var(--color-bg-secondary);
  padding-bottom: 120rpx;
}

/* Banner区域 */
.banner-section {
  position: relative;
  width: 100%;
  height: 600rpx;
  overflow: hidden;
  
  .banner-image {
    width: 100%;
    height: 100%;
    display: block;
  }
  
  .banner-text {
    position: absolute;
    bottom: 60rpx;
    left: 50%;
    transform: translateX(-50%);
    font-size: var(--font-size-2xl);
    font-weight: 600;
    color: var(--color-text-primary);
    text-align: center;
    white-space: nowrap;
  }
}

/* 作品卡片列表 */
.cards-section {
  padding: var(--spacing-lg);
  display: flex;
  flex-direction: column;
  gap: var(--spacing-lg);
}

.asset-card {
  position: relative;
  width: 100%;
  height: 600rpx;
  border-radius: var(--radius-lg);
  overflow: hidden;
  background-color: var(--color-bg-primary);
  box-shadow: var(--shadow-md);
  
  .card-image {
    width: 100%;
    height: 100%;
    display: block;
  }
  
  .card-action {
    position: absolute;
    bottom: 0;
    left: 0;
    right: 0;
    display: flex;
    align-items: center;
    justify-content: center;
    padding: var(--spacing-lg);
    background: linear-gradient(to top, rgba(0, 0, 0, 0.6), transparent);
    
    .action-btn {
      padding: 16rpx 48rpx;
      background-color: rgba(0, 0, 0, 0.8);
      color: var(--color-text-white);
      font-size: var(--font-size-base);
      font-weight: 500;
      border-radius: var(--radius-full);
      backdrop-filter: blur(10px);
    }
  }
}

/* 加载和空状态 */
.loading-wrapper,
.empty-wrapper {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 120rpx 0;
  
  .loading-text,
  .empty-text {
    font-size: 28rpx;
    color: #999999;
    margin-top: 24rpx;
  }
  
  .empty-emoji {
    font-size: 120rpx;
  }
}
</style>
