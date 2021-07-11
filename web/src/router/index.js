import {createRouter, createWebHistory} from "vue-router";
import Home from "../views/Home.vue";


const routes = [
    {
        path: '/',
        name: 'Home',
        component: Home,
    },
    {
        path: "/404",
        name: "NotFound",
        component: () => import("../views/NotFound.vue")
    },
    {
        path: "/login",
        name: "Login",
        component: () => import("../views/Login.vue")
    },
    {
        path: '/admin',
        name: 'Admin',
        component: () => import("../views/Admin.vue"),
        children: [
            {
                path: 'urls',
                name: "Url",
                component: () => import("../views/Url.vue")
            }, {
                path: 'urls/:slug/analytics',
                name: "UrlAnalytics",
                props: true,
                component: () => import("../views/url/Analytics.vue")
            },
        ]
    },
];

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes
})


export default router;