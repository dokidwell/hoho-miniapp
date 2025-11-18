<template>
  <view class="page">
    <!-- 顶部搜索栏 -->
    <view class="header">
      <view class="search-box">
        <text class="search-icon">🔍</text>
        <input class="search-input" placeholder="搜索..." v-model="keyword" />
      </view>
      <view class="filter-btn" @click="showFilter">
        <text class="filter-icon">⚙️</text>
      </view>
    </view>

    <!-- 集换专区标题 -->
    <view class="section-title">集换专区</view>

    <!-- 分类Tab -->
    <scroll-view class="category-tabs" scroll-x>
      <view 
        class="tab-item" 
        v-for="cat in categories" 
        :key="cat.value"
        :class="{ 'tab-active': currentCategory === cat.value }"
        @click="selectCategory(cat.value)">
        {{ cat.label }}
      </view>
    </scroll-view>

    <!-- 作品网格（2列） -->
    <scroll-view class="asset-grid" scroll-y @scrolltolower="loadMore">
      <!-- 加载中 -->
      <view v-if="loading && listings.length === 0" class="loading-wrapper">
        <text class="loading-text">加载中...</text>
      </view>

      <!-- 作品列表 -->
      <view v-else-if="listings.length > 0" class="grid-wrapper">
        <view v-for="item in listings" :key="item.id" class="asset-card" @click="goToDetail(item.asset.id)">
          <!-- 作品图片 -->
          <view class="asset-image-wrapper">
            <view class="asset-image-placeholder">
              <text class="placeholder-emoji">🖼️</text>
            </view>
            
            <!-- 左上角标签 -->
            <view v-if="item.is_official" class="asset-tag">
              <text class="tag-text">合集作品</text>
            </view>
            
            <!-- 左下角可兑数量 -->
            <view class="asset-available">
              <text class="available-icon">⏰</text>
              <text class="available-text">{{ item.available_count || 0 }}份可兑</text>
            </view>
          </view>

          <!-- 作品信息 -->
          <view class="asset-info">
            <view class="asset-name">{{ item.asset.name }}</view>
            <view class="asset-price-row">
              <text class="price-value number-display">{{ formatPrice(item.price) }}</text>
              <text class="price-unit">起</text>
            </view>
            <view class="asset-supply">
              <text class="supply-icon">📦</text>
              <text class="supply-text">{{ item.asset.total_supply }}份</text>
            </view>
          </view>
        </view>
      </view>

      <!-- 空状态 -->
      <view v-else class="empty-wrapper">
        <text class="empty-emoji">📭</text>
        <text class="empty-text">暂无挂售作品</text>
      </view>
    </scroll-view>

    <!-- 底部导航栏 -->
    <TabBar :active="2" />
  </view>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import TabBar from '@/components/TabBar/TabBar.vue'
import request from '@/api/request'
import { API_ENDPOINTS } from '@/api/config'
import { formatPoints } from '@/utils/format'

const keyword = ref('')
const loading = ref(false)
const listings = ref([])
const page = ref(1)
const pageSize = ref(20)
const currentCategory = ref('all')

const categories = ref([
  { label: '全部', value: 'all' },
  { label: 'PFP', value: 'pfp' },
  { label: '合集作品', value: 'collection' },
  { label: '衍生作品', value: 'derivative' },
  { label: '其他', value: 'other' }
])

onMounted(() => {
  fetchListings()
})

// 获取挂售列表
async function fetchListings() {
  loading.value = true
  try {
    const res = await request.get(API_ENDPOINTS.TRADE.GET_LISTINGS, {
      page: page.value,
      page_size: pageSize.value,
      category: currentCategory.value === 'all' ? '' : currentCategory.value,
      sort_by: 'price',
      sort_order: 'asc'
    })
    listings.value = res.list || []
  } catch (error) {
    uni.showToast({
      title: error.message || '加载失败',
      icon: 'none'
    })
  } finally {
    loading.value = false
  }
}

// 选择分类
function selectCategory(value) {
  currentCategory.value = value
  page.value = 1
  listings.value = []
  fetchListings()
}

// 加载更多
function loadMore() {
  if (!loading.value) {
    page.value++
    fetchListings()
  }
}

// 格式化价格
function formatPrice(price) {
  return formatPoints(price, 2)
}

// 跳转到详情页
function goToDetail(assetId) {
  uni.navigateTo({
    url: `/pages/asset-detail/index?id=${assetId}`
  })
}

