<template>
  <el-card shadow="hover" class="box-card">
    <div>
      <el-row>
        <el-col :span="22">
          <span class="book_title"></span>
        </el-col>
        <el-col :span="2">
          <a href="/#/novel/chapters/list" class="el-button el-button--primary">章节列表</a>
        </el-col>
      </el-row>
      <el-divider>修改章节</el-divider>
    </div>
    <el-form ref="dataForm" :model="form" label-position="left" label-width="200px" size="small">
      <el-tabs type="border-card">
        <el-form-item label="标题" prop="title">
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
        <el-form-item label="内容" prop="desc">
          <el-col :span="19">
            <tinymce v-model="form.desc" :height="200" :menubar="menubar" :toolbar="toolbar" />
          </el-col>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="updateData">保存</el-button>
          <el-button @click="prev">返回</el-button>
        </el-form-item>
      </el-tabs>
    </el-form>
  </el-card>
</template>
<script>
import Tinymce from "@/components/Tinymce";
import { updateChapter, getChapter } from "@/api/lime-admin/chapter";
export default {
  name: "CreateBook",
  components: { Tinymce },
  data() {
    return {
      form: {
        title: "",
        chapter_no: 0,
        is_vip: 0,
        desc: ""
      },
      is_vips: [
        {
          value: 0,
          label: "全部"
        },
        {
          value: 1,
          label: "是"
        },
        {
          value: 2,
          label: "否"
        }
      ],
      menubar: "",
      toolbar: ['undo redo formatselect bold italic backcolor alignleft aligncenter alignright alignjustify bullist numlist outdent indent removeformat']
    };
  },
  mounted() {
    this.getChapterInfo();
  },
  methods: {
    resetForm() {
      // 重置
      this.$refs["dataForm"].resetFields();
    },
    prev() {
      this.$router.go(-1);
    },
    async getChapterInfo() {
      // 获取列表
      this.loading = true;
      try {
        var id = this.$route.query.chapter_id;
        const list = await getChapter(id);
        this.form = list.data.result;
      } finally {
        this.loading = false;
      }
    },
    updateData() {
      this.$refs["dataForm"].validate(valid => {
        if (valid) {
          this.form.novel_id = this.$route.query.book_id;
          this.form.chapter_id = this.$route.query.chapter_id;
          updateChapter(this.form.chapter_id,this.form).then(() => {
            this.dialogFormVisible = false;
            this.$notify({
              title: "成功",
              message: "保存成功",
              type: "success",
              duration: 2000
            });
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
</style>