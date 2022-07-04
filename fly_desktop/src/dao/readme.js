// 读取指定目标的文件
const { readFileSync } = require('fs');
// let base_path = "D:\\3.Node.js\\worspace-javaScript\\flydev\\fly_desktop\\static\\config.yaml";
let base_path = "D:\\3.Node.js\\worspace-javaScript\\flydev\\fly_desktop\\src\\dao\\target.txt";

function readData() {
    // 读取指定目标的文件
    return readFileSync(base_path, 'utf8');
}

export default readData;