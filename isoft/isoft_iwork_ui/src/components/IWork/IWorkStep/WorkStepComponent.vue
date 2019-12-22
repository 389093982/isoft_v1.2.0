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
        <Collapse>
          <Panel v-for="module_name in module_names" :name="module_name">
            {{module_name}}
            <div slot="content">
              <div v-if="work.module_name == module_name" v-for="work in works"
                   draggable="true" @dragstart="dragstart($event, 'work_name__' + work.work_name)">
                <Tag><span @click="$router.push({path:'/iwork/workList',query:{work_name:work.work_name}})">{{ work.work_name }}</span></Tag>
              </div>
            </div>
          </Panel>

          <Panel name="all">
            全部
            <div slot="content">
              <div v-for="work in works"
                   draggable="true" @dragstart="dragstart($event, 'work_name__' + work.work_name)">
                <Tag><span @click="$router.push({path:'/iwork/workList',query:{work_name:work.work_name}})">{{ work.work_name }}</span></Tag>
              </div>
            </div>
          </Panel>
        </Collapse>
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
        module_names:[],
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
          let module_names = this.works.map(work => work.module_name);
          this.module_names = this.getUniqueArr(module_names);
        }
      },
      getUniqueArr: function(arr){
        var x = new Set(arr);
        return [...x];
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
