<template>
  <ISimpleConfirmModal ref="modal" :modal-width="1000" @handleSubmit="handleSubmit" modal-title="占位符管理">
    <IBaseChooser ref="placement_chooser" chooser-title="占位符选择">
      <Placement :chooserMode="true" @choosePlacement="choosePlacement"/>
    </IBaseChooser>

    <Form ref="formInline" :model="formInline" :rules="ruleInline" :label-width="100">
      <Row>
        <Col span="12">
          <FormItem label="占位符">
            <Input type="text" readonly="readonly" v-model="formInline.placement" placeholder="placement" style="width: 80%;"/>
            <Button type="success" @click="$refs.placement_chooser.showModal()">选择</Button>
          </FormItem>
          <FormItem prop="title" label="标题">
            <Input type="text" v-model="formInline.title" placeholder="title" style="width: 80%;"/>
          </FormItem>
          <FormItem prop="content"  label="内容">
            <Input type="text" v-model="formInline.content" placeholder="content" style="width: 80%;"/>
          </FormItem>
        </Col>
        <Col span="12">
          <FormItem prop="imgpath"  label="图片">
            <Input type="text" readonly="readonly" v-model="formInline.imgpath" placeholder="imgpath" style="width: 80%;"/>
            <IFileUpload @uploadComplete="uploadComplete" action="/api2/iwork/fileUpload/fileUpload" uploadLabel="上传"/>
          </FormItem>
          <FormItem prop="linked_refer"  label="链接关键词">
            <Input type="text" v-model="formInline.linked_refer" placeholder="linked_refer" style="width: 80%;"/>
          </FormItem>
        </Col>
      </Row>
    </Form>
  </ISimpleConfirmModal>
</template>

<script>
  import ISimpleConfirmModal from "../../Common/modal/ISimpleConfirmModal"
  import IBaseChooser from "../../Common/IBaseChooser"
  import Placement from "./Placement"
  import IFileUpload from "../../IFile/IFileUpload"
  import {AddElement} from "../../../api"

  export default {
    name: "EditElement",
    components:{IBaseChooser,Placement,IFileUpload,ISimpleConfirmModal},
    data(){
      return {
        formInline: {
          placement:'',
          title: '',
          content: '',
          imgpath: '',
          linked_refer: '',
        },
        ruleInline: {
          title: [
            { required: true, message: 'Please fill in the title.', trigger: 'blur' },
          ],
        }
      }
    },
    methods:{
      showModal:function(){
        this.$refs.modal.showModal();
      },
      handleSubmit() {
        this.$refs['formInline'].validate(async (valid) => {
          if (valid) {
            const result = await AddElement(this.formInline.placement, this.formInline.title, this.formInline.content,
              this.formInline.imgpath, this.formInline.linked_refer);
            if(result.status=="SUCCESS"){
              this.$emit("refreshElementList");
              this.$Message.success('提交成功!');
            }else{
              this.$Message.error('提交失败!' + result.errorMsg);
            }
          } else {
            this.$Message.error('校验不通过!');
          }
        })
      },
      choosePlacement:function (placement_name) {
        this.formInline.placement = placement_name;
        this.$refs.placement_chooser.hideModal();
      },
      uploadComplete: function (result) {
        if(result.status == "SUCCESS"){
          this.formInline.imgpath = result.fileServerPath;
        }
      },
      initFormData:function (element) {
        this.formInline.placement = element.placement;
        this.formInline.title = element.title;
        this.formInline.content = element.content;
        this.formInline.imgpath = element.img_path;
        this.formInline.linked_refer = element.linked_refer;
      }
    }
  }
</script>

<style scoped>

</style>
