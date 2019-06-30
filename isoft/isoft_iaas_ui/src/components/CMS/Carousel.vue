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

    <Table :columns="columns1" :data="carousels" size="small"></Table>
    <Page :total="total" :page-size="offset" show-total show-sizer :styles="{'text-align': 'center','margin-top': '10px'}"
          @on-change="handleChange" @on-page-size-change="handlePageSizeChange"/>
  </div>
</template>

<script>
  import {FilterCarousels} from "../../api"

  export default {
    name: "Carousel",
    data () {
      return {
        // 当前页
        current_page:1,
        // 总页数
        total:1,
        // 每页记录数
        offset:10,
        // 搜索条件
        search:"",
        carousels: [],
        columns1: [
          {
            title: 'placement',
            key: 'placement',
            width:180
          },
          {
            title: 'title',
            key: 'title',
            width:200
          },
          {
            title: 'content',
            key: 'content',
            width:400
          },
          {
            title: 'linked_refer',
            key: 'linked_refer',
            width:200
          },
        ],
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
      },
      refreshCarouselList:async function () {
        const result = await FilterCarousels(this.offset, this.current_page, this.search);
        if(result.status=="SUCCESS"){
          this.carousels = result.carousels;
          this.total = result.paginator.totalcount;
        }
      },
      handleChange(page){
        this.current_page = page;
        this.refreshCarouselList();
      },
      handlePageSizeChange(pageSize){
        this.offset = pageSize;
        this.refreshCarouselList();
      },
    },
    mounted(){
      this.refreshCarouselList();
    }
  }
</script>

<style scoped>

</style>
