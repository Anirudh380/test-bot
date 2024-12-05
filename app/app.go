package app

import (
	"encoding/csv"
	"fmt"
	"os"
	"sync"
)

func getInstrumentKeys() []string {
	file, err := os.Open("complete.csv")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return []string{}
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		return []string{}
	}
	res := []string{}
	for i, r := range records {
		if i == 0 {
			continue
		}
		res = append(res, r[0])
	}
	return res
}

type App interface {
	Start()
}

func NewTestBotApp() App {
	return &impl{}
}

type impl struct {
}

func (i *impl) Start() {
	var wg sync.WaitGroup
	res := getInstrumentKeys()
	fmt.Println(res)

	wg.Wait()
}
