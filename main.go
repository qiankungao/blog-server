package main

import (
	"blog-server/global"
	"blog-server/internal/model"
	"blog-server/internal/routers"
	"blog-server/pkg/setting"
	"github.com/gin-gonic/gin"
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
}

func main() {
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
