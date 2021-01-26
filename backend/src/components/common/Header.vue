<template>
    <el-header style="">
        <el-row>
            <el-col :span="17">
                <el-breadcrumb separator="/">
                    <el-breadcrumb-item v-for="(item, index) in breadcrumb" :key="index">
                        {{item.name}}
                    </el-breadcrumb-item>
                </el-breadcrumb>
            </el-col>
            <el-col :span="7" class="userinfo">
                <span>{{user.username}}</span>
                <span @click="logout()">退出</span>
            </el-col>
        </el-row>
    </el-header>
</template>

<script>
import AuthApi from "@/api/auth";
import routes_config from "@/config/routes";

export default {
    data() {
        return {
            user: {username: "demo"},
            breadcrumb: this.$route.matched
        }
    },
    mounted() {
        this.user = AuthApi.user()
    },
    watch: {
        $route: function() {
            this.breadcrumb = this.$route.matched
        }
    },
    methods: {
        // initBreadcrumb() {
        //     this.breadcrumb = this.$route.matched
        //     window.console.log(this.breadcrumb)
        // },
        logout() {
            if(AuthApi.logout()) {
                this.$router.push({name: routes_config.LOGIN_ROUTE_NAME});
            }
        }
    }
}
</script>

<style scoped lang="scss">
  @import "@/assets/css/variables.sass"; 
 .el-header {
    text-align: right;
    height: 50px!important;
    font-size: 12px!important;
    line-height: 50px;
    border-bottom: 1px solid rgb(238, 241, 246);
 }
 .el-breadcrumb{
    font-size: 12px!important;
    height: 100%!important;
    line-height: inherit!important;
 }
 .menu-icon{
    padding-top: 5px;
    & i{
        font-size: 25px;margin-right: 10px;
    }
 }
 .userinfo {
    color: $blue-color;
    & span:last-child {
        margin-left: 10px;
        cursor: pointer;
    }
 }
</style>