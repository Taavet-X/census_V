package utils

import (
	"encoding/csv"
	"os"
	"time"
)

func ExportToCSV(data [][]string) (string, error) {
	fileName := "source_data" + time.Now().Format("20060102150405") + ".csv"
	file, err := os.Create(fileName)
	if err != nil {
		return "", err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, record := range data {
		if err := writer.Write(record); err != nil {
			return "", err
		}
	}
	return fileName, nil
}
