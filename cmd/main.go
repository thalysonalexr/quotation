package main

import (
	"os"
	"path/filepath"
	"quotation/quotation"
	"time"

	"github.com/fatih/color"
)

func readCommandSaveByPath() (string, error) {
	if len(os.Args) == 2 {
		command := os.Args[1]
		return command, nil
	}
	path, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return filepath.FromSlash(path + "/tmp"), nil
}

func main() {
	ini := time.Now()
	defaultPath, err := readCommandSaveByPath()
	if err != nil {
		color.Red(err.Error())
		os.Exit(1)
	}
	color.Green("Init process to get quotation, please waiting...")
	if err := quotation.RunQuotation(defaultPath); err != nil {
		color.Red(err.Error())
	}
	color.Yellow("Finish process %f secs\n", time.Since(ini).Seconds())
}
