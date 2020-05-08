import {
  login,
  logout,
  getMenu
  // getUserInfo
} from '@/api/admin/login'
import {
  getToken,
  setToken,
  removeToken
} from '@/utils/auth'
// import { MessageBox } from 'element-ui'
import {
// menusConverter
} from '@/utils'

const user = {
  state: {
    user: '',
    status: '',
    code: '',
    token: getToken(),
    name: '',
    avatar: '',
    introduction: '',
    loadedMenus: false,
    roles: [],
    setting: {
      articlePlatform: []
    }
  },

  mutations: {
    SET_CODE: (state, code) => {
      state.code = code
    },
    SET_TOKEN: (state, token) => {
      state.token = token
      setToken(token)
    },
    SET_INTRODUCTION: (state, introduction) => {
      state.introduction = introduction
    },
    SET_SETTING: (state, setting) => {
      state.setting = setting
    },
    SET_STATUS: (state, status) => {
      state.status = status
    },
    SET_NAME: (state, name) => {
      state.name = name
    },
    SET_AVATAR: (state, avatar) => {
      state.avatar = avatar
    },
    SET_ROLES: (state, roles) => {
      state.roles = roles
    },
    SET_LOADED_MENUS: (state, flag) => {
      state.loadedMenus = flag
    }
  },

  actions: {
    // 用户名登录
    LoginByUsername({
      commit
    }, userInfo) {
      return new Promise((resolve, reject) => {
        // request api
        login(userInfo).then(response => {
          const data = response.data
          commit('SET_TOKEN', data.access_token)
          commit('SET_NAME', data.username)
          resolve()
        }).catch(error => {
          reject(error)
        })
      })
    },

    GetMenu({
      commit
    }) {
      return new Promise((resolve, reject) => {
        getMenu().then(response => {
          // 拉取了菜单的设置
          commit('SET_LOADED_MENUS', true)
          const powerMenus = response.data.result || []
          if (powerMenus.length > 0) {
            resolve(response)
          } else {
            // 后端返回菜单列表为空, 无权限
            reject(' user has insufficient permissions!')
          }
        })
      })
    },

    // 登出
    LogOut({
      commit,
      state
    }) {
      return new Promise((resolve, reject) => {
        logout(state.token).then(() => {
          commit('SET_TOKEN', '')
          commit('SET_ROLES', [])
          commit('SET_MENUS', [])
          removeToken()
          resolve()
        }).catch(error => {
          reject(error)
        })
      })
    },

    // 前端 登出
    FedLogOut({
      commit
    }) {
      return new Promise(resolve => {
        commit('SET_TOKEN', '')
        removeToken()
        resolve()
      })
    }
  }
}

export default user
