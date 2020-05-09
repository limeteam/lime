import request from './request'
// import qs from 'qs'

/**
 * 获取用户信息接口
 */
export function getInfo() {
  return request({
    url: 'v1/users/info',
    method: 'GET',
    params: {
      domain: 'zeus-config'
    }
  })
}

export function login(data) {
  return request({
    url: 'v1/users/login',
    method: 'POST',
    data
  })
}

export function logout() {
  return request({
    url: 'v1/users/logout',
    method: 'GET',
    params: {
      domain: 'zeus-config'
    }
  })
}

