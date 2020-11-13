<template>
    <div class="header">
        <div class="left">
            <ul>
                <li><a href="" class="logo"><Logo /></a></li>
                <li><a href="" class="active">首页</a></li>
                <li><a href="">课程</a></li>
                <li><a href="">博客</a></li>
            </ul>
        </div>
        <div class="right">
            <ul>
                <li><a href="">注册</a></li>
                <li><a href="">登录</a></li>
            </ul>
        </div>
    </div>
</template>

<script>
export default {
    
}
</script>

<style lang="scss" scoped>
a{
    color: $main-font-color;
    text-decoration: none;
    padding: 20px 15px;
    display: block;
    &.logo {
        padding: 0;
    }
    &.active {
        color: $main-font-color-active;
    }
}
.header {
    background: $white;
    position: fixed;
    z-index: 99;
    display: flex;
    justify-content: space-between;
    width: 100%;
    padding: 0 50px;
    box-sizing: border-box;
    box-shadow: -2px -4px 10px #333;

    .left > ul, .right > ul{
        display: flex;
        justify-content: space-between;
        align-items: center;
    }
}
</style>
