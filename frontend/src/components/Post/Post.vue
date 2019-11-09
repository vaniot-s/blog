<template>
  <div>
    <a-row>
      <a-col :xs="24" :sm="24" :md="24" :lg="24" :xl="24">
        <a-list :grid="{ gutter: 16, column: 1 }" >
          <a-list-item >
            <a-card size="default">
              <Breadcrumb></Breadcrumb>
            </a-card>
          </a-list-item>
          <a-list-item>
            <a-card title="Comment" size="default">
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
          <a-list-item>
            <a-card title="Comment" size="default">
              <a-comment>
                <a-avatar
                  slot="avatar"
                  src="https://zos.alipayobjects.com/rmsportal/ODTLcjxAfvqbxHnVXCYX.png"
                  alt="Han Solo"
                />
                <div slot="content">
                  <a-form-item>
                    <a-textarea :rows="4" @change="handleChange" :value="value"></a-textarea>
                  </a-form-item>
                  <a-form-item>
                    <a-button htmlType="submit" :loading="submitting" @click="handleSubmit" type="primary">
                      Add Comment
                    </a-button>
                  </a-form-item>
                </div>
              </a-comment>
            </a-card>
          </a-list-item>
          <a-list-item>
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
          </a-list-item>
        </a-list>
      </a-col>
    </a-row>
  </div>
</template>

<script>
    import Carousel from "@/components/Carousel";
    import Breadcrumb from "@/components/Breadcrumb";

    export default {
        name: "Post",
        components: {Breadcrumb, Carousel }
    }
</script>

<style scoped>

</style>
