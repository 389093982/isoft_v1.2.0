<template>
  <div>
    <Button type="success" size="small" @click="showRemark = !showRemark">显示备注</Button>

    <Button type="warning" size="small" v-if="placement" @click="showPlacementTypeDesc = !showPlacementTypeDesc">
      当前占位符类型:{{placement.placement_type}}
    </Button>
    <p v-for="(value, key) in placement_types" v-show="showPlacementTypeDesc"><span style="display:inline-block;width: 100px;">{{key}}:</span><Tag v-for="(val, index) in value">{{val}}</Tag></p>

    <Form ref="formInline" :model="formInline" :rules="ruleInline" :label-width="100">
      <Row>
        <Col span="12">
          <div class="remark" v-show="showRemark && checkShow('placement_name')">备注：placement_name:仅做命名区分,不用于界面展示</div>
          <FormItem label="占位符" v-show="checkShow('placement_name')">
            <Input type="text" readonly="readonly" v-model="formInline.placement" placeholder="placement"/>
          </FormItem>
          <div class="remark" v-show="showRemark && checkShow('element_label')">备注：element_label:界面块元素名称,用于界面展示</div>
          <FormItem prop="element_label" v-show="checkShow('element_label')" label="页面显示名称">
            <Input type="text" v-model="formInline.element_label" placeholder="element_label"/>
          </FormItem>
          <FormItem prop="navigation_parent_id" v-show="checkShow('navigation_parent_id')" label="父级关联 id">
            <Input type="text" readonly="readonly" v-model="formInline.navigation_parent_id" placeholder="navigation_parent_id" style="width: 80%;"/>
            <Poptip v-model="visible_choose_element" placement="left-start" width="400" @on-popper-show="showChooseElement">
              <a href="javascript:;">选择父级</a>
              <div slot="content" style="width: 100%;">
                <Tag style="margin: 5px;float: left;" v-for="element in elements">
                  <span @click="chooseElement(element)">{{element.element_label}}</span>
                </Tag>
              </div>
            </Poptip>
          </FormItem>
          <FormItem prop="content" v-show="checkShow('content')" label="内容">
            <Input type="textarea" :rows="3" v-model="formInline.content" placeholder="content"/>
          </FormItem>
        </Col>
        <Col span="12">
          <div class="remark" v-show="showRemark && checkShow('element_name')">备注：element_name:仅做命名区分,不用于界面展示</div>
          <FormItem prop="element_name" v-show="checkShow('element_name')" label="元素名称">
            <Input type="text" v-model="formInline.element_name" placeholder="element_name"/>
          </FormItem>
          <FormItem prop="navigation_level" v-show="checkShow('navigation_level')" label="导航级别">
            <Input type="text" readonly="readonly" v-model="formInline.navigation_level" placeholder="navigation_level"/>
          </FormItem>
          <FormItem prop="imgpath" v-show="checkShow('imgpath')" label="图片">
            <Input type="text" readonly="readonly" v-model="formInline.imgpath" placeholder="imgpath" style="width: 80%;"/>
            <IFileUpload @uploadComplete="uploadComplete" action="/api/iwork/httpservice/fileUpload" uploadLabel="上传"/>
          </FormItem>
          <div class="remark" v-show="showRemark && checkShow('linked_refer')">备注：linked_refer:链接地址、链接关键词等</div>
          <FormItem prop="linked_refer" v-show="checkShow('linked_refer')" label="链接信息">
            <Input type="text" v-model="formInline.linked_refer" placeholder="linked_refer"/>
          </FormItem>
        </Col>
      </Row>
      <Row>
        <FormItem prop="md_content" v-show="checkShow('md_content')" label="markdown内容">
          <mavon-editor v-model="formInline.md_content" :toolbars="toolbars" :ishljs = "true" style="z-index: 1;"/>
        </FormItem>
      </Row>
      <Row>
        <FormItem>
          <Button type="success" @click="handleSubmit" style="margin-right: 6px">提交</Button>
          <Button type="warning" @click="handleGoBack" style="margin-right: 6px">返回</Button>
        </FormItem>
      </Row>
    </Form>
  </div>
</template>

