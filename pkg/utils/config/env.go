package config

import (
	"go-mitra-akosta/pkg/utils/logger"

	"github.com/spf13/viper"
)

type Config struct {
	Server *serverConfig
	Mysql  *dbMysqlServer
	Jwt    *jwt
}

type serverConfig struct {
	AppName  string
	AppStage string
	AppKey   string
	Port     string
	Version  string
}

type jwt struct {
	Secret string
}

type dbMysqlServer struct {
	MysqlDriver   string
	MysqlName     string
	MysqlUserName string
	MysqlPassword string
	MysqlHost     string
	MysqlPort     string
}

type Environment struct {
	Log      logger.Logging
	FileName string
	PathFile string
	TypeFile string
}

func NewEnvironment(log *logger.Logging, fileName string, pathFile string, typeFile string) *Environment {
	return &Environment{
		Log:      *log,
		FileName: fileName,
		PathFile: pathFile,
		TypeFile: typeFile,
	}
}

func (e *Environment) LoadConfig() *viper.Viper {
	v := viper.New()
	v.SetConfigName(e.FileName)
	v.AddConfigPath(e.PathFile)

	v.SetConfigType(e.TypeFile)
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			e.Log.ErrorLog("config file not found")
			return nil
		}
		e.Log.ErrorLog("failed read config file:", e.FileName)
		return nil
	}

	return v
}

func (e *Environment) ParseConfig(v *viper.Viper) *Config {
	var c Config

	error := v.Unmarshal(&c)

	if error != nil {
		e.Log.ErrorLog("error ParseConfig: ", error)
		return nil
	}

	return &c
}
