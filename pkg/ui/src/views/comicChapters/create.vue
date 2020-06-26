<template>
  <el-card shadow="hover" class="box-card">
    <div>
      <el-row>
        <el-col :span="22">
          <span class="book_title"></span>
        </el-col>
        <el-col :span="2">
          <a href="/#/comics/chapters" class="el-button el-button--primary">章节列表</a>
        </el-col>
      </el-row>
      <el-divider>新增章节</el-divider>
    </div>
    <el-form ref="dataForm" :model="form" label-position="left" label-width="200px" size="small">
      <el-tabs type="border-card">
        <el-form-item label="章节封面" prop="cover">
          <el-upload
            :data="dataObj"
            :multiple="true"
            :before-upload="beforeUpload"
            accept="image/jpeg, image/gif, image/png, image/bmp"
            :on-success="handleAvatarSuccess"
            action="https://upload-z2.qiniup.com"
            drag
          >
            <i class="el-icon-upload" />
            <div class="el-upload__text">
              将文件拖到此处，或
              <em>点击上传</em>
            </div>
          </el-upload>
          <img v-if="imageUrl" :src="imageUrl" class="avatar" />
        </el-form-item>
        <el-form-item label="章节名称" prop="title">
          <el-col :span="17">
            <el-input v-model="form.title" />
          </el-col>
        </el-form-item>
        <el-form-item label="章节序号" prop="chapter_no">
          <el-input-number v-model="form.chapter_no" :step="1" step-strictly></el-input-number>
        </el-form-item>
        <el-form-item label="是否收费" prop="is_vip">
          <el-select v-model="form.is_vip" class="filter-item" placeholder="请选择">
            <el-option
              v-for="item in is_vips"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="上传类型" prop="upload_type">
          <el-select v-model="form.upload_type" class="filter-item" placeholder="请选择">
            <el-option
              v-for="item in uploadTypes"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="章节图" prop="desc">
          <el-upload
            :data="dataObj"
            action="https://upload-z2.qiniup.com"
            accept="image/jpeg, image/gif, image/png, image/bmp"
            list-type="picture-card"
            :before-upload="beforeUpload"
            :on-success="handleAvatarSuccess"
            :auto-upload="false"
          >
            <i slot="default" class="el-icon-plus"></i>
            <div slot="file" slot-scope="{file}">
              <img class="el-upload-list__item-thumbnail" :src="file.url" alt />
              <span class="el-upload-list__item-actions">
                <span class="el-upload-list__item-preview" @click="handlePictureCardPreview(file)">
                  <i class="el-icon-zoom-in"></i>
                </span>
                <span
                  v-if="!disabled"
                  class="el-upload-list__item-delete"
                  @click="handleDownload(file)"
                >
                  <i class="el-icon-download"></i>
                </span>
                <span
                  v-if="!disabled"
                  class="el-upload-list__item-delete"
                  @click="handleRemove(file)"
                >
                  <i class="el-icon-delete"></i>
                </span>
              </span>
            </div>
          </el-upload>
          <el-dialog :visible.sync="dialogVisible">
            <img width="100%" :src="dialogImageUrl" alt />
          </el-dialog>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="createData">保存</el-button>
          <el-button @click="prev">返回</el-button>
        </el-form-item>
      </el-tabs>
    </el-form>
  </el-card>
</template>
<script>
import { createChapter } from "@/api/lime-admin/comicChapter";
export default {
  name: "CreateChapters",
  data() {
    return {
      form: {
        title: "",
        comic_id: 0,
        chapter_no: 0,
        is_vip: 2,
        upload_type: 1,
        desc: "",
        cover: ""
      },
      is_vips: [
        {
          value: 1,
          label: "是"
        },
        {
          value: 2,
          label: "否"
        }
      ],
      uploadTypes: [
        {
          value: 1,
          label: "多图上传"
        },
        {
          value: 2,
          label: "链接上传"
        }
      ],
      dataObj: { token: "", key: "" },
      domain: "",
      imageUrl: "",
      dialogImageUrl: "",
      dialogVisible: false,
      disabled: false
    };
  },
  mounted() {},
  methods: {
    resetForm() {
      // 重置
      this.$refs["dataForm"].resetFields();
    },
    prev() {
      this.$router.go(-1);
    },
    handleRemove(file) {
      console.log(file);
    },
    handlePictureCardPreview(file) {
      this.dialogImageUrl = file.url;
      this.dialogVisible = true;
    },
    handleDownload(file) {
      console.log(file);
    },
    beforeUpload(file) {
      const _self = this;
      return new Promise((resolve, reject) => {
        getQiniuToken()
          .then(response => {
            const token = response.data.token;
            _self._data.dataObj.token = token;
            _self._data.domain = response.data.domain;
            _self._data.dataObj.key = "comic/chapter/" + file.name;
            resolve(true);
          })
          .catch(err => {
            console.log(err);
            reject(false);
          });
      });
    },
    handleAvatarSuccess(res, file) {
      this.imageUrl = this.domain + res.key;
      this.form.faceicon = this.imageUrl;
    },
    createData() {
      this.$refs["dataForm"].validate(valid => {
        if (valid) {
          this.form.comic_id = parseInt(this.$route.query.comic_id);
          createChapter(this.form).then(() => {
            this.dialogFormVisible = false;
            this.$notify({
              title: "成功",
              message: "新增成功",
              type: "success",
              duration: 2000
            });
            this.$store.dispatch("tagsView/delView", this.$route);
            this.$router.push({ path: "/comics/chapters?comic_id=" + this.form.comic_id });
          });
        }
      });
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

.avatar-uploader .el-upload {
  border: 1px dashed #d9d9d9;
  border-radius: 6px;
  cursor: pointer;
  position: relative;
  overflow: hidden;
}
.avatar-uploader .el-upload:hover {
  border-color: #409eff;
}
.avatar-uploader-icon {
  font-size: 28px;
  color: #8c939d;
  width: 178px;
  height: 178px;
  line-height: 178px;
  text-align: center;
}
.avatar {
  width: 178px;
  height: 178px;
  display: block;
}
</style>
