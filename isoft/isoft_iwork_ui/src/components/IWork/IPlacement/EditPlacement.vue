<template>
  <div>
    <Button type="success" size="small" @click="showRemark = !showRemark">显示备注</Button>

    <Form ref="formInline" :model="formInline" :rules="ruleInline" :label-width="100">
      <Row>
        <Col span="12">
          <div class="remark" v-show="showRemark">备注：placement_name:仅做命名区分,不用于界面展示</div>
          <FormItem prop="placement_name" label="占位符名称">
            <Input type="text" v-model="formInline.placement_name" placeholder="placement_name"/>
          </FormItem>
          <div class="remark" v-show="showRemark">备注：placement_label:界面块元素顶级名称,用于界面展示</div>
          <FormItem prop="placement_label" label="页面显示名称">
            <Input type="text" v-model="formInline.placement_label" placeholder="placement_label"/>
          </FormItem>
          <FormItem prop="placement_type" label="占位符类型">
            <Select v-model.trim="formInline.placement_type" placeholder="placement_type">
              <Option value="all">默认类型 -- all</Option>
              <Option value="text_link">文字链接 -- text_link</Option>
              <Option value="text_event">文字事件 -- text_event</Option>
              <Option value="img_text_link">图文链接 -- img_text_link</Option>
              <Option value="img_text_event">图文事件 -- img_text_event</Option>
            </Select>
          </FormItem>
        </Col>
        <Col span="12">
          <div class="remark" v-show="showRemark">备注：placement_desc:仅做说明,不用于界面展示</div>
          <FormItem prop="placement_desc" label="占位符描述">
            <Input type="text" v-model="formInline.placement_desc" placeholder="placement_desc"/>
          </FormItem>
          <div class="remark" v-show="showRemark">备注：element_limit:占位符元素最大显示数量,小于1默认1000</div>
          <FormItem prop="element_limit" label="元素显示数量">
            <Input type="number" v-model="formInline.element_limit" placeholder="element_limit"/>
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
  import {EditPlacement,QueryPlacementById} from "../../../api"

  export default {
    name: "EditPlacement",
    data(){
      const _validatePlacementName = (rule, value, callback) =>  {
        if (value === '') {
          callback(new Error('占位符名称不能为空!'));
        } else if (value.indexOf("placement_") < 0) {
          callback(new Error('占位符名称必须以 placement_ 开头!'));
        } else {
          callback();
        }
      };
      return {
        showRemark:true,
        formInline: {
          placement_name:'',
          placement_desc:'',
          placement_label:'',
          element_limit:'-1',
          placement_type:'',
        },
        ruleInline: {
          placement_name: [
            { required: true, validator: _validatePlacementName,  trigger: 'blur' },
          ],
          placement_desc: [
            { required: true, message: 'Please fill in the placement_desc.', trigger: 'blur' },
          ],
          placement_label: [
            { required: true, message: 'Please fill in the placement_label.', trigger: 'blur' },
          ],
        }
      }
    },
    methods:{
      handleSubmit() {
        this.$refs['formInline'].validate(async (valid) => {
          if (valid) {
            let id = this.$route.query.id == undefined ? 0 : this.$route.query.id;
              const result = await EditPlacement(id, this.formInline.placement_name, this.formInline.placement_desc,
                this.formInline.placement_label, this.formInline.element_limit, this.formInline.placement_type);
              if(result.status == "SUCCESS"){
                this.$Message.success('提交成功!');
              }else{
                this.$Message.error(result.errorMsg);
              }
            } else {
            this.$Message.error('校验不通过!');
          }
        })
      },
      handleGoBack:function(){
        this.$router.push({ path: '/iwork/placementList'});
      },
      refreshPlacement:async function (id) {
        const result = await QueryPlacementById(id);
        if(result.status == "SUCCESS"){
          let placement = result.placement;
          this.formInline.placement_name = placement.placement_name;
          this.formInline.placement_desc = placement.placement_desc;
          this.formInline.placement_label = placement.placement_label;
          this.formInline.placement_type = placement.placement_type;
          this.formInline.element_limit = placement.element_limit;
        }
      }
    },
    mounted(){
      if(this.$route.query.id != undefined && this.$route.query.id > 0){
        this.refreshPlacement(this.$route.query.id);
      }
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
