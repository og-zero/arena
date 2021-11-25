package config

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v3"
)

type (
	Config interface{}

	config struct {
		data map[string]interface{}
		path string
	}
)

var dir = getCFGDir()

func Load(path ...string) (conf *config) {
	if len(path) > 0 {
		conf = &config{
			path: path[0],
		}

		conf.path = fmt.Sprintf("%s/server.yaml", conf.path)
		conf.bindYAML()
	}

	fmt.Printf("Configs:\n%v\n", conf)
	return conf
}

func (conf *config) bindYAML() {
	body, err := ioutil.ReadFile(conf.path)
	if err != nil {
		fmt.Println()
		panic(err)
	}

	data := map[string]interface{}{}
	yaml.Unmarshal(body, &data)
	conf.data = data
}

func getCFGDir() string {
	// /Users/zero/code/gd/go/arena/configs
	return os.Getenv("ARENA_CONFIGS_DIR")
}
