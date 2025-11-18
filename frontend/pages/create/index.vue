<template>
  <view class="page">
    <!-- 顶部导航 -->
    <view class="navbar">
      <view class="back-btn" @click="goBack" v-if="currentStep > 1">
        <text class="back-icon">←</text>
      </view>
      <view class="navbar-title">创作</view>
      <view class="placeholder"></view>
    </view>

    <!-- 步骤指示器 -->
    <view class="steps">
      <view class="step-item" :class="{ 'step-active': currentStep >= 1 }">
        <view class="step-number">1</view>
        <text class="step-label">上传作品</text>
      </view>
      <view class="step-line" :class="{ 'step-active': currentStep >= 2 }"></view>
      <view class="step-item" :class="{ 'step-active': currentStep >= 2 }">
        <view class="step-number">2</view>
        <text class="step-label">填写信息</text>
      </view>
      <view class="step-line" :class="{ 'step-active': currentStep >= 3 }"></view>
      <view class="step-item" :class="{ 'step-active': currentStep >= 3 }">
        <view class="step-number">3</view>
        <text class="step-label">确认提交</text>
      </view>
    </view>

    <!-- 步骤1：上传作品 -->
    <view v-if="currentStep === 1" class="step-content">
      <view class="upload-section">
        <view class="upload-box" @click="chooseImage">
          <view v-if="formData.image_url" class="preview-wrapper">
            <text class="preview-emoji">🖼️</text>
            <text class="preview-text">已选择图片</text>
          </view>
          <view v-else class="upload-placeholder">
            <text class="upload-emoji">📤</text>
            <text class="upload-text">点击上传作品</text>
            <text class="upload-hint">支持图片、视频、音频</text>
          </view>
        </view>
      </view>
      
      <view class="btn-group">
        <button class="btn btn-primary" @click="nextStep" :disabled="!formData.image_url">下一步</button>
      </view>
    </view>

    <!-- 步骤2：填写信息 -->
    <view v-if="currentStep === 2" class="step-content">
      <view class="form-section">
        <view class="form-item">
          <text class="form-label">作品名称</text>
          <input class="form-input" v-model="formData.name" placeholder="请输入作品名称" />
        </view>
        
        <view class="form-item">
          <text class="form-label">作品描述</text>
          <textarea class="form-textarea" v-model="formData.description" placeholder="请输入作品描述" />
        </view>
        
        <view class="form-item">
          <text class="form-label">制作数量</text>
          <input class="form-input" type="number" v-model="formData.total_supply" placeholder="请输入制作数量" />
        </view>
        
        <view class="form-item">
          <text class="form-label">分类</text>
          <picker mode="selector" :range="categories" range-key="label" @change="onCategoryChange">
            <view class="picker-value">
              {{ selectedCategory ? selectedCategory.label : '请选择分类' }}
            </view>
          </picker>
        </view>
      </view>
      
      <view class="btn-group">
        <button class="btn btn-secondary" @click="prevStep">上一步</button>
        <button class="btn btn-primary" @click="nextStep" :disabled="!isStep2Valid">下一步</button>
      </view>
    </view>

    <!-- 步骤3：确认提交 -->
    <view v-if="currentStep === 3" class="step-content">
      <view class="preview-section">
        <view class="preview-title">作品预览</view>
        <view class="preview-image-large">
          <text class="preview-emoji-large">🖼️</text>
        </view>
        
        <view class="preview-info">
          <view class="info-row">
            <text class="info-label">作品名称</text>
            <text class="info-value">{{ formData.name }}</text>
          </view>
          <view class="info-row">
            <text class="info-label">作品描述</text>
            <text class="info-value">{{ formData.description }}</text>
          </view>
          <view class="info-row">
            <text class="info-label">制作数量</text>
            <text class="info-value">{{ formData.total_supply }}份</text>
          </view>
          <view class="info-row">
            <text class="info-label">分类</text>
            <text class="info-value">{{ selectedCategory?.label }}</text>
          </view>
        </view>
      </view>
      
      <view class="btn-group">
        <button class="btn btn-secondary" @click="prevStep">上一步</button>
        <button class="btn btn-primary" @click="submitCreate" :loading="submitting" :disabled="submitting">
          {{ submitting ? '提交中...' : '提交审核' }}
        </button>
      </view>
    </view>

    <!-- 底部导航栏 -->
    <TabBar :active="1" />
  </view>
