<template>
  <div style="margin:50px 200px;">
    <div style="text-align: center;padding-left: 50px;">
      <span style="height: 60px;line-height: 60px;font-size: 16px;color: #000;">用户注册</span>
      <span>
        <a href="/#/sso/login/" style="font-size: 15px;font-weight: inherit;">已有账号,前去登录</a>
      </span>
    </div>
    <Form ref="formValidate" :model="formValidate" :rules="ruleValidate" :label-width="80">
      <FormItem label="用户名" prop="username">
        <Input v-model.trim="formValidate.username" placeholder="请输入注册邮箱"></Input>
      </FormItem>
      <FormItem label="用户昵称" prop="nickname">
        <Input v-model.trim="formValidate.nickname" placeholder="请输入用户昵称"></Input>
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
        <div @click="handleSubmit('formValidate')" class="submitBtn">注册</div>
      </FormItem>
    </Form>
  </div>
</template>

<script>
  import {validateEmail} from "../../../tools"
  import {Regist} from "../../../api"

  export default {
    name: "Regist",
    components:{},
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
          nickname: '',
          passwd: '',
          repasswd: '',
          protocol:[],
        },
        ruleValidate: {
          username: [
            { required: true, validator: _validateUserName,  trigger: 'blur' }
          ],
          nickname: [
            { required: true, validator: checkEmptyValidator("用户昵称不能为空!"), trigger: 'blur' }
          ],
          passwd: [
            { required: true, validator: checkEmptyValidator("密码不能为空!"), trigger: 'blur' },
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
        var _this = this;
        const result = await Regist(this.formValidate.username,this.formValidate.passwd,this.formValidate.nickname);
        if(result.status=="SUCCESS"){
          this.$Message.success('注册成功!');
          // 注册成功延迟 2s 跳往登录页面
          setTimeout(function(){
            _this.$router.push({path:'/sso/login'});
          },2000);

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
  a:hover {
    color: #E4393C;
    text-decoration: underline;
  }
  .submitBtn{
    width: 100%;height: 40px;display: block;line-height: 40px;
    font-size: 16px;font-weight: 800;cursor: pointer;color: #fff;
    background: #3f89ec;border: 0;text-align: center;
  }
</style>
