<template>
  <Modal
    v-model="showModuleModal"
    width="950"
    title="选择模块"
    :footer-hide="true"
    :transfer="false"
    :mask-closable="false"
    :styles="{top: '20px'}">
    <div style="min-height: 380px;">
      <IModuleList :show-chooser-btn="true" @chooseModule="chooseModule"/>
    </div>
  </Modal>
</template>

<script>
  import IModuleList from "../IModule/IModuleList"

  export default {
    name: "ModuleChooser",
    components:{IModuleList},
    data(){
      return {
        showModuleModal:false,
      }
    },
    methods:{
      showModal:function () {
        this.showModuleModal = true;
      },
      chooseModule:function (module) {
        this.hideModal();
        this.$emit('chooseModule',module);
      },
      hideModal:function () {
        this.showModuleModal = false;
      },
    }
  }
</script>

<style scoped>

</style>
