package aladhan

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sort"
	"time"
)

type Timing struct {
	Name string
	When time.Time
}

func GetTimings(date time.Time, city, country string, method Method) ([]Timing, error) {
	url := fmt.Sprintf("http://api.aladhan.com/v1/timingsByCity/%s?city=%s&country=%s&method=%d", date.Format("2006-01-02"), city, country, method)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errors.New("aladhan: non-200 status")
	}
	defer resp.Body.Close()

	body := struct {
		Data struct {
			Timings map[string]string
		}
	}{}
	err = json.NewDecoder(resp.Body).Decode(&body)
	if err != nil {
		return nil, err
	}

	timings := []Timing{}
	names := []string{"Fajr", "Dhuhr", "Asr", "Maghrib", "Isha"}
	for _, name := range names {
		var hour, minute int
		fmt.Sscanf(body.Data.Timings[name], "%2d:%2d", &hour, &minute)

		timings = append(timings, Timing{
			Name: name,
			When: time.Date(date.Year(), date.Month(), date.Day(), hour, minute, 0, 0, time.Local),
		})
	}

	sort.Slice(timings, func(i, j int) bool {
		return timings[i].When.Before(timings[j].When)
	})

	return timings, nil
}
