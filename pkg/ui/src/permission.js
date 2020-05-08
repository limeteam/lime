import router from './router'
import store from './store'
import {
  Message
} from 'element-ui'
import {
  getDomainHost
} from '@/utils'
import NProgress from 'nprogress' // progress bar
import 'nprogress/nprogress.css' // progress bar style
import {
  getToken
} from '@/utils/auth' // getToken from cookie

NProgress.configure({
  showSpinner: false
}) // NProgress Configuration

// 判断路由权限方法
// permission judge function
// function hasPermission(roles, permissionRoles) {
//   if (roles.indexOf('admin') >= 0) return true // admin permission passed directly
//   if (!permissionRoles) return true
//   return roles.some(role => permissionRoles.indexOf(role) >= 0)
// }

const whiteList = ['/login', '/auth-redirect'] // no redirect whitelist

router.beforeEach((to, from, next) => {
  // start progress bar
  NProgress.start()
  // determine if there has token
  if (getToken()) {
    /* has token*/
    if (to.path === '/login') {
      next({
        path: '/'
      })
      // if current page is dashboard will not trigger	afterEach hook, so manually handle it
      NProgress.done()
    } else {
      // 判断是否拉取了菜单
      if (!store.getters.loadedMenus) {
        // 拉取菜单
        store
          .dispatch('GetMenu')
          .then(res => {
            // 权限菜单
            const powerMenus = res.data.result
            // 生成可访问菜单
            store.dispatch('GenerateRoutes', powerMenus).then(() => {
              // 动态添加可访问路由表
              router.addRoutes(store.getters.addRouters)
              // hack方法 确保addRoutes已完成 ,set the replace: true so the navigation will not leave a history record
              next({
                ...to,
                replace: true
              })
            })
          })
          .catch(error => {
            // 拉取失败，返回登录
            store.dispatch('FedLogOut').then(() => {
              Message.error(error || 'Verification failed, please login again')
              next({
                path: '/'
              })
            })
          })
          .finally(res => {})
      } else {
        next()
      }
    }
  } else {
    /* has no token*/
    if (whiteList.indexOf(to.path) !== -1) {
      // 在免登录白名单，直接进入
      next()
    } else {
      /**
       * TODO: 修改登录授权逻辑
       * @author Wendell Sheh
       * 如果本地开发环境直接到达登录接口
       * 如果打包环境 自动跳转到鉴权中心做登录操作
       */
      const domainHost = getDomainHost()
      const currURL = encodeURIComponent(location.href)
      if (process.env.NODE_ENV === 'production') {
        if (location.host.startsWith('stage.')) {
          location.href = `//stage.auth.${domainHost.domain}/login?redirectURL=${currURL}`
        } else {
          location.href = `//auth.${domainHost.domain}/login?redirectURL=${currURL}`
        }
      } else {
        location.href = `//localhost:9527/login?redirectURL=${currURL}`
      }
      NProgress.done()
    }
  }
})

router.afterEach(() => {
  // finish progress bar
  NProgress.done()
})
