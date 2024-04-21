package config

type Imports struct {
	Set map[string]bool
}

var NO_IMPORTS = Imports{
	Set: map[string]bool{},
}
