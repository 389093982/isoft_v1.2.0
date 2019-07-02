<template>
  <div>
    <IFileUpload @uploadComplete="uploadComplete" action="/api/iwork/fileUpload/" uploadLabel="上传"/>
  </div>
</template>

<script>
  import IFileUpload from "../../Common/file/IFileUpload"

  export default {
    name: "File",
    components:{IFileUpload},
    methods:{
      uploadComplete: function (result) {
        if(result.status == "SUCCESS"){
          alert(result.filepath);
        }
      },
    }
  }
</script>

<style scoped>

</style>
