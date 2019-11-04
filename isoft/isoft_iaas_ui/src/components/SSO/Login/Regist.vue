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
              2. 考虑到在线网校产品服务的重要性，用户同意:
              （1）提供及时、详尽及准确的个人资料。
              （2）不断更新注册资料，符合及时、详尽准确的要求。所有原始键入的资料将引用为注册资料。
              3. 用户可授权在线网校向第三方透露其注册资料，否则在线网校不能公开用户的姓名、住址、出件地址、电子邮箱、帐号。除非:
              （1）事先获得用户明确授权后，用户要求在线网校或授权某人通过电子邮件服务或其他方式透露这些信息。
              （2）相应的法律、法规要求以及按照有关政府主管部门的要求，需要在线网校提供用户的个人资料。
              （3）为了维护公众以及在线网校合法利益。
              （4）在线网校可能会与第三方合作向用户提供相关的网络服务，在此情况下，如该第三方同意承担与在线网校同等的保护用户隐私的责任，则在线网校有权将用户的注册资料等提供给该第三方。
              （5）在不透露单个用户隐私资料的前提下，在线网校有权对整个用户数据库进行分析并对用户数据库进行商业上的利用。
              4. 如果用户提供的资料不准确，不真实，不合法有效，在线网校保留结束用户使用在线网校各项服务的权利。
              用户在享用在线网校各项服务的同时，同意接受在线网校提供的各类信息服务。
              5. 在线网校定义的信息内容包括:文字、软件、声音、相片、录像、图表；在广告中全部内容；在线网校为用户提供的商业信息等，所有这些内容受版权、商标权、和其它知识产权及所有权法律的保护。所以，用户只能在在线网校授权下才能使用这些内容，而不能擅自复制、修改、编撰这些内容、或创造与内容有关的衍生产品。
              6. 如果用户未遵守本服务条款的任何一项，在线网校有权利即终止提供一切服务，并保留通过法律手段追究责任的权利。
              7. 使用在线网校提供的服务由用户自己承担风险，在适用法律允许的最大范围内，在线网校在任何情况下不就因使用或不能使用在线网校提供的服务所发生的特殊的、意外的、直接或间接的损失承担赔偿责任。即使已事先被告知该损害发生的可能性。
              8. 用户须明白，使用在线网校提供的服务涉及到Internet服务和电信增值服务，可能会受到各个环节不稳定因素的影响。因此服务存在因不可抗力、计算机病毒或黑客攻击、国家相关行业主管部门及电信运营商的调整、系统不稳定、用户所在位置、用户关机以及其他任何技术、互联网络、通信线路原因等造成的服务中断或不能满足用户要求的风险。用户须承担以上风险，在线网校不作担保。对因此导致用户不能发送、上传和接受阅读消息、或接发错消息，或无法实现其他通讯条件，在线网校不承担任何责任。
              9. 用户须明白，在使用在线网校提供的服务存在有来自任何他人的包括威胁性的、诽谤性的、令人反感的或非法的内容或行为或对他人权利的侵犯（包括知识产权）的匿名或冒名的信息的风险，用户须承担以上风险，在线网校和合作公司对服务不作任何类型的担保，不论是明确的或隐含的，包括所有有关信息真实性、适用性、所有权和非侵权性的默示担保和条件，对因此导致任何因用户不正当或非法使用服务产生的直接、间接、偶然、特殊及后续的损害，不承担任何责任。
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
  import {validatePasswd} from "../../../tools"
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
