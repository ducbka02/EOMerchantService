// "host": "171.244.142.143",
// "port": "3306",
// "user": "username",
// "pass": "Ems@2019",
package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/spf13/viper"

	_httpDelivery "merchant-service/merchant/delivery/http"
	_merchantRepo "merchant-service/merchant/repository"
	_merchantUsecase "merchant-service/merchant/usecase"
	middleware "merchant-service/middleware"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		fmt.Println("Service run on DEBUG mode")
	}
}

func main() {
	dbHost := viper.GetString(`database.host`)
	dbPort := viper.GetString(`database.port`)
	dbUser := viper.GetString(`database.user`)
	dbPass := viper.GetString(`database.pass`)
	dbName := viper.GetString(`database.name`)
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	dsn := fmt.Sprintf("%s", connection)
	dbConn, err := sql.Open(`mysql`, dsn)
	if err != nil && viper.GetBool("debug") {
		fmt.Println(err)
	}
	err = dbConn.Ping()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	fmt.Println("Connected...")

	defer func() {
		err := dbConn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	e := echo.New()
	middL := middleware.InitMiddleware()
	e.Use(middL.CORS)
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hi! I am a merchant-service")
	})

	// Routes
	articleRepo := _merchantRepo.NewMysqlMerchantRepository(dbConn)
	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second
	articleUsecase := _merchantUsecase.NewMerchantUsecase(articleRepo, timeoutContext)
	_httpDelivery.NewMerchantHandler(e, articleUsecase)

	e.Logger.Fatal(e.Start(":1080"))
}
