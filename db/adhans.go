package db

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/boltdb/bolt"
	"github.com/hjr265/deen/aladhan"
	"github.com/hjr265/deen/cfg"
	"github.com/hjr265/deen/model"
)

type Adhans struct {
	db *bolt.DB
}

var _ Store = &Adhans{}

func OpenAdhans() (*Adhans, error) {
	db, err := bolt.Open(cfg.Current.Database.Path, 0666, &bolt.Options{})
	if err != nil {
		return nil, err
	}
	return &Adhans{db}, nil
}

func (s Adhans) GetNext(method aladhan.Method, country, city string, when time.Time) (r *model.Adhan, err error) {
	err = s.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(fmt.Sprintf("adhan:%d:%s:%s", method, city, country)))
		if b == nil {
			return nil
		}

		c := b.Cursor()
		for k, v := c.Seek([]byte(when.Format(time.RFC3339))); k != nil; k, v = c.Next() {
			adhan := model.Adhan{}
			err := json.Unmarshal(v, &adhan)
			if err != nil {
				return err
			}
			if adhan.When.Format("2006-01-02") == when.Format("2006-01-02") || adhan.When.Format("2006-01-02") == when.AddDate(0, 0, 1).Format("2006-01-02") {
				r = &adhan
				break
			}
		}

		return nil
	})
	return
}

func (s Adhans) Put(adhan model.Adhan) error {
	return s.db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(fmt.Sprintf("adhan:%d:%s:%s", adhan.Method, adhan.City, adhan.Country)))
		if err != nil {
			return err
		}

		v, err := json.Marshal(adhan)
		if err != nil {
			return err
		}

		return b.Put([]byte(adhan.When.Format(time.RFC3339)), v)
	})
}
