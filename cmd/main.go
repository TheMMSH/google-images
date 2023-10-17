package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/spf13/viper"
	"google-images/crypt"
	"google-images/googleapis"
	"google-images/img"
	"google-images/pg"
	"google-images/server"
	"google-images/service"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func main() {
	setUpViper()
	cr := crypt.New()
	db := pg.New(getGormDb())

	router := gin.New()
	controller := server.Server{
		DownloaderService: service.NewIDownloaderService(
			getGoogleApi(),
			img.New(viper.GetUint("img.width"), viper.GetUint("img.height")),
			cr,
			db,
		),
		PresenterService: service.NewIPresenterService(cr, db),
	}

	controller.SetRoutes(router)
	router.Run(":" + getEnv("port", "9000"))
}

func getGoogleApi() googleapis.IGoogleApiService {
	apiKey := viper.GetString("google.apiKey")
	sEID := viper.GetString("google.searchEngineId")
	if apiKey == "" || sEID == "" {
		log.Fatalf("please ensure that apiKey and searchEngineId exists in config file (use conf/example-config.yaml as an example")
	}
	return googleapis.New(apiKey, sEID)
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func setUpViper() {
	viper.SetConfigName(getEnv("CONFIG_NAME", "config"))
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./conf")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Fatal error config file: %w \n", err)
	}
}

func getGormDb() *gorm.DB {
	db, err := gorm.Open(postgres.Open(viper.GetString("postgresSource")))
	if err != nil {
		log.Fatalf("%v failed to initialize gorm DB", err)
	}
	return db
}
