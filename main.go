package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/joho/godotenv"
	cron "github.com/lovego/redis-cron"

	"quotation/quotation"

	"github.com/fatih/color"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln(err)
	}

	pool := quotation.ConfigRedis()
	c := cron.New(pool)

	c.AddFunc(os.Getenv("CRON_DESCRIPTOR"), func() {
		color.Yellow("Init job func getQuotation")
		path, err := os.Getwd()
		if err != nil {
			color.Red(err.Error())
		} else if err := quotation.RunQuotation(path + "/tmp"); err != nil {
			color.Red(err.Error())
		}
	})

	c.Start()
	color.Green("Init process to get quotation, please waiting job run...")

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, os.Kill)
	<-signalChan

	c.Stop()
}
