package main

import "github.com/gin-gonic/gin"

func RegisterExpensesEndpoint(group *gin.RouterGroup) {
	group.GET("/", GetExpenses)
}
