<template>
  <div class="isoft_bg_white isoft_pd10">
    <div class="vipCenterUp">
      <div style="font-size: 30px;padding: 70px 20px 30px 38%">
        <span style="color: #ff6600">会员尊享，学而不止</span>
      </div>
      <div style="font-size: 30px;padding: 70px 54% 30px 46%">
        <Button shape="circle" size="large" @click="toRecharge()" style="color: #ff6600;font-size: 20px;">
          开通会员
        </Button>
      </div>
    </div>
    <div>
      <Table stripe :columns="columns1" :data="data1"></Table>
    </div>
    <div class="vipCenterDown">
      <div class="demo-split">
        <Split v-model="split1">
          <div slot="left" class="demo-split-pane">
            <a><p style="color: #ff6600">非会员简介:</p></a>
          </div>
          <div slot="right" class="demo-split-pane">
            <a><p style="color: #ff6600;font-size: 15px">会员介绍:</p></a>
          </div>
        </Split>
      </div>
    </div>
  </div>
</template>
<script>
  import {CheckHasLoginConfirmDialog} from "../../tools"

  export default {
    name: "VipIntroduction",
    data () {
      return {
        split1: 0.5,
        columns1: [
          {
            title: '非会员',
            key: 'notVip'
          },
          {
            title: '会员',
            key: 'isVip'
          },
          {
            title: '尊享特权',
            key: 'privilege'
          }
        ],
        data1: [
          {
            notVip: '学习时间有限',
            isVip: '学习时间不限',
            privilege: '随时随地都可以登录学习',
          },
          {
            notVip: '视频数量20，长度30分钟',
            isVip: '视频数量200，长度120分钟',
            privilege: '上传视频自己开课',
          },
          {
            notVip: '视频数量20，长度30分钟',
            isVip: '视频数量200，长度120分钟',
            privilege: '上传视频自己开课',
          },

        ]
      }
    },
    methods:{
      toRecharge:function () {
        CheckHasLoginConfirmDialog(this, "/vipcenter/recharge/");
      }
    },
  }
</script>
<style>
  .vipCenterUp{
    width: 100%;
    height:246px;
    background-image:url("../../../static/images/vipCenter/vipCenterImg.png");
    background-position-x: center;
    background-position-y: center;
    margin-bottom: 10px;
  }
  .vipCenterDown{
    width: 100%;
    height:500px;
    background-position-x: center;
    background-position-y: center;
  }
  .demo-split{
    height: 500px;
    border: 1px solid #dcdee2;
  }
  .demo-split-pane{
    padding: 10px;
  }
</style>
