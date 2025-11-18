// API配置文件
const API_BASE_URL = process.env.NODE_ENV === 'production' 
  ? 'https://api.hoho.app'
  : 'http://localhost:8080'

const API_VERSION = 'v1'

export const API_CONFIG = {
  baseURL: `${API_BASE_URL}/api/${API_VERSION}`,
  timeout: 10000,
  headers: {
    'Content-Type': 'application/json'
  }
}

// API端点定义
export const API_ENDPOINTS = {
  // 用户相关
  USER: {
    REGISTER: '/users/register',
    LOGIN: '/users/login',
    LOGOUT: '/users/logout',
    PROFILE: '/users/profile',
    UPDATE_PROFILE: '/users/profile',
    VERIFY_IDENTITY: '/users/verify-identity',
    GET_POINTS: '/users/points'
  },
  
  // 藏品相关
  ASSET: {
    LIST: '/assets',
    CREATE: '/assets',
    GET_DETAIL: (id) => `/assets/${id}`,
    UPDATE: (id) => `/assets/${id}`,
    DELETE: (id) => `/assets/${id}`,
    GET_INSTANCES: (id) => `/assets/${id}/instances`,
    GET_JINGTAN: '/assets/jingtan',
    SYNC_JINGTAN: '/assets/jingtan/sync'
  },
  
  // 交易相关
  TRADE: {
    LIST: '/trades',
    GET_DETAIL: (id) => `/trades/${id}`,
    CREATE_LISTING: '/listings',
    GET_LISTINGS: '/listings',
    GET_LISTING_DETAIL: (id) => `/listings/${id}`,
    CANCEL_LISTING: (id) => `/listings/${id}/cancel`,
    EXECUTE_TRADE: '/trades/execute',
    GET_HISTORY: '/trades/history'
  },
  
  // 积分相关
  POINTS: {
    GET_BALANCE: '/points/balance',
    GET_HISTORY: '/points/history',
    AIRDROP: '/points/airdrop'
  },
  
  // 社区事件
  EVENT: {
    LIST: '/events',
    GET_DETAIL: (id) => `/events/${id}`
  },
  
  // 第三方账户
  THIRD_PARTY: {
    LIST: '/third-party',
    BIND: '/third-party/bind',
    UNBIND: (platform) => `/third-party/${platform}/unbind`,
    GET_JINGTAN_ASSETS: '/third-party/jingtan/assets'
  }
}

export default API_CONFIG
