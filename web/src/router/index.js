import {createRouter, createWebHistory} from "vue-router";
import Home from "../views/Home.vue";
import Cookies from 'js-cookie'

const checkLogin = (to, from, next) => {
    if (typeof Cookies.get("uid") === 'undefined') {
        window.location.href = '/login'
        return
    }

    next()
}

const routes = [
    {
        path: '/',
        name: 'Home',
        beforeEnter: checkLogin,
        component: Home,
    },
    {
        path: '/admin',
        name: 'Admin',
        component: () => import("../views/Admin.vue"),
        children: [
            {
                path: 'links',
                name: "Link",
                component: () => import("../views/Link.vue")
            },
        ]
    },
];

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes
})


export default router;