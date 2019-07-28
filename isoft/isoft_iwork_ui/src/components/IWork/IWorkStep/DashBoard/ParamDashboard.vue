<template>
  <Modal v-model="showModal" fullscreen title="参数可视">
    <ul>
      <li v-for="workstep in worksteps" style="float: left;">
        <WorkStepParamDashboard :workstep="workstep"/>
      </li>
    </ul>
  </Modal>
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
    data(){
      return {
        showModal:true,
      }
    },
    methods:{
      showParamDashboard:function (flag) {
        this.showModal = flag;
      }
    }
  }
</script>

<style scoped>

</style>
