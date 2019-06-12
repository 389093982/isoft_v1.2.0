<template>
  <Modal
    v-model="showTemplateModal"
    width="950"
    title="选择模板"
    :footer-hide="true"
    :transfer="false"
    :mask-closable="false"
    :styles="{top: '20px'}">
    <div style="min-height: 380px;">
      <Template :show-chooser-btn="true" @chooseTemplate="chooseTemplate"/>
    </div>
  </Modal>
</template>

<script>
  import Template from "../../IAssistant/Template"

  export default {
    name: "TemplateChooser",
    components:{Template},
    data(){
      return {
        showTemplateModal:false,
      }
    },
    methods:{
      showModal:function () {
        this.showTemplateModal = true;
      },
      chooseTemplate:function (template) {
        this.$emit('chooseTemplate',template);
      },
      hideModal:function () {
        this.showTemplateModal = false;
      }
    }
  }
</script>

<style scoped>

</style>
