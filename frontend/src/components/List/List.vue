<template>
  <a-list :grid="{ gutter: 16, column: 1 }" :dataSource="data">
    <a-list-item slot="renderItem" slot-scope="item, index">
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
        <a-card-meta title="item" description="This is the description">
          <a-avatar
            slot="avatar"
            src="https://zos.alipayobjects.com/rmsportal/ODTLcjxAfvqbxHnVXCYX.png"
          />
        </a-card-meta>
      </a-card>
    </a-list-item>
  </a-list>
</template>

<script>
    const data = [
        {
            title: 'Title 1',
        },
        {
            title: 'Title 2',
        },
        {
            title: 'Title 3',
        },
        {
            title: 'Title 4',
        },
    ];
    export default {
        name: "List",
        data() {
            return {
                data,
            };
        },
    };
</script>

<style scoped>

</style>
