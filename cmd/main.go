package main

import (
	"quotation/quotation"
	"time"

	"github.com/fatih/color"
)

func main() {
	ini := time.Now()
	color.Green("Init process to get quotation, please waiting...")
	if err := quotation.RunQuotation(); err != nil {
		color.Red(err.Error())
	}
	color.Yellow("Finish process %f secs\n", time.Since(ini).Seconds())
}
