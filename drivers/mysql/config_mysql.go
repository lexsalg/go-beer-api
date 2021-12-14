package mysql

import (
	"beerapi/drivers/mysql/beers"
	"beerapi/drivers/mysql/users"
	"fmt"
	"log"

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

func (configDB *ConfigDB) InitDB() *gorm.DB {
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

func MigrateDB(db *gorm.DB) {
	_ = db.AutoMigrate(&users.Users{}, &beers.Beers{})
}
