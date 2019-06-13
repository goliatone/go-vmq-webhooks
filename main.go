package main

import (
	"fmt"
	"os"

	routes "github.com/goliatone/go-vmq-webhooks/router"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var port string

func main() {
	viper.SetConfigName("global")
	viper.AddConfigPath("./config")

	err := viper.ReadInConfig()

	if err != nil {
		fmt.Println("Fatal error config file: %s \n", err)
		panic("Fatal error")
	}

	port = viper.GetString("port")

	p := os.Getenv("PORT")
	if len(p) > 0 {
		port = p
	}

	router := gin.Default()

	v1 := router.Group("/")
	routes.WebHooks(v1.Group("/"))

	fmt.Printf("Running sever on port: %s\n", port)

	err = router.Run("0.0.0.0" + port)
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
