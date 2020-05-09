import request from './request'

export function getConfigList() {
  return request({
    url: 'configlist/list',
    method: 'GET'
  })
}
