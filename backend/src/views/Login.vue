<template>
    <div class="bg">
        <div class="box">
            <a-card style="padding: 30px; border-radius: 10px">
                <a-form-model ref="form" :layout="layout" :model="form" v-bind="formItemLayout">
                    <div class="name">VKE-ADMIN 管理系统</div>
                    <a-form-model-item has-feedback label="账号" prop="username">
                        <a-input v-model="form.username" autocomplete="off" />
                    </a-form-model-item>
                    <a-form-model-item has-feedback label="密码" prop="password">
                        <a-input v-model="form.password" type="password" autocomplete="off" />
                    </a-form-model-item>
                    <div class="btn-wrapper">
                        <a-button type="primary" class="login" @click="submitForm('ruleForm')">
                            登录
                        </a-button>
                    </div>
                    <div class="forget"><a href="#">忘记密码?</a></div>
                </a-form-model>
            </a-card>
        </div>
    </div>
</template>

<script>

export default {
    data() {
        return {
            form: {
                username: "",
                password: ""
            },
            layout: "horizontal"
        }
    },
    computed: {
        formItemLayout() {
            return this.layout === 'horizontal'
                ? {
                    labelCol: { span: 4 },
                    wrapperCol: { span: 17 },
                }
                : {};
        },
    }
}
</script>

<style scoped>
.bg {
    background: url("~@/assets/img/login-bg.jpg");
    background-repeat:no-repeat;
    background-position:top;

    height: 100vh;
    display: flex;
    align-items: center;
    justify-content: center;
}
.btn-wrapper {
    text-align: center;
    margin-bottom: 30px;
}
.login {
    width: 70%;
}
.name {
    text-align: center;
    margin-bottom: 30px;
}
.forget {
    text-align: center;
}
</style>