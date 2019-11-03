<template>
  <div class="container">
    <el-row :gutter="32">
      <el-col :xs="32" :sm="32" :lg="32" >
        <el-card class="box-card">
          <div slot="header" class="clearfix">
            <span>文章标题</span>
          </div>
          <div v-for="o in 4" :key="o" class="text item">
            {{'列表内容 ' + o }}
          </div>
        </el-card>
      </el-col>
      <el-col :xs="32" :sm="32" :lg="32" >
        <el-card class="box-card">
          <div slot="header" class="clearfix">
            <span>文章标题</span>
          </div>
          <div v-for="o in 4" :key="o" class="text item">
            {{'列表内容 ' + o }}
          </div>
        </el-card>
      </el-col>
      <el-col :xs="32" :sm="32" :lg="32" >
        <el-card class="box-card">
          <div slot="header" class="clearfix">
            <span>文章标题</span>
          </div>
          <div v-for="o in 4" :key="o" class="text item">
            {{'列表内容 ' + o }}
          </div>
        </el-card>
      </el-col>
      <el-col :xs="32" :sm="32" :lg="32" >
        <el-card class="box-card">
          <div slot="header" class="clearfix">
            <span>文章标题</span>
          </div>
          <div v-for="o in 4" :key="o" class="text item">
            {{'列表内容 ' + o }}
          </div>
        </el-card>
      </el-col>
      <el-col :xs="32" :sm="32" :lg="32" >
        <el-card class="box-card">
          <div slot="header" class="clearfix">
            <span>文章标题</span>
          </div>
          <div v-for="o in 4" :key="o" class="text item">
            {{'列表内容 ' + o }}
          </div>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script>
    export default {
        name: "List"
    }
</script>

<style scoped>

</style>
