<template>
<div>
<Post></Post>
  <List></List>
</div>
</template>

<script>
    import Post from "../Post/Post";
    import List from "../List/List";
    export default {
        name: "Container",
        components: {List, Post}
    }
</script>

<style scoped>

</style>
