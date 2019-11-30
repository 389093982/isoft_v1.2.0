<template>
  <div class="box" @mouseenter="mouseenter" :style="styles" @click="onclick">
    <img :src="srcImg" :style="styles" class="hoverScaleImg" @error="defImg()"/>
  </div>
</template>

<script>
  export default {
    name: "HoverBigImg",
    props:{
      srcImg:{
        type:String,
        default:"../../../assets/2.jpeg",
      },
      width:{
        type:String,
        default:"154px",
      },
      height:{
        type:String,
        default:"86px",
      }
    },
    data(){
      return {
        defaultImg: require('../../../assets/default.png'),
      }
    },
    methods:{
      defImg(){
        let img = event.srcElement;
        img.src = this.defaultImg;
        img.onerror = null; //防止闪图
      },
      mouseenter:function () {
        this.$emit("mouseenter");
      },
      onclick:function () {
        this.$emit("onclick");
      }
    },
    computed: {

      styles () {
        let style = {};
        style.width = this.width;
        style.height = this.height;
        return style;
      }
    },
  }
</script>

<style scoped>
  .box{
    overflow: hidden;
  }
  .box img{
    cursor: pointer;
    transition: all 2s;
  }
  .box img:hover{
    transform: scale(1.4);
  }
</style>
