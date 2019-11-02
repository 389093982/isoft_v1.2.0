<template>
  <div class="box" @mouseenter="mouseenter">
    <img :src="srcImg" width="154px;" height="86px;" class="hoverScaleImg"/>
  </div>
</template>

<script>
  export default {
    name: "HoverBigImg",
    props:{
      srcImg:{
        type:String,
        default:"../../../assets/2.jpeg",
      }
    },
    methods:{
      mouseenter:function () {
        this.$emit("mouseenter");
      }
    }
  }
</script>

<style scoped>
  .box{
    width: 154px;
    height: 86px;
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
