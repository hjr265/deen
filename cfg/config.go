package cfg

import (
	"os/user"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

type Configuration struct {
	Database struct {
		Path string
	}

	Adhan struct {
		Method  int
		City    string
		Country string
	}
}

func (c *Configuration) SetDefaults() error {
	u, err := user.Current()
	if err != nil {
		return err
	}
	c.Database.Path = filepath.Join(u.HomeDir, "/.deendb")
	return nil
}

var Current Configuration

func Load() error {
	Current.SetDefaults()

	u, err := user.Current()
	if err != nil {
		return err
	}
	_, err = toml.DecodeFile(filepath.Join(u.HomeDir, "/.config/deen/config.toml"), &Current)
	return err
}
