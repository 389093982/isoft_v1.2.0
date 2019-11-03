<template>
  <div class="header">
    <Carousel autoplay v-model="value" loop dots="inside">
      <CarouselItem>
        <div class="demo-carousel">
          <img class="demo-img" height="250px;" src="../../assets/ilearning/course/1.jpg"/>
        </div>
      </CarouselItem>
      <CarouselItem>
        <div class="demo-carousel">
          <img class="demo-img" height="250px;" src="../../assets/ilearning/course/2.jpg"/>
        </div>
      </CarouselItem>
      <CarouselItem>
        <div class="demo-carousel">
          <img class="demo-img" height="250px;" src="../../assets/ilearning/course/3.jpg">
        </div>
      </CarouselItem>
      <CarouselItem>
        <div class="demo-carousel">
          <img class="demo-img" height="250px;" src="../../assets/ilearning/course/4.jpg">
        </div>
      </CarouselItem>
    </Carousel>
  </div>
</template>

<script>

  export default {
    name: 'IndexCarousel',
    data(){
      return {
        value:0,
      }
    }
  }
</script>

<style lang="stylus" scoped>
  .demo-img
    width: 100%
</style>
