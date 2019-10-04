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
      },
      fontWeight:{
        type:[String,Number],
        default:"normal",
      },
      initialColor:{
        type:String,
        default:'#626262',
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
        style.fontWeight = this.fontWeight;
        style.color = this.initialColor;
        return style;
      }
    },
  }
</script>

<style scoped>
  a:hover{
    color: red;
  }
</style>
