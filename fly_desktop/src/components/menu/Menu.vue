<!-- 这个是菜单的视窗组件 -->
<template>
    <div class="menu-root">
        <!-- 图标 -->
        <div class="menu-icon">
            <img src="../../assets/logo.png" alt="图标">
        </div>
        <!-- 菜单栏 -->
        <ul class="menu-list">
            <li>设置（S）</li>
            <li>插件（C）</li>
            <li>终端（T）</li>
            <li>帮助（H）</li>
        </ul>
        <!-- 常用操作 -->
        <div class="menu-operation">
            <el-tooltip content="更多选项" placement="bottom" effect="dark" :open-delay="open_delay">
                <div class="operation-item" @click="extend"></div>
            </el-tooltip>

            <el-tooltip content="最小化" placement="bottom" effect="dark" :open-delay="open_delay">
                <div class="operation-item" @click="min_win"></div>
            </el-tooltip>

            <el-tooltip content="最大化" placement="bottom" effect="dark" :open-delay="open_delay">
                <div class="operation-item" @click="max_win"></div>
            </el-tooltip>

            <el-tooltip content="关闭" placement="bottom" effect="dark" :open-delay="open_delay">
                <div class="operation-item" @click="close_win"></div>
            </el-tooltip>

        </div>
    </div>
</template>

<script>
const { remote } = window.require('electron')
export default {
    data() {
        return {
            open_delay: 700
        }
    },
    methods: {
        // 窗口最大化
        max_win() {
            const win = remote.getCurrentWindow();
            if (win.isMaximized()) { // 判断 窗口是否已最大化
                win.restore();// 恢复原窗口大小
            } else {
                win.maximize();  //最大化窗口
            }
        },
        // 窗口最小化
        min_win() {
            remote.getCurrentWindow().minimize();
        },
        // 窗口关闭
        close_win() {
            remote.getCurrentWindow().close();
        },
        // 菜单展开收起
        extend() {
            // 重新启动应用
            remote.app.relaunch({ args: process.argv.slice(1) });
            /* if (this.$refs.menu_root.classList.contains('menu-extend')) {
                this.$refs.menu_root.classList.remove('menu-extend');
            } else {
                this.$refs.menu_root.classList.add('menu-extend');
            } */
        }
    },
    mounted() {

    },
}

</script>

<style lang='less' scoped>
.menu-root {
    padding: 5px 10px 5px 10px;
    background: #b37feb;
    color: #fff;
    width: 100%;
    position: relative;
    font-size: @Menu_Font_Size;
    z-index: 999;

    .menu-icon {
        -webkit-app-region: no-drag;
        position: absolute;
        top: 0;
        left: 0;
        z-index: 5;
    }

    .menu-list {
        position: absolute;
        left: @Menu_H + 10px;
        top: 0;
        display: flex;
        justify-content: center;
        justify-items: center;
        align-items: center;

        &>li {
            -webkit-app-region: no-drag;
            margin-right: 5px;
            border-radius: 5px;
            cursor: pointer;
            line-height: @Menu_H;

            &:hover {
                background: rgba(0, 0, 0, 0.1);
            }
        }
    }

    .menu-icon>img {
        width: auto;
        height: @Menu_H;
    }

    .menu-operation {
        position: absolute;
        right: 0;
        top: 0;
        display: flex;
        justify-content: center;
        justify-items: center;
        align-items: center;
        border-radius: 5px;
        cursor: pointer;
        height: @Menu_H;

        &>.operation-item {
            @This_Width: @Menu_H * 0.5;
            -webkit-app-region: no-drag;
            margin-right: 10px;
            line-height: @Menu_H;
            cursor: pointer;
            width: @This_Width;
            height: @This_Width;
            background: #fff;
            border-radius: @This_Width;
        }

        //最后一个div背景色为红色
        &>.operation-item:last-child {
            background: #ff0000;

            // 当选中时颜色变深
            &:hover {
                background: #ff000066;
            }
        }

        //倒数第二个为黄色
        &>.operation-item:nth-last-child(3) {
            background: #ffc000;

            &:hover {
                background: #ffc00066;
            }
        }

        //倒数第三个为蓝色
        &>.operation-item:nth-last-child(2) {
            background: #95de64;

            &:hover {
                background: #95de6466;
            }
        }

        //倒数第四个为绿色
        &>.operation-item:nth-last-child(4) {
            background: #00b3ff;

            &:hover {
                background: #00b3ff66;
            }
        }

    }
}
</style>
