<template>
  <div>
    <Tree :data="data3" :load-data="loadData" show-checkbox></Tree>
  </div>
</template>

<script>
  export default {
    name: "WorkStepTree",
    data(){
      return {
        data3: [
          {
            title: 'parent',
            loading: false,
            children: []
          }
        ]
      }
    },
    methods: {
      loadData (item, callback) {
        const data = [
          {
            title: 'children',
            loading: false,
            children: []
          },
          {
            title: 'children',
            loading: false,
            children: []
          }
        ];
        callback(data);
      }
    }
  }
</script>

<style scoped>

</style>
