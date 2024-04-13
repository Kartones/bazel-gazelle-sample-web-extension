package generate

import (
	"fmt"

	"github.com/bazelbuild/bazel-gazelle/rule"
	extensionConfig "github.com/kartones/bazel-gazelle-sample-web-extension/gazelle/config"
	"github.com/kartones/bazel-gazelle-sample-web-extension/gazelle/constants"
	extensionFile "github.com/kartones/bazel-gazelle-sample-web-extension/gazelle/file"
)

func MakeJSTestRules(args RuleArgs, webConfig *extensionConfig.WebConfig) ([]*extensionConfig.Imports, []*rule.Rule) {
	rules := []*rule.Rule{}
	rulesImports := []*extensionConfig.Imports{}

	for _, sourceFile := range args.Sources.JavascriptTests {
		newRule := makeBaseRule(BaseRuleArgs{
			Kind:    "js_test",
			Name:    fmt.Sprintf("test_%s", extensionFile.FileBaseName(sourceFile)),
			Package: args.Rel,
		}, webConfig)

		newRule.SetAttr(constants.ATTR_ENTRY_POINT, sourceFile)
		newRule.SetAttr(constants.ATTR_SIZE, "small")

		// TODO: process imports
		imports := extensionConfig.Imports{
			Set: make(map[string]bool),
		}
		// TODO: process data attribute

		rules = append(rules, newRule)
		rulesImports = append(rulesImports, &imports)
	}

	return rulesImports, rules
}
