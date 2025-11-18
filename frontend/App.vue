<template>
  <view class="app">
    <router-view />
  </view>
</template>

<script setup>
import { onLaunch, onShow, onHide } from '@dcloudio/uni-app'

onLaunch(() => {
  console.log('HOHO小程序启动')
})

onShow(() => {
  console.log('HOHO小程序显示')
})

onHide(() => {
  console.log('HOHO小程序隐藏')
})
</script>

<style lang="scss">
/* 全局重置 */
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

/* 白色主题 - 浅色配色 */
page {
  background-color: #F5F5F5;
  color: #000000;
  font-family: -apple-system, BlinkMacSystemFont, 'SF Pro Display', 'PingFang SC', 'Hiragino Sans GB', 'Microsoft YaHei', sans-serif;
  font-size: 28rpx;
  line-height: 1.6;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}

.app {
  width: 100%;
  min-height: 100vh;
}

/* 颜色变量 - 白色主题 */
:root {
  /* 主色调 */
  --color-primary: #000000;
  --color-primary-light: #333333;
  --color-accent: #0066FF;
  
  /* 功能色 */
  --color-success: #52C41A;
  --color-warning: #FAAD14;
  --color-error: #FF4D4F;
  --color-info: #1890FF;
  
  /* 文本色 */
  --color-text-primary: #000000;
  --color-text-secondary: #666666;
  --color-text-tertiary: #999999;
  --color-text-disabled: #CCCCCC;
  --color-text-white: #FFFFFF;
  
  /* 背景色 */
  --color-bg-primary: #FFFFFF;
  --color-bg-secondary: #F5F5F5;
  --color-bg-tertiary: #FAFAFA;
  --color-bg-hover: #F0F0F0;
  --color-bg-dark: #000000;
  
  /* 边框色 */
  --color-border: #E8E8E8;
  --color-border-light: #F0F0F0;
  --color-border-dark: #D9D9D9;
  
  /* 字体大小 */
  --font-size-xs: 20rpx;
  --font-size-sm: 24rpx;
  --font-size-base: 28rpx;
  --font-size-lg: 32rpx;
  --font-size-xl: 36rpx;
  --font-size-2xl: 40rpx;
  --font-size-3xl: 48rpx;
  
  /* 间距 */
  --spacing-xs: 8rpx;
  --spacing-sm: 16rpx;
  --spacing-md: 24rpx;
  --spacing-lg: 32rpx;
  --spacing-xl: 48rpx;
  --spacing-2xl: 64rpx;
  
  /* 圆角 */
  --radius-sm: 8rpx;
  --radius-md: 12rpx;
  --radius-lg: 16rpx;
  --radius-xl: 24rpx;
  --radius-full: 9999rpx;
  
  /* 阴影 */
  --shadow-sm: 0 2rpx 8rpx rgba(0, 0, 0, 0.04);
  --shadow-md: 0 4rpx 16rpx rgba(0, 0, 0, 0.08);
  --shadow-lg: 0 8rpx 24rpx rgba(0, 0, 0, 0.12);
}

/* 通用布局 */
.container {
  min-height: 100vh;
  background-color: var(--color-bg-secondary);
}

.page-padding {
  padding: var(--spacing-lg);
}

/* 卡片样式 */
.card {
  background-color: var(--color-bg-primary);
  border-radius: var(--radius-lg);
  padding: var(--spacing-lg);
  margin-bottom: var(--spacing-md);
  box-shadow: var(--shadow-sm);
}

.card-hover {
  transition: all 0.3s ease;
  
  &:active {
    transform: scale(0.98);
    opacity: 0.9;
  }
}

/* 按钮样式 */
.btn {
  display: flex;
  align-items: center;
  justify-content: center;
  height: 88rpx;
  border-radius: var(--radius-full);
  font-size: var(--font-size-lg);
  font-weight: 500;
  transition: all 0.3s ease;
  border: none;
  
  &.btn-primary {
    background-color: var(--color-primary);
    color: var(--color-text-white);
    
    &:active {
      opacity: 0.8;
    }
  }
  
  &.btn-secondary {
    background-color: var(--color-bg-primary);
    color: var(--color-text-primary);
    border: 1px solid var(--color-border);
    
    &:active {
      background-color: var(--color-bg-hover);
    }
  }
  
  &.btn-accent {
    background-color: var(--color-accent);
    color: var(--color-text-white);
    
    &:active {
      opacity: 0.8;
    }
  }
  
  &.btn-small {
    height: 64rpx;
    font-size: var(--font-size-base);
    padding: 0 var(--spacing-lg);
  }
  
  &.btn-disabled {
    opacity: 0.5;
    pointer-events: none;
  }
}

/* 文本样式 */
.text-primary { color: var(--color-text-primary); }
.text-secondary { color: var(--color-text-secondary); }
.text-tertiary { color: var(--color-text-tertiary); }
.text-white { color: var(--color-text-white); }
.text-accent { color: var(--color-accent); }
.text-success { color: var(--color-success); }
.text-error { color: var(--color-error); }
.text-warning { color: var(--color-warning); }

.text-bold { font-weight: 600; }
.text-medium { font-weight: 500; }
.text-normal { font-weight: 400; }

