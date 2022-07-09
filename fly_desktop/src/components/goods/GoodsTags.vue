<!-- 这个是人物tag，当然其他物品也行哦 -->
<template>
    <div>
        <div class="operation">
            <!-- 点击加号，然后出现新的一行，用Tap键位可以追加更新的一行，使用回车键位，完成整个更新操作 -->
            <i class="el-icon-circle-plus-outline el_plus" @click="addPlus()"></i>
            <i :class="{ 'el-icon-edit': is_et, 'el-icon-check': !is_et }" @click="edit_item"></i>
        </div>

        <div class="block character-tag">

            <!-- 人物标签 -->
            <span v-for="(item, index) in 6" :key="index" :class="getTagClass">
                <button v-if="is_del" class="delete is-small"></button>
                <div class="is-build-data">Primary</div>
                <div class="is-build">
                    <input class="input is-link" type="text" placeholder="标签:">
                </div>
            </span>

        </div>

        <!-- 添加tag的弹出窗 -->
        <el-dialog title="请输入标签的名称" :visible.sync="dialogVisible" width="30%" :append-to-body="true">
            <input class="input is-link" type="text" placeholder="标签:">
            <span slot="footer" class="dialog-footer">
                <el-button class="but" @click="dialogVisible = false">取 消</el-button>
                <el-button class="but" type="primary" @click="dialogVisible = false">确 定</el-button>
            </span>
        </el-dialog>

    </div>
</template>

<script>
export default {
    data() {
        return {
            is_et: true,
            is_del: false,
            tag_color: ['is-primary', 'is-link', 'is-info', 'is-success', 'is-warning', 'is-danger'],
            tag_class: ['tag', 'is-light'],
            dialogVisible: false,
        }
    },
    methods: {
        addPlus() {
            // 使用弹窗
            this.dialogVisible = true;
        },
        edit_item() {
            this.is_del = !this.is_del;
        },
        // tag颜色随机函数
        getTagColor() {
            let color = this.tag_color[Math.floor(Math.random() * this.tag_color.length)];
            return color;
        },
    },
    // 计算属性
    computed: {
        getTagClass() {
            return this.tag_class;
        }
    },
    mounted() {
        document.querySelectorAll('.tag').forEach(item => {
            item.classList.add(this.getTagColor());
        });
    }
}

</script>

<style lang='less' scoped>
.but {
    padding: 10px;
}

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
</style>
