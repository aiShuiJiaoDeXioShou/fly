<!-- 物品分类的模板 -->
<template>
    <div class="goods_stl_root">
        <div :class="card_load_an">
            <!-- 删除按钮 -->
            <button @click="deleteOne" :class="{ 'delete': true, 'goods_delete_color': !is_good_side }"></button>
            <div class="goods_positive" v-if="!is_good_side">
                <!-- 我们点进去还能够看到他的子类,点击头像查看子类 -->
                <figure @click="flip_card" class="image is-128x128 block">
                    <img src="../../assets/icon.png" alt="test">
                </figure>

                <item-form-vue />

                <div class="block character-tag">

                    <!-- 人物标签 -->
                    <span class="tag is-primary is-light">
                        <button class="delete is-small"></button>
                        <div class="is-build-data">Primary</div>
                        <div class="is-build">
                            <input class="input is-link" type="text" placeholder="标签:">
                        </div>
                    </span>

                    <span class="tag is-link is-light">
                        <button class="delete is-small"></button>
                        <div class="is-build-data">Link</div>
                        <div class="is-build">
                            <input class="input is-link" type="text" placeholder="标签:">
                        </div>
                    </span>

                    <span class="tag is-info is-light">
                        <button class="delete is-small"></button>
                        <div class="is-build-data">Info</div>
                        <div class="is-build">
                            <input class="input is-link" type="text" placeholder="标签:">
                        </div>
                    </span>

                    <span class="tag is-success is-light">
                        <button class="delete is-small"></button>
                        <div class="is-build-data">Success</div>
                        <div class="is-build">
                            <input class="input is-link" type="text" placeholder="标签:">
                        </div>
                    </span>

                    <span class="tag is-warning is-light">
                        <button class="delete is-small"></button>
                        <div class="is-build-data">Warning</div>
                        <div class="is-build">
                            <input class="input is-link" type="text" placeholder="标签:">
                        </div>
                    </span>

                    <span class="tag is-danger is-light">
                        <button class="delete is-small"></button>
                        <div class="is-build-data">Danger</div>
                        <div class="is-build">
                            <input class="input is-link" type="text" placeholder="标签:">
                        </div>
                    </span>

                </div>
            </div>

            <div class="good_side" v-if="is_good_side">
                <!-- 这一面主要是这个详情信息了 -->
                <el-tabs v-model="activeName" @tab-click="handleClick">
                    <el-tab-pane label="属性" name="fourth">
                        <item-form-vue></item-form-vue>
                    </el-tab-pane>
                    <el-tab-pane label="身体" name="first">
                        <item-form-vue></item-form-vue>
                    </el-tab-pane>
                    <el-tab-pane label="背景" name="second">
                        <item-text-vue></item-text-vue>
                    </el-tab-pane>
                    <el-tab-pane label="背包" name="third">
                        <item-form-vue></item-form-vue>
                    </el-tab-pane>
                </el-tabs>

                <div id="good_side_content">
                    <!-- <item-form-vue></item-form-vue> -->
                </div>

            </div>

        </div>
    </div>
</template>

<script>
import $ from 'jquery';
import ItemFormVue from './ItemForm.vue';
// 导入特殊的列点击效果，双击开启编辑，或者双击关闭编辑
import { columnClick } from '../../comm/special_action/sp_click_action.js';
import ItemTextVue from './ItemText.vue';


export default {
    props: ["local_bg_color"],
    data() {
        return {
            is_good_side: false,
            card_load_an: ['goods_stl_box', 'box', 'notification', 'animate__animated'],
            activeName: 'fourth',
        }
    },
    methods: {
        // 设置背景颜色的函数
        setBgColor() {
            if (this.local_bg_color !== undefined) {
                $('.goods_stl_box').css('background-color', this.local_bg_color)
            }
        },

        flip_card() {
            // 点击该卡片时赋予翻转动画
            this.card_load_an.push('animate__flipInX')
            // 0.7秒过后删除flipInY动画
            setTimeout(() => {
                this.card_load_an.pop()
            }, 700)
            // 翻页过后显示反面
            this.is_good_side = !this.is_good_side
        },
        deleteOne() {
            if (this.is_good_side) {
                this.flip_card()
                return
            }
            // 弹出对话框
            this.$confirm('此操作将永久删除该卡片, 是否继续?', '提示', {
                confirmButtonText: '确定',
                cancelButtonText: '取消',
                type: 'warning'
            }).then(() => {
                this.$message({
                    type: 'success',
                    message: '删除成功!'
                });
            }).catch(() => {
                this.$message({
                    type: 'info',
                    message: '已取消删除'
                });
            });
        },
        editAfter() {
            console.log("编辑了这个值");
        },
        handleClick(tab, event) {
            console.log(tab, event);
        }

    },
    mounted() {
        columnClick({
            afterEditingFunction() {
                console.log("编辑了这个值");
            }
        })
        this.setBgColor()
    },
    components: {
        ItemFormVue,
        ItemTextVue
    }
}

</script>

<style lang='less'>
.goods_stl_box {
    padding: 10px;
    background-color: #69c0ff;
    min-width: 20vw;
    max-width: 74vw;
}

.goods_delete_color {
    background-color: #b5f5ec;
}

.goods_stl_root {
    min-height: 100vh;
    height: 100%;

    .delete {
        z-index: 1;
    }
}


.goods_positive {



    // 人物标签样式
    .character-tag {
        display: flex;
        flex-wrap: wrap;


        &>span {
            display: block;
            margin-right: 8px;
            margin-top: 5px;
            position: relative;

            .delete {
                position: absolute;
                top: -7px;
                right: 0;
            }
        }
    }
}

.is-build {
    display: none;
}
</style>
