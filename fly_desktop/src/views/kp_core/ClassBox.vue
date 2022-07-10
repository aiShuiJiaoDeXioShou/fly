<!-- 分类盒子 -->
<template>
    <div class="class_box_root">
        <!-- 分类盒子，分类盒子里面还能够放另一个分类盒子，还有任意卡片，什么元素都能够放里面 -->
        <web-box-vue></web-box-vue>
    </div>
</template>

<script>
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
        WebBoxVue
    },
    methods: {
    }
}

</script>

<style lang='less' scoped>
.class_box_root {
    width: 100%;
    height: 100%;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
}
</style>
