<template>
  <Row :gutter="20" style="padding-top: 20px;">
    <Col span="12" v-for="history in historys">
      <IBeautifulLink2 @onclick="$router.push({path:'/ilearning/course_detail',query:{course_id:history.history_link}})">
        {{history.history_desc}}
      </IBeautifulLink2>
      <span style="float: right;font-size: 12px;"><Time :time="history.last_updated_time" :interval="1"/></span>
      <Divider/>
    </Col>
  </Row>
</template>

<script>
  import {ShowCourseHistory} from "../../../api"
  import IBeautifulLink2 from "../../Common/link/IBeautifulLink2";

  export default {
    name: "RecentlyViewed",
    components: {IBeautifulLink2},
    data(){
      return {
        current_page:1,
        offset:10,
        total:0,
        historys:[],
      }
    },
    methods:{
      searchFunc:function (data) {
        this.$router.push({ path: '/ilearning/course_search', query: { search: data }});
      },
      async refreshRecentlyViewed(){
        const data = await ShowCourseHistory(this.offset, this.current_page);
        if(data.status == "SUCCESS"){
          this.historys = data.historys;
          this.total = data.paginator.totalcount;
        }
      },
    },
    mounted:function () {
      this.refreshRecentlyViewed();
    }
  }
</script>

<style scoped>

</style>
