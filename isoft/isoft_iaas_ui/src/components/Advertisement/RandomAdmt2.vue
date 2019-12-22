<template>
  <div v-if="advertisements && advertisements.length > 0" style="margin-top: 5px;">
    <Row :gutter="10">
      <Col span="12" v-for="(advertisement, index) in advertisements" style="height: 80px;margin-bottom: 10px;">
        <a target="_blank" :href="advertisement.linked_refer" :title="advertisement.advertisement_label">
          <img :src="advertisement.linked_img" width="120px;" height="80px;"/>
          <div class="advertisement_label">{{advertisement.advertisement_label}}</div>
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
  a .advertisement_label {
    padding-left: 10px;
    background-color: rgba(0,0,0,0.6);
    color: white;
    width: 120px;
    height: 30px;
    position: relative;
    transition: all ease-in 1s;
    display: block;
    top: -37px;
  }
  a:hover .advertisement_label {
    display: none;
    top: 0px;

  }
</style>
