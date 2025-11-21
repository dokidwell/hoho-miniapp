import request from './request.js'

/**
 * 公告相关API
 */

// 获取公告列表
export function getAnnouncements(params) {
  return request({
    url: '/announcements',
    method: 'GET',
    params
  })
}

// 获取公告详情
export function getAnnouncementDetail(id) {
  return request({
    url: `/announcements/${id}`,
    method: 'GET'
  })
}

// 标记公告为已读
export function markAnnouncementAsRead(id) {
  return request({
    url: `/announcements/${id}/read`,
    method: 'POST'
  })
}

// 获取未读公告数量
export function getUnreadCount() {
  return request({
    url: '/announcements/unread/count',
    method: 'GET'
  })
}
