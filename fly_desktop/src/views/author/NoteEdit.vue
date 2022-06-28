<!-- 文本编辑器 -->
<template>
    <div>
        <mavon-editor ref=md @imgAdd="imgAdd" v-model="text" @change="changeData" style="min-height: 800px">
            <template slot="right-toolbar-before">
                <el-button icon="el-icon-s-promotion" circle type="button" @click="release()" aria-hidden="true"
                    title="发布">
                </el-button>
            </template>
        </mavon-editor>
    </div>
</template>

<script>
import axios from 'axios';
export default {
    data() {
        return {
            text: "",
            render: "",
            upload: ""
        }
    },
    methods: {
        changeData(_, render) {
            // value中是文本值,render是渲染出的html文本
            this.render = render;
        },
        imgAdd(pos, $file) {
            // 第一步.将图片上传到服务器.
            var formdata = new FormData();
            formdata.append('image', $file);
            axios({
                url: this.upload,
                method: 'post',
                data: formdata,
                headers: { 'Content-Type': 'multipart/form-data' },
            }).then((url) => {
                // 这个url 是文件上传成功之后的返回值
                console.log(url.data.data);
                this.$refs.md.$img2Url(pos, url.data.data);
            })
        }
    },
}

</script>

<style lang='less' scoped>

</style>
