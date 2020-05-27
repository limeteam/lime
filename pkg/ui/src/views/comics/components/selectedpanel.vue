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
        <el-link @click='JumpUrl("sorts",0)' :class="this.$route.query.sorts == 0 || this.$route.query.sorts == null  ? 'selected': 'unselected'">全部</el-link>
        <el-link @click='JumpUrl("sorts",1)' :class="this.$route.query.sorts == 1 ? 'selected': 'unselected'">日常</el-link>
        <el-link @click='JumpUrl("sorts",2)' :class="this.$route.query.sorts == 2 ? 'selected': 'unselected'">搞笑</el-link>
        <el-link @click='JumpUrl("sorts",3)' :class="this.$route.query.sorts == 3 ? 'selected': 'unselected'">热血</el-link>
        <el-link @click='JumpUrl("sorts",4)' :class="this.$route.query.sorts == 4 ? 'selected': 'unselected'">古风</el-link>
        <el-link @click='JumpUrl("sorts",5)' :class="this.$route.query.sorts == 5 ? 'selected': 'unselected'">猎奇</el-link>
        <el-link @click='JumpUrl("sorts",6)' :class="this.$route.query.sorts == 6 ? 'selected': 'unselected'">同人</el-link>
        <el-link @click='JumpUrl("sorts",7)' :class="this.$route.query.sorts == 7 ? 'selected': 'unselected'">都市</el-link>
        <el-link @click='JumpUrl("sorts",8)' :class="this.$route.query.sorts == 8 ? 'selected': 'unselected'">冒险</el-link>
        <el-link @click='JumpUrl("sorts",9)' :class="this.$route.query.sorts == 9 ? 'selected': 'unselected'">恋爱</el-link>
        <el-link @click='JumpUrl("sorts",10)' :class="this.$route.query.sorts == 10 ? 'selected': 'unselected'">校园</el-link>
        <el-link @click='JumpUrl("sorts",11)' :class="this.$route.query.sorts == 11 ? 'selected': 'unselected'">日漫</el-link>
        <el-link @click='JumpUrl("sorts",12)' :class="this.$route.query.sorts == 12 ? 'selected': 'unselected'">治愈</el-link>
        <el-link @click='JumpUrl("sorts",13)' :class="this.$route.query.sorts == 13 ? 'selected': 'unselected'">玄幻</el-link>
        <el-link @click='JumpUrl("sorts",14)' :class="this.$route.query.sorts == 14 ? 'selected': 'unselected'">奇幻</el-link>
        <el-link @click='JumpUrl("sorts",15)' :class="this.$route.query.sorts == 15 ? 'selected': 'unselected'">恐怖</el-link>
        <el-link @click='JumpUrl("sorts",16)' :class="this.$route.query.sorts == 16 ? 'selected': 'unselected'">悬疑</el-link>
      </div>
    </el-col>
  </el-row>
  <el-row>
    <el-col :span="1">
      <div>状态:</div>
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
      <div>属性:</div>
    </el-col>
    <el-col :span="23">
      <div>
        <el-link @click='JumpUrl("flag",0)' :class="this.$route.query.flag == 0 || this.$route.query.flag == null ? 'selected': 'unselected'">全部</el-link>
        <el-link @click='JumpUrl("flag",1)'  :class="this.$route.query.flag == 1 ? 'selected': 'unselected'">免费</el-link>
        <el-link @click='JumpUrl("flag",2)'  :class="this.$route.query.flag == 2 ? 'selected': 'unselected'">热门</el-link>
        <el-link @click='JumpUrl("flag",3)'  :class="this.$route.query.flag == 3 ? 'selected': 'unselected'">会员</el-link>
        <el-link @click='JumpUrl("flag",4)'  :class="this.$route.query.flag == 4 ? 'selected': 'unselected'">推荐</el-link>
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
      this.$router.push({ path: "/comics/comic?" + paramStr });
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
