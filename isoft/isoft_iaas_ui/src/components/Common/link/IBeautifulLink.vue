<template>
  <a @click="onclick" style="display:inline-block;height: inherit;">
    <slot></slot>
  </a>
</template>

<script>
  export default {
    name: "IBeautifulLink",
    methods:{
      onclick:function () {
        this.$emit('onclick');
      }
    }
  }
</script>

<style scoped>
  a{
    border-bottom: 2px solid rgba(255, 0, 0, 0);
  }
  a:hover{
    border-bottom: 2px solid red;
  }
</style>
