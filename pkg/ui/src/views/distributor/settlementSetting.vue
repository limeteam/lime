<template>
  <div class="app-container">
    <el-tabs v-model="activeName" @tab-click="handleClick">
      <el-tab-pane label="基本设置" name="first">基本设置</el-tab-pane>
      <el-tab-pane label="结算设置" name="second">结算设置</el-tab-pane>
      <el-tab-pane label="申请协议" name="third">申请协议</el-tab-pane>
      <el-tab-pane label="文案样式" name="fourth">文案样式</el-tab-pane>
      <el-tab-pane label="推送通知" name="five">推送通知</el-tab-pane>
    </el-tabs>
  </div>
</template>
<script>
export default {
  data() {
    return {
      activeName: "second"
    };
  },
  methods: {
    handleClick(tab, event) {
      if(tab.name == 'first'){
        this.$router.push({ path: "/distributor/distributionSetting" });
      }
    }
  }
};
</script>