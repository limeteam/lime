<template>
  <el-card shadow="hover" class="box-card">
    <div>
      <el-row>
        <el-col :span="22">
          <span class="book_title"></span>
        </el-col>
        <el-col :span="2">
          <el-button type="primary" @click="onUserList">用户列表</el-button>
        </el-col>
      </el-row>
      <el-divider>编辑用户</el-divider>
    </div>
    <el-form
      ref="dataForm"
      :model="form"
      label-position="left"
      label-width="120px"
      style="width: 500px; margin-left:50px;"
    >
      <el-form-item label="昵称" prop="username">
        <el-input v-model="form.username" />
      </el-form-item>
      <el-form-item label="手机" prop="mobile">
        <el-input v-model="form.mobile" />
      </el-form-item>
      <el-form-item label="性别" prop="sex">
        <el-select v-model="form.sex" class="filter-item" placeholder="请选择">
          <el-option v-for="item in genders" :key="item.key" :label="item.display_name" :value="item.key" />
        </el-select>
      </el-form-item>
      <el-form-item label="头像" prop="faceicon">
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
      <el-form-item label="状态" prop="status">
        <el-select v-model="form.status" class="filter-item" placeholder="请选择">
          <el-option
            v-for="item in users_status"
            :key="item.key"
            :label="item.display_name"
            :value="item.key"
          />
        </el-select>
      </el-form-item>
      <!-- <el-form-item label="免验登录" prop="is_robot">
        <el-select v-model="form.is_robot" class="filter-item" placeholder="请选择">
          <el-option
            v-for="item in robots"
            :key="item.key"
            :label="item.display_name"
            :value="item.key"
          />
        </el-select>
        <div class="help-block">免验登录的用户可以输入任意验证码登录,此功能只应用于手机号登录</div>
      </el-form-item> -->
      <el-form-item>
        <el-button type="primary" @click="updateData">保存</el-button>
        <el-button>返回</el-button>
      </el-form-item>
    </el-form>
  </el-card>
</template>
<script>
import { getUser,updateUser } from "@/api/lime-admin/users";
import { getQiniuToken } from "@/api/lime-admin/upload";
import { USERS_GENDER, USERS_STATUS, USERS_ROBOTS } from "./emun/index.js";
export default {
  name: "CreateUser",
  data() {
    return {
      form: {
        username: "",
        sex: 1,
        mobile: "",
        faceicon: "",
        is_robot: "",
        gender: 1,
        upload_file: ""
      },
      robots: USERS_ROBOTS,
      genders: USERS_GENDER,
      users_status: USERS_STATUS,
      dataObj: { token: "", key: "" },
      domain: "",
      imageUrl: ""
    };
  },
  mounted() {
      this.getUsers();
  },
  methods: {
    onUserList() {
      this.$router.push({ path: "/users" });
    },
    resetForm() {
      // 重置
      this.$refs["dataForm"].resetFields();
    },
    async getUsers() {
      // 获取用户信息
      this.loading = true;
      try {
        var id = this.$route.query.id;
        const list = await getUser(id);
        this.form = list.data.result;
        if (list.data.result.faceicon.indexOf("http") !== -1){
            this.imageUrl = list.data.result.faceicon;
        }else{
            this.imageUrl = process.env.VUE_APP_CONFIG_API + list.data.result.faceicon;
        }
      } finally {
        this.loading = false;
      }
    },
    updateData() {
      this.$refs["dataForm"].validate(valid => {
        if (valid) {
          updateUser(this.form.id,this.form).then(() => {
            this.dialogFormVisible = false;
            this.$notify({
              title: "成功",
              message: "更新成功",
              type: "success",
              duration: 2000
            });
            this.$store.dispatch("tagsView/delView", this.$route);
            this.$router.push({ path: "/users" });
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
    handleAvatarSuccess(res, file) {
      this.imageUrl = this.domain + res.key;
      this.form.faceicon = this.imageUrl;
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