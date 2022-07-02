<!-- 用于展示首页菜单 -->
<template>
    <div class="card-root">
        <article class="card card--1">
            <div class="card__info-hover">
                <div class="default_open">
                    <svg t="1656660024229" :class="{card__like:true,is_default_card:card.is_default}" @click="default_open" viewBox="0 0 1024 1024"
                        version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="3043" width="200" height="200">
                        <path
                            d="M923 283.6c-13.4-31.1-32.6-58.9-56.9-82.8-24.3-23.8-52.5-42.4-84-55.5-32.5-13.5-66.9-20.3-102.4-20.3-49.3 0-97.4 13.5-139.2 39-10 6.1-19.5 12.8-28.5 20.1-9-7.3-18.5-14-28.5-20.1-41.8-25.5-89.9-39-139.2-39-35.5 0-69.9 6.8-102.4 20.3-31.4 13-59.7 31.7-84 55.5-24.4 23.9-43.5 51.7-56.9 82.8-13.9 32.3-21 66.6-21 101.9 0 33.3 6.8 68 20.3 103.3 11.3 29.5 27.5 60.1 48.2 91 32.8 48.9 77.9 99.9 133.9 151.6 92.8 85.7 184.7 144.9 188.6 147.3l23.7 15.2c10.5 6.7 24 6.7 34.5 0l23.7-15.2c3.9-2.5 95.7-61.6 188.6-147.3 56-51.7 101.1-102.7 133.9-151.6 20.7-30.9 37-61.5 48.2-91 13.5-35.3 20.3-70 20.3-103.3 0.1-35.3-7-69.6-20.9-101.9zM512 814.8S156 586.7 156 385.5C156 283.6 240.3 201 344.3 201c73.1 0 136.5 40.8 167.7 100.4C543.2 241.8 606.6 201 679.7 201c104 0 188.3 82.6 188.3 184.5 0 201.2-356 429.3-356 429.3z"
                            p-id="3044"></path>
                    </svg>
                    <span v-if="!card.is_default">设为默认</span>
                    <span v-if="card.is_default">取消默认</span>
                </div>

                <div class="card__clock-info">
                    <svg class="card__clock" viewbox="0 0 24 24">
                        <path
                            d="M12,20A7,7 0 0,1 5,13A7,7 0 0,1 12,6A7,7 0 0,1 19,13A7,7 0 0,1 12,20M19.03,7.39L20.45,5.97C20,5.46 19.55,5 19.04,4.56L17.62,6C16.07,4.74 14.12,4 12,4A9,9 0 0,0 3,13A9,9 0 0,0 12,22C17,22 21,17.97 21,13C21,10.88 20.26,8.93 19.03,7.39M11,14H13V8H11M15,1H9V3H15V1Z">
                        </path>
                    </svg>
                    <span class="card__time">
                        15 min
                    </span>
                </div>
            </div>
            <div class="card__img" :style="{'background-image':`url('${card.image}')`}">
            </div>
            <a class="card_link" href="#">
                <div class="card__img--hover" :style="{'background-image':`url('${card.image}')`}">
                </div>
            </a>
            <div class="card__info">
                <span class="card__category">
                    Let's go
                </span>
                <h3 class="card__title">
                    {{card.describe}}
                </h3>
                <span class="card__by">
                    by
                    <router-link class="card__author" :to="card.path" title="author">
                        {{card.title}}
                    </router-link>
                </span>
            </div>
        </article>
    </div>
</template>

<script>
export default {
    props: ["card"],
    data() {
        return {
        }
    },
    methods: {
        default_open() {
            console.log("default");
            this.card.is_default = !this.card.is_default;
        }
    }
}

</script>

<style lang='less' scoped>
@Card_H : 150px;
@Card_W_Min : 220px;
@Card_H_Min : 250px;
/* .card-root {
    max-width: 30vw;
    min-width: 275px;
    height: 30vh;
    box-shadow: 0 15px 25px rgb(0, 0, 0, 0.1);
    backdrop-filter: blur(30px);
} */


* {
    box-sizing: border-box;
}

.card {
    @Size:15vw;
    margin-right: 25px;
    transition: all .4s cubic-bezier(0.175, 0.885, 0, 1);
    min-width: @Card_W_Min;
    width: @Size;
    height: @Size*1.5;
    min-height: @Card_H_Min * 1.5;
    position: relative;
    border-radius: 12px;
    overflow: hidden;
    box-shadow: 0px 13px 10px -7px rgba(0, 0, 0, 0.1);
}

.default_open {
    margin-right: 57px;
    margin-top: 2px;
    display: flex;
    &>*{
        display: block;
    }
    &>span{
        margin-left: 3px;
        font-size: 12px;
    }
}

.card__like {
    width: 24px;
    height: 24px;
}

.card__like:hover {
    cursor: pointer;
    fill: #ffa39e;
}

.is_default_card {
    fill: #ffa39e;
}

.card__clock {
    width: 23px;
    height: 23px;
    vertical-align: middle;
    fill: #AD7D52;
}

.card__time {
    font-size: 12px;
    color: #AD7D52;
    vertical-align: middle;
    margin-left: 5px;
}

.card__clock-info {
    float: right;
}

.card__img {
    visibility: hidden;
    background-size: cover;
    background-position: center;
    background-repeat: no-repeat;
    width: 100%;
    height: 235px;
    border-top-left-radius: 12px;
    border-top-right-radius: 12px;
}

.card__info-hover {
    position: absolute;
    padding: 16px;
    width: 100%;
    opacity: 0;
    display: flex;
    top: 0;
    z-index: 999;
}

.card__img--hover {
    transition: 0.2s all ease-out;
    background-size: cover;
    background-position: center;
    background-repeat: no-repeat;
    width: 100%;
    position: absolute;
    height: 235px;
    border-top-left-radius: 12px;
    border-top-right-radius: 12px;
    top: 0;
}

.card:hover {
    box-shadow: 0px 30px 18px -8px rgba(0, 0, 0, 0.1);
    transform: scale(1.10, 1.10);
}

.card__info {
    z-index: 2;
    border-bottom-left-radius: 12px;
    border-bottom-right-radius: 12px;
    padding: 16px 24px 24px 24px;
}

.card__category {
    font-family: 'Raleway', sans-serif;
    text-transform: uppercase;
    font-size: 13px;
    letter-spacing: 2px;
    font-weight: 500;
    color: #868686;
}

.card__title {
    margin-top: 5px;
    margin-bottom: 10px;
    font-family: 'Roboto Slab', serif;
}

.card__by {
    font-size: 12px;
    font-family: 'Raleway', sans-serif;
    font-weight: 500;
}

.card__author {
    font-weight: 600;
    text-decoration: none;
    color: #AD7D52;
}

.card:hover .card__img--hover {
    height: 100%;
    opacity: 0.3;
}

.card:hover .card__info {
    background-color: transparent;
    position: relative;
}

.card:hover .card__info-hover {
    opacity: 1;
}
</style>
