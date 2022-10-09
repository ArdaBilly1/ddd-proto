package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type MysqlConfig struct {
	Host         string
	Port         string
	Username     string
	Password     string
	DatabaseName string

	ConnStatement string
}

func (MysqlConfig *MysqlConfig) GetStatement() *MysqlConfig {
	MysqlConfig.ConnStatement = "%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local"
	return MysqlConfig
}

func (MysqlConfig *MysqlConfig) SetConfigMysql() *MysqlConfig {
	db := viper.Sub("database.mysql")

	MysqlConfig.Host = db.GetString("host")
	MysqlConfig.Port = db.GetString("port")
	MysqlConfig.Username = db.GetString("username")
	MysqlConfig.Password = db.GetString("password")
	MysqlConfig.DatabaseName = db.GetString("database")

	return MysqlConfig
}

var MysqlDB *gorm.DB

func (MysqlConfig *MysqlConfig) ConnectMysql() *gorm.DB {
	var driver gorm.Dialector
	var isDebug = logger.Error

	dsn := fmt.Sprintf(
		MysqlConfig.GetStatement().ConnStatement,
		MysqlConfig.GetStatement().Username,
		MysqlConfig.GetStatement().Password,
		MysqlConfig.GetStatement().Host,
		MysqlConfig.GetStatement().Port,
		MysqlConfig.GetStatement().DatabaseName,
	)

	driver = mysql.Open(dsn)

	MysqlDB = MysqlConfig.establishConnection(driver, isDebug)

	sqlDB, errDb := MysqlDB.DB()

	if errDb != nil {
		panic(errDb)
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	defer func() {
		sqlDB.Ping()
	}()

	return MysqlDB
}

func (MysqlConfig *MysqlConfig) establishConnection(driver gorm.Dialector, isDebug logger.LogLevel) *gorm.DB {
	dbLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Millisecond,
			LogLevel:      isDebug,
			Colorful:      true,
		},
	)

	connection, errConn := gorm.Open(driver, &gorm.Config{
		PrepareStmt:     true,
		DryRun:          false,
		Logger:          dbLogger,
		CreateBatchSize: 500,
	})

	if errConn != nil {
		panic(errConn)
	}

	return connection
}
