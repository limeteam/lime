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
            <el-input v-model="formData.name" placeholder="书名/作者/来源" class="filter-item" />
          </el-form-item>
          <el-form-item>
            <el-button
              :loading="loading"
              type="primary"
              icon="el-icon-search"
              @click="onSubmit"
              class="filter-item"
            >搜索</el-button>
          </el-form-item>
          <el-form-item style="text-align: right;width: 60%;">
            <el-button type="primary" @click="onCreateNovel">新增小说</el-button>
          </el-form-item>
          <selectedpanel />
        </el-form>
      </div>
      <!-- 过滤条件 -->
      <!--列表-->
      <el-table v-loading="loading" :data="items.list" border style="width: 100%">
        <el-table-column prop="cover" label="封面" width="100px" align="center">
          <template slot-scope="scope">
            <img
              :src="base_url + scope.row.cover"
              style="width: 75px;height:100px"
            />
          </template>
        </el-table-column>
        <el-table-column prop="name" label="类型/名称" width="220px" align="left">
          <template slot-scope="scope">
            <span class="basic-info">
              <a v-bind:href="scope.row.url">[{{ scope.row.category_name }}] {{scope.row.name}}</a>&nbsp;
              <small>[{{ scope.row.status}}]</small>
            </span>
            <br />
            <small
              class="author"
            >作&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;者：{{scope.row.author}}</small>
            <br />
            <small class="last-chapter">最新章节：第二章 星星之火</small>
            <br />
            <small class="update-time">更新时间：{{ scope.row.chapter_updated_at}}</small>
          </template>
        </el-table-column>
        <el-table-column label="所属频道" prop="channel_id" width="80px" />
        <el-table-column label="来源" prop="source" width="100px" />
        <el-table-column label="千字价格" prop="thousand_characters_price" width="80px" />
        <el-table-column label="每章价格" prop="chapter_price" width="80px" />
        <el-table-column label="操作" prop="operation" fixed="right">
          <template slot-scope="{row}">
            <el-button
              v-if="row.status!='published'"
              icon="el-icon-notebook-2"
              size="mini"
              type="success"
              @click="handleJumpChapater(row.id)"
            >章节管理</el-button>
            <el-button icon="el-icon-edit" type="primary" size="mini" @click="handleUpdate(row)">编辑</el-button>
            <el-button
              v-if="row.status ==0"
              icon="el-icon-close"
              type="primary"
              size="mini"
              @click="handleModifyStatus(row,1)"
            >下架</el-button>
            <el-button
              v-if="row.status==1"
              icon="el-icon-close"
              type="primary"
              size="mini"
              @click="handleModifyStatus(row,0)"
            >上架</el-button>
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
import { categoryList } from "@/api/lime-admin/category";
import {
  BookList,
  createBook,
  updateBook,
  deleteBook,
  updatestatus
} from "@/api/lime-admin/book";
import Pagination from "@/components/Pagination";
import { CATEGORY_CHANNEL } from "./emun/index.js";
import selectedpanel from "./components/selectedpanel";
export default {
  name: "BooksList",
  filters: {
    statuss: function(value) {
      Number(value);
      return CATEGORY_CHANNEL[value];
    }
  },
  components: { Pagination, selectedpanel },
  data() {
    return {
      listQuery: {
        page: 1,
        skip: 0,
        limit: 20,
        importance: undefined,
        title: undefined,
        type: undefined,
        sort: '+id'
      },
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
        page_size: 10,
        url: ""
      },
      base_url: process.env.VUE_APP_CONFIG_API,
      loading: false,
      items: {
        cur_page: 1,
        total_items: 2,
        list: []
      },
      statusMap: {
        0 : "连载中",
        1 : "已完结",
        2 : "太监"
      }
    };
  },
  computed: {},
  created() {},
  watch: {
    $route() {
      this.getBookList();
    }
  },
  mounted() {
    this.getBookList();
  },
  methods: {
    onSubmit() {
      // 查询按钮
      this.formData.page = 1;
      this.getBookList();
    },
    onCreateNovel() {
      this.$router.push({ path: "/novel/create"});
    },
    on_refresh() {
      this.getBookList();
    },
    handleJumpChapater(id){
      this.$router.push({ path: "/novel/chapters?book_id=" + id });
    },
    handleUpdate(row) {
      this.$router.push({ path: "/novel/update?id=" + row.id });
    },
    handleModifyStatus(row, status) {
      updatestatus(row.id, { status: status })
            .then(() => {
              this.$notify({
                title: "成功",
                message: "操作成功",
                type: "success",
                duration: 2000
              });
              this.getBookList();
            })
            .catch(res => {
              this.$message.error(res.msg);
            });
    },
    handleDelete(row) {
      this.$confirm("此操作将永久删除数据, 是否继续?", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning"
      })
        .then(() => {
          deleteBook(row.id, { id: row.id })
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
        var qStr = "";
        if (this.$route.query.channel_id > 0){
          qStr = 'channel_id=' + this.$route.query.channel_id + ','
        }
        if (this.$route.query.cid > 0){
           qStr += 'cid=' + this.$route.query.cid + ','
        }
        if (this.$route.query.status > 0){
          qStr += 'status=' + this.$route.query.status + ','
        }
        if (this.$route.query.online_status > 0){
          qStr += 'online_status=' + this.$route.query.online_status + ','
        }
        if (this.$route.query.is_sensitivity > 0){
          qStr += 'is_sensitivity=' + this.$route.query.is_sensitivity + ','
        }
        if (this.$route.query.flag > 0){
          qStr += 'flag=' + this.$route.query.flag + ','
        }
        if (this.$route.query.word > 0){
          qStr += 'word=' + this.$route.query.word + ','
        }
        if (this.$route.query.sort > 0){
          qStr += 'sort=' + this.$route.query.sort + ','
        }
        if (this.formData.name != null) {
          qStr += 'name=' + this.formData.name + ','
        }
        this.listQuery.skip = (this.listQuery.page - 1) * this.listQuery.limit
        this.listQuery.q = qStr.slice(0,-1)
        const categorys = await categoryList();
        var categories = new Object()
        
        for (var i = 0; i < categorys.data.list.length; i++) {
          var id = categorys.data.list[i].id
          categories[id] = categorys.data.list[i].name
        }
        const list = await BookList(this.listQuery);
        this.items.list = list.data.result;
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
          v.category_name = categories[v.category_id];
          v.status = this.statusMap[v.status];
          v.chapter_updated_at = this.$moment(v.chapter_updated_at).format('YYYY-MM-DD HH:mm:ss');
          v.url = "/#/novel/view?id="+v.id
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
.pagination-block {
  padding: 20px 0;
  text-align: center;
}
</style>

