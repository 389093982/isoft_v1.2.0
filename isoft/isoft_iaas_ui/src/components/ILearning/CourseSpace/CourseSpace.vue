<template>
  <div style="background: #FFFFFF;padding: 10px;">
    <Row style="background: #FFFFFF;">
      <Col span="6" style="padding: 2px;background: #f8f8f9;">
        <Card title="我的课程空间">
          <p style="padding: 5px;">
            <IBeautifulLink2 @onclick="$router.push({path:'/ilearning/mine/course_space/newCourse'})">我要开课</IBeautifulLink2>
          </p>
          <p style="padding: 5px;">
            <IBeautifulLink2 @onclick="$router.push({path:'/ilearning/mine/course_space/myCourseList'})">我的课程</IBeautifulLink2>
          </p>
          <p style="padding: 5px;">
            <IBeautifulLink2 @onclick="$router.push({path:'/ilearning/mine/course_space/recentlyViewed'})">最近浏览</IBeautifulLink2>
          </p>
          <p style="padding: 5px;"><IBeautifulLink2 @onclick="">我的收藏</IBeautifulLink2></p>
        </Card>
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
  export default {
    name: "CourseSpace",
    components: {IBeautifulLink2},
  }
</script>

<style scoped>

</style>
