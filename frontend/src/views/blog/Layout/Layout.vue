<template>
  <a-layout>
    <Nav></Nav>
    <a-layout-content :style="{ padding: '0 100px', marginTop: '64px' }">
      <Container></Container>
      <Bottom></Bottom>
    </a-layout-content>
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
.blog-layout {
  background-color: #F1F1F1;
}
</style>
