<template>
  <div>
    <a>热门用户</a>
    <a>积分排行榜</a>

    <Row v-for="user in users">{{user.user_name}}</Row>
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

</style>
