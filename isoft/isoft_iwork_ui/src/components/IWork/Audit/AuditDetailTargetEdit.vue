<template>
  <div>
    <Tabs :animated="false" name="tab_level_2" style="width: 80%;">
      <TabPane v-for="(item, index) in update_cases" :label="item.case_name ? item.case_name : '场景 ' + (index + 1)" tab="tab_level_2">
        场景名称: <span style="color: #00ce00;">参考案例：生效、失效、审核通过、内容不合法等中文或英文</span>
        <span style="margin-left: 100px;">删除按钮颜色设置：<ColorPicker v-model="item.case_color" size="small" alpha recommend/></span>
        <Button type="error" size="small" @click="handleRemove(index)">删除</Button>
        <Input type="text" v-model="item.case_name" placeholder="请输入场景名称" style="margin: 5px 0;"></Input>
        场景更新sql: <span style="color: #00ce00;">参考案例：update blog set status = 1 where id = :id</span>
        <Input type="textarea" :rows="6" v-model="item.update_sql" placeholder="请输入 update_sql" style="margin: 5px 0;"></Input>
        场景描述: <span style="color: #00ce00;">参考案例：根据 id 更新状态为 1(生效状态)</span>
        <Input type="textarea" :rows="5" v-model="item.update_desc" placeholder="请输入描述" style="margin: 10px 0;"></Input>

      </TabPane>
    </Tabs>
    <div>
      <Button type="success" size="small" @click="handleAdd">新增场景</Button>
      <Button type="info" size="small" @click="handleSubmit">保存场景</Button>
    </div>
  </div>
</template>

<script>
  import {EditAuditTaskTarget,QueryTaskDetail} from "../../../api"

  export default {
    name: "AuditDetailTargetEdit",
    data(){
      return {
        update_cases: [
          {
            case_name:'',
            update_sql:'',
            update_desc:'',
            case_color:'',
          }
        ]
      }
    },
    methods:{
      handleSubmit:async function () {
        const result = await EditAuditTaskTarget(this.$route.query.task_name, JSON.stringify(this.update_cases));
        if(result.status == "SUCCESS"){
          this.$Message.success("保存成功！");
          this.$router.go(0);     // 强制刷新页面
        }else{
          this.$Message.error("保存失败!");
        }
      },
      handleAdd () {
        this.update_cases.push({
          case_name:'',
            update_sql: '',
            update_desc: '',
          case_color: '#19be6b',
        });
      },
      handleRemove (index) {
        // 删除一个数组元素
        this.update_cases.splice(index, 1);
      },
      refreshAuditDetail:async function () {
        const result = await QueryTaskDetail(this.$route.query.task_name);
        if(result.status == "SUCCESS"){
          this.update_cases = result.taskDetail.update_cases;
          if(this.update_cases == null){
            this.update_cases = [];
            this.handleAdd();
          }
        }
      }
    },
    mounted(){
      this.refreshAuditDetail();
    }
  }
</script>

<style scoped>

</style>
