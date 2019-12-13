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
        <div style="margin: 40px 0px 0px 0px;min-height: 400px;padding:25px;
          background: linear-gradient(to right, rgba(255, 0, 0, 0.05), rgba(0, 0, 255, 0.06));">
          <h3 style="color: #7800ff;padding: 5px;"><Icon type="ios-paper"></Icon>账号注册协议</h3>
          <hr>
          <div style="font-size: 12px;font-family: Tahoma, Helvetica, 'Microsoft Yahei', 微软雅黑, Arial, STHeiti;">
            <p style="line-height: 30px;">
              在线网校注册条款
              一．服务条款的确认和接纳
              在线网校及其涉及到的产品、相关软件的所有权和运作权归在北京市海淀区航天信息培训学校（以下简称在线网校）所有， 在线网校享有对在线网校上一切活动的监督、提示、检查、纠正及处罚等权利。用户通过注册程序阅读本服务条款并点击"同意"按钮完成注册，即表示用户与在线网校已达成协议，自愿接受本服务条款的所有内容。如果用户不同意服务条款的条件，则不能获得使用服务以及注册成为在线网校用户的权利。
              二．服务保护条款
              1. 在线网校运用自己的操作系统通过国际互联网络为用户提供各项服务，用户必须:
              （1）提供设备，包括个人电脑一台、调制解调器一个及配备上网装置。
              （2）个人上网和支付与此服务有关的电话费用。
            </p>
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
