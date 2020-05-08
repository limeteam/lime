/**
 * 任务列表
 */
import {
  taskInfo
} from '@/api/lime-admin/task'

const task = {
  state: {
    id: '1', // 任务id
    tasklist: {
      task_name: '', // 任务名称
      status: '1'
    }
  },
  mutations: {
    SET_WITHDRAWAL(state, data) {
      state.withdrawal = data
    }
  },
  actions: {
    getTaskInfo({
      commit
    }, id) {
      return new Promise((resolve, reject) => {
        taskInfo(id).then(res => {
          commit('SET_WITHDRAWAL', res.data)
          resolve(res.data)
        }).catch(error => {
          reject(error)
        })
      })
    }
  }
}

export default task
