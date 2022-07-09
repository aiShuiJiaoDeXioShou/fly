<!--  -->
<template>
    <div id="web_box_root" class="box">
        <div class="box_icon"></div>
        <!-- 这个是集合的图片部分 -->
        <div class="box_image">
            <figure class="image_ image is-128x128">
                <img style="width: 100%; height: auto;" src="../../assets/icon.png" alt="图片">
                <div class="web_box_title">你好世界</div>
            </figure>
            <div class="box_image_action">
                <i @click="openDetails" @mouseout="is_open_folder = !is_open_folder"
                    @mouseover="is_open_folder = !is_open_folder"
                    :class="{ 'el-icon-folder': !is_open_folder, 'el-icon-folder-opened': is_open_folder }"></i>
                <i @click="time_click" class="el-icon-alarm-clock"></i>
                <i @click="statistics" class="el-icon-data-analysis"></i>
                <i @click="extend" class="el-icon-files"></i>
            </div>
        </div>

        <!-- 时间统计 -->
        <div class="hour animate__animated animate__fadeInUp" v-if="is_move_image_box">
            <h4 class="title is-5">花开花谢，大家一起这么久了...</h4>
            <p class="number">
                1 天 8 小时 45 分钟 23秒
            </p>
        </div>

        <!-- 这个是内容部分 -->
        <div :class="box_content" v-if="!is_box_extend">
            <!-- 内容的小导航 -->
            <el-tabs v-model="tabMenu" class="box_content_menu" v-if="!is_content_side">
                <el-tab-pane label="介绍" name="1">
                    <div class="item_content animate__animated animate__fadeInUp">
                        因为加载图像可能需要几秒钟（或者根本不需要），所以请使用图像容器指定精确大小的容器，这样布局就不会因为图像加载或图像错误而中断。
                        因为加载图像可能需要几秒钟（或者根本不需要），所以请使用图像容器指定精确大小的容器，这样布局就不会因为图像加载或图像错误而中断。
                        因为加载图像可能需要几秒钟（或者根本不需要），所以请使用图像容器指定精确大小的容器，这样布局就不会因为图像加载或图像错误而中断。
                        因为加载图像可能需要几秒钟（或者根本不需要），所以请使用图像容器指定精确大小的容器，这样布局就不会因为图像加载或图像错误而中断。
                        因为加载图像可能需要几秒钟（或者根本不需要），所以请使用图像容器指定精确大小的容器，这样布局就不会因为图像加载或图像错误而中断。
                        因为加载图像可能需要几秒钟（或者根本不需要），所以请使用图像容器指定精确大小的容器，这样布局就不会因为图像加载或图像错误而中断。
                        因为加载图像可能需要几秒钟（或者根本不需要），所以请使用图像容器指定精确大小的容器，这样布局就不会因为图像加载或图像错误而中断。
                        因为加载图像可能需要几秒钟（或者根本不需要），所以请使用图像容器指定精确大小的容器，这样布局就不会因为图像加载或图像错误而中断。
                    </div>
                </el-tab-pane>
                <el-tab-pane label="经历" name="2">
                    <div class="item_content animate__animated animate__fadeInUp"></div>
                </el-tab-pane>
            </el-tabs>
            <!-- 背面统计面板 -->
            <div v-if="is_content_side">

            </div>
        </div>

        <!-- 下载，收藏，观看,分享 -->
        <div :class="box_extend" v-if="is_box_extend">
            <div class="box_extend_item">
                <el-button type="primary" circle icon="el-icon-download"></el-button>
            </div>
            <div class="box_extend_item">
                <el-button type="warning" circle icon="el-icon-star-off"></el-button>
            </div>
            <div class="box_extend_item">
                <el-button type="success" circle icon="el-icon-video-camera"></el-button>
            </div>
            <div class="box_extend_item">
                <el-button type="info" circle icon="el-icon-share"></el-button>
            </div>
        </div>
    </div>
</template>

