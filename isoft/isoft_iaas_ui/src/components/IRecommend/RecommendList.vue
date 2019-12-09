<template>
    <div>
      <p>易推广</p>
      <HotRecommendList></HotRecommendList>
      <Advertisements :advertisement-name01=AdvertisementName1 :is-show="true" :advertisements="Advertisement"></Advertisements>
      <Advertisements :advertisement-name01=AdvertisementName2 :is-show="false"></Advertisements>
      <Advertisements :advertisement-name01=AdvertisementName3 :is-show="false"></Advertisements>
      <div>
        <button>我要发布</button>
      </div>
      <div>
        <button>我要接单</button>
      </div>
    </div>
</template>

<script>
    import HotRecommendList from "./HotRecommendList";
    import Advertisements from "./Advertisements";
    export default {
        name: "RecommendList",
      components: {Advertisements, HotRecommendList},
      data(){
          return {
            AdvertisementName1:"测试广告1",
            AdvertisementName2:"测试广告2",
            AdvertisementName3:"测试广告3",
            Advertisement:["7777","8888","9999"]
          }
      }
    }
</script>

<style scoped>

</style>
