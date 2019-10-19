<template>
  <div>
    <Select v-model="current_filter_id" style="width:400px" @on-change="chooseFilter">
      <Option v-for="filterWork in filterWorks" :value="filterWork.id" :key="filterWork.work_name">
        {{filterWork.work_name}}
      </Option>
    </Select>

    <div v-for="module in modules" style="list-style: none;margin-top: 20px;">
      <Row style="border-bottom: 1px solid rgba(0,34,232,0.16);">
        <Col span="4">
          <Tag><span>{{module.module_name}}</span></Tag>
          <Button type="success" size="small" @click="handleCheckAll(module.module_name)">全选</Button>
        </Col>
        <Col span="20">
          <CheckboxGroup v-model="checkedFilterWorks">
            <Checkbox :label="moduleWork.work_name" v-for="moduleWork in works" v-if="moduleWork.module_name == module.module_name"></Checkbox>
          </CheckboxGroup>
        </Col>
      </Row>
    </div>

    <Button type="success" size="small" @click="saveFilters" style="margin-top: 20px;">保存</Button>
  </div>
</template>

<script>
  import {GetAllFiltersAndWorks} from "../../../api"
  import {SaveFilters} from "../../../api"
  import {oneOf} from "../../../tools"

  export default {
    name: "IFilterList",
    data(){
      return {
        filterWorks:[],
        filters:[],
        modules:[],
        works: [],
        moduleWorks:[],
        current_filter_id:-1,
        checkedFilterWorks:[],
      }
    },
    methods:{
      refreshFilterList:async function () {
        const result = await GetAllFiltersAndWorks();
        if(result.status == "SUCCESS"){
          this.filterWorks = result.filterWorks;
          this.filters = result.filters;
          this.modules = result.modules;
          this.works = result.works;
        }
      },
      handleCheckAll:function (module_name) {
        let work_names = this.works.filter(work => work.module_name == module_name).map(work => work.work_name);
        var allInflag = true;
        for(var i=0; i<work_names.length; i++){
          let work_name = work_names[i];
          if(!oneOf(work_name, this.checkedFilterWorks)){
            allInflag = false;
            break;
          }
        }
        if(allInflag){
          this.checkedFilterWorks = this.checkedFilterWorks.filter(checkWork => !oneOf(checkWork, work_names));
        }else{
          let addWorks = work_names.filter(work_name => !oneOf(work_name, this.checkedFilterWorks));
          for(var i=0; i<addWorks.length; i++){
            let work_name = addWorks[i];
            this.checkedFilterWorks.push(work_name);
          }
        }
      },
      saveFilters:async function () {
        if(this.current_filter_id < 0){
          this.$Message.error('请选择 filter!');
        }else{
          const result = await SaveFilters(this.current_filter_id, JSON.stringify(this.checkedFilterWorks));
          if(result.status == "SUCCESS"){
            this.$Message.success("保存成功！");
            this.refreshFilterList();
          }else{
            this.$Message.error(result.errorMsg);
          }
        }
      },
      chooseFilter:function () {
        this.checkedFilterWorks = this.filters.filter(filter => filter.filter_work_id == this.current_filter_id).map(filter => filter.work_name);
      }
    },
    mounted(){
      this.refreshFilterList();
    },
  }
</script>

<style scoped>

</style>
