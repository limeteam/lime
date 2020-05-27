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
          <el-form-item label="书名" prop="novel_id">
            <el-input v-model="formData.novel_id" placeholder />
          </el-form-item>
          <el-form-item label="来源" prop="source">
            <el-select v-model="formData.source" placeholder="状态" style="width: 100px">
              <el-option
                v-for="(item,index) in sources"
                :key="index"
                :label="item"
                :value="index"
              />
            </el-select>
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
        <el-table-column label="用户" prop="username" width="100px" />
        <el-table-column label="作品" prop="novel_id" width="80px" />
        <el-table-column label="内容" prop="content" width="123px" />
        <el-table-column label="点赞数" prop="likes" width="80px" />
        <el-table-column label="来源" prop="source" width="123px" />
        <el-table-column label="评论时间" prop="created_at" width="123px" />
        <el-table-column label="操作" prop="operation" fixed="right">
          <template slot-scope="{row}">
            <el-button
              icon="el-icon-delete"
              size="mini"
              type="danger"
              @click="handleDelete(row)"
            >删除</el-button>
            <el-button
              icon="el-icon-view"
              size="mini"
              type="success"
              @click="handleClick(row)"
            >封禁此用户</el-button>
          </template>
        </el-table-column>
      </el-table>
      <Pagination
        v-show="items.total_items>10"
        :total="items.total_items"
        :page.sync="formData.page"
        :limit.sync="formData.page_size"
        @pagination="getCommentsList"
      />
      <!--列表 -->
    </el-card>
  </div>
</template>

<script>
import {
  CommentList,
  deleteComment
} from "@/api/lime-admin/comment";
import Pagination from "@/components/Pagination";
export default {
  name: "CommentsList",
  components: { Pagination },
  data() {
    return {
      formData: {
        id: 0,
        novel_id: 0,
        username: "",
        content: "",
        likes: 0,
        source: "",
        page: 1,
        page_size: 10
      },
      temp: {
        id: undefined,
        novel_id: 0,
        username: "",
        content: "",
        likes: 0,
        source: ""
      },
      loading: false,
      items: {
        cur_page: 1,
        total_items: 2,
        list: []
      },
      dialogStatus: "",
      sources: {
          0: "全部",
          1: "PC",
          2: "Wap",
          3: "Android",
          4: "IOS",
          5: "公众号",
          6: "小程序",
      },
    };
  },
  computed: {},
  created() {},
  mounted() {
    this.getCommentsList();
  },
  methods: {
    onSubmit() {
      // 查询按钮
      this.formData.page = 1;
      this.getCommentsList();
    },
    on_refresh() {
      this.getCommentsList();
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
    },
    handleClick(row) {
      this.$store.dispatch('tagsView/delView', this.$route)
      this.$router.push({ path: "/novel/books/?cid" + row.id });
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
          this.getCommentsList();
        }).catch((res) => {
          this.$message.error(res.msg)
        })
      }).catch(() => {})
    },
    currentChange(index) {
      // 分页
      this.formData.page = index;
      this.getCommentsList();
    },
    async getCommentsList() {
      // 获取列表
      this.loading = true;
      try {
        const list = await CommentList(this.formData);
        this.items.list = list.data.list;
        for (const v of this.items.list) {
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

