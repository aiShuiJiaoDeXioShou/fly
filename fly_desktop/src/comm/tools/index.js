import html2canvas from 'html2canvas'
import $ from 'jquery'

// 将普通的html元素转图片
export function htmlToImg(id, target) {
    // 获取想要转化为图片的dom元素
    const imageDom = document.querySelector(id);
    $(imageDom).css('letter-spacing', '10px')

    // 获取image的宽度和高度
    const width = imageDom.offsetWidth
    const height = imageDom.offsetHeight
    html2canvas(imageDom, {
        y: window.pageYOffset,
        windowHeight: document.body.scrollHeight,
        scale: window.devicePixelRatio < 3 ? window.devicePixelRatio : 2,
        useCORS: true,// 允许跨域
        allowTaint: true,
    }).then(canvas => {
        let image = new Image();
        image.src = canvas.toDataURL("image/png", 1.0);
        // 把文字排版还回来
        $(imageDom).css('letter-spacing', '0px')
        image.width = width;
        image.height = height;
        // 下载该图片
        let loadUrl = image.src.replace('image/png', 'image/octet-stream');
        window.location.href = loadUrl;
        document.querySelector(target).appendChild(image);
    });
}