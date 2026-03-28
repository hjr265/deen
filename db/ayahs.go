package db

import (
	"encoding/json"
	"fmt"

	"github.com/boltdb/bolt"
	"github.com/hjr265/deen/model"
)

type Ayahs struct {
	db *bolt.DB
}

var _ Store = &Ayahs{}

func OpenAyahs() (*Ayahs, error) {
	db, err := open()
	if err != nil {
		return nil, err
	}
	return &Ayahs{db}, nil
}

func (s Ayahs) Get(edition, reference string) (*model.Ayah, error) {
	var ayah *model.Ayah
	err := s.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(fmt.Sprintf("ayah:%s", edition)))
		if b == nil {
			return nil
		}

		v := b.Get([]byte(reference))
		if v == nil {
			return nil
		}

		ayah = &model.Ayah{}
		return json.Unmarshal(v, ayah)
	})
	return ayah, err
}

func (s Ayahs) Put(reference string, ayah model.Ayah) error {
	return s.db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(fmt.Sprintf("ayah:%s", ayah.Edition)))
		if err != nil {
			return err
		}

		v, err := json.Marshal(ayah)
		if err != nil {
			return err
		}

		return b.Put([]byte(reference), v)
	})
}
