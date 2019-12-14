<template>
  <div>
    <div id="header">
      <div class="login-link" id="login_link" style="position: absolute;top: 25px;right: 100px;">
        <span style="font-size: 15px;font-weight: inherit;">已有账号,前去<a href="/#/sso/login/">登录</a></span>
      </div>
    </div>
    <div style="height: 20px;display: block;width:100%;background: linear-gradient(to right, red, blue);opacity:0.1;">
    </div>

    <Row>
      <Col span="14">
        <Forget v-if="$route.query.pattern && $route.query.pattern == 2"/>
        <Regist v-else/>
      </Col>
      <Col span="8">
        <div style="margin: 40px 0px;min-height: 400px;padding:25px;
          background: linear-gradient(to right, rgba(255, 0, 0, 0.05), rgba(0, 0, 255, 0.06));">
          <h3 style="color: #7800ff;padding: 5px;"><Icon type="ios-paper"></Icon>账号注册协议</h3>
          <hr>
          <div style="font-size: 12px;font-family: Tahoma, Helvetica, 'Microsoft Yahei', 微软雅黑, Arial, STHeiti;">
            <p style="line-height: 30px;">1、为简化注册流程,注册不需要用户接收验证码校验,请自行保证邮箱信息正确</p>
            <p style="line-height: 30px;">2、账号注册需要使用真实邮箱进行注册,禁止冒用无效邮箱、他人邮箱注册,以防账号存在安全问题</p>
            <p style="line-height: 30px;">3、修改用户昵称或者密码时需要先获取验证码,验证码有效时间 5 min</p>
            <p style="line-height: 30px;">4、新用户注册账号赠送 7 天 VIP 会员体验资格</p>
            <p style="line-height: 30px;">5、需要继续完善咯...</p>
          </div>
        </div>
      </Col>
    </Row>
    <LoginFooter/>
  </div>
</template>

<script>
  import LoginFooter from "./LoginFooter"
  import Regist from "./Regist";
  import Forget from "./Forget";

  export default {
    name: "RegistOrForget",
    components:{Regist, Forget, LoginFooter},
    data(){
      return {}
    },
    methods:{

    }
  }
</script>

<style scoped>
  #header {
    background-color: rgb(255, 255, 255);
    text-align:center;
    height:70px;
    padding:5px;
  }
  a:hover {
    color: #E4393C;
    text-decoration: underline;
  }
</style>
