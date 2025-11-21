import request from './request.js'

/**
 * 平台账户（阳光账户）相关API
 */

// 获取平台账户信息
export function getPlatformAccount() {
  return request({
    url: '/platform/account',
    method: 'GET'
  })
}

// 获取平台交易记录
export function getPlatformTransactions(params) {
  return request({
    url: '/platform/transactions',
    method: 'GET',
    params
  })
}

// 获取平台统计数据
export function getPlatformStats(params) {
  return request({
    url: '/platform/stats',
    method: 'GET',
    params
  })
}

// 获取透明公示数据
export function getTransparentLedger(type, params) {
  return request({
    url: `/platform/ledger/${type}`,
    method: 'GET',
    params
  })
}
