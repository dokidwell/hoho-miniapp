// 管理后台API配置
const API_BASE_URL = 'https://api.hohopark.com';

const API_ENDPOINTS = {
  // 认证
  AUTH: {
    LOGIN: '/admin/login',
    LOGOUT: '/admin/logout'
  },
  
  // 创作审核
  CREATION: {
    LIST: '/admin/creations',
    DETAIL: '/admin/creations/:id',
    APPROVE: '/admin/creations/:id/approve',
    REJECT: '/admin/creations/:id/reject'
  },
  
  // 任务管理
  TASK: {
    LIST: '/admin/tasks',
    CREATE: '/admin/tasks',
    UPDATE: '/admin/tasks/:id',
    DELETE: '/admin/tasks/:id',
    TOGGLE: '/admin/tasks/:id/toggle'
  },
  
  // 公告管理
  ANNOUNCEMENT: {
    LIST: '/admin/announcements',
    CREATE: '/admin/announcements',
    UPDATE: '/admin/announcements/:id',
    DELETE: '/admin/announcements/:id',
    TOGGLE_PIN: '/admin/announcements/:id/toggle-pin'
  },
  
  // 阳光账户
  PLATFORM_ACCOUNT: {
    SUMMARY: '/admin/platform-account',
    TRANSACTIONS: '/admin/platform-account/transactions'
  },
  
  // 系统配置
  CONFIG: {
    GET: '/admin/config',
    UPDATE: '/admin/config'
  },
  
  // 统计数据
  STATS: {
    DASHBOARD: '/admin/stats/dashboard',
    USERS: '/admin/stats/users',
    ARTWORKS: '/admin/stats/artworks',
    TRANSACTIONS: '/admin/stats/transactions'
  }
};

// HTTP请求封装
class APIClient {
  constructor(baseURL) {
    this.baseURL = baseURL;
    this.token = localStorage.getItem('admin_token');
  }
  
  setToken(token) {
    this.token = token;
    localStorage.setItem('admin_token', token);
  }
  
  clearToken() {
    this.token = null;
    localStorage.removeItem('admin_token');
  }
  
  async request(method, url, data = null) {
    const headers = {
      'Content-Type': 'application/json'
    };
    
    if (this.token) {
      headers['Authorization'] = `Bearer ${this.token}`;
    }
    
    const options = {
      method,
      headers
    };
    
    if (data && (method === 'POST' || method === 'PUT' || method === 'PATCH')) {
      options.body = JSON.stringify(data);
    }
    
    try {
      const response = await fetch(this.baseURL + url, options);
      const result = await response.json();
      
      if (!response.ok) {
        throw new Error(result.message || '请求失败');
      }
      
      return result;
    } catch (error) {
      console.error('API请求错误:', error);
      throw error;
    }
  }
  
  get(url) {
    return this.request('GET', url);
  }
  
  post(url, data) {
    return this.request('POST', url, data);
  }
  
  put(url, data) {
    return this.request('PUT', url, data);
  }
  
  delete(url) {
    return this.request('DELETE', url);
  }
}

const apiClient = new APIClient(API_BASE_URL);

export { API_ENDPOINTS, apiClient };
