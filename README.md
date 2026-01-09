# AppGofig (AppConfig for Go)

Using a struct (and one description map) as single source of truth to add configuration to Go applications.

> [!note]
> This is a very simplistic approach to adding configuration to your applications. It is intended to be one step above just using a config map or something similar.

# Install

Should be a simple go get command:

```bash
go get github.com/kzoltner/appgofig
```

# How to use

First, create your Config and your ConfigDescriptions (optional, only needed for documentation):

```go
type Config struct {
	MyOwnSetting    int64  `default:"10" env:"MY_OWN_SETTING"`
	MyStringSetting string `default:"hello" env:"MY_STRING_SETTING" req:"true"`
}

var configDescriptions map[string]string = map[string]string{
	"MyOwnSetting":    "This is just a simple example description so this map is not empty",
	"MyStringSetting": "This is just a string setting that is empty but required.",
}
```

Next, instantiate your config and then read the config values by calling `ReadConfig()`:

```go
cfg := &Config{}

if err := appgofig.ReadConfig(cfg, appgofig.ReadModeEnvThenYaml); err != nil {
	log.Fatal(err)
}

appgofig.LogConfig(cfg, os.Stdout)
```

Now, using your config should be as easy as accessing the struct itself:

```go
log.Println(cfg.MyOwnSetting)
```

# The `Config` struct

The `Config` struct determines your whole configuration. You can name it whatever you want.
The following tags are usable:

| Tag Name  | Content                                                                               |
| --------- | ------------------------------------------------------------------------------------- |
| `env`     | Key used for Environment Variables. If this is empty, it defaults to the field name   |
| `default` | String representation of a default value. Otherwise an empty string is used.          |
| `req`     | If set to "true", this config setting cannot be empty. Only applies to string values. |

Example entry:

```go
type Config struct {
	MyOwnSetting string `env:"ENV_MY_OWN_SETTING" default:"myDefaultValue" req:"true"`
}
```

> [!important]
> Due to my own needs, only four types are allowed: `string`, `int64`, `float64` and `bool`.

# ReadModes

There are four read modes available:

| ReadMode                     | Description                                                          |
| ---------------------------- | -------------------------------------------------------------------- |
| `appgofig.ReadModeEnvOnly`     | Only uses environment to read values                                 |
| `appgofig.ReadModeYamlOnly`    | Only use yaml file                                                   |
| `appgofig.ReadModeEnvThenYaml` | First read env, then apply yaml (overwriting env values if present)  |
| `appgofig.ReadModeYamlThenEnv` | First read yaml, then apply env (overwriting yaml values if present) |

# Documentation
Two methods are provided to automatically create documentation about your configuration.
Check the `example` folder for how they could look like.

## Markdown 
Using `WriteToMarkdownFile()` you can generate a markdown file containing a simple table.

```go
if err := appgofig.WriteToMarkdownFile(cfg, configDescriptions, "example/MarkdownExample.md"); err != nil {
	log.Fatal(err)
}
```

## Example config.yaml
Similarly, using `WriteToYamlExampleFile()` will get you a config.yml example with comments explaining each entry

```go
if err := appgofig.WriteToYamlExampleFile(cfg, configDescriptions, "example/ConfigYamlExample.yaml"); err != nil {
	log.Fatal(err)
}
```
