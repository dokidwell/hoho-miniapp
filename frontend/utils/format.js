/**
 * 格式化积分显示（8位小数）
 * @param {string|number} points - 积分值
 * @param {number} decimals - 小数位数，默认8位
 * @returns {string} 格式化后的积分
 */
export function formatPoints(points, decimals = 8) {
  if (points === null || points === undefined || points === '') return '0.00000000'
  const num = parseFloat(points)
  if (isNaN(num)) return '0.00000000'
  return num.toFixed(decimals)
}

/**
 * 手机号脱敏
 * @param {string} phone - 手机号
 * @returns {string} 脱敏后的手机号
 */
export function maskPhone(phone) {
  if (!phone) return '未绑定'
  return phone.replace(/(\d{3})\d{4}(\d{4})/, '$1****$2')
}

/**
 * 时间格式化
 * @param {string|number} timestamp - 时间戳
 * @returns {string} 格式化后的时间
 */
export function formatTime(timestamp) {
  if (!timestamp) return '-'
  const date = new Date(timestamp)
  const year = date.getFullYear()
  const month = String(date.getMonth() + 1).padStart(2, '0')
  const day = String(date.getDate()).padStart(2, '0')
  const hour = String(date.getHours()).padStart(2, '0')
  const minute = String(date.getMinutes()).padStart(2, '0')
  return `${year}-${month}-${day} ${hour}:${minute}`
}

/**
 * 短Hash
 * @param {string} hash - 完整Hash
 * @returns {string} 短Hash
 */
export function shortHash(hash) {
  if (!hash) return '-'
  return `${hash.slice(0, 6)}...${hash.slice(-4)}`
}

/**
 * 事件类型名称
 * @param {string} type - 事件类型
 * @returns {string} 事件类型名称
 */
export function getEventTypeName(type) {
  const map = {
    'mint': '铸造',
    'trade': '交易',
    'airdrop': '空投',
    'listing': '挂售',
    'cancel_listing': '取消挂售',
  }
  return map[type] || type
}

/**
 * 计算手续费
 * @param {string|number} price - 价格
 * @param {number} rate - 费率
 * @returns {string} 手续费
 */
export function calculateFee(price, rate) {
  if (!price) return '0.00000000'
  const num = parseFloat(price)
  if (isNaN(num)) return '0.00000000'
  return (num * rate).toFixed(8)
}

/**
 * 计算实际收入（扣除平台费和版税）
 * @param {string|number} price - 价格
 * @returns {string} 实际收入
 */
export function calculateReceived(price) {
  if (!price) return '0.00000000'
  const num = parseFloat(price)
  if (isNaN(num)) return '0.00000000'
  const platformFee = num * 0.025
  const royalty = num * 0.025
  return (num - platformFee - royalty).toFixed(8)
}
