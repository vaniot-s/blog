'use strict'
// 将会被depend的cli sever自动加载

const port = process.env.port || process.env.npm_config_port || 89// dev port


//　导出一个包含了选项的对象
// vue.config.js　文档地址:
module.exports = {
  // 部署应用包时的基本 URL
  publicPath: './',
  // 当运行 vue-cli-service build 时生成的生产环境构建文件的目录。注意目标目录在构建之前会被清除 (构建时传入 --no-clean 可关闭该行为)
  outputDir: 'ui',
  assetsDir: 'static',

  // 是否在开发环境下通过 eslint-loader 在每次保存时 lint 代码。这个值会在 @vue/cli-plugin-eslint 被安装之后生效。
  lintOnSave: true,

  devServer: {
    port: port,
    open: true,
    overlay: {
      warnings: false,
      errors: true
    },
    // 请求的代理地址
    proxy: {
      '/api': {
        target: 'http://127.0.0.1:53008/',
        changeOrigin: true,
        pathRewrite: {'^/api': ''}
      },
    }
  }
}
