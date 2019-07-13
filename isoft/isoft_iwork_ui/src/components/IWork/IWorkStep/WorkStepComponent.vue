<template>
  <Drawer title="全部组件" placement="left" :closable="true" :mask="false" v-model="showComponetDrawer">
    <span v-for="default_work_step_type in default_work_step_types" style="margin: 5px;float: left;"
          draggable="true" @dragstart="dragstart($event, default_work_step_type.name)">
     <Tag>{{default_work_step_type.name}}</Tag>
    </span>
  </Drawer>
</template>

<script>
  export default {
    name: "WorkStepComponent",
    data(){
      return {
        showComponetDrawer:false,
        default_work_step_types: this.GLOBAL.default_work_step_types,
      }
    },
    methods:{
      dragstart:function(event, transferData){
        event.dataTransfer.setData("Text", transferData);
      },
      toggleShow:function () {
        this.showComponetDrawer = !this.showComponetDrawer;
      }
    }
  }
</script>

<style scoped>

</style>
