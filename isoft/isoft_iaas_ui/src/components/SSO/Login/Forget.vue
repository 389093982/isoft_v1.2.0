<template>
  <div style="margin:50px 200px;">
    <div style="text-align: center;padding-left: 50px;">
      <span style="height: 60px;line-height: 60px;font-size: 16px;color: #000;">修改账号</span>
      <span>
        <a href="/#/sso/login/" style="font-size: 15px;font-weight: inherit;">已有账号,前去登录</a>
      </span>
    </div>
    <Form ref="formValidate" :model="formValidate" :rules="ruleValidate" :label-width="80">
      <FormItem label="用户名" prop="username">
        <Input v-model.trim="formValidate.username" placeholder="请输入注册邮箱"></Input>
      </FormItem>
      <FormItem label="验证码" prop="verifycode">
        <Input v-model.trim="formValidate.verifycode" placeholder="请输入验证码"></Input>
      </FormItem>
      <FormItem label="修改密码" prop="passwd">
        <Input v-model.trim="formValidate.passwd" type="password" placeholder="请输入密码"></Input>
      </FormItem>
      <FormItem label="确认密码" prop="repasswd">
        <Input v-model.trim="formValidate.repasswd" type="password" placeholder="请输入确认密码"></Input>
      </FormItem>
      <FormItem>
        <Row :gutter="10">
          <Col span="12">
            <Button type="primary" @click="getVerifyCode('formValidate')" :disabled="VerDisableFlag" size="large">{{VerifyCodeButtonDesc}}</Button>
          </Col>
          <Col span="12">
            <Button type="primary" @click="handleSubmit('formValidate')" size="large">&nbsp;&nbsp;&nbsp;&nbsp;提交&nbsp;&nbsp;&nbsp;&nbsp;</Button>
          </Col>
        </Row>
      </FormItem>
    </Form>
  </div>
</template>

<script>
  import {validateEmail} from "../../../tools"
  import {CreateVerifyCode,ModifyPwd} from "../../../api"

  export default {
    name: "Forget",
    data(){
      const _validateUserName = (rule, value, callback) => {
        if (value === '') {
          callback(new Error('用户名不能为空!'));
        } else if (!validateEmail(value)) {
          callback(new Error('邮箱不合法!'));
        } else {
          callback();
        }
      };
      const checkEmptyValidator = function (emptyMsg) {
        return (rule, value, callback) => {
          if (value === '') {
            callback(new Error(emptyMsg));
          } else {
            callback();
          }
        };
      };
      // 确认密码校验 validatePassCheck
      const validatePassCheck = (rule, value, callback) =>  {
        if (value === '') {
          callback(new Error('确认密码不能为空!'));
        } else if (value !== this.formValidate.passwd) {
          callback(new Error('两次输入密码不一致!'));
        } else {
          callback();
        }
      };
      return {
        VerDisableFlag:false,
        totalTime:30,
        VerifyCodeButtonDesc:'获取验证码',
        formValidate: {
          username: '',
          verifycode: '',
          passwd: '',
          repasswd: '',
        },
        ruleValidate: {
          username: [
            { required: true, validator: _validateUserName,  trigger: 'blur' }
          ],
          verifycode: [
            { required: true, validator: checkEmptyValidator("验证码不能为空!"),  trigger: 'blur' }
          ],
          passwd: [
            { required: true, validator: checkEmptyValidator("密码不能为空!"), trigger: 'blur' },
          ],
          repasswd: [        // 确认密码校验 validatePassCheck
            { required: true, validator: validatePassCheck, trigger: 'blur' }
          ],
        }
      }
    },
    methods:{
      getVerifyCode:function(name){
        this.$refs[name].validateField('username', async (err) => {
          if (!err) {
            // 校验通过则进行注册
            this.createVerifyCode(this.formValidate.username);
          } else {
            this.$Message.error('信息校验失败!');
          }
        });
      },
      handleSubmit: function (name) {
        this.$refs[name].validate(async (valid) => {
          if (valid) {
            // 校验通过则进行注册
            this.modifyPwd();
          } else {
            this.$Message.error('信息校验失败!');
          }
        })
      },
      createVerifyCode:async function (username) {
        const result = await CreateVerifyCode(username);
        if(result.status == "SUCCESS"){
          this.$Message.success("验证码发送成功,请注意查收!");
          //这里进行30秒的置灰设置
          this.VerDisableFlag = true;
          this.VerifyCodeButtonDesc = this.totalTime + 's后重新获取';//展示30
          let clock = window.setInterval(() => {
            this.totalTime--;
            this.VerifyCodeButtonDesc = this.totalTime + 's后重新获取';
            if (this.totalTime < 0) {//当倒计时小于0时清除定时器
              window.clearInterval(clock);
              this.VerifyCodeButtonDesc = '获取验证码';
              this.VerDisableFlag = false;
              this.totalTime = 30
            }
          },1000);
        }else{
          this.$Message.error(result.errorMsg);
        }
      },
      modifyPwd:async function () {
        var _this = this;
        const result = await ModifyPwd(this.formValidate.username, this.formValidate.passwd, this.formValidate.verifycode);
        if(result.status == "SUCCESS"){
          this.$Message.success("密码修改成功!");
          // 注册成功延迟 2s 跳往登录页面
          setTimeout(function(){
            _this.$router.push({path:'/sso/login'});
          },2000);
        }else{
          this.$Message.error(result.errorMsg);
        }
      }
    }
  }
</script>

<style scoped>
  .submitBtn{
    width: 100%;height: 40px;display: block;line-height: 40px;
    font-size: 16px;font-weight: 800;cursor: pointer;color: #fff;
    background: #3f89ec;border: 0;text-align: center;
  }
</style>
