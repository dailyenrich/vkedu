<template>
    <el-container class="wrapper">
        <el-container style="flex-direction: column;">
            <el-main>
                <div class="content">
                    <el-card shadow="always" class="login-box">
                        <div slot="header" class="clearfix">
                            <span class="card-title">欢迎登陆</span>
                        </div>
                        <el-form ref="form" :model="form">
                            <el-form-item>
                                <el-input size="mini" prefix-icon="el-icon-user-solid" v-model="form.username"></el-input>
                            </el-form-item>
                            <el-form-item>
                                <el-input type="password" size="mini" prefix-icon="el-icon-view" v-model="form.password"></el-input>
                            </el-form-item>
                            <el-button size="mini" type="primary" @click="login('form')">立即登陆</el-button>
                        </el-form>
                    </el-card>
                </div>
            </el-main>
            <Footer/>
        </el-container>
    </el-container>
</template>

<script>
import AuthApi from "@/api/auth";
import UserModel from "@/model/user";
import routes_config from "@/config/routes";

export default {
    data() {
        return {
            form: {
                username: "",
                password: ""
            }
        }
    },
    methods: {
        login() {
            if(!this.form.username || !this.form.password) {
                this.$message.error('账号密码不能为空');
                return false;
            }
            UserModel.setUsername(this.form.username)
            UserModel.setPassword(this.form.password)

            // this.$message.success("登陆成功~~ ");
            // this.$router.push({name: routes_config.DEFAULT_ROUTE_NAME})
            let _this = this;
            AuthApi.login(UserModel, function(r) {
                if(!r) {
                    _this.$message.error("登陆失败~~ ");
                    return
                }

                _this.$message.success("登陆成功~~ ");
                _this.$router.push({name: routes_config.DEFAULT_ROUTE_NAME})
            })
        }
    }
}
</script>

<style scoped>

    .wrapper {
        min-width: 1024px;
        height: 100%;
        background: url("~@/assets/img/login-bg.jpg")
    }
    .login-box {
        width: 300px;
    }
    .el-form-item{
        display: flex;
        justify-content: center;
    }
    .el-form{
        text-align: center;
    }
    .el-input{
        width: 250px;
    }
    .el-button{
        width: 250px;
        margin: 0 0 15px;
    }
    .content {
        display: flex;
        flex-direction: row;
        justify-content: flex-end;
        align-items: center;
        min-height: 100%;
        margin-right: 350px;
    }
    .clearfix {
        text-align: center;
    }
</style>5