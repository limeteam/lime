<template>
  <div class="panel">
    <div class="panel-title">
      <h1 class="panel-title__text">
        {{ title }}
      </h1><slot name="title" />
    </div>
    <el-main class="panel-contain">
      <slot />
    </el-main>
  </div>
</template>

<script>
export default {
  name: 'Panel',
  props: {
    title: {
      type: String,
      default: '标题'
    }
  }
}
</script>

<style lang="scss" scoped>
  .panel {
    margin-bottom: 20px;

    &:last-of-type{
      margin-bottom: 0;
    }

    &-title {
      height: 44px;
      background-color: rgba(250, 250, 250, 1);
      border-width: 1px;
      border-style: solid;
      border-color: rgba(233, 233, 233, 1);
      border-radius: 3px 3px 0 0;
      padding-left: 24px;
      margin: 0;
      display: flex;
      align-items: center;

      &__text{
        font-weight: 700;
        font-style: normal;
        color: #333333;
        font-size: 14px;
        text-align: left;
        line-height: 44px;
      }
    }

    &-contain {
      background: #fff;
      border-radius: 0 0 3px 3px;
    }
  }
</style>
