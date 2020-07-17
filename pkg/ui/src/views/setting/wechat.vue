<template>
  <div class="app-container">
    <el-tabs v-model="activeName" @tab-click="handleClick">
      <el-tab-pane label="基本设置" name="first">
        <div>
          <el-card class="box-card">
            <div slot="header" class="clearfix">
              <span>基本配置</span>
              <el-button
                style="float: right; padding: 3px 0"
                type="text"
                @click="handleEdit"
                v-if="disableEditBase == 1"
              >编辑</el-button>
              <el-button
                style="float: right; padding: 3px 0"
                type="text"
                @click="saveData"
                v-else
              >保存</el-button>
            </div>
            <div class="text item">
              <el-form
                ref="dataForm"
                :model="form"
                label-position="left"
                label-width="150px"
                size="small"
              >
                <el-form-item label="URL" prop="url">
                  <el-col :span="20">
                    <el-input v-model="form.url" :disabled="disableEditBase == 1" />
                  </el-col>
                  <el-col :span="4" style="padding-left:10px;">
                    <button
                      class="el-icon-document-copy"
                      v-clipboard:copy="form.url"
                      v-clipboard:success="onCopy"
                      v-clipboard:error="onError"
                      title="点击复制"
                    />
                  </el-col>
                </el-form-item>
                <el-form-item label="Token" prop="token">
                  <el-col :span="17">
                    <el-input v-model="form.token" :disabled="disableEditBase == 1" />
                  </el-col>
                  <el-col :span="2" style="padding-left:10px;">
                    <button
                      class="el-icon-document-copy"
                      v-clipboard:copy="form.token"
                      v-clipboard:success="onCopy"
                      v-clipboard:error="onError"
                      title="点击复制"
                    />
                  </el-col>
                  <el-col :span="2" style="padding-left:5px;">
                    <el-button @click="genRandom" :disabled="disableEditBase == 1">随机生成</el-button>
                  </el-col>
                </el-form-item>
                <el-form-item label="EncodingAESKey" prop="encodingaeskey">
                  <el-col :span="20">
                    <el-input v-model="form.encodingaeskey" :disabled="disableEditBase == 1" />
                  </el-col>
                  <el-col :span="4" style="padding-left:10px;">
                    <button
                      class="el-icon-document-copy"
                      v-clipboard:copy="form.encodingaeskey"
                      v-clipboard:success="onCopy"
                      v-clipboard:error="onError"
                      title="点击复制"
                    />
                  </el-col>
                </el-form-item>
                <el-form-item label="消息加解密方式" prop="msg_encrypt_type">
                  <el-radio-group v-model="form.msg_encrypt_type">
                    <el-radio :label="1" disabled>明文模式</el-radio>
                    <el-radio :label="2" disabled>兼容模式</el-radio>
                    <el-radio :label="3" disabled>安全模式（推荐）</el-radio>
                  </el-radio-group>
                </el-form-item>
              </el-form>
            </div>
          </el-card>
        </div>
        <div style="padding-top:30px;">
          <el-card class="box-card">
            <div slot="header" class="clearfix">
              <span>公众号开发信息</span>
              <el-button style="float: right; padding: 3px 0" type="text" @click="saveData">编辑</el-button>
            </div>
            <div class="text item">
              <el-form
                ref="dataForm2"
                :model="form"
                label-position="left"
                label-width="200px"
                size="small"
              >
                <el-form-item label="开发者ID(AppID)" prop="title">
                  <el-col :span="20">
                    <el-input v-model="form.title" :disabled="true" />
                  </el-col>
                </el-form-item>
                <el-form-item label="开发者密码(AppSecret)" prop="title">
                  <el-col :span="17">
                    <el-input v-model="form.title" :disabled="true" />
                  </el-col>
                </el-form-item>
                <el-form-item label="IP白名单" prop="title">
                  <el-col :span="20">
                    <el-input v-model="form.title" :disabled="true" />
                  </el-col>
                </el-form-item>
              </el-form>
            </div>
          </el-card>
        </div>
      </el-tab-pane>
      <el-tab-pane label="菜单配置" name="second"></el-tab-pane>
      <el-tab-pane label="菜单链接" name="third">菜单链接</el-tab-pane>
    </el-tabs>
  </div>
</template>
<script>
import { wechatSetting, getWetchatSetting,getTokenAndKey } from "@/api/lime-admin/setting";

export default {
  name: "WechatSetting",
  data() {
    return {
      activeName: "first",
      disableEditBase: 1,
      form: {
        url: "",
        token: "",
        encodingaeskey: "",
        msg_encrypt_type: 1
      }
    };
  },
  mounted() {
    this.getSetting();
  },
  methods: {
    handleClick(tab, event) {
      //   if (tab.name == "first") {
      //     this.$router.push({ path: "/distributor/distributionSetting" });
      //   }
    },
    handleEdit() {
      this.disableEditBase = 0;
    },
    onCopy: function() {
      alert("复制成功");
    },

    onError: function() {
      alert("复制失败");
    },
    saveData() {
      this.disableEditBase = 1;
      this.$refs["dataForm"].validate(valid => {
        if (valid) {
          var data = {
            config_code: "wechatSetting",
            config_value: this.form
          };
          wechatSetting("wechatSetting", data).then(() => {
            this.$notify({
              title: "成功",
              message: "保存成功",
              type: "success",
              duration: 2000
            });
          });
        }
      });
    },
    async getSetting() {
      this.loading = true;
      try {
        const list = await getWetchatSetting("wechatSetting");
        if (list.data.result !== null) {
          this.form = list.data.result;
        }
      } finally {
        this.loading = false;
      }
    },
    async genRandom() {
      this.loading = true;
      try {
        const list = await getTokenAndKey();
        if (list.data.result !== null) {
          this.form.url = list.data.receiveUrl;
          this.form.token = list.data.token;
          this.form.encodingaeskey = list.data.EncodingAESKey;
          this.form.msg_encrypt_type = 3;
        }
      } finally {
        this.loading = false;
      }
    }
  }
};
</script>
<style>
.text {
  font-size: 14px;
}

.item {
  margin-bottom: 18px;
}

.clearfix:before,
.clearfix:after {
  display: table;
  content: "";
}
.clearfix:after {
  clear: both;
}

.box-card {
  width: 680px;
}
</style>