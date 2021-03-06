package models

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
	"strings"
)

//交易订单结构体
type Order struct {
	Id                   int64
	OrderId              string `orm:"unique"` //支付订单号
	OrgOrderId           string //原交易订单号
	PayStyle             string //支付方式(微信支付)
	TransType            string //交易类型（SALE/REFUND）
	MerchantNo           string //商户号
	ProductId            string //商品ID
	ProductDesc          string //商品描述
	TransTime            string //交易时间
	TransAmount          int64  //交易金额
	TransCurrCode        string //交易币种
	OrderResultCode      string //下单业务结果
	OrderResultDesc      string //下单结果描述
	OrderErrCode         string //下单错误代码
	OrderErrCodeDes      string //下单错误代码描述
	CodeUrl              string //付款二维码
	PayResultCode        string //付款业务结果代码
	PayResultDesc        string //付款业务结果描述
	PayErrCode           string //付款错误代码
	PayErrCodeDesc       string //付款错误代码描述
	TransactionId        string //微信官方支付订单号
	WechatCash           string //微信零钱支付
	BankType             string //付款银行code
	BankName             string //付款银行名称
	PayTimeEnd           string //支付完成时间
	RefundReason         string //退货原因
	RefundReqResultCode  string //退货请求业务结果代码
	RefundReqResultDesc  string //退货请求业务结果描述
	RefundReqErrCode     string //退货请求错误代码
	RefundReqErrCodeDesc string //退货请求错误代码描述
	RefundStatus         string //退款状态
	RefundSuccessTime    string //退款成功时间
	RefundedAmount       int64  //已退金额
}

//下单
func (this *Order) Pay(o orm.Ormer, order Order) error {
	//1.支付参数验证
	validation := this.PayParamValidation(order)
	if validation != nil {
		return validation
	}
	//2.入库
	_, e := o.Insert(&order)
	if e != nil {
		return errors.New(fmt.Sprintf(fmt.Sprintf("支付订单%v入库失败,失败原因:%v", order.OrderId, e.Error())))
	}
	return nil
}

//退货
func (this *Order) Refund(o orm.Ormer, order Order) error {
	//1.退货参数验证
	validation := this.RefundParamValidation(order)
	if validation != nil {
		return validation
	}
	//2.查询原交易订单
	originalOrder := Order{}
	originalOrder.OrderId = order.OrgOrderId
	originalOrder.TransType = "SALE"
	err := o.Read(&originalOrder, "OrderId", "TransType")
	if err != nil {
		return errors.New(fmt.Sprintf("原交易订单%v查询失败！", originalOrder.OrderId))
	}
	//3.查看原交易订单是否支付成功
	if originalOrder.PayResultCode != "SUCCESS" {
		return errors.New(fmt.Sprintf("原交易订单付款未成功,不允许退货！"))
	}
	//4.判断退货金额是否大于可退金额
	returnableAmount := originalOrder.TransAmount - originalOrder.RefundedAmount
	refundAmount := order.TransAmount
	if refundAmount > returnableAmount {
		return errors.New("退货金额大于可退金额！")
	}
	//5.入库
	_, e := o.Insert(&order)
	if e != nil {
		return errors.New(fmt.Sprintf("退货订单%v入库失败,失败原因:%v", order.OrderId, e.Error()))
	}
	//6.更新原交易已退金额
	originalOrder.RefundedAmount = originalOrder.RefundedAmount + refundAmount
	_, e = o.Update(&originalOrder, "RefundedAmount")
	if e != nil {
		return errors.New(fmt.Sprintf("退货订单%v更新原交易订单%v已退金额失败,失败原因:%v", order.OrderId, order.OrgOrderId, e.Error()))
	}
	return nil
}

//支付订单参数验证
func (this *Order) PayParamValidation(order Order) error {
	if len(strings.TrimSpace(order.OrderId)) == 0 {
		return errors.New("交易订单为空！")
	}
	if len(strings.TrimSpace(order.PayStyle)) == 0 {
		return errors.New("支付方式不能为空！")
	}
	if order.TransType != "SALE" {
		return errors.New("交易类型不正确！")
	}
	if len(strings.TrimSpace(order.MerchantNo)) == 0 {
		return errors.New("商户号为空！")
	}
	if len(strings.TrimSpace(order.ProductId)) == 0 {
		return errors.New("下单的产品代码为空！")
	}
	if len(strings.TrimSpace(order.ProductDesc)) == 0 {
		return errors.New("下单的产品描述为空！")
	}
	if len(strings.TrimSpace(order.TransTime)) == 0 {
		return errors.New("交易时间为空！")
	}
	if order.TransAmount <= 0 {
		return errors.New("交易金额不正确！")
	}
	if order.TransCurrCode != "CNY" {
		return errors.New("交易币种不正确，目前只支持CNY！")
	}
	return nil
}

//退货参数验证
func (this *Order) RefundParamValidation(order Order) error {
	if len(strings.TrimSpace(order.OrderId)) == 0 {
		return errors.New("退货订单为空！")
	}
	if len(strings.TrimSpace(order.OrgOrderId)) == 0 {
		return errors.New("原交易订单为空！")
	}
	if order.TransType != "REFUND" {
		return errors.New("交易类型不正确！")
	}
	if len(strings.TrimSpace(order.MerchantNo)) == 0 {
		return errors.New("商户号为空！")
	}
	if len(strings.TrimSpace(order.ProductId)) == 0 {
		return errors.New("退货的产品代码为空！")
	}
	if len(strings.TrimSpace(order.ProductDesc)) == 0 {
		return errors.New("退货的产品描述为空！")
	}
	if len(strings.TrimSpace(order.TransTime)) == 0 {
		return errors.New("交易时间为空！")
	}
	if order.TransAmount <= 0 {
		return errors.New("交易金额不正确！")
	}
	if order.TransCurrCode != "CNY" {
		return errors.New("交易币种不正确，目前只支持CNY！")
	}
	if len(strings.TrimSpace(order.RefundReason)) == 0 {
		return errors.New("退货原因为空！")
	}
	return nil
}
