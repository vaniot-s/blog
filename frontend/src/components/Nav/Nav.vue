<template>
  <a-layout-header :style="{ position: 'fixed', zIndex: 1, width: '100%', background: '#fff', padding: 0 }">
    <a-row >
      <a-col :xs="4" :sm="4" :md="4" :lg="4" :xl="4"></a-col>
      <a-col :xs="16" :sm="16" :md="16" :lg="16" :xl="16">
        <a-row >
          <div class="logo" />
          <a-menu
            theme="light"
            mode="horizontal"
            :defaultSelectedKeys="['1']"
            :style="{ lineHeight: '64px' }"
          >
            <a-menu-item key="1">首页</a-menu-item>
            <a-menu-item key="2">分类</a-menu-item>
            <a-menu-item key="3">归档</a-menu-item>
          </a-menu>
        </a-row>
      </a-col>
      <a-col :xs="4" :sm="4" :md="4" :lg="4" :xl="4"></a-col>
    </a-row>

  </a-layout-header>
</template>

<script>
    export default {
        name: "Nav"
    }
</script>

<style scoped>
  .logo {
    width: 120px;
    height: 31px;
    background: rgba(255, 255, 255, 0.2);
    margin: 16px 24px 16px 0  ;
    float: left;
  }
</style>
