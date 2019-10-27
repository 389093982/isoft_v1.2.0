<template>
  <div style="margin: 20px;">
    <Form ref="formValidate" :model="formValidate" :rules="ruleValidate" :label-width="80">
      <FormItem label="分类名称" prop="catalog_name">
          <Input v-model="formValidate.catalog_name" placeholder="Enter catalog name..."/>
      </FormItem>
      <FormItem label="分类简介" prop="catalog_desc">
        <Input v-model="formValidate.catalog_desc" type="textarea" :rows="4" placeholder="Enter catalog desc..."></Input>
      </FormItem>
      <FormItem>
        <Button type="success" size="small" @click="handleSubmit('formValidate')">Submit</Button>
        <Button type="warning" size="small" @click="handleReset('formValidate')">Cancel</Button>
      </FormItem>
    </Form>
  </div>
</template>

<script>
  import {BlogCatalogEdit} from "../../api"

  export default {
    name: "BlogCatalogEdit",
    data () {
      return {
        formValidate: {
          catalog_name: '',
          catalog_desc: '',
        },
        ruleValidate: {
          catalog_name: [
            { required: true, message: '分类名称不能为空', trigger: 'blur' }
          ],
          catalog_desc: [
            { required: true, message: '分类描述不能为空', trigger: 'blur' }
          ],
        },
      }
    },
    methods:{
      handleSubmit (name) {
        var _this = this;
        this.$refs[name].validate(async (valid) => {
          if (valid) {
            const result = await BlogCatalogEdit(_this.formValidate.catalog_name, _this.formValidate.catalog_desc);
            if(result.status == "SUCCESS"){
              _this.$Message.success('提交成功!');
              this.$emit("handleSuccess");
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
      }
    }
  }
</script>

<style scoped>

</style>
