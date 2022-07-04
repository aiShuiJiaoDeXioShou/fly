/**
* @作者: 林河
* @date: 2022/6/27
* @describe:
**/
package service

import (
	"flye/service/config"
	"flye/service/netreptile/qidiansort"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os/exec"
	"runtime"
	"strings"
	"syscall"
)

func ServiceApi(e *gin.Engine) {
	rg := e.Group("/api")
	qidiansort.NewQiDianSort(rg)
	config.NewConfigService(rg)
}

func StartService() {
	fmt.Println("开始启动服务器...")
	e := gin.Default()
	e.Static("/dist", "./dist")

	// 启动服务器的所有Api
	ServiceApi(e)

	e.GET("/", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusMovedPermanently, "/dist/index.html")
	})

	// 跳转到首页
	//JumpIndex()

	err := e.Run(":41")
	if err != nil {
		panic("滴滴，端口被占领，请联系程序员更换端口号！")
	}
}

// 跳转到首页
func JumpIndex() {
	// 如果服务器启动成功则打开首页
	// 无GUI调用
	if strings.Contains("windows", runtime.GOOS) {
		cmd := exec.Command(`cmd`, `/c`, `start`, `http://localhost:41/`)
		cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
		err := cmd.Start()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("如果打开首页失败，看是不是浏览器版本太低了，请升级浏览器！")
		fmt.Println("将这段链接直接粘贴到浏览器导航栏当中，即可跳转到首页！")
		fmt.Println("http://localhost:41/")
	} else {
		cmd := exec.Command(`xdg-open`, `http://localhost:41/`).Start()
		// 如果启动linux出错那么肯定是macos
		if cmd != nil {
			exec.Command(`open`, `http://localhost:41/`).Start()
		}
	}

}
