<template>
  <div class="app-container">
    <el-card shadow="hover" class="box-card">
      <div class="filter-container">
        <!-- 过滤条件 -->
        <el-form
          ref="formData"
          :inline="true"
          :model="formData"
          size="small"
          class="demo-form-inline"
        >
          <el-form-item label="推荐人" prop="name">
            <el-input v-model="formData.name" placeholder />
          </el-form-item>
          <el-form-item label="分销商" prop="name">
            <el-input v-model="formData.name" placeholder />
          </el-form-item>
          <el-form-item label="手机" prop="name">
            <el-input v-model="formData.name" placeholder />
          </el-form-item>
          <el-form-item label="分销商等级" prop="name">
            <el-input v-model="formData.name" placeholder />
          </el-form-item>
          <el-form-item>
            <el-button :loading="loading" type="primary" @click="onSubmit">搜索</el-button>
          </el-form-item>
        </el-form>
      </div>
      <!-- 过滤条件 -->
      <!--列表-->
      <el-table v-loading="loading" :data="items.list" border style="width: 100%">
        <el-table-column label="ID" prop="id" width="80px" />
        <el-table-column label="推荐人" prop="name" width="100px" />
        <el-table-column label="分销商" prop="comic_num" width="80px" />
        <el-table-column label="姓名/手机" prop="sort" width="123px" />
        <el-table-column label="分销等级/下级人数" prop="sort" width="123px" />
        <el-table-column label="可用佣金/已提现佣金" prop="sort" width="123px" />
        <el-table-column label="状态" prop="sort" width="123px" />
        <el-table-column label="操作" prop="operation" fixed="right">
          <template slot-scope="{row}">
            <el-button icon="el-icon-edit" type="primary" size="mini" @click="handleUpdate(row)">编辑</el-button>
            <el-button
              icon="el-icon-delete"
              size="mini"
              type="danger"
              @click="handleDelete(row)"
            >删除</el-button>
          </template>
        </el-table-column>
      </el-table>
      <Pagination
        v-show="items.total_items>10"
        :total="items.total_items"
        :page.sync="formData.page"
        :limit.sync="formData.page_size"
        @pagination="getDistributorList"
      />
      <!--列表 -->
    </el-card>
    <el-dialog :title="textMap[dialogStatus]" :visible.sync="dialogFormVisible">
      <el-form
        ref="dataForm"
        :rules="rules"
        :model="temp"
        label-position="left"
        label-width="100px"
        style="width: 300px; margin-left:50px;"
      >
        <el-form-item label="分类名" prop="name">
          <el-input v-model="temp.name" />
        </el-form-item>
        <el-form-item label="分类展示排序" prop="sort">
          <el-input-number v-model="temp.sort" :step="1" step-strictly></el-input-number>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogFormVisible = false">取消</el-button>
        <el-button type="primary" @click="dialogStatus==='create'?createData():updateData()">确认</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import {
  getDistributor,
  distributorList,
  updateDistributor,
  deleteDistributor
} from "@/api/lime-admin/distributor";
import Pagination from "@/components/Pagination";

export default {
  name: "DistributorList",
  components: { Pagination },
  data() {
    return {
      formData: {
        name: "",
        channel_id: 0,
        novel_num: 0,
        sort: 0,
        page: 1,
        page_size: 10
      },
      temp: {
        id: undefined,
        name: "",
        channel_id: 1,
        novel_num: 0,
        sort: 0
      },
      loading: false,
      items: {
        cur_page: 1,
        total_items: 2,
        list: []
      },
      channels: ["全部", "男生", "女生"],
      dialogFormVisible: false,
      dialogStatus: "",
      textMap: {
        update: "编辑",
        create: "新增"
      },
      rules: {
        type: [
          { required: true, message: "type is required", trigger: "change" }
        ],
        timestamp: [
          {
            type: "date",
            required: true,
            message: "timestamp is required",
            trigger: "change"
          }
        ],
        title: [
          { required: true, message: "title is required", trigger: "blur" }
        ]
      }
    };
  },
  computed: {},
  created() {},
  mounted() {
    this.getDistributorList();
  },
  methods: {
    onSubmit() {
      // 查询按钮
      this.formData.page = 1;
      this.getDistributorList();
    },
    on_refresh() {
      this.getDistributorList();
    },
    timeChange(val) {
      // 时间选择
      if (val === null) {
        this.formData.start_time = "";
        this.formData.end_time = "";
      } else {
        this.formData.start_time = val[0];
        this.formData.end_time = val[1];
      }
    },
    resetTemp() {
      this.temp = {
        id: undefined,
        channel_id: 0,
        sort: 0,
        name: ""
      };
    },
    resetForm() {
      // 重置
      this.$refs["formData"].resetFields();
      this.formData.channel_id = 0;
      this.formData.name = "";
      this.formData.sort = 0;
    },
    handleClick(row) {
      this.$store.dispatch('tagsView/delView', this.$route)
      this.$router.push({ path: "/novel/books/?cid" + row.id });
    },
    handleDelete(row) {
      this.$confirm('此操作将永久删除数据, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        updateDistributor(row.id, { id: row.id }).then(() => {
          this.$notify({
            title: '成功',
            message: '删除成功',
            type: 'success',
            duration: 2000
          })
          this.getDistributorList();
        }).catch((res) => {
          this.$message.error(res.msg)
        })
      }).catch(() => {})
    },
    currentChange(index) {
      // 分页
      this.formData.page = index;
      this.getDistributorList();
    },
    async getDistributorList() {
      // 获取列表
      this.loading = true;
      try {
        const list = await distributorList(this.formData);
        this.items.list = list.data.list;
        this.items.total_items = list.data.total;
      } finally {
        this.loading = false;
      }
    }
  }
};
</script>

<style lang="scss" scoped>
.category_list {
  margin: 20px;
}
.pagination-block {
  padding: 20px 0;
  text-align: center;
}
</style>

