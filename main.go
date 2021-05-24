package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/hjr265/deen/adhan"
	"github.com/hjr265/deen/aladhan"
	"github.com/hjr265/deen/cfg"
	cli "github.com/jawher/mow.cli"
)

var app *cli.Cli

func main() {
	err := cfg.Load()
	if err != nil {
		log.Fatalf("cfg: failed to load: %s", err)
	}

	app = cli.App("deen", "Command line companion to get prayer timings, Quran verses, and more")

	app.Command("adhan:next", "Show adhan timing of next prayer", cmdAdhanNext)

	app.Run(os.Args)
}

func cmdAdhanNext(cmd *cli.Cmd) {
	cmd.Action = func() {
		adhan, err := adhan.Next(time.Now(), cfg.Current.Adhan.City, cfg.Current.Adhan.Country, aladhan.Methods[cfg.Current.Adhan.Method])
		if err != nil {
			log.Fatalf("adhan: failed to get next: %s", err)
		}

		if adhan != nil {
			fmt.Printf("%s %s\n", adhan.Name, adhan.When.Format("15:04"))
		}
	}
}
