<template>
  <div v-if="advertisements && advertisements.length > 0" style="margin-top: 5px;">
    <Row :gutter="10">
      <Col span="12" v-for="(advertisement, index) in advertisements">
        <a target="_blank" :href="advertisement.linked_refer" :title="advertisement.linked_refer">
          <img :src="advertisement.linked_img" width="130px;" height="80px;"/>
        </a>
      </Col>
    </Row>
  </div>
</template>

<script>
  import {GetRandomAdvertisement} from "../../api"

  export default {
    name: "RandomAdmt2",
    data(){
      return {
        advertisements:null,
      }
    },
    methods:{
      refreshRandomAdvertisement:async function(){
        var _this = this;
        const result = await GetRandomAdvertisement(4);
        if(result.status == "SUCCESS"){
          this.advertisements = result.advertisements;
        }
      },
    },
    mounted(){
      this.refreshRandomAdvertisement();
    }
  }
</script>

<style scoped>
  a:hover img{
    border:1px solid red;
  }
  *{
    box-sizing: border-box;
  }
</style>
