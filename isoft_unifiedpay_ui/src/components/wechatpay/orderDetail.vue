<!--订单详情-->
<template>
  <div>
    <Row>
      <Col span="6"><div>订单号:<a>{{Order.OrderId}}</a></div></Col>
      <Col span="6"><div>原交易订单:<a>{{Order.OrgOrderId}}</a></div></Col>
      <Col span="6"><div>支付方式:<a>{{Order.PayStyle}}</a></div></Col>
      <Col span="6"><div>交易类型:<a>{{Order.TransType}}</a></div></Col>
    </Row>
    <Row>
      <Col span="6"><div>商户号:<a>{{Order.MerchantNo}}</a></div></Col>
      <Col span="6"><div>商品ID:<a>{{Order.ProductId}}</a></div></Col>
      <Col span="6"><div>商品描述:<a>{{Order.ProductDesc}}</a></div></Col>
      <Col span="6"><div>交易时间:<a>{{Order.TransTime}}</a></div></Col>
    </Row>
    <Row>
      <Col span="6"><div>交易金额:<a>{{Order.TransAmount}}</a></div></Col>
      <Col span="6"><div>交易币种:<a>{{Order.TransCurrCode}}</a></div></Col>
      <Col span="6"><div>下单业务结果:<a>{{Order.OrderResultCode}}</a></div></Col>
      <Col span="6"><div>下单结果描述:<a>{{Order.OrderResultDesc}}</a></div></Col>
    </Row>
    <Row>
      <Col span="6"><div>下单错误代码:<a>{{Order.OrderErrCode}}</a></div></Col>
      <Col span="6"><div>下单错误代码描述:<a>{{Order.OrderErrCodeDes}}</a></div></Col>
      <Col span="6"><div>付款二维码:<a>{{Order.CodeUrl}}</a></div></Col>
      <Col span="6"><div>付款业务结果代码:<a>{{Order.PayResultCode}}</a></div></Col>
    </Row>
    <Row>
      <Col span="6"><div>付款业务结果描述:<a>{{Order.PayResultDesc}}</a></div></Col>
      <Col span="6"><div>付款错误代码:<a>{{Order.PayErrCode}}</a></div></Col>
      <Col span="6"><div>付款错误代码描述:<a>{{Order.PayErrCodeDesc}}</a></div></Col>
      <Col span="6"><div>微信官方支付订单号:<a>{{Order.TransactionId}}</a></div></Col>
    </Row>
    <Row>
      <Col span="6"><div>微信零钱支付:<a>{{Order.WechatCash}}</a></div></Col>
      <Col span="6"><div>付款银行code:<a>{{Order.BankType}}</a></div></Col>
      <Col span="6"><div>付款银行名称:<a>{{Order.BankName}}</a></div></Col>
      <Col span="6"><div>支付完成时间:<a>{{Order.PayTimeEnd}}</a></div></Col>
    </Row>
    <Row>
      <Col span="6"><div>退货原因:<a>{{Order.RefundReason}}</a></div></Col>
      <Col span="6"><div>退货请求业务结果代码:<a>{{Order.RefundReqResultCode}}</a></div></Col>
      <Col span="6"><div>退货请求业务结果描述:<a>{{Order.RefundReqResultDesc}}</a></div></Col>
      <Col span="6"><div>退货请求错误代码:<a>{{Order.RefundReqErrCode}}</a></div></Col>
    </Row>
    <Row>
      <Col span="6"><div>退货请求错误代码描述:<a>{{Order.RefundReqErrCodeDesc}}</a></div></Col>
      <Col span="6"><div>退款状态:<a>{{Order.RefundStatus}}</a></div></Col>
      <Col span="6"><div>退款成功时间:<a>{{Order.RefundSuccessTime}}</a></div></Col>
      <Col span="6"><div>已退金额:<a>{{Order.RefundedAmount}}</a></div></Col>
    </Row>

  </div>
</template>

<script>
	export default {
		name: "orderDetail",
    props:{
		  Order:{

      }
    }
	}
</script>

<style scoped>

</style>
