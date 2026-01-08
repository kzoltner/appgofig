package main

import (
	"log"

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
	// Creating the AppGofig itself
	gofig, err := appgofig.NewAppGofig[Config](configDescriptions)
	if err != nil {
		log.Fatal(err)
	}

	// Initialize it, which actually reads the values
	if err := gofig.Init(); err != nil {
		log.Fatal(err)
	}

	// Get the config struct itself - this should be the same as your struct - but with values added
	cfg := gofig.GetConfig()
	log.Println(cfg.MyOwnSetting)
	log.Println(cfg.MyStringSetting)

	// printing actual (masked) config to console
	gofig.LogToConsole()

	// Documentation functions - can be integrated into automation to create Markdown/ExampleYaml
	appgofig.CreateMarkdownFile[Config]("./example/MarkdownExample.md", configDescriptions)
	appgofig.CreateYamlExampleFile[Config]("./example/ConfigYamlExample.yaml", configDescriptions)
}
