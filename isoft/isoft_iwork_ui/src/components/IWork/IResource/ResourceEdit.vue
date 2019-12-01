<template>
  <!-- 按钮触发模态框 -->
  <!-- ref 的作用是为了在其它地方方便的获取到当前子组件 -->
  <ISimpleBtnTriggerModal ref="triggerModal" btn-text="新增资源信息" btn-size="small" modal-title="新增/编辑资源信息" modal-top="50px" :modal-width="600">
    <!-- 表单信息 -->
    <Form ref="formValidate" :model="formValidate" :rules="ruleValidate" :label-width="140">
      <FormItem label="resource_name" prop="resource_name">
        <Input v-model.trim="formValidate.resource_name" placeholder="请输入 resource_name"></Input>
      </FormItem>
      <FormItem label="resource_type" prop="resource_type">
        <Select v-model="formValidate.resource_type" placeholder="请选择 resource_type">
          <Option value="db">db</Option>
          <Option value="sftp">sftp</Option>
          <Option value="ssh">ssh</Option>
        </Select>
      </FormItem>
      <FormItem label="resource_url" prop="resource_url">
        <Input v-model.trim="formValidate.resource_url" placeholder="请输入 resource_url"></Input>
      </FormItem>
      <FormItem label="resource_dsn" prop="resource_dsn">
        <Input v-model.trim="formValidate.resource_dsn" placeholder="请输入 resource_dsn"></Input>
      </FormItem>
      <FormItem label="resource_username" prop="resource_username">
        <Input v-model.trim="formValidate.resource_username" placeholder="请输入 resource_username"></Input>
      </FormItem>
      <FormItem label="resource_password" prop="resource_password">
        <Input v-model.trim="formValidate.resource_password" placeholder="请输入 resource_password"></Input>
      </FormItem>
      <FormItem>
        <Button type="success" @click="handleSubmit('formValidate')" style="margin-right: 6px">Submit</Button>
        <Button type="warning" @click="handleReset('formValidate')" style="margin-right: 6px">Reset</Button>
      </FormItem>
    </Form>
  </ISimpleBtnTriggerModal>
</template>

<script>
  import ISimpleBtnTriggerModal from "../../Common/modal/ISimpleBtnTriggerModal"
  import {EditResource,GetResourceById} from "../../../api"

  export default {
    name: "ResourceEdit",
    components:{ISimpleBtnTriggerModal},
    data(){
      return {
        formValidate: {
          resource_id:-1,
          resource_name: '',
          resource_type: '',
          resource_url: '',
          resource_dsn: '',
          resource_username: '',
          resource_password: '',
          env_name: '',
        },
        ruleValidate: {
          resource_name: [
            { required: true, message: 'resource_name 不能为空!', trigger: 'blur' }
          ],
          resource_type: [
            { required: true, message: 'resource_type 不能为空!', trigger: 'change' }
          ],
          resource_url: [
            { required: true, message: 'resource_url 不能为空!', trigger: 'blur' }
          ],
          resource_dsn: [
            { required: true, message: 'resource_dsn 不能为空!', trigger: 'blur' }
          ],
          resource_username: [
            { required: true, message: 'resource_username 不能为空!', trigger: 'blur' }
          ],
          resource_password: [
            { required: true, message: 'resource_password 不能为空!', trigger: 'blur' }
          ],
        },
      }
    },
    methods:{
      initData:async function(resource_id){
        this.$refs.triggerModal.triggerClick();
        const result = await GetResourceById(resource_id);
        if(result.status == "SUCCESS"){
          this.formValidate.resource_id =result.resource.id;
          this.formValidate.resource_name =result.resource.resource_name;
          this.formValidate.resource_type =result.resource.resource_type;
          this.formValidate.resource_url =result.resource.resource_url;
          this.formValidate.resource_dsn =result.resource.resource_dsn;
          this.formValidate.resource_username =result.resource.resource_username;
          this.formValidate.resource_password =result.resource.resource_password;
          this.formValidate.env_name =result.resource.env_name;
        }
      },
      handleSubmit (name) {
        this.$refs[name].validate(async (valid) => {
          if (valid) {
            const result = await EditResource(this.formValidate.resource_id,this.formValidate.resource_name,
              this.formValidate.resource_type,this.formValidate.resource_url,this.formValidate.resource_dsn,
              this.formValidate.resource_username,this.formValidate.resource_password);
            if(result.status == "SUCCESS"){
              this.$Message.success('提交成功!');
              // 调用子组件隐藏 modal (this.refs.xxx.子组件定义的方法())
              this.$refs.triggerModal.hideModal();
              // 通知父组件添加成功
              this.$emit('handleSuccess');
            }else{
              this.$Message.error('提交失败!');
            }
          }
        })
      },
      handleReset (name) {
        this.$refs[name].resetFields();
      },
    }
  }
</script>

<style scoped>

</style>



