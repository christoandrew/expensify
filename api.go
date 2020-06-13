package main

import (
	exporters "expensify/prometheus"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetExpenses(context *gin.Context) {
	db := GetDB()
	var expenses []Expense
	db.Find(&expenses)
	exporters.ExpensesRequestCounter.Inc()
	context.JSON(http.StatusOK, expenses)
}
