package appgofig

import (
	"log"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/joho/godotenv"
	"go.yaml.in/yaml/v4"
)

// createConfigMap create a ConfigEntry mapped to every name within T while applying tags and also setting the default if present within the map
func createConfigMap[T StructOnly]() map[string]*ConfigEntry {
	// create map for config entries
	configEntries := make(map[string]*ConfigEntry)

	// loop over all entries in T to initialize all entries
	t := reflect.TypeFor[T]()
	for k := 0; k < t.NumField(); k++ {
		f := t.Field(k)

		entry := &ConfigEntry{
			EnvKey:     f.Tag.Get("env"),
			FieldName:  f.Name,
			FieldType:  f.Type.Kind(),
			IsRequired: f.Tag.Get("req") == "true",
			RawInput:   f.Tag.Get("default"),
		}

		configEntries[entry.FieldName] = entry
	}

	return configEntries
}

// applyEnvConfigIfPresent applies environment values to the configuration while loading .env files first
func applyEnvConfigIfPresent(configEntries map[string]*ConfigEntry) {
	// load .env
	// error is ignored on purpose, as not having .env is not an issue
	if err := godotenv.Load(); err != nil {
		log.Println("AppGofig: No .env file. Ignored.")
	}

	// overwrite existing entries with env entries
	for _, entry := range configEntries {
		val, isPresent := os.LookupEnv(entry.EnvKey)
		if isPresent {
			entry.RawInput = strings.TrimSpace(val)
		}
	}
}

// applyYamlConfigIfPresent checks for (config/)config.y(a)ml files and applies the first one found to configuration
func applyYamlConfigIfPresent(configEntries map[string]*ConfigEntry) error {
	// check for a config.yml or config.yaml in root directory or within a config folder
	// any value here overwrites the rest
	YAML_PATHS := []string{"config.yml", "config.yaml", "config/config.yml", "config/config.yaml"}

	yamlFilePath := ""
	for _, path := range YAML_PATHS {
		if _, err := os.Stat(path); err == nil {
			yamlFilePath = path
			break
		}
	}

	if len(yamlFilePath) == 0 {
		log.Println("AppGofig: Did not find config.y(a)ml or config/config.y(a)ml. Ignored.")
		return nil
	}

	log.Printf("AppGofig: Using yaml %q", yamlFilePath)
	data, err := os.ReadFile(filepath.Clean(yamlFilePath))
	if err != nil {
		return err
	}

	yamlMap := make(map[string]string)

	if err := yaml.Unmarshal(data, &yamlMap); err != nil {
		return err
	}

	for _, entry := range configEntries {
		// here both key variants apply - if someone likes them or something, I don't know
		val, isPresent := yamlMap[entry.EnvKey]
		if isPresent {
			entry.RawInput = strings.TrimSpace(val)
		}

		// yamlKey takes priority though
		val, isPresent = yamlMap[entry.FieldName]
		if isPresent {
			entry.RawInput = strings.TrimSpace(val)
		}
	}

	return nil
}
