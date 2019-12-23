<template>
  <div>
    <div>
      <br><div>订单管理 => 下单付款</div><br>
    </div>
    <Form :label-width="300">
      <FormItem label="进入仓库">
        <Icon type="ios-arrow-round-forward" style="font-size: 25px"/>
        <Tooltip content="点击进入" placement="top-end">
          <a><Icon type="md-home" style="font-size: 30px"  @click="repositoryModal = true"/></a>
        </Tooltip>
      </FormItem>
      <FormItem label="商品ID">
        <Input v-model="ProductId" placeholder="商户ID" style="width: 300px" readonly></Input>
      </FormItem>
      <FormItem label="商品描述">
        <Input v-model="ProductDesc" placeholder="商品描述" number style="width: 300px" readonly></Input>
      </FormItem>
      <FormItem label="金额">
        <Input v-model="TransAmount" placeholder="金额" style="width: 300px" readonly></Input>
      </FormItem>
      <FormItem label="币种">
        <Input v-model="TransCurrCode" placeholder="币种" style="width: 300px" readonly></Input>
      </FormItem>
      <FormItem>
        <Button type="primary">下单</Button>
        <Button style="margin-left: 8px">重置</Button>
      </FormItem>
    </Form>
    <Modal title="中央仓库" fullscreen  v-model="repositoryModal">
      <showGoods></showGoods>
    </Modal>
  </div>
</template>
<script>
  import showGoods from '../centralRepository/showGoods'
  export default {
    name:"payOrder",
    components:{showGoods,},
    data (){
      return{
        ProductId:'',
        ProductDesc:'',
        TransAmount:'',
        TransCurrCode:'',
        repositoryModal:false,
        centralRepository:{}
      }
    }
  }
</script>

<style scoped>

</style>
