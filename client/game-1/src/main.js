import Vue from 'vue'
import axios from 'axios'
import VueAxios from 'vue-axios'
import Utils from './utils'
Vue.use(VueAxios,axios);
Vue.use(Utils)

import App from './App.vue'
import router from './router'
import store from './store'
Vue.config.productionTip = false

// @text Vue组件
import { Button, Form, Input, Icon, Checkbox, Row, Col, Alert } from 'ant-design-vue';
Vue.component(Button.name, Button)
Vue.component(Input.name, Input)
Vue.component(Icon.name, Icon)
Vue.component(Checkbox.name, Checkbox)
Vue.component(Row.name, Row)
Vue.component(Col.name, Col)
Vue.component(Alert.name, Alert)


// @text 实力化app对象
new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
