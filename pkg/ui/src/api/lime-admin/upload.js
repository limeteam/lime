import request from './request'

export function getQiniuToken() {
  return request({
    url: '/admin/upload/qiniuToken',
    method: 'GET'
  })
}
