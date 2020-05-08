import request from './request'
import qs from 'qs'

/**
 * 登录接口
 */
export function login(data) {
  return request({
    url: '/user/login',
    method: 'post',
    data: qs.stringify(data)
  })
}

// 退出
export function logout() {
  return request({
    url: '/login/logout',
    method: 'post'
  })
}

// 集成系统 暂弃
// 获取用户信息
export function getUserInfo() {
  return request({
    url: '/user/info',
    method: 'GET'
  })
}

/**
 * lime系统
 * 获取授权的菜单接口
 */
export function getMenu() {
  return request({
    url: 'v1/user/menu',
    method: 'GET',
    params: {
      domain: 'lime'
    }
  })
}

/**
 * 获取用户权限列表接口
 */
export function permList() {
  return request({
    url: 'v1/user/perm/list',
    method: 'GET',
    params: {
      code: 'lime'
    }
  })
}
