<template>
  <view class="page">
    <!-- 顶部导航 -->
    <view class="navbar">
      <text class="navbar-title">{{ isEdit ? '编辑地址' : '新增地址' }}</text>
    </view>

    <!-- 表单内容 -->
    <view class="form-container">
      <!-- 收货人 -->
      <view class="form-item">
        <text class="label">收货人</text>
        <input 
          class="input" 
          v-model="form.name" 
          placeholder="请输入收货人姓名"
          maxlength="20"
        />
      </view>

      <!-- 手机号 -->
      <view class="form-item">
        <text class="label">手机号</text>
        <input 
          class="input" 
          v-model="form.phone" 
          placeholder="请输入手机号"
          type="number"
          maxlength="11"
        />
      </view>

      <!-- 所在地区 -->
      <view class="form-item" @click="selectRegion">
        <text class="label">所在地区</text>
        <view class="region-selector">
          <text class="region-text" :class="{ 'placeholder': !regionText }">
            {{ regionText || '请选择省/市/区' }}
          </text>
          <text class="arrow">›</text>
        </view>
      </view>

      <!-- 详细地址 -->
      <view class="form-item address-item">
        <text class="label">详细地址</text>
        <textarea 
          class="textarea" 
          v-model="form.address" 
          placeholder="街道、楼牌号等"
          maxlength="100"
          :auto-height="true"
        />
      </view>

      <!-- 设为默认地址 -->
      <view class="form-item switch-item">
        <text class="label">设为默认地址</text>
        <switch :checked="form.isDefault" @change="onDefaultChange" color="#007AFF" />
      </view>
    </view>

    <!-- 底部按钮 -->
    <view class="footer">
      <button class="delete-btn" v-if="isEdit" @click="deleteAddress">删除地址</button>
      <button class="save-btn" @click="saveAddress">保存</button>
    </view>
  </view>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { onLoad } from '@dcloudio/uni-app'
import request from '@/api/request'

// 页面参数
const addressId = ref(null)
const isEdit = computed(() => !!addressId.value)

// 表单数据
const form = ref({
  name: '',
  phone: '',
  province: '',
  city: '',
  district: '',
  address: '',
  isDefault: false
})

// 地区文本
const regionText = computed(() => {
  if (form.value.province && form.value.city && form.value.district) {
    return `${form.value.province} ${form.value.city} ${form.value.district}`
  }
  return ''
})

// 页面加载
onLoad((options) => {
  if (options.id) {
    addressId.value = options.id
    loadAddress()
  }
})

// 加载地址详情
const loadAddress = async () => {
  try {
    const res = await request.get(`/addresses/${addressId.value}`)
    form.value = res.data
  } catch (error) {
    uni.showToast({
      title: error.message || '加载失败',
      icon: 'none'
    })
  }
}

// 选择地区
const selectRegion = () => {
  uni.navigateTo({
    url: `/pages/region-picker/index?province=${form.value.province}&city=${form.value.city}&district=${form.value.district}`
  })
}

// 监听地区选择返回
onMounted(() => {
  uni.$on('regionSelected', (region) => {
    form.value.province = region.province
    form.value.city = region.city
    form.value.district = region.district
  })
})

// 切换默认地址
const onDefaultChange = (e) => {
  form.value.isDefault = e.detail.value
}

// 验证表单
const validateForm = () => {
  if (!form.value.name) {
    uni.showToast({ title: '请输入收货人姓名', icon: 'none' })
    return false
  }

  if (!form.value.phone) {
    uni.showToast({ title: '请输入手机号', icon: 'none' })
    return false
  }

  if (!/^1[3-9]\d{9}$/.test(form.value.phone)) {
    uni.showToast({ title: '请输入正确的手机号', icon: 'none' })
    return false
  }

  if (!form.value.province || !form.value.city || !form.value.district) {
    uni.showToast({ title: '请选择所在地区', icon: 'none' })
    return false
  }

  if (!form.value.address) {
    uni.showToast({ title: '请输入详细地址', icon: 'none' })
    return false
  }

  return true
}

// 保存地址
const saveAddress = async () => {
  if (!validateForm()) return

  try {
    uni.showLoading({ title: '保存中...' })
    
    if (isEdit.value) {
      await request.put(`/addresses/${addressId.value}`, form.value)
    } else {
      await request.post('/addresses', form.value)
    }

    uni.hideLoading()
    uni.showToast({ title: '保存成功', icon: 'success' })

    setTimeout(() => uni.navigateBack(), 1500)
  } catch (error) {
    uni.hideLoading()
    uni.showToast({ title: error.message || '保存失败', icon: 'none' })
  }
}

// 删除地址
const deleteAddress = () => {
  uni.showModal({
    title: '提示',
    content: '确定要删除这个地址吗？',
    success: async (res) => {
      if (res.confirm) {
        try {
          uni.showLoading({ title: '删除中...' })
          await request.delete(`/addresses/${addressId.value}`)
          uni.hideLoading()
          uni.showToast({ title: '删除成功', icon: 'success' })
          setTimeout(() => uni.navigateBack(), 1500)
        } catch (error) {
          uni.hideLoading()
          uni.showToast({ title: error.message || '删除失败', icon: 'none' })
        }
      }
    }
  })
}
</script>

<style scoped>
.page {
  min-height: 100vh;
  background-color: #f5f5f5;
  padding-bottom: 200rpx;
}

.navbar {
  background-color: #fff;
  padding: 30rpx;
  text-align: center;
  border-bottom: 1rpx solid #eee;
}

.navbar-title {
  font-size: 36rpx;
  font-weight: 600;
}

.form-container {
  margin-top: 20rpx;
  background-color: #fff;
}

.form-item {
  display: flex;
  align-items: center;
  padding: 30rpx;
  border-bottom: 1rpx solid #f0f0f0;
}

.form-item.address-item {
  align-items: flex-start;
}

.form-item.switch-item {
  justify-content: space-between;
}

.label {
  font-size: 30rpx;
  color: #333;
  width: 180rpx;
  flex-shrink: 0;
}

.input {
  flex: 1;
  font-size: 30rpx;
  color: #333;
}

.textarea {
  flex: 1;
  font-size: 30rpx;
  color: #333;
  min-height: 120rpx;
}

.region-selector {
  flex: 1;
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.region-text {
  font-size: 30rpx;
  color: #333;
}

.region-text.placeholder {
  color: #999;
}

.arrow {
  font-size: 40rpx;
  color: #999;
  font-weight: 300;
}

.footer {
  position: fixed;
  bottom: 0;
  left: 0;
  right: 0;
  background-color: #fff;
  padding: 20rpx 30rpx;
  padding-bottom: calc(20rpx + env(safe-area-inset-bottom));
  box-shadow: 0 -4rpx 12rpx rgba(0, 0, 0, 0.05);
  display: flex;
  gap: 20rpx;
}

.delete-btn {
  flex: 1;
  height: 88rpx;
  line-height: 88rpx;
  background-color: #fff;
  border: 1rpx solid #ff4444;
  color: #ff4444;
  font-size: 32rpx;
  border-radius: 8rpx;
}

.save-btn {
  flex: 2;
  height: 88rpx;
  line-height: 88rpx;
  background-color: #007AFF;
  color: #fff;
  font-size: 32rpx;
  border-radius: 8rpx;
}
</style>
