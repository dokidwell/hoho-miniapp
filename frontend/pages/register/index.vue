<template>
  <view class="page">
    <!-- 顶部导航 -->
    <view class="navbar">
      <view class="back-btn" @click="goBack">
        <text class="back-icon">←</text>
      </view>
      <view class="navbar-title">注册账号</view>
      <view class="placeholder"></view>
    </view>

    <!-- Logo区域 -->
    <view class="logo-section">
      <view class="logo">🎨</view>
      <text class="welcome-text">欢迎加入HOHO Park</text>
    </view>

    <!-- 注册表单 -->
    <view class="form-section">
      <view class="form-item">
        <view class="input-wrapper">
          <text class="input-icon">📱</text>
          <input 
            class="form-input" 
            type="number" 
            v-model="phone" 
            placeholder="请输入手机号"
            maxlength="11"
          />
        </view>
      </view>
      
      <view class="form-item">
        <view class="input-wrapper">
          <text class="input-icon">🔒</text>
          <input 
            class="form-input" 
            type="text" 
            v-model="password" 
            :password="!showPassword" 
            placeholder="请设置密码（6-20位）"
          />
          <text class="toggle-password" @click="showPassword = !showPassword">
            {{ showPassword ? '👁️' : '🙈' }}
          </text>
        </view>
      </view>
      
      <view class="form-item">
        <view class="input-wrapper">
          <text class="input-icon">🔐</text>
          <input 
            class="form-input" 
            type="text" 
            v-model="confirmPassword" 
            :password="!showConfirmPassword" 
            placeholder="请再次输入密码"
          />
          <text class="toggle-password" @click="showConfirmPassword = !showConfirmPassword">
            {{ showConfirmPassword ? '👁️' : '🙈' }}
          </text>
        </view>
      </view>
      
      <button class="btn btn-primary" @click="handleRegister" :loading="registering" :disabled="registering">
        {{ registering ? '注册中...' : '立即注册' }}
      </button>
      
      <view class="form-footer">
        <text class="link-text" @click="goToLogin">已有账号？立即登录</text>
      </view>
    </view>

    <!-- 底部提示 -->
    <view class="bottom-tip">
      <text class="tip-text">注册即表示同意</text>
      <text class="tip-link">《用户协议》</text>
      <text class="tip-text">和</text>
      <text class="tip-link">《隐私政策》</text>
    </view>
  </view>
</template>

<script setup>
import { ref } from 'vue'
import request from '@/api/request'
import { API_ENDPOINTS } from '@/api/config'

const phone = ref('')
const password = ref('')
const confirmPassword = ref('')
const showPassword = ref(false)
const showConfirmPassword = ref(false)
const registering = ref(false)

// 注册
async function handleRegister() {
  // 验证手机号
  if (!phone.value || phone.value.length !== 11) {
    uni.showToast({
      title: '请输入正确的手机号',
      icon: 'none'
    })
    return
  }
  
  // 验证密码
  if (!password.value || password.value.length < 6 || password.value.length > 20) {
    uni.showToast({
      title: '密码长度为6-20位',
      icon: 'none'
    })
    return
  }
  
  // 验证确认密码
  if (password.value !== confirmPassword.value) {
    uni.showToast({
      title: '两次密码不一致',
      icon: 'none'
    })
    return
  }
  
  registering.value = true
  try {
    await request.post(API_ENDPOINTS.USER.REGISTER, {
      phone: phone.value,
      password: password.value
    })
    
    uni.showToast({
      title: '注册成功',
      icon: 'success'
    })
    
    // 延迟跳转到登录页
    setTimeout(() => {
      goToLogin()
    }, 1500)
  } catch (error) {
    uni.showToast({
      title: error.message || '注册失败',
      icon: 'none',
      duration: 2000
    })
  } finally {
    registering.value = false
  }
}

// 跳转到登录页
function goToLogin() {
  uni.navigateBack()
}

// 返回
function goBack() {
  uni.navigateBack()
}
</script>

<style lang="scss" scoped>
.page {
  min-height: 100vh;
  background: linear-gradient(135deg, #f5f7fa 0%, #c3cfe2 100%);
  display: flex;
  flex-direction: column;
}

.navbar {
  display: flex;
  align-items: center;
  justify-content: space-between;
  height: 88rpx;
  padding: 0 32rpx;
  background-color: transparent;
  
  .back-btn {
    width: 64rpx;
    height: 64rpx;
    display: flex;
    align-items: center;
    justify-content: center;
    
    .back-icon {
      font-size: 48rpx;
      color: #000000;
      font-weight: 300;
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

.logo-section {
  display: flex;
  flex-direction: column;
  align-items: center;
  margin: 80rpx 0 100rpx;
  
  .logo {
    font-size: 100rpx;
    margin-bottom: 24rpx;
  }
  
  .welcome-text {
    font-size: 32rpx;
    font-weight: 600;
    color: #000000;
    letter-spacing: 2rpx;
  }
}

.form-section {
  padding: 0 64rpx;
  
  .form-item {
    margin-bottom: 32rpx;
    
    .input-wrapper {
      display: flex;
      align-items: center;
      background-color: #FFFFFF;
      border-radius: 16rpx;
      padding: 0 32rpx;
      height: 96rpx;
      box-shadow: 0 4rpx 12rpx rgba(0, 0, 0, 0.05);
      
      .input-icon {
        font-size: 40rpx;
        margin-right: 24rpx;
      }
      
      .form-input {
        flex: 1;
        font-size: 28rpx;
        color: #000000;
      }
      
      .toggle-password {
        font-size: 40rpx;
        padding: 0 8rpx;
      }
    }
  }
  
  .btn {
    width: 100%;
    height: 96rpx;
    border-radius: 16rpx;
    font-size: 32rpx;
    font-weight: 600;
    border: none;
    margin-top: 16rpx;
    
    &.btn-primary {
      background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
      color: #FFFFFF;
      
      &:disabled {
        opacity: 0.6;
      }
    }
  }
  
  .form-footer {
    display: flex;
    justify-content: center;
    margin-top: 48rpx;
    
    .link-text {
      font-size: 28rpx;
      color: #667eea;
      font-weight: 500;
    }
  }
}

.bottom-tip {
  position: fixed;
  bottom: 64rpx;
  left: 0;
  right: 0;
  display: flex;
  justify-content: center;
  align-items: center;
  flex-wrap: wrap;
  padding: 0 64rpx;
  
  .tip-text {
    font-size: 22rpx;
    color: #999999;
  }
  
  .tip-link {
    font-size: 22rpx;
    color: #667eea;
    margin: 0 4rpx;
  }
}
</style>
