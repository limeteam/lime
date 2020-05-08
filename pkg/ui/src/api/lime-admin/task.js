import request from './request'

// 获取任务列表
export function taskList(params) {
  return request({
    url: '/tasks',
    method: 'get',
    params
  })
}

// 获取任务详情
export function taskInfo(id) {
  return request({
    url: '/tasks/' + id + '/detail',
    method: 'get'
  })
}

// 查询所有rules
export function getRules() {
  return request({
    url: '/rules',
    method: 'get'
  })
}

// 根据id停止任务
export function stopTask(id) {
  return request({
    url: '/tasks/' + id + '/stop',
    method: 'put'
  })
}
// 根据id启动非定时任务
export function startTask(id) {
  return request({
    url: '/tasks/' + id + '/start',
    method: 'put'
  })
}
// 根据id重启定时任务
export function restartTask(id) {
  return request({
    url: '/tasks/' + id + '/restart',
    method: 'put'
  })
}
// 添加数据
export function saveTask(data) {
  return request({
    url: '/tasks',
    method: 'post',
    data
  })
}

// 修改数据
export function updateTask(id, data) {
  return request({
    url: '/tasks/' + id,
    method: 'put',
    data
  })
}

