<template>
  <div class="layout">
    <Menu mode="horizontal" :theme="theme1" active-name="1">
      <div class="layout-nav">
        <MenuItem name="1">
          <IBeautifulLink2 @onclick="$router.push({path:'/ilearning/index'})">精品课程</IBeautifulLink2>
        </MenuItem>
        <MenuItem name="2">
          <IBeautifulLink2 @onclick="$router.push({path:'/iblog/blog_list'})">话题博客</IBeautifulLink2>
        </MenuItem>
        <MenuItem name="3">
          <IBeautifulLink2 @onclick="$router.push({path:'/iblog/book_list'})">热门书单</IBeautifulLink2>
        </MenuItem>
        <MenuItem name="4">
          <IBeautifulLink2 @onclick="$router.push({path:'/ifound/found_list'})">发现频道</IBeautifulLink2>
        </MenuItem>
        <MenuItem name="5">
          <IBeautifulLink2 @onclick="$router.push({path:'/ifound/found_list'})">热门活动</IBeautifulLink2>
        </MenuItem>
        <MenuItem name="6">
          <IBeautifulLink2 @onclick="$router.push({path:'/igood/good_list'})">我要赚钱</IBeautifulLink2>
        </MenuItem>
        <Submenu name="7">
          <template slot="title">
            <span v-if="loginUserName">{{loginUserName}}</span>
            <span v-else>登录</span>
          </template>
          <MenuGroup title="账号管理">
            <MenuItem name="7-1" @click.native="cancelUser">注销</MenuItem>
            <MenuItem name="7-2" @click.native="cancelUser">切换账号</MenuItem>
          </MenuGroup>
        </Submenu>
        <MenuItem name="8">
          <IBeautifulLink2 @onclick="$router.push({path:'/background/cms/placement_list'})">管理控制台</IBeautifulLink2>
        </MenuItem>
      </div>
    </Menu>
  </div>
</template>

<script>
  import {getCookie,delCookie} from '../../tools/index'
  import {LoginAddr} from "../../api"
  import IBeautifulLink2 from "../Common/link/IBeautifulLink2";

  export default {
    name: "Header",
    components: {IBeautifulLink2},
    data () {
      return {
        theme1: 'light',
        loginUserName:'',
      }
    },
    methods:{
      cancelUser() {
        delCookie("tokenString");
        delCookie("userName");
        delCookie("isLogin");
        this.loginUserName = "";
        window.location.href = LoginAddr + "?redirectUrl=" + window.location.href;
      }
    },
    mounted:function(){
      this.loginUserName = getCookie("userName");
    },
  }
</script>

<style scoped>
  .layout-nav{
    width: 860px;
    margin: 0 auto;
    margin-right: 5px;
  }
</style>

