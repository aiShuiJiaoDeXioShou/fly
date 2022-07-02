<!--  -->
<template>
    <div>
        <div id="div_raindrops">
            <div class="my_raindrops_item">
            </div>
        </div>
    </div>
</template>

<script>
export default {
    data() {
        return {
        }
    },
    methods: {
        raindrop() {
            const div_raindrops = document.querySelector('#div_raindrops');
            for (let i = 0; i < div_raindrops.clientWidth / 5; i++) {
                let raindrops_item = document.createElement('div');
                raindrops_item.className = 'my_raindrops_item';
                raindrops_item.style.backgroundImage = 'linear-gradient(transparent,rgb(' + Math.random() * 255 + ',' + Math.random() * 255 + ',' + Math.random() * 255 + '))';
                raindrops_item.style.animationDelay = Math.random() * 10 + 's';
                raindrops_item.style.left = i * 5 + 'px';
                raindrops_item.style.animationTimingFunction = 'cubic-bezier(' + (Math.random() * 0.4 + 0.6) + ',' + (Math.random() * 0.2) + ',' + (Math.random() * 0.2 + 0.8) + ',' + (Math.random() * 0.2 + 0.4) + ')';
                div_raindrops.appendChild(raindrops_item);
            }
        }
    },
    mounted() {
        this.raindrop();
    }
}

</script>

<style lang='css' scoped>
* {
    margin: 0px;
    padding: 0px;
}

#div_raindrops {
    position: fixed;
    width: 100vw;
    height: 100vh;
    overflow: hidden;
}

@keyframes my_raindrops {
    from {
        transform: translateY(-200px);
    }

    to {
        transform: translateY(100vh);
    }
}

.my_raindrops_item {
    position: absolute;
    width: 5px;
    height: 30px;
    border-radius: 0% 0% 5px 5px;
    transform: translateY(-200px);
    animation-name: my_raindrops;
    animation-duration: 5s;
    animation-iteration-count: infinite;
}


</style>
