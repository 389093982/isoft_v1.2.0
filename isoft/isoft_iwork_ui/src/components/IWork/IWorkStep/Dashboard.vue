<template>
  <div>
    流程调用成功率
    <i-circle :percent="80">
      <span class="demo-Circle-inner" style="font-size:24px">80%</span>
    </i-circle>


    work id : {{workId}}

    调度次数：1
    调度成功率：1
    节点分布：1


  </div>
</template>

<script>
  export default {
    name: "Dashboard",
    props:{
      workId:{
        type:String,
        default:"-1",
      }
    }
  }
</script>

<style scoped>

</style>
