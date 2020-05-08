import {
  asyncRouterMap,
  constantRouterMap
} from '@/router'

// import { type } from 'os'
const permission = {
  state: {
    routers: constantRouterMap,
    addRouters: []
  },
  mutations: {
    SET_ROUTERS: (state, routers) => {
      state.addRouters = routers
      state.routers = constantRouterMap.concat(routers)
    }
  },
  actions: {
    GenerateRoutes({
      commit
    }, powerMenus) {
      // TODO:
      // 修改权限判断 Wendell Sheh
      return new Promise(resolve => {
        const accessedRouters = filterAsyncRouter(asyncRouterMap, powerMenus)
        // const accessedRouters = asyncRouterMap // 不校验权限
        commit('SET_ROUTERS', accessedRouters)
        resolve()
      })
    }
  }
}

/**
 * 通过meta.role判断是否与当前用户权限匹配
 * @param menus
 * @param route
 */
function hasPermission(menus, route) {
  return menus.some(item => {
    return item.url === route.path
  })
}

/**
 * 递归过滤异步路由表，返回符合用户角色权限的路由表
 * @param routes asyncRouterMap
 * @param menus
 */
function filterAsyncRouter(routes, menus) {
  const res = []
  // debugger
  routes.forEach(route => {
    const tmp = {
      ...route
    }
    if (hasPermission(menus, tmp)) {
      if (tmp.children && tmp.children.length > 0) {
        tmp.children = filterAsyncRouter(tmp.children, menus)
      }
      res.push(tmp)
    }
  })
  return res
}

export default permission
