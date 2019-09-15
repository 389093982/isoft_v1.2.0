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
        <Col span="12" style="padding-top: 30px;">
          <p style="margin-bottom: 20px;">加入时间：{{user.created_time}}</p>

          <h3>{{user.user_name}}</h3>
          <p>这家伙很懒，什么个性签名都没有留下</p>
        </Col>
        <Col span="6" style="padding-top: 100px;text-align: right;">
          <Button @click="$router.push({path:'/user/mine/detail',query:{username:'mine'}})">编辑个人资料</Button>
        </Col>
      </Row>
    </div>

    <div>
      <Row>
        <Col span="16">
          AAAAA
        </Col>
        <Col span="8">
          <HotUser/>
        </Col>
      </Row>
    </div>
  </div>
</template>

<script>
  import {GetUserDetail} from "../../api"
  import HotUser from "./HotUser"
  import {GetLoginUserName} from "../../tools"

  export default {
    name: "UserDetail",
    components: {HotUser},
    data(){
      return {
        user:null,
      }
    },
    methods:{
      refreshUserDetail:async function () {
        let userName = this.$route.query.username == 'mine' ? GetLoginUserName() : this.$route.query.username;
        const result = await GetUserDetail(userName);
        if(result.status == "SUCCESS"){
          this.user = result.user;
        }
      }
    },
    mounted(){
      if(this.$route.query.username != undefined && this.$route.query.username != null){
        this.refreshUserDetail();
      }
    },
    watch: {
      '$route':'refreshUserDetail'
    },
  }
</script>

<style scoped>

</style>
