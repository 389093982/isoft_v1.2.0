<template>
  <div id="login_area">
    <div id="login_header" style="float: right;width: 94%">
      <h4>
        密码登录
        <a style="float: right;width: 30%;color:#2E82FF" href="/sso/login">访问首页</a>
      </h4>
    </div>

    <div id="login_content" style="margin: 0 auto;padding:20px;">
      <input class="focus" name="username" placeholder="请输入用户名" type="text" style="width: 100%;height: 40px;" required/>
      <input type="password" style="display:none">
      <input class="focus" name="passwd" placeholder="请输入密码" type="password"
             style="width: 100%;height: 40px;margin-top:20px;" autocomplete="new-password" required/>
      <span id="error_msg" v-if="showError">{{errorMsg}}</span>
      <p>
        <input id="submit" type="submit" value="登录" @click="login">
        <router-link :to="{path:'/sso/forget', query: { pattern: 2 }}" style="float: right;color: #2e82ff;">忘记密码？</router-link>
      </p>
    </div>
    <div id="login_footer">
      <router-link :to="{path:'/sso/user/loginProblem'}" style="float: left;color: #2e82ff;">常见登录问题</router-link>
      <router-link :to="{path:'/sso/regist', query: { pattern: 1 }}" style="float: right;color: #2e82ff;">立即注册</router-link>
    </div>
  </div>
</template>

<script>
  import {Login} from "../../../api"
  import {setCookie} from "../../../tools"
  import {strSplit} from "../../../tools"

  export default {
    name: "LoginForm",
    data(){
      return {
        showError:false,
        errorMsg:"登录失败!",
      }
    },
    methods:{
      login:async function () {
        let redirectUrl="";
        var arr = strSplit(window.location.href, "?redirectUrl=");
        if (arr.length == 2){
          redirectUrl = arr[1];
        }
        var username = $("input[name='username']").val();
        var passwd = $("input[name='passwd']").val();
        var result = await Login(username, passwd, decodeURIComponent(redirectUrl));
        if(result.loginSuccess == true || result.loginSuccess == "SUCCESS"){
          setCookie("tokenString",result.tokenString,365,result.domain);
          setCookie("userName",username,365,result.domain);
          setCookie("nickName",result.nickName,365,result.domain);
          setCookie("isLogin","isLogin",365,result.domain);
          setCookie("roleName",result.roleName,365,result.domain);
          let expireSecond = new Date().getTime() + result.expireSecond * 1000;     // 时间戳
          setCookie("expireSecond",expireSecond,365,result.domain);
          // 跳往需要跳转的页面,并设置cookie
          window.location.href = result.redirectUrl;
        }else{
          this.showError = true;
          this.errorMsg = result.errorMsg;
        }
      }
    }
  }
</script>

<style scoped>
  #login_area{
    width: 320px;
    height: 450px;
    background: #ffffff;
    float: right;
    margin-top: 20px;
    margin-right: 150px;
    position: relative;
  }
  #login_header{
    height: 80px;
    line-height: 80px;
    text-align: left;
    font-size: 16px;
    color: #000;
  }
  #submit{
    width: 100%;
    height: 40px;
    margin-top:30px;
    margin-bottom:20px;
    display: block;
    line-height: 40px;
    font-size: 16px;
    font-weight: 800;
    cursor: pointer;
    color: #fff;
    background: #3f89ec;
    border: 0;
  }
  #login_footer{
    color: #2e82ff;
    position:absolute;
    bottom:0;
    width:100%;
    height: 60px;
    line-height: 60px;
    background: #f0f6ff;
    padding: 0 28px;
  }
  #error_msg{
    font-size: 12px;
    color:red;
    float: right;
    margin-top: 5px;
  }
  a {
    color: #666;
    text-decoration: none;
  }
  a:hover {
    color: #E4393C;
    text-decoration: underline;
  }
  .focus:focus {
    background-color: #ffffff;
    border-color: #2c5bff;
  }
</style>
