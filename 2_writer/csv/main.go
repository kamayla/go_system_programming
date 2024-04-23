package main

import (
	"encoding/csv"
	"log"
	"os"
)

func main() {
	file, err := os.Create("hoge.csv")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	w := csv.NewWriter(file)

	records := [][]string{
		[]string{"名前", "年齢", "出身地", "性別"},
		[]string{"ippei", "42", "新潟", "男性"},
		[]string{"manko", "18", "沖縄", "女性"},
	}

	w.WriteAll(records)
	w.Write([]string{"さとし", "33", "福井", "男"})
	w.Write([]string{"さとし", "33", "福井", "男"})
	w.Write([]string{"さとし", "33", "福井", "男"})
	w.Flush()

	// Flushしてないから書き込まれない
	w.Write([]string{"さとし", "33", "福井", "男"})
	w.Write([]string{"さとし", "33", "福井", "男"})
	w.Write([]string{"さとし", "33", "福井", "男"})

}
