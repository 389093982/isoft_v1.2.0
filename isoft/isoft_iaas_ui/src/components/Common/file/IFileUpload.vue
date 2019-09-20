<template>
<span>
  <Button :size="btnSize" type="success" @click="fileUploadModal = true">{{ uploadLabel }}</Button>
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
      autoHideModal:{
        type:Boolean,
        default:false,
      },
      btnSize:{
        type:String,
        default:'default',
      },
      uploadLabel: {
        type: String,
        default: '文件上传'
      },
      action: {
        type: String,
        default: ''
      },
      extraData: {
        type:[Object,Number,String],
        default:null,
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
          // 传递 extraData 数据
          result.extraData = this.extraData;
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
        if(this.autoHideModal){
          this.hideModal();
        }
      },
      hideModal(){
        this.fileUploadModal = false;
      }
    }
  }
</script>

<style scoped>

</style>
