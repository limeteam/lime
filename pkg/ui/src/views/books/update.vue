<template>
  <el-card shadow="hover" class="box-card">
    <div>
      <el-row>
        <el-col :span="22">
          <span class="book_title"></span>
        </el-col>
        <el-col :span="2">
          <a href="/#/novel/books" class="el-button el-button--primary">小说列表</a>
        </el-col>
      </el-row>
      <el-divider>修改小说</el-divider>
    </div>
    <el-form
      ref="dataForm"
      :model="form"
      label-position="left"
      label-width="120px"
      style="width: 500px; margin-left:50px;"
    >
      <el-form-item label="所属频道" prop="channel_id">
        <el-select v-model="form.channel_id" class="filter-item" placeholder="请选择">
          <el-option
            v-for="item in category_channels"
            :key="item.key"
            :label="item.display_name"
            :value="item.key"
          />
        </el-select>
      </el-form-item>
      <el-form-item label="所属分类" prop="category_id">
        <el-select v-model="form.category_id" class="filter-item" placeholder="请选择">
          <el-option v-for="item in categorys" :key="item.id" :label="item.name" :value="item.id" />
        </el-select>
      </el-form-item>
      <el-form-item label="小说名" prop="name">
        <el-input v-model="form.name" />
      </el-form-item>
      <el-form-item label="原名" prop="old_name">
        <el-input v-model="form.old_name" />
      </el-form-item>
      <el-form-item label="来源" prop="source">
        <el-input v-model="form.source" />
      </el-form-item>
      <el-form-item label="作者" prop="author">
        <el-input v-model="form.author" />
      </el-form-item>
      <el-form-item label="简介" prop="desc">
        <el-input type="textarea" v-model="form.desc"></el-input>
      </el-form-item>
      <el-form-item label="千字价格" prop="thousand_characters_price">
        <el-input-number v-model="form.thousand_characters_price" :step="1" step-strictly></el-input-number>
        <div class="help-block">免费小说设置成0</div>
      </el-form-item>
      <el-form-item label="每章价格" prop="chapter_price">
        <el-input-number v-model="form.chapter_price" :step="1" step-strictly></el-input-number>
        <div class="help-block">免费小说设置成0</div>
      </el-form-item>
      <el-form-item label="评分" prop="score">
        <el-input-number v-model="form.score" :step="1" step-strictly></el-input-number>
        <div class="help-block">可设值0~100，前端展示1.0~10.0分（设置值/10）</div>
      </el-form-item>
      <el-form-item label="浏览数" prop="views">
        <el-input-number v-model="form.views" :step="1" step-strictly></el-input-number>
        <div class="help-block">浏览数会乘以一定系数在前端以热度值形式显示</div>
      </el-form-item>
      <el-form-item label="收藏数" prop="collect_num">
        <el-input-number v-model="form.collect_num" :step="1" step-strictly></el-input-number>
      </el-form-item>
      <el-form-item label="属性" prop="attribute">
        <el-checkbox-group v-model="form.attribute">
          <el-checkbox
            v-for="item in book_atts"
            :key="item.key"
            :label="item.display_name"
            :value="item.key"
          />
        </el-checkbox-group>
      </el-form-item>
      <el-form-item label="状态" prop="status">
        <el-select v-model="form.status" class="filter-item" placeholder="请选择">
          <el-option
            v-for="item in book_status"
            :key="item.key"
            :label="item.display_name"
            :value="item.key"
          />
        </el-select>
      </el-form-item>
      <el-form-item label="上架状态" prop="online_status">
        <el-select v-model="form.online_status" class="filter-item" placeholder="请选择">
          <el-option
            v-for="item in book_online_status"
            :key="item.key"
            :label="item.display_name"
            :value="item.key"
          />
        </el-select>
        <div class="help-block">下架状态小说将在各端不可见</div>
      </el-form-item>
      <el-form-item label="是否敏感作品" prop="is_sensitivity">
        <el-select v-model="form.is_sensitivity" class="filter-item" placeholder="请选择">
          <el-option
            v-for="item in book_is_sensitivity"
            :key="item.key"
            :label="item.display_name"
            :value="item.key"
          />
        </el-select>
        <div class="help-block">敏感小说不会出现在书库和搜索结果中，只能通过推荐位和推送等形式触达用户</div>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="updateData">修改</el-button>
        <el-button>取消</el-button>
      </el-form-item>
    </el-form>
  </el-card>
</template>
<script>
import { categoryList } from "@/api/lime-admin/category";
import { createBook, getBook, updateBook } from "@/api/lime-admin/book";
import {
  CATEGORY_CHANNEL,
  BOOK_ATTRS,
  BOOK_ONLINE_STATUS,
  BOOK_STATUS,
  BOOK_IS_SENSITIVITYS
} from "./emun/index.js";
export default {
  name: "UpdateBook",
  data() {
    return {
      form: {
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
        attribute: [],
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
        is_sensitivity: 0
      },
      category_channels: CATEGORY_CHANNEL,
      book_atts: BOOK_ATTRS,
      book_online_status: BOOK_ONLINE_STATUS,
      book_status: BOOK_STATUS,
      book_is_sensitivity: BOOK_IS_SENSITIVITYS,
      categorys: []
    };
  },
  mounted() {
    this.getCategorys();
    this.getNovels();
  },
  methods: {
    resetForm() {
      // 重置
      this.$refs["dataForm"].resetFields();
      this.formData.channel_id = 0;
      this.formData.name = "";
      this.formData.sort = 0;
    },
    updateData() {
      this.$refs["dataForm"].validate(valid => {
        if (valid) {
          updateBook(this.form.id, this.form).then(() => {
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
    async getCategorys() {
      // 获取列表
      this.loading = true;
      try {
        const list = await categoryList([]);
        this.categorys  = list.data.list;
      } finally {
        this.loading = false;
      }
    },
    async getNovels() {
      // 获取列表
      this.loading = true;
      try {
        var id = this.$route.query.id;
        const list = await getBook(id);
        this.form  = list.data.result;
      } finally {
        this.loading = false;
      }
    }
  }
};
</script>
<style lang="scss" scoped>
.book_title {
  font-weight: 700 !important;
  font-size: 16px;
  padding: 20px 20px;
}
.help-block {
  display: block;
  margin-top: 5px;
  margin-bottom: 10px;
  color: #737373;
}
</style>