// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'
import store from './store'

import {CheckSSOLogin} from "./tools"

// 引用全局静态数据
import global_ from './components/GlobalData'     //引用文件
Vue.prototype.GLOBAL = global_                    //挂载到Vue实例上面,通过 this.GLOBAL.xxx 访问全局变量

// 使用 iview
import iView from 'iview'
import 'iview/dist/styles/iview.css'
Vue.use(iView);

// 使用 vue-markdown
import mavonEditor from 'mavon-editor'
import 'mavon-editor/dist/css/index.css'
Vue.use(mavonEditor)

// 使用全局的 animate.css
import animated from 'animate.css'
Vue.use(animated)

// 使用全局的 hover.css
import hovered from 'hover.css'
Vue.use(hovered)

//------------------------------------------------------------------
// 注册自定义公共组件
import IBeautifulLink from "./components/Common/link/IBeautifulLink"
Vue.component('IBeautifulLink', IBeautifulLink);
import "./assets/css/isoft_common.css"
//------------------------------------------------------------------

Vue.config.productionTip = false


router.beforeEach((to, from, next) => {
  /* 路由发生变化修改页面title */
  if (to.meta.title) {
    document.title = to.meta.title;
  }else{
    document.title = "地耳 App";
  }
  // LoadingBar 加载进度条
  iView.LoadingBar.start();

  CheckSSOLogin(to, from, next);
});

router.afterEach(route => {
  // LoadingBar 加载进度条
  iView.LoadingBar.finish();
});

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  components: { App },
  template: '<App/>',
  store, // 使用上vuex
});
