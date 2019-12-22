// 所有POST请求在此发送
import axios from 'axios'
import Qs from 'qs'

export function POST(url, params) {
  //返回异步请求结果
  return new Promise(function(resolve,reject){
    //统一去首尾空格
    let strings = Object.keys(params);
    let trimParams = {};
    for (let i in strings) {
      trimParams[strings[i]] = params[strings[i]].trim()
    }
    //发送请求并返回数据
    axios.post(url, Qs.stringify(trimParams))
      .then(res => {
        // console.log(res.data);
        resolve(res.data);
      })
      .catch(err => {
        console.log(err);
        reject(err);
      });
  });


}


