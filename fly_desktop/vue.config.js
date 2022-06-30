const path = require('path')

function addStyleResource(rule) {
  rule.use('style-resource')
    .loader('style-resources-loader')
    .options({
      patterns: [
        // 需要全局导入的less路径，自己修改，我这里引入了两个less文件
        path.resolve(__dirname, './src/assets/css/global.less'),
        path.resolve(__dirname, './src/assets/css/variables.less')
      ],
    })
}

module.exports = {

  publicPath: './',
  lintOnSave: true,
  pages: {
    index: {
      entry: 'src/main.js'
    }
  },
  chainWebpack: config => {
    const types = ['vue-modules', 'vue', 'normal-modules', 'normal']
    types.forEach(type => addStyleResource(config.module.rule('less').oneOf(type)))
  },
  configureWebpack: {
    //关闭 webpack 的性能提示
    performance: {
      hints: false
    }
  },
  devServer: {
    proxy: {
      '/api': {
        target: 'http://127.0.0.1:41/', //后端接口地址
        changeOrigin: true, //是否允许跨越
        pathRewrite: {
          '^/api': '/api' //重写,
        },
      }
    }
  },
}