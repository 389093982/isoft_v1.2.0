<template>
  <div>
    <EditAdvertisement ref="editAdvertisement"/>

    <div v-if="advertisements && advertisements.length > 0" class="isoft_bg_white" style="padding: 10px 10px 0 10px;">
      <Row>
        <Col span="4">友情链接名称</Col>
        <Col span="4">链接类型</Col>
        <Col span="4">链接地址</Col>
        <Col span="4">链接图片</Col>
        <Col span="4">联系人</Col>
        <Col span="4">操作</Col>
      </Row>
      <Row v-for="(advertisement,index) in advertisements">
        <Col span="4">{{advertisement.advertisement_label}}</Col>
        <Col span="4">{{advertisement.linked_type}}</Col>
        <Col span="4">{{advertisement.linked_refer}}</Col>
        <Col span="4">{{advertisement.linked_img}}</Col>
        <Col span="4">{{loginUserName}}</Col>
        <Col span="4"><a @click="editAdvertisement(advertisement.id)">编辑</a></Col>
      </Row>
    </div>

  </div>
</template>

<script>
  import {GetPersonalAdvertisement} from "../../api"
  import {GetLoginUserName} from "../../tools"
  import EditAdvertisement from "./EditAdvertisement"

  export default {
    name: "Manage",
    components: {EditAdvertisement},
    data(){
      return {
        advertisements:null,
      }
    },
    methods:{
      refreshPersonalAdvertisement:async function () {
        const result = await GetPersonalAdvertisement();
        if(result.status == "SUCCESS"){
          this.advertisements = result.advertisements;
        }else{
          this.$Message.error(result.errorMsg);
        }
      },
      editAdvertisement:function (id) {
        this.$refs.editAdvertisement.initData(id);
      }
    },
    computed:{
      loginUserName(){
        return GetLoginUserName();
      }
    },
    mounted(){
      this.refreshPersonalAdvertisement();
    }
  }
</script>

<style scoped>
  @import "../../assets/css/isoft_common.css";

</style>
