package global

import (
	"github.com/minio/minio-go/v7"
	"github.com/uidea/artwork-backend/pkg/logger"
	"github.com/uidea/artwork-backend/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	Logger          *logger.Logger
	JWTSetting      *setting.JWTSettingS
	MinioSetting    *setting.MinioSettingS
	MinioClient     *minio.Client
)
