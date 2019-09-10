<template>
  <span>
    您还未登录奥,前去 <IBeautifulLink2 @onclick="cancelUser">登录</IBeautifulLink2>
  </span>
</template>

<script>
  import IBeautifulLink2 from "../Common/link/IBeautifulLink2";
  import {LoginAddr} from "../../api"
  import {delCookie} from '../../tools/index'

  export default {
    name: "ForwardLogin",
    components: {IBeautifulLink2},
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
