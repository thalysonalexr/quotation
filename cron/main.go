package main

import (
	"os"

	"quotation/quotation"

	"github.com/fatih/color"
	"github.com/jasonlvhit/gocron"
)

func getQuotation() error {
	color.Yellow("Init job to get quotation")
	path, err := os.Getwd()
	if err != nil {
		color.Red(err.Error())
		return err
	}
	if err := quotation.RunQuotation(path + "/tmp"); err != nil {
		color.Red(err.Error())
		return err
	}
	return nil
}

func main() {
	s := gocron.NewScheduler()
	s.Every(5).Minutes().Do(getQuotation)
	<-s.Start()
}
