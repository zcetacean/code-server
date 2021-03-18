import Vue from 'vue'
import App from './App'
import req from './REQ.js'
import share from './SHARE.js'
import uView from "uview-ui"

Vue.use(uView)
Vue.mixin(share)
Vue.prototype.$req = req
Vue.config.productionTip = false

App.mpType = 'app'

const app = new Vue({
    ...App
})
app.$mount()
