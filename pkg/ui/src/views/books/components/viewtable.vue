<template>
  <table class="viewTable" :style="styleObject" v-if="s_showByRow">
    <tr v-for="index in rowCount">
      <td class="column">{{tableData[index*2-2].key}}</td>
      <td>{{tableData[index*2-2].value}}</td>
      <td class="column">{{tableData[index*2-1] !== undefined ? tableData[index*2-1].key : ''}}</td>
      <td>{{tableData[index*2-1] !== undefined ? tableData[index*2-1].value : ''}}</td>
    </tr>
  </table>
  <table class="viewTable" :style="styleObject" v-else>
    <tr v-for="index in rowCount">
      <td class="column">{{tableData[index-1].key}}</td>
      <td>{{tableData[index-1].value}}</td>
      <td
        class="column"
      >{{tableData[rowCount+index-1] !== undefined ? tableData[rowCount+index-1].key : ''}}</td>
      <td>{{tableData[rowCount+index-1] !== undefined ? tableData[rowCount+index-1].value : ''}}</td>
    </tr>
  </table>
</template>

<script>
export default {
  data() {
    return {
      styleObject: {},
      s_showByRow: true,
    };
  },
  props: ["tableData", "tableStyle", "showByRow"],
  computed: {
    rowCount: function() {
      return Math.ceil(this.tableData.length / 2);
    }
  },
  created() {
    this.styleObject = this.tableStyle;
    if (this.showByRow !== undefined) {
      this.s_showByRow = this.showByRow;
    }
  }
};
</script>


<style>
　　.viewTable,
.viewTable tr,
.viewTable tr td {
  border: 1px solid #e6eaee;
}
　　.viewTable {
  font-size: 12px;
  color: #71787e;
}
　　.viewTable tr td {
  border: 1px solid #e6eaee;
  width: 150px;
  height: 35px;
  line-height: 35px;
  box-sizing: border-box;
  padding: 0 10px;
}
　　.viewTable tr td.column {
  background-color: #eff3f6;
  color: #393c3e;
}
</style>