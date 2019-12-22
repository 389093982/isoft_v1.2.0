<template>
  <div>
    <IBeautifulCard title="用户排行榜">
      <div slot="content" style="padding: 10px;">
        <Row v-for="user in users" :gutter="10">
          <Col span="4">
            <img width="30" height="30" :src="user.small_icon" @error="defImg()">
          </Col>
          <Col span="10" class="isoft_inline_ellipsis">
            <IBeautifulLink @onclick="$router.push({path:'/user/detail',query:{username:user.user_name}})">{{user.user_name}}</IBeautifulLink>
          </Col>
          <Col span="10" class="small_font_size">
            用户积分数：{{user.user_points}}
          </Col>
        </Row>
      </div>
    </IBeautifulCard>
  </div>
</template>

<script>
  import {GetHotUsers} from "../../api"
  import IBeautifulCard from "../../components/Common/card/IBeautifulCard";
  import IBeautifulLink from "../Common/link/IBeautifulLink";

  export default {
    name: "HotUser",
    components: {IBeautifulLink, IBeautifulCard},
    data(){
      return {
        users:[],
        defaultImg: require('../../assets/default.png'),
      }
    },
    methods:{
      defImg(){
        let img = event.srcElement;
        img.src = this.defaultImg;
        img.onerror = null; //防止闪图
      },
      refreshHotUsers:async function () {
        const result = await GetHotUsers();
        if (result.status == "SUCCESS"){
          this.users = result.users;
        }
      }
    },
    mounted(){
      this.refreshHotUsers();
    }
  }
</script>

<style scoped>


  .small_font_size {
    font-size: 12px;
  }
</style>
