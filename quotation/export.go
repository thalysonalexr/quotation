package quotation

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"
)

// NewFilenameWithDate create filename with datetime
func NewFilenameWithDate(filename, ext string) string {
	now := time.Now()
	hour, min, sec := now.Clock()
	year, month, day := now.Local().Date()

	date := strconv.Itoa(day) + "-" + month.String() + "-" + strconv.Itoa(year)
	clock := strconv.Itoa(hour) + "_" + strconv.Itoa(min) + "_" + strconv.Itoa(sec)
	return filename + "-" + date + "_" + clock + "." + ext
}

// ExportCsv export data to csv
func ExportCsv(data [][]string, filePath string) error {
	fmt.Println(filePath)
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
