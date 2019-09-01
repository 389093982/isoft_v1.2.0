<template>
  <a @click="onclick" style="display:inline-block;height: inherit;">
    <slot></slot>
  </a>
</template>

<script>
  export default {
    name: "IBeautifulLink2",
    methods:{
      onclick:function () {
        this.$emit('onclick');
      }
    }
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
