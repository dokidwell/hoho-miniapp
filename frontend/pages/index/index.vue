<template>
  <view class="page-container">
    <!-- 自定义导航栏 -->
    <view class="navbar">
      <view class="navbar-left">
        <view class="search-box">
          <uni-icons type="search" size="18" color="#999"></uni-icons>
          <input type="text" placeholder="搜索..." class="search-input" />
        </view>
      </view>
      <view class="navbar-right">
        <uni-icons type="more-filled" size="24" color="#333"></uni-icons>
        <view class="filter-toggle">
          <uni-icons type="circle" size="20" color="#3a8fff"></uni-icons>
        </view>
      </view>
    </view>

    <!-- 分类标签 -->
    <view class="category-tabs">
      <view
        class="tab-item"
        :class="{ active: activeCategory === 'all' }"
        @click="activeCategory = 'all'"
      >
        全部
      </view>
      <view
        class="tab-item"
        :class="{ active: activeCategory === 'pfp' }"
        @click="activeCategory = 'pfp'"
      >
        PFP
      </view>
      <view
        class="tab-item"
        :class="{ active: activeCategory === 'collection' }"
        @click="activeCategory = 'collection'"
      >
        合集作品
      </view>
      <view
        class="tab-item"
        :class="{ active: activeCategory === 'derivative' }"
        @click="activeCategory = 'derivative'"
      >
        衍生作品
      </view>
    </view>

    <!-- 藏品网格 -->
    <view class="assets-grid">
      <view
        class="asset-card"
        v-for="asset in assets"
        :key="asset.id"
        @click="goToAssetDetail(asset.id)"
      >
        <!-- 藏品标签 -->
        <view class="asset-label">
          <text class="label-text">官方合作作品</text>
        </view>

        <!-- 藏品图片 -->
        <image :src="asset.image_url" class="asset-image" mode="aspectFill"></image>

        <!-- 藏品信息 -->
        <view class="asset-info">
          <view class="asset-name">{{ asset.name }}</view>
          <view class="asset-meta">
            <view class="meta-item">
              <text class="meta-label">{{ asset.total_supply }}份可兑</text>
            </view>
          </view>
          <view class="asset-price">
            <text class="price-value">18500</text>
            <text class="price-unit">起</text>
          </view>
          <view class="asset-quantity">
            <uni-icons type="circle" size="14" color="#333"></uni-icons>
            <text class="quantity-text">5000份</text>
          </view>
        </view>
      </view>
    </view>

    <!-- 加载更多 -->
    <view class="load-more" v-if="!loading && hasMore">
      <text @click="loadMore">加载更多</text>
    </view>

    <!-- 加载中 -->
    <view class="loading" v-if="loading">
      <uni-load-more status="loading"></uni-load-more>
    </view>
  </view>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useAssetStore } from '../../stores/asset'

const assetStore = useAssetStore()

const activeCategory = ref('all')
const assets = ref([])
const loading = ref(false)
const hasMore = ref(true)
const page = ref(1)
const pageSize = ref(20)

// 获取藏品列表
const fetchAssets = async (isLoadMore = false) => {
  if (loading.value) return

  loading.value = true
  try {
    const response = await assetStore.fetchAssets({
      page: page.value,
      page_size: pageSize.value,
      status: 'approved'
    })

    if (isLoadMore) {
      assets.value = [...assets.value, ...response.list]
    } else {
      assets.value = response.list
    }

    hasMore.value = response.list.length === pageSize.value
    page.value++
  } catch (error) {
    console.error('Failed to fetch assets:', error)
    uni.showToast({
      title: '加载失败',
      icon: 'none'
    })
  } finally {
    loading.value = false
  }
}

// 加载更多
const loadMore = () => {
  if (hasMore.value && !loading.value) {
    fetchAssets(true)
  }
}

// 跳转到藏品详情
const goToAssetDetail = (assetId) => {
  uni.navigateTo({
    url: `/pages/asset-detail/index?id=${assetId}`
  })
}

// 页面加载
onMounted(() => {
  fetchAssets()
})
</script>

<style lang="scss" scoped>
.page-container {
  width: 100%;
  min-height: 100vh;
  background-color: #f5f5f5;
  padding-bottom: 20px;
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

.navbar-left {
  flex: 1;
}

.search-box {
  display: flex;
  align-items: center;
  padding: 8px 12px;
  background-color: #f0f0f0;
  border-radius: 20px;
  gap: 8px;
}

.search-input {
  flex: 1;
  background: transparent;
  border: none;
  font-size: 14px;
  color: #333;

  &::placeholder {
    color: #999;
  }
}

.navbar-right {
  display: flex;
  align-items: center;
  gap: 12px;
  margin-left: 12px;
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

// 分类标签
.category-tabs {
  display: flex;
  padding: 12px 16px;
  background-color: #ffffff;
  gap: 8px;
  overflow-x: auto;
  border-bottom: 1px solid #f0f0f0;

  &::-webkit-scrollbar {
    display: none;
  }
}

.tab-item {
  padding: 8px 16px;
  border: 1px solid #e8e8e8;
  border-radius: 20px;
  font-size: 14px;
  color: #666;
  white-space: nowrap;
  transition: all 0.3s ease;

  &.active {
    background-color: #000;
    color: #fff;
    border-color: #000;
  }
}

// 藏品网格
.assets-grid {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 12px;
  padding: 12px;
  background-color: #f5f5f5;
}

.asset-card {
  background-color: #ffffff;
  border-radius: 12px;
  overflow: hidden;
  cursor: pointer;
  transition: transform 0.3s ease;

  &:active {
    transform: scale(0.98);
  }
}

.asset-label {
  position: absolute;
  top: 8px;
  left: 8px;
  z-index: 10;
  background-color: rgba(0, 0, 0, 0.7);
  padding: 4px 8px;
  border-radius: 4px;
}

.label-text {
  font-size: 12px;
  color: #fff;
}

.asset-image {
  width: 100%;
  height: 180px;
  display: block;
  background-color: #f0f0f0;
  position: relative;
}

.asset-info {
  padding: 12px;
}

.asset-name {
  font-size: 14px;
  font-weight: 600;
  color: #000;
  margin-bottom: 8px;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.asset-meta {
  display: flex;
  gap: 8px;
  margin-bottom: 8px;
}

.meta-item {
  display: flex;
  align-items: center;
  gap: 4px;
}

.meta-label {
  font-size: 12px;
  color: #666;
  background-color: #f5f5f5;
  padding: 2px 6px;
  border-radius: 4px;
}

.asset-price {
  display: flex;
  align-items: baseline;
  gap: 4px;
  margin-bottom: 8px;
}

.price-value {
  font-size: 16px;
  font-weight: 600;
  color: #000;
}

.price-unit {
  font-size: 12px;
  color: #666;
}

.asset-quantity {
  display: flex;
  align-items: center;
  gap: 4px;
  font-size: 12px;
  color: #666;
}

.quantity-text {
  color: #666;
}

// 加载更多
.load-more {
  text-align: center;
  padding: 20px;

  text {
    color: #3a8fff;
    font-size: 14px;
  }
}

.loading {
  padding: 20px;
  text-align: center;
}
</style>
