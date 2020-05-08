/**
 * 通用 ajax 请求方法
 * TODO:
 */
import axios from 'axios'
import store from '@/store'
// 不全局错误提示的枚举
import {
  NOT_SHOW_MSG,
  TOKEN_EXPIRATION
} from '../enum'
import {
  Message,
  MessageBox
} from 'element-ui'

// create an axios instance
const service = axios.create({
  // withCredentials: true,
  // baseURL: process.env.BASE_API, // api 的 base_url
  baseURL: process.env['VUE_APP_ADMIN'],
  // request timeout
  timeout: 15000
  // 默认表单提交的方式
  // headers: { 'content-Type': 'application/x-www-form-urlencoded' }
  // config.headers['content-Type'] = 'application/x-www-form-urlencoded'
})

// request interceptor
service.interceptors.request.use(
  config => {
    // Do something before request is sent
    const token = store.getters.token
    if (token) {
      // 让每个请求携带token
      config.headers['Authorization'] = 'Bearer ' + token
    }
    return config
  },
  error => {
    // Do something with request error
    console.warn(error) // for debug
    Promise.reject(error)
  }
)

// TODO: 以下需要具体修改
// response interceptor
service.interceptors.response.use(
  /**
   * response => response,
   * 下面的注释为通过在response里，自定义code来标示请求状态
   * 当code返回如下情况则说明权限有问题，登出并返回到登录页
   * 如想通过 xmlhttprequest 来状态码标识 逻辑可写在下面error中
   * 以下代码均为样例，请结合自生需求加以修改，若不需要，则可删除
   */
  response => {
    // 常见业务错误
    const res = response.data
    console.log(res)

    // 102001003 需要
    if (res.code !== 200) {
      // 登录失效
      if (TOKEN_EXPIRATION.includes(res.code)) {
        // 请自行在引入 MessageBox
        MessageBox.alert('token expired', 'Tips', {
          confirmButtonText: 'login',
          cancelButtonText: 'cancel',
          type: 'warning'
        }).then(() => {
          store.dispatch('FedLogOut').then(() => {
            // 为了重新实例化vue-router对象 避免bug
            location.reload()
          })
        })
        return Promise.reject('token expired！')
      }

      // 如果不全局提示错误，返回失败，catch 处理
      if (NOT_SHOW_MSG.includes(res.code)) {
        return Promise.reject(res)
      }

      // 其他错误异常
      Message({
        message: res.msg || 'server error!',
        type: 'error',
        duration: 5 * 1000
      })
      return Promise.reject(res)
    } else {
      // 通过
      return response.data
    }
  },
  error => {
    console.log('server error:' + error) // for debug
    return Promise.reject(error)
  }
)

export default service
