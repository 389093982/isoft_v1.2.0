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
        <div style="margin:80px;margin-left:200px;margin-right: 200px;">
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
                  <label>阅读并接受</label><a href="#" @click="showUserProtocol">《Isoft用户协议》</a>
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
          <h3 style="color: #7800ff;padding: 5px;"><Icon type="ios-paper"></Icon>账号特权</h3>
          <hr>
          <div style="font-size: 12px;font-family: Tahoma, Helvetica, 'Microsoft Yahei', 微软雅黑, Arial, STHeiti;">
            <p style="line-height: 30px;">1、初次注册账号送30小时免费学习时间</p>
            <p style="line-height: 30px;">2、初次注册账号送3000积分</p>
            <p style="line-height: 30px;">3、初次注册账号送云笔记使用特权</p>
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
  import {validatePasswd} from "../../../tools"
  import ISimpleSearch from "../../Common/search/ISimpleSearch"
  import IBeautifulLink2 from "../../Common/link/IBeautifulLink2";

  export default {
    name: "Regist",
    components:{IBeautifulLink2, LoginFooter,ISimpleSearch},
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
        } else if (!validatePasswd(value)) {
          callback(new Error('最少6位，至少1个大小写字母，数字和特殊字符!'));
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
            { validator: _validateUserName, trigger: 'blur' }
          ],
          passwd: [
            { validator: _validatePasswd, trigger: 'blur' },
          ],
          repasswd: [        // 确认密码校验 validatePassCheck
            { validator: validatePassCheck, trigger: 'blur' }
          ],
          protocol: [
            { required: true, type: 'array', min: 1, message: '用户协议必须同意!', trigger: 'change' },
          ],
        }
      }
    },
    methods:{
      showUserProtocol:function(){
        this.$Modal.info({
          title: '用户协议',
          content: this.GLOBAL.user_protocol,
        });
      },
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
