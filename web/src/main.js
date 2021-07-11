import {createApp} from 'vue'
import App from './App.vue'
import router from "./router";
import './index.less'

// import {
//     Button,
//     Card,
//     Col,
//     Form,
//     Input,
//     Layout, List,
//     Menu,
//     PageHeader,
//     Pagination,
//     Row,
//     Switch,
//     Table, Tabs,
//     Tag
// } from "ant-design-vue";

const app = createApp(App);

app.use(router);
// app.use(Button)
//     .use(Menu)
//     .use(Layout)
//     .use(Table)
//     .use(Pagination)
//     .use(Tag)
//     .use(Switch)
//     .use(Form)
//     .use(Input)
//     .use(Row)
//     .use(Col)
//     .use(PageHeader)
//     .use(Card)
//     .use(List)
//     .use(Tabs)

app.mount('#app')