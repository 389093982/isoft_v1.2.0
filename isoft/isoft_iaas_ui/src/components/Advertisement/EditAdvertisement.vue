<template>
  <div class="isoft_bg_white" style="padding: 10px 10px 0 10px;">
    <Row :gutter="10">
      <Col span="12">
        <Form ref="formValidate" :model="formValidate" :rules="ruleValidate" :label-width="80">
          <FormItem label="链接名称" prop="advertisement_label">
            <Input v-model="formValidate.advertisement_label" placeholder="Enter advertisement_label..."/>
          </FormItem>
          <FormItem label="链接类型" prop="linked_type">
            <Input v-model="formValidate.linked_type" placeholder="Enter linked_type..."></Input>
          </FormItem>
          <FormItem label="链接地址" prop="linked_refer">
            <Input v-model="formValidate.linked_refer" placeholder="Enter linked_refer..."></Input>
          </FormItem>
          <FormItem label="链接图片" prop="linked_img">
            <Input v-model="formValidate.linked_img" placeholder="点击右边图标选择图片"
                   readonly="readonly" icon="ios-document" @on-click="editLinkedImg"></Input>
            <IFileUpload ref="fileUpload" :show-button="false" @uploadComplete="uploadComplete"
                         action="/api/iwork/httpservice/fileUpload" uploadLabel="上传链接图片"/>
          </FormItem>
          <FormItem>
            <Button type="success" size="small" @click="handleSubmit('formValidate')">Submit</Button>
            <Button type="warning" size="small" @click="handleReset('formValidate')">Cancel</Button>
          </FormItem>
        </Form>
      </Col>
      <Col span="12">
        这块再来点别的东西
      </Col>
    </Row>
  </div>
</template>

<script>
  import {EditAdvertisement,QueryAdvertisementById} from "../../api"
  import {handleSpecial} from "../../tools"
  import IFileUpload from "../Common/file/IFileUpload"

  export default {
    name: "EditAdvertisement",
    components: {IFileUpload},
    data(){
      return {
        formValidate: {
          id:-1,
          advertisement_label: '',
          linked_type: '',
          linked_refer: '',
          linked_img: '',
        },
        ruleValidate: {
          advertisement_label: [
            { required: true, message: '链接名称不能为空', trigger: 'blur' }
          ],
          linked_type: [
            { required: true, message: '链接类型不能为空', trigger: 'blur' }
          ],
          linked_refer: [
            { required: true, message: '链接地址不能为空', trigger: 'blur' }
          ],
        },
      }
    },
    methods:{
      uploadComplete: function (result) {
        if(result.status == "SUCCESS"){
          this.formValidate.linked_img = result.fileServerPath;
          this.$refs.fileUpload.hideModal();
        }
      },
      editLinkedImg:function(){
        this.$refs.fileUpload.showModal();
      },
      handleSubmit (name) {
        var _this = this;
        this.$refs[name].validate(async (valid) => {
          if (valid) {
            const result = await EditAdvertisement(_this.formValidate.id, _this.formValidate.advertisement_label,
              _this.formValidate.linked_type, _this.formValidate.linked_refer, handleSpecial(_this.formValidate.linked_img));
            if(result.status == "SUCCESS"){
              _this.$Message.success('提交成功!');
              _this.$emit("handleSubmit");
            }else{
              _this.$Message.error('提交失败!');
            }
          } else {
            _this.$Message.error('验证失败!');
          }
        })
      },
      handleReset (name) {
        this.$refs[name].resetFields();
      },
      initData:async function (id) {
        const result = await QueryAdvertisementById(id);
        if(result.status == "SUCCESS"){
          this.formValidate.id = result.advertisement.id;
          this.formValidate.advertisement_label = result.advertisement.advertisement_label;
          this.formValidate.linked_type = result.advertisement.linked_type;
          this.formValidate.linked_refer = result.advertisement.linked_refer;
          this.formValidate.linked_img = result.advertisement.linked_img;
        }else{
          this.$Message.error(result.errorMsg);
        }
      }
    }
  }
</script>

<style scoped>


</style>
