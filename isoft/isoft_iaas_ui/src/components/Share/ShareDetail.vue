<template>
  <div v-if="share">
    <IShowMarkdown v-if="share.content" :content="share.content"/>

    <hr>
    <!-- 评论模块 -->
    <IEasyComment :theme_pk="share.id" theme_type="share_theme_type" style="margin-top: 50px;"/>
  </div>
</template>

<script>
  import {ShowShareDetail} from "../../api"
  import IShowMarkdown from "../Common/markdown/IShowMarkdown"
  import IEasyComment from "../Comment/IEasyComment"

  export default {
    name: "ShareDetail",
    components:{IShowMarkdown,IEasyComment},
    data(){
      return {
        share:null,
      }
    },
    methods:{
      refreshShareDetail:async function () {
        const result = await ShowShareDetail(this.$route.query.share_id);
        if(result.status=="SUCCESS"){
          this.share = result.share;
        }
      }
    },
    mounted:function () {
      this.refreshShareDetail();
    }
  }
</script>

<style scoped>

</style>
