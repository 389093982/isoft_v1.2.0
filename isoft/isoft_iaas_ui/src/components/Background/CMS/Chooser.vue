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
      <slot/>
    </div>
  </Modal>
</template>

<script>
  export default {
    name: "Chooser",
    data(){
      return {
        showTemplateModal:false,
      }
    },
    methods:{
      showModal:function () {
        this.showTemplateModal = true;
      },
      hideModal:function () {
        this.showTemplateModal = false;
      }
    }
  }
</script>

<style scoped>

</style>
