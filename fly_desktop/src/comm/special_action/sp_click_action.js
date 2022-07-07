import $ from 'jquery'

// 设置列点击效果
export function columnClick(clickArgs) {
    clickArgs = {
        my: clickArgs.my || '.is-build-data',
        you: clickArgs.you || '.is-build',
        afterEditingFunction: clickArgs.afterEditingFunction || function () {
            // 你是猪
        },
    }
    // 带有默认值的传参
    $(clickArgs.my).dblclick(function () {
        this.style.display = 'none'
        let that = $(this).next()
        let value = this.innerText
        that.children('input').val(value)
        that.css('display', 'block')
    })

    // 设置按下回车键恢复常态
    $(clickArgs.you).keydown(function (event) {
        if (event.keyCode === 13) {
            recovery(clickArgs, this)
        }
    })

    // 设置双击输入框恢复常态
    $(clickArgs.you).dblclick(function () {
        recovery(clickArgs, this)
    })
}

function recovery(clickArgs, el) {
    let that = $(el).prev()
    that.css('display', 'block')
    $(el).css('display', 'none')
    // 调用编辑后函数
    clickArgs.afterEditingFunction(el, that)
}