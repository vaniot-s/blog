<template>
  <div class="blog-layout">
    <Nav></Nav>
    <Bottom></Bottom>
  </div>
</template>

<script>
  import Nav from "@/components/Nav/Nav";
  import Bottom from  "@/components/Bottom"
  export default {
      name: "Layout",
      components: {Bottom, Nav},
  }
</script>

<style scoped>
.blog-layout {
  background-color: #F0F0F0;
}
</style>
