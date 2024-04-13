package config

import (
	"path/filepath"
)

type WebConfig struct {
	ExtensionEnabled bool
	Quiet            bool
	Verbose          bool
}

type WebConfigs map[string]*WebConfig

func (webConfigs *WebConfigs) ParentForPackage(pkg string) *WebConfig {
	dir := filepath.Dir(pkg)
	if dir == "." {
		dir = ""
	}
	parent := (map[string]*WebConfig)(*webConfigs)[dir]
	return parent
}

func NewWebConfig() *WebConfig {
	return &WebConfig{
		ExtensionEnabled: true,
	}
}

// New configurations inherit desired values from the parent.
// It is best to copy maps to avoid mutating the parent.
func (parent *WebConfig) NewChild() *WebConfig {
	child := NewWebConfig()

	child.ExtensionEnabled = parent.ExtensionEnabled
	child.Quiet = parent.Quiet
	child.Verbose = parent.Verbose

	return child
}
