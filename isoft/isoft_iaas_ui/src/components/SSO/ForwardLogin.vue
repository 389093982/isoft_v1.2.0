<template>
  <span>
    您还未登录奥,前去 <IBeautifulLink @onclick="cancelUser">登录</IBeautifulLink>
  </span>
</template>

<script>
  import IBeautifulLink from "../Common/link/IBeautifulLink";
  import {LoginAddr} from "../../api"
  import {delCookie} from '../../tools/index'

  export default {
    name: "ForwardLogin",
    components: {IBeautifulLink},
    methods: {
      cancelUser() {
        delCookie("tokenString");
        delCookie("userName");
        delCookie("isLogin");
        window.location.href = LoginAddr + "?redirectUrl=" + window.location.href;
      }
    }
  }
</script>

<style scoped>

</style>
