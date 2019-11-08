<template>
  <a-layout :style="{ background: '#F1F1F1' }">
    <Nav></Nav>
    <a-layout-content >
      <a-row :style="{  marginTop: '84px', }">
        <a-col :xs="4" :sm="4" :md="4" :lg="4" :xl="4"></a-col>
        <a-col :xs="16" :sm="16" :md="16" :lg="16" :xl="16">
          <Container></Container>
        </a-col>
        <a-col :xs="4" :sm="4" :md="4" :lg="4" :xl="4"></a-col>
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
