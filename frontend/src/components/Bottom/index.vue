<template>
  <a-layout-footer class="copyright">
    <a-row >
    <a-col :xs="4" :sm="4" :md="4" :lg="4" :xl="4"></a-col>
    <a-col :xs="16" :sm="16" :md="16" :lg="16" :xl="16">
      <a-row >
        <a-col > ©2019 Your Vaniot.</a-col>
        <a-col > Template design by Vaniot.</a-col>
      </a-row>
    </a-col>
    <a-col :xs="4" :sm="4" :md="4" :lg="4" :xl="4"></a-col>
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
