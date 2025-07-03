package tomlread

import (
	"log"

	"github.com/BurntSushi/toml"
)

func ReadTomlConfig(filename string) interface{} {
	var f interface{}
	if _, err := toml.DecodeFile(filename, &f); err != nil {
		log.Println(err)
	}
	return f
}
