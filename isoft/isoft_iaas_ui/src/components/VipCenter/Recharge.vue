<template>
  <div class="isoft_bg_white isoft_pd10">
    <div>
      <a @click="prePage()" style="color: #ff6600"><Icon type="ios-undo-outline"  style="font-size: 40px"/>会员中心</a>
    </div>

    <div class="demo-split">
      <div class="demo-split-pane">
        <div >
          <img class="vipLogo" src="../../../static/images/vipCenter/vipLogo.png" height="41" width="41"/>
          <span class="payCenter">支付中心</span>
        </div>
        <div class="vipDesc">
          <Row>
            <a style="color: #ff6600">当前账号:  </a>{{loginUserName}}
          </Row><br>
          <Row>
            <a style="color: #ff6600">开通时长:  </a>{{openingTime}}&nbsp;&nbsp;&nbsp;
            <RadioGroup v-model="openingTime" type="button">
              <Radio label=" 1个月 / ¥60 " openingTime></Radio>
              <Radio label=" 3个月 / ¥150 " openingTime></Radio>
              <Radio label=" 6个月 / ¥240 " openingTime></Radio>
              <Radio label=" 1年 / ¥360 " openingTime></Radio>
            </RadioGroup>
          </Row><br>
          <Row>
            <a style="color: #ff6600">支付方式:  </a>
            <RadioGroup v-model="payType">
              <Radio label="微信支付" border></Radio>
            </RadioGroup>
          </Row><br>
          <Row>
            <a style="color: #ff6600">支付金额:  </a><a style="font-size: 20px ;color: #ff6900">{{payMoney}}</a>元
          </Row><br>
        </div>
        <div class="qrcode">
          (这里生成二维码)
        </div>
      </div>
    </div>
  </div>
</template>

<script>
  import {GetLoginUserName} from "../../tools"

	export default {
		name: "Recharge",
    data(){
      return {
        openingTime: ' 1个月 / ¥60 ',
        payType:'微信支付',
        payMoney:'60',
      }
    },
    computed:{
		  loginUserName:function () {
        return GetLoginUserName();
      }
    },
    methods:{
      prePage:function () {
        window.location.href = "#/vipcenter/vipIntroduction/";
      }
    },
	}
</script>

<style scoped>
  .vipLogo{
    background-position-x: center;
    background-position-y: center;
  }
  .vipDesc{
    float: right;width: 70%;
  }
  .payCenter{
    float: right;width: 55%; color: #ff6900;font-size: 25px
  }
  .qrcode{
    float: right;
    width: 55%;
  }
</style>
