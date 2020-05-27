<template>
<el-card class="box-card">
  <el-row>
    <el-col :span="1">
      <div>来源:</div>
    </el-col>
    <el-col :span="23">
      <div>
        <el-link @click='JumpUrl("source",0)' :class="this.$route.query.source == 0 || this.$route.query.source == null? 'selected': 'unselected'">全部</el-link>
        <el-link @click='JumpUrl("source",1)' :class="this.$route.query.source == 1 ? 'selected': 'unselected'">PC</el-link>
        <el-link @click='JumpUrl("source",2)' :class="this.$route.query.source == 2 ? 'selected': 'unselected'">Wap</el-link>
        <el-link @click='JumpUrl("source",3)' :class="this.$route.query.source == 2 ? 'selected': 'unselected'">Android</el-link>
        <el-link @click='JumpUrl("source",4)' :class="this.$route.query.source == 2 ? 'selected': 'unselected'">IOS</el-link>
        <el-link @click='JumpUrl("source",5)' :class="this.$route.query.source == 2 ? 'selected': 'unselected'">公众号</el-link>
        <el-link @click='JumpUrl("source",6)' :class="this.$route.query.source == 2 ? 'selected': 'unselected'">小程序</el-link>
      </div>
    </el-col>
  </el-row>
  <el-row>
    <el-col :span="1">
      <div>渠道:</div>
    </el-col>
    <el-col :span="23">
      <div>
        <el-link @click='JumpUrl("from_market",0)' :class="this.$route.query.from_market == 0 || this.$route.query.from_market == null  ? 'selected': 'unselected'">全部</el-link>
        <el-link @click='JumpUrl("from_market",1)' :class="this.$route.query.from_market == 1 ? 'selected': 'unselected'">360</el-link>
        <el-link @click='JumpUrl("from_market",2)' :class="this.$route.query.from_market == 2 ? 'selected': 'unselected'">wandoujia</el-link>
        <el-link @click='JumpUrl("from_market",3)' :class="this.$route.query.from_market == 3 ? 'selected': 'unselected'">yingyongbao</el-link>
        <el-link @click='JumpUrl("from_market",4)' :class="this.$route.query.from_market == 4 ? 'selected': 'unselected'">baiduzhushou</el-link>
        <el-link @click='JumpUrl("from_market",5)' :class="this.$route.query.from_market == 5 ? 'selected': 'unselected'">oppo</el-link>
        <el-link @click='JumpUrl("from_market",6)' :class="this.$route.query.from_market == 6 ? 'selected': 'unselected'">vivo</el-link>
        <el-link @click='JumpUrl("from_market",7)' :class="this.$route.query.from_market == 7 ? 'selected': 'unselected'">xiaomi</el-link>
        <el-link @click='JumpUrl("from_market",8)' :class="this.$route.query.from_market == 8 ? 'selected': 'unselected'">huawei</el-link>
        <el-link @click='JumpUrl("from_market",9)' :class="this.$route.query.from_market == 9 ? 'selected': 'unselected'">IOS</el-link>
        <el-link @click='JumpUrl("from_market",10)' :class="this.$route.query.from_market == 10 ? 'selected': 'unselected'">meizu</el-link>
        <el-link @click='JumpUrl("from_market",11)' :class="this.$route.query.from_market == 11 ? 'selected': 'unselected'">pp</el-link>
      </div>
    </el-col>
  </el-row>
  <el-row>
    <el-col :span="1">
      <div>性别:</div>
    </el-col>
    <el-col :span="23">
      <div>
        <el-link @click='JumpUrl("gender",0)' :class="this.$route.query.gender == 0 || this.$route.query.status == null? 'selected': 'unselected'">全部</el-link>
        <el-link @click='JumpUrl("gender",1)' :class="this.$route.query.gender == 1 ? 'selected': 'unselected'">男</el-link>
        <el-link @click='JumpUrl("gender",2)' :class="this.$route.query.gender == 2 ? 'selected': 'unselected'">女</el-link>
      </div>
    </el-col>
  </el-row>
  <el-row>
    <el-col :span="1">
      <div>状态:</div>
    </el-col>
    <el-col :span="23">
      <div>
        <el-link @click='JumpUrl("status",0)'  :class="this.$route.query.status == 0 || this.$route.query.online_status == null? 'selected': 'unselected'">全部</el-link>
        <el-link @click='JumpUrl("status",1)' :class="this.$route.query.status == 1 ? 'selected': 'unselected'">正常</el-link>
        <el-link @click='JumpUrl("status",2)' :class="this.$route.query.status == 2 ? 'selected': 'unselected'">封禁</el-link>
      </div>
    </el-col>
  </el-row>
  <el-row>
    <el-col :span="1">
      <div>排序:</div>
    </el-col>
    <el-col :span="23">
      <div>
        <el-link @click='JumpUrl("sort",0)' :class="this.$route.query.sort == 0 || this.$route.query.sort == null ? 'selected': 'unselected'" :underline="false">全部</el-link>
        <el-link @click='JumpUrl("sort",1)' :class="this.$route.query.sort == 1 ? 'selected': 'unselected'">金币</el-link>
        <el-link @click='JumpUrl("sort",2)' :class="this.$route.query.sort == 2 ? 'selected': 'unselected'">现金</el-link>
      </div>
    </el-col>
  </el-row>
</el-card>
</template>
<script>
export default {
  name: 'selectedpanel',
  data() {
    return {
    }
  },
  methods: {
    JumpUrl(paramKey,current_id){
      var paramStr = "";
      var allParam  = this.$route.query
      for (const keys in allParam) {
        if (paramKey !== keys && allParam[keys] > 0){
            paramStr += keys + "=" + allParam[keys] + "&"
        }
      }
      paramStr += paramKey + "="+ current_id
      this.$router.push({ path: "/users?" + paramStr });
    },
  }

}
</script>
<style lang="scss" scoped>
.el-row {
    margin-bottom: 10px;
    &:last-child {
      margin-bottom: 0;
    }
  }
.el-link {
  font-size: 16px;
  margin-right: 10px;
  padding: 1px 5px;
}
.selected {
  color: #fff;
  background-color: #32c5d2;
  border-color: #32c5d2;
  border-radius: 4px;
}
.unselected {
  color: #1890ff;
}
</style>
