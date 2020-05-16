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
          <el-form-item label="是否收费" prop="is_vip">
            <el-select v-model="formData.is_vip" placeholder="请选择" style="width: 100px">
              <el-option
                v-for="(item,index) in is_vips"
                :key="index"
                :label="item"
                :value="index"
              />
            </el-select>
          </el-form-item>
          <el-form-item label="章节序号" prop="id_sort">
            <el-select v-model="formData.id_sort" placeholder="请选择" style="width: 100px">
              <el-option
                v-for="(item,index) in id_sort"
                :key="index"
                :label="item"
                :value="index"
              />
            </el-select>
          </el-form-item>
          <el-form-item label="标题" prop="title">
            <el-input v-model="formData.title" placeholder />
          </el-form-item>
          <el-form-item>
            <el-button :loading="loading" type="primary" @click="onSubmit">搜索</el-button>
            <el-form-item style="text-align: right;width: 50%;">
              <el-button type="primary" @click="onCreateChapter">新增章节</el-button>
          </el-form-item>
          </el-form-item>
        </el-form>
      </div>
      <!-- 过滤条件 -->
      <!--列表-->
      <el-table v-loading="loading" :data="items.list" border style="width: 100%">
        <el-table-column label="章节序号" prop="id" width="80px" />
        <el-table-column label="标题" prop="title" width="100px" />
        <el-table-column label="是否收费" prop="is_vip" width="80px" />
        <el-table-column label="字数" prop="text_num" width="100px" />
        <el-table-column label="创建时间" prop="created_at" width="123px" />
        <el-table-column label="更新时间" prop="updated_at" width="123px" />
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
              @click="handleModifyStatus(row,'published')"
            >分类小说</el-button>
          </template>
        </el-table-column>
      </el-table>
      <Pagination
        v-show="items.total_items>10"
        :total="items.total_items"
        :page.sync="formData.page"
        :limit.sync="formData.page_size"
        @pagination="getChaptersList"
      />
      <!--列表 -->
    </el-card>
  </div>
</template>

<script>
import { categoryList } from "@/api/lime-admin/category";
import {
  ChapterList,
  createChapter,
  updateChapter,
  deleteChapter
} from "@/api/lime-admin/chapter";
import Pagination from "@/components/Pagination";

export default {
  name: "ChaptersList",
  components: { Pagination },
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
        title: "",
        is_vip: 0,
        text_num: 0,
        page: 1,
        page_size: 10
      },
      temp: {
        id: undefined,
        title: "",
        is_vip: 0,
        text_num: 0
      },
      loading: false,
      items: {
        cur_page: 1,
        total_items: 2,
        list: []
      },
      is_vips: ["全部", "是", "否"],
      id_sort: ["升序", "降序"],
      dialogFormVisible: false,
      dialogStatus: "",
    };
  },
  computed: {},
  created() {},
  mounted() {
    this.getChaptersList();
  },
  methods: {
    onSubmit() {
      // 查询按钮
      this.formData.page = 1;
      this.getChaptersList();
    },
    on_refresh() {
      this.getChaptersList();
    },
    onCreateChapter() {
      var book_id = this.$route.query.book_id
      this.$router.push({ path: "/novel/chapters/create?book_id=" + book_id });
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
      // 详情
      console.log(row.id);
      this.$router.push({
        path: "/taskmanage/details/" + row.id
        // query: { id: row.id }
      });
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
            this.getChaptersList();
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
      this.$router.push({ path: "/novel/chapters/update?book_id=" + row.novel_id +"&chapter_id="+ row.id });
    },
    updateData() {
      this.$refs["dataForm"].validate(valid => {
        if (valid) {
          const tempData = Object.assign({}, this.temp);
          tempData.timestamp = +new Date(tempData.timestamp); // change Thu Nov 30 2017 16:41:05 GMT+0800 (CST) to 1512031311464
          updateCategory(this.temp.id, tempData).then(() => {
            this.getChaptersList();
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
      this.$confirm("此操作将永久删除数据, 是否继续?", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning"
      })
        .then(() => {
          deleteChapter(row.id, { id: row.id })
            .then(() => {
              this.$notify({
                title: "成功",
                message: "删除成功",
                type: "success",
                duration: 2000
              });
              this.getChaptersList();
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
      this.getChaptersList();
    },
    async getChaptersList() {
      // 获取列表
      this.loading = true;
      try {
        this.listQuery.skip = (this.listQuery.page - 1) * this.listQuery.limit
        this.listQuery.q = 'novel_id=' + this.$route.query.book_id
        const list = await ChapterList(this.listQuery);
        this.items.list = list.data.result;
        for (const v of this.items.list) {
          v.created_at = this.$moment(v.created_at).format('YYYY-MM-DD HH:mm:ss');
          v.updated_at = this.$moment(v.updated_at).format('YYYY-MM-DD HH:mm:ss');
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

