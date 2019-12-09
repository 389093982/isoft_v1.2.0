<template>
    <div>
      <div @click=showTitle(AdvertisementName)>点击我</div>
      <div v-if="isShow">{{AdvertisementName01}}</div>
      <div v-for="item in Advertisements" v-if="isShow">{{item}}</div>
    </div>
</template>

<script>
    export default {
        name: "Advertisements",
      data(){
          return{
            AdvertisementName:"广告名称0"
          }
      },
      props:{
          Advertisements:{
            type:Array,
            default:["1111","2222","3333"]
          },
          isShow:{
            type:Boolean,
            default:true,
          },
          AdvertisementName01:{
            type: String,
            default: "广告名称01",
        }
      },
    methods:{
          showTitle: function (title) {
             alert(title)
          }
    }


    }
</script>

<style scoped>

</style>
