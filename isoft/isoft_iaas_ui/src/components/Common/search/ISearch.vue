<template>
  <div id="course_search" style="height: 45px;">
    <div style="position: relative;float: right;">
      <div id="_search">
        <input class="search_input" title="请输入搜索文字" placeholder="请输入搜索文字..." v-model.trim="search_data">
        <input class="submit" type="submit" title="提交" @click="submitFunc">
      </div>
    </div>
  </div>
</template>

<script>
  export default {
    name: "Search",
    data(){
      return {
        search_data:"",
      }
    },
    methods:{
      submitFunc:function () {
        this.$emit("submitFunc", this.search_data);
      }
    }
  }
</script>

<style scoped>
  .search_input{
    padding-left: 16px;
    float: left;
    width: 350px;
    background: #efefef;
    vertical-align: middle;
    height: 40px;
    box-shadow: none;
    -webkit-appearance: none;
    border: none;
    outline: 0;
    color: #333;
    font-size: 14px;
    transition: all ease-in 500ms;
  }
  .search_input:focus{
    width: 450px;
    box-shadow: 0 0 3px rgba(142, 142, 142, 0.5);
  }
  .submit{
    background: url('../../../assets/search.png') 50% 50% no-repeat;
    background-size: 24px;
    position: absolute;
    right: 10px;
    top: 15px;
    margin-top: -7px;
    width: 24px;
    height: 24px;
    border: none;
    color: transparent;
    box-sizing: initial;
    cursor: pointer;
    outline: 0;
  }
</style>
