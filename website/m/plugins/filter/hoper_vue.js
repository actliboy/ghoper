import Vue from 'vue'
import {Toast} from 'vant'
//import 'vant/lib/index.css';
//import Mint from 'mint-ui';
//import 'mint-ui/lib/style.css';
//import iView from 'iview';
//import 'iview/dist/styles/iview.css';

Vue.filter('dateFormat', function (value) {
    let date = new Date(value);
    return date.getFullYear() + '-' + (date.getMonth() + 1) + '-' + date.getDate() + ' ' + date.getHours() + ':' + date.getMinutes() + ':' + date.getSeconds();
});
import lioHint from '../../pages/common/hint'
Vue.component('lio-hint',lioHint);

//Vue.use(iView);
//Vue.use(Mint);
Vue.use(Toast);
