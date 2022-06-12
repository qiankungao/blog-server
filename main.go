package main

import (
	"blog-server/global"
	"blog-server/internal/model"
	"blog-server/internal/routers"
	"blog-server/pkg/logger"
	"blog-server/pkg/setting"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/natefinch/lumberjack"
	"log"
	"net/http"
	"time"
)

func init() {
	if err := setupSetting(); err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}
	if err := setupDBEngine(); err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}
	if err := setupLogger(); err != nil {
		log.Fatalf("init.setupLogger err: %v", err)
	}
}

type funcType func(c int) int

//函数作为返回值和作为返回值的这个函数能够调用之前函数的参数，就是Go语言闭包的精髓

func main() {
	for _, c := range countBy() {
		fmt.Println(c())
	}
	fmt.Println("启动成........")
}

//每次调用闭包函数所处的环境都是相互独立的,
func countBy() []func() int {
	var arr []func() int
	for i := 1; i <= 3; i++ {
		func(n int) {
			arr = append(arr, func() int {
				return n
			})
		}(i)

	}
	return arr
}

//@Title 博客系统
//@version 1.0
//@description Go 变成之旅
func Run() {
	gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewRouter()
	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeOut,
		WriteTimeout:   global.ServerSetting.WriteTimeOut,
		MaxHeaderBytes: 1 << 20,
	}
	_ = s.ListenAndServe()
}

func setupSetting() error {
	setting, err := setting.NewSetting()
	if err != nil {
		return err
	}
	if err = setting.ReadSection("Server", &global.ServerSetting); err != nil {
		return err
	}
	if err = setting.ReadSection("App", &global.AppSettings); err != nil {
		return err
	}
	if err = setting.ReadSection("DataBase", &global.DataBaseSettings); err != nil {
		return err
	}
	global.ServerSetting.ReadTimeOut *= time.Second
	global.ServerSetting.WriteTimeOut *= time.Second
	return nil
}

func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DataBaseSettings)
	if err != nil {
		return err
	}
	return nil
}

func setupLogger() error {
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  global.AppSettings.LogSavePath + "/" + global.AppSettings.LogFileName + global.AppSettings.LogFileExt,
		MaxSize:   600,
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)
	return nil
}
