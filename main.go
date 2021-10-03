package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/razanlrahardjo/hacktiv8/app/routes"
	"github.com/spf13/viper"
	"log"
	"os"
	"strings"
)

// @title Master API
// @version 1.0.0
// @description API Documentation
// @termsOfService http://localhost:9000
// @contact.name Developer
// @contact.email razanlrahardjo@gmail.com
// @host http://localhost:9000
// @schemes http
// @BasePath /api/v1/master
func main() {
	app := fiber.New(fiber.Config{
		Prefork: viper.GetString("PREFORK") == "true",
	})

	routes.Handle(app)
	log.Fatal(app.Listen(":" + viper.GetString("PORT")))
}

func init() {
	// load local env
	localEnv := viper.New()
	localEnv.SetConfigType("dotenv")
	localEnv.SetConfigFile(".env")
	if err := localEnv.ReadInConfig(); nil == err {
		localEnvKeys := localEnv.AllKeys()
		for i := range localEnvKeys {
			viper.Set(localEnvKeys[i], localEnv.Get(localEnvKeys[i]))
		}
	}
	keys := viper.AllKeys()
	for i := range keys {
		stringValue := viper.GetString(keys[i])
		if stringValue == "" {
			if value := viper.Get(keys[i]); nil != value {
				stringValue = fmt.Sprintf("%v", value)
			}
		}
		os.Setenv(strings.ToUpper(keys[i]), stringValue)
	}
}
