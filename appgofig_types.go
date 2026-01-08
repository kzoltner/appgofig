package appgofig

import "reflect"

type AppGofig[T StructOnly] struct {
	Cfg          *T
	Descriptions map[string]string
}

type ConfigEntry struct {
	EnvKey     string // Config Key for Environment based inputs
	FieldName  string
	FieldType  reflect.Kind
	IsRequired bool

	RawInput string
	Value    any
}

type StructOnly any
