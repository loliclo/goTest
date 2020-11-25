package tomlexec

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

// TConfig ...
type TConfig struct {
	TestToml *TomlConfig
}

// TomlConfig ...
type TomlConfig struct {
	No   int
	Name string
}

// Exec ...
func Exec() {
	var cfg TConfig
	_, err := toml.DecodeFile("./config/testConfig.toml", &cfg)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(cfg.TestToml.Name)
}
