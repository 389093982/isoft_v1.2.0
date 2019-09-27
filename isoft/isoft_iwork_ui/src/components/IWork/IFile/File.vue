<template>
  <div>
    <IFileUpload ref="fileUpload" @uploadComplete="uploadComplete" action="/api/iwork/httpservice/fileUpload" uploadLabel="文件上传测试"/>

    <Input :readonly="true" v-model.trim="uploadFilePath" readonly placeholder="上传文件路径"style="width: 300px;"></Input>

    <a :href="uploadFilePath" v-if="uploadFilePath">下载链接地址：{{uploadFilePath}}</a>
  </div>
</template>

<script>
  import IFileUpload from "../../Common/file/IFileUpload"

  export default {
    name: "File",
    components:{IFileUpload},
    data(){
      return {
        uploadFilePath: '',
      }
    },
    methods:{
      uploadComplete: function (result) {
        if(result.status == "SUCCESS"){
          this.uploadFilePath = result.fileServerPath;
          this.$refs.fileUpload.hideModal();
        }
      },
    }
  }
</script>

<style scoped>

</style>
