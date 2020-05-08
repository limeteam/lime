<template>
  <div class="login-container">
    <el-form ref="loginForm" :model="loginForm" :rules="loginRules" class="login-form" auto-complete="on" label-position="left">
      <div class="title-container">
        <span class="title">{{ $t('login.title') }}</span>
        <lang-select class="set-language" />
      </div>

      <el-form-item prop="username">
        <span class="svg-container">
          <svg-icon icon-class="user" />
        </span>
        <el-input
          v-model="loginForm.username"
          :placeholder="$t('login.username')"
          name="username"
          type="text"
          auto-complete="off"
          @keyup.native="staffCode(loginForm.username)"
        />
      </el-form-item>
      <el-form-item prop="password">
        <span class="svg-container">
          <svg-icon icon-class="password" />
        </span>
        <el-input
          v-model="loginForm.password"
          :type="passwordType"
          :placeholder="$t('login.password')"
          name="password"
          auto-complete="on"
          @keyup.enter.native="handleLogin"
        />
        <span class="show-pwd" @click="showPwd">
          <svg-icon icon-class="eye" />
        </span>
      </el-form-item>
      <el-button :loading="loading" type="primary" class="login-btn" style="" @click.native.prevent="handleLogin">
        {{ $t('login.logIn') }}
      </el-button>
    </el-form>
  </div>
</template>

<script>
import { isvalidUsername } from '@/utils/validate'
import LangSelect from '@/components/LangSelect'
/* import SocialSign from './socialsignin' */

export default {
  name: 'Login',
  components: { LangSelect/* , SocialSign  */ },
  data() {
    const validateUsername = (rule, value, callback) => {
      if (!isvalidUsername(value)) {
        callback(new Error('请输入用户名'))
      } else {
        callback()
      }
    }
    const validatePassword = (rule, value, callback) => {
      if (value.length < 6) {
        callback(new Error('密码不少于 6 位'))
      } else {
        callback()
      }
    }
    return {
      loginForm: {
        username: '',
        password: ''
      },
      loginRules: {
        username: [{ required: true, trigger: 'blur', validator: validateUsername }],
        password: [{ required: true, trigger: 'blur', validator: validatePassword }]
      },
      passwordType: 'password',
      loading: false,
      showDialog: false,
      redirect: undefined
    }
  },
  watch: {
    $route: {
      handler: function(route) {
        this.redirect = route.query && route.query.redirect
      },
      immediate: true
    }
  },
  mounted() {
    const username = localStorage.getItem('sys_username')
    if (username) {
      this.username = username
    }
  },
  created() {
    // window.addEventListener('hashchange', this.afterQRScan)
  },
  destroyed() {
    // window.removeEventListener('hashchange', this.afterQRScan)
  },
  methods: {
    staffCode(val) {
      this.loginForm.username = val.replace(/[\u4e00-\u9fa5]/g, '').trim()
    },
    // 显示密码
    showPwd() {
      if (this.passwordType === 'password') {
        this.passwordType = ''
      } else {
        this.passwordType = 'password'
      }
    },
    // 登录方法
    handleLogin() {
      this.$refs.loginForm.validate(valid => {
        if (valid) {
          this.loading = true
          this.$store.dispatch('LoginByUsername', this.loginForm).then(() => {
            // 登录成功
            this.loading = false
            this.$router.push({ path: this.redirect || '/' })
          }).catch(() => {
            this.loading = false
          })
        } else {
          console.log('error submit!!')
          return false
        }
      })
    }
  }
}
</script>

<style rel="stylesheet/scss" lang="scss">
  /* 修复input 背景不协调 和光标变色 */
  /* Detail see https://github.com/PanJiaChen/vue-element-admin/pull/927 */

  $bg:#283443;
  $bg: #f0f2f5;
  $cursor: #fff;

  @supports (-webkit-mask: none) and (not (cater-color: $cursor)) {
    .login-container .el-input input{
      &::first-line {
        background-color: #fff;
        color: #666;
      }
    }
  }
.login-container{

    /* reset element-ui css */
    .el-form-item__content{
      height: 40px;
      font-size: 16px;
      border-radius: 4px;
      background: transparent;
    }

    .el-input {
      display: inline-block;
      height: 40px;
      line-height: 40px;
      background: transparent;
      input {
        padding: 0 40px 0 40px;
        height: 40px;
        font-size: 16px;
        color: #333;
        border: 1px solid #ccc;
        background: transparent;
        &:-webkit-autofill {
          -webkit-box-shadow: 0 0 0px 1000px #fff inset !important;
          -webkit-text-fill-color: #333 !important;
        }
      }
      .el-input__inner:focus{
        border-color: #40a9ff;
        outline: 0;
        box-shadow: 0 0 0 2px rgba(24,144,255,.2);
        border-right-width: 1px!important;
      }
    }
    .el-form-item {
      position: relative;
      border: 0;
      background: #fff;
      border-radius: 5px;
      color: #454545;
    }
  }
</style>

<style rel="stylesheet/scss" lang="scss" scoped>
$bg: #2d3a4b;
$bg: #f0f2f5;
$dark_gray:#889aa4;
$light_gray:#eee;

.login-container {
  position: fixed;
  height: 100%;
  width: 100%;
  background-color: $bg;
  background-image: url(https://gw.alipayobjects.com/zos/rmsportal/TVYTbAXWheQpRcWDaDMu.svg);
    background-repeat: no-repeat;
    background-position: center 110px;
    background-size: 100%;
  .login-form {
    position: absolute;
    left: 0;
    right: 0;
    width: 450px;
    max-width: 100%;
    padding: 35px 35px 15px 35px;
    margin: 120px auto;

    .login-btn{
      width:100%;
      padding: 10px 0;
      font-size: 16px;
      margin-bottom: 25px;
    }
  }
  .tips {
    font-size: 14px;
    color: #fff;
    margin-bottom: 10px;
    span {
      &:first-of-type {
        margin-right: 16px;
      }
    }
  }
  .svg-container {
    position: absolute;
    z-index: 100;
    padding: 6px 5px 6px 15px;
    color: #bbb;
    vertical-align: middle;
    width: 30px;
    display: inline-block;
  }
  .title-container {
    position: relative;
    height: 44px;
    line-height:44px;
    text-align : center;
    margin-bottom :  20px;
    margin-right: 16px;
    margin-left: -25px;
    vertical-align: top;

    .logo{
      height: 44px;
      vertical-align: top;
      margin-right: 16px;
    }
    .title {
      font-size: 26px;
      color: rgba(0,0,0,.85);
      margin: 0px auto 40px auto;
      text-align: center;
      font-weight: bold;
    }
    .set-language {
      color: #999;
      position: absolute;
      top: 5px;
      right: 0px;
    }
  }
  .show-pwd {
    position: absolute;
    right: 10px;
    top: 7px;
    font-size: 16px;
    color: #333;
    cursor: pointer;
    user-select: none;
  }
  .thirdparty-button {
    position: absolute;
    right: 35px;
    bottom: 28px;
  }
}
</style>
