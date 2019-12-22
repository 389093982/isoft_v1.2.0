<template>
  <div v-if="showAdv" class="rightSuspension isoft_bg_white">
    <div class="itemArea">问题反馈</div>
    <div class="itemArea">问题反馈</div>
    <div class="itemArea">问题反馈</div>
    <div class="itemArea">问题反馈</div>
    <div class="itemArea">问题反馈</div>
  </div>
</template>

<script>
  export default {
    name: "RandomAdmt",
    data(){
      return {
        advFunc:null,
        showAdv:false,
      }
    },
    mounted(){
      var _this = this;
      this.advFunc = setInterval(function(){
        _this.showAdv = true;
        setTimeout(function () {
          _this.showAdv = false;
        }, 3000);
      },2000);
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
</style>
