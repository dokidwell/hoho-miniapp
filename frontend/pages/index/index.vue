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

    <!-- 藏品卡片列表 -->
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

    <!-- 底部导航栏 -->
    <view class="tabbar safe-area-bottom">
      <view class="tabbar-item tabbar-active">
        <image class="tabbar-icon" src="/static/icons/tab-works-active.png" mode="aspectFit" />
        <text class="tabbar-label">作品</text>
      </view>
      <view class="tabbar-item" @click="navigateTo('/pages/create/index')">
        <image class="tabbar-icon" src="/static/icons/tab-create.png" mode="aspectFit" />
        <text class="tabbar-label">创作</text>
      </view>
      <view class="tabbar-item" @click="navigateTo('/pages/jijhuan/index')">
        <image class="tabbar-icon" src="/static/icons/tab-exchange.png" mode="aspectFit" />
        <text class="tabbar-label">集换</text>
      </view>
      <view class="tabbar-item" @click="navigateTo('/pages/ecology/index')">
        <image class="tabbar-icon" src="/static/icons/tab-ecology.png" mode="aspectFit" />
        <text class="tabbar-label">生态</text>
      </view>
      <view class="tabbar-item" @click="navigateTo('/pages/profile/index')">
        <image class="tabbar-icon" src="/static/icons/tab-profile.png" mode="aspectFit" />
        <text class="tabbar-label">我的</text>
      </view>
    </view>
  </view>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useAssetStore } from '@/stores/asset'

const assetStore = useAssetStore()

const assets = ref([
  {
    id: 1,
    image_url: '/static/images/asset-1.png',
    name: '藏品1'
  },
  {
    id: 2,
    image_url: '/static/images/asset-2.png',
    name: '藏品2'
  },
  {
    id: 3,
    image_url: '/static/images/asset-3.png',
    name: '藏品3'
  }
])

const handleCardClick = (asset) => {
  uni.navigateTo({
    url: `/pages/asset-detail/index?id=${asset.id}`
  })
}

const navigateTo = (url) => {
  uni.switchTab({ url })
}

onMounted(async () => {
  // TODO: 从API加载藏品列表
  // const res = await assetStore.fetchAssets()
  // assets.value = res.list
})
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

/* 藏品卡片列表 */
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

/* 底部导航栏 */
.tabbar {
  position: fixed;
  bottom: 0;
  left: 0;
  right: 0;
  display: flex;
  align-items: center;
  justify-content: space-around;
  height: 120rpx;
  background-color: var(--color-bg-primary);
  border-top: 1px solid var(--color-border);
  z-index: 1000;
  
  .tabbar-item {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    gap: 8rpx;
    flex: 1;
    height: 100%;
    color: var(--color-text-tertiary);
    
    .tabbar-icon {
      width: 48rpx;
      height: 48rpx;
    }
    
    .tabbar-label {
      font-size: var(--font-size-xs);
    }
    
    &.tabbar-active {
      color: var(--color-text-primary);
      font-weight: 500;
    }
  }
}
</style>
