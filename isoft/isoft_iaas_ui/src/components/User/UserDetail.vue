<template>
  <div>
    <div v-if="user" >
      <div style="min-height: 140px;background-color: #96a1a9;background-size: cover;
        background-position: 50%;background-repeat: no-repeat;"></div>
      <Row style="min-height: 150px;background-color: #ffffff;padding: 20px;">
        <Col span="6" style="top:-100px;">
          <img width="150" height="150" v-if="user.small_icon" :src="user.small_icon">
          <img width="150" height="150" v-else src="../../../src/assets/sso/default_user_small_icon.jpg">
        </Col>
        <Col span="12" style="padding-top: 100px;">
          <h3>{{user.user_name}}</h3>
          <p>这家伙很懒，什么个性签名都没有留下</p>
        </Col>
        <Col span="6" style="padding-top: 100px;text-align: right;">
          <Button>编辑个人资料</Button>
        </Col>
      </Row>
    </div>
  </div>
</template>

<script>
  import {GetUserDetail} from "../../api"

  export default {
    name: "UserDetail",
    data(){
      return {
        user:null,
      }
    },
    methods:{
      refreshUserDetail:async function (userName) {
        const result = await GetUserDetail(userName);
        if(result.status == "SUCCESS"){
          this.user = result.user;
        }
      }
    },
    mounted(){
      if(this.$route.query.username != undefined && this.$route.query.username != null){
        this.refreshUserDetail(this.$route.query.username);
      }
    }
  }
</script>

<style scoped>

</style>
