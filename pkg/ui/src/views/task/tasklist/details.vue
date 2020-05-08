<template>
  <div class="app-container">
    <el-row v-loading="loadData" border fit highlight-current-row>
      <el-col :span="12">
        <el-form ref="form" :model="form" :rules="rules" label-position="right" label-width="150px">
          <el-form-item :label="$t('task.taskInfo.name')" prop="task_name">
            <el-input v-model="form.task_name" placeholder="please enter the content" :disabled="is_disable" />
          </el-form-item>
          <el-form-item :label="$t('task.taskInfo.rule')" prop="task_rule_name">
            <el-select v-model="form.task_rule_name" placeholder="please select" :disabled="is_disable">
              <el-option
                v-for="item in ruleOpts"
                :key="item"
                :label="item"
                :value="item"
              />
            </el-select>
          </el-form-item>
          <el-form-item :label="$t('task.taskInfo.desc')">
            <el-input v-model="form.task_desc" type="textarea" placeholder="please enter the content" :rows="2" />
          </el-form-item>
          <el-form-item :label="$t('task.taskInfo.cron')">
            <el-input v-model="form.cron_spec" style="width:calc(100% - 150px);padding-right:10px" placeholder="compatible with crontab" />
            <el-button style="width:100px" type="primary" @click="showDialog">
              生成 cron
            </el-button>
          </el-form-item>
          <el-dialog title="生成 cron" :visible.sync="showCron">
            <vcrontab :expression="expression" @hide="showCron=false" @fill="crontabFill" />
          </el-dialog>
          <el-form-item :label="$t('task.taskInfo.proxy')">
            <el-input v-model="form.proxy_urls" placeholder="compatible with socks5,http,https; separated by commas" />
          </el-form-item>
          <el-form-item :label="$t('task.taskInfo.agent')">
            <el-input v-model="form.opt_user_agent" placeholder="user agent" />
          </el-form-item>
          <el-form-item :label="$t('task.taskInfo.maxDepth')">
            <el-input-number v-model="form.opt_max_depth" :controls="false" />
          </el-form-item>
          <el-form-item :label="$t('task.taskInfo.allowDomains')">
            <el-input v-model="form.opt_allowed_domains" placeholder="default empty, not limited" />
          </el-form-item>
          <el-form-item :label="$t('task.taskInfo.urlFilter')">
            <el-input v-model="form.opt_url_filters" placeholder="default empty, not limited, support regex" />
          </el-form-item>
          <el-form-item :label="$t('task.taskInfo.maxBody')">
            <el-input-number v-model="form.opt_max_body_size" :controls="false" class="fl" />
          </el-form-item>
          <el-form-item :label="$t('task.taskInfo.requestTimeout')">
            <el-input-number v-model="form.opt_request_timeout" :controls="false" class="fl" />
          </el-form-item>
          <el-form-item :label="$t('task.taskInfo.limitEn')">
            <el-checkbox v-model="form.limit_enable" />
          </el-form-item>
          <el-form-item :label="$t('task.taskInfo.limitDomainGlob')">
            <el-input v-model="form.limit_domain_glob" placeholder="default *" />
          </el-form-item>
          <el-form-item :label="$t('task.taskInfo.limitDelay')">
            <el-input-number v-model="form.limit_delay" :controls="false" />
          </el-form-item>
          <el-form-item :label="$t('task.taskInfo.limitRandomDelay')">
            <el-input-number v-model="form.limit_random_delay" :controls="false" />
          </el-form-item>
          <el-form-item :label="$t('task.taskInfo.limitPara')">
            <el-input-number v-model="form.limit_parallelism" :controls="false" />
          </el-form-item>

          <el-form-item>
            <el-button type="primary" :loading="on_submit_loading" :disabled="submit_disable" @click="on_submit_form">
              {{ routeID ? $t('task.taskInfo.update') : $t('task.taskInfo.add') }}
            </el-button>
          </el-form-item>
        </el-form>
      </el-col>
    </el-row>
  </div>
</template>

<script>
import { taskInfo, getRules, updateTask, saveTask } from '@/api/lime-admin/task'
import waves from '@/directive/waves'
import vcrontab from 'vcrontab'
export default {
  name: 'TaskCreate',
  directives: {
    waves
  },
  components: { vcrontab },
  data() {
    return {
      expression: '',
      showCron: false,
      is_disable: this.$route.params.id >= 1,
      form: {
        cron_spec: '',
        opt_user_agent: navigator.userAgent,
        limit_enable: true,
        auto_migrate: true,
        limit_parallelism: 1,
        opt_request_timeout: 10
      },
      showExportDB: false,
      ruleOpts: [],
      exportDBList: [],
      routeID: this.$route.params.id,
      loadData: false,
      on_submit_loading: false,
      submit_disable: false,
      rules: {
        task_name: [{ required: true, message: 'task name should not be empty', trigger: 'blur' }],
        task_rule_name: [{ required: true, message: 'please select rule name', trigger: 'change' }]
      }
    }
  },
  created() {
    this.getRules()
    this.routeID && this.getTaskInfo()
  },
  methods: {
    // 获取数据
    getTaskInfo() {
      this.loadData = true
      taskInfo(this.routeID).then(response => {
        const data = response.data.result
        this.form = data
        this.loadData = false
        this.outputTypeChange(data.output_type)
        setTimeout(() => {
          this.loadData = false
        }, 1.5 * 1000)
      })
    },
    crontabFill(value) {
      value = value.replace(/[?]/g, '*')
      this.form.cron_spec = value
    },
    showDialog() {
      this.expression = this.form.cron_spec
      this.showCron = true
    },
    // 获取所有rules
    getRules() {
      getRules().then(response => {
        const data = response.data.list
        this.ruleOpts = data
        setTimeout(() => {
        }, 1.5 * 1000)
      })
    },
    // 提交
    on_submit_form() {
      this.$refs.form.validate((valid) => {
        if (!valid) return false
        this.on_submit_loading = true
        if (this.routeID) {
          updateTask(this.routeID, this.form).then((response) => {
            const data = response.data
            this.$message.success('task update success!  taskID:' + data.id + '  2s redirect to task list page!')
            this.on_submit_loading = false
            this.submit_disable = true
            setTimeout(() => this.$router.push({ name: 'taskList' }), 2000)
          }).catch(() => {
            this.on_submit_loading = false
          })
        } else {
          saveTask(this.form).then((response) => {
            const data = response.data
            this.$message.success('task create success!  taskID:' + data.id + '  2s redirect to task list page!')
            this.on_submit_loading = false
            this.submit_disable = true
            setTimeout(() => this.$router.push({ name: 'taskList' }), 3000)
          }).catch(() => {
            this.on_submit_loading = false
          })
        }
      })
    }
  }
}
</script>
