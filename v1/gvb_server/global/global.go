package global

import (
	"gvb_server/config"

	"github.com/cc14514/go-geoip2"
	"github.com/olivere/elastic/v7"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	// 配置文件全局变量
	Config *config.Config
	//数据库全局变量
	DB     *gorm.DB
	//日志全局变量
	Log *logrus.Logger
	//全局mysql日志
	MysqlLog logger.Interface
	//全局redis
	Redis *redis.Client
	//全局es
	ESClient *elastic.Client
	//通过ip读取城市
	AddrDB *geoip2.DBReader
)

