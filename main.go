package main

import (
	"beerapi/app/middleware/auth"
	handlerHealth "beerapi/app/presenter/health"
	handlerUsers "beerapi/app/presenter/users"
	"beerapi/app/routes"
	serviceHealth "beerapi/bussiness/health"
	serviceUsers "beerapi/bussiness/users"
	mysqlDrivers "beerapi/drivers/mysql"
	repositoryUsers "beerapi/drivers/mysql/users"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("test-config")
	viper.AddConfigPath("./app/config/")
	viper.AutomaticEnv()
	viper.SetConfigType("json")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func main() {
	configDB := mysqlDrivers.ConfigDB{
		DBUsername: viper.GetString(`database.user`),
		DBPassword: viper.GetString(`database.pass`),
		DBHost:     viper.GetString(`database.host`),
		DBPort:     viper.GetString(`database.port`),
		DBDatabase: viper.GetString(`database.name`),
	}

	configJWT := auth.ConfigJWT{
		SecretJWT:       viper.GetString(`jwt.secret`),
		ExpiresDuration: viper.GetInt(`jwt.expired`),
	}

	//initial DB
	db := configDB.InitDB()
	//Migrate DB
	mysqlDrivers.MigrateDB(db)
	//Init Fiber Framework
	app := fiber.New()

	//factory of domain

	//health
	healthService := serviceHealth.NewService()
	healthHandler := handlerHealth.NewHandler(healthService)

	//user
	userRepository := repositoryUsers.NewRepositoryMySQL(db)
	userService := serviceUsers.NewService(userRepository, &configJWT)
	userHandler := handlerUsers.NewHandler(userService)

	//routes handler
	routesInit := routes.HandlerList{
		HealthHandler: *healthHandler,
		UserHandler:   *userHandler,
	}
	routesInit.Routes(app)
	log.Fatal(app.Listen(viper.GetString("server.address")))
}
