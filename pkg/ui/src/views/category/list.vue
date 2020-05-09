<template>
  <div class="category_list">
    <el-card shadow="hover" class="box-card">
      <!-- 过滤条件 -->
      <el-form ref="formData" :inline="true" :model="formData" size="small" class="demo-form-inline">
        <el-form-item label="所属频道" prop="channel_id">
          <el-select v-model="formData.channel_id" placeholder="状态" style="width: 100px">
            <el-option v-for="(item,index) in channels" :key="index" :label="item" :value="index" />
          </el-select>
        </el-form-item>
        <el-form-item label="分类名" prop="name">
          <el-input v-model="formData.name" placeholder="" />
        </el-form-item>
        <el-form-item>
          <el-button :loading="loading" type="primary" @click="onSubmit">
            搜索
          </el-button>
        </el-form-item>
      </el-form>
      <!-- 过滤条件 -->
      <!--列表-->
      <el-table v-loading="loading" :data="items.list" border style="width: 100%">
        <el-table-column label="ID" prop="id" />
        <el-table-column label="分类名" prop="name" />
        <el-table-column label="小说数" prop="novel_num" />
        <el-table-column label="所属频道" prop="channel_id" />
        <el-table-column label="分类展示排序" prop="sort" />
        <el-table-column label="操作" fixed="right">
          <template slot-scope="props">
            <el-button v-if="props.row.optionbutton & 0b00100" type="danger" size="small" @click="stop(props.row)">
              编辑
            </el-button>
          </template>
        </el-table-column>
      </el-table>
      <Pagination v-show="items.total_items>10" :total="items.total_items" :page.sync="formData.page" :limit.sync="formData.page_size" @pagination="getCategoryList" />
      <!--列表 -->
    </el-card>
  </div>
</template>

<script>
import { categoryList } from '@/api/lime-admin/category'
import Pagination from '@/components/Pagination'
import { CATEGORY_CHANNEL } from './emun/index.js'

export default {
  name: 'CategoryList',
  filters: {
    statuss: function(value) {
      Number(value)
      return CATEGORY_CHANNEL[value]
    }
  },
  components: { Pagination },
  data() {
    return {
      formData: {
        name: '',
        channel_id: 0,
        novel_num: 0,
        sort: '',
        page: 1,
        page_size: 10
      },
      loading: false,
      items: {
        cur_page: 1,
        total_items: 2,
        list: []
      },
      channels: [
        '全部',
        '男生',
        '女生'
      ],
      SETime: []

    }
  },
  computed: {
  },
  created() {
  },
  mounted() {
    this.getCategoryList()
  },
  methods: {
    onSubmit() {
      // 查询按钮
      this.formData.page = 1
      this.getCategoryList()
    },
    on_refresh() {
      this.getCategoryList()
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
      this.getCategoryList()
    },
    async getCategoryList() {
      // 获取列表
      this.loading = true
      try {
        const list = await categoryList(this.formData)
        this.items.list = list.data.list
        for (const v of this.items.list) {
          v.isCron = 'No'
          if (v.cron_spec) {
            v.isCron = 'Yes'
          }
        }
        this.items.total_items = list.data.total
      } finally {
        this.loading = false
      }
    },
  }
}
</script>

<style lang="scss" scoped>
  .category_list {
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

