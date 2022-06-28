/**
* @作者: 杨腾
* @date: 2022/6/23
* @describe: 定义这个项目的日记管理策略
**/
package mylog

import (
	"bytes"
	"io"
	"log"
	"os"
)

// 错误但是不影响程序的运行
func Wrong(message string, tag string) {
	// 将定义的日记输出到指定的文件中
	writer1 := &bytes.Buffer{}
	writer2 := os.Stdout
	writer3, err := os.OpenFile("log.txt", os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		log.Fatalf("create file log.txt failed: %v", err)
	}
	logger := log.New(io.MultiWriter(writer1, writer2, writer3), "", log.Ldate|log.Ldate|log.Llongfile)
	logger.Printf("来自于< %s >的错误信息:%s\n", tag, message)
}

// 严重的错误，直接退出程序
func Err(message string, tag string) {
	// 将定义的日记输出到指定的文件中
	writer1 := &bytes.Buffer{}
	writer2 := os.Stdout
	writer3, err := os.OpenFile("log.txt", os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		log.Fatalf("create file log.txt failed: %v", err)
	}
	logger := log.New(io.MultiWriter(writer1, writer2, writer3), "", log.Ldate|log.Ldate|log.Llongfile)
	logger.Printf("【错误】来自于< %s >的错误信息:%s\n", tag, message)
	logger.Fatalf("【错误】来自于< %s >的错误信息:%s\n", tag, message)
}
