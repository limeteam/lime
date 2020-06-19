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
          <el-form-item label="搜索" prop="keyword">
            <el-input v-model="formData.keyword" placeholder="UID/用户名/昵称/手机号" class="filter-item" />
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
            <el-button type="primary" @click="onCreateUser">新增用户</el-button>
          </el-form-item>
          <selectedpanel />
        </el-form>
      </div>
      <!-- 过滤条件 -->
      <!--列表-->
      <el-table v-loading="loading" :data="items.list" border style="width: 100%">
        <el-table-column label="昵称(UID)" width="150px" align="center">
          <template slot-scope="scope">
            <img :src="scope.row.faceicon" style="width: 25px;" />
            &nbsp;{{scope.row.username}}&nbsp;
            <span
              style="color:#00A1CB"
            >({{scope.row.id}})</span>
          </template>
        </el-table-column>
        <el-table-column label="手机" prop="mobile" width="220px" />
        <el-table-column label="微信" align="center" width="80px">
          <template slot-scope="{row}">
            <span
              v-if="row.wechat!==''&& row.wechat.openid !==''"
              class="link-type"
              @click="handleWechatView(row.wechat)"
            >已绑定</span>
            <span v-else>--</span>
          </template>
        </el-table-column>
        <el-table-column label="性别" prop="sex" width="100px" />
        <el-table-column label="来源" prop="source" width="80px" />
        <el-table-column label="是否会员" prop="is_vip" width="80px" />
        <el-table-column label="金币余额" prop="coin" width="80px" />
        <el-table-column label="现金余额(元)" prop="amount" width="80px" />
        <el-table-column label="注册时间" prop="create_time" width="80px" />
        <el-table-column label="今日阅读时长" prop="today_reading_time" width="80px" />
        <el-table-column label="累计阅读时长" prop="total_reading_time" width="80px" />
        <el-table-column label="状态" prop="status" width="80px">
          <template slot-scope="{row}">
            <el-button v-if = "row.status == 1" size="mini" type="success" @click="handleModifyStatus(row,0)">正常</el-button>
            <el-button v-else size="mini" type="danger" @click="handleModifyStatus(row,1)">封禁</el-button>
          </template>
        </el-table-column>
        <el-table-column label="好友" prop="friends" width="80px" />
        <el-table-column label="操作" prop="operation" fixed="right">
          <template slot-scope="{row}">
            <el-button icon="el-icon-edit" size="mini" type="danger" @click="handleUpdate(row)">编辑</el-button>
          </template>
        </el-table-column>
      </el-table>
      <Pagination
        v-show="items.total_items>10"
        :total="items.total_items"
        :page.sync="formData.page"
        :limit.sync="formData.page_size"
        @pagination="getUsersList"
      />
      <!--列表 -->
    </el-card>
    <el-dialog :visible.sync="dialogWechatVisible" title="微信信息">
      <el-table :data="wechatData" border fit highlight-current-row style="width: 100%">
        <el-table-column prop="openid" label="Openid" />
        <el-table-column prop="nickname" label="昵称" />
        <el-table-column prop="headimgurl" label="头像">
          <template slot-scope="scope">
            <img :src="scope.row.headimgurl" style="width: 25px;" />
          </template>
        </el-table-column>
      </el-table>
      <span slot="footer" class="dialog-footer">
        <el-button type="primary" @click="dialogWechatVisible = false">关闭</el-button>
      </span>
    </el-dialog>
  </div>
</template>
<script>
import { userList, createUser, updateUser,updatestatus } from "@/api/lime-admin/users";
import Pagination from "@/components/Pagination";
import selectedpanel from "./components/selectedpanel";
import { USERS_GENDER, USERS_STATUS, USERS_ROBOTS } from "./emun/index.js";
export default {
  name: "UsersList",
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
        sort: "+id"
      },
      formData: {
        id: undefined,
        name: "",
        status: 0,
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
        0: "正常",
        1: "封禁"
      },
      vipMap: {
        0: "否",
        1: "是"
      },
      genderMap: {
        0: "未知",
        1: "男",
        2: "女"
      },
      dialogWechatVisible: false,
      wechatData: []
    };
  },
  computed: {},
  created() {},
  watch: {
    $route() {
      this.getUsersList();
    }
  },
  mounted() {
    this.getUsersList();
  },
  methods: {
    onSubmit() {
      // 查询按钮
      this.formData.page = 1;
      this.getUsersList();
    },
    onCreateUser() {
      this.$router.push({ path: "/users/create" });
    },
    on_refresh() {
      this.getUsersList();
    },
    handleUpdate(row) {
      this.$router.push({ path: "/users/update?id=" + row.id });
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
          this.getUsersList();
        })
        .catch(res => {
          this.$message.error(res.msg);
        });
    },
    currentChange(index) {
      // 分页
      this.formData.page = index;
      this.getUsersList();
    },
    handleWechatView(wechat) {
      this.wechatData = [wechat];
      this.dialogWechatVisible = true;
    },
    async getUsersList() {
      // 获取列表
      this.loading = true;
      try {
        var qStr = "";
        if (this.$route.query.source > 0) {
          qStr = "source=" + this.$route.query.source + ",";
        }
        if (this.$route.query.from_market > 0) {
          qStr += "from_market=" + this.$route.query.from_market + ",";
        }
        if (this.$route.query.gender > 0) {
          qStr += "gender=" + this.$route.query.gender + ",";
        }
        if (this.$route.query.status > 0) {
          qStr += "status=" + this.$route.query.status + ",";
        }
        if (this.$route.query.sort > 0) {
          qStr += "sort=" + this.$route.query.sort + ",";
        }
        if (this.formData.keyword != null && this.formData.keyword !== "") {
          qStr += "keyword=" + this.formData.keyword + ",";
        }
        this.listQuery.skip = (this.listQuery.page - 1) * this.listQuery.limit;
        this.listQuery.q = qStr.slice(0, -1);
        const list = await userList(this.listQuery);
        this.items.list = list.data.result;
        for (const v of this.items.list) {
          v.is_vip = this.vipMap[v.is_vip];
          v.sex = this.genderMap[v.sex];
          v.create_time = this.$moment(v.create_time).format(
            "YYYY-MM-DD HH:mm:ss"
          );
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

