<template>
  <div>
    <a>热门用户</a>
    <a>积分排行榜</a>
    <Row v-for="user in users">
      <Col span="4">
        <img width="30" height="30" v-if="user.small_icon" :src="user.small_icon">
        <img width="30" height="30" v-else src="../../../src/assets/sso/default_user_small_icon.jpg">
      </Col>
      <Col span="10">
        {{user.user_name}}
      </Col>
      <Col span="10" class="small_font_size">
        用户积分数：{{user.user_points}}
      </Col>
    </Row>
  </div>
</template>

<script>
  import {GetHotUsers} from "../../api"

  export default {
    name: "HotUser",
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
