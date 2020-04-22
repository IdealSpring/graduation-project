package config

import (
	"github.com/spf13/viper"
	"log"
)

// Application 配置
type Application struct {
	Env           string
	Name          string
	Host          string
	Port          string
	LogPath       string
	IsInit        bool
}
// Database 配置
type Database struct {
	Database string
	Dbtype   string
	Host     string
	Port     int
	Username string
	Password string
}

var ApplicationConfig = new(Application)
var DatabaseConfig = new(Database)

// 初始化 Application 配置
func initApplicationConfig(cfg *viper.Viper) *Application {
	return &Application{
		Env:           cfg.GetString("env"),
		Name:          cfg.GetString("name"),
		Host:          cfg.GetString("host"),
		Port:          cfg.GetString("port"),
		LogPath:       cfg.GetString("logPath"),
		IsInit:        cfg.GetBool("isInit"),
	}
}

// 初始化数据库配置
func initDatabaseConfig(cfg *viper.Viper) *Database {
	return &Database{
		Database: cfg.GetString("database"),
		Dbtype:   cfg.GetString("dbtype"),
		Host:     cfg.GetString("host"),
		Port:     cfg.GetInt("port"),
		Username: cfg.GetString("username"),
		Password: cfg.GetString("password"),
	}
}

// 设置配置文件中 isinit 为 true
func SetApplicationDatabaseIsInit() {
	viper.AddConfigPath("./config")
	viper.Set("settings.application.isInit", true)
	viper.WriteConfig()
}

// 加载配置信息
func init() {
	viper.SetConfigName("settings") // 设置配置文件名(不需要写后缀名)
	viper.AddConfigPath("./config") // 在工作目录中查找配置 - [可选配置]
	err := viper.ReadInConfig()     // 加载和读取配置文件
	if err != nil {
		log.Println("load and read config err:", err)
	}

	// 从配置文件中初始化数据库连接参数
	cfgDatabase := viper.Sub("settings.database")
	if cfgDatabase == nil {
		panic("config not found settings.database")
	}
	DatabaseConfig = initDatabaseConfig(cfgDatabase)

	// 从配置文件中初始化gin服务参数
	cfgApplication := viper.Sub("settings.application")
	if cfgApplication == nil {
		panic("config not found settings.application")
	}
	ApplicationConfig = initApplicationConfig(cfgApplication)
}


