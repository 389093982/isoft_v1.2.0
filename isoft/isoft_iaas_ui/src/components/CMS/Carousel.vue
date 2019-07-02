<template>
  <div>
    <Chooser ref="placement_chooser">
      <placement :chooser="true" @choosePlacement="choosePlacement"/>
    </Chooser>

    <Form ref="formInline" :model="formInline" :rules="ruleInline" :label-width="100">
      <Row>
        <Col span="12">
          <FormItem label="占位符">
            <Input type="text" readonly="readonly" v-model="formInline.placement" placeholder="placement" style="width: 80%;"/>
            <Button type="success" @click="$refs.placement_chooser.showModal()">选择</Button>
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
  import {UpdateCarouselStatus} from "../../api"
  import IFileUpload from "../IFile/IFileUpload"
  import Chooser from "./Chooser"
  import Placement from "./Placement"
  import MultiClickButton from "../Common/button/MultiClickButton"

  export default {
    name: "Carousel",
    components:{IFileUpload,Chooser,Placement,MultiClickButton},
    data () {
      var _this = this;
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
            title: 'status',
            key: 'status',
            width:100,
            render: (h,params)=> {
              return h('div', {
                style:{
                  color: this.carousels[params.index].status == 1 ?  'green' : (this.carousels[params.index].status == 0 ? 'red' : 'grey'),
                }
              },
              this.carousels[params.index].status == 1 ?  '启用' : (this.carousels[params.index].status == 0 ? '停用' : '失效'))
            }
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
          {
            title: '操作',
            key: 'operate',
            width:250,
            render: (h,params)=> {
              return h('div',[
                h(MultiClickButton,{
                  props:{
                    btnCounts: 3,
                    btnTypes: ['primary','success','warning'],
                    btnShows: [true, true, true],
                    btnBindDatas: [1, 0, -1],
                    btnTexts: ['启用', '停用', '失效'],
                  },
                  on:{
                    handleClick:async function (index, bindData) {
                      alert(index);
                      alert(bindData);
                      const result = await UpdateCarouselStatus(_this.carousels[params.index].id, bindData);
                      if(result.status == "SUCCESS"){
                        _this.refreshCarouselList();
                      }else{
                        _this.$Message.error("状态更新失败!");
                      }
                    }
                  }
                })
              ]);
            }
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
      },
      choosePlacement:function (placement_name) {
        this.formInline.placement = placement_name;
        this.$refs.placement_chooser.hideModal();
      }
    },
    mounted(){
      this.refreshCarouselList();
    }
  }
</script>

<style scoped>

</style>
