import request from './request.js'

/**
 * 出价（心愿单）相关API
 */

// 创建出价
export function createOffer(data) {
  return request({
    url: '/offers',
    method: 'POST',
    data
  })
}

// 获取我的出价列表
export function getMyOffers(params) {
  return request({
    url: '/offers/my',
    method: 'GET',
    params
  })
}

// 获取收到的出价列表
export function getReceivedOffers(params) {
  return request({
    url: '/offers/received',
    method: 'GET',
    params
  })
}

// 接受出价
export function acceptOffer(id) {
  return request({
    url: `/offers/${id}/accept`,
    method: 'POST'
  })
}

// 拒绝出价
export function rejectOffer(id, data) {
  return request({
    url: `/offers/${id}/reject`,
    method: 'POST',
    data
  })
}

// 取消出价
export function cancelOffer(id) {
  return request({
    url: `/offers/${id}/cancel`,
    method: 'POST'
  })
}

// 获取出价详情
export function getOfferDetail(id) {
  return request({
    url: `/offers/${id}`,
    method: 'GET'
  })
}

// 获取作品的出价列表
export function getArtworkOffers(artworkInstanceId, params) {
  return request({
    url: `/offers/artwork/${artworkInstanceId}`,
    method: 'GET',
    params
  })
}
