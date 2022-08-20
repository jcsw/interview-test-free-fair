package sys

import (
	"encoding/csv"
	os "os"
)

func LoadCsvData(fileName string) [][]string {
	dataFile, err := os.Open(fileName)
	if err != nil {
		LogError("Error on open file", err)
	}
	defer dataFile.Close()

	csvReader := csv.NewReader(dataFile)
	data, err := csvReader.ReadAll()
	if err != nil {
		LogError("Error on read file", err)
	}

	return data
}
