import Vue from "vue"
import VueRouter from "vue-router"
import routes from "@/routes"
import routes_config from "./routes"
import AuthApi from "@/api/auth"

Vue.use(VueRouter)

let white_list = [routes_config.LOGIN_ROUTE_NAME]

let router = new VueRouter({
    mode: 'history',
    routes,
});

router.beforeEach((to, from, next) => {
    if (white_list.indexOf(to.name) != -1 && AuthApi.check()) {
       next({name: routes_config.DEFAULT_ROUTE_NAME}) 
    }

    if (white_list.indexOf(to.name) == -1 && !AuthApi.check()) {
       next({name: routes_config.LOGIN_ROUTE_NAME}) 
    }
    next()
})

export default router;