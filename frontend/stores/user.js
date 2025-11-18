import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import { request } from '../api/request'
import { API_ENDPOINTS } from '../api/config'

export const useUserStore = defineStore('user', () => {
  // 状态
  const user = ref(null)
  const token = ref(uni.getStorageSync('token') || '')
  const points = ref(null)
  const isLoggedIn = computed(() => !!token.value)

  // 登录
  const login = async (phone, password) => {
    try {
      const response = await request.post(API_ENDPOINTS.USER.LOGIN, {
        phone,
        password
      })
      token.value = response.token
      user.value = response.user
      uni.setStorageSync('token', response.token)
      return response
    } catch (error) {
      throw error
    }
  }

  // 注册
  const register = async (phone, password, confirmPassword) => {
    try {
      const response = await request.post(API_ENDPOINTS.USER.REGISTER, {
        phone,
        password,
        confirmPassword
      })
      token.value = response.token
      user.value = response.user
      uni.setStorageSync('token', response.token)
      return response
    } catch (error) {
      throw error
    }
  }

  // 获取个人资料
  const fetchProfile = async () => {
    try {
      const response = await request.get(API_ENDPOINTS.USER.PROFILE)
      user.value = response
      return response
    } catch (error) {
      throw error
    }
  }

  // 更新个人资料
  const updateProfile = async (data) => {
    try {
      const response = await request.put(API_ENDPOINTS.USER.UPDATE_PROFILE, data)
      user.value = response
      return response
    } catch (error) {
      throw error
    }
  }

  // 获取积分余额
  const fetchPoints = async () => {
    try {
      const response = await request.get(API_ENDPOINTS.USER.GET_POINTS)
      points.value = response
      return response
    } catch (error) {
      throw error
    }
  }

  // 实名认证
  const verifyIdentity = async (realName, idNumber) => {
    try {
      const response = await request.post(API_ENDPOINTS.USER.VERIFY_IDENTITY, {
        realName,
        idNumber
      })
      user.value = response
      return response
    } catch (error) {
      throw error
    }
  }

  // 登出
  const logout = async () => {
    try {
      await request.post(API_ENDPOINTS.USER.LOGOUT)
    } catch (error) {
      console.error('Logout error:', error)
    } finally {
      token.value = ''
      user.value = null
      points.value = null
      uni.removeStorageSync('token')
      uni.navigateTo({ url: '/pages/auth/login' })
    }
  }

  // 检查登录状态
  const checkLoginStatus = () => {
    if (!token.value) {
      uni.navigateTo({ url: '/pages/auth/login' })
      return false
    }
    return true
  }

  return {
    user,
    token,
    points,
    isLoggedIn,
    login,
    register,
    fetchProfile,
    updateProfile,
    fetchPoints,
    verifyIdentity,
    logout,
    checkLoginStatus
  }
})
