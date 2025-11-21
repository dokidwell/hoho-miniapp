import request from './request.js'

/**
 * 任务相关API
 */

// 获取任务列表
export function getTasks(params) {
  return request({
    url: '/tasks',
    method: 'GET',
    params
  })
}

// 获取我的任务完成情况
export function getMyTaskCompletions(params) {
  return request({
    url: '/tasks/my',
    method: 'GET',
    params
  })
}

// 完成任务
export function completeTask(taskId, data) {
  return request({
    url: `/tasks/${taskId}/complete`,
    method: 'POST',
    data
  })
}

// 领取任务奖励
export function claimTaskReward(taskId) {
  return request({
    url: `/tasks/${taskId}/claim`,
    method: 'POST'
  })
}

// 每日签到
export function dailyCheckIn() {
  return request({
    url: '/tasks/checkin',
    method: 'POST'
  })
}

// 获取签到状态
export function getCheckInStatus() {
  return request({
    url: '/tasks/checkin/status',
    method: 'GET'
  })
}

// 获取任务统计
export function getTaskStats() {
  return request({
    url: '/tasks/stats',
    method: 'GET'
  })
}
