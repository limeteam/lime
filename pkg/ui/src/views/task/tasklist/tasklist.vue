<template>
  <div class="task_list">
    <el-card shadow="hover" class="box-card">
      <!-- 过滤条件 -->
      <el-form ref="formData" :inline="true" :model="formData" size="small" class="demo-form-inline">
        <el-form-item :label="$t('task.taskmanage.condition.task_name')" prop="task_name">
          <el-input v-model="formData.task_name" :placeholder="$t('common.ok')" />
        </el-form-item>
        <el-form-item :label="$t('task.taskmanage.condition.status')" prop="status">
          <el-select v-model="formData.status" placeholder="状态" style="width: 100px">
            <el-option v-for="(item,index) in taskStatus" :key="index" :label="item" :value="index" />
          </el-select>
        </el-form-item>
        <el-form-item prop="time_type">
          <el-select v-model="formData.time_type" style="width: 100px">
            <el-option :value="1" :label="$t('task.taskmanage.condition.created_at')" />
            <el-option :value="2" :label="$t('task.taskmanage.condition.updated_at')" />
          </el-select>
        </el-form-item>
        <el-form-item prop="SETime">
          <el-date-picker v-model="SETime" :start-placeholder="$t('task.taskmanage.condition.time_placeholder')" :end-placeholder="$t('task.taskmanage.condition.time_placeholder')" value-format="yyyy-MM-dd HH:mm:ss" type="datetimerange" @change="timeChange" />
        </el-form-item>
        <el-form-item>
          <el-button :loading="loading" type="primary" @click="onSubmit">
            {{ $t('task.taskmanage.condition.inquire') }}
          </el-button>
          <el-button @click="resetForm">
            {{ $t('task.taskmanage.condition.Reset') }}
          </el-button>
          <el-button icon="el-icon-refresh" @click="on_refresh">
            {{ $t('task.taskmanage.condition.on_refresh') }}
          </el-button>
          <router-link :to="{name: 'taskCreate'}" tag="span">
            <el-button type="primary" icon="el-icon-edit" size="small">
              {{ $t('task.taskmanage.condition.create') }}
            </el-button>
          </router-link>
        </el-form-item>
      </el-form>
      <!-- 过滤条件 -->
      <!--列表-->
      <el-table v-loading="loading" :data="items.list" border style="width: 100%">
        <el-table-column :label="$t('task.taskmanage.table.task_name')" prop="task_name" />
        <el-table-column :label="$t('task.taskmanage.table.status')" prop="status" />
        <el-table-column :label="$t('task.taskmanage.table.counts')" prop="counts" />
        <el-table-column :label="$t('task.taskmanage.table.cron_spec')" prop="cron_spec" />
        <el-table-column :label="$t('task.taskmanage.condition.created_at')" prop="created_at" />
        <el-table-column :label="$t('task.taskmanage.table.operation')" fixed="right">
          <template slot-scope="props">
            <!-- <el-button type="text" size="small" @click="handleClick(scope.row)">
              {{ $t('task.taskmanage.table.details') }}
            </el-button> -->
            <router-link :to="{name: 'taskDetail', params: {id: props.row.id}}" tag="span">
              <el-button v-if="props.row.optionbutton & 0b01000" type="warning" size="small" icon="edit">
                {{ $t('task.taskmanage.table.edit') }}
              </el-button>
            </router-link>
            <el-button v-if="props.row.optionbutton & 0b00100" type="danger" size="small" @click="stop(props.row)">
              {{ $t('task.taskmanage.table.stop') }}
            </el-button>
            <el-button v-if="props.row.optionbutton & 0b00010" type="success" size="small" @click="start(props.row)">
              {{ $t('task.taskmanage.table.start') }}
            </el-button>
            <el-button v-if="props.row.optionbutton & 0b00001" type="success" size="small" @click="restart(props.row)">
              {{ $t('task.taskmanage.table.restart') }}
            </el-button>
          </template>
        </el-table-column>
      </el-table>
      <Pagination v-show="items.total_items>10" :total="items.total_items" :page.sync="formData.page" :limit.sync="formData.page_size" @pagination="getTaskList" />
      <!--列表 -->
    </el-card>
  </div>
</template>

<script>
import { taskList, stopTask, startTask, restartTask } from '@/api/lime-admin/task'
import Pagination from '@/components/Pagination'
import { TASKTATUS } from './emun/index.js'

