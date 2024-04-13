package generate

import (
	"github.com/bazelbuild/bazel-gazelle/rule"
	extensionConfig "github.com/kartones/bazel-gazelle-sample-web-extension/gazelle/config"
	"github.com/kartones/bazel-gazelle-sample-web-extension/gazelle/constants"
	extensionFile "github.com/kartones/bazel-gazelle-sample-web-extension/gazelle/file"
)

func MakeJSLibraryRules(args RuleArgs, webConfig *extensionConfig.WebConfig) ([]*extensionConfig.Imports, []*rule.Rule) {
	rules := []*rule.Rule{}
	rulesImports := []*extensionConfig.Imports{}

	for _, sourceFile := range args.Sources.Javascript {
		newRule := makeBaseRule(BaseRuleArgs{
			Kind:    "js_library",
			Name:    extensionFile.FileBaseName(sourceFile),
			Package: args.Rel,
		}, webConfig)

		newRule.SetAttr(constants.ATTR_SOURCES, []string{sourceFile})

		// TODO: process imports
		imports := extensionConfig.Imports{
			Set: make(map[string]bool),
		}

		rules = append(rules, newRule)
		rulesImports = append(rulesImports, &imports)
	}

	return rulesImports, rules
}
