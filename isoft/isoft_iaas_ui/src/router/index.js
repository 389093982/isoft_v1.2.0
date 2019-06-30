import Vue from 'vue'
import Router from 'vue-router'

import {ISSOReouter} from "./isso"
import {ILearningRouters} from "./ilearning"

Vue.use(Router);

function getRootRouters () {
  return [{
    path: '/',
    redirect: '/ilearning/index'
  }]
};


function getAllRouters() {
  let allRouters = [];
  [].push.apply(allRouters, ILearningRouters);
  [].push.apply(allRouters, ISSOReouter);
  [].push.apply(allRouters, getRootRouters());
  return allRouters;
}


export default new Router({
  // # 主要用来区分前后台应用, history 模式需要使用 nginx 代理
  // History 模式,去除vue项目中的 #
  mode: 'history',
  routes: getAllRouters(),
})
