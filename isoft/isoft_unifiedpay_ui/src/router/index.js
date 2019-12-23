//==============================================开始引入组件============================================================
import Vue from 'vue'
import Router from 'vue-router'
const baseLayout = () => import('../components/base/baseLayout');
const orderList  = () => import('../components/wechatpay/orderList');
const payOrder  = () => import('../components/wechatpay/payOrder');
const orderRefund  = () => import('../components/wechatpay/orderRefund');
const test  = () => import('../components/wechatpay/test');

Vue.use(Router);

//==============================================引入二级模块============================================================
//微信支付
const wechatpayRouters = [
  {path:'orderList',component:orderList},
  {path:'payOrder',component:payOrder},
  {path:'orderRefund',component:orderRefund},
  {path:'test',component:test},
];


//==============================================引入一级模块============================================================
export default new Router({
  //去掉#号
  mode:'history',
  routes: [
    {path:'/wechatpay',component:baseLayout,children:wechatpayRouters},//微信支付
    { path: '*', redirect: '/wechatpay/payOrder' }
  ]
})
