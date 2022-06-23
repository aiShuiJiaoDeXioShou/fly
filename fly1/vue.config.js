module.exports = {
    publicPath: './',
    lintOnSave: true,
    pages: {
        index: {
            entry:'src/main.js'
        }
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
      }
}