</template>

<script setup>
import { ref, computed } from 'vue'
import TabBar from '@/components/TabBar/TabBar.vue'
import request from '@/api/request'
import { API_ENDPOINTS } from '@/api/config'
import { chooseAndUploadImage } from '@/utils/upload'

const currentStep = ref(1)
const submitting = ref(false)

const formData = ref({
  name: '',
  description: '',
  image_url: '',
  total_supply: '',
  category: ''
})

const categories = ref([
  { label: '艺术', value: 'art' },
  { label: '收藏', value: 'collectible' },
  { label: '游戏', value: 'game' },
  { label: '音乐', value: 'music' },
  { label: '其他', value: 'other' }
])

const selectedCategory = ref(null)

const isStep2Valid = computed(() => {
  return formData.value.name && 
         formData.value.description && 
         formData.value.total_supply && 
         selectedCategory.value
})

// 选择图片
async function chooseImage() {
  try {
    uni.showLoading({ title: '上传中...' })
    
    const urls = await chooseAndUploadImage(
      { count: 1 },
      (current, total, progress) => {
        uni.showLoading({ title: `上传中 ${progress}%` })
      }
    )
    
    formData.value.image_url = urls[0]
    uni.hideLoading()
    uni.showToast({ title: '上传成功', icon: 'success' })
  } catch (error) {
    uni.hideLoading()
    uni.showToast({
      title: error.message || '上传失败',
      icon: 'none'
    })
  }
}

// 分类选择
function onCategoryChange(e) {
  const index = e.detail.value
  selectedCategory.value = categories.value[index]
  formData.value.category = selectedCategory.value.value
}

// 下一步
function nextStep() {
  if (currentStep.value < 3) {
    currentStep.value++
  }
}

// 上一步
function prevStep() {
  if (currentStep.value > 1) {
    currentStep.value--
  }
}

// 返回
function goBack() {
  if (currentStep.value > 1) {
    prevStep()
  } else {
    uni.navigateBack()
  }
}

// 提交铸造请求
async function submitCreate() {
  submitting.value = true
  try {
    // 先上传图片
    let imageUrl = formData.value.image_url
    if (imageUrl && imageUrl.startsWith('http://tmp/')) {
      imageUrl = await request.upload('/upload', formData.value.image_url, 'file')
    }
    
    // 提交铸造请求
    await request.post(API_ENDPOINTS.ASSET.CREATE, {
      name: formData.value.name,
      description: formData.value.description,
      image_url: imageUrl,
      total_supply: parseInt(formData.value.total_supply),
      category: formData.value.category
    })
    
    uni.showToast({
      title: '提交成功，等待审核',
      icon: 'success',
      duration: 2000
    })
    
    setTimeout(() => {
      uni.switchTab({
        url: '/pages/index/index'
      })
    }, 2000)
  } catch (error) {
    uni.showToast({
      title: error.message || '提交失败',
      icon: 'none',
      duration: 2000
    })
  } finally {
    submitting.value = false
  }
}
</script>

<style lang="scss" scoped>
.page {
  min-height: 100vh;
  background-color: #F5F5F5;
  padding-bottom: 140rpx;
}

.navbar {
  display: flex;
  align-items: center;
  justify-content: space-between;
  height: 88rpx;
  padding: 0 32rpx;
  background-color: #FFFFFF;
  border-bottom: 1px solid #E8E8E8;
  
  .back-btn {
    width: 64rpx;
    height: 64rpx;
    display: flex;
    align-items: center;
    justify-content: center;
    
    .back-icon {
      font-size: 48rpx;
      color: #000000;
    }
  }
  
  .navbar-title {
    font-size: 32rpx;
    font-weight: 600;
    color: #000000;
  }
  
  .placeholder {
    width: 64rpx;
  }
}

