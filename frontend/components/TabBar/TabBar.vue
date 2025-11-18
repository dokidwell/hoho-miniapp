<template>
  <view class="tabbar safe-area-bottom">
    <view 
      class="tabbar-item" 
      :class="{ 'tabbar-active': active === 0 }"
      @click="switchTab('/pages/index/index')"
    >
      <view class="tabbar-icon-wrapper">
        <view class="tabbar-icon" :class="{ 'active': active === 0 }">📱</view>
      </view>
      <text class="tabbar-label">作品</text>
    </view>
    
    <view 
      class="tabbar-item" 
      :class="{ 'tabbar-active': active === 1 }"
      @click="switchTab('/pages/create/index')"
    >
      <view class="tabbar-icon-wrapper">
        <view class="tabbar-icon" :class="{ 'active': active === 1 }">✨</view>
      </view>
      <text class="tabbar-label">创作</text>
    </view>
    
    <view 
      class="tabbar-item" 
      :class="{ 'tabbar-active': active === 2 }"
      @click="switchTab('/pages/jijhuan/index')"
    >
      <view class="tabbar-icon-wrapper">
        <view class="tabbar-icon" :class="{ 'active': active === 2 }">🔄</view>
      </view>
      <text class="tabbar-label">集换</text>
    </view>
    
    <view 
      class="tabbar-item" 
      :class="{ 'tabbar-active': active === 3 }"
      @click="switchTab('/pages/ecology/index')"
    >
      <view class="tabbar-icon-wrapper">
        <view class="tabbar-icon" :class="{ 'active': active === 3 }">🌱</view>
      </view>
      <text class="tabbar-label">生态</text>
    </view>
    
    <view 
      class="tabbar-item" 
      :class="{ 'tabbar-active': active === 4 }"
      @click="switchTab('/pages/profile/index')"
    >
      <view class="tabbar-icon-wrapper">
        <view class="tabbar-icon" :class="{ 'active': active === 4 }">👤</view>
      </view>
      <text class="tabbar-label">我的</text>
    </view>
  </view>
</template>

<script setup>
const props = defineProps({
  active: {
    type: Number,
    default: 0
  }
})

const switchTab = (url) => {
  uni.switchTab({ url })
}
</script>

<style lang="scss" scoped>
.tabbar {
  position: fixed;
  bottom: 0;
  left: 0;
  right: 0;
  display: flex;
  align-items: center;
  justify-content: space-around;
  height: 120rpx;
  background-color: #FFFFFF;
  border-top: 1px solid #E8E8E8;
  z-index: 1000;
  
  .tabbar-item {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    gap: 8rpx;
    flex: 1;
    height: 100%;
    color: #999999;
    transition: all 0.3s ease;
    
    .tabbar-icon-wrapper {
      display: flex;
      align-items: center;
      justify-content: center;
      
      .tabbar-icon {
        font-size: 48rpx;
        opacity: 0.6;
        transition: all 0.3s ease;
        
        &.active {
          opacity: 1;
          transform: scale(1.1);
        }
      }
    }
    
    .tabbar-label {
      font-size: 20rpx;
      transition: all 0.3s ease;
    }
    
    &.tabbar-active {
      color: #000000;
      font-weight: 500;
      
      .tabbar-label {
        font-weight: 600;
      }
    }
  }
}

.safe-area-bottom {
  padding-bottom: constant(safe-area-inset-bottom);
  padding-bottom: env(safe-area-inset-bottom);
}
</style>