// 显示筛选
function showFilter() {
  uni.showToast({
    title: '筛选功能开发中',
    icon: 'none'
  })
}
</script>

<style lang="scss" scoped>
.page {
  min-height: 100vh;
  background-color: #F5F5F5;
  padding-bottom: 140rpx;
}

.header {
  display: flex;
  align-items: center;
  gap: 16rpx;
  padding: 24rpx 32rpx;
  background-color: #FFFFFF;
  
  .search-box {
    flex: 1;
    display: flex;
    align-items: center;
    background-color: #F5F5F5;
    border-radius: 48rpx;
    padding: 0 32rpx;
    height: 72rpx;
    
    .search-icon {
      font-size: 32rpx;
      margin-right: 16rpx;
    }
    
    .search-input {
      flex: 1;
      font-size: 28rpx;
    }
  }
  
  .filter-btn {
    width: 72rpx;
    height: 72rpx;
    display: flex;
    align-items: center;
    justify-content: center;
    background-color: #F5F5F5;
    border-radius: 50%;
    
    .filter-icon {
      font-size: 36rpx;
    }
  }
}

.section-title {
  font-size: 36rpx;
  font-weight: 700;
  color: #000000;
  padding: 32rpx 32rpx 24rpx;
  background-color: #FFFFFF;
}

.category-tabs {
  display: flex;
  padding: 0 32rpx 24rpx;
  background-color: #FFFFFF;
  white-space: nowrap;
  
  .tab-item {
    display: inline-block;
    padding: 12rpx 32rpx;
    margin-right: 16rpx;
    background-color: #F5F5F5;
    border-radius: 48rpx;
    font-size: 26rpx;
    color: #666666;
    transition: all 0.3s ease;
    
    &.tab-active {
      background-color: #000000;
      color: #FFFFFF;
      font-weight: 600;
    }
  }
}

.asset-grid {
  height: calc(100vh - 400rpx);
  padding: 24rpx 32rpx;
}

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

.grid-wrapper {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 24rpx;
}

.asset-card {
  background-color: #FFFFFF;
  border-radius: 16rpx;
  overflow: hidden;
  
  .asset-image-wrapper {
    position: relative;
    width: 100%;
    height: 300rpx;
    
    .asset-image-placeholder {
      width: 100%;
      height: 100%;
      background: linear-gradient(135deg, #f5f7fa 0%, #c3cfe2 100%);
      display: flex;
      align-items: center;
      justify-content: center;
      
      .placeholder-emoji {
        font-size: 80rpx;
      }
    }
    
    .asset-tag {
      position: absolute;
      top: 12rpx;
      left: 12rpx;
      background-color: rgba(0, 0, 0, 0.7);
      border-radius: 8rpx;
      padding: 6rpx 12rpx;
      
      .tag-text {
        font-size: 20rpx;
        color: #FFFFFF;
      }
    }
    
    .asset-available {
      position: absolute;
      bottom: 12rpx;
      left: 12rpx;
      background-color: rgba(0, 0, 0, 0.7);
      border-radius: 8rpx;
      padding: 6rpx 12rpx;
      display: flex;
      align-items: center;
      gap: 6rpx;
      
      .available-icon {
        font-size: 20rpx;
      }
      
      .available-text {
        font-size: 20rpx;
        color: #FFFFFF;
      }
    }
  }
  
  .asset-info {
    padding: 20rpx;
    
    .asset-name {
      font-size: 26rpx;
      font-weight: 600;
      color: #000000;
      margin-bottom: 12rpx;
      overflow: hidden;
      text-overflow: ellipsis;
      white-space: nowrap;
    }
    
    .asset-price-row {
      display: flex;
      align-items: baseline;
      gap: 6rpx;
      margin-bottom: 8rpx;
      
      .price-value {
        font-size: 32rpx;
        font-weight: 700;
        color: #000000;
      }
      
      .price-unit {
        font-size: 22rpx;
        color: #999999;
      }
    }
    
    .asset-supply {
      display: flex;
      align-items: center;
      gap: 6rpx;
      
      .supply-icon {
        font-size: 20rpx;
      }
      
      .supply-text {
        font-size: 22rpx;
        color: #999999;
      }
    }
  }
}

.number-display {
  font-family: 'Monaco', 'Menlo', 'Consolas', monospace;
}
</style>
