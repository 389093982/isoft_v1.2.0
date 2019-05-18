// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'
import store from './store'

// 引用全局静态数据
import global_ from './components/GlobalData'     //引用文件
Vue.prototype.GLOBAL = global_                    //挂载到Vue实例上面,通过 this.GLOBAL.xxx 访问全局变量

// 使用 iview
import iView from 'iview'
import 'iview/dist/styles/iview.css'
Vue.use(iView);

Vue.config.productionTip = false

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  components: { App },
  template: '<App/>',
  store, // 使用上vuex
});
