package quotation

import (
	"log"
	"os"
	"path/filepath"
	"time"

	pb "github.com/cheggaaa/pb/v3"
	"github.com/fatih/color"
)

func saveDataCSV(quotations *[]Quotation) error {
	var lines [][]string
	ini := time.Now()
	bar := pb.Full.Start(len(*quotations))

	lines = append(lines, *GetDescriptionKeys())

	for _, q := range *quotations {
		line := ToArrayString(&q)
		lines = append(lines, *line)
		bar.Increment()
	}

	filename := newFilenameWithDate("quotation") + ".csv"
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}

	err = ExportCsv(lines, filepath.FromSlash(currentDir+"/tmp/"+filename))
	if err != nil {
		return err
	}

	bar.Finish()
	color.Cyan("Process save CSV finish. Duration %f secs\n", time.Since(ini).Seconds())
	return nil
}

// RunQuotation run process of download and export quotation to csv
func RunQuotation() error {
	q, err := getQuotationOptionsTableNonce()
	if err != nil {
		log.Fatalln(err)
	}

	var page int = 1
	var quotations []Quotation
	bar := pb.Full.Start(50)

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
		if !unpack.HasMore {
			color.Yellow("No more pages to request")
			break
		}
		page++
		quotations = append(quotations, unpack.Options...)
	}

	if err = saveDataCSV(&quotations); err != nil {
		color.Red("Error to export data csv", err)
		return err
	}
	bar.Finish()
	return nil
}
