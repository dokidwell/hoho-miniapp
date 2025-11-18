<template>
  <view class="page">
    <view class="navbar">
      <text class="navbar-title">收货地址</text>
      <view class="navbar-right" @click="addAddress">
        <text class="add-btn">+ 新增</text>
      </view>
    </view>

    <view class="address-list">
      <view 
        v-for="address in addresses" 
        :key="address.id"
        class="address-item"
      >
        <view class="address-info">
          <view class="address-header">
            <text class="receiver-name">{{ address.receiver_name }}</text>
            <text class="receiver-phone">{{ address.phone }}</text>
          </view>
          <text class="address-detail">{{ address.province }} {{ address.city }} {{ address.district }} {{ address.detail }}</text>
          <view v-if="address.is_default" class="default-tag">默认</view>
        </view>
        <view class="address-actions">
          <text class="action-btn" @click="editAddress(address.id)">编辑</text>
          <text class="action-btn delete" @click="deleteAddress(address.id)">删除</text>
        </view>
      </view>
    </view>

    <view v-if="addresses.length === 0" class="empty-wrapper">
      <text class="empty-emoji">📍</text>
      <text class="empty-text">暂无收货地址</text>
    </view>
  </view>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import request from '@/api/request'
import { API_ENDPOINTS } from '@/api/config'

const addresses = ref([])

onMounted(() => {
  fetchAddresses()
})

async function fetchAddresses() {
  try {
    const res = await request.get(API_ENDPOINTS.USER.ADDRESSES)
    addresses.value = res.list || []
  } catch (error) {
    console.error('获取地址失败:', error)
  }
}

function addAddress() {
  uni.navigateTo({
    url: '/pages/address-edit/index'
  })
}

function editAddress(id) {
  uni.navigateTo({
    url: `/pages/address-edit/index?id=${id}`
  })
}

function deleteAddress(id) {
  uni.showModal({
    title: '提示',
    content: '确定要删除这个地址吗？',
    success: async (res) => {
      if (res.confirm) {
        try {
          await request.delete(`${API_ENDPOINTS.USER.ADDRESSES}/${id}`)
          uni.showToast({ title: '删除成功' })
          fetchAddresses()
        } catch (error) {
          uni.showToast({ title: '删除失败', icon: 'none' })
        }
      }
    }
  })
}
</script>

<style lang="scss" scoped>
.page {
  min-height: 100vh;
  background-color: #F5F5F5;
}

.navbar {
  height: 88rpx;
  background-color: #FFFFFF;
  display: flex;
  align-items: center;
  justify-content: center;
  border-bottom: 1rpx solid #F0F0F0;
  position: relative;
  
  .navbar-title {
    font-size: 32rpx;
    font-weight: 600;
    color: #000000;
  }
  
  .navbar-right {
    position: absolute;
    right: 32rpx;
    
    .add-btn {
      font-size: 28rpx;
      color: #667eea;
    }
  }
}

.address-list {
  padding: 24rpx 32rpx;
}

.address-item {
  background-color: #FFFFFF;
  border-radius: 16rpx;
  padding: 24rpx;
  margin-bottom: 16rpx;
  
  .address-info {
    margin-bottom: 16rpx;
    
    .address-header {
      display: flex;
      gap: 24rpx;
      margin-bottom: 12rpx;
      
      .receiver-name {
        font-size: 28rpx;
        font-weight: 600;
        color: #000000;
      }
      
      .receiver-phone {
        font-size: 28rpx;
        color: #666666;
      }
    }
    
    .address-detail {
      font-size: 26rpx;
      color: #666666;
      line-height: 1.6;
      display: block;
      margin-bottom: 12rpx;
    }
    
    .default-tag {
      display: inline-block;
      padding: 4rpx 12rpx;
      background-color: #667eea;
      color: #FFFFFF;
      font-size: 22rpx;
      border-radius: 6rpx;
    }
  }
  
  .address-actions {
    display: flex;
    gap: 32rpx;
    padding-top: 16rpx;
    border-top: 1rpx solid #F0F0F0;
    
    .action-btn {
      font-size: 26rpx;
      color: #667eea;
      
      &.delete {
        color: #F44336;
      }
    }
  }
}

.empty-wrapper {
  padding: 200rpx 0;
  text-align: center;
  
  .empty-emoji {
    font-size: 120rpx;
    display: block;
    margin-bottom: 24rpx;
  }
  
  .empty-text {
    font-size: 28rpx;
    color: #999999;
    display: block;
  }
}
</style>
