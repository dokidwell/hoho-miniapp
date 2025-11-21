<template>
  <view class="login-page">
    <!-- 背景插画 -->
    <view class="background">
      <!-- 这里应该使用用户提供的插画背景图 -->
      <image src="/static/images/login-bg.png" mode="aspectFill"></image>
    </view>
    
    <!-- 主标语 -->
    <view class="slogan">
      <text class="line1">和你浪费的时间，</text>
      <text class="line2">比其它时间都好。</text>
    </view>
    
    <!-- 登录方式切换 -->
    <view class="login-container">
      <!-- 微信快捷登录模式 -->
      <view v-if="loginMode === 'wechat'" class="wechat-login">
        <button class="primary-btn" open-type="getPhoneNumber" @getphonenumber="getPhoneNumber">
          手机号快捷登录
        </button>
        <button class="secondary-btn" @click="switchToPhone">
          使用其他手机号登录
        </button>
        
        <view class="agreement">
          <checkbox-group @change="onAgreementChange">
            <label class="agreement-label">
              <checkbox :checked="agreed" color="#1890ff" />
              <text>登录注册即代表您已阅读并同意</text>
            </label>
          </checkbox-group>
          <view class="links">
            <text class="link" @click="openAgreement('user')">《用户协议》</text>
            <text>和</text>
            <text class="link" @click="openAgreement('privacy')">《隐私条款》</text>
          </view>
        </view>
      </view>
      
      <!-- 手机号验证码登录模式 -->
      <view v-else class="phone-login">
        <view class="input-group">
          <text class="label">手机号</text>
          <input 
            type="number" 
            maxlength="11" 
            placeholder="输入手机号" 
            v-model="phone"
            class="input"
          />
        </view>
        
        <view class="input-group">
          <text class="label">验证码</text>
          <input 
            type="number" 
            maxlength="6" 
            placeholder="6位验证码" 
            v-model="code"
            class="input"
          />
          <button 
            class="code-btn" 
            :disabled="codeSending || countdown > 0"
            @click="sendCode"
          >
            {{ countdown > 0 ? `${countdown}s` : '获取验证码' }}
          </button>
        </view>
        
        <view class="agreement">
          <checkbox-group @change="onAgreementChange">
            <label class="agreement-label">
              <checkbox :checked="agreed" color="#1890ff" />
              <text>已经阅读并同意</text>
              <text class="link" @click.stop="openAgreement('user')">《用户协议》</text>
              <text>和</text>
              <text class="link" @click.stop="openAgreement('privacy')">《隐私条款》</text>
            </label>
          </checkbox-group>
        </view>
        
        <button class="primary-btn" @click="phoneLogin">
          手机号快捷登录
        </button>
      </view>
    </view>
  </view>
</template>

<script>
import request from '@/api/request'
import { API_ENDPOINTS } from '@/api/config'

export default {
  data() {
    return {
      loginMode: 'wechat', // wechat | phone
      agreed: false,
      phone: '',
      code: '',
      codeSending: false,
      countdown: 0,
      countdownTimer: null
    }
  },
  
  onLoad() {
    // 检查是否已登录
    const token = uni.getStorageSync('token')
    if (token) {
      this.redirectToHome()
    }
  },
  
  onUnload() {
    if (this.countdownTimer) {
      clearInterval(this.countdownTimer)
    }
  },
  
  methods: {
    switchToPhone() {
      this.loginMode = 'phone'
    },
    
    onAgreementChange(e) {
      this.agreed = e.detail.value.length > 0
    },
    
    openAgreement(type) {
      const urls = {
        user: '/pages/agreement/user',
        privacy: '/pages/agreement/privacy'
      }
      uni.navigateTo({
        url: urls[type]
      })
    },
    
    async getPhoneNumber(e) {
      if (!this.agreed) {
        uni.showToast({
          title: '请先阅读并同意用户协议和隐私条款',
          icon: 'none'
        })
        return
      }
      
      if (e.detail.errMsg !== 'getPhoneNumber:ok') {
        uni.showToast({
          title: '获取手机号失败',
          icon: 'none'
        })
        return
      }
      
      try {
        uni.showLoading({ title: '登录中...' })
        
        // 获取微信登录code
        const { code } = await uni.login()
        
        // 调用后端登录接口
        const res = await request.post(API_ENDPOINTS.USER.LOGIN, {
          code,
          encrypted_data: e.detail.encryptedData,
          iv: e.detail.iv
        })
        
        // 保存token和用户信息
        uni.setStorageSync('token', res.token)
        uni.setStorageSync('userInfo', res.user)
        
        uni.hideLoading()
        uni.showToast({
          title: '登录成功',
          icon: 'success'
        })
        
        setTimeout(() => {
          this.redirectToHome()
        }, 1000)
      } catch (error) {
        uni.hideLoading()
        uni.showToast({
          title: error.message || '登录失败',
          icon: 'none'
        })
      }
    },
    
    async sendCode() {
      if (!this.phone) {
        uni.showToast({
          title: '请输入手机号',
          icon: 'none'
        })
        return
      }
      
      if (!/^1[3-9]\d{9}$/.test(this.phone)) {
        uni.showToast({
          title: '请输入正确的手机号',
          icon: 'none'
        })
        return
      }
      
      this.codeSending = true
      try {
        // 调用发送验证码接口
        await request.post(API_ENDPOINTS.USER.SEND_CODE, {
          phone: this.phone
        })
        
        uni.showToast({
          title: '验证码已发送',
          icon: 'success'
        })
        
        // 开始倒计时
        this.countdown = 60
        this.countdownTimer = setInterval(() => {
          this.countdown--
          if (this.countdown <= 0) {
            clearInterval(this.countdownTimer)
          }
        }, 1000)
      } catch (error) {
        uni.showToast({
          title: error.message || '发送失败',
          icon: 'none'
        })
      } finally {
        this.codeSending = false
      }
    },
    
    async phoneLogin() {
      if (!this.agreed) {
        uni.showToast({
          title: '请先阅读并同意用户协议和隐私条款',
          icon: 'none'
        })
        return
      }
      
      if (!this.phone) {
        uni.showToast({
          title: '请输入手机号',
          icon: 'none'
        })
        return
      }
      
      if (!this.code) {
        uni.showToast({
          title: '请输入验证码',
          icon: 'none'
        })
        return
      }
      
      uni.showLoading({ title: '登录中...' })
      try {
        const res = await request.post(API_ENDPOINTS.USER.LOGIN_BY_CODE, {
          phone: this.phone,
          code: this.code
        })
        
        // 保存token和用户信息
        uni.setStorageSync('token', res.token)
        uni.setStorageSync('userInfo', res.user)
        
        uni.hideLoading()
        uni.showToast({
          title: '登录成功',
          icon: 'success'
        })
        
        setTimeout(() => {
          this.redirectToHome()
        }, 1000)
      } catch (error) {
        uni.hideLoading()
        uni.showToast({
          title: error.message || '登录失败',
          icon: 'none'
        })
      }
    },
    
    redirectToHome() {
      uni.reLaunch({
        url: '/pages/index/index'
      })
    }
  }
}
</script>

