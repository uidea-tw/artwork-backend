package main

import (
	"log"
	"net/http"
	"time"

	"github.com/uidea/artwork-backend/global"
	"github.com/uidea/artwork-backend/internal/model"
	"github.com/uidea/artwork-backend/internal/routers"
	"github.com/uidea/artwork-backend/pkg/logger"
	"github.com/uidea/artwork-backend/pkg/setting"
	"gopkg.in/natefinch/lumberjack.v2"
)

func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}

	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}

	err = setupLogger()
	if err != nil {
		log.Fatalf("init.setupLogger err: %v", err)
	}
}

func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)

	if err != nil {
		return err
	}
	return nil
}

func setupSetting() error {
	setting, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = setting.ReadConfigField("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = setting.ReadConfigField("App", &global.AppSetting)
	if err != nil {
		return err
	}
	err = setting.ReadConfigField("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}
	err = setting.ReadConfigField("JWT", &global.JWTSetting)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	router := routers.NewRouter()

	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    time.Second * global.ServerSetting.ReadTimeout,
		WriteTimeout:   time.Second * global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}

func setupLogger() error {
	fileName := global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  fileName,
		MaxSize:   500,
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)

	return nil
}
