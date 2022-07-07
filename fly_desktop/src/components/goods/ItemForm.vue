<!-- 表单类型的可选项组件 -->
<template>
    <div class="item_from_root">
        <div>
            <i class="el-icon-circle-plus-outline el_plus"></i>
        </div>

        <div class="block columns">
            <div class="column is-2">姓名:</div>
            <div class="column is-build-data">林河</div>
            <div class="column is-build">
                <input class="input is-link" type="text" placeholder="姓名:">
            </div>
        </div>
        <div class="block columns">
            <div class="column is-2">简述:</div>
            <div class="column is-build-data">只要我相信正义，我就是无敌的化身，一个不想活的男人。</div>
            <div class="column is-build">
                <input class="input is-link" type="text" placeholder="简述:">
            </div>
        </div>

    </div>
</template>

<script>
// 导入特殊的列点击效果，双击开启编辑，或者双击关闭编辑
import { columnClick } from '../../comm/special_action/sp_click_action.js';
export default {
    data() {
        return {
        }
    },
    mounted() {
        columnClick({
            afterEditingFunction() {
                console.log("itemform 编辑了这个值");
            }
        })
    }
}

</script>

<style lang='less' scoped>
.el_plus {
    transition: .5s all;
    font-size: 20px;
    cursor: pointer;
}

// 旋转四十五度
.el_plus:hover {
    // 动漫效果一秒
    transition: .5s all;
    transform: rotate(45deg);
    color: #bae7ff;
}

.item_from_root {
    padding: 0 0 20px 0px;
}
</style>
