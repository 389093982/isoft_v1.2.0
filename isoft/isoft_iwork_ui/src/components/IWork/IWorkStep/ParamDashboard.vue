<template>
  <div>
    <Row>
      <Col span="12" v-for="workstep in worksteps">
        <WorkStepParamDashboard :workstep="workstep"/>
      </Col>
    </Row>
  </div>
</template>

<script>
  import WorkStepParamDashboard from "./WorkStepParamDashboard"

  export default {
    name: "ParamDashboard",
    components:{WorkStepParamDashboard},
    props:{
      worksteps:{
        type:Array,
        default:[],
      }
    },
  }
</script>

<style scoped>

</style>
