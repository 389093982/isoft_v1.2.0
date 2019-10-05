<template>
  <span>
    <slot v-if="initSuccess"></slot>
  </span>
</template>

<script>
  export default {
    name: "isoft-lazy",
    props:{
      time:{
        required: false,
        default: 0,
      }
    },
    data(){
      return {
        initSuccess:false,
      }
    },
    created(){
      this.initSlot();
    },
    methods:{
      initSlot:function () {
        let _this = this;
        setTimeout(function () {
          _this.initSuccess = true;
        }, Number(_this.time || 0));
      }
    }
  }
</script>

<style scoped>

</style>
