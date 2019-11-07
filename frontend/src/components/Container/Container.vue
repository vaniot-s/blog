<template>
<div class="container">
  <a-row>
    <a-col :xs="24" :sm="24" :md="24" :lg="16" :xl="16">
      <Post></Post>
      <Post></Post>
      <Post></Post>
      <Post></Post>
      <List></List>
    </a-col>
    <a-col :xs="0" :sm="0" :md="0" :lg="2" :xl="2"  ></a-col>
    <a-col :xs="24" :sm="24" :md="24" :lg="6" :xl="6"  >
      <a-card hoverable size="default">
        <img
          alt="example"
          src="https://gw.alipayobjects.com/zos/rmsportal/JiqGstEfoWAOHiTxclqi.png"
          slot="cover"
        />
        <template class="ant-card-actions" slot="actions">
          <a-icon type="setting" />
          <a-icon type="edit" />
          <a-icon type="ellipsis" />
        </template>
        <a-card-meta title="Card title" description="This is the description">
          <a-avatar
            slot="avatar"
            src="https://zos.alipayobjects.com/rmsportal/ODTLcjxAfvqbxHnVXCYX.png"
          />
        </a-card-meta>
      </a-card>
    </a-col>
  </a-row>
　
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
