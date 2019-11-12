<template>
  <a-layout :style="{ background: '#F1F1F1' }">
    <Nav></Nav>
    <a-layout-content >
      <a-row :style="{  marginTop: '84px', }">
        <a-col :xs="{ span: 22, offset: 1 }" :sm="{ span: 20, offset: 2 }" :md="{ span: 20, offset: 2 }" :lg="{ span: 20, offset: 2 }" :xl="{ span: 18, offset: 3 }">
          <Container></Container>
        </a-col>
      </a-row>
    </a-layout-content>
    <Bottom></Bottom>
  </a-layout>
</template>

<script>
  import Nav from "@/components/Nav/Nav";
  import Bottom from  "@/components/Bottom"
  import Container from "@/components/Container/Container";
  export default {
      name: "Layout",
      components: {Container, Bottom, Nav},
  }
</script>

<style scoped>
</style>
