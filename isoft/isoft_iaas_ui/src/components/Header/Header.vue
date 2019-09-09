<template>
  <div class="layout">
    <Menu mode="horizontal" :theme="theme1" active-name="1">
      <div class="layout-nav">
        <Submenu name="1">
          <template slot="title">
            <Icon type="ios-stats" />
            精品应用
          </template>
          <MenuGroup title="博客天地">
            <MenuItem name="1-1"><router-link to="/iblog/blog_list">博客天地</router-link></MenuItem>
          </MenuGroup>
          <MenuGroup title="在线学习系统">
            <MenuItem name="1-5"><router-link to="/ilearning/index">精品课程</router-link></MenuItem>
            <MenuItem name="1-6"><router-link to="/inote/inote_list">云笔记</router-link></MenuItem>
          </MenuGroup>
          <MenuGroup title="友情链接">
            <MenuItem name="1-7"><router-link to="/share/list">共享空间</router-link></MenuItem>
          </MenuGroup>
        </Submenu>
        <MenuItem name="2">
          <Icon type="ios-construct" />
          综合设置
        </MenuItem>
        <Submenu name="3">
          <template slot="title">
            <Icon type="ios-people" />
            <span v-if="loginUserName">{{loginUserName}}</span>
            <span v-else>登录</span>
          </template>
          <MenuGroup title="账号管理">
            <MenuItem name="3-1" @click.native="cancelUser">注销</MenuItem>
            <MenuItem name="3-2" @click.native="cancelUser">切换账号</MenuItem>
          </MenuGroup>
        </Submenu>
        <MenuItem name="4">
          <Icon type="ios-construct" />
          <router-link to="/background/cms/placement_list">管理控制台</router-link>
        </MenuItem>
      </div>
    </Menu>
  </div>
</template>

<script>
  import {getCookie} from '../../tools/index'
  import {delCookie} from '../../tools/index'
  import {LoginAddr} from "../../api"

  export default {
    name: "Header",
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
    width: 620px;
    margin: 0 auto;
    margin-right: 20px;
  }
</style>