<script>
  import IBaseChooser from "../../Common/IBaseChooser"
  import Placement from "./Placement"
  import Element from "./Element"
  import IFileUpload from "../../Common/file/IFileUpload"
  import {EditElement,FilterElementByPlacement,QueryElementById,QueryPlacementByName} from "../../../api"
  import {checkEmpty,oneOf} from "../../../tools"

  export default {
    name: "EditElement",
    components:{IBaseChooser,Placement,Element,IFileUpload},
    data(){
      const _validateElementName = (rule, value, callback) => {
        if (value === '') {
          callback(new Error('元素名称不能为空!'));
        } else if (value.indexOf("element_") < 0) {
          callback(new Error('元素名称必须以 element_ 开头'));
        } else {
          callback();
        }
      };
      return {
        placement_types:{
          'all':['placement_name','element_name','element_label','navigation_level','navigation_parent_id','content','imgpath','linked_refer','md_content'],
          'text_link':['placement_name','element_name','element_label','linked_refer'],
          'text_event':['placement_name','element_name','element_label','navigation_level','navigation_parent_id','content','imgpath','linked_refer','md_content'],
          'img_text_link':['placement_name','element_name','element_label'],
          'img_text_event':['placement_name','element_name','element_label','navigation_level','navigation_parent_id','content','imgpath','linked_refer','md_content'],
        },
        showPlacementTypeDesc:false,
        showRemark:true,
        placement:null,
        toolbars: {
          bold: true, // 粗体
          italic: true, // 斜体
          header: true, // 标题
          underline: true, // 下划线
          // mark: true, // 标记
          superscript: true, // 上角标
          quote: true, // 引用
          ol: true, // 有序列表
          link: true, // 链接
          imagelink: true, // 图片链接
          help: true, // 帮助
          code: true, // code
          subfield: true, // 是否需要分栏
          fullscreen: true, // 全屏编辑
          readmodel: true, // 沉浸式阅读
          undo: true, // 上一步
          trash: true, // 清空
          save: true, // 保存（触发events中的save事件）
          navigation: true // 导航目录
        },
        visible_choose_element:false,
        elements:[],
        formInline: {
          placement:'',
          element_label: '',
          element_name:'',
          navigation_level:0,  // 元素层级
          navigation_parent_id:0,   // 父级元素 id
          content: '',
          imgpath: '',
          linked_refer: '',
          md_content: '',
        },
        ruleInline: {
          element_name: [
            { required: true, validator: _validateElementName,  trigger: 'blur' }
          ],
          element_label: [
            { required: true, message: 'Please fill in the element_label.', trigger: 'blur' },
          ],
        }
      }
    },
    methods:{
      checkShow:function(name){
        return this.placement != null && oneOf(name, this.placement_types[this.placement.placement_type]);
      },
      handleSubmit() {
        this.$refs['formInline'].validate(async (valid) => {
          if (valid) {
            let id = this.$route.query.id == undefined ? 0 : this.$route.query.id;
            const result = await EditElement(id, this.formInline.placement, this.formInline.element_name, this.formInline.navigation_level,
              this.formInline.navigation_parent_id, this.formInline.element_label, this.formInline.content, this.formInline.md_content,
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
        this.$router.push({ path: '/iwork/elementList', query: { placement_name: this.formInline.placement }});
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
          this.formInline.element_name = element.element_name;
          this.formInline.element_label = element.element_label;
          this.formInline.content = element.content;
          this.formInline.md_content = element.md_content;
          this.formInline.imgpath = element.img_path;
          this.formInline.linked_refer = element.linked_refer;
        }
      },
      refreshPlacement:async function (placement_name) {
        const result = await QueryPlacementByName(placement_name);
        if(result.status == "SUCCESS"){
          this.placement = result.placement;
        }
      }
    },
    mounted(){
      this.formInline.placement = this.$route.query.placement_name;
      if(this.$route.query.id != undefined && this.$route.query.id > 0){
        this.refreshElement(this.$route.query.id);
      }
      this.refreshPlacement(this.$route.query.placement_name);
    }
  }
</script>

<style scoped>
  .remark{
    text-align: right;
    color: green;
    font-size: 12px;
  }
</style>
