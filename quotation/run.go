package quotation

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	pb "github.com/cheggaaa/pb/v3"
	"github.com/fatih/color"
)

func saveDataCSV(quotations *[]Quotation, dir string) error {
	var lines [][]string
	ini := time.Now()
	bar := pb.Full.Start(len(*quotations))

	lines = append(lines, *GetDescriptionKeys())

	for _, q := range *quotations {
		line := ToArrayString(&q)
		lines = append(lines, *line)
		bar.Increment()
	}
	var err error
	path := os.Getenv("PATH_SAVE_FILES")
	file := NewFilenameWithDate("quotation", "csv")

	if path == "" {
		path = dir
	}

	err = ExportCsv(lines, filepath.FromSlash(path+"/"+file))
	if err != nil {
		return err
	}

	bar.Finish()
	color.Cyan("Process save CSV finish. Duration %f secs\n", time.Since(ini).Seconds())
	return nil
}

// RunQuotation run process of download and export quotation to csv
func RunQuotation(dir string) error {
	const max, min = 100, 50
	q, err := getQuotationOptionsTableNonce()
	if err != nil {
		log.Fatalln(err)
	}

	var page int = 1
	var quotations []Quotation
	bar := pb.Full.Start(rand.Intn(max - min))

	for {
		data, err := GetQuotation(q.AjaxURL, makeFormData(page, q.CotacoesOpcoesTableNonce))
		if err != nil {
			color.Red("Error to get quotation in page %d\n", page)
		}
		bar.Increment()
		unpack, err := NewQuotationResponse(data)
		if err != nil {
			color.Red(err.Error())
			break
		}
		color.Green("Request of page {%d} with sucessfully", page)
		CallClear()
		if !unpack.HasMore {
			color.Yellow("No more pages to request")
			break
		}
		page++
		quotations = append(quotations, unpack.Options...)
	}

	if err = saveDataCSV(&quotations, dir); err != nil {
		fmt.Println(err.Error())
		color.Red("Error to export data csv", err)
		return err
	}
	bar.Finish()
	return nil
}
