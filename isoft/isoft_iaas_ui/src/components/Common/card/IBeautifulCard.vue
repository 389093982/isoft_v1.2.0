<template>
  <div class="box_bg">
    <div class="box_area">
      <slot name="header">
        <div class="title">
          <span class="leftHeaderLeft" style="float: left;">
            <slot name="header_left">{{title}}</slot>
          </span>
          <span style="float: right;">
            <slot name="header_right"></slot>
          </span>

        </div>
      </slot>
      <slot name="content"></slot>
    </div>
  </div>
</template>

<script>
  export default {
    name: "IBeautifulCard",
    props: {
      title: {
        type: String,
        default: '默认标题'
      },
    },
  }
</script>

<style scoped>
 .box_bg {
   padding: 2px;
   margin: 2px;
 }
 .box_area {
   background-color: #ffffff;
 }
 .title {
   height: 40px;
   line-height: 40px;
   padding: 0 15px;
   font-size: 16px;
   text-align: left;
   background: linear-gradient(to right, rgba(255, 0, 0, 0.15), rgba(255, 0, 0, 0.03));
 }
 
 .leftHeaderLeft{
    cursor: pointer;
    transition: all 2s ease-in;
 }
 .leftHeaderLeft:hover{
    color: #00b500;
 }
</style>
