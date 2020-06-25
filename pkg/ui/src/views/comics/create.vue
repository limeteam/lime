<template>
  <el-card shadow="hover" class="box-card">
    <div>
      <el-row>
        <el-col :span="22">
          <span class="comic_title"></span>
        </el-col>
        <el-col :span="2">
          <el-button type="primary" @click="onComicList">漫画列表</el-button>
        </el-col>
      </el-row>
      <el-divider>新增漫画</el-divider>
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
      <el-form-item label="所属分类" prop="category_ids">
        <el-checkbox-group v-model="form.category_ids">
          <el-checkbox
            v-for="item in categorys"
            :key="item.id"
            :label="item.name"
            :value="item.id"
          />
        </el-checkbox-group>
      </el-form-item>
      <el-form-item label="漫画名" prop="name">
        <el-input v-model="form.name" />
      </el-form-item>
      <el-form-item label="漫画原名" prop="old_name">
        <el-input v-model="form.old_name" />
      </el-form-item>
      <el-form-item label="来源" prop="source">
        <el-input v-model="form.source" />
      </el-form-item>
      <el-form-item label="作者" prop="author">
        <el-input v-model="form.author" />
      </el-form-item>
      <el-form-item label="作品简介" prop="desc">
        <el-input type="textarea" v-model="form.desc"></el-input>
      </el-form-item>
      <el-form-item label="横版封面" prop="vertical_cover">
        <el-upload
          class="cover-uploader"
          :data="dataObj"
          :before-upload="beforeUpload"
          accept="image/jpeg, image/gif, image/png, image/bmp"
          :on-success="handleVerticalCoverSuccess"
          action="https://upload-z2.qiniup.com"
        >
          <img v-if="verticalImageUrl" :src="verticalImageUrl" class="vertical_cover" />
          <i v-else class="el-icon-plus cover-uploader-icon"></i>
        </el-upload>
        <div class="help-block">建议大小600*400</div>
      </el-form-item>

      <el-form-item label="竖版封面" prop="horizontal_cover">
        <el-upload
          class="cover-uploader"
          :data="dataObj"
          :before-upload="beforeUpload"
          accept="image/jpeg, image/gif, image/png, image/bmp"
          :on-success="handleHorizontalCoverSuccess"
          action="https://upload-z2.qiniup.com"
        >
          <img v-if="horizontalImageUrl" :src="horizontalImageUrl" class="horizontal_cover" />
          <i v-else class="el-icon-plus cover-uploader-icon"></i>
        </el-upload>
        <div class="help-block">建议大小300*400</div>
      </el-form-item>
      <el-form-item label="免费章节数" prop="free_chapter">
        <el-input-number v-model="form.free_chapter" :step="1" step-strictly></el-input-number>
      </el-form-item>
      <el-form-item label="章节定价" prop="chapter_price">
        <el-input-number v-model="form.chapter_price" :step="1" step-strictly></el-input-number>
        <div class="help-block">免费小说设置成0</div>
      </el-form-item>
      <el-form-item label="浏览数" prop="views">
        <el-input-number v-model="form.views" :step="1" step-strictly></el-input-number>
        <div class="help-block">浏览数会乘以一定系数在前端以热度值形式显示</div>
      </el-form-item>
      <el-form-item label="收藏数" prop="collect_num">
        <el-input-number v-model="form.collect_num" :step="1" step-strictly></el-input-number>
      </el-form-item>
      <el-form-item label="属性" prop="flag">
        <el-checkbox-group v-model="form.flag">
          <el-checkbox
            v-for="item in comic_atts"
            :key="item.key"
            :label="item.display_name"
            :value="item.key"
          />
        </el-checkbox-group>
      </el-form-item>
      <el-form-item label="状态" prop="status">
        <el-select v-model="form.status" class="filter-item" placeholder="请选择">
          <el-option
            v-for="item in comic_status"
            :key="item.key"
            :label="item.display_name"
            :value="item.key"
          />
        </el-select>
      </el-form-item>
      <el-form-item label="上架状态" prop="online_status">
        <el-select v-model="form.online_status" class="filter-item" placeholder="请选择">
          <el-option
            v-for="item in comic_online_status"
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
            v-for="item in comic_is_sensitivity"
            :key="item.key"
            :label="item.display_name"
            :value="item.key"
          />
        </el-select>
        <div class="help-block">敏感小说不会出现在书库和搜索结果中，只能通过推荐位和推送等形式触达用户</div>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="createData">立即创建</el-button>
        <el-button>取消</el-button>
      </el-form-item>
    </el-form>
  </el-card>
