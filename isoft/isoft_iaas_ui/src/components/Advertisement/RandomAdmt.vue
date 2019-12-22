<template>
  <div v-if="advertisements && showAdv" class="rightSuspension isoft_bg_white">
    <div v-for="(advertisement, index) in advertisements">
      <a target="_blank" :href="advertisement.linked_refer" :title="advertisement.linked_refer">
        <img :src="advertisement.linked_img" width="95px;" height="50px;" @click="gotoAdvertisement(advertisement)"/>
      </a>

    </div>
  </div>
</template>

<script>
  import {GetRandomAdvertisement} from "../../api"

  export default {
    name: "RandomAdmt",
    data(){
      return {
        advFunc:null,
        showAdv:false,
        advertisements:null,
      }
    },
    methods:{
      refreshRandomAdvertisement:async function(){
        var _this = this;
        const result = await GetRandomAdvertisement();
        if(result.status == "SUCCESS"){
          this.advertisements = result.advertisements;
          this.showAdv = true;
          setTimeout(function () {
            _this.showAdv = false;
          }, 3000);
        }
      },
    },
    mounted(){
      var _this = this;
      this.advFunc = setInterval(function(){
        _this.refreshRandomAdvertisement();
      },5000);
    },
    beforeDestroy(){
      clearInterval(this.advFunc);
    }
  }
</script>

<style scoped>
  @import "../../assets/css/isoft_common.css";

  .rightSuspension {
    position: fixed;
    width: 95px;
    height: 100%;
    right: 0px;
    top: 0px;
    z-index: 1000;
  }
  a:hover img{
    border:1px solid red;
  }
</style>
