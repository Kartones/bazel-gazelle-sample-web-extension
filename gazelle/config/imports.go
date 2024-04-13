package config

type Imports struct {
	Set map[string]bool
}

var NO_IMPORTS = Imports{
	Set: map[string]bool{},
}

func ReadFileAndParse(filePath string) *Imports {
	fileImports := Imports{
		Set: make(map[string]bool),
	}

	// TODO: Imports are managed here

	return &fileImports
}
