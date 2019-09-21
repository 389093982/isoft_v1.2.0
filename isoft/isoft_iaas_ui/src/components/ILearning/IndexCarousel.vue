<template>
  <div class="header">
    <Carousel autoplay v-model="value" loop>
      <CarouselItem>
        <div class="demo-carousel">
          <img class="demo-img" height="250px;" src="http://img1.qunarzz.com/piao/fusion/1806/fc/e47aa3e1c67bbc02.jpg_750x200_0f3eecf8.jpg"/>>
        </div>
      </CarouselItem>
      <CarouselItem>
        <div class="demo-carousel">
          <img class="demo-img" height="250px;" src="http://img1.qunarzz.com/piao/fusion/1806/fc/e47aa3e1c67bbc02.jpg_750x200_0f3eecf8.jpg"/>>
        </div>
      </CarouselItem>
      <CarouselItem>
        <div class="demo-carousel">
          <img class="demo-img" height="250px;" src="http://img1.qunarzz.com/piao/fusion/1806/de/df09fa61aed2d502.jpg_750x200_67472739.jpg">
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
