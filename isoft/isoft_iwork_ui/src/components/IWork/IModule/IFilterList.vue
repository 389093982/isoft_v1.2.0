<template>
  <div>
    AAAAAAAAAA
  </div>
</template>

<script>
  import {GetAllFilterWorks} from "../../../api"

  export default {
    name: "IFilterList",
    methods:{
      refreshFilterList:async function () {
        const result = await GetAllFilterWorks();
        alert(result);
      }
    },
    mounted(){
      this.refreshFilterList();
    }
  }
</script>

<style scoped>

</style>
