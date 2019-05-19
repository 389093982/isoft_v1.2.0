<template>
  <span>
    <Row>
      <ParamMappingAdd @handleSubmit="paramMappingAdd"/>
      <div v-for="paramMapping in paramMappings">
        <Row>
          <Col span="12">{{paramMapping.paramMappingName}} -- {{paramMapping.paramMappingType}}</Col>
          <Col span="12" style="text-align: right;">
            <Button type="success" size="small" @click="handleDelete(paramMapping.paramMappingName)" style="margin-left: 6px">删除</Button>
          </Col>
        </Row>
        <hr/>
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
