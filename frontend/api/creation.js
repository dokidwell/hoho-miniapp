import request from './request.js'

/**
 * 创作相关API
 */

// 创建创作
export function createCreation(data) {
  return request({
    url: '/creations',
    method: 'POST',
    data
  })
}

// 上传创作图片
export function uploadCreationImage(filePath) {
  return new Promise((resolve, reject) => {
    uni.uploadFile({
      url: request.baseURL + '/creations/upload',
      filePath: filePath,
      name: 'image',
      header: {
        'Authorization': 'Bearer ' + uni.getStorageSync('token')
      },
      success: (res) => {
        const data = JSON.parse(res.data)
        if (data.error) {
          reject(data.error)
        } else {
          resolve(data)
        }
      },
      fail: reject
    })
  })
}

// 获取我的创作列表
export function getMyCreations(params) {
  return request({
    url: '/creations/my',
    method: 'GET',
    params
  })
}

// 获取创作详情
export function getCreationDetail(id) {
  return request({
    url: `/creations/${id}`,
    method: 'GET'
  })
}

// 获取创作统计
export function getCreationStats() {
  return request({
    url: '/creations/stats',
    method: 'GET'
  })
}

// 获取社区创作列表（已发布的）
export function getCommunityCreations(params) {
  return request({
    url: '/creations/community',
    method: 'GET',
    params
  })
}

// 购买创作作品
export function purchaseCreation(creationId, data) {
  return request({
    url: `/creations/${creationId}/purchase`,
    method: 'POST',
    data
  })
}
