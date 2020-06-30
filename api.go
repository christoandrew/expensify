package main

import (
	exporters "expensify/prometheus"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func GetExpenses(context *gin.Context) {
	logger, _ := NewLogger()
	defer logger.Sync()
	db := GetDB()
	var expenses []Expense
	db.Find(&expenses)
	logger.Info(
		"Loaded Expenses",
		zap.String("url", context.FullPath()),
		zap.String("clientip", context.ClientIP()),
		zap.String("method", context.Request.Method),
		zap.String("useragent", context.Request.UserAgent()),
	)
	exporters.ExpensesRequestCounter.Inc()
	context.JSON(http.StatusOK, expenses)
}
