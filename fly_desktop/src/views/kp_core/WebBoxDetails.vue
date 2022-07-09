<!-- 这个是分类盒子详情 -->
<template>
    <div class="web_box_details_root">
        <div class="web_box_menu_root">

            <div class="select_item">
                <div :class="{ 'select_box': true, 'is_select': is_model }" @click="is_model = !is_model">
                    <i class="el-icon-menu"></i>
                </div>
                <div :class="{ 'select_box': true, 'is_select': !is_model }" @click="is_model = !is_model">
                    <i class="el-icon-s-unfold"></i>
                </div>
            </div>


            <div class="level-item go_through_motions" v-if="is_model">
                <div class="field has-addons">
                    <p class="control">
                        <input class="input" type="text" placeholder="支持正则">
                    </p>
                    <p class="control">
                        <button class="button" style="padding:0px 5px 0px 5px ;">
                            Search
                        </button>
                    </p>
                </div>
            </div>


            <div class="filter_select go_through_motions" v-if="is_model">
                <!-- 选择过滤的模块项 -->
                <el-select v-model="value" placeholder="请选择">
                    <el-option v-for="item in options" :key="item.value" :label="item.label" :value="item.value">
                    </el-option>
                </el-select>
            </div>

            <div class="attribute_view go_through_motions" v-if="is_model">
                <el-tag :key="tag" v-for="tag in dynamicTags" closable :disable-transitions="false"
                    @close="handleClose(tag)">
                    {{ tag }}
                </el-tag>
                <el-input class="input-new-tag" v-if="inputVisible" v-model="inputValue" ref="saveTagInput" size="small"
                    @keyup.enter.native="handleInputConfirm" @blur="handleInputConfirm">
                </el-input>
                <el-button v-else class="button-new-tag" size="small" @click="showInput">+ New Tag</el-button>
            </div>

        </div>

        <div class="items" v-if="is_model">
            <div class="item" v-for="(item, index) in 5" :key="index">
                <goods-s-t-l-vue />
            </div>
        </div>


        <div class="combination_panel" v-if="!is_model">
            <nav class="go_through_motions panel_filter is-info panel">
                <p class="panel-heading">
                    Repositories
                </p>
                <div class="panel-block">
                    <p class="control has-icons-left">
                        <input class="input" type="text" placeholder="Search">
                        <span class="icon is-left">
                            <i class="fas fa-search" aria-hidden="true"></i>
                        </span>
                    </p>
                </div>
                <p class="panel-tabs">
                    <a class="is-active">All</a>
                    <a>Public</a>
                    <a>Private</a>
                    <a>Sources</a>
                    <a>Forks</a>
                </p>
                <a class="panel-block is-active">
                    <span class="panel-icon">
                        <i class="fas fa-book" aria-hidden="true"></i>
                    </span>
                    bulma
                </a>
                <a class="panel-block">
                    <span class="panel-icon">
                        <i class="fas fa-book" aria-hidden="true"></i>
                    </span>
                    marksheet
                </a>
                <a class="panel-block">
                    <span class="panel-icon">
                        <i class="fas fa-book" aria-hidden="true"></i>
                    </span>
                    minireset.css
                </a>
                <a class="panel-block">
                    <span class="panel-icon">
                        <i class="fas fa-book" aria-hidden="true"></i>
                    </span>
                    jgthms.github.io
                </a>
                <a class="panel-block">
                    <span class="panel-icon">
                        <i class="fas fa-code-branch" aria-hidden="true"></i>
                    </span>
                    daniellowtw/infboard
                </a>
                <a class="panel-block">
                    <span class="panel-icon">
                        <i class="fas fa-code-branch" aria-hidden="true"></i>
                    </span>
                    mojs
                </a>
                <label class="panel-block">
                    <input type="checkbox">
                    remember me
                </label>
                <div class="panel-block">
                    <button class="button is-link is-outlined is-fullwidth">
                        Reset all filters
                    </button>
                </div>
            </nav>

            <div class="panel_content tile is-parent is-vertical">
                <div class="tile is-child">
                    <goods-s-t-l-vue />
                </div>
                <div class="tile is-child">
                    <goods-s-t-l-vue />
                </div>
                <div class="tile is-child">
                    <goods-s-t-l-vue />
                </div>
                <div class="tile is-child">
                    <goods-s-t-l-vue />
                </div>

            </div>
        </div>
    </div>
</template>

<script>
import GoodsSTLVue from '../../components/goods/GoodsSTL.vue'

export default {
    data() {
        return {
            is_model: true,
            dynamicTags: ['标签一', '标签二', '标签三'],
            inputVisible: false,
            inputValue: '',
            value: '',
            options: []
        }
    },
    methods: {
        handleClose(tag) {
            this.dynamicTags.splice(this.dynamicTags.indexOf(tag), 1);
        },

        showInput() {
            this.inputVisible = true;
            this.$nextTick(() => {
                this.$refs.saveTagInput.$refs.input.focus();
            });
        },

        handleInputConfirm() {
            let inputValue = this.inputValue;
            if (inputValue) {
                this.dynamicTags.push(inputValue);
            }
            this.inputVisible = false;
            this.inputValue = '';
        }
    },
    mounted() {
    },
    components: {
        GoodsSTLVue,
    }
}

</script>

<style lang='less' scoped>
.web_box_details_root {
    min-height: 100vh;
    height: 100%;
    background-color: #d5ebe1;
    width: 80vw;

    .web_box_menu_root {
        display: flex;

        .attribute_view {
            margin: 14px 0px 0px 14px;
            min-width: 40%;
            max-width: 50%;
            border-radius: 3px;
            display: flex;
            overflow: scroll;
        }

        .filter_select {
            margin: 14px;
            text-align: center;
            line-height: 20px;
        }
    }

    .items {
        display: flex;
        flex-direction: row;
        flex-wrap: wrap;
        justify-content: space-around;
        align-items: center;
        padding: 10px;

        .item {
            margin: 5px;
        }
    }

    .select_item {
        margin: 10px;
        width: 100px;
        border-radius: 3px;
        padding: 5px 10px 5px 10px;
        background: #fff;
        display: flex;
        // 居中
        justify-content: center;
        align-items: center;

        .select_box {
            padding: 5px;
        }

        .is_select {
            background-color: rgba(0, 0, 0, 0.2);
            border-radius: 3px;
        }

    }

    .el-tag {
        padding: 0 5px;
        margin-left: 10px;
    }

    .button-new-tag {
        margin-left: 10px;
        height: 32px;
    }

    .input-new-tag {
        width: 90px;
        margin-left: 10px;
        vertical-align: bottom;
    }

    .combination_panel {
        display: flex;
        flex-wrap: wrap;
        padding: 10px;

        .panel_filter {
            width: 300px;
            margin-right: 20px;
        }

        .panel_content {
            width: 65%;
        }
    }


}
</style>
