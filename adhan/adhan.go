package adhan

import (
	"log"
	"time"

	"github.com/hjr265/deen/aladhan"
	"github.com/hjr265/deen/db"
	"github.com/hjr265/deen/model"
)

func Next(when time.Time, city, country string, method aladhan.Method) (*model.Adhan, error) {
	adhans, err := db.OpenAdhans()
	if err != nil {
		return nil, err
	}

	adhan, err := adhans.GetNext(method, country, city, when)
	if err != nil {
		return nil, err
	}
	if adhan != nil {
		return adhan, nil
	}

	log.Println("Downloading adhan timings")

	err = syncTimings(adhans, when, city, country, method)
	if err != nil {
		return nil, err
	}
	err = syncTimings(adhans, when.AddDate(0, 0, 1), city, country, method)
	if err != nil {
		return nil, err
	}

	return adhans.GetNext(method, country, city, when)
}

func syncTimings(adhans *db.Adhans, when time.Time, city, country string, method aladhan.Method) error {
	timings, err := aladhan.GetTimings(when, city, country, method)
	if err != nil {
		return err
	}

	for _, timing := range timings {
		err = adhans.Put(model.Adhan{
			Name:    timing.Name,
			When:    timing.When,
			City:    city,
			Country: country,
			Method:  method,
		})
		if err != nil {
			return err
		}
	}

	return nil
}
