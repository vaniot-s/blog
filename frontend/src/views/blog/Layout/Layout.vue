<template>
  <div class="blog-layout">
    <Nav></Nav>
    <Container></Container>
    <Bottom></Bottom>
  </div>
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
