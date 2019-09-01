<template>
  <div>
    <Form ref="formInline" :model="formInline" :rules="ruleInline" :label-width="100">
      <Row>
        <Col span="12">
          <FormItem prop="placement_name" label="占位符名称">
            <Input type="text" v-model="formInline.placement_name" placeholder="placement_name" style="width: 80%;"/>
          </FormItem>
        </Col>
        <Col span="12">
          <FormItem prop="placement_desc" label="占位符描述">
            <Input type="text" v-model="formInline.placement_desc" placeholder="placement_desc" style="width: 80%;"/>
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
      return {
        formInline: {
          placement_name:'',
          placement_desc:'',
        },
        ruleInline: {
          placement_name: [
            { required: true, message: 'Please fill in the placement_name.', trigger: 'blur' },
          ],
          placement_desc: [
            { required: true, message: 'Please fill in the placement_desc.', trigger: 'blur' },
          ],
        }
      }
    },
    methods:{
      handleSubmit() {
        this.$refs['formInline'].validate(async (valid) => {
          if (valid) {
            let id = this.$route.query.id == undefined ? -1 : this.$route.query.id;
              const result = await EditPlacement(id, this.formInline.placement_name, this.formInline.placement_desc);
              if(result.status == "SUCCESS"){
                this.$Message.success('提交成功!');
              }else{
                this.$Message.error('提交失败!');
              }
            } else {
            this.$Message.error('校验不通过!');
          }
        })
      },
      handleGoBack:function(){
        this.$router.push({ path: '/background/cms/placement_list'});
      },
      refreshPlacement:async function (id) {
        const result = await QueryPlacementById(id);
        if(result.status == "SUCCESS"){
          let placement = result.placement;
          this.formInline.placement_name = placement.placement_name;
          this.formInline.placement_desc = placement.placement_desc;
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

</style>
