package main

import (
	"github.com/gin-gonic/gin"
	"github.com/little-go/road-to-go/blog/global"
	"github.com/little-go/road-to-go/blog/internal/routers"
	setting "github.com/little-go/road-to-go/blog/pkg/setting"
	"log"
	"net/http"
	"time"
)

func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting failed: %v", err)
	}
}

func main() {
	gin.SetMode(global.ServerSetting.RunMode)

	r := routers.NewRouter()
	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        r,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	_ = s.ListenAndServe()
}

func setupSetting() error {
	st, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = st.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = st.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	err = st.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}

	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second

	return nil
}
