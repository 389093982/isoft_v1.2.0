<template>
  <div>
    <Form ref="formInline" :model="formInline" :rules="ruleInline">
      <FormItem prop="task_name">
        <Input type="text" v-model="formInline.task_name" placeholder="task_name">
        </Input>
      </FormItem>
      <FormItem prop="task_desc">
        <Input type="textarea" :rows="5" v-model="formInline.task_desc" placeholder="task_desc">
        </Input>
      </FormItem>
      <FormItem>
        <Button type="success" @click="handleSubmit('formInline')" :style="{'float':'right'}">确认</Button>
      </FormItem>
    </Form>
  </div>
</template>

<script>
  export default {
    name: "AuditEdit",
    data () {
      return {
        formInline: {
          task_name: '',
          task_desc: ''
        },
        ruleInline: {
          task_name: [
            { required: true, message: 'Please fill in the task_name', trigger: 'blur' }
          ],
          task_desc: [
            { required: true, message: 'Please fill in task_desc.', trigger: 'blur' },
          ]
        }
      }
    },
    methods: {
      handleSubmit(name) {
        this.$refs[name].validate((valid) => {
          if (valid) {
            this.$emit("handleSucesss", this.formInline.task_name, this.formInline.task_desc);
            this.$refs[name].resetFields();
          } else {
            this.$Message.error('校验失败!');
          }
        })
      }
    }
  }
</script>

<style scoped>

</style>
