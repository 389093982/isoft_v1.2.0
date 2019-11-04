<template>
  <a @click="onclick" :style="styles">
    <slot></slot>
  </a>
</template>

<script>
  export default {
    name: "IBeautifulLink",
    props:{
      inLine:{
        type:Boolean,
        default:true,
      },
      fontWeight:{
        type:[String,Number],
        default:"normal",
      },
    },
    methods:{
      onclick:function () {
        this.$emit('onclick');
      }
    },
    computed: {
      styles () {
        let style = {};
        if(this.inLine){
          style.display = 'inline-block';
          style.height = 'inherit';
        }
        style.fontWeight = this.fontWeight;
        return style;
      }
    },
  }
</script>

<style scoped>
  a{
    color:#626262;
  }
  a:hover{
    color: red;
  }
</style>
