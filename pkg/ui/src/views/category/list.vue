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
          <el-form-item label="所属频道" prop="channel_id">
            <el-select v-model="formData.channel_id" placeholder="状态" style="width: 100px">
              <el-option
                v-for="(item,index) in channels"
                :key="index"
                :label="item"
                :value="index"
              />
            </el-select>
          </el-form-item>
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
            >新增分类</el-button>
          </el-form-item>
        </el-form>
      </div>
      <!-- 过滤条件 -->
      <!--列表-->
      <el-table v-loading="loading" :data="items.list" border style="width: 100%">
        <el-table-column label="ID" prop="id" width="80px" />
        <el-table-column label="分类名" prop="name" width="100px" />
        <el-table-column label="小说数" prop="novel_num" width="80px" />
        <el-table-column label="所属频道" prop="channel_id" width="100px" />
        <el-table-column label="分类展示排序" prop="sort" width="123px" />
        <el-table-column label="操作" prop="operation" fixed="right">
          <template slot-scope="{row}">
            <el-button icon="el-icon-edit" type="primary" size="mini" @click="handleUpdate(row)">编辑</el-button>
            <el-button
              v-if="row.status!='deleted'"
              icon="el-icon-delete"
              size="mini"
              type="danger"
              @click="handleDelete(row)"
            >删除</el-button>
            <el-button
              v-if="row.status!='published'"
              icon="el-icon-view"
              size="mini"
              type="success"
              @click="handleClick(row)"
            >分类小说</el-button>
          </template>
        </el-table-column>
      </el-table>
      <Pagination
        v-show="items.total_items>10"
        :total="items.total_items"
        :page.sync="formData.page"
        :limit.sync="formData.page_size"
        @pagination="getCategoryList"
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
        <el-form-item label="所属频道" prop="channel_id">
          <el-select v-model="temp.channel_id" class="filter-item" placeholder="请选择">
            <el-option
              v-for="item in category_channels"
              :key="item.key"
              :label="item.display_name"
              :value="item.key"
            />
          </el-select>
        </el-form-item>
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
  categoryList,
  createCategory,
  updateCategory,
  deleteCategory
} from "@/api/lime-admin/category";
import Pagination from "@/components/Pagination";
import { CATEGORY_CHANNEL } from "./emun/index.js";

export default {
  name: "CategoryList",
  filters: {
    statuss: function(value) {
      Number(value);
      return CATEGORY_CHANNEL[value];
    }
  },
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
      category_channels: CATEGORY_CHANNEL,
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
    this.getCategoryList();
  },
  methods: {
    onSubmit() {
      // 查询按钮
      this.formData.page = 1;
      this.getCategoryList();
    },
    on_refresh() {
      this.getCategoryList();
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
    handleCreate() {
      this.resetTemp();
      this.dialogStatus = "create";
      this.dialogFormVisible = true;
      this.$nextTick(() => {
        this.$refs["dataForm"].clearValidate();
      });
    },
    createData() {
      this.$refs["dataForm"].validate(valid => {
        if (valid) {
          this.temp.id = parseInt(Math.random() * 100) + 1024; // mock a id
          createCategory(this.temp).then(() => {
            this.getCategoryList();
            this.dialogFormVisible = false;
            this.$notify({
              title: "成功",
              message: "新增成功",
              type: "success",
              duration: 2000
            });
          });
        }
      });
    },
    handleUpdate(row) {
      this.temp = Object.assign({}, row); // copy obj
      this.temp.timestamp = new Date(this.temp.timestamp);
      this.dialogStatus = "update";
      this.dialogFormVisible = true;
      this.$nextTick(() => {
        this.$refs["dataForm"].clearValidate();
      });
    },
    updateData() {
      this.$refs["dataForm"].validate(valid => {
        if (valid) {
          const tempData = Object.assign({}, this.temp);
          tempData.timestamp = +new Date(tempData.timestamp); // change Thu Nov 30 2017 16:41:05 GMT+0800 (CST) to 1512031311464
          updateCategory(this.temp.id, tempData).then(() => {
            this.getCategoryList();
            this.dialogFormVisible = false;
            this.$notify({
              title: "成功",
              message: "更新成功",
              type: "success",
              duration: 2000
            });
          });
        }
      });
    },
    handleDelete(row) {
      this.$confirm('此操作将永久删除数据, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        deleteCategory(row.id, { id: row.id }).then(() => {
          this.$notify({
            title: '成功',
            message: '删除成功',
            type: 'success',
            duration: 2000
          })
          this.getCategoryList();
        }).catch((res) => {
          this.$message.error(res.msg)
        })
      }).catch(() => {})
    },
    currentChange(index) {
      // 分页
      this.formData.page = index;
      this.getCategoryList();
    },
    async getCategoryList() {
      // 获取列表
      this.loading = true;
      try {
        const list = await categoryList(this.formData);
        this.items.list = list.data.list;
        for (const v of this.items.list) {
          switch (v.channel_id) {
            case 1:
              v.channel_id = "男生";
              break;
            case 2:
              v.channel_id = "女生";
              break;
            default:
              v.channel_id = "全部";
          }
        }
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