.text-xs { font-size: var(--font-size-xs); }
.text-sm { font-size: var(--font-size-sm); }
.text-base { font-size: var(--font-size-base); }
.text-lg { font-size: var(--font-size-lg); }
.text-xl { font-size: var(--font-size-xl); }
.text-2xl { font-size: var(--font-size-2xl); }
.text-3xl { font-size: var(--font-size-3xl); }

/* 数字样式（积分显示 - 8位小数） */
.number-display {
  font-family: 'SF Mono', 'Monaco', 'Consolas', 'Courier New', monospace;
  font-weight: 600;
  letter-spacing: 0.5px;
  font-variant-numeric: tabular-nums;
}

/* 标签样式 */
.tag {
  display: inline-flex;
  align-items: center;
  padding: 4rpx 12rpx;
  border-radius: var(--radius-sm);
  font-size: var(--font-size-xs);
  
  &.tag-success {
    background-color: rgba(82, 196, 26, 0.1);
    color: var(--color-success);
  }
  
  &.tag-error {
    background-color: rgba(255, 77, 79, 0.1);
    color: var(--color-error);
  }
  
  &.tag-warning {
    background-color: rgba(250, 173, 20, 0.1);
    color: var(--color-warning);
  }
  
  &.tag-info {
    background-color: rgba(24, 144, 255, 0.1);
    color: var(--color-info);
  }
  
  &.tag-default {
    background-color: var(--color-bg-secondary);
    color: var(--color-text-secondary);
  }
}

/* 分割线 */
.divider {
  height: 1px;
  background-color: var(--color-border);
  margin: var(--spacing-lg) 0;
}

/* Flex布局工具类 */
.flex { display: flex; }
.flex-col { flex-direction: column; }
.flex-center { align-items: center; justify-content: center; }
.flex-between { justify-content: space-between; align-items: center; }
.flex-around { justify-content: space-around; align-items: center; }
.flex-start { justify-content: flex-start; align-items: center; }
.flex-end { justify-content: flex-end; align-items: center; }
.flex-1 { flex: 1; }
.flex-wrap { flex-wrap: wrap; }

/* 对齐工具类 */
.text-center { text-align: center; }
.text-left { text-align: left; }
.text-right { text-align: right; }

/* 间距工具类 */
.mt-0 { margin-top: 0; }
.mt-xs { margin-top: var(--spacing-xs); }
.mt-sm { margin-top: var(--spacing-sm); }
.mt-md { margin-top: var(--spacing-md); }
.mt-lg { margin-top: var(--spacing-lg); }
.mt-xl { margin-top: var(--spacing-xl); }

.mb-0 { margin-bottom: 0; }
.mb-xs { margin-bottom: var(--spacing-xs); }
.mb-sm { margin-bottom: var(--spacing-sm); }
.mb-md { margin-bottom: var(--spacing-md); }
.mb-lg { margin-bottom: var(--spacing-lg); }
.mb-xl { margin-bottom: var(--spacing-xl); }

.ml-xs { margin-left: var(--spacing-xs); }
.ml-sm { margin-left: var(--spacing-sm); }
.ml-md { margin-left: var(--spacing-md); }
.ml-lg { margin-left: var(--spacing-lg); }

.mr-xs { margin-right: var(--spacing-xs); }
.mr-sm { margin-right: var(--spacing-sm); }
.mr-md { margin-right: var(--spacing-md); }
.mr-lg { margin-right: var(--spacing-lg); }

.p-0 { padding: 0; }
.p-xs { padding: var(--spacing-xs); }
.p-sm { padding: var(--spacing-sm); }
.p-md { padding: var(--spacing-md); }
.p-lg { padding: var(--spacing-lg); }
.p-xl { padding: var(--spacing-xl); }

.px-sm { padding-left: var(--spacing-sm); padding-right: var(--spacing-sm); }
.px-md { padding-left: var(--spacing-md); padding-right: var(--spacing-md); }
.px-lg { padding-left: var(--spacing-lg); padding-right: var(--spacing-lg); }

.py-sm { padding-top: var(--spacing-sm); padding-bottom: var(--spacing-sm); }
.py-md { padding-top: var(--spacing-md); padding-bottom: var(--spacing-md); }
.py-lg { padding-top: var(--spacing-lg); padding-bottom: var(--spacing-lg); }

/* 加载状态 */
.loading {
  display: flex;
  align-items: center;
  justify-content: center;
  padding: var(--spacing-2xl) 0;
  color: var(--color-text-secondary);
}

/* 空状态 */
.empty-state {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 128rpx 0;
  color: var(--color-text-secondary);
  
  .empty-icon {
    font-size: 96rpx;
    margin-bottom: var(--spacing-md);
    opacity: 0.5;
  }
  
  .empty-text {
    font-size: var(--font-size-base);
  }
}

/* 图片样式 */
image {
  display: block;
  width: 100%;
  height: auto;
}

/* 安全区域适配 */
.safe-area-bottom {
  padding-bottom: constant(safe-area-inset-bottom);
  padding-bottom: env(safe-area-inset-bottom);
}

/* 输入框样式 */
input, textarea {
  font-family: inherit;
  font-size: inherit;
  color: inherit;
}

/* 滚动条隐藏 */
::-webkit-scrollbar {
  display: none;
}
</style>
