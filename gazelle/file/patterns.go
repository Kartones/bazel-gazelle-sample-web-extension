package file

import (
	"fmt"
	"regexp"
	"strings"
)

var JSExtensions = []string{
	".js",
	".jsx",
	".mjs",
	".cjs",
}

var TSExtensions = []string{
	".ts",
	".tsx",
	".mts",
	".cts",
}

var jsTestExtensions = []string{
	".test.js",
	".test.jsx",
	".test.mjs",
	".test.cjs",
}

var tsTestExtensions = []string{
	".test.ts",
	".test.tsx",
	".test.mts",
	".test.cts",
}

var (
	JSTestExtensionsPattern *regexp.Regexp
	TSTestExtensionsPattern *regexp.Regexp
	JSExtensionsPattern     *regexp.Regexp
	TSExtensionsPattern     *regexp.Regexp
	trimExtPattern          *regexp.Regexp
	fileBaseNamePattern     *regexp.Regexp
)

func init() {
	JSTestExtensionsPattern = extensionPattern(jsTestExtensions, "")
	TSTestExtensionsPattern = extensionPattern(tsTestExtensions, "")
	JSExtensionsPattern = extensionPattern(JSExtensions, "")
	TSExtensionsPattern = extensionPattern(TSExtensions, "")

	escaped := make([]string, len(TSExtensions)+len(JSExtensions))
	for i, ext := range append(TSExtensions, JSExtensions...) {
		escaped[i] = regexp.QuoteMeta(ext)
	}

	trimExtPattern = regexp.MustCompile(
		fmt.Sprintf(`(\S+)(%s)$`,
			strings.Join(escaped, "|"),
		),
	)
	fileBaseNamePattern = regexp.MustCompile(`([^\.]+)`)
}

func extensionPattern(extensions []string, extensionPrefix string) *regexp.Regexp {
	escaped := make([]string, len(extensions))
	for i := range extensions {
		escaped[i] = fmt.Sprintf("(%s%s$)", regexp.QuoteMeta(extensionPrefix), regexp.QuoteMeta(extensions[i]))
	}
	return regexp.MustCompile(strings.Join(escaped, "|"))
}
