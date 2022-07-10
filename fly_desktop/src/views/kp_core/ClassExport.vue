<!-- 分类盒子 -->
<template>
    <div>
        <button class="button is-dark" @click="toImage">html转图片</button>
        <!-- 这个盒子应该能够自由扩展，我们应该怎么写呢，它应该有简介跟详情 -->
        <!-- 底部操作栏应该是 详情 删除 修改 管理子类 -->
        <div class="class_box_root" id="class_box_root">
            <goods-sTL></goods-sTL>
        </div>

        <!-- 显示捕获的image图片 -->
        <el-dialog title="提示" width="100%" :visible.sync="dialogImageVisible">
            <div id="showBoxItemImage"></div>
            <span slot="footer" class="dialog-footer">
                <el-button @click="dialogImageVisible = false">取 消</el-button>
                <el-button type="primary" @click="dialogImageVisible = false">确 定</el-button>
            </span>
        </el-dialog>

        <!-- 分类盒子，分类盒子里面还能够放另一个分类盒子，还有任意卡片，什么元素都能够放里面 -->
        <web-box-vue></web-box-vue>
    </div>
</template>

<script>
import GoodsSTL from '../../components/goods/GoodsSTL.vue';
import { htmlToImg } from '../../comm/tools';
import WebBoxVue from '../../components/card/WebBox.vue';

export default {
    data() {
        return {
            dialogImageVisible: false,
        }
    },
    mounted() {

    },
    components: {
        GoodsSTL,
        WebBoxVue
    },
    methods: {
        toImage() {
            htmlToImg('#goods_stl_root', '#showBoxItemImage');
            this.dialogImageVisible = true;
        }
    }
}

</script>

<style lang='less' scoped>
#showBoxItemImage {
    text-align: center;
}

#showBoxItemImage img {
    width: 100%;
    height: auto;
}
</style>
