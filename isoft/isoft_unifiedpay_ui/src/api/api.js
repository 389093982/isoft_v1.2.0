// 所有界面的跨站请求在此统一管理，统一为POST异步请求
import {POST} from './POST'
const baseUrl = '/wechatPayApi';

export const QueryOrder = (params) => POST(baseUrl+'/QueryOrder',params);
export const ShowLastedOrders = (params) => POST(baseUrl+'/ShowLastedOrders',params);