<style scoped>
.login-page {
  position: relative;
  width: 100vw;
  height: 100vh;
  overflow: hidden;
}

.background {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  z-index: 0;
  background: linear-gradient(180deg, #2c3e50 0%, #34495e 100%);
}

.background image {
  width: 100%;
  height: 100%;
  opacity: 0.9;
}

.slogan {
  position: absolute;
  top: 280rpx;
  left: 80rpx;
  z-index: 1;
}

.line1,
.line2 {
  display: block;
  font-size: 48rpx;
  font-weight: 600;
  color: #fff;
  line-height: 1.5;
  text-shadow: 0 2px 8px rgba(0, 0, 0, 0.3);
}

.login-container {
  position: absolute;
  bottom: 120rpx;
  left: 0;
  right: 0;
  padding: 0 60rpx;
  z-index: 2;
}

.wechat-login,
.phone-login {
  width: 100%;
}

.input-group {
  position: relative;
  background: rgba(255, 255, 255, 0.15);
  backdrop-filter: blur(10px);
  border-radius: 16rpx;
  padding: 32rpx;
  margin-bottom: 24rpx;
  display: flex;
  align-items: center;
  border: 1px solid rgba(255, 255, 255, 0.2);
}

.label {
  font-size: 28rpx;
  color: #fff;
  width: 120rpx;
  flex-shrink: 0;
}

.input {
  flex: 1;
  font-size: 28rpx;
  color: #fff;
  background: transparent;
}

.input::placeholder {
  color: rgba(255, 255, 255, 0.5);
}

.code-btn {
  background: transparent;
  border: none;
  font-size: 24rpx;
  color: #1890ff;
  padding: 0;
  margin: 0;
  line-height: 1;
  flex-shrink: 0;
}

.code-btn::after {
  border: none;
}

.code-btn[disabled] {
  color: rgba(255, 255, 255, 0.5);
}

.primary-btn {
  width: 100%;
  height: 96rpx;
  line-height: 96rpx;
  background: rgba(24, 144, 255, 0.9);
  backdrop-filter: blur(10px);
  border-radius: 48rpx;
  font-size: 32rpx;
  color: #fff;
  font-weight: 600;
  border: none;
  margin-bottom: 24rpx;
}

.primary-btn::after {
  border: none;
}

.secondary-btn {
  width: 100%;
  height: 96rpx;
  line-height: 96rpx;
  background: rgba(255, 255, 255, 0.15);
  backdrop-filter: blur(10px);
  border-radius: 48rpx;
  font-size: 28rpx;
  color: #fff;
  border: 1px solid rgba(255, 255, 255, 0.3);
  margin-bottom: 48rpx;
}

.secondary-btn::after {
  border: none;
}

.agreement {
  text-align: center;
  font-size: 24rpx;
  color: rgba(255, 255, 255, 0.8);
  margin-bottom: 32rpx;
}

.agreement checkbox-group {
  display: inline-block;
}

.agreement-label {
  display: flex;
  align-items: center;
  justify-content: center;
  flex-wrap: wrap;
}

.agreement checkbox {
  margin-right: 8rpx;
}

.agreement text {
  line-height: 1.8;
}

.agreement .links {
  margin-top: 8rpx;
  display: flex;
  align-items: center;
  justify-content: center;
}

.link {
  color: #1890ff;
  text-decoration: underline;
  margin: 0 4rpx;
}
</style>
