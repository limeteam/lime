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
          <el-form-item label="分类名" prop="name">
            <el-input v-model="formData.name" placeholder />
          </el-form-item>
          <el-form-item>
            <el-button :loading="loading" type="primary" @click="onSubmit">搜索</el-button>
            <el-button
              style="margin-left: 400px;"
              type="primary"
              icon="el-icon-edit"
              @click="handleCreate"
            >添加等级</el-button>
          </el-form-item>
        </el-form>
      </div>
      <!-- 过滤条件 -->
      <!--列表-->
      <el-table v-loading="loading" :data="items.list" border style="width: 100%">
        <el-table-column label="ID" prop="id" width="80px" />
        <el-table-column label="等级名称" prop="name" width="100px" />
        <el-table-column label="升级条件" prop="upgrade_conditions" width="80px" />
        <el-table-column label="降级条件" prop="degradation_conditions" width="123px" />
        <el-table-column label="权重" prop="weight" width="123px" />
        <el-table-column label="操作" prop="operation" fixed="right">
          <template slot-scope="{row}">
            <el-button icon="el-icon-edit" type="primary" size="mini" @click="handleUpdate(row)">编辑</el-button>
            <el-button icon="el-icon-delete" size="mini" type="danger" @click="handleDelete(row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
      <Pagination
        v-show="items.total_items>10"
        :total="items.total_items"
        :page.sync="formData.page"
        :limit.sync="formData.page_size"
        @pagination="getLevelList"
      />
      <!--列表 -->
    </el-card>
  </div>
</template>

<script>
import { DistributorLevelList } from "@/api/lime-admin/distributorLevel";
import Pagination from "@/components/Pagination";

export default {
  name: "distributorLevelList",
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
      loading: false,
      items: {
        cur_page: 1,
        total_items: 2,
        list: []
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
    this.getLevelList();
  },
  methods: {
    onSubmit() {
      // 查询按钮
      this.formData.page = 1;
      this.getLevelList();
    },
    handleCreate() {
      this.$router.push({ path: "/distributor/distributorLevelCreate" });
    },
    handleUpdate(row){
      this.$router.push({ path: "/distributor/distributorLevelUpdate?id=" + row.id });
    },
    handleDelete(row) {
      this.$confirm("此操作将永久删除数据, 是否继续?", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning"
      })
        .then(() => {
          deleteCategory(row.id, { id: row.id })
            .then(() => {
              this.$notify({
                title: "成功",
                message: "删除成功",
                type: "success",
                duration: 2000
              });
              this.getLevelList();
            })
            .catch(res => {
              this.$message.error(res.msg);
            });
        })
        .catch(() => {});
    },
    currentChange(index) {
      // 分页
      this.formData.page = index;
      this.getLevelList();
    },
    async getLevelList() {
      // 获取列表
      this.loading = true;
      try {
        const list = await DistributorLevelList(this.formData);
        this.items.list = list.data.result;
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

