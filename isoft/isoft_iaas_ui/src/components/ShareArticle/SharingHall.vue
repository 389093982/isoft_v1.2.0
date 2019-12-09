<template>
    <div>
      <div>分享大厅</div>
      <div>
        <SharingUser :article="article"></SharingUser><br>
      </div>
    </div>
</template>

<script>
    import SharingUser from "../ShareArticle/SharingUser"
    export default {
      props:{
        article:{
          title:String,
          content:String
        },
      },
      name: "SharingHall",
      components:{SharingUser},
    }
</script>

<style scoped>

</style>
