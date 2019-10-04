<template>
  <div style="background: #FFFFFF;padding: 10px;">
    <Row style="background: #FFFFFF;">
      <Col span="6" style="padding: 2px;background: #f8f8f9;">
        <IBeautifulCard title="我的课程空间">
          <div slot="content">
            <p style="padding: 5px 5px 5px 20px;">
              <IBeautifulLink2 @onclick="$router.push({path:'/ilearning/mine/course_space/newCourse'})">我要开课</IBeautifulLink2>
            </p>
            <p style="padding: 5px 5px 5px 20px;">
              <IBeautifulLink2 @onclick="$router.push({path:'/ilearning/mine/course_space/myCourseList'})">我的课程</IBeautifulLink2>
            </p>
            <p style="padding: 5px 5px 5px 20px;">
              <IBeautifulLink2 @onclick="$router.push({path:'/ilearning/mine/course_space/recentlyViewed'})">最近浏览</IBeautifulLink2>
            </p>
          </div>
        </IBeautifulCard>
      </Col>
      <Col span="18">
        <div style="padding: 2px;background: #f8f8f9;margin-left: 10px;">
          <div style="padding: 10px;background: #FFFFFF;min-height: 500px;">
            <router-view/>
          </div>
        </div>
      </Col>
    </Row>
  </div>
</template>

<script>
  import IBeautifulLink2 from "../../Common/link/IBeautifulLink2";
  import IBeautifulCard from "../../Common/card/IBeautifulCard";
  import RecentlyViewed from "./RecentlyViewed";
  export default {
    name: "CourseSpace",
    components: {RecentlyViewed, IBeautifulCard, IBeautifulLink2},
  }
</script>

<style scoped>

</style>
