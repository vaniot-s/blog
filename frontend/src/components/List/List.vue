<template>
  <div class="container">
    <a-list itemLayout="horizontal" >
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
<!--      <a-list-item slot="renderItem" slot-scope="item, index">-->
<!--        <a slot="actions">edit</a>-->
<!--        <a slot="actions">more</a>-->
<!--        <a-list-item-meta-->
<!--          description="Ant Design, a design language for background applications, is refined by Ant UED Team"-->
<!--        >-->
<!--          <a slot="title" href="https://vue.ant.design/">{{item.name.last}}</a>-->
<!--          <a-avatar-->
<!--            slot="avatar"-->
<!--            src="https://zos.alipayobjects.com/rmsportal/ODTLcjxAfvqbxHnVXCYX.png"-->
<!--          />-->
<!--        </a-list-item-meta>-->
<!--        <div>content</div>-->
<!--      </a-list-item>-->
    </a-list>
  </div>
</template>

<script>
    // import reqwest from 'reqwest';
    //
    // const fakeDataUrl = 'https://randomuser.me/api/?results=5&inc=name,gender,email,nat&noinfo';

    export default {
        name: "List",
        data() {
            return {
                loading: true,
                loadingMore: false,
                showLoadingMore: true,
                data: [],
            };
        },
        mounted() {
            this.getData(res => {
                this.loading = false;
                this.data = res.results;
            });
        },
        methods: {
            getData(callback) {
                reqwest({
                    url: fakeDataUrl,
                    type: 'json',
                    method: 'get',
                    contentType: 'application/json',
                    success: res => {
                        callback(res);
                    },
                });
            },
            onLoadMore() {
                this.loadingMore = true;
                this.getData(res => {
                    this.data = this.data.concat(res.results);
                    this.loadingMore = false;
                    this.$nextTick(() => {
                        window.dispatchEvent(new Event('resize'));
                    });
                });
            },
        },
    };
</script>

<style scoped>

</style>
