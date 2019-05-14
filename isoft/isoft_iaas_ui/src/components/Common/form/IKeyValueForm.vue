<template>
  <!-- 表单信息 -->
  <Form ref="formValidate" :model="formValidate" :rules="ruleValidate" :label-width="150">
    <FormItem :label="formKeyLabel" prop="formkey">
      <Input v-model.trim="formValidate.formkey" :placeholder="formKeyPlaceholder"></Input>
    </FormItem>
    <FormItem :label="formValueLabel" prop="formvalue">
      <Input v-model.trim="formValidate.formvalue" type="textarea" :rows="4" :placeholder="formValuePlaceholder"></Input>
    </FormItem>
    <FormItem>
      <Button type="success" @click="handleSubmit('formValidate')" style="margin-right: 6px">Submit</Button>
      <Button type="warning" @click="handleReset('formValidate')" style="margin-right: 6px">Reset</Button>
    </FormItem>
  </Form>
</template>

<script>
  export default {
    name: "IKeyValueForm",
    props:{
      formKeyLabel:{
        type:String,
        default:'formKeyLabel',
      },
      formValueLabel:{
        type:String,
        default:'formValueLabel',
      },
      formKeyPlaceholder:{
        type:String,
        default:'请输入 formKey',
      },
      formValuePlaceholder:{
        type:String,
        default:'请输入 formValue',
      },
    },
    data(){
      return {
        formValidate: {
          formmode: null,    // 用户回显使用
          formkey: '',
          formvalue: '',
        },
        ruleValidate: {
          formkey: [
            { required: true, message: 'formkey 不能为空!', trigger: 'blur' }
          ],
          formvalue: [
            { required: true, message: 'formvalue 不能为空!', trigger: 'blur' }
          ],
        },
      }
    },
    methods:{
      handleSubmit (name) {
        this.$refs[name].validate((valid) => {
          if (valid) {
            this.$emit("handleSubmit", this.formValidate.formmode, this.formValidate.formkey, this.formValidate.formvalue);
          }
        })
      },
      handleSubmitSuccess (msg){
        this.$Message.success(msg);
      },
      handleSubmitError (msg){
        this.$Message.error(msg);
      },
      initFormData (formmode, formkey, formvalue){
        this.formValidate.formmode = formmode;
        this.formValidate.formkey = formkey;
        this.formValidate.formvalue = formvalue;
      },
      handleReset (name) {
        this.$refs[name].resetFields();
      },
    }
  }
</script>

<style scoped>

</style>
