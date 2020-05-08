<template>
  <div class="block">
    <el-pagination
      :background="true"
      :current-page.sync="currentPage"
      :page-sizes="pageSizes"
      :page-size="pageSize"
      :total="totalList"
      layout="total, sizes, prev, pager, next, jumper"
      @size-change="handleSizeChange"
      @current-change="handleCurrentChange"
    />
  </div>
</template>

<script>
/*
使用办法：
<div v-show="isPage" class="bottom">
        <Paging v-bind="pages" class="paging" @passpagesize="passPageSize"/>
      </div>
      pages: {// 分页
        getPageSizes: 10, // 有多少个分页
        getPageSize: 10, // 每个页多少条
        getTotalList: 0,
        currentPage: 1
      },
      isPage: false
      passPageSize: function(payload) {
      const w = payload.Sizejudge
      if (w === 0) {
        this.formData.page = 1
        this.formData.size = payload.sizeNum
        this.getHouseList()
      } else {
        this.formData.page = payload.sizeNum
        this.getHouseList()
      }
    }
*/
export default {
  name: 'Paging',
  props: {
    getPageSizes: {
      type: Number,
      default: 10
    },
    getPageSize: {
      type: Number,
      default: 10
    },
    getTotalList: {
      type: Number,
      default: 0
    },
    getCurrentPage: {
      type: Number,
      default: 1
    }
  },
  data() {
    return {
      currentPage: this.getCurrentPage,
      pageSize: this.getPageSize
    }
  },
  computed: {
    totalList: function() {
      return parseInt(this.getTotalList)
    },
    pageSizes: function() {
      const all = parseInt(this.getTotalList)
      const one = parseInt(this.getPageSizes)
      const n = parseInt(all / one)
      let result = []
      if (n > 0) {
        for (let i = 1; i <= n; i++) {
          result.push(one * i)
        }
        return result
      } else {
        result = [one]
        return result
      }
    }
  },
  methods: {
    handleCurrentChange(val) { // 当前第几页
      this.$emit('passpagesize', { sizeNum: val, Sizejudge: 1 })
    },
    handleSizeChange(val) { // 每页多少条
      this.$emit('passpagesize', { sizeNum: val, Sizejudge: 0 })
    }
  }
  // })
}
</script>

