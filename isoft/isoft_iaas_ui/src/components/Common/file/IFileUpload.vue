<template>
<span>
  <Button v-if="showButton" :size="size" type="success" @click="fileUploadModal = true">{{ uploadLabel }}</Button>
  <Modal
    v-model="fileUploadModal"
    width="500"
    :title="uploadLabel"
    :mask-closable="false">
    <div>
      <Upload
        ref="upload"
        multiple
        :on-success="uploadComplete"
        :action="action">
        <Button icon="ios-cloud-upload-outline">{{ uploadLabel }}</Button>
      </Upload>
    </div>
  </Modal>
</span>
</template>

<script>
  export default {
    name: "IFileUpload",
    props: {
      showButton: {
        type: Boolean,
        default: true
      },
      uploadLabel: {
        type: String,
        default: '文件上传'
      },
      action: {
        type: String,
        default: ''
      },
      size: {
        type: String,
        default: 'default'
      }
    },
    data () {
      return {
        // 文件上传 modal
        fileUploadModal: false,
      }
    },
    methods:{
      uploadComplete(result, file) {
        if(result.status=="SUCCESS"){
          // 父子组件通信
          this.$emit('uploadComplete',result);
          this.$Notice.success({
            title: '文件上传成功',
            desc: '文件 ' + file.name + ' 上传成功!'
          });
        }else{
          this.$Notice.error({
            title: '文件上传失败',
            desc: '文件 ' + file.name + ' 上传失败!'
          });
        }
      },
      showModal(){
        this.fileUploadModal = true;
      },
      hideModal(){
        this.fileUploadModal = false;
      }
    }
  }
</script>

<style scoped>

</style>
