package global

import (
	"blog-server/pkg/logger"
	"blog-server/pkg/setting"
)

var (
	ServerSetting    *setting.ServerSettings
	AppSettings      *setting.AppSettings
	DataBaseSettings *setting.DataBaseSettings
	Logger           *logger.Logger
)
