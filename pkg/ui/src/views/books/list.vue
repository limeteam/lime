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
          <el-form-item label="搜索" prop="name">
            <el-input v-model="formData.name" placeholder="书名/作者/来源" />
          </el-form-item>
          <el-form-item>
            <el-button :loading="loading" type="primary" @click="onSubmit">搜索</el-button>
            <a href="/#/novel/create" class="el-button el-button--primary">新增小说</a>
          </el-form-item>
          <selectedpanel/>
        </el-form>
      </div>
      <!-- 过滤条件 -->
      <!--列表-->
      <el-table v-loading="loading" :data="items.list" border style="width: 100%">
        <el-table-column label="封面" prop="id" width="80px" />
        <el-table-column label="类型/名称" prop="name" width="100px" />
        <el-table-column label="所属频道" prop="novel_num" width="80px" />
        <el-table-column label="来源" prop="channel_id" width="100px" />
        <el-table-column label="千字价格" prop="sort" width="123px" />
        <el-table-column label="每章价格" prop="sort" width="80px" />
        <el-table-column label="操作" prop="operation" fixed="right">
          <template slot-scope="{row}">
            <el-button
              v-if="row.status!='published'"
              icon="el-icon-notebook-2"
              size="mini"
              type="success"
              @click="handleModifyStatus(row,'published')"
            >章节管理</el-button>
            <el-button icon="el-icon-edit" type="primary" size="mini" @click="handleUpdate(row)">编辑</el-button>
            <el-button icon="el-icon-close" type="primary" size="mini" @click="handleDown(row)">下架</el-button>
            <el-button
              v-if="row.status!='deleted'"
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
        @pagination="getBookList"
      />
      <!--列表 -->
    </el-card>
  </div>
</template>
<script>
import {
  categoryList
} from "@/api/lime-admin/category";
import {
  BookList,
  createBook,
  updateBook,
  deleteBook
} from "@/api/lime-admin/book";
import Pagination from "@/components/Pagination";
import { CATEGORY_CHANNEL } from "./emun/index.js";
import selectedpanel from './components/selectedpanel';
export default {
  name: "BooksList",
  filters: {
    statuss: function(value) {
      Number(value);
      return CATEGORY_CHANNEL[value];
    }
  },
  components: { Pagination,selectedpanel},
  data() {
    return {
      formData: {
        id: undefined,
        name: "",
        old_name: "",
        channel_id: 0,
        category_id: 0,
        desc: "",
        cover: "",
        author: "",
        source: "",
        split_rule: "",
        upload_file: "",
        status: 0,
        attribute: 0,
        chapter_price: 0,
        thousand_characters_price: 0,
        score: 0,
        text_num: 0,
        chapter_num: 0,
        chapter_id: 0,
        chapter_title: "",
        views: 0,
        collect_num: 0,
        online_status: 0,
        is_sensitivity: 0,
        page: 1,
        page_size: 10
      },
      loading: false,
      items: {
        cur_page: 1,
        total_items: 2,
        list: []
      },
    };
  },
  computed: {},
  created() {},
  mounted() {
    this.getBookList();
  },
  methods: {
    onSubmit() {
      // 查询按钮
      this.formData.page = 1;
      this.getBookList();
    },
    on_refresh() {
      this.getBookList();
    },
    handleClick(row) {
    },
    handleUpdate(row) {
      console.log(row)
      this.$router.push({path: '/novel/update?id=' + row.id})
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
              this.getBookList();
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
      this.getBookList();
    },
    async getBookList() {
      // 获取列表
      this.loading = true;
      try {
        const list = await BookList(this.formData);
        this.items.list = list.data.list;
        console.log(this.items.list);
        this.items.total_items = list.data.total;
      } finally {
        this.loading = false;
      }
    }
  }
};
</script>

<style lang="scss" scoped>
.pagination-block {
  padding: 20px 0;
  text-align: center;
}
</style>

