/**
 * 图片上传工具
 * 支持腾讯云COS和阿里云OSS
 */

import { API_BASE_URL } from '@/api/config'
import request from '@/api/request'

/**
 * 上传配置
 * 请根据实际使用的云存储服务修改
 */
const UPLOAD_CONFIG = {
  // 上传方式: 'server' | 'cos' | 'oss'
  // 'server' - 直接上传到后端服务器
  // 'cos' - 腾讯云COS
  // 'oss' - 阿里云OSS
  method: 'server',
  
  // 服务器上传地址
  serverUrl: `${API_BASE_URL}/api/v1/upload`,
  
  // 腾讯云COS配置
  cos: {
    bucket: 'your-bucket-name',
    region: 'ap-guangzhou',
    secretId: '',
    secretKey: ''
  },
  
  // 阿里云OSS配置
  oss: {
    bucket: 'your-bucket-name',
    region: 'oss-cn-hangzhou',
    accessKeyId: '',
    accessKeySecret: ''
  },
  
  // 图片限制
  maxSize: 10 * 1024 * 1024, // 10MB
  allowedTypes: ['image/jpeg', 'image/png', 'image/gif', 'image/webp']
}

/**
 * 选择图片
 * @param {Object} options - 选择选项
 * @returns {Promise<Array>} 图片临时路径数组
 */
export function chooseImage(options = {}) {
  return new Promise((resolve, reject) => {
    uni.chooseImage({
      count: options.count || 1,
      sizeType: options.sizeType || ['compressed'],
      sourceType: options.sourceType || ['album', 'camera'],
      success: (res) => {
        resolve(res.tempFilePaths)
      },
      fail: (err) => {
        reject(err)
      }
    })
  })
}

/**
 * 上传图片到服务器
 * @param {String} filePath - 图片临时路径
 * @param {Function} onProgress - 上传进度回调
 * @returns {Promise<String>} 图片URL
 */
async function uploadToServer(filePath, onProgress) {
  return new Promise((resolve, reject) => {
    const uploadTask = uni.uploadFile({
      url: UPLOAD_CONFIG.serverUrl,
      filePath: filePath,
      name: 'file',
      header: {
        'Authorization': `Bearer ${uni.getStorageSync('token')}`
      },
      success: (res) => {
        if (res.statusCode === 200) {
          const data = JSON.parse(res.data)
          if (data.url) {
            resolve(data.url)
          } else {
            reject(new Error(data.message || '上传失败'))
          }
        } else {
          reject(new Error('上传失败'))
        }
      },
      fail: (err) => {
        reject(err)
      }
    })

    // 监听上传进度
    if (onProgress) {
      uploadTask.onProgressUpdate((res) => {
        onProgress(res.progress)
      })
    }
  })
}

/**
 * 上传图片到腾讯云COS
 * @param {String} filePath - 图片临时路径
 * @param {Function} onProgress - 上传进度回调
 * @returns {Promise<String>} 图片URL
 */
async function uploadToCOS(filePath, onProgress) {
  // 1. 从后端获取临时密钥
  const credentials = await request.get('/upload/cos-credentials')
  
  // 2. 生成文件名
  const fileName = `images/${Date.now()}_${Math.random().toString(36).slice(2)}.jpg`
  
  // 3. 使用uni-app的上传API
  return new Promise((resolve, reject) => {
    const uploadTask = uni.uploadFile({
      url: `https://${UPLOAD_CONFIG.cos.bucket}.cos.${UPLOAD_CONFIG.cos.region}.myqcloud.com/${fileName}`,
      filePath: filePath,
      name: 'file',
      formData: {
        key: fileName,
        'x-cos-security-token': credentials.sessionToken,
        'x-cos-meta-fileid': fileName
      },
      header: {
        'Authorization': credentials.authorization
      },
      success: (res) => {
        if (res.statusCode === 200) {
          const url = `https://${UPLOAD_CONFIG.cos.bucket}.cos.${UPLOAD_CONFIG.cos.region}.myqcloud.com/${fileName}`
          resolve(url)
        } else {
          reject(new Error('上传失败'))
        }
      },
      fail: (err) => {
        reject(err)
      }
    })

    if (onProgress) {
      uploadTask.onProgressUpdate((res) => {
        onProgress(res.progress)
      })
    }
  })
}

