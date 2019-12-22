<template>
  <!-- 需要动画元素放在transition标签内部，在样式里添加类名属性 -->
  <!-- 多个元素使用transition-group -->
  <transition name="fadeInOut">
    <div v-if="advertisements && showAdv" class="rightSuspension isoft_bg_white">
      <div v-for="(advertisement, index) in advertisements">
        <a target="_blank" :href="advertisement.linked_refer" :title="advertisement.advertisement_label">
          <img :src="advertisement.linked_img" width="95px;" height="50px;"/>
        </a>
      </div>
    </div>
  </transition>
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
        const result = await GetRandomAdvertisement(10);
        if(result.status == "SUCCESS"){
          this.advertisements = result.advertisements;
          this.showAdv = true;
          setTimeout(function () {
            _this.showAdv = false;
          }, 10000);
        }
      },
    },
    mounted(){
      var _this = this;
      this.advFunc = setInterval(function(){
        _this.refreshRandomAdvertisement();
      },10000);
    },
    beforeDestroy(){
      clearInterval(this.advFunc);
    }
  }
</script>

<style scoped>

  .fadeInOut-enter,.fadeInOut-leave-to{
    opacity: 0;
  }
  .fadeInOut-enter-to,.fadeInOut-leave{
    opacity: 1;
  }
  .fadeInOut-enter-active,.fadeInOut-leave-active{
    transition: all 1s;
  }

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
