// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'

//Librerias para que funcione axios
//Axios sirve para hacer las peticiones al web service
import axios from 'axios'
import VueAxios from 'vue-axios'
 
Vue.use(VueAxios, axios)



Vue.config.productionTip = false

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  template: '<App/>',
  components: { App }
})
