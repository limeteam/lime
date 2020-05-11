<template>
<el-card class="box-card">
  <el-row>
    <el-col :span="1">
      <div>频道:</div>
    </el-col>
    <el-col :span="23">
      <div>
        <el-link :href="`${current_link}`" class="wenzi">全部</el-link>
        <el-link :href="`${current_link}&channel_id=1`" type="primary">男生</el-link>
        <el-link :href="`${current_link}&channel_id=2`" type="primary">女生</el-link>
      </div>
    </el-col>
  </el-row>
  <el-row>
    <el-col :span="1">
      <div>分类:</div>
    </el-col>
    <el-col :span="23">
      <div>
        <el-link :href="`${current_link}`" class="wenzi">全部</el-link>
        <el-link :href="`${current_link}&cid=1`" type="primary">都市现实</el-link>
        <el-link :href="`${current_link}&cid=2`" type="primary">悬疑灵异</el-link>
        <el-link :href="`${current_link}&cid=3`" type="primary">军事历史</el-link>
        <el-link :href="`${current_link}&cid=4`" type="primary">玄幻奇幻</el-link>
        <el-link :href="`${current_link}&cid=5`" type="primary">武侠仙侠</el-link>
        <el-link :href="`${current_link}&cid=6`" type="primary">网游竞技</el-link>
        <el-link :href="`${current_link}&cid=7`" type="primary">科幻未来</el-link>
      </div>
    </el-col>
  </el-row>
  <el-row>
    <el-col :span="1">
      <div>连载:</div>
    </el-col>
    <el-col :span="23">
      <div>
        <el-link :href="`${current_link}`" class="wenzi">全部</el-link>
        <el-link :href="`${current_link}&status=1`" type="primary">连载中</el-link>
        <el-link :href="`${current_link}&status=2`" type="primary">全完结</el-link>
      </div>
    </el-col>
  </el-row>
  <el-row>
    <el-col :span="1">
      <div>上架:</div>
    </el-col>
    <el-col :span="23">
      <div>
        <el-link :href="`${current_link}`"  class="wenzi">全部</el-link>
        <el-link :href="`${current_link}&online_status=1`" type="primary">已上架</el-link>
        <el-link :href="`${current_link}&online_status=2`" type="primary">已下架</el-link>
      </div>
    </el-col>
  </el-row>
  <el-row>
    <el-col :span="1">
      <div>敏感:</div>
    </el-col>
    <el-col :span="23">
      <div>
        <el-link :href="`${current_link}`" class="wenzi">全部</el-link>
        <el-link :href="`${current_link}&is_sensitivity=1`" type="primary">不敏感</el-link>
        <el-link :href="`${current_link}&is_sensitivity=2`" type="primary">敏感</el-link>
      </div>
    </el-col>
  </el-row>
  <el-row>
    <el-col :span="1">
      <div>属性:</div>
    </el-col>
    <el-col :span="23">
      <div>
        <el-link :href="`${current_link}&sort=2`" class="wenzi">全部</el-link>
        <el-link :href="`${current_link}&flag=1`" type="primary">免费</el-link>
        <el-link :href="`${current_link}&flag=2`" type="primary">新书</el-link>
        <el-link :href="`${current_link}&flag=3`" type="primary">热门</el-link>
        <el-link :href="`${current_link}&flag=4`" type="primary">会员</el-link>
        <el-link :href="`${current_link}&flag=5`" type="primary">爽文</el-link>
        <el-link :href="`${current_link}&flag=6`" type="primary">精选</el-link>
        <el-link :href="`${current_link}&flag=7`" type="primary">大神</el-link>
      </div>
    </el-col>
  </el-row>
  <el-row>
    <el-col :span="1">
      <div>字数:</div>
    </el-col>
    <el-col :span="23">
      <div>
        <el-link :href="`${current_link}`" class="wenzi">全部</el-link>
        <el-link :href="`${current_link}&word=1`" type="primary">20万字以下</el-link>
        <el-link :href="`${current_link}&word=2`" type="primary">20-30万字</el-link>
        <el-link :href="`${current_link}&word=3`" type="primary">30-50万字</el-link>
        <el-link :href="`${current_link}&word=4`" type="primary">50-100万字</el-link>
        <el-link :href="`${current_link}&word=5`" type="primary">100-200万字</el-link>
        <el-link :href="`${current_link}&word=6`" type="primary">200字以上</el-link>
      </div>
    </el-col>
  </el-row>
  <el-row>
    <el-col :span="1">
      <div>排序:</div>
    </el-col>
    <el-col :span="23">
      <div>
        <el-link :href="`${current_link}`" class="wenzi" :underline="false">全部</el-link>
        <el-link :href="`${current_link}&sort=1`" type="primary">最新上架</el-link>
        <el-link :href="`${current_link}&sort=2`" type="primary">最近更新</el-link>
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
  },
  watch: {
    // 如果路由有变化，会再次执行该方法
    '$route': 'fetchData'
  },
  methods: {
    fetchData () {
      this.current_link =  document.location.href
    }
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
.wenzi {
  color: #fff;
  background-color: #32c5d2;
  border-color: #32c5d2;
  border-radius: 4px;
}
</style>