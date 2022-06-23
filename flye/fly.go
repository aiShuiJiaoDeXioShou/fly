package main

import (
	"fmt"
	"net/http"
	"os/exec"
	"runtime"
	"strings"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	StartService()

}

func StartService() {
	e := gin.Default()
	e.Static("/dist", "./dist")
	rg := e.Group("/api")
	rg.GET("/ok", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			// 这个只有零跟-1，零表示成功，-1表示失败
			"code":    0,
			"message": "success",
			"data":    "你是猪",
			"date":    time.Now(),
		})
	})
	e.GET("/", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusMovedPermanently, "/dist/index.html")
	})
	JumpIndex()
	err := e.Run(":41")
	if err != nil {
		panic("滴滴，端口被占领，请联系程序员更换端口号！")
	}
}

func JumpIndex() {
	// 如果服务器启动成功则打开首页
	// 无GUI调用
	if strings.Contains("windows", runtime.GOOS) {
		cmd := exec.Command(`cmd`, `/c`, `start`, `http://localhost:41/`)
		cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
		cmd.Start()
		fmt.Println(runtime.GOOS)
		fmt.Println(runtime.GOARCH)
	} else {
		cmd := exec.Command(`xdg-open`, `http://localhost:41/`).Start()
		// 如果启动linux出错那么肯定是macos
		if cmd != nil {
			exec.Command(`open`, `http://localhost:41/`).Start()
		}
	}

}
