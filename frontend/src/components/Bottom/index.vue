<template>
  <a-layout-footer class="copyright">
    <a-row :gutter="24">
      <a-col :xs="24" :sm="24" :md="12" :lg="12" :xl="12"> ©2019 Your Vaniot.</a-col>
      <a-col :xs="24" :sm="24" :md="12" :lg="12" :xl="12"> Template design by Vaniot.</a-col>
    </a-row>
  </a-layout-footer>
</template>

<script>
    export default {
        name: "Bottom"
    }
</script>

<style scoped>
  .copyright {
    background: #333;
    color: #ccc;
    padding: 20px 0;
    font-size: 0.738rem;
  }
  .ant-row div{
    text-align: center;
  }
</style>
