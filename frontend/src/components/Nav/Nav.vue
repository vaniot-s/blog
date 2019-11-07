<template>
  <a-layout-header class="header">
    <div class="logo" />
    <a-menu
      theme="dark"
      mode="horizontal"
      :defaultSelectedKeys="['2']"
      :style="{ lineHeight: '64px' }"
    >
      <a-menu-item key="1">首页</a-menu-item>
      <a-menu-item key="2">分类</a-menu-item>
      <a-menu-item key="3">归档</a-menu-item>
    </a-menu>
  </a-layout-header>
</template>

<script>
    export default {
        name: "Nav"
    }
</script>

<style scoped>
</style>