<script>
export default {
    data() {
        return {
            tabMenu: '1',
            is_open_folder: false,
            is_move_image_box: false,
            box_content: ['box_content', 'notification', 'animate__animated'],
            is_content_side: false,
            is_box_extend: false,
            box_extend: ['box_extend', 'notification', 'animate__animated']
        }
    },
    methods: {
        time_click() {
            let imageBox = document.querySelector(".box_image")
            if (this.is_move_image_box) {
                imageBox.classList.remove('move_200px')
                this.is_move_image_box = false
                return
            }
            imageBox.classList.add('move_200px')
            this.is_move_image_box = true
        },
        statistics() {
            // 点击该卡片时赋予翻转动画
            this.box_content.push('animate__flipInX')
            // 0.7秒过后删除flipInY动画
            setTimeout(() => {
                this.box_content.pop()
            }, 700)
            // 翻页过后显示反面
            this.is_content_side = !this.is_content_side
        },
        extend() {
            if (this.is_box_extend) {
                this.box_extend.push('animate__flipInX')
                this.is_box_extend = false
                this.box_extend.pop()
                setTimeout(() => {
                    this.box_content.pop()
                }, 700)
                return
            }

            // 点击该卡片时赋予淡出动画
            this.box_content.push('animate__flipInX')
            // 0.7秒过后删除animate__fadeInUp动画
            this.is_box_extend = true
            setTimeout(() => {
                this.box_content.pop()
            }, 700)
        },
        // 打开详情页面
        openDetails() {
            this.$router.push({
                path: '/web_box_details'
            })
        }
    }
}

</script>

<style lang='less' scoped>
@web_box_root_height: 30vh;



#web_box_root {
    display: flex;
    flex-direction: row;
    min-width: 200px;
    min-height: 100px;
    height: @web_box_root_height;
    max-width: 60vw;
    max-height: 30vh;
    background-color: #adc6ff;
    padding: 10px 20px 10px 20px;
    margin: 5px;
    border-radius: 3.5px;
    position: relative;

    .box_icon {
        background: url('../../assets/phl.png') no-repeat;
        background-size: contain;
        background-position: center;
        width: 130px;
        height: 100px;
        position: absolute;
        top: -53px;
        left: -39px;
    }

    .box_content {
        width: 80vw;
        .global_card_transparent();


        .box_content_menu {
            font-size: 13px;
            letter-spacing: 2px;

            .item_content {
                // 首行缩进，隐藏超出内容
                max-height: 120px;
                text-indent: 2vw;
                // 超出部分隐藏
                overflow-y: scroll;
                text-overflow: ellipsis;
            }

        }
    }

    .box_extend {
        width: 80vw;
        .global_card_transparent();
    }

    .box_image {
        width: 20vw;
        margin-right: 20px;
        background-color: rgba(255, 255, 255, 0.2);
        box-shadow: 0 15px 20px rgba(0, 0, 0, 0.2);
        transition: 1s all;

        .image_ {
            position: relative;

            .web_box_title {
                position: absolute;
                bottom: 0;
                left: 0;
                width: 100%;
                height: 0%;
                background-color: rgba(0, 0, 0, 0.5);
                color: #fff;
                text-align: center;
                text-shadow: 0 0 10px #000;
                transition: .5s all;
                font-size: 0;
            }

            &:hover {
                .web_box_title {
                    height: 20%;
                    transition: .5s all;
                    z-index: 1;
                    opacity: 1;
                    font-size: 18px;
                }
            }
        }
    }



    .box_image_action {
        display: flex;
        flex-wrap: wrap;
        margin-top: 10px;

        &>i {
            display: block;
            font-size: 15px;
            cursor: pointer;

            &:hover {
                color: #1890ff;
            }

            margin-right: 5px;
        }
    }

    .hour {
        .title {
            font-size: 10px;
            font-weight: bold;
            color: #6b798e;
        }

        .number {
            width: 100px;
            font-size: 15px;
            font-weight: bold;
            color: #6b798e;
            border-radius: 0;
            // 背景颜色为无色
            background-color: transparent;
        }

        position: absolute;
        left: 0;
        bottom: 20px;
    }

    .box_extend {
        width: 80vw;
        display: flex;
        flex-wrap: wrap;
        // 垂直居中
        justify-content: center;
        align-items: center;

        &>.box_extend_item {
            margin-right: 10px;

            button {
                width: 40px;
                height: 40px;
                text-align: center;

            }
        }
    }

}

.move_200px {
    transform: translateY(-100px);
    transition: 1s all;
}

@media (max-width: 580px) {
    #web_box_root {
        flex-direction: column;
        height: auto;
        max-height: 100vh;

        .box_content {
            width: 100%;
        }

        .box_extend {
            width: 100%;
        }

        .hour {
            top: 70px;
        }
    }
}
</style>
