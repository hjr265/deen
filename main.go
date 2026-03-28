package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/hjr265/deen/adhan"
	"github.com/hjr265/deen/aladhan"
	"github.com/hjr265/deen/cfg"
	"github.com/hjr265/deen/quran"
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
	app.Command("quran:ayah", "Show an ayah of the Quran", cmdQuranAyah)

	app.Run(os.Args)
}

func cmdQuranAyah(cmd *cli.Cmd) {
	cmd.Spec = "REFERENCE"
	var (
		reference = cmd.StringArg("REFERENCE", "", "Ayah reference (e.g. 2:255)")
	)
	cmd.Action = func() {
		ayahs, err := quran.Ayah(*reference, cfg.Current.Quran.Editions)
		if err != nil {
			log.Fatalf("quran: failed to get ayah: %s", err)
		}

		for _, ayah := range ayahs {
			fmt.Printf("[%s %d:%d, %s] %s\n", ayah.SurahNameEn, ayah.SurahNumber, ayah.NumberInSurah, ayah.EditionNameEn, ayah.Text)
		}
	}
}

func cmdAdhanNext(cmd *cli.Cmd) {
	cmd.Action = func() {
		adhan, err := adhan.Next(time.Now(), cfg.Current.Adhan.City, cfg.Current.Adhan.Country, aladhan.Methods[cfg.Current.Adhan.Method])
		if err != nil {
			log.Fatalf("adhan: failed to get next: %s", err)
		}

		if adhan != nil {
			tf := "15:04"
			if cfg.Current.Adhan.TimeFormat == "12h" {
				tf = "3:04 PM"
			}
			fmt.Printf("%s %s\n", adhan.Name, adhan.When.Format(tf))
		}
	}
}
