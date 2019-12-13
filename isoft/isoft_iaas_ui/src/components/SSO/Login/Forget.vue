<template>
  <div style="margin:50px 200px;">
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
            <div @click="getVerifyCode('formValidate')" class="submitBtn">获取验证码</div>
          </Col>
          <Col span="12">
            <div @click="handleSubmit('formValidate')" class="submitBtn">提交</div>
          </Col>
        </Row>
      </FormItem>
    </Form>
  </div>
</template>

<script>
  import {validateEmail} from "../../../tools"

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
            this.createVerifyCode();
          } else {
            this.$Message.error('信息校验失败!');
          }
        });
      },
      handleSubmit: function (name) {
        this.$refs[name].validate(async (valid) => {
          alert(valid);
          if (valid) {
            // 校验通过则进行注册
            this.forgetPwd();
          } else {
            this.$Message.error('信息校验失败!');
          }
        })
      },
      createVerifyCode:function () {
        alert(11111);
      },
      forgetPwd:function () {
        alert(11111);
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
