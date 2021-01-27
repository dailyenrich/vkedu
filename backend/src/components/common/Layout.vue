<template>
    <el-container>
        <Menu :collapse="false"/>
        <el-container style="flex-direction: column;">
            <Header :breadcrumb="this.breadcrumb"/>
            <el-main>
                <div class="content">
                  <router-view></router-view>
                </div>
            </el-main>
            <Footer/>
        </el-container>
    </el-container>
</template>

<script>
import Header from "./Header";
import Menu from "./Menu";
import Footer from "./Footer";

export default {
    props: {
        "breadcrumb": Array
    },
    components: {
        Header,
        Menu,
        Footer
    },
    updated() {
    }
}
</script>

<style scoped lang="scss">
    @import "@/assets/css/variables.sass";
    .el-breadcrumb {
        padding: 10px 0px;
    }
    .el-main{
        background: #f7f7f7;
    }
    .el-container{
        height: 100%;
        min-width: 1024px;
    }
    .content{
        background: $white-color;
        padding: 15px;
    }
</style>
