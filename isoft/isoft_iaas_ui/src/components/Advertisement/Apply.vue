<template>
  <div>
    <div class="isoft_bg_white" style="padding: 10px 10px 0 10px;">
      <p>广告投放说明</p>
      <p class="isoft_font12">1、用户可以在本站点发布广告，每个广告位对应一个链接地址，跳往用户指定站点;</p>
      <p class="isoft_font12">不接受违法、色情、赌博类站点的广告，一经发现永久查封，欢迎大家举报;</p>
      <p class="isoft_font12">广告位租出后不接受退订，未到期不能更换广告位;</p>
      <p class="isoft_font12">广告到期，续租广告者，请及时联系管理员;</p>
      <Button type="success" @click="manageAdvertisement"
              style="position: relative;float:right;right: 20px;bottom: 50px;">管理我的广告</Button>
    </div>

    <div class="isoft_bg_white" style="margin: 5px 0 5px 0;padding: 10px 10px 0 10px;">
      热门广告列表
    </div>

    <div class="isoft_bg_white" style="margin: 5px 0 5px 0;padding: 10px 10px 0 10px;">
      广告费用说明
    </div>
  </div>
</template>

<script>
  import {CheckHasLoginConfirmDialog} from "../../tools"

  export default {
    name: "Apply",
    data(){
      return {

      }
    },
    methods:{
      manageAdvertisement:function () {
        CheckHasLoginConfirmDialog(this, "/advertisement/manage");
      }
    }
  }
</script>

<style scoped>


</style>
