package main

import (
	"fmt"
	"go.uber.org/zap"
	"log"
	"time"
)

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func formatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d% 02d/%02d", year, month, day)
}

func NewLogger() (*zap.Logger, error) {
	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{
		LogFile,
	}
	return cfg.Build()
}
