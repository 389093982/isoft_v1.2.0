<template>
  <div>
    <Form ref="formInline" :model="formInline" :rules="ruleInline" :label-width="100">
      <Row>
        <Col span="12">
          <FormItem label="占位符">
            <Select v-model="formInline.placement" style="width: 80%;">
              <Option value="beijing">New York</Option>
              <Option value="shanghai">London</Option>
              <Option value="shenzhen">Sydney</Option>
            </Select>
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
            <IFileUpload @uploadComplete="uploadComplete" action="/api/cms/fileUpload/" uploadLabel="上传"/>
          </FormItem>
          <FormItem prop="linked_refer"  label="链接关键词">
            <Input type="text" v-model="formInline.linked_refer" placeholder="linked_refer" style="width: 80%;"/>
          </FormItem>
          <FormItem>
            <Button type="primary" @click="handleSubmit('formInline')">提交</Button>
          </FormItem>
        </Col>
      </Row>
    </Form>

    <Table :columns="columns1" :data="carousels" size="small"></Table>
    <Page :total="total" :page-size="offset" show-total show-sizer :styles="{'text-align': 'center','margin-top': '10px'}"
          @on-change="handleChange" @on-page-size-change="handlePageSizeChange"/>
  </div>
</template>

<script>
  import {FilterCarousels} from "../../api"
  import {AddCarousel} from "../../api"
  import IFileUpload from "../IFile/IFileUpload"

  export default {
    name: "Carousel",
    components:{IFileUpload},
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
            title: 'image_path',
            key: 'image_path',
            width:400
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
          content: [
            { required: true, message: 'Please fill in the content.', trigger: 'blur' },
          ],
          linked_refer: [
            { required: true, message: 'Please fill in the linked_refer.', trigger: 'blur' },
          ]
        }
      }
    },
    methods: {
      handleSubmit(name) {
        this.$refs[name].validate(async (valid) => {
          if (valid) {
            const result = await AddCarousel(this.formInline.placement, this.formInline.title, this.formInline.content,
              this.formInline.imgpath, this.formInline.linked_refer);
            if(result.status=="SUCCESS"){
             this.refreshCarouselList();
            }
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
      uploadComplete: function (result) {
        if(result.status == "SUCCESS"){
          this.formInline.imgpath = result.filepath;
        }
      }
    },
    mounted(){
      this.refreshCarouselList();
    }
  }
</script>

<style scoped>

</style>
