package alquran

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type Ayah struct {
	Text          string
	NumberInSurah int
	Surah         struct {
		Number                 int
		Name                   string
		EnglishName            string
		EnglishNameTranslation string
	}
	Edition struct {
		Identifier  string
		Language    string
		Name        string
		EnglishName string
	}
}

func GetAyah(reference string, editions []string) ([]Ayah, error) {
	url := fmt.Sprintf("https://api.alquran.cloud/v1/ayah/%s/editions/%s", reference, strings.Join(editions, ","))
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, errors.New("alquran: non-200 status")
	}

	body := struct {
		Data []Ayah
	}{}
	err = json.NewDecoder(resp.Body).Decode(&body)
	if err != nil {
		return nil, err
	}

	return body.Data, nil
}
