<template>
  <div class="layout">
    <Menu mode="horizontal" :theme="theme1" active-name="1">
      <div class="layout-nav">
        <MenuItem name="1">
          <IBeautifulLink @onclick="$router.push({path:'/ilearning/index'})">精品课程</IBeautifulLink>
        </MenuItem>
        <MenuItem name="2">
        <IBeautifulLink @onclick="$router.push({path:'/iblog/blog_list'})">话题博客</IBeautifulLink>
      </MenuItem>
        <MenuItem name="3">
          <IBeautifulLink @onclick="$router.push({path:'/iblog/book_list'})">热门书单</IBeautifulLink>
        </MenuItem>
        <MenuItem name="4">
          <IBeautifulLink @onclick="$router.push({path:'/ifound/found_list'})">发现频道</IBeautifulLink>
        </MenuItem>
        <MenuItem name="5">
          <IBeautifulLink @onclick="$router.push({path:'/ifound/found_list'})">热门活动</IBeautifulLink>
        </MenuItem>
        <MenuItem name="6">
          <IBeautifulLink @onclick="$router.push({path:'/background/cms/configuration'})">管理控制台</IBeautifulLink>
        </MenuItem>
        <Submenu name="7">
          <template slot="title">
            <span>更多内容</span>
          </template>
          <MenuGroup title="更多内容">
            <MenuItem name="7-1">全部菜单项</MenuItem>
            <MenuItem name="7-2">精选内容</MenuItem>
            <MenuItem name="7-3">更多内容</MenuItem>
          </MenuGroup>
        </Submenu>
        <Submenu name="8">
          <template slot="title">
            <span v-if="loginUserName">{{loginUserName}}</span>
            <span v-else>未登录</span>
          </template>
          <MenuGroup title="账号管理">
            <MenuItem name="8-1" @click.native="cancelUser">前往登录</MenuItem>
            <MenuItem name="8-2" @click.native="cancelUser">切换账号</MenuItem>
            <MenuItem name="8-3" @click.native="cancelUser">注销</MenuItem>
          </MenuGroup>
        </Submenu>
        <MenuItem name="9">
          <IBeautifulLink @onclick="$router.push({path:'/shareArticle/shareArticlePlace'})">作文分享</IBeautifulLink>
        </MenuItem>
        <MenuItem name="10">
          <IBeautifulLink @onclick="$router.push({path:'/job/jobList'})">求职招聘</IBeautifulLink>
        </MenuItem>
        <MenuItem name="11">
          <IBeautifulLink @onclick="$router.push({path:'/vipcenter/vipIntroduction'})">会员中心</IBeautifulLink>
        </MenuItem>
      </div>
    </Menu>

    <LevelTwoHeader/>
  </div>
</template>

<script>
  import {getCookie, delCookie, CheckHasLogin} from '../../tools/index'
  import {LoginAddr} from "../../api"
  import IBeautifulLink from "../Common/link/IBeautifulLink";
  import LevelTwoHeader from "./LevelTwoHeader";

  export default {
    name: "Header",
    components: {LevelTwoHeader, IBeautifulLink},
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
      if (CheckHasLogin()){
        this.loginUserName = getCookie("nickName");
      }
    },
  }
</script>

<style scoped>
  .layout-nav{
    float: right;
    width: 1200px;
  }
</style>

