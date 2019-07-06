<template>
  <!-- 设置 transfer 属性为 false 后,保证好两个 Modal 的前后顺序,可以解决顺序问题 -->
  <Modal
    v-model="showFormModal"
    width="800"
    title="快捷函数"
    :footer-hide="true"
    :transfer="false"
    :mask-closable="false"
    :styles="{top: '20px'}">
    <Scroll height="380">
      <Table :columns="columns1" :data="funcs" size="small"></Table>
    </Scroll>
  </Modal>
</template>

<script>
  export default {
    name: "QuickFuncList",
    data(){
      return {
        showFormModal:false,
        funcs:this.GLOBAL.quick_funcs,
        columns1: [
          {
            title: 'funcDesc',
            key: 'funcDesc',
          },
          {
            title: 'funcDemo',
            key: 'funcDemo',
          },
        ],
      }
    },
    methods:{
      showModal: function(){
        this.showFormModal = true;
      }
    }
  }
</script>

<style scoped>

</style>
