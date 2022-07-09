<!-- 表单类型的可选项组件 -->
<template>
    <div class="item_from_root">
        <div class="operation">
            <!-- 点击加号，然后出现新的一行，用Tap键位可以追加更新的一行，使用回车键位，完成整个更新操作 -->
            <i class="el-icon-circle-plus-outline el_plus" @click="addPlus()"></i>
            <i :class="{ 'el-icon-edit': is_et, 'el-icon-check': !is_et }" @click="edit_item"></i>
        </div>

        <div class="block columns">
            <div class="remove_item">
                <i class="el-icon-close"></i>
            </div>
            <div class="column is-2">姓名:</div>
            <div class="column is-build-data">林河</div>
            <div class="column is-build">
                <input class="input is-link" type="text" placeholder="姓名:">
            </div>
        </div>

        <div class="block columns">
            <div class="remove_item">
                <i class="el-icon-close"></i>
            </div>
            <div class="column is-2">简述:</div>
            <div class="column is-build-data">只要我相信正义，我就是无敌的化身，一个不想活的男人。</div>
            <div class="column is-build">
                <input class="input is-link" type="text" placeholder="简述:">
            </div>
        </div>

        <div class="block columns" v-if="is_add">
            <div class="column is-2">
                <input id="item_name" class="input is-link" type="text" placeholder="简述:">
            </div>
            <div class="column is-8">
                <input @keydown.enter="addItemData" @keydown.tab="tabNewRow" class="input is-link" type="text"
                    placeholder="简述:">
            </div>
        </div>

    </div>
</template>

<script>
// 导入特殊的列点击效果，双击开启编辑，或者双击关闭编辑
import { columnClick } from '../../comm/special_action/sp_click_action.js';
import $ from 'jquery';
export default {
    data() {
        return {
            is_et: true,
            is_add: false,
        }
    },
    mounted() {
        this.addItemClick();
    },
    methods: {
        edit_item() {
            // 获取item_from_root下的所有block属性
            let blocks = $('.item_from_root').find('.block');

            // 向这里面添加一个css元素
            blocks.toggleClass('global_card_transparent');
            blocks.toggleClass('go_shake');
            blocks.find('.remove_item').toggle();
            this.is_et = !this.is_et;
        },
        addPlus() {
            this.is_add = true;
        },
        // 添加元素点击事件
        addItemClick() {
            columnClick({
                afterEditingFunction() {
                    console.log("itemform 编辑了这个值");
                }
            })
        },
        addItemData() {
            this.is_add = false;
            console.log("添加数据了，你没看错！");
        },
        tabNewRow() {
            console.log("关闭了此行添加了下一行数据");
            // 零点五秒之后
            setTimeout(() => {
                // 获取input id 为item_name的值
                let item_name = $('#item_name')
                // item_name 设为下一个焦点
                item_name.focus();
            }, 10);
        }
    }
}

</script>

<style lang='less' scoped>
.el_plus {
    transition: .5s all;
    font-size: 15px;
    cursor: pointer;
}

// 旋转四十五度
.el_plus:hover {
    // 动漫效果一秒
    transition: .5s all;
    transform: rotate(45deg);
}

.item_from_root {
    padding: 0 0 20px 0px;

    .operation {
        display: flex;
        margin-bottom: 10px;
        font-size: 15px;
        color: #f5f5f5;

        &>i {
            display: block;
            margin-right: 5px;
        }

        &>i:hover {
            color: #bae7ff;
        }
    }

    .remove_item {
        color: red;
        margin-right: 5px;
        display: none;
    }
}
</style>
