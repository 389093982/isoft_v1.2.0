<template>
  <div>
    <EditAdvertisement ref="editAdvertisement" @handleSubmit="handleAdvertisementSubmit"/>

    <div class="isoft_bg_white" style="padding: 10px 10px 0 10px;">
      <IBeautifulCard title="我的广告清单">
        <div slot="content" style="padding: 10px;">
          <div v-if="advertisements && advertisements.length > 0">
            <Row>
              <Col span="4">友情链接名称</Col>
              <Col span="4">链接类型</Col>
              <Col span="4">链接地址</Col>
              <Col span="4">链接图片</Col>
              <Col span="4">联系人</Col>
              <Col span="2">广告状态</Col>
              <Col span="2">操作</Col>
            </Row>
            <Row v-for="(advertisement,index) in advertisements">
              <!-- 加空格是防止没有内容而不占空间 -->
              <Col span="4">{{advertisement.advertisement_label}}&nbsp;</Col>
              <Col span="4">{{advertisement.linked_type}}&nbsp;</Col>
              <Col span="4">{{advertisement.linked_refer}}&nbsp;</Col>
              <Col span="4" class="isoft_inline_ellipsis">
                <span :title="advertisement.linked_img">{{advertisement.linked_img}}</span>&nbsp;
              </Col>
              <Col span="4">{{loginUserName}}&nbsp;</Col>
              <Col span="2">已发布</Col>
              <Col span="2"><a @click="editAdvertisement(advertisement.id)">编辑</a></Col>
            </Row>
          </div>
        </div>
      </IBeautifulCard>
    </div>
  </div>
</template>

<script>
  import {GetPersonalAdvertisement} from "../../api"
  import {GetLoginUserName} from "../../tools"
  import EditAdvertisement from "./EditAdvertisement"
  import IBeautifulCard from "../Common/card/IBeautifulCard"

  export default {
    name: "Manage",
    components: {IBeautifulCard, EditAdvertisement},
    data(){
      return {
        advertisements:null,
      }
    },
    methods:{
      handleAdvertisementSubmit:function(){
        this.refreshPersonalAdvertisement();
      },
      refreshPersonalAdvertisement:async function () {
        const result = await GetPersonalAdvertisement();
        if(result.status == "SUCCESS"){
          this.advertisements = result.advertisements;
        }else{
          this.$Message.error(result.errorMsg);
        }
      },
      editAdvertisement:function (id) {
        this.$refs.editAdvertisement.initData(id);
      }
    },
    computed:{
      loginUserName(){
        return GetLoginUserName();
      }
    },
    mounted(){
      this.refreshPersonalAdvertisement();
    }
  }
</script>

<style scoped>


</style>
