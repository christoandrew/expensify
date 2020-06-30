package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// type H map[string]interface{}

func AuthMiddleWare() gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}

func LoggerMiddleWare() gin.HandlerFunc{
	return func(context *gin.Context) {

		GetLogger().WithFields(logrus.Fields{
			"endpoint": context.FullPath(),
		}).Info("expenses_endpoint")
	}
}
