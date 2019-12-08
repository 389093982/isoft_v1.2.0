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
        <div style="margin:80px 200px;">
          <div style="text-align: center;padding-left: 50px;">
            <span style="height: 60px;line-height: 60px;font-size: 16px;color: #000;">用户注册</span>
            <span>
            <a href="/#/sso/login/" style="font-size: 15px;font-weight: inherit;">已有账号,前去登录</a>
          </span>
          </div>
          <Form ref="formValidate" :model="formValidate" :rules="ruleValidate" :label-width="80">
            <FormItem label="用户名" prop="username">
              <Input v-model.trim="formValidate.username" placeholder="请输入用户名"></Input>
            </FormItem>
            <FormItem label="密码" prop="passwd">
              <Input v-model.trim="formValidate.passwd" type="password" placeholder="请输入密码"></Input>
            </FormItem>
            <FormItem label="确认密码" prop="repasswd">
              <Input v-model.trim="formValidate.repasswd" type="password" placeholder="请输入确认密码"></Input>
            </FormItem>
            <FormItem label="用户协议" prop="protocol">
              <CheckboxGroup v-model="formValidate.protocol">
                <Checkbox label="用户协议">
                  <label>阅读并接受</label>《用户协议》
                </Checkbox>
              </CheckboxGroup>
            </FormItem>
            <FormItem>
              <div @click="handleSubmit('formValidate')" style="width: 100%;height: 40px;display: block;line-height: 40px;
                font-size: 16px;font-weight: 800;cursor: pointer;color: #fff;background: #3f89ec;border: 0;text-align: center;">注册</div>
            </FormItem>
          </Form>
        </div>
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
  import {Regist} from "../../../api"
  import {validateUserName} from "../../../tools"
  import ISimpleSearch from "../../Common/search/ISimpleSearch"
  import IBeautifulLink from "../../Common/link/IBeautifulLink";

  export default {
    name: "Regist",
    components:{IBeautifulLink, LoginFooter,ISimpleSearch},
    data(){
      const _validateUserName = (rule, value, callback) => {
        if (value === '') {
          callback(new Error('用户名不能为空!'));
        } else if (!validateUserName(value)) {
          callback(new Error('6至20位，以字母开头，字母，数字，减号，下划线!'));
        } else {
          callback();
        }
      };
      const _validatePasswd = (rule, value, callback) => {
        if (value === '') {
          callback(new Error('密码不能为空!'));
        } else {
          callback();
        }
      };
      // 确认密码校验 validatePassCheck
      const validatePassCheck = (rule, value, callback) =>  {
        if (value === '') {
          callback(new Error('请输入密码进行确认!'));
        } else if (value !== this.formValidate.passwd) {
          callback(new Error('两次输入密码不一致!'));
        } else {
          callback();
        }
      };
      return {
        formValidate: {
          username: '',
          passwd: '',
          repasswd: '',
          protocol:[],
        },
        ruleValidate: {
          username: [
            { required: true, validator: _validateUserName, trigger: 'blur' }
          ],
          passwd: [
            { required: true, validator: _validatePasswd, trigger: 'blur' },
          ],
          repasswd: [        // 确认密码校验 validatePassCheck
            { required: true, validator: validatePassCheck, trigger: 'blur' }
          ],
          protocol: [
            { required: true, type: 'array', min: 1, message: '用户协议必须同意!', trigger: 'change' },
          ],
        }
      }
    },
    methods:{
      handleSubmit: function (name) {
        this.$refs[name].validate(async (valid) => {
          if (valid) {
            // 校验通过则进行注册
            this.regist();
          } else {
            this.$Message.error('注册信息校验失败!');
          }
        })
      },
      regist:async function () {
        var username = this.formValidate.username;
        var passwd = this.formValidate.passwd;
        const result = await Regist(username,passwd);
        if(result.status=="SUCCESS"){
          this.$Message.success('注册成功!');
          this.$router.push({path:'/sso/login'});
        }else{
          if(result.errorMsg == "regist_exist"){
            this.$Message.error("该用户已经被注册!");
          }else if(result.errorMsg == "regist_failed"){
            this.$Message.error("注册失败,请联系管理员获取账号!");
          }else {
            this.$Message.error(result.errorMsg);
          }
        }
      },
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
