package database

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var dsn string

func init() {
	viper.SetConfigName("application")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}

	dbHost := viper.GetString("database.host")
	dbPort := viper.GetString("database.port")
	dbUsername := viper.GetString("database.username")
	dbPassword := viper.GetString("database.password")
	dbName := viper.GetString("database.dbname")

	//    dsn := "username:password@tcp(hostname:port)/database?charset=utf8mb4&parseTime=True&loc=Local"
	dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUsername, dbPassword, dbHost, dbPort, dbName)
	log.Info("database connection url ", dsn)
}

func connectionDB() error {
	var err error
	config := gorm.Config{
		PrepareStmt: true,
	}
	db, err = gorm.Open(mysql.Open(dsn), &config)
	sqlDB, err := db.DB()
	if err != nil {
		return nil
	}
	sqlDB.SetMaxOpenConns(20)
	sqlDB.SetMaxIdleConns(5)

	return nil
}

func CloseDB() error {
	dbSQL, err := db.DB()
	if err != nil {
		return err
	}
	err = dbSQL.Close()
	if err != nil {
		return err
	}
	return nil
}

func GetDB() *gorm.DB {
	err := connectionDB()
	if err != nil {
		log.Fatalf("connect database failed %s", dsn)
	}

	return db
}
