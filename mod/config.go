package mod

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

const (
	SecurityStrategy
)

type Config struct {
	Mod struct {
		Security struct {
			Strategy string
		}
	}
	Config struct {
		Driver string
		Config map[string]interface{}
	}
}

func ParseConf() (*Config, error) {
	p := os.Getenv("MOD_CONF")
	if p == "" {
		p = "mod.yml"
	}

	if _, err := os.Stat(p); os.IsNotExist(err) {
		return nil, fmt.Errorf(`mod: config file not exist: %s`, p)
	}
	data, err := ioutil.ReadFile(p)
	if err != nil {
		return nil, fmt.Errorf(`mod: parse config file failed: %s`, err.Error())
	}
	c := new(Config)
	if err := yaml.Unmarshal(data, c); err != nil {
		log.Fatalf("error: %v", err)
	}
	return c, nil
}
