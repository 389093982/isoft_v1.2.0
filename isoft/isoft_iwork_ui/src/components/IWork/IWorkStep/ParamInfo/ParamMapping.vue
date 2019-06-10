<template>
  <span>
    <Row>
      <ParamMappingAdd @handleSubmit="paramMappingAdd" v-show="paramMappings.length == 0"/>
      <div v-for="paramMapping in paramMappings">
        <table>
          <tr>
            <td>
              <Input style="width: 200px;" size="small" type="text" v-model="paramMapping.paramMappingName"/>
            </td>
            <td>
              <Select style="width: 100px;" size="small" v-model="paramMapping.paramMappingType" :transfer="true">
                <Option value="default">default</Option>
                <Option value="objects">objects</Option>
              </Select>
            </td>
            <td>
              <Button type="success" size="small" @click="handleDelete(paramMapping.paramMappingName)" style="margin-left: 6px">删除</Button>
              <ParamMappingAdd @handleSubmit="paramMappingAdd"/>
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
      paramMappings:{
        type:Array,
        default:[],
      }
    },
    methods:{
      paramMappingAdd:function (paramMappingName, paramMappingType) {
        var exist = false;
        for(var i=0; i<this.paramMappings.length; i++){
          var paramMapping = this.paramMappings[i];
          if(paramMapping.paramMappingName == paramMappingName){
            exist = true;
            paramMapping.paramMappingType = paramMappingType;
            this.paramMappings[i] = paramMapping;
          }
        }
        if(!exist){
          this.paramMappings.push({"paramMappingName":paramMappingName, "paramMappingType":paramMappingType});
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
