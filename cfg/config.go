package cfg

import (
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

type Configuration struct {
	Database struct {
		Path string `toml:"path"`
	} `toml:"database"`

	Adhan struct {
		Method     int    `toml:"method"`
		City       string `toml:"city"`
		Country    string `toml:"country"`
		TimeFormat string `toml:"time_format"`
	} `toml:"adhan"`

	Quran struct {
		Editions []string `toml:"editions"`
	} `toml:"quran"`
}

func (c *Configuration) SetDefaults() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	c.Database.Path = filepath.Join(home, ".local", "share", "deen", "deen.db")
	c.Adhan.TimeFormat = "24h"
	c.Quran.Editions = []string{"quran-uthmani", "en.sahih"}
	return nil
}

var Current Configuration

func Load(configPath string) error {
	Current.SetDefaults()

	if configPath == "" {
		configDir, err := os.UserConfigDir()
		if err != nil {
			return err
		}
		configPath = filepath.Join(configDir, "deen", "config.toml")
	}

	_, err := toml.DecodeFile(configPath, &Current)
	return err
}
