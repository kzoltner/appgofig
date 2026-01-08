# AppGofig (AppConfig for Go)

Using a struct (and some maps) as single source of truth to add configuration to Go applications.

> [!note]
> This is a very simplistic approach to adding configuration to your applications. It is intended to be one step above just using a config map or something.

## What it does

# Install

Should be a simple go get command:

```bash
go get github.com/kzoltner/appgofig
```

# How to use

First, create a new AppGofig instance with your Config and your ConfigDescriptions:

```go
type Config struct {
	MyOwnSetting    int64  `default:"10" env:"MY_OWN_SETTING"`
}

var configDescriptions map[string]string = map[string]string{
	"MyOwnSetting":    "This is just a simple example description so this map is not empty",
}

// Creating the AppGofig itself
gofig, err := appgofig.NewAppGofig[Config](configDescriptions)
if err != nil {
	log.Fatal(err)
}
```

Next, actuall initialize the config by calling `Init()` - only now values are added:

```go
// Initialize it, which actually reads the values
if err := gofig.Init(); err != nil {
	log.Fatal(err)
}
```

The actual Config is now available using `GetConfig`:

```go
cfg := gofig.GetConfig()
```

This should now have all the fields from the Config struct but with values added.