export default {
  name: 'TaskList',
  filters: {
    statuss: function(value) {
      Number(value)
      return TASKTATUS[value]
    }
  },
  components: { Pagination },
  data() {
    return {
      formData: {
        task_name: '',
        status: 0,
        time_type: 1,
        start_time: '',
        end_time: '',
        page: 1,
        page_size: 10
      },
      loading: false,
      items: {
        cur_page: 1,
        total_items: 2,
        list: []
      },
      taskStatus: [
        this.$t('task.taskmanage.status.all'),
        this.$t('task.taskmanage.status.TaskStatusRunning'),
        this.$t('task.taskmanage.status.TaskStatusPaused'),
        this.$t('task.taskmanage.status.TaskStatusStopped'),
        this.$t('task.taskmanage.status.TaskStatusUnexceptedExited'),
        this.$t('task.taskmanage.status.TaskStatusCompleted'),
        this.$t('task.taskmanage.status.TaskStatusRunningTimeout')
      ],
      SETime: []

    }
  },
  computed: {
  },
  created() {
  },
  mounted() {
    this.getTaskList()
  },
  methods: {
    onSubmit() {
      // 查询按钮
      this.formData.page = 1
      this.getTaskList()
    },
    on_refresh() {
      this.getTaskList()
    },
    timeChange(val) {
      // 时间选择
      if (val === null) {
        this.formData.start_time = ''
        this.formData.end_time = ''
      } else {
        this.formData.start_time = val[0]
        this.formData.end_time = val[1]
      }
    },
    resetForm() {
      // 重置
      this.$refs['formData'].resetFields()
      this.SETime = []
      this.formData.start_time = ''
      this.formData.end_time = ''
    },
    handleClick(row) {
      // 详情
      console.log(row.id)
      this.$router.push({
        path: '/taskmanage/details/' + row.id
        // query: { id: row.id }
      })
    },
    currentChange(index) {
      // 分页
      this.formData.page = index
      this.getTaskList()
    },
    async getTaskList() {
      // 获取列表
      this.loading = true
      try {
        const list = await taskList(this.formData)
        this.items.list = list.data.list
        for (const v of this.items.list) {
          v.isCron = 'No'
          if (v.cron_spec) {
            v.isCron = 'Yes'
          }
          // 操作按钮，用5位2进制数表示，每位控制一个按钮是否显示
          // ----0----0----0----0----0----
          // ----|----|----|----|----|----
          // ---详情--修改-停止--启动-重启---
          switch (v.status) {
            case '未知状态':
              v.optionbutton = 0b10000
              break
            case '运行中':
              v.optionbutton = 0b11100
              break
            case '暂停':
              v.optionbutton = 0b10100
              break
            case '停止':
              v.optionbutton = 0b11010
              if (v.cron_spec) { v.optionbutton = 0b11001 }
              break
            case '异常退出':
              v.optionbutton = 0b11010
              if (v.cron_spec) { v.optionbutton = 0b11100 }
              break
            case '完成':
              v.optionbutton = 0b11010
              if (v.cron_spec) { v.optionbutton = 0b11100 }
              break
            case '运行超时':
              v.optionbutton = 0b11010
              if (v.cron_spec) { v.optionbutton = 0b1110 }
              break
            default:
              v.optionbutton = 0b10000
              break
          }
        }
        this.items.total_items = list.data.total
      } finally {
        this.loading = false
      }
    },
    stop(item) {
      this.$confirm('The operating will stop this task, contine?', 'NOTICE', {
        confirmButtonText: 'Yes',
        cancelButtonText: 'Cancel',
        type: 'warning'
      }).then(() => {
        this.load_data = true
        stopTask(item.id).then(() => {
          this.getTaskList()
          this.$message.success('Success!')
        }).catch(() => {
          this.$message.error('Some error!')
        })
      })
    },
    // 非定时任务启动
    start(item) {
      this.$confirm('start the task?', 'NOTICE', {
        confirmButtonText: 'Yes',
        cancelButtonText: 'Cancel',
        type: 'warning'
      }).then(() => {
        this.load_data = true
        startTask(item.id).then(() => {
          this.getTaskList()
          this.$message.success('Success!')
        }).catch(() => {
          this.$message.error('Some error!')
        })
      })
    },
    // 定时任务重启
    restart(item) {
      this.$confirm('restart timed task?', 'NOTICE', {
        confirmButtonText: 'Yes',
        cancelButtonText: 'Cancel',
        type: 'warning'
      }).then(() => {
        this.load_data = true
        restartTask(item.id).then(() => {
          this.getTaskList()
          this.$message.success('Success!')
        }).catch(() => {
          this.$message.error('Some error!')
        })
      })
    }
  }
}
</script>

<style lang="scss" scoped>
  .task_list {
    margin: 20px;

    .con-title {
      color: #333333;
      font-size: 16px;
      font-weight: bold;
      padding: 20px 10px;
      background: #f5f7fa;
    }
    .reject {
      background: #F04134;
    }
  }

  .pagination-block {
    padding: 20px 0;
    text-align: center;
  }

</style>

