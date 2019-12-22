<template>
  <div>
    <Card style="width:350px">
      <p slot="title">
        <Icon type="ios-card" />
        VIP会员
      </p>
      <a href="#" slot="extra" @click.prevent="changeLimit">
        <Icon type="md-cart" />
        下单
      </a>
      <div>这是是简单介绍</div>
    </Card>
  </div>

</template>
<script>
  export default {
    name: "showGoods",
    data () {
      return {

      }
    },
    methods: {

    }
  }
</script>

<style scoped>

</style>



