// 全局导入electron模块
const electron = require('electron')

module.exports = {
    install(Vue) {
        Object.defineProperties(Vue.prototype, {
            $electron: {
                get() {
                    return electron
                },
            },
        })
    },
}