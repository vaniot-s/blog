<template>
  <a-card title="Friendlink" :bordered="false" >
    <a-list itemLayout="horizontal" :dataSource="data">
      <a-list-item slot="renderItem" slot-scope="item, index">
        <a-list-item-meta
          description="Ant Design, a design language for background applications, is refined by Ant UED Team"
        >
          <a slot="title" href="https://vue.ant.design/">{{item.title}}</a>
        </a-list-item-meta>
      </a-list-item>
    </a-list>
  </a-card>
</template>

<script>
    const data = [
        {
            title: 'Ant Design Title 1',
        },
        {
            title: 'Ant Design Title 2',
        },
        {
            title: 'Ant Design Title 3',
        },
        {
            title: 'Ant Design Title 4',
        },
    ];
    export default {
        name: "FriendLink",
        data() {
            return {
                data,
            };
        },
    }
</script>

<style scoped>

</style>
