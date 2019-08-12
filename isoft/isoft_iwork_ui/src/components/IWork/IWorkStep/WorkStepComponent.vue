<template>
  <Drawer title="全部组件" placement="left" :closable="true" :mask="false" v-model="showComponentDrawer">
    <Tabs value="name1">
      <TabPane label="组件" name="name1">
        <span v-for="default_work_step_type in nodeMetas"
                draggable="true" @dragstart="dragstart($event, 'work_type__' + default_work_step_type.name)">
         <Tag v-if="showComponent(default_work_step_type.name)">{{default_work_step_type.name}}</Tag>
        </span>
      </TabPane>

      <TabPane label="流程" name="name2">
        <div v-for="work in works"
                draggable="true" @dragstart="dragstart($event, 'work_name__' + work.work_name)">
          <Tag>{{ work.module_name }} ~~~ {{ work.work_name }}</Tag>
        </div>
      </TabPane>
    </Tabs>
  </Drawer>
</template>

<script>
  import {GetAllFiltersAndWorks, GetMetaInfo} from "../../../api"
  import {oneOf} from "../../../tools"

  export default {
    name: "WorkStepComponent",
    data(){
      return {
        showComponentDrawer:false,
        nodeMetas: [],
        works:[],
      }
    },
    methods:{
      showComponent:function(name){
        return !oneOf(name, ['work_start',"work_end"]);   // 开始和结束节点不能添加和拖拽
      },
      dragstart:function(event, transferData){
        event.dataTransfer.setData("Text", transferData);
      },
      toggleShow:function () {
        this.showComponentDrawer = !this.showComponentDrawer;
      },
      refreshAllWorks:async function () {
        const result = await GetAllFiltersAndWorks();
        if(result.status == "SUCCESS"){
          this.works = result.works;
        }
      },
      refreshNodeMetas:async function () {
        const result = await GetMetaInfo("nodeMetas");
        if(result.status == "SUCCESS"){
          this.nodeMetas = result.nodeMetas;
        }
      }
    },
    mounted(){
      this.refreshAllWorks();
      this.refreshNodeMetas();
    }
  }
</script>

<style scoped>

</style>
