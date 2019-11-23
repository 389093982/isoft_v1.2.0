<template>
  <div>
    <div v-if="user" >
      <div style="min-height: 140px;background: linear-gradient(to right, rgba(29,255,139,0.14), rgba(206,54,255,0.23));
          background-size: cover;background-position: 50%;background-repeat: no-repeat;"></div>

      <Row style="min-height: 150px;background-color: #ffffff;padding: 20px;">
        <Col span="6" style="top:-100px;">
          <img width="150" height="150" v-if="user.small_icon" :src="user.small_icon" @error="defImg()">
          <p style="margin: 0 0 0 40px;" v-if="$route.query.username == 'mine'">
            <IFileUpload ref="fileUpload" @uploadComplete="uploadComplete" action="/api/iwork/httpservice/fileUpload" uploadLabel="上传头像"/>
          </p>
        </Col>
        <Col span="12" style="padding-top: 30px;">
          <p style="margin-bottom: 20px;">加入时间：{{user.created_time}}</p>

          <h3>{{user.user_name}}</h3>
          <p>这家伙很懒，什么个性签名都没有留下</p>
        </Col>
        <Col span="6" style="padding-top: 100px;text-align: right;">
          <Button @click="$router.push({ path: '/iblog/mine/blog_edit'})">发&nbsp;&nbsp;&nbsp;&nbsp;帖</Button>
          <Button @click="$router.push({path:'/user/mine/detail',query:{username:'mine'}})">编辑个人资料</Button>
        </Col>
      </Row>
    </div>

    <div style="min-height: 150px;background-color: #ffffff;margin: 10px 0 0 0; padding:10px;">
      <Row :gutter="10">
        <Col span="16">
          <UserAbout :user-name="_userName"/>
        </Col>
        <Col span="8">
          <HotUser/>
        </Col>
      </Row>
    </div>
  </div>
</template>

<script>
  import {GetUserDetail,UpdateUserIcon} from "../../api"
  import HotUser from "./HotUser"
  import {GetLoginUserName} from "../../tools"
  import IFileUpload from "../Common/file/IFileUpload"
  import UserAbout from "./UserAbout";

  export default {
    name: "UserDetail",
    components: {UserAbout, HotUser,IFileUpload},
    data(){
      return {
        user:null,
        defaultImg: require('../../assets/default.png'),
      }
    },
    methods:{
      defImg(){
        let img = event.srcElement;
        img.src = this.defaultImg;
        img.onerror = null; //防止闪图
      },
      uploadComplete: async function (data) {
        if(data.status == "SUCCESS"){
          this.$refs.fileUpload.hideModal();
          let uploadFilePath = data.fileServerPath;
          const result = await UpdateUserIcon(GetLoginUserName(), uploadFilePath);
          if(result.status == "SUCCESS"){
            this.refreshUserDetail();
          }
        }
      },
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
    computed:{
      _userName:function () {
        return this.$route.query.username == 'mine' ? GetLoginUserName() : this.$route.query.username;
      }
    },
    watch: {
      '$route':'refreshUserDetail'
    },
  }
</script>

<style scoped>

</style>
