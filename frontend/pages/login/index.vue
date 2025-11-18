<template>
  <view class="page">
    <!-- Logo区域 -->
    <view class="logo-section">
      <view class="logo">🎨</view>
      <text class="app-name">HOHO Park</text>
      <text class="app-slogan">数字藏品集换平台</text>
    </view>

    <!-- 登录表单 -->
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
            placeholder="请输入密码"
          />
          <text class="toggle-password" @click="showPassword = !showPassword">
            {{ showPassword ? '👁️' : '🙈' }}
          </text>
        </view>
      </view>
      
      <button class="btn btn-primary" @click="handleLogin" :loading="logging" :disabled="logging">
        {{ logging ? '登录中...' : '登录' }}
      </button>
      
      <view class="form-footer">
        <text class="link-text" @click="goToRegister">还没有账号？立即注册</text>
      </view>
    </view>

    <!-- 底部提示 -->
    <view class="bottom-tip">
      <text class="tip-text">登录即表示同意</text>
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
const showPassword = ref(false)
const logging = ref(false)

// 登录
async function handleLogin() {
  // 验证手机号
  if (!phone.value || phone.value.length !== 11) {
    uni.showToast({
      title: '请输入正确的手机号',
      icon: 'none'
    })
    return
  }
  
  // 验证密码
  if (!password.value) {
    uni.showToast({
      title: '请输入密码',
      icon: 'none'
    })
    return
  }
  
  logging.value = true
  try {
    const res = await request.post(API_ENDPOINTS.USER.LOGIN, {
      phone: phone.value,
      password: password.value
    })
    
    // 保存token和用户信息
    uni.setStorageSync('token', res.token)
    uni.setStorageSync('userInfo', res.user)
    
    uni.showToast({
      title: '登录成功',
      icon: 'success'
    })
    
    // 延迟跳转到首页
    setTimeout(() => {
      uni.switchTab({
        url: '/pages/index/index'
      })
    }, 1500)
  } catch (error) {
    uni.showToast({
      title: error.message || '登录失败',
      icon: 'none',
      duration: 2000
    })
  } finally {
    logging.value = false
  }
}

// 跳转到注册页
function goToRegister() {
  uni.navigateTo({
    url: '/pages/register/index'
  })
}
</script>

<style lang="scss" scoped>
.page {
  min-height: 100vh;
  background: linear-gradient(135deg, #f5f7fa 0%, #c3cfe2 100%);
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 120rpx 64rpx 64rpx;
}

.logo-section {
  display: flex;
  flex-direction: column;
  align-items: center;
  margin-bottom: 120rpx;
  
  .logo {
    font-size: 120rpx;
    margin-bottom: 32rpx;
  }
  
  .app-name {
    font-size: 48rpx;
    font-weight: 700;
    color: #000000;
    margin-bottom: 16rpx;
    letter-spacing: 4rpx;
  }
  
  .app-slogan {
    font-size: 24rpx;
    color: #666666;
    letter-spacing: 2rpx;
  }
}

.form-section {
  width: 100%;
  
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
