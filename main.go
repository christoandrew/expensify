package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const LogFile = "nginx/log/expensify.log"

func main() {
	defer CloseDB()
	Boostrap()
	router := gin.Default()
	router.Use(LoggerMiddleWare())
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))
	api := router.Group("/api")
	api.Use(LoggerMiddleWare())
	api.Use(AuthMiddleWare())
	RegisterExpensesEndpoint(api.Group("/expenses"))

	_ = router.Run(":9087")

}
