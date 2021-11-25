package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type (
	Config interface{
		
	}

	config struct {
		data map[string]
		path string
	}
)

var (
	dir = getCFGDir()
	cfg = Load(dir)
)

func Load(path ...string) (conf *config) {
	if len(path) > 0 {
		conf = &config{
			path: path[0],
		}

		conf.data = fmt.Sprintf("%s/server.yaml", conf.path)
		yaml.Unmarshal()
	} 

	fmt.Printf("Configs:\n%v\n", conf)
	return conf
}

func getCFGDir() string {
	// /Users/zero/code/gd/go/arena/configs
	return os.Getenv("ARENA_CONFIGS_DIR")
}
