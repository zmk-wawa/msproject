package config

import (
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"log"
	"os"
	"project-common/logs"
)

var C = InitConfig()

// 路由的Server
type ServerConfig struct {
	Name string
	Addr string
}

type Config struct {
	viper *viper.Viper
	SC    *ServerConfig
}

func (c *Config) ReadServerConfig() {
	sc := &ServerConfig{}
	sc.Name = c.viper.GetString("server.name")
	sc.Addr = c.viper.GetString("server.addr")
	c.SC = sc
}

func (c *Config) ReadRedisConfig() *redis.Options {
	//rd := redis.Options{}
	//rd.Addr = c.viper.GetString("redis.host") + ":" + c.viper.GetString("redis.port")
	//rd.Password = c.viper.GetString("redis.password")
	//rd.DB = c.viper.GetInt("redis.db")
	//
	//return &rd
	return &redis.Options{
		Addr:     c.viper.GetString("redis.host") + ":" + c.viper.GetString("redis.port"),
		Password: c.viper.GetString("redis.password"),
		DB:       c.viper.GetInt("redis.db"),
	}
}

func (c *Config) InitZapLog() {
	lc := &logs.LogConfig{
		DebugFileName: c.viper.GetString("zap.debugFileName"),
		InfoFileName:  c.viper.GetString("zap.infoFileName"),
		WarnFileName:  c.viper.GetString("zap.warnFileName"),
		MaxSize:       c.viper.GetInt("maxSize"),
		MaxAge:        c.viper.GetInt("maxAge"),
		MaxBackups:    c.viper.GetInt("maxBackups"),
	}
	err := logs.InitLogger(lc)
	if err != nil {
		log.Fatalln(err)
	}
}

func InitConfig() *Config {
	// 1.找到对应的配置文件： 格式+路径+read
	v := viper.New()
	conf := &Config{viper: v}

	workDir, _ := os.Getwd()
	// 名称,格式,路径
	conf.viper.SetConfigName("config")
	conf.viper.SetConfigType("yaml")
	// 可以加多个路径，根据优先级选择
	conf.viper.AddConfigPath(workDir + "/config")
	err := conf.viper.ReadInConfig()
	if err != nil {
		log.Fatalln(err)
	}

	conf.ReadServerConfig()
	conf.InitZapLog()
	return conf
}
