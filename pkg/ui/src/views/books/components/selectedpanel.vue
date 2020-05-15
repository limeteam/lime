<template>
<el-card class="box-card">
  <el-row>
    <el-col :span="1">
      <div>频道:</div>
    </el-col>
    <el-col :span="23">
      <div>
        <el-link @click='JumpUrl("channel_id",0)' :class="this.$route.query.channel_id == 0 || this.$route.query.channel_id == null? 'selected': 'unselected'">全部</el-link>
        <el-link @click='JumpUrl("channel_id",1)' :class="this.$route.query.channel_id == 1 ? 'selected': 'unselected'">男生</el-link>
        <el-link @click='JumpUrl("channel_id",2)' :class="this.$route.query.channel_id == 2 ? 'selected': 'unselected'">女生</el-link>
      </div>
    </el-col>
  </el-row>
  <el-row>
    <el-col :span="1">
      <div>分类:</div>
    </el-col>
    <el-col :span="23">
      <div>
        <el-link @click='JumpUrl("cid",0)' :class="this.$route.query.cid == 0 || this.$route.query.cid == null  ? 'selected': 'unselected'">全部</el-link>
        <el-link @click='JumpUrl("cid",1)' :class="this.$route.query.cid == 1 ? 'selected': 'unselected'">都市现实</el-link>
        <el-link @click='JumpUrl("cid",2)' :class="this.$route.query.cid == 2 ? 'selected': 'unselected'">悬疑灵异</el-link>
        <el-link @click='JumpUrl("cid",3)' :class="this.$route.query.cid == 3 ? 'selected': 'unselected'">军事历史</el-link>
        <el-link @click='JumpUrl("cid",4)' :class="this.$route.query.cid == 4 ? 'selected': 'unselected'">玄幻奇幻</el-link>
        <el-link @click='JumpUrl("cid",5)' :class="this.$route.query.cid == 5 ? 'selected': 'unselected'">武侠仙侠</el-link>
        <el-link @click='JumpUrl("cid",6)' :class="this.$route.query.cid == 6 ? 'selected': 'unselected'">网游竞技</el-link>
        <el-link @click='JumpUrl("cid",7)' :class="this.$route.query.cid == 7 ? 'selected': 'unselected'">科幻未来</el-link>
      </div>
    </el-col>
  </el-row>
  <el-row>
    <el-col :span="1">
      <div>连载:</div>
    </el-col>
    <el-col :span="23">
      <div>
        <el-link @click='JumpUrl("status",0)' :class="this.$route.query.status == 0 || this.$route.query.status == null? 'selected': 'unselected'">全部</el-link>
        <el-link @click='JumpUrl("status",1)' :class="this.$route.query.status == 1 ? 'selected': 'unselected'">连载中</el-link>
        <el-link @click='JumpUrl("status",2)' :class="this.$route.query.status == 2 ? 'selected': 'unselected'">全完结</el-link>
      </div>
    </el-col>
  </el-row>
  <el-row>
    <el-col :span="1">
      <div>上架:</div>
    </el-col>
    <el-col :span="23">
      <div>
        <el-link @click='JumpUrl("online_status",0)'  :class="this.$route.query.online_status == 0 || this.$route.query.online_status == null? 'selected': 'unselected'">全部</el-link>
        <el-link @click='JumpUrl("online_status",1)' :class="this.$route.query.online_status == 1 ? 'selected': 'unselected'">已上架</el-link>
        <el-link @click='JumpUrl("online_status",2)' :class="this.$route.query.online_status == 2 ? 'selected': 'unselected'">已下架</el-link>
      </div>
    </el-col>
  </el-row>
  <el-row>
    <el-col :span="1">
      <div>敏感:</div>
    </el-col>
    <el-col :span="23">
      <div>
        <el-link @click='JumpUrl("is_sensitivity",0)' :class="this.$route.query.is_sensitivity == 0 || this.$route.query.is_sensitivity == null? 'selected': 'unselected'">全部</el-link>
        <el-link @click='JumpUrl("is_sensitivity",1)' :class="this.$route.query.is_sensitivity == 1 ? 'selected': 'unselected'">不敏感</el-link>
        <el-link @click='JumpUrl("is_sensitivity",2)' :class="this.$route.query.is_sensitivity == 2 ? 'selected': 'unselected'">敏感</el-link>
      </div>
    </el-col>
  </el-row>
  <el-row>
    <el-col :span="1">
      <div>属性:</div>
    </el-col>
    <el-col :span="23">
      <div>
        <el-link @click='JumpUrl("flag",0)' :class="this.$route.query.flag == 0 || this.$route.query.flag == null ? 'selected': 'unselected'">全部</el-link>
        <el-link @click='JumpUrl("flag",1)'  :class="this.$route.query.flag == 1 ? 'selected': 'unselected'">免费</el-link>
        <el-link @click='JumpUrl("flag",2)'  :class="this.$route.query.flag == 2 ? 'selected': 'unselected'">新书</el-link>
        <el-link @click='JumpUrl("flag",3)'  :class="this.$route.query.flag == 3 ? 'selected': 'unselected'">热门</el-link>
        <el-link @click='JumpUrl("flag",4)'  :class="this.$route.query.flag == 4 ? 'selected': 'unselected'">会员</el-link>
        <el-link @click='JumpUrl("flag",5)'  :class="this.$route.query.flag == 5 ? 'selected': 'unselected'">爽文</el-link>
        <el-link @click='JumpUrl("flag",6)'  :class="this.$route.query.flag == 6 ? 'selected': 'unselected'">精选</el-link>
        <el-link @click='JumpUrl("flag",7)'  :class="this.$route.query.flag == 7 ? 'selected': 'unselected'">大神</el-link>
      </div>
    </el-col>
  </el-row>
  <el-row>
    <el-col :span="1">
      <div>字数:</div>
    </el-col>
    <el-col :span="23">
      <div>
        <el-link @click='JumpUrl("word",0)' :class="this.$route.query.word == 0 || this.$route.query.word == null ? 'selected': 'unselected'">全部</el-link>
        <el-link @click='JumpUrl("word",1)' :class="this.$route.query.word == 1 ? 'selected': 'unselected'">20万字以下</el-link>
        <el-link @click='JumpUrl("word",2)' :class="this.$route.query.word == 2 ? 'selected': 'unselected'">20-30万字</el-link>
        <el-link @click='JumpUrl("word",3)' :class="this.$route.query.word == 3 ? 'selected': 'unselected'">30-50万字</el-link>
        <el-link @click='JumpUrl("word",4)' :class="this.$route.query.word == 4 ? 'selected': 'unselected'">50-100万字</el-link>
        <el-link @click='JumpUrl("word",5)' :class="this.$route.query.word == 5 ? 'selected': 'unselected'">100-200万字</el-link>
        <el-link @click='JumpUrl("word",6)' :class="this.$route.query.word == 6 ? 'selected': 'unselected'">200字以上</el-link>
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
        <el-link @click='JumpUrl("sort",1)' :class="this.$route.query.sort == 1 ? 'selected': 'unselected'">最新上架</el-link>
        <el-link @click='JumpUrl("sort",2)' :class="this.$route.query.sort == 2 ? 'selected': 'unselected'">最近更新</el-link>
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
      current_link: "",
    }
  },

  created () {
    this.fetchData()
    // this.getChannel(1)
  },
  watch: {
    // 如果路由有变化，会再次执行该方法
    '$route': 'fetchData'
  },
  methods: {
    fetchData () {
      this.current_link =  document.location.href
    },
    JumpUrl(paramKey,current_id){
      var paramStr = "";
      var allParam  = this.$route.query
      for (const keys in allParam) {
        if (paramKey !== keys && allParam[keys] > 0){
            paramStr += keys + "=" + allParam[keys] + "&"
        }
      }
      paramStr += paramKey + "="+ current_id
      this.$router.push({ path: "/novel/books?" + paramStr });
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