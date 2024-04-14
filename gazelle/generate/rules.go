package generate

import (
	"github.com/bazelbuild/bazel-gazelle/rule"
	extensionConfig "github.com/kartones/bazel-gazelle-sample-web-extension/gazelle/config"
)

// TODO: process TS tests
var ENABLED_RULES = []func(RuleArgs, *extensionConfig.WebConfig) ([]*extensionConfig.Imports, []*rule.Rule){
	MakeJSLibraryRules,
	MakeJSTestRules,
	MakeTSProjectRules,
}

type IdentifiedSources struct {
	Javascript      []string
	Typescript      []string
	JavascriptTests []string
	TypescriptTests []string
}

type IdentifiedImports struct {
	Javascript      []extensionConfig.Imports
	JavascriptTests []extensionConfig.Imports
	Typescript      []extensionConfig.Imports
	TypescriptTests []extensionConfig.Imports
}

type RuleArgs struct {
	Sources IdentifiedSources
	Imports IdentifiedImports
	Rel     string
}

func GenerateAllRules(availableRules []func(RuleArgs, *extensionConfig.WebConfig) ([]*extensionConfig.Imports, []*rule.Rule), args RuleArgs, webConfig *extensionConfig.WebConfig) ([]*extensionConfig.Imports, []*rule.Rule) {
	rules := []*rule.Rule{}
	rulesImports := []*extensionConfig.Imports{}

	for _, rulesGenerator := range availableRules {
		imports, generatedRules := rulesGenerator(args, webConfig)
		rules = append(rules, generatedRules...)
		rulesImports = append(rulesImports, imports...)
	}

	return rulesImports, rules
}
