package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

/*
	"fmt"
	"os"
	"time"
	"io/ioutil"
	"strings"
	"encoding/csv"
	"strconv"
*/
var watchedPath = os.Getenv("GOPATH") + "/source"

func main() {
	runtime.GOMAXPROCS(4)
	var tracker int64 = 1
	fmt.Println("begin Number", tracker)
	for {
		d, err := os.Open(watchedPath)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			files, _ := d.Readdir(-1)
			for i, file := range files {
				filepath := watchedPath + "/" + file.Name()
				fmt.Println("loop1:", i)
				fmt.Println("file path for file:", file.Name(), "is", filepath)
				f, _ := os.Open(filepath)
				data, _ := ioutil.ReadAll(f)
				f.Close()
				os.Remove(filepath)
				go func(data string, i int) {

					reader := csv.NewReader(strings.NewReader(data))
					records, _ := reader.ReadAll()
					for _, record := range records {
						fmt.Println("loop2:", i)
						invoice := new(Invoice)
						invoice.Number = record[0]
						invoice.Amount, _ = strconv.ParseFloat(record[1], 64)
						invoice.PurchaseOrderNumber, _ = strconv.Atoi(record[2])
						unixTime, _ := strconv.ParseInt(record[3], 10, 64)
						invoice.InvoiceDate = time.Unix(unixTime, 0)
						tracker++
						fmt.Printf("Record number: %v.Recieve invoice:'%v' for $%.2f and submitted on %v \n", tracker, invoice.Number, invoice.Amount, invoice.InvoiceDate)
					}

				}(string(data), i)

			}
		}

	}

}

type Invoice struct {
	Number              string
	Amount              float64
	PurchaseOrderNumber int
	InvoiceDate         time.Time
}
