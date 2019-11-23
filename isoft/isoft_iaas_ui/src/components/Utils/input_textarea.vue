<template>
  <div>
    <Input v-model="content" :rows="rows" :maxlength="maxlength" show-word-limit="true" type="textarea" placeholder="Enter something..." style="width: 100%" />
  </div>
</template>

<script>
  export default {
    name: "input_textarea",
    props:{
      content:String,
      rows:Number,
      maxlength:Number,
    },
  }
</script>

<style scoped>

</style>




