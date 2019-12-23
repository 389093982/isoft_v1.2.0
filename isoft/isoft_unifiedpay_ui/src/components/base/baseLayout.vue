<template>
  <div class="layout">
    <Layout>
      <Header>
        <Menu mode="horizontal" theme="dark" active-name="1">
          <!--<div class="layout-logo" style="font-size: 20px">&nbsp;&nbsp;领&nbsp;&nbsp;创&nbsp;&nbsp;科&nbsp;&nbsp;技</div>-->
          <div class="layout-nav">
            <MenuItem name="1">
              <Icon type="logo-yen" />
              微信
            </MenuItem>
            <MenuItem name="2">
              <Icon type="md-checkmark-circle" />
              支付宝
            </MenuItem>
            <MenuItem name="3">
              <Icon type="md-cloud-outline" />
              云闪付
            </MenuItem>
            <MenuItem name="4">
              <Icon type="logo-apple" />
              Apple Pay
            </MenuItem>
          </div>
        </Menu>
      </Header>
      <Layout>
        <Sider hide-trigger :style="{background: '#fff'}">
          <Menu active-name="1-1" theme="light" width="auto" :open-names="['1']">
            <Submenu name="1">
              <template slot="title">
                <Icon type="ios-navigate"></Icon>
                {{mainMenu1}}
              </template>
              <MenuItem name="1-1"><router-link to="/wechatpay/payOrder">{{mainMenu1_01}}</router-link></MenuItem>
              <MenuItem name="1-2"><router-link to="/wechatpay/orderList">{{mainMenu1_02}}</router-link></MenuItem>
            </Submenu>
            <Submenu name="2">
              <template slot="title">
                <Icon type="ios-keypad"></Icon>
                {{mainMenu2}}
              </template>
              <MenuItem name="2-1"><router-link to="/wechatpay/test">{{mainMenu2_01}}</router-link></MenuItem>
              <MenuItem name="2-2"><router-link to="/wechatpay/test">{{mainMenu2_02}}</router-link></MenuItem>
            </Submenu>
            <Submenu name="3">
              <template slot="title">
                <Icon type="ios-analytics"></Icon>
                {{mainMenu3}}
              </template>
              <MenuItem name="3-1"><router-link to="/wechatpay/test">{{mainMenu3_01}}</router-link></MenuItem>
            </Submenu>
          </Menu>
        </Sider>
        <Layout :style="{padding: '0 24px 24px'}">
          <Breadcrumb :style="{margin: '0px 0'}">
            <!--<BreadcrumbItem>{{navigationMenu1}}</BreadcrumbItem>-->
            <!--<BreadcrumbItem>{{navigationMenu2}}</BreadcrumbItem>-->
          </Breadcrumb>
          <Content :style="{padding: '0px', minHeight: '490px', background: '#f5f7f9'}">
            <router-view></router-view>
          </Content>
        </Layout>
      </Layout>
    </Layout>
  </div>
</template>
<script>
  export default {
    name:"baseLayout",
    data() {
      return {
        //菜单
        mainMenu1:'订单管理',
          mainMenu1_01:'下单付款',
          mainMenu1_02:'订单查询',
        mainMenu2:'交易统计',
          mainMenu2_01:'今日交易统计',
          mainMenu2_02:'历史交易统计',
        mainMenu3:'系统配置',
          mainMenu3_01:'微信支付配置',
        //导航
        navigationMenu1:'订单管理',//导航栏
        navigationMenu2:'下单付款',
      }
    },
    methods:{
      changeNavigation:function (p2) {
        alert(p2);
        this.navigationMenu2 = p2;
      }
    },


  }
</script>
<style scoped>
  .layout{
    border: 1px solid #d7dde4;
    background: #f5f7f9;
    position: relative;
    border-radius: 4px;
    overflow: hidden;
  }
  .layout-logo{
    width: 135px;
    height: 55px;
    background: #cdb5f9;
    border-radius: 3px;
    float: left;
    position: relative;
    top: 5px;
    left: -20px;
  }
  .layout-nav{
    width: 40%;
    margin: 0 auto;
    margin-right: 20px;
  }
</style>
