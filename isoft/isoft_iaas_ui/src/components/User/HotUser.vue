<template>
  <IBeautifulCard title="用户排行榜">
    <div slot="content" style="padding: 10px;">
      <Row v-for="user in users">
        <Col span="4">
          <img width="30" height="30" v-if="user.small_icon" :src="user.small_icon">
          <img width="30" height="30" v-else src="../../../src/assets/sso/default_user_small_icon.jpg">
        </Col>
        <Col span="10">
          <IBeautifulLink2 @onclick="$router.push({path:'/user/detail',query:{username:user.user_name}})">{{user.user_name}}</IBeautifulLink2>
        </Col>
        <Col span="10" class="small_font_size">
          用户积分数：{{user.user_points}}
        </Col>
      </Row>
    </div>
  </IBeautifulCard>
</template>

<script>
  import {GetHotUsers} from "../../api"
  import IBeautifulCard from "../../components/Common/card/IBeautifulCard";
  import IBeautifulLink2 from "../Common/link/IBeautifulLink2";

  export default {
    name: "HotUser",
    components: {IBeautifulLink2, IBeautifulCard},
    data(){
      return {
        users:[],
      }
    },
    methods:{
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
