package main

import (
	"github.com/jinzhu/configor"
)

type Source struct {
	Name   string
	Type   string
	Prefix string
	Coin   string
}

type CheckHeightListItem struct {

}

var Config = struct {
	Name    string `default:"app_name"`
	Sources []Source
}{}

func InitConfig(files string) {
	err := configor.Load(&Config, files)
	if err != nil {
		panic(err)
	}
}
