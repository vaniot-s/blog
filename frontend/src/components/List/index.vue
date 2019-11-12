<template>
  <a-list :grid="{ gutter: 16, column: 1 }">
    <a-list-item  >
      <Carousel></Carousel>
    </a-list-item>
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
    import Carousel from "@/components/Carousel/index";


    export default {
        name: "List",
        components: {Carousel}
    };
</script>

<style scoped>

</style>
