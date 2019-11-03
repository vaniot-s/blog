<template>
  <div id="copyright">
    <div class="container">
      <el-row :gutter="32">
        <el-col :xs="24" :sm="24" :lg="10" >©2019 Your Vaniot.</el-col>
        <el-col :xs="24" :sm="24" :lg="10" >Template design by Vaniot.</el-col>
      </el-row>
    </div>
  </div>
</template>

<script>
    export default {
        name: "Bottom"
    }
</script>

<style scoped>
  #copyright {
    background: #333;
    color: #ccc;
    padding: 20px 0;
    font-size: 0.738rem;
  }

  #copyright p {
    margin: 0;
  }
</style>
