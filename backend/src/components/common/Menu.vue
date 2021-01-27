<template>
  <el-aside width="200px">
    <el-menu background-color="#1f4f9e" text-color="#fff" active-text-color="#1f4f9e" :router="true" :default-active="default_active">
      <el-submenu :index="item.url" v-for="(item, index) in Menu" :key="index">
        <template slot="title"><i class="el-icon-message"></i>{{item.menu}}</template>
          <el-menu-item v-for="(children, i) in item.children" :key="i" :index="children.url" :to="children.url">{{children.menu}}</el-menu-item>
      </el-submenu>
    </el-menu>
  </el-aside>    
</template>

<script>
import Menu from "@/data/menu";

export default {
  props: ["colloapse"],
  data() {
    return {
        Menu,
        default_active: "/"
    }
  },
  mounted() {
    this.default_active = this.$route.fullPath
  },
}
</script>

<style lang="scss" scope>
  @import "@/assets/css/variables.sass"; 
  .el-aside{
    background-color: $blue-color;
  }
  .el-menu{
    border-right: none!important;
    & .el-menu-item{
      color: $white-color!important;
    }
    & i{
      color: $white-color!important;
    }
    .el-submenu:hover{
      background: $s-blue-color!important;
    }
  }
  .el-menu-item.is-active {
    background-color: #193f7e!important;
  }
</style>