package db

import (
	"os"
	"path/filepath"

	"github.com/boltdb/bolt"
	"github.com/hjr265/deen/cfg"
)

type Store interface{}

func open() (*bolt.DB, error) {
	if err := os.MkdirAll(filepath.Dir(cfg.Current.Database.Path), 0755); err != nil {
		return nil, err
	}
	return bolt.Open(cfg.Current.Database.Path, 0666, &bolt.Options{})
}
