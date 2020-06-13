package main

import (
	"expensify/db"
	exporters "expensify/prometheus"
	"github.com/prometheus/client_golang/prometheus"
)

func Boostrap() {
	InitDB()
	models := []interface{}{Expense{}, User{}, Category{}}
	registry := db.NewModelRegistry()
	registry.RegisterAll(models...)
	Migrate(registry)
	prometheus.MustRegister(exporters.ExpensesRequestCounter)
}
