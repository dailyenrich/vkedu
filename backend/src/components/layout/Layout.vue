<template>
  <a-layout id="components-layout-demo-custom-trigger">
    <a-layout-sider
      :trigger="null"
      collapsible
      v-model="collapsed"
    >
      <div class="logo"><img :src="logo" alt=""></div>
      <a-menu theme="dark" mode="inline" :defaultSelectedKeys="['5']" :default-open-keys="['sub1']">
        <a-sub-menu key="sub1">
          <span slot="title"><a-icon type="mail" /><span>Navigation One</span></span>
          <a-menu-item key="5">
            Option 5
          </a-menu-item>
          <a-menu-item key="6">
            Option 6
          </a-menu-item>
          <a-menu-item key="7">
            Option 7
          </a-menu-item>
          <a-menu-item key="8">
            Option 8
          </a-menu-item>
        </a-sub-menu>
        <a-sub-menu key="sub2">
          <span slot="title"><a-icon type="book" /><span>Navigation Two</span></span>
          <a-menu-item key="9">
            Option 9
          </a-menu-item>
          <a-menu-item key="10">
            Option 10
          </a-menu-item>
          <a-menu-item key="11">
            Option 11
          </a-menu-item>
          <a-menu-item key="12">
            Option 12
          </a-menu-item>
        </a-sub-menu>
      </a-menu>
    </a-layout-sider>
    <a-layout>
      <a-layout-header style="background: #fff; padding: 0">
        <a-menu
            theme="dark"
            mode="horizontal"
            :defaultSelectedKeys="['1-1']"
        >
            <a-icon
            class="trigger"
            :type="collapsed ? 'menu-unfold' : 'menu-fold'"
            @click="()=> collapsed = !collapsed"
            />
            <a-menu-item key="1-1">nav 1</a-menu-item>
            <a-menu-item key="2-1">nav 2</a-menu-item>
            <a-menu-item key="3-1">nav 3</a-menu-item>
        </a-menu>
      </a-layout-header>
      <a-layout-content :style="{ margin: '16px', padding: '16px', background: '#fff', minHeight: '280px' }">
          <router-view/>
      </a-layout-content>
    </a-layout>
  </a-layout>
</template>
<script>
import logo from "@/assets/logo.png"
export default {
  data(){
    return {
      collapsed: false,
      logo,
    }
  },
}
</script>
<style scoped>
.ant-layout-header {
    height: 46px!important;
}
#components-layout-demo-custom-trigger {
    height: 100vh;
}
#components-layout-demo-custom-trigger .trigger {
  font-size: 18px;
  padding: 0 24px;
  cursor: pointer;
  transition: color .3s;
}

#components-layout-demo-custom-trigger .trigger:hover {
  color: #1890ff;
}

#components-layout-demo-custom-trigger .logo {
  height: 32px;
  background: rgba(255,255,255,.2);
  margin: 5px;
  text-align: center;
}
#components-layout-demo-custom-trigger .logo img{
  max-height: 32px;
  max-width: 100%;
}
</style>
