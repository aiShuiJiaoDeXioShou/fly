<!-- 分类盒子 -->
<template>
    <div>
        <!-- 这个盒子应该能够自由扩展，我们应该怎么写呢，它应该有简介跟详情 -->
        <!-- 底部操作栏应该是 详情 删除 修改 管理子类 -->
        <div class="class_box_root">

        </div>
    </div>
</template>

<script>
export default {
    data() {
        return {
        }
    },
    mounted() {
        console.log("你是住");
    }
}

</script>

<style lang='less' scoped>
.class_box_root {
    background: red;
    height: 500px;
    width: 600px;
}
</style>
