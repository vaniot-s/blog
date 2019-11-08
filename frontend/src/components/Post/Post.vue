<template>
  <div>
    <a-row>
      <a-col :xs="24" :sm="24" :md="24" :lg="24" :xl="24">
        <Carousel></Carousel>
      </a-col>
      <a-col :xs="24" :sm="24" :md="24" :lg="24" :xl="24">
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
      <a-col :xs="24" :sm="24" :md="24" :lg="24" :xl="24">
        <a-card title="Comment" size="default">
          <a-comment>
            <span slot="actions">Reply to</span>
            <a slot="author">Han Solo</a>
            <a-avatar
              slot="avatar"
              src="https://zos.alipayobjects.com/rmsportal/ODTLcjxAfvqbxHnVXCYX.png"
              alt="Han Solo"
            />
            <p slot="content">
              We supply a series of design principles, practical patterns and high quality design resources
              (Sketch and Axure).
            </p>
            <a-comment>
              <span slot="actions">Reply to</span>
              <a slot="author">Han Solo</a>
              <a-avatar
                slot="avatar"
                src="https://zos.alipayobjects.com/rmsportal/ODTLcjxAfvqbxHnVXCYX.png"
                alt="Han Solo"
              />
              <p slot="content">
                We supply a series of design principles, practical patterns and high quality design
                resources (Sketch and Axure).
              </p>
              <a-comment>
                <span slot="actions">Reply to</span>
                <a slot="author">Han Solo</a>
                <a-avatar
                  slot="avatar"
                  src="https://zos.alipayobjects.com/rmsportal/ODTLcjxAfvqbxHnVXCYX.png"
                  alt="Han Solo"
                />
                <p slot="content">
                  We supply a series of design principles, practical patterns and high quality design
                  resources (Sketch and Axure).
                </p>
              </a-comment>
              <a-comment>
                <span slot="actions">Reply to</span>
                <a slot="author">Han Solo</a>
                <a-avatar
                  slot="avatar"
                  src="https://zos.alipayobjects.com/rmsportal/ODTLcjxAfvqbxHnVXCYX.png"
                  alt="Han Solo"
                />
                <p slot="content">
                  We supply a series of design principles, practical patterns and high quality design
                  resources (Sketch and Axure).
                </p>
              </a-comment>
            </a-comment>
          </a-comment>
        </a-card>

      </a-col>
    </a-row>

  </div>
</template>

<script>
    import Carousel from "@/components/Carousel";

    export default {
        name: "Post",
        components: { Carousel }
    }
</script>

<style scoped>

</style>
