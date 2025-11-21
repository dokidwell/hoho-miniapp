/**
 * 性能优化工具
 * 提供防抖、节流、图片优化等功能
 */

class Performance {
  /**
   * 防抖函数
   * @param {Function} func - 要执行的函数
   * @param {number} wait - 等待时间(ms)
   * @param {boolean} immediate - 是否立即执行
   * @returns {Function} 防抖后的函数
   */
  static debounce(func, wait = 300, immediate = false) {
    let timeout

    return function executedFunction(...args) {
      const context = this

      const later = () => {
        timeout = null
        if (!immediate) func.apply(context, args)
      }

      const callNow = immediate && !timeout

      clearTimeout(timeout)
      timeout = setTimeout(later, wait)

      if (callNow) func.apply(context, args)
    }
  }

  /**
   * 节流函数
   * @param {Function} func - 要执行的函数
   * @param {number} limit - 时间限制(ms)
   * @returns {Function} 节流后的函数
   */
  static throttle(func, limit = 300) {
    let inThrottle

    return function executedFunction(...args) {
      const context = this

      if (!inThrottle) {
        func.apply(context, args)
        inThrottle = true

        setTimeout(() => {
          inThrottle = false
        }, limit)
      }
    }
  }

  /**
   * 图片压缩
   * @param {string} filePath - 图片路径
   * @param {Object} options - 压缩选项
   * @returns {Promise} Promise对象
   */
  static compressImage(filePath, options = {}) {
    const {
      quality = 80,
      maxWidth = 1080,
      maxHeight = 1080
    } = options

    return new Promise((resolve, reject) => {
      uni.compressImage({
        src: filePath,
        quality: quality,
        compressedWidth: maxWidth,
        compressedHeight: maxHeight,
        success: (res) => {
          resolve(res.tempFilePath)
        },
        fail: (err) => {
          reject(err)
        }
      })
    })
  }

  /**
   * 图片懒加载
   * @param {string} imageUrl - 图片URL
   * @param {string} placeholder - 占位图
   * @returns {string} 处理后的图片URL
   */
  static lazyLoadImage(imageUrl, placeholder = '/static/placeholder.png') {
    if (!imageUrl) {
      return placeholder
    }

    // 检查图片是否在缓存中
    const cacheKey = `img_cache_${imageUrl}`
    const cached = uni.getStorageSync(cacheKey)

    if (cached) {
      return imageUrl
    }

    // 返回占位图，实际图片后台加载
    this.preloadImage(imageUrl).then(() => {
      uni.setStorageSync(cacheKey, true)
    })

    return placeholder
  }

  /**
   * 预加载图片
   * @param {string} imageUrl - 图片URL
   * @returns {Promise} Promise对象
   */
  static preloadImage(imageUrl) {
    return new Promise((resolve, reject) => {
      uni.getImageInfo({
        src: imageUrl,
        success: (res) => {
          resolve(res)
        },
        fail: (err) => {
          reject(err)
        }
      })
    })
  }

  /**
   * 批量预加载图片
   * @param {Array} imageUrls - 图片URL数组
   * @returns {Promise} Promise对象
   */
  static preloadImages(imageUrls) {
    const promises = imageUrls.map(url => this.preloadImage(url))
    return Promise.all(promises)
  }

  /**
   * 数据缓存
   * @param {string} key - 缓存键
   * @param {*} data - 缓存数据
   * @param {number} expire - 过期时间(ms)
   */
  static setCache(key, data, expire = 300000) { // 默认5分钟
    const cacheData = {
      data: data,
      timestamp: Date.now(),
      expire: expire
    }

    try {
      uni.setStorageSync(key, cacheData)
    } catch (e) {
      console.error('设置缓存失败:', e)
    }
  }

  /**
   * 获取缓存
   * @param {string} key - 缓存键
   * @returns {*} 缓存数据，如果过期或不存在则返回null
   */
  static getCache(key) {
    try {
      const cacheData = uni.getStorageSync(key)

      if (!cacheData) {
        return null
      }

      const { data, timestamp, expire } = cacheData

      // 检查是否过期
      if (Date.now() - timestamp > expire) {
        uni.removeStorageSync(key)
        return null
      }

      return data
    } catch (e) {
      console.error('获取缓存失败:', e)
      return null
    }
  }

  /**
   * 清除缓存
   * @param {string} key - 缓存键，不传则清除所有缓存
   */
  static clearCache(key) {
    try {
      if (key) {
        uni.removeStorageSync(key)
      } else {
        uni.clearStorageSync()
      }
    } catch (e) {
      console.error('清除缓存失败:', e)
    }
  }

  /**
   * 延迟执行
   * @param {number} ms - 延迟时间(ms)
   * @returns {Promise} Promise对象
   */
  static delay(ms) {
    return new Promise(resolve => setTimeout(resolve, ms))
  }

  /**
   * 重试机制
   * @param {Function} func - 要执行的函数
   * @param {number} retries - 重试次数
   * @param {number} delay - 重试延迟(ms)
   * @returns {Promise} Promise对象
   */
  static async retry(func, retries = 3, delay = 1000) {
    try {
      return await func()
    } catch (error) {
      if (retries <= 0) {
        throw error
      }

      await this.delay(delay)
      return this.retry(func, retries - 1, delay * 2) // 指数退避
    }
  }

  /**
   * 格式化文件大小
   * @param {number} bytes - 字节数
   * @returns {string} 格式化后的大小
   */
  static formatFileSize(bytes) {
    if (bytes === 0) return '0 B'

    const k = 1024
    const sizes = ['B', 'KB', 'MB', 'GB', 'TB']
    const i = Math.floor(Math.log(bytes) / Math.log(k))

    return parseFloat((bytes / Math.pow(k, i)).toFixed(2)) + ' ' + sizes[i]
  }

  /**
   * 获取性能指标
   * @returns {Object} 性能指标对象
   */
  static getPerformanceMetrics() {
    const performance = uni.getPerformance()

    if (!performance) {
      return null
    }

    return {
      // 导航时间
      navigationStart: performance.navigationStart,
      // 页面加载完成时间
      loadComplete: performance.loadEventEnd - performance.navigationStart,
      // DOM解析时间
      domParsing: performance.domContentLoadedEventEnd - performance.domContentLoadedEventStart,
      // 资源加载时间
      resourceLoading: performance.loadEventStart - performance.domContentLoadedEventEnd
    }
  }

  /**
   * 记录性能日志
   * @param {string} label - 标签
   * @param {Function} func - 要执行的函数
   * @returns {*} 函数执行结果
   */
  static async measurePerformance(label, func) {
    const startTime = Date.now()

    try {
      const result = await func()
      const endTime = Date.now()
      const duration = endTime - startTime

      console.log(`[Performance] ${label}: ${duration}ms`)

      return result
    } catch (error) {
      const endTime = Date.now()
      const duration = endTime - startTime

      console.error(`[Performance] ${label} failed after ${duration}ms:`, error)

      throw error
    }
  }

  /**
   * 批处理
   * @param {Array} items - 要处理的项目数组
   * @param {Function} handler - 处理函数
   * @param {number} batchSize - 批次大小
   * @returns {Promise} Promise对象
   */
  static async batchProcess(items, handler, batchSize = 10) {
    const results = []

    for (let i = 0; i < items.length; i += batchSize) {
      const batch = items.slice(i, i + batchSize)
      const batchResults = await Promise.all(batch.map(handler))
      results.push(...batchResults)

      // 避免阻塞UI
      await this.delay(10)
    }

    return results
  }
}

export default Performance
