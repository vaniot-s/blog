<template>
<div class="container" >
  <a-row>
    <a-col :xs="24" :sm="24" :md="24" :lg="17" :xl="17">
      <router-view :key="key" />
    </a-col>
    <a-col :xs="0" :sm="0" :md="0" :lg="1" :xl="1"  ></a-col>
    <a-col :xs="24" :sm="24" :md="24" :lg="5" :xl="5"  >
      <a-list :grid="{ gutter: 16, column: 1 }" >
        <a-list-item >
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
        </a-list-item>
        <a-list-item >
          <TagsCard></TagsCard>
        </a-list-item>
        <a-list-item >
          <ArchesCard></ArchesCard>
        </a-list-item>
        <a-list-item >
          <FriendLink></FriendLink>
        </a-list-item>
      </a-list>
    </a-col>
  </a-row>
</div>
</template>

<script>
    import TagsCard from "@/components/TagsCard";
    import ArchesCard from "@/components/ArchesCard";
    import FriendLink from "@/components/FriendLink/index";
    export default {
        name: "Container",
        components: {FriendLink, TagsCard, ArchesCard},
        computed: {
            key() {
                return this.$route.path
            }
        }
    }
</script>

<style scoped>
.container{
  margin-top: 20px;
}
</style>
