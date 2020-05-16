<template>
  <el-card shadow="hover" class="box-card">
    <div>
      <el-row>
        <el-col :span="22">
          <span class="book_title"></span>
        </el-col>
        <el-col :span="2">
          <el-button type="primary" @click="onNovelList">小说列表</el-button>
        </el-col>
      </el-row>
      <div class="tips">
        <aside>
          温馨提示
          <br />上传作品时，程序会后台上传并拆分章节，异步上传处理章节会根据文本大小有一定的时间延迟，可稍后在
          <a href="/#/novel/books">小说管理</a> 中查看章节是否完整，上传文件只支持上传txt文本。
          章节拆分支持三种模式： 固定字数分章是按照2000字/章对文本进行拆分成章；
          标签分章需要上传固定格式的小说文本，章节文本之间用 ####lime#### 标识符分隔，可点击下载查看模板；
          智能分章程序会自动根据小说内容自动进行分章，建议选择智能分章模式；智能分章可能会出现分章不准确，可手动去章节管理进行修正。
        </aside>
      </div>
      <el-divider>上传小说</el-divider>
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
      <el-form-item label="封面" prop="cover">
        <el-upload
          class="cover-uploader"
          action
          :auto-upload="false"
          :on-change="handleImgSuccess"
          :show-file-list="false"
          :on-success="handleAvatarSuccess"
          :before-upload="beforeAvatarUpload"
        >
          <img v-if="imageUrl" :src="imageUrl" class="cover" />
          <i v-else class="el-icon-plus cover-uploader-icon"></i>
        </el-upload>
        <div class="help-block">建议大小225*300</div>
      </el-form-item>
      <el-form-item label="属性" prop="flag">
        <el-checkbox-group v-model="form.flag">
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
      <el-form-item label="拆分规则" prop="split_rule">
          <el-radio-group v-model="form.split_rule">
          <el-radio
            v-for="item in rules"
            :key="item.key"
            :label="item.display_name"
            :value="item.key"
          />
        </el-radio-group>
      </el-form-item>
      <el-form-item label="上传文件" prop="upload_file">
        <el-upload
          class="upload-demo"
          action=""
          :on-change="handleBookFileSuccess"
          :file-list="fileList"
        >
          <el-button size="small" type="primary">点击上传</el-button>
        </el-upload>
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
        <el-button type="primary" @click="createData">保存</el-button>
        <el-button>取消</el-button>
      </el-form-item>
    </el-form>
  </el-card>
</template>
<script>
import { categoryList } from "@/api/lime-admin/category";
import { createBook, uploadcover, uploadbookfile } from "@/api/lime-admin/book";
import {
  CATEGORY_CHANNEL,
  BOOK_ATTRS,
  BOOK_ONLINE_STATUS,
  BOOK_STATUS,
  BOOK_IS_SENSITIVITYS,
  UPLOAD_RULES
} from "./emun/index.js";
export default {
  name: "UploadBook",
  data() {
    return {
      form: {
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
        flag: [],
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
      rules: UPLOAD_RULES,
      book_online_status: BOOK_ONLINE_STATUS,
      book_status: BOOK_STATUS,
      book_is_sensitivity: BOOK_IS_SENSITIVITYS,
      categorys: [],
      fileList: [],
      imageUrl: ""
    };
  },
  mounted() {
    this.getCategorys();
  },
  methods: {
    onNovelList() {
      this.$router.push({ path: "/novel/books" });
    },
    resetForm() {
      // 重置
      this.$refs["dataForm"].resetFields();
    },
    createData() {
      this.$refs["dataForm"].validate(valid => {
        if (valid) {
          createBook(this.form).then(() => {
            this.dialogFormVisible = false;
            this.$notify({
              title: "成功",
              message: "保存成功",
              type: "success",
              duration: 2000
            });
            this.$store.dispatch("tagsView/delView", this.$route);
            this.$router.push({ path: "/novel/books" });
          });
        }
      });
    },
    handleAvatarSuccess(res, file) {
      this.imageUrl = URL.createObjectURL(file.raw);
      console.log(this.imageUrl);
    },
    beforeAvatarUpload(file) {
      const isJPG = file.type === "image/jpeg";
      const isLt2M = file.size / 1024 / 1024 < 10;
      if (!isJPG) {
        this.$message.error("上传头像图片只能是 JPG 格式!");
      }
      if (!isLt2M) {
        this.$message.error("上传头像图片大小不能超过 2MB!");
      }
      return isJPG && isLt2M;
    },
    async handleImgSuccess(file, fileList) {
      if (!file) return;
      let formData = new FormData();
      formData.append("file", file.raw);
      uploadcover(formData).then(res => {
        this.form.cover = res.result;
        this.dialogFormVisible = false;
        this.$notify({
          title: "成功",
          message: "上传成功",
          type: "success",
          duration: 2000
        });
      });
      this.imageUrl = URL.createObjectURL(file.raw);
    },

    async handleBookFileSuccess(file, fileList) {
      if (!file) return;
      let formData = new FormData();
      formData.append("file", file.raw);
      uploadbookfile(formData).then(res => {
        this.form.upload_file = res.result;
        // this.dialogFormVisible = false;
        this.$notify({
          title: "成功",
          message: "上传成功",
          type: "success",
          duration: 2000
        });
      });
    },
    async getCategorys() {
      // 获取列表
      this.loading = true;
      try {
        const list = await categoryList([]);
        this.categorys = list.data.list;
      } finally {
        this.loading = false;
      }
    }
  }
};
</script>
<style>
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
.cover-uploader .el-upload {
  border: 1px dashed #d9d9d9;
  border-radius: 6px;
  cursor: pointer;
  position: relative;
  overflow: hidden;
}
.cover-uploader .el-upload:hover {
  border-color: #409eff;
}
.cover-uploader-icon {
  font-size: 28px;
  color: #8c939d;
  width: 255px;
  height: 300px;
  line-height: 255px;
  text-align: center;
}
.cover {
  width: 255px;
  height: 300px;
  display: block;
}
.tips {
  padding-top: 10px;
}
aside {
  font-size: 13px;
}
</style>