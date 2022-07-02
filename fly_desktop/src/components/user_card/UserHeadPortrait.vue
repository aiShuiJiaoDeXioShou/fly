<!-- 这个是那个动态的名片部分 -->
<template>
    <div :class="{ envelope: true, envelope_click: is_icon_click }">
        <h1>
            <div class="touxian envelope-icon" @click="is_icon_click = !is_icon_click">
                <el-dropdown @command="dropdownhandle">

                    <!-- 这个是头像部分 -->
                    <el-avatar class="head-portrait" shape="circle" :size="size" :src="squareUrl"></el-avatar><br />

                    <div v-show="!this.is_icon_click" style="color">{{ user.user_name }}</div>
                    <el-dropdown-menu slot="dropdown">
                        <el-dropdown-item v-if="isExistence" :disabled="false" command="login">登入</el-dropdown-item>
                        <el-dropdown-item v-if="isExistence" command="register">注册</el-dropdown-item>
                        <el-dropdown-item divided v-if="!isExistence" command="logout">退出登陆</el-dropdown-item>
                    </el-dropdown-menu>
                </el-dropdown>
            </div>
            <div class="color-bar"></div>
            <div class="color-bar"></div>
            <div class="color-bar"></div>
            <div class="color-bar"></div>
            <div class="color-bar"></div>
            <div class="color-bar"></div>
            <div class="color-horizontal-bar"></div>
            <div id="span1" class="color-horizontal-bar"></div>
            <div class="color-horizontal-bar2"></div>
            <div id="span2" class="color-horizontal-bar2"></div>
        </h1>
        <div :class="{ body: true, body_click: is_icon_click }">
            <div class="text-div">
                <div>
                    <span class="name">名称：</span>
                    <span> : </span>
                    <span class="data">林河</span>
                </div>
                <div>
                    <span class="name">身份：</span>
                    <span> : </span>
                    <span class="data">rp</span>
                </div>
                <div>
                    <span class="name">身份背景：</span>
                    <span> : </span>
                    <span class="data">终末的交响曲是最无语伦比的疯狂。</span>
                </div>
            </div>
            <div class="operation">
                <ul>
                    <li>
                        <a href="#"><i></i></a>
                    </li>
                    <li>
                        <a href="#"><i></i></a>
                    </li>
                    <li>
                        <a href="#"><i></i></a>
                    </li>
                    <li>
                        <a href="#"><i></i></a>
                    </li>
                </ul>
            </div>
            <div class="img-box" v-if="false">
                <!-- <img src="C:\Users\有天道\Pictures\Saved Pictures\heilishi.jpeg" alt="" /> -->
            </div>
        </div>
    </div>
</template>

<script>
import Cookie from 'js-cookie'
export default {
    data() {
        return {
            squareUrl: require("../../assets/icon.png"),
            size: 100,
            user: {
                user_name: "",
            },
            isExistence: true,
            is_icon_click: false,
            dialogTableVisible: false
        };
    },
    methods: {
        login() {
            console.log("登入");
            this.$router.push({ path: "/login", query: { isLogin: "显示" } });
        },
        register() {
            console.log("注册");
            this.$router.push({ path: "/login", query: { isLogin: "隐藏" } });
        },
        logout() {
            console.log("退出登陆");
        },
        dropdownhandle(item) {
            if (item === "login") {
                this.login();
            } else if (item === "register") {
                this.register();
            } else if (item === "logout") {
                this.logout();
            }
        },
    },
    components: {

    },
    created() {
        Cookie.get("user_id")

    }
};
</script>

<style lang="less" scoped>
@Head_Width: 40vw;

.envelope {
    width: @Head_Width;
    height: 40px;
    background-color: rgba(255, 255, 255, 0.2);
    box-shadow: 0 15px 20px rgba(0, 0, 0, 0.2);
    backdrop-filter: blur(10px);
    position: relative;
    margin: 0 auto;
    transition: all 2s;
    padding: 10px;
    z-index: 4;
}

@media (max-width: 580px) {
    .envelope {
        width: 70vw;
    }
}

.body {
    width: @Head_Width;
    height: 0px;
    overflow: hidden;
    position: relative;
    transition: all 0.5s;
}

.img-box {
    position: absolute;
    bottom: 0;
    left: 0;
    z-index: -3;
}

.img-box img {
    width: 100%;
    height: 100%;
    border-radius: 5%;
}

.envelope .color-bar {
    width: 120%;
    transform: translateX(-8%);
    height: 5px;
    margin-bottom: 1px;
    border-radius: 3px;
}

.envelope .color-bar:nth-child(1) {
    background: linear-gradient(135deg, #fce38a, #f38181);
}

.envelope .color-bar:nth-child(2) {
    background: linear-gradient(135deg, #f54ea2, #ff7676);
}

.envelope .color-bar:nth-child(3) {
    background: linear-gradient(135deg, #622774, #c53364);
}

.envelope .color-bar:nth-child(4) {
    background: linear-gradient(135deg, #622774, #c53364);
}

.envelope .color-bar:nth-child(5) {
    background: linear-gradient(135deg, #65799b, #5e2563);
}

.envelope .color-bar:nth-child(6) {
    background: linear-gradient(135deg, #184e68, #57ca85);
}

.envelope-icon:hover {
    cursor: pointer;
}

.touxian {
    position: absolute;
    top: -40px;
    left: 40%;
    z-index: 2;
}

.touxian:hover .head-portrait {
    animation: turnX 1s infinite;
    animation-iteration-count: 1;
    animation-fill-mode: forwards;
}

.color-horizontal-bar {
    height: 60px;
    width: 5px;
    position: absolute;
    top: -10px;
    left: -10px;
    background: #17ead9;
    margin-left: 5px;
    z-index: 3;
    transition: all 0.5s;
}

#span1 {
    background: #fff;
    left: -20px;
}

.color-horizontal-bar2 {
    height: 60px;
    width: 5px;
    position: absolute;
    top: -10px;
    right: -10px;
    background: #17ead9;
    margin-left: 5px;
    z-index: 3;
    transition: all 0.5s;
}

#span2 {
    background: #6078ea;
    top: -10px;
    right: -20px;
}

.envelope:hover .color-horizontal-bar {
    transform: rotate(-45deg);
    transition: all 0.5s;
}

.envelope:hover #span1 {
    transform: rotate(45deg);
    transition: all 0.5s;
}

.envelope:hover .color-horizontal-bar2 {
    transform: rotate(45deg);
    transition: all 0.5s;
}

.envelope:hover #span2 {
    transform: rotate(-45deg);
    transition: all 0.5s;
}

.body_click {
    height: 200px;
    transition: all 1.5s;
}

.envelope_click {
    transition: all 1s;
    height: 250px;
}

.text-div .name {
    font-size: larger;
    color: #fff;
}

.text-div div {
    margin-bottom: 10px;
    padding: 10px;
    border-bottom: 1px solid #fff;
}

.text-div .data {
    font-size: medium;
}
</style>
