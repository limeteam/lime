<template>
  <p class="p-title">
    {{ title }}
  </p>
</template>
<script>
export default {
  props: {
    title: {
      type: String,
      default: null
    }
  }
}
</script>
<style>
.p-title{
      height: 30px;
    background: #f5f7fa;
    color: #909399;
    line-height: 30px;
    padding-left: 10px;
    font-size: 14px;
}
</style>
