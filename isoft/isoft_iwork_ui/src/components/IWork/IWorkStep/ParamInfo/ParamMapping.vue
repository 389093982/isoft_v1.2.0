<template>
  <span>
    <Row>
      <ParamMappingAdd @handleSubmit="paramMappingAdd" v-show="paramMappings.length == 0"/>
      <div v-for="(paramMapping, index) in paramMappings">
        <table>
          <tr>
            <td>
              变量名：<Input style="width: 150px;" size="small" type="text" v-model="paramMapping.paramMappingName"/>
            </td>
            <td v-if="workStepType == 'work_start'">
              默认值：<Input style="width: 150px;" size="small" type="text" v-model="paramMapping.paramMappingDefault"/>
            </td>
            <td>
              <Icon type="md-arrow-round-up" @click="toggleLocation(index, -1)"/>
              <Icon type="md-arrow-round-down" @click="toggleLocation(index, 1)"/>
              <Button type="error" size="small" @click="handleDelete(paramMapping.paramMappingName)" style="margin-left: 6px">删除</Button>
              <ParamMappingAdd @handleSubmit="paramMappingAdd"/>
            </td>
          </tr>
          <tr>
            <td colspan="3" v-if="workStepType == 'work_start'" style="padding: 2px 0 8px 0;">
               <Checkbox v-model="paramMapping.paramMappingCleanXss">CleanXss</Checkbox>
               <Checkbox v-model="paramMapping.paramMappingSafePageNo">SafePageNo</Checkbox>
               <Checkbox v-model="paramMapping.paramMappingSafePageSize">SafePageSize</Checkbox>
            </td>
          </tr>
        </table>
      </div>
    </Row>
  </span>
</template>

<script>
  import ParamMappingAdd from "./ParamMappingAdd"

  export default {
    name: "ParamMapping",
    components:{ParamMappingAdd},
    props:{
      workStepType:{
        type:String,
        default: '',
      },
      paramMappings:{
        type:Array,
        default:[],
      }
    },
    methods:{
      toggleLocation:function(index, direction){
        let index1 = index;
        let index2 = index + direction;
        if(index1 < 0 || index2 < 0 || index1 >= this.paramMappings.length || index2 >= this.paramMappings.length){
          return;
        }
        let paramMapping1 = this.paramMappings[index1];
        let paramMapping2 = this.paramMappings[index2];
        // 进行替换
        this.paramMappings.splice(index1, 1, paramMapping2);
        this.paramMappings.splice(index2, 1, paramMapping1);
      },
      paramMappingAdd:function (paramMappingName) {
        var exist = false;
        for(var i=0; i<this.paramMappings.length; i++){
          var paramMapping = this.paramMappings[i];
          if(paramMapping.paramMappingName == paramMappingName){
            exist = true;
            this.paramMappings[i] = paramMapping;
          }
        }
        if(!exist){
          this.paramMappings.push({"paramMappingName":paramMappingName});
        }
      },
      handleDelete:function (paramMappingName) {
        for(var i=0; i<this.paramMappings.length; i++){
          var paramMapping = this.paramMappings[i];
          if(paramMapping.paramMappingName == paramMappingName){
            this.paramMappings.splice(i, 1);
          }
        }
      }
    }
  }
</script>

<style scoped>

</style>
