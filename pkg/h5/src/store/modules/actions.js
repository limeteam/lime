import * as types from './types'
import state from './state'
import { login, logout, userInfo } from '@/api'
import { setToken, removeToken } from '@/utils/auth'
import md5 from "md5";
const actions = {
    getType({ commit }, type) {
        return new Promise((resolve, reject) => {
            state.types.map(item => {
                if (item.type == type) {
                    resolve(item.title)
                }
            })
        })
    },
    toggleCover({ commit }) {
        commit(types.TITLE_COVER)
    },
    currentCpt({ commit }, num) {
        commit(types.CURRENT_CPT, num)
    },
    preCpt({ commit }) {
        return new Promise((resolve, reject) => {
            resolve(true);
            commit(types.PRE_CPT)
        })
    },
    nextCpt({ commit }, max) {
        return new Promise((resolve, reject) => {
            resolve(true);
            commit(types.NEXT_CPT, max)
        })
    },
    switchNight({ commit }) {
        commit(types.SWITCH_STYLE);
    },
    bigSize({ commit }) {
        commit(types.BIG_SIZE);
    },
    smallSize({ commit }) {
        commit(types.SMALL_SIZE);
    },
    font({ commit }, bool) {
        commit('FONT', bool)
    },
    menus({ commit }) {
        commit(types.MENU);
    },
    currentStyle({ commit }, index) {
        commit(types.DAY_STYLE, index);
    },
    getInfo({ commit }, info) {
        commit(types.PERSON_INFO, info);
    },
    // user login
    login({ commit }, userInfo) {
        const { username, password } = userInfo
        return new Promise((resolve, reject) => {
            login({ username: username.trim(), password: md5(password) }).then(response => {
                const { data } = response
                commit('SET_TOKEN', data.result.access_token)
                setToken(data.result.access_token)
                resolve()
            }).catch(error => {
                reject(error)
            })
        })
    },

    // get user info
    getInfo({ commit, state }) {
        return new Promise((resolve, reject) => {
            userInfo(state.token).then(response => {
                const { data } = response

                if (!data) {
                    reject('Verification failed, please Login again.')
                }

                const { username, faceicon } = data.result

                commit('SET_NAME', username)
                commit('SET_FACEICON', faceicon)
                resolve(data)
            }).catch(error => {
                reject(error)
            })
        })
    },

    // user logout
    logout({ commit, state, dispatch }) {
        return new Promise((resolve, reject) => {
            logout(state.token).then(() => {
                commit('SET_TOKEN', '')
                removeToken()
                resetRouter()
                dispatch('tagsView/delAllViews', null, { root: true })

                resolve()
            }).catch(error => {
                reject(error)
            })
        })
    },

    // remove token
    resetToken({ commit }) {
        return new Promise(resolve => {
            commit('SET_TOKEN', '')
            removeToken()
            resolve()
        })
    }
}
export default actions;
