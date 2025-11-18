import { API_CONFIG } from './config'

/**
 * 统一的HTTP请求方法
 */
class Request {
  constructor(config = {}) {
    this.config = {
      ...API_CONFIG,
      ...config
    }
    this.interceptors = {
      request: [],
      response: []
    }
  }

  /**
   * 发送请求
   */
  async request(method, url, data = null, options = {}) {
    const requestUrl = `${this.config.baseURL}${url}`
    
    // 构建请求配置
    const requestConfig = {
      url: requestUrl,
      method: method.toUpperCase(),
      timeout: options.timeout || this.config.timeout,
      header: {
        ...this.config.headers,
        ...options.headers
      }
    }

    // 添加请求体
    if (data && ['POST', 'PUT', 'PATCH'].includes(requestConfig.method)) {
      requestConfig.data = data
    } else if (data && requestConfig.method === 'GET') {
      // GET请求的参数添加到URL中
      const params = new URLSearchParams(data).toString()
      if (params) {
        requestConfig.url += `?${params}`
      }
    }

    // 获取token并添加到header
    const token = uni.getStorageSync('token')
    if (token) {
      requestConfig.header['Authorization'] = `Bearer ${token}`
    }

    return new Promise((resolve, reject) => {
      uni.request({
        ...requestConfig,
        success: (response) => {
          const { data, statusCode } = response
          
          // 检查HTTP状态码
          if (statusCode >= 200 && statusCode < 300) {
            // 检查业务状态码
            if (data.code === 0 || data.success) {
              resolve(data.data || data)
            } else {
              // 业务错误
              reject({
                code: data.code,
                message: data.message || '请求失败',
                data: data.data
              })
            }
          } else if (statusCode === 401) {
            // 未授权，清除token并跳转到登录页
            uni.removeStorageSync('token')
            uni.navigateTo({ url: '/pages/auth/login' })
            reject({
              code: 401,
              message: '登录已过期，请重新登录'
            })
          } else {
            reject({
              code: statusCode,
              message: data.message || '请求失败',
              data: data
            })
          }
        },
        fail: (error) => {
          reject({
            code: -1,
            message: error.errMsg || '网络请求失败',
            error
          })
        }
      })
    })
  }

  /**
   * GET请求
   */
  get(url, data = null, options = {}) {
    return this.request('GET', url, data, options)
  }

  /**
   * POST请求
   */
  post(url, data = null, options = {}) {
    return this.request('POST', url, data, options)
  }

  /**
   * PUT请求
   */
  put(url, data = null, options = {}) {
    return this.request('PUT', url, data, options)
  }

  /**
   * PATCH请求
   */
  patch(url, data = null, options = {}) {
    return this.request('PATCH', url, data, options)
  }

  /**
   * DELETE请求
   */
  delete(url, data = null, options = {}) {
    return this.request('DELETE', url, data, options)
  }

  /**
   * 文件上传
   */
  upload(url, filePath, name = 'file', formData = {}) {
    return new Promise((resolve, reject) => {
      const uploadUrl = `${this.config.baseURL}${url}`
      const token = uni.getStorageSync('token')
      
      const header = {
        ...this.config.headers
      }
      if (token) {
        header['Authorization'] = `Bearer ${token}`
      }

      uni.uploadFile({
        url: uploadUrl,
        filePath,
        name,
        formData,
        header,
        success: (response) => {
          const data = JSON.parse(response.data)
          if (response.statusCode >= 200 && response.statusCode < 300) {
            if (data.code === 0 || data.success) {
              resolve(data.data || data)
            } else {
              reject({
                code: data.code,
                message: data.message || '上传失败'
              })
            }
          } else {
            reject({
              code: response.statusCode,
              message: data.message || '上传失败'
            })
          }
        },
        fail: (error) => {
          reject({
            code: -1,
            message: error.errMsg || '上传失败'
          })
        }
      })
    })
  }
}

// 创建并导出请求实例
export const request = new Request()

export default request