.steps {
  display: flex;
  align-items: center;
  padding: 48rpx 64rpx;
  background-color: #FFFFFF;
  
  .step-item {
    display: flex;
    flex-direction: column;
    align-items: center;
    gap: 12rpx;
    
    .step-number {
      width: 56rpx;
      height: 56rpx;
      border-radius: 50%;
      background-color: #E8E8E8;
      color: #999999;
      display: flex;
      align-items: center;
      justify-content: center;
      font-size: 24rpx;
      font-weight: 600;
      transition: all 0.3s ease;
    }
    
    .step-label {
      font-size: 22rpx;
      color: #999999;
      transition: all 0.3s ease;
    }
    
    &.step-active {
      .step-number {
        background-color: #000000;
        color: #FFFFFF;
      }
      
      .step-label {
        color: #000000;
        font-weight: 600;
      }
    }
  }
  
  .step-line {
    flex: 1;
    height: 2rpx;
    background-color: #E8E8E8;
    margin: 0 16rpx;
    transition: all 0.3s ease;
    
    &.step-active {
      background-color: #000000;
    }
  }
}

.step-content {
  padding: 32rpx;
}

.upload-section {
  .upload-box {
    background-color: #FFFFFF;
    border-radius: 16rpx;
    height: 500rpx;
    display: flex;
    align-items: center;
    justify-content: center;
    border: 2rpx dashed #CCCCCC;
    
    .upload-placeholder,
    .preview-wrapper {
      display: flex;
      flex-direction: column;
      align-items: center;
      gap: 16rpx;
      
      .upload-emoji,
      .preview-emoji {
        font-size: 120rpx;
      }
      
      .upload-text,
      .preview-text {
        font-size: 28rpx;
        color: #666666;
        font-weight: 500;
      }
      
      .upload-hint {
        font-size: 24rpx;
        color: #999999;
      }
    }
  }
}

.form-section {
  background-color: #FFFFFF;
  border-radius: 16rpx;
  padding: 32rpx;
  
  .form-item {
    margin-bottom: 32rpx;
    
    &:last-child {
      margin-bottom: 0;
    }
    
    .form-label {
      display: block;
      font-size: 28rpx;
      color: #000000;
      font-weight: 600;
      margin-bottom: 16rpx;
    }
    
    .form-input,
    .picker-value {
      width: 100%;
      height: 88rpx;
      background-color: #F5F5F5;
      border-radius: 12rpx;
      padding: 0 24rpx;
      font-size: 28rpx;
      color: #000000;
      display: flex;
      align-items: center;
    }
    
    .form-textarea {
      width: 100%;
      min-height: 200rpx;
      background-color: #F5F5F5;
      border-radius: 12rpx;
      padding: 24rpx;
      font-size: 28rpx;
      color: #000000;
    }
    
    .picker-value {
      color: #999999;
    }
  }
}

.preview-section {
  background-color: #FFFFFF;
  border-radius: 16rpx;
  padding: 32rpx;
  
  .preview-title {
    font-size: 32rpx;
    font-weight: 600;
    color: #000000;
    margin-bottom: 24rpx;
  }
  
  .preview-image-large {
    width: 100%;
    height: 400rpx;
    background: linear-gradient(135deg, #f5f7fa 0%, #c3cfe2 100%);
    border-radius: 16rpx;
    display: flex;
    align-items: center;
    justify-content: center;
    margin-bottom: 32rpx;
    
    .preview-emoji-large {
      font-size: 160rpx;
    }
  }
  
  .preview-info {
    .info-row {
      display: flex;
      justify-content: space-between;
      padding: 24rpx 0;
      border-bottom: 1rpx solid #F0F0F0;
      
      &:last-child {
        border-bottom: none;
      }
      
      .info-label {
        font-size: 28rpx;
        color: #666666;
      }
      
      .info-value {
        font-size: 28rpx;
        color: #000000;
        font-weight: 500;
        max-width: 400rpx;
        text-align: right;
      }
    }
  }
}

.btn-group {
  display: flex;
  gap: 24rpx;
  margin-top: 48rpx;
  
  .btn {
    flex: 1;
    height: 88rpx;
    border-radius: 12rpx;
    font-size: 30rpx;
    font-weight: 600;
    border: none;
    
    &.btn-primary {
      background-color: #000000;
      color: #FFFFFF;
      
      &:disabled {
        opacity: 0.4;
      }
    }
    
    &.btn-secondary {
      background-color: #F5F5F5;
      color: #000000;
    }
  }
}
</style>
