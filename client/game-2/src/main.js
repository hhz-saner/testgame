import Vue from 'vue'
import axios from 'axios'
import VueAxios from 'vue-axios'
import Utils from "./utils";

Vue.use(VueAxios, axios);
Vue.use(Utils);

import "./assets/app.css";

import App from './App.vue'
import router from './router'
import store from './store'

Vue.config.productionTip = false

import {
    Row,
    Col,
    Input,
    Button,
    Checkbox,
    Icon,
    InputSearch
} from 'ant-design-vue';

Vue.component(Row.name, Row)
Vue.component(Col.name, Col)
Vue.component(Input.name, Input)
Vue.component(Button.name, Button)
Vue.component(Icon.name, Icon)
Vue.component(Checkbox.name, Checkbox)
Vue.component(Input.Search.name, Input.Search)



new Vue({
    router,
    store,
    render: h => h(App)
}).$mount('#app')