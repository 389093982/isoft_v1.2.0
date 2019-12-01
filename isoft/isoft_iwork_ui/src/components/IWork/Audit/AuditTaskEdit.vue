<template>
  <div>
    <Form ref="formInline" :model="formInline" :rules="ruleInline" :label-width="100">
      <FormItem prop="task_name" label="任务名称">
        <Input type="text" v-model.trim="formInline.task_name" placeholder="task_name">
        </Input>
      </FormItem>
      <FormItem prop="task_desc" label="任务描述">
        <Input type="textarea" :rows="5" v-model.trim="formInline.task_desc" placeholder="task_desc">
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
  import {startsWith} from "../../../tools"

  export default {
    name: "AuditTaskEdit",
    data () {
      const validateTaskName = (rule, value, callback) => {
        if (value === '') {
          callback(new Error("任务名称不能为空!"));
        } else if (!startsWith(value, "task_audit_")) {
          callback(new Error("任务名称必须以 'task_audit_' 开头!"));
        } else if (value.length < 10) {
          callback(new Error("任务名称太短,必须大于 10 个字符!"));
        } else {
          callback();
        }
      };
      return {
        formInline: {
          task_name: '',
          task_desc: ''
        },
        ruleInline: {
          task_name: [
            { required: true, message: '请输入任务名称', trigger: 'blur' },
            { validator: validateTaskName, trigger: 'blur' }
          ],
          task_desc: [
            { required: true, message: '请输入任务描述', trigger: 'blur' },
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
