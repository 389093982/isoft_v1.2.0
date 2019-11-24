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
  import {EditAuditTask} from "../../../api"

  export default {
    name: "AuditTaskEdit",
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
        this.$refs[name].validate(async (valid) => {
          if (valid) {
            const result = await EditAuditTask(this.formInline.task_name, this.formInline.task_desc);
            if(result.status == "SUCCESS"){
              this.$Message.success('提交成功!');
              this.$emit("handleSucess");
              this.$refs[name].resetFields();
            }else{
              this.$Message.error('提交失败!');
            }
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
