<template>
  <view class="page-container">
    <!-- 自定义导航栏 -->
    <view class="navbar">
      <view class="navbar-title">集换专区</view>
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

    <!-- 排序和筛选 -->
    <view class="filter-bar">
      <view class="filter-item" @click="showSortMenu = !showSortMenu">
        <text>排序</text>
        <uni-icons type="down" size="16" color="#666"></uni-icons>
      </view>
      <view class="filter-item" @click="showPriceFilter = !showPriceFilter">
        <text>价格</text>
        <uni-icons type="down" size="16" color="#666"></uni-icons>
      </view>
      <view class="filter-item" @click="showMoreFilter = !showMoreFilter">
        <text>更多</text>
        <uni-icons type="down" size="16" color="#666"></uni-icons>
      </view>
    </view>

    <!-- 挂售列表 -->
    <view class="listings-grid">
      <view
        class="listing-card"
        v-for="listing in listings"
        :key="listing.id"
        @click="goToListingDetail(listing.id)"
      >
        <!-- 藏品图片 -->
        <image :src="listing.asset.image_url" class="listing-image" mode="aspectFill"></image>

        <!-- 藏品信息 -->
        <view class="listing-info">
          <view class="asset-name">{{ listing.asset.name }}</view>
          <view class="serial-number">#{{ listing.serial_number }}</view>
          <view class="price-section">
            <view class="price">
              <text class="price-value">{{ listing.price }}</text>
              <text class="price-unit">积分</text>
            </view>
          </view>
          <view class="seller-info">
            <image :src="listing.seller.avatar_url" class="seller-avatar"></image>
            <text class="seller-name">{{ listing.seller.nickname }}</text>
          </view>
        </view>

        <!-- 购买按钮 -->
        <view class="buy-button" @click.stop="buyNow(listing.id)">
          <text>立即兑换</text>
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

const activeCategory = ref('all')
const listings = ref([])
const loading = ref(false)
const hasMore = ref(true)
const page = ref(1)
const pageSize = ref(20)

const showSortMenu = ref(false)
const showPriceFilter = ref(false)
const showMoreFilter = ref(false)

// 模拟数据
const mockListings = [
  {
    id: 1,
    asset: {
      id: 1,
      name: '真我HOHO',
      image_url: 'https://via.placeholder.com/300x300?text=HOHO1'
    },
    serial_number: 1958,
    price: '18500.00000000',
    seller: {
      id: 1,
      nickname: '卖家昵称',
      avatar_url: 'https://via.placeholder.com/40x40?text=User'
    }
  },
  {
    id: 2,
    asset: {
      id: 1,
      name: '真我HOHO',
      image_url: 'https://via.placeholder.com/300x300?text=HOHO2'
    },
    serial_number: 1959,
    price: '18500.00000000',
    seller: {
      id: 2,
      nickname: '卖家昵称2',
      avatar_url: 'https://via.placeholder.com/40x40?text=User2'
    }
  }
]

// 获取挂售列表
const fetchListings = async (isLoadMore = false) => {
  if (loading.value) return

  loading.value = true
  try {
    // TODO: 调用API获取真实数据
    // const response = await request.get(API_ENDPOINTS.TRADE.GET_LISTINGS, {
    //   page: page.value,
    //   page_size: pageSize.value
    // })

    // 使用模拟数据
    if (isLoadMore) {
      listings.value = [...listings.value, ...mockListings]
    } else {
      listings.value = mockListings
    }

    hasMore.value = mockListings.length === pageSize.value
    page.value++
  } catch (error) {
    console.error('Failed to fetch listings:', error)
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
    fetchListings(true)
  }
}

// 跳转到挂售详情
const goToListingDetail = (listingId) => {
  uni.navigateTo({
    url: `/pages/asset-detail/index?listingId=${listingId}`
  })
}

// 立即兑换
const buyNow = (listingId) => {
  uni.showModal({
    title: '确认兑换',
    content: '确认用积分兑换这个藏品吗？',
    success: (res) => {
      if (res.confirm) {
        // TODO: 调用API执行交易
        uni.showToast({
          title: '兑换成功',
          icon: 'success'
        })
      }
    }
  })
}

// 页面加载
onMounted(() => {
  fetchListings()
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

.navbar-title {
  font-size: 18px;
  font-weight: 600;
  color: #000;
}

.navbar-right {
  display: flex;
  align-items: center;
  gap: 12px;
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

// 筛选条
.filter-bar {
  display: flex;
  padding: 12px 16px;
  background-color: #ffffff;
  gap: 12px;
  border-bottom: 1px solid #f0f0f0;
}

.filter-item {
  display: flex;
  align-items: center;
  gap: 4px;
  padding: 6px 12px;
  border: 1px solid #e8e8e8;
  border-radius: 20px;
  font-size: 14px;
  color: #666;
  cursor: pointer;
}

// 挂售列表
.listings-grid {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 12px;
  padding: 12px;
  background-color: #f5f5f5;
}

.listing-card {
  background-color: #ffffff;
  border-radius: 12px;
  overflow: hidden;
  cursor: pointer;
  transition: transform 0.3s ease;
  position: relative;

  &:active {
    transform: scale(0.98);
  }
}

.listing-image {
  width: 100%;
  height: 180px;
  display: block;
  background-color: #f0f0f0;
}

.listing-info {
  padding: 12px;
}

.asset-name {
  font-size: 14px;
  font-weight: 600;
  color: #000;
  margin-bottom: 4px;
}

.serial-number {
  font-size: 12px;
  color: #999;
  margin-bottom: 8px;
}

.price-section {
  margin-bottom: 8px;
}

.price {
  display: flex;
  align-items: baseline;
  gap: 4px;
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

.seller-info {
  display: flex;
  align-items: center;
  gap: 8px;
  margin-bottom: 8px;
}

.seller-avatar {
  width: 24px;
  height: 24px;
  border-radius: 50%;
  background-color: #f0f0f0;
}

.seller-name {
  font-size: 12px;
  color: #666;
  flex: 1;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.buy-button {
  width: 100%;
  padding: 8px;
  background-color: #000;
  color: #fff;
  text-align: center;
  font-size: 14px;
  font-weight: 600;
  border-radius: 0 0 12px 12px;
  cursor: pointer;
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
