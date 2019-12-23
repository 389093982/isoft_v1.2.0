<template>
  <div>
    <br><div>订单管理 => 订单查询</div><br>
    <div>
      <Input v-model="OrderId" placeholder="订单号" style="width: 250px"/>
      <Select v-model="TransType" placeholder="交易类型" style="width:100px">
        <Option v-for="item in TransTypeList" :value="item.value" :key="item.value">{{ item.label }}</Option>
      </Select>
      <Input v-model="ProductDesc" placeholder="商品描述" style="width: 150px"/>
      <DatePicker v-model="TransTime" type="date" placeholder="交易日期" style="width: 140px"></DatePicker>
      <Input v-model="TransAmount" placeholder="交易金额" style="width: 100px"/>
      <Button type="primary" shape="circle" icon="ios-search" @click="queryOrder"></Button>
    </div>
    <div>
      <Table width="1020" height="430" border :columns="columns" :data="orders"></Table>
    </div>
    <Modal title="订单详情" v-model="orderDetailModal" :mask-closable="false" width="1200">
      <orderDetail :Order=orderDetailModalData></orderDetail>
    </Modal>
  </div>
</template>
<script>
  import {QueryOrder,ShowLastedOrders} from '../../api/api'
  import orderDetail from './orderDetail'
  export default {
    name: "orderList",
    components:{orderDetail,},
    data () {
      return {
        OrderId:'',
        TransType:'',
        TransTypeList:[{value:'SALE',label:'SALE'},{value:'REFUND',label:'REFUND'}],
        ProductDesc:'',
        TransTime:'',
        TransAmount:'',
        orderDetailModal: false,
        columns: [
          {title: '订单号',                   key: 'OrderId',                 width: 250,           fixed: 'left'},
          {title: '交易类型',                 key: 'TransType',               width: 100},
          {title: '商品描述',                 key: 'ProductDesc',             width: 150},
          {title: '交易时间',                 key: 'TransTime',               width: 150},
          {title: '交易金额',                 key: 'TransAmount',             width: 100},
          {title: '下单结果描述',             key: 'OrderResultDesc',        width: 140},
          {title: '付款业务结果描述',         key: 'PayResultDesc',          width: 150},
          {title: '退款请求业务结果描述',    key: 'RefundReqResultDesc',    width: 160},
          {title: '退款状态',                 key: 'RefundStatus',            width: 100},
          {title: '已退金额',                 key: 'RefundedAmount',          width: 100},
          {title: '操作',                     key: 'action',                   width: 120,         fixed: 'right',
            render: (h, params) => {
              return h('div', [
                h('Button', {
                  props: {type: 'text', size: 'small'},
                  on:{
                    click: () => {
                      this.orderDetailModal = true;
                      for (let i = 0; i < this.orders.length; i++) {
                        if (this.orders[i].OrderId===this.orders[params.index].OrderId){
                          this.orderDetailModalData = this.orders[i];
                        }
                      }
                    }
                  }
                }, '查看'),
              ]);
            }
          }
        ],
        orders: [],
        orderDetailModalData:{}
      }
    },
    // created:{
    //
    // },
    mounted (){
      //初始化界面，展示几条订单
      this.showLastedOrders()
    },
    methods:{
      queryOrder:async function () {
        if (this.OrderId.length===0 && this.TransType.length===0 && this.ProductDesc.length===0 && this.TransTime.length===0 && this.TransAmount.length===0){
          this.$Message.error('查询条件不能全部为空！');
          return false;
        }
        let params = {
          'OrderId':this.OrderId,
          'TransType':this.TransType,
          'ProductDesc':this.ProductDesc,
          'TransTime':new Date(this.TransTime).format("yyyyMMdd"),
          'TransAmount':this.TransAmount,
        };
        let orders = await QueryOrder(params);
        let ordersObj = JSON.parse(orders);
        if (ordersObj.length === 0) {
          this.$Message.warning("查不到数据！")
        }else{
          this.$Message.success("查询成功！");
          for (let i = 0; i < ordersObj.length; i++) {
            ordersObj[i].TransAmount = ordersObj[i].TransAmount/100;  //金额从分转为元
            ordersObj[i].RefundedAmount = ordersObj[i].RefundedAmount/100  //金额从分转为元
          }
          this.orders = ordersObj;
        }
      },

      showLastedOrders:async function() {
        let params = {
          'count':'8',
        };
        let orders = await ShowLastedOrders(params);
        let ordersObj = JSON.parse(orders);
        if (ordersObj.length === 0) {
          this.$Message.warning("查不到数据！")
        }else{
          for (let i = 0; i < ordersObj.length; i++) {
            ordersObj[i].TransAmount = ordersObj[i].TransAmount/100;  //金额从分转为元
            ordersObj[i].RefundedAmount = ordersObj[i].RefundedAmount/100  //金额从分转为元
          }
          this.orders = ordersObj;
        }
      }


    }
  }
</script>

<style scoped>

</style>
