package main

import (
	"bufio"
	"encoding/csv"
	"github.com/jinzhu/gorm"
	"io"
	"os"
)

type Expense struct {
	gorm.Model
	Date        string
	Description string
	Debit       string
	Credit      string
	Balance     string
}

type User struct {
	gorm.Model
	FirstName string
	LastName  string
	Username  string
	Password  string
	Email     string
}

type Category struct {
	Description string
}

func ImportCsv(filename string) {
	file, err := os.Open(filename)
	CheckError(err)
	r1, err := bufio.NewReader(file).ReadSlice('\n')
	CheckError(err)
	_, err = file.Seek(int64(len(r1)), io.SeekStart)
	CheckError(err)
	reader, _ := csv.NewReader(file).ReadAll()

	for i := range reader {
		GetDB().Create(&Expense{
			Date:        reader[i][1],
			Description: reader[i][2],
			Debit:       reader[i][3],
			Credit:      reader[i][4],
			Balance:     reader[i][5],
		})
	}
}
