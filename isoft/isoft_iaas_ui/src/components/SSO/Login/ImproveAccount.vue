<template>
  <div id="nav">
    依据《网络安全法》，为保障您的账户安全和正常使用，请尽快完善账号信息！新版本全面为您的账号安全护航！
    <a href="javascript:;" id="account" @click="$router.push({path:'/sso/regist', query: { pattern: 1 }})">完善账号资料</a>
  </div>
</template>

<script>
  import IBeautifulLink from "../../Common/link/IBeautifulLink";
  export default {
    name: "ImproveAccount",
    components: {IBeautifulLink}
  }
</script>

<style scoped>
 #nav{
   font-size: 12px;
   text-align: center;
   margin-bottom: 0px;
   color: #8a6d3b;
   background-color: #fcf8e3;
   border-color: #faebcc;
   padding: 15px;
   border: 1px solid transparent;
   border-radius: 4px;
   box-sizing: border-box;
 }
  #account{
    color: #66512c;
    font-weight: bold;
    text-decoration: underline;
    background-color: transparent;
  }
</style>
