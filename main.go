package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/uidea/artwork-backend/global"
	"github.com/uidea/artwork-backend/internal/model"
	"github.com/uidea/artwork-backend/internal/routers"
	"github.com/uidea/artwork-backend/pkg/app"
	"github.com/uidea/artwork-backend/pkg/logger"
	"github.com/uidea/artwork-backend/pkg/setting"
	"gopkg.in/natefinch/lumberjack.v2"
)

type MinioConfig struct {
	Endpoint        string
	AccessKeyID     string
	SecretAccessKey string
}

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

	err = setupMinio()
	if err != nil {
		log.Fatalf("init.setupMinio err: %v", err)
	}

}

func setupMinio() error {
	appEnv := os.Getenv("APP_ENV")
	var conf *MinioConfig

	if appEnv == "production" {
		conf = &MinioConfig{
			Endpoint:        os.Getenv("MINIO_ENDPOINT"),
			AccessKeyID:     os.Getenv("MINIO_ACCESSKEYID"),
			SecretAccessKey: os.Getenv("MINIO_SECRETACCESSKEY"),
		}
	} else {
		conf = &MinioConfig{
			Endpoint:        global.MinioSetting.Endpoint,
			AccessKeyID:     global.MinioSetting.AccessKeyID,
			SecretAccessKey: global.MinioSetting.SecretAccessKey,
		}
	}

	minioClient, err := minio.New(conf.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(conf.AccessKeyID, conf.SecretAccessKey, ""),
		Secure: false,
	})

	if err != nil {
		log.Fatalln(err)
	}
	global.MinioClient = minioClient
	err = app.CheckBuckets(minioClient)
	if err != nil {
		log.Fatalln(err)
	}
	return nil
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
	err = setting.ReadConfigField("Minio", &global.MinioSetting)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	router := routers.NewRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = global.ServerSetting.HttpPort
	}

	s := &http.Server{
		Addr:           ":" + port,
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
