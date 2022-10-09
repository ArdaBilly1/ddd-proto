package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/viper"
)

type Service struct {
	Addr  string
	Host  string
	Port  string
	Name  string
	Debug bool
}

func init() {
	LoadEnvar()

	mysql := new(MysqlConfig)
	mysql.SetConfigMysql().ConnectMysql()

	// RunMigrations(MysqlDB)
}

func LoadEnvar() {
	viper.SetEnvPrefix("ardamock")
	viper.BindEnv("env")
	viper.BindEnv("app_path")

	dir, _ := os.Getwd()
	AppPath := dir

	cfg := "config"

	viper.SetConfigName(cfg)
	viper.SetConfigType("json")
	viper.AddConfigPath(AppPath)

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("can't reading format file: %s", err))
	}
}

func (service *Service) SetConfig() *Service {
	service.Host = viper.GetString("app_host")
	service.Port = viper.GetString("app_port")
	service.Name = viper.GetString("app_name")
	service.Debug, _ = strconv.ParseBool(os.Getenv("APP_DEBUG"))
	service.Addr = fmt.Sprintf("%v:%v", service.Host, service.Port)

	return service
}
