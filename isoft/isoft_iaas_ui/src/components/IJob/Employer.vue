<template>
  <div>
    <div class="isoft_bg_white isoft_pd10">
      <p>公司名称：
        <Input class="input" v-if="editMode" v-model="formInline.corporate_ame"/>
        <span v-else>{{formInline.corporate_ame}}</span>
      </p>
      <p>公司主页：
        <Input class="input" v-if="editMode" v-model="formInline.corporate_site"/>
        <span v-else>{{formInline.corporate_site}}</span>
      </p>
      <p>公司规模：
        <Input class="input" v-if="editMode" v-model="formInline.corporate_size"/>
        <span v-else>{{formInline.corporate_size}}</span>
      </p>
      <p>招聘类型：
        <Input class="input" v-if="editMode" v-model="formInline.job_type"/>
        <span v-else>{{formInline.job_type}}</span>
      </p>

      <Button style="position: relative;float: right;right: 10px;bottom: 35px;" @click="editMode = !editMode">编辑模式</Button>
    </div>

    <div class="isoft_bg_white isoft_pd10 isoft_top10">
      <p>公司简介：北京博思致新互联网科技有限责任公司xxxxxx</p>
      <p>职位简介：xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx</p>
    </div>
  </div>
</template>

<script>
  export default {
    name: "Employer",
    data(){
      return {
        editMode:false,
        formInline: {
          corporate_ame: '待完善',
          corporate_site: '待完善',
          corporate_size: '待完善',
          job_type: '待完善',
        },
      }
    }
  }
</script>

<style scoped>

</style>
