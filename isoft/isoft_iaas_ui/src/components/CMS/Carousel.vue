<template>
  <div>
    <Form ref="formInline" :model="formInline" :rules="ruleInline" :label-width="100">
      <Row>
        <Col span="12">
          <FormItem label="占位符">
            <Select v-model="formInline.select" style="width: 80%;">
              <Option value="beijing">New York</Option>
              <Option value="shanghai">London</Option>
              <Option value="shenzhen">Sydney</Option>
            </Select>
          </FormItem>
          <FormItem prop="user" label="图片">
            <Input type="text" v-model="formInline.user" placeholder="Username" style="width: 80%;"/>
          </FormItem>
          <FormItem prop="password"  label="标题">
            <Input type="password" v-model="formInline.password" placeholder="Password" style="width: 80%;"/>
          </FormItem>
        </Col>
        <Col span="12">
          <FormItem prop="password"  label="内容">
            <Input type="password" v-model="formInline.password" placeholder="Password" style="width: 80%;"/>
          </FormItem>
          <FormItem prop="password"  label="图片">
            <Input type="password" v-model="formInline.password" placeholder="Password" style="width: 80%;"/>
          </FormItem>
          <FormItem prop="password"  label="链接关键词">
            <Input type="password" v-model="formInline.password" placeholder="Password" style="width: 80%;"/>
          </FormItem>
        </Col>
      </Row>
      <FormItem>
        <Button type="primary" @click="handleSubmit('formInline')">Signin</Button>
      </FormItem>
    </Form>
  </div>
</template>

<script>
  export default {
    name: "Carousel",
    data () {
      return {
        formInline: {
         select:'',
          user: '',
          password: ''
        },
        ruleInline: {
          user: [
            { required: true, message: 'Please fill in the user name', trigger: 'blur' }
          ],
          password: [
            { required: true, message: 'Please fill in the password.', trigger: 'blur' },
            { type: 'string', min: 6, message: 'The password length cannot be less than 6 bits', trigger: 'blur' }
          ]
        }
      }
    },
    methods: {
      handleSubmit(name) {
        this.$refs[name].validate((valid) => {
          if (valid) {
            this.$Message.success('Success!');
          } else {
            this.$Message.error('Fail!');
          }
        })
      }
    }
  }
</script>

<style scoped>

</style>
