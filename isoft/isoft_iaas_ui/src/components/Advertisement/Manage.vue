<template>
  <div>
    <EditAdvertisement/>

    <div v-if="advertisements && advertisements.length > 0" class="isoft_bg_white" style="padding: 10px 10px 0 10px;">
      <Row>
        <Col span="6">友情链接名称</Col>
        <Col span="6">链接类型</Col>
        <Col span="6">链接地址</Col>
        <Col span="6">链接图片</Col>
      </Row>
      <Row v-for="(advertisement,index) in advertisements">
        <Col span="6">{{advertisement.advertisement_label}}</Col>
        <Col span="6">{{advertisement.linked_type}}</Col>
        <Col span="6">{{advertisement.linked_refer}}</Col>
        <Col span="6">{{advertisement.linked_img}}</Col>
      </Row>
    </div>

  </div>
</template>

<script>
  import {GetPersonalAdvertisement} from "../../api"
  import EditAdvertisement from "./EditAdvertisement";

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
