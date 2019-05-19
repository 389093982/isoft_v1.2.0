<template>
  <ISimpleBtnTriggerModal ref="triggerModal" btn-text="新增" btn-size="small" modal-title="新增" :modal-width="500">
    <Input style="margin-bottom: 20px;" v-model.trim="paramMappingName" placeholder="请输入新增的字段名称"></Input>
    <Select v-model="paramMappingType">
      <Option value="default">default</Option>
      <Option value="objects">objects</Option>
    </Select>
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
        paramMappingType:"default",
      }
    },
    methods:{
      handleSubmit:function () {
        this.$emit("handleSubmit", this.paramMappingName, this.paramMappingType);
        this.$refs.triggerModal.hideModal();
      }

    }
  }
</script>

<style scoped>

</style>
