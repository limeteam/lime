import request from './request'

export function wechatSetting(code, data) {
  return request({
    url: '/admin/configs/' + code + '/distributor',
    method: 'post',
    data: data
  })
}

export function getWetchatSetting(code, data) {
  return request({
    url: '/admin/distributor/' + code,
    method: 'GET'
  })
}

export function getTokenAndKey() {
  return request({
    url: '/admin/setting/getTokenAndEncodingAESKey',
    method: 'GET'
  })
}