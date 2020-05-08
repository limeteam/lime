<template>
  <div class="info-bar">
    <p>
      <el-icon class="info-bar__icon el-icon-warning" />
      <slot name="prependText" /><span class="info-bar__number">{{ total }}</span><slot name="appendText" />
    </p>
  </div>
</template>

<script>
export default {
  name: 'InfoBar',
  props: {
    total: { type: Number, default: 0 }
  }
}
</script>

<style lang="scss" scoped>
  .info-bar{
    background: inherit;
    background-color: #e6f7ff;
    -webkit-box-sizing: border-box;
    box-sizing: border-box;
    border: 1px solid #bae7ff;
    border-radius: 4px;
    line-height: 40px;
    padding: 0 15px;
    -webkit-box-shadow: none;
    box-shadow: none;
    font-family: "\5FAE\8F6F\96C5\9ED1 Regular","\5FAE\8F6F\96C5\9ED1";
    font-weight: 400;
    font-style: normal;
    font-size: 14px;
    margin: 16px;

    p{margin: 0}

    &__icon{
      color: blue;
      margin-right: 8px;
    }

    &__number{
      margin: 0 3px;
      color: red;
    }
  }
</style>
