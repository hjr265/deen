package cfg

import "os/user"

type Configuration struct {
	Database struct {
		Path string
	}
}

func (c *Configuration) SetDefaults() error {
	u, err := user.Current()
	if err != nil {
		return err
	}
	c.Database.Path = u.HomeDir + "/.deendb"
	return nil
}

var Current Configuration

func Load() error {
	Current.SetDefaults()

	return nil
}
