<template>
  <div>
    <IBaseChooser ref="placement_chooser" chooser-title="占位符选择">
      <Placement :chooserMode="true" @choosePlacement="choosePlacement"/>
    </IBaseChooser>

    <Form ref="formInline" :model="formInline" :rules="ruleInline" :label-width="100">
      <Row>
        <Col span="12">
          <FormItem label="占位符">
            <Input type="text" readonly="readonly" v-model="formInline.placement" placeholder="placement" style="width: 80%;"/>
            <Button type="success" @click="$refs.placement_chooser.showModal()">选择</Button>
          </FormItem>
          <FormItem prop="navigation_level" label="导航级别">
            <Input type="text" readonly="readonly" v-model="formInline.navigation_level" placeholder="navigation_level" style="width: 80%;"/>
          </FormItem>
          <FormItem prop="navigation_parent_id"  label="父级关联 id">
            <Input type="text" readonly="readonly" v-model="formInline.navigation_parent_id" placeholder="navigation_parent_id" style="width: 80%;"/>
            <Poptip v-model="visible_choose_element" placement="left-start" width="400" @on-popper-show="showChooseElement">
              <a href="javascript:;">选择父级</a>
              <div slot="content" style="width: 100%;">
                <Tag style="margin: 5px;float: left;" v-for="element in elements">
                  <span @click="chooseElement(element)">{{element.title}}</span>
                </Tag>
              </div>
            </Poptip>
          </FormItem>
        </Col>
        <Col span="12">
          <FormItem prop="title" label="标题">
            <Input type="text" v-model="formInline.title" placeholder="title" style="width: 80%;"/>
          </FormItem>
          <FormItem prop="content"  label="内容">
            <Input type="text" v-model="formInline.content" placeholder="content" style="width: 80%;"/>
          </FormItem>
          <FormItem prop="imgpath"  label="图片">
            <Input type="text" readonly="readonly" v-model="formInline.imgpath" placeholder="imgpath" style="width: 80%;"/>
            <IFileUpload @uploadComplete="uploadComplete" action="/api2/iwork/fileUpload/fileUpload" uploadLabel="上传"/>
          </FormItem>
          <FormItem prop="linked_refer"  label="链接关键词">
            <Input type="text" v-model="formInline.linked_refer" placeholder="linked_refer" style="width: 80%;"/>
          </FormItem>
          <FormItem>
            <Button type="success" @click="handleSubmit" style="margin-right: 6px">提交</Button>
            <Button type="warning" @click="handleGoBack" style="margin-right: 6px">返回</Button>
          </FormItem>
        </Col>
      </Row>
    </Form>
  </div>
</template>

<script>
  import IBaseChooser from "../../Common/IBaseChooser"
  import Placement from "./Placement"
  import Element from "./Element"
  import IFileUpload from "../../IFile/IFileUpload"
  import {EditElement,FilterElementByPlacement,QueryElementById} from "../../../api"
  import {checkEmpty} from "../../../tools"

  export default {
    name: "EditElement",
    components:{IBaseChooser,Placement,Element,IFileUpload},
    data(){
      return {
        visible_choose_element:false,
        elements:[],
        formInline: {
          placement:'',
          navigation_level:0,  // 元素层级
          navigation_parent_id:0,   // 父级元素 id
          title: '',
          content: '',
          imgpath: '',
          linked_refer: '',
        },
        ruleInline: {
          title: [
            { required: true, message: 'Please fill in the title.', trigger: 'blur' },
          ],
        }
      }
    },
    methods:{
      handleSubmit() {
        this.$refs['formInline'].validate(async (valid) => {
          if (valid) {
            let id = this.$route.query.id == undefined ? -1 : this.$route.query.id;
            const result = await EditElement(id, this.formInline.placement, this.formInline.navigation_level,
              this.formInline.navigation_parent_id, this.formInline.title, this.formInline.content,
              this.formInline.imgpath, this.formInline.linked_refer);
            if(result.status=="SUCCESS"){
              this.$Message.success('提交成功!');
            }else{
              this.$Message.error('提交失败!' + result.errorMsvg);
            }
          } else {
            this.$Message.error('校验不通过!');
          }
        })
      },
      handleGoBack:function(){
        let search = checkEmpty(this.formInline.placement) ? "" : this.formInline.placement;
        this.$router.push({ path: '/background/cms/element_list', query: { search: search }});
      },
      choosePlacement:function (placement_name) {
        this.formInline.placement = placement_name;
        this.$refs.placement_chooser.hideModal();
      },
      chooseElement:function(element){
        this.formInline.navigation_level = element.navigation_level + 1;
        this.formInline.navigation_parent_id = element.id;
        this.visible_choose_element = false;
      },
      showChooseElement:async function(){
        if(!checkEmpty(this.formInline.placement)){
          const result = await FilterElementByPlacement(this.formInline.placement);
          if(result.status == "SUCCESS"){
            this.elements = result.elements;
          }
        }
      },
      uploadComplete: function (result) {
        if(result.status == "SUCCESS"){
          this.formInline.imgpath = result.fileServerPath;
        }
      },
      refreshElement:async function (id) {
        const result = await QueryElementById(id);
        if(result.status == "SUCCESS"){
          let element = result.element;
          this.formInline.placement = element.placement;
          this.formInline.title = element.title;
          this.formInline.content = element.content;
          this.formInline.imgpath = element.img_path;
          this.formInline.linked_refer = element.linked_refer;
        }
      }
    },
    mounted(){
      if(this.$route.query.id != undefined && this.$route.query.id > 0){
        this.refreshElement(this.$route.query.id);
      }
    }
  }
</script>

<style scoped>

</style>
