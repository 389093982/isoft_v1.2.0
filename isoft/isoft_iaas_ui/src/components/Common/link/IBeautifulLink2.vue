<template>
  <a @click="onclick" :style="styles">
    <slot></slot>
  </a>
</template>

<script>
  export default {
    name: "IBeautifulLink2",
    props:{
      inLine:{
        type:Boolean,
        default:true,
      }
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
        return style;
      }
    },
  }
</script>

<style scoped>
  a{
    color: #626262;
  }
  a:hover{
    color: red;
  }
</style>
