<!-- 查看起点小说的排行榜 -->
<template>
    <div class="book_sort" v-loading="loading">

        <table class="yuep table">
            <tr>
                <th>序号</th>
                <th>书名</th>
                <th>月票数</th>
            </tr>
            <tr v-for="(item, index) in list" :key="index">
                <td>{{ item.sequenceId }}</td>
                <td>{{ item.novelName }}</td>
                <td>{{ item.monthTicket }}</td>
            </tr>
        </table>

    </div>
</template>

<script>
import axios from 'axios'
export default {
    data() {
        return {
            list: [],
            loading: false
        }
    },
    created() {
        this.loading = true
        axios.get('/api/qidian/yuep')
            .then(response => {
                this.list = response.data
                this.loading = false
                console.log(response)
            })
            .catch(error => {
                console.log(error)
            })
    },
}

</script>

<style lang='less' scoped>
.yuep {
  margin: 0 auto;
}
.book_sort {
    min-width: 30%;
    min-height: 40%;
}
</style>
