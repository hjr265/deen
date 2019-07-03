package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/hjr265/deen/adhan"
	"github.com/hjr265/deen/aladhan"
	"github.com/hjr265/deen/cfg"
)

func main() {
	flag.Parse()

	err := cfg.Load()
	if err != nil {
		log.Fatalf("cfg: failed to load: %s", err)
	}

	switch flag.Arg(0) {
	case "adhan:next":
		adhan, err := adhan.Next(time.Now(), cfg.Current.Adhan.City, cfg.Current.Adhan.Country, aladhan.Methods[cfg.Current.Adhan.Method])
		if err != nil {
			log.Fatalf("adhan: failed to get next: %s", err)
		}

		if adhan != nil {
			fmt.Printf("%s %s\n", adhan.Name, adhan.When.Format("15:04"))
		}
	}
}
