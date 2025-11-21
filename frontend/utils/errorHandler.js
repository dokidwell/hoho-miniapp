/**
 * 统一错误处理工具
 * 提供友好的用户提示和错误日志
 */

class ErrorHandler {
  /**
   * 处理API错误
   * @param {Error} error - 错误对象
   * @param {Object} options - 配置选项
   */
  static handleApiError(error, options = {}) {
    const {
      showToast = true,
      duration = 2000,
      logError = true
    } = options

    let message = '操作失败，请稍后重试'
    let code = 0

    if (error.response) {
      // 服务器返回错误
      const { status, data } = error.response
      code = status

      switch (status) {
        case 400:
          message = data.detail || data.message || '请求参数错误'
          break
        case 401:
          message = '请先登录'
          // 跳转到登录页
          setTimeout(() => {
            uni.navigateTo({
              url: '/pages/login/index'
            })
          }, 1500)
          break
        case 403:
          message = '没有权限执行此操作'
          break
        case 404:
          message = '请求的资源不存在'
          break
        case 409:
          message = data.detail || '操作冲突，请刷新后重试'
          break
        case 422:
          message = data.detail || '数据验证失败'
          break
        case 429:
          message = '操作太频繁，请稍后再试'
          break
        case 500:
          message = '服务器错误，请稍后重试'
          break
        case 502:
        case 503:
        case 504:
          message = '服务暂时不可用，请稍后重试'
          break
        default:
          message = data.message || '操作失败'
      }
    } else if (error.request) {
      // 请求已发出但没有收到响应
      message = '网络连接失败，请检查网络'
      code = -1
    } else {
      // 其他错误
      message = error.message || '操作失败'
      code = -2
    }

    // 显示提示
    if (showToast) {
      uni.showToast({
        title: message,
        icon: 'none',
        duration: duration
      })
    }

    // 记录错误日志
    if (logError) {
      this.logError({
        code,
        message,
        error,
        timestamp: new Date().toISOString()
      })
    }

    return {
      code,
      message,
      error
    }
  }

  /**
   * 处理业务错误
   * @param {string} message - 错误消息
   * @param {Object} options - 配置选项
   */
  static handleBusinessError(message, options = {}) {
    const {
      showToast = true,
      duration = 2000,
      icon = 'none'
    } = options

    if (showToast) {
      uni.showToast({
        title: message,
        icon: icon,
        duration: duration
      })
    }

    this.logError({
      type: 'business',
      message,
      timestamp: new Date().toISOString()
    })
  }

  /**
   * 记录错误日志
   * @param {Object} errorInfo - 错误信息
   */
  static logError(errorInfo) {
    // 开发环境打印到控制台
    if (process.env.NODE_ENV === 'development') {
      console.error('[Error]', errorInfo)
    }

    // 生产环境可以上报到错误监控平台
    if (process.env.NODE_ENV === 'production') {
      // TODO: 上报到错误监控平台（如Sentry）
      // this.reportToSentry(errorInfo)
    }

    // 保存到本地存储（最多保存100条）
    try {
      let errorLogs = uni.getStorageSync('error_logs') || []
      errorLogs.unshift(errorInfo)
      
      // 只保留最近100条
      if (errorLogs.length > 100) {
        errorLogs = errorLogs.slice(0, 100)
      }
      
      uni.setStorageSync('error_logs', errorLogs)
    } catch (e) {
      console.error('保存错误日志失败:', e)
    }
  }

  /**
   * 获取错误日志
   * @returns {Array} 错误日志列表
   */
  static getErrorLogs() {
    try {
      return uni.getStorageSync('error_logs') || []
    } catch (e) {
      return []
    }
  }

  /**
   * 清空错误日志
   */
  static clearErrorLogs() {
    try {
      uni.removeStorageSync('error_logs')
    } catch (e) {
      console.error('清空错误日志失败:', e)
    }
  }

  /**
   * 处理表单验证错误
   * @param {Object} errors - 验证错误对象
   * @returns {string} 第一个错误消息
   */
  static handleValidationErrors(errors) {
    if (!errors || typeof errors !== 'object') {
      return '表单验证失败'
    }

    const firstError = Object.values(errors)[0]
    const message = Array.isArray(firstError) ? firstError[0] : firstError

    uni.showToast({
      title: message,
      icon: 'none',
      duration: 2000
    })

    return message
  }

  /**
   * 显示成功提示
   * @param {string} message - 成功消息
   * @param {number} duration - 显示时长
   */
  static showSuccess(message, duration = 2000) {
    uni.showToast({
      title: message,
      icon: 'success',
      duration: duration
    })
  }

  /**
   * 显示加载中
   * @param {string} title - 加载提示文字
   */
  static showLoading(title = '加载中...') {
    uni.showLoading({
      title: title,
      mask: true
    })
  }

  /**
   * 隐藏加载中
   */
  static hideLoading() {
    uni.hideLoading()
  }

  /**
   * 显示确认对话框
   * @param {Object} options - 配置选项
   * @returns {Promise} Promise对象
   */
  static showConfirm(options = {}) {
    const {
      title = '提示',
      content = '',
      confirmText = '确定',
      cancelText = '取消',
      confirmColor = '#667eea'
    } = options

    return new Promise((resolve, reject) => {
      uni.showModal({
        title,
        content,
        confirmText,
        cancelText,
        confirmColor,
        success: (res) => {
          if (res.confirm) {
            resolve(true)
          } else {
            resolve(false)
          }
        },
        fail: (err) => {
          reject(err)
        }
      })
    })
  }
}

export default ErrorHandler
