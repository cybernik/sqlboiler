package cmds

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

var cfg = struct {
	Postgres struct {
		User   string `toml:"user"`
		Pass   string `toml:"pass"`
		Host   string `toml:"host"`
		Port   int    `toml:"port"`
		DBName string `toml:"dbname"`
	} `toml:"postgres"`
}{}

func init() {
	_, err := toml.DecodeFile("config.toml", &cfg)
	if err == nil {
		return
	}

	if os.IsNotExist(err) {
		return
	}

	if err != nil {
		fmt.Println("Failed to decode toml configuration file:", err)
	}
}