</template>
<script>
import { categoryList } from "@/api/lime-admin/comicCategory";
import { createComic } from "@/api/lime-admin/comic";
import { getQiniuToken } from "@/api/lime-admin/upload";
import {
  CATEGORY_CHANNEL,
  COMIC_ATTRS,
  COMIC_ONLINE_STATUS,
  COMIC_STATUS,
  COMIC_IS_SENSITIVITYS
} from "./emun/index.js";
export default {
  name: "CreateComic",
  data() {
    return {
      form: {
        name: "",
        old_name: "",
        channel_id: 1,
        category_ids: [],
        desc: "",
        cover: "",
        author: "",
        source: "",
        vertical_cover: "",
        horizontal_cover: "",
        status: 0,
        flag: [],
        chapter_price: 0,
        free_chapter: 0,
        score: 0,
        chapter_id: 0,
        views: 0,
        collect_num: 0,
        online_status: 0,
        is_sensitivity: 0
      },
      category_channels: CATEGORY_CHANNEL,
      comic_atts: COMIC_ATTRS,
      comic_online_status: COMIC_ONLINE_STATUS,
      comic_status: COMIC_STATUS,
      comic_is_sensitivity: COMIC_IS_SENSITIVITYS,
      categorys: [],
      fileList: [],
      dataObj: { token: "", key: "" },
      domain: "",
      verticalImageUrl: '',
      horizontalImageUrl: ''
    };
  },
  mounted() {
    this.getCategorys();
  },
  methods: {
    onComicList() {
      this.$router.push({ path: "/comics/comic"});
    },
    resetForm() {
      // 重置
      this.$refs["dataForm"].resetFields();
    },
    createData() {
      this.$refs["dataForm"].validate(valid => {
        if (valid) {
          createComic(this.form).then(() => {
            this.dialogFormVisible = false;
            this.$notify({
              title: "成功",
              message: "新增成功",
              type: "success",
              duration: 2000
            });
            this.$store.dispatch('tagsView/delView', this.$route)
            this.$router.push({ path: "/comics/comic" });
          });
        }
      });
    },
    beforeUpload(file) {
      const _self = this;
      return new Promise((resolve, reject) => {
        getQiniuToken()
          .then(response => {
            const token = response.data.token;
            _self._data.dataObj.token = token;
            _self._data.domain = response.data.domain;
            _self._data.dataObj.key = "faceicon/" + file.name;
            resolve(true);
          })
          .catch(err => {
            console.log(err);
            reject(false);
          });
      });
    },
    handleVerticalCoverSuccess(res, file) {
      this.verticalImageUrl = this.domain + res.key;
      this.form.vertical_cover = this.verticalImageUrl;
    },
    handleHorizontalCoverSuccess(res, file) {
      this.horizontalImageUrl = this.domain + res.key;
      this.form.horizontal_cover = this.horizontalImageUrl;
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
    async getCategorys() {
      // 获取列表
      this.loading = true;
      try {
        const list = await categoryList([]);
        this.categorys  = list.data.list;
      } finally {
        this.loading = false;
      }
    }
  }
};
</script>
<style lang="scss" scoped>
.comic_title {
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
.vertical_cover {
  width: 255px;
  height: 300px;
  display: block;
}
.horizontal_cover {
  width: 255px;
  height: 300px;
  display: block;
}
</style>