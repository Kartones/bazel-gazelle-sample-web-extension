package parse

import (
	"log"
	"os"

	extensionConfig "github.com/kartones/bazel-gazelle-sample-web-extension/gazelle/config"
	extensionUtils "github.com/kartones/bazel-gazelle-sample-web-extension/gazelle/utils"
)

func ReadFileAndParse(filePath string) *extensionConfig.Imports {
	fileImports := extensionConfig.Imports{
		Set: make(map[string]bool),
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf(extensionUtils.Err("Error reading %s: %v", filePath, err))
	}

	imports, err := ParseJS(data)
	if err != nil {
		log.Fatalf(extensionUtils.Err("Error parsing %s: %v", filePath, err))
	}

	for _, imp := range imports {
		fileImports.Set[imp] = true
	}

	return &fileImports
}
