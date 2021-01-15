package quotation

import (
	"encoding/csv"
	"os"
	"strconv"
	"time"
)

func newFilenameWithDate(filename string) string {
	now := time.Now()
	year, month, day := now.Local().Date()
	hour, min, sec := now.Clock()
	date := strconv.Itoa(day) + "-" + month.String() + "-" + strconv.Itoa(year)
	clock := strconv.Itoa(hour) + "_" + strconv.Itoa(min) + "_" + strconv.Itoa(sec)
	return filename + "-" + date + "_" + clock
}

// ExportCsv export data to csv
func ExportCsv(data [][]string, filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	writer.Comma = ';'
	defer writer.Flush()

	for i := range data {
		if err := writer.Write(data[i]); err != nil {
			return err
		}
	}
	return nil
}
