<template>
  <div>
    <div class="isoft_bg_white isoft_pd10">
      <Row :gutter="10">
        <Col span="9"><Input v-model="jobNameSearch" placeholder="输入职位或者公司关键词搜索您感兴趣的职位"/></Col>
        <Col span="6"><Input v-model="jobPlaceSearch" placeholder="北京、上海、广州、深圳"/></Col>
        <Col span="6"><Input v-model="jobSalaySearch" placeholder="薪酬范围"/></Col>
        <Col span="3"><Button type="success">搜索职位</Button></Col>
      </Row>
    </div>

    <div class="isoft_bg_white isoft_pd10 isoft_top10">
      <Row>
        <Col span="12">搜索职位:{{jobNameSearch}} {{jobPlaceSearch}} {{jobSalaySearch}}</Col>
        <Col span="12" style="text-align: right;padding-right: 40px;">
          <Button @click="$router.push({path:'/job/employee'})">我要求职</Button>
          <Button @click="$router.push({path:'/job/employer'})">我要招聘</Button>
        </Col>
      </Row>
    </div>
  </div>
</template>

<script>
  export default {
    name: "JobList",
    data(){
      return {
        jobNameSearch:'',
        jobPlaceSearch:'',
        jobSalaySearch:'',
      }
    }
  }
</script>

<style scoped>

</style>