/**
 * 上传图片到阿里云OSS
 * @param {String} filePath - 图片临时路径
 * @param {Function} onProgress - 上传进度回调
 * @returns {Promise<String>} 图片URL
 */
async function uploadToOSS(filePath, onProgress) {
  // 1. 从后端获取STS临时凭证
  const credentials = await request.get('/upload/oss-credentials')
  
  // 2. 生成文件名
  const fileName = `images/${Date.now()}_${Math.random().toString(36).slice(2)}.jpg`
  
  // 3. 使用uni-app的上传API
  return new Promise((resolve, reject) => {
    const uploadTask = uni.uploadFile({
      url: `https://${UPLOAD_CONFIG.oss.bucket}.${UPLOAD_CONFIG.oss.region}.aliyuncs.com`,
      filePath: filePath,
      name: 'file',
      formData: {
        key: fileName,
        policy: credentials.policy,
        OSSAccessKeyId: credentials.accessKeyId,
        signature: credentials.signature,
        'x-oss-security-token': credentials.securityToken
      },
      success: (res) => {
        if (res.statusCode === 200 || res.statusCode === 204) {
          const url = `https://${UPLOAD_CONFIG.oss.bucket}.${UPLOAD_CONFIG.oss.region}.aliyuncs.com/${fileName}`
          resolve(url)
        } else {
          reject(new Error('上传失败'))
        }
      },
      fail: (err) => {
        reject(err)
      }
    })

    if (onProgress) {
      uploadTask.onProgressUpdate((res) => {
        onProgress(res.progress)
      })
    }
  })
}

/**
 * 验证图片
 * @param {String} filePath - 图片路径
 * @returns {Promise<Boolean>}
 */
async function validateImage(filePath) {
  return new Promise((resolve, reject) => {
    uni.getFileInfo({
      filePath: filePath,
      success: (res) => {
        // 检查文件大小
        if (res.size > UPLOAD_CONFIG.maxSize) {
          reject(new Error(`图片大小不能超过${UPLOAD_CONFIG.maxSize / 1024 / 1024}MB`))
          return
        }
        resolve(true)
      },
      fail: (err) => {
        reject(err)
      }
    })
  })
}

/**
 * 上传单张图片
 * @param {String} filePath - 图片临时路径
 * @param {Function} onProgress - 上传进度回调
 * @returns {Promise<String>} 图片URL
 */
export async function uploadImage(filePath, onProgress) {
  try {
    // 验证图片
    await validateImage(filePath)
    
    // 根据配置选择上传方式
    switch (UPLOAD_CONFIG.method) {
      case 'cos':
        return await uploadToCOS(filePath, onProgress)
      case 'oss':
        return await uploadToOSS(filePath, onProgress)
      case 'server':
      default:
        return await uploadToServer(filePath, onProgress)
    }
  } catch (error) {
    console.error('上传失败:', error)
    throw error
  }
}

/**
 * 上传多张图片
 * @param {Array} filePaths - 图片临时路径数组
 * @param {Function} onProgress - 上传进度回调 (current, total)
 * @returns {Promise<Array>} 图片URL数组
 */
export async function uploadImages(filePaths, onProgress) {
  const urls = []
  const total = filePaths.length
  
  for (let i = 0; i < total; i++) {
    try {
      const url = await uploadImage(filePaths[i], (progress) => {
        if (onProgress) {
          onProgress(i + 1, total, progress)
        }
      })
      urls.push(url)
    } catch (error) {
      console.error(`第${i + 1}张图片上传失败:`, error)
      throw error
    }
  }
  
  return urls
}

/**
 * 选择并上传图片
 * @param {Object} options - 选择选项
 * @param {Function} onProgress - 上传进度回调
 * @returns {Promise<Array>} 图片URL数组
 */
export async function chooseAndUploadImage(options = {}, onProgress) {
  try {
    // 选择图片
    const filePaths = await chooseImage(options)
    
    // 上传图片
    const urls = await uploadImages(filePaths, onProgress)
    
    return urls
  } catch (error) {
    console.error('选择或上传图片失败:', error)
    throw error
  }
}

export default {
  chooseImage,
  uploadImage,
  uploadImages,
  chooseAndUploadImage
}
