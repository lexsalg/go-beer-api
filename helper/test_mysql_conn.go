package helper

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ConfigDB struct {
	DBUsername string
	DBPassword string
	DBHost     string
	DBPort     string
	DBDatabase string
}

func InitTestingDB() *gorm.DB {
	configDB := readEnv()
	destination := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		configDB.DBUsername,
		configDB.DBPassword,
		configDB.DBHost,
		configDB.DBPort,
		configDB.DBDatabase,
	)

	db, err := gorm.Open(mysql.Open(destination), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func readEnv() ConfigDB {
	viper.SetConfigName("test-config")
	viper.AddConfigPath("")
	viper.AutomaticEnv()
	viper.SetConfigType("json")
	err := viper.ReadInConfig()
	if err != nil {
		return ConfigDB{
			DBUsername: "test",
			DBPassword: "test",
			DBHost:     "localhost",
			DBPort:     "3306",
			DBDatabase: "beerapp_test",
		}
	}

	return ConfigDB{
		DBUsername: viper.GetString(`database_testing.user`),
		DBPassword: viper.GetString(`database_testing.pass`),
		DBHost:     viper.GetString(`database_testing.host`),
		DBPort:     viper.GetString(`database_testing.port`),
		DBDatabase: viper.GetString(`database_testing.name`),
	}
}
