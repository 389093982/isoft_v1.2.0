<template>
  <ISimpleBtnTriggerModal ref="triggerModal" btn-text="新增" btn-size="small" modal-title="新增" :modal-width="500">
    <Input style="margin-bottom: 20px;" v-model.trim="paramMappingName" placeholder="请输入新增的字段名称"></Input>
    <div>
      <div style="margin-top: 20px;margin-bottom: 10px;">推荐字段：</div>
      <CheckboxGroup v-model="choosed_recommend_fields">
        <Checkbox v-for="recommend_field in recommend_fields" :label="recommend_field">
          <span>{{recommend_field}}</span>
        </Checkbox>
      </CheckboxGroup>
    </div>

    <Row style="text-align: right;">
      <Button type="success" @click="handleSubmit()" style="margin-top: 6px">Submit</Button>
    </Row>
  </ISimpleBtnTriggerModal>
</template>

<script>
  import ISimpleBtnTriggerModal from "../../../Common/modal/ISimpleBtnTriggerModal"

  export default {
    name: "ParamMappingAdd",
    components:{ISimpleBtnTriggerModal},
    data(){
      return {
        paramMappingName:"",
        recommend_fields:["status","result","errorMsg"],
        choosed_recommend_fields:[],
      }
    },
    methods:{
      handleSubmit:function () {
        this.choosed_recommend_fields.forEach(choosed_recommend_field => {
          this.$emit("handleSubmit", choosed_recommend_field);
        });
        this.paramMappingName.split(",").forEach(_paramMappingName => {
          if(_paramMappingName.trim() != ""){
            this.$emit("handleSubmit", _paramMappingName.trim());
          }
        });
        this.$refs.triggerModal.hideModal();
      }

    }
  }
</script>

<style scoped>

</style>
