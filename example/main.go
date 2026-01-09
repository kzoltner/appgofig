package main

import (
	"log"
	"os"

	"github.com/kzoltner/appgofig"
)

type Config struct {
	MyOwnSetting    int64  `default:"10" env:"MY_OWN_SETTING"`
	MyStringSetting string `default:"hello" env:"MY_STRING_SETTING" req:"true"`
}

var configDescriptions map[string]string = map[string]string{
	"MyOwnSetting":    "This is just a simple example description so this map is not empty",
	"MyStringSetting": "This is just a string setting that is empty but required.",
}

func main() {
	cfg := &Config{}

	if err := appgofig.ReadConfig(cfg, appgofig.EnvThenYaml); err != nil {
		log.Fatal(err)
	}

	appgofig.LogConfig(cfg, os.Stdout)

	log.Println(cfg.MyOwnSetting)

	appgofig.WriteToMarkdownFile(cfg, configDescriptions, "example/MarkdownExample.md")
	appgofig.WriteToYamlExampleFile(cfg, configDescriptions, "example/ConfigYamlExample.yaml")
}
