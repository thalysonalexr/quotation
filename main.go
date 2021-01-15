package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/gocraft/work"
	"github.com/joho/godotenv"

	"quotation/quotation"

	"github.com/fatih/color"
)

// Context job
type Context struct{}

func (c *Context) getQuotation(job *work.Job) error {
	color.Yellow("Init job " + job.Name + ":" + job.ID)
	if err := quotation.RunQuotation(); err != nil {
		color.Red(err.Error())
		return err
	}
	return nil
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln(err)
	}

	pool := quotation.ConfigRedis()
	wp := work.NewWorkerPool(Context{}, 1, "quotation_workspace", pool)
	wp.PeriodicallyEnqueue(os.Getenv("CRON_DESCRIPTOR"), "get_quotation")
	wp.Job("get_quotation", (*Context).getQuotation)
	color.Green("Init process to get quotation, please waiting...")
	wp.Start()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, os.Kill)
	<-signalChan

	wp.Stop()
}
