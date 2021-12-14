package main

import (
	"beerapi/app/config"
	"beerapi/app/middleware/auth"
	handlerBeers "beerapi/app/presenter/beers"
	handlerHealth "beerapi/app/presenter/health"
	handlerUsers "beerapi/app/presenter/users"
	"beerapi/app/routes"
	serviceBeers "beerapi/bussiness/beers"
	serviceHealth "beerapi/bussiness/health"
	serviceUsers "beerapi/bussiness/users"
	mysqlDrivers "beerapi/drivers/mysql"
	repositoryBeers "beerapi/drivers/mysql/beers"
	repositoryUsers "beerapi/drivers/mysql/users"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
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
	keyCurrency := config.CurrencyApi{
		Url:    viper.GetString(`currency.url`),
		ApiKey: viper.GetString(`currency.key`),
	}

	//initial DB
	db := configDB.InitDB()
	//Migrate DB
	mysqlDrivers.MigrateDB(db)
	//Init Fiber Framework
	app := fiber.New()
	app.Use(logger.New())
	app.Use(recover.New())

	//factory of domain

	//health
	healthService := serviceHealth.NewService()
	healthHandler := handlerHealth.NewHandler(healthService)

	//user
	userRepository := repositoryUsers.NewRepositoryMySQL(db)
	userService := serviceUsers.NewService(userRepository, &configJWT)
	userHandler := handlerUsers.NewHandler(userService)

	//beer
	beerRepository := repositoryBeers.NewRepositoryMySQL(db)
	beerService := serviceBeers.NewService(beerRepository, &keyCurrency)
	beerHandler := handlerBeers.NewHandler(beerService)

	//routes handler
	routesInit := routes.HandlerList{
		HealthHandler: *healthHandler,
		UserHandler:   *userHandler,
		BeerHandler:   *beerHandler,
	}
	routesInit.Routes(app)
	log.Fatal(app.Listen(viper.GetString("server.address")))
}
