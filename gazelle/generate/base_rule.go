package generate

import (
	"log"

	"github.com/bazelbuild/bazel-gazelle/rule"
	extensionConfig "github.com/kartones/bazel-gazelle-sample-web-extension/gazelle/config"
	extensionUtils "github.com/kartones/bazel-gazelle-sample-web-extension/gazelle/utils"
)

var invalidTargetNames = []string{
	"all",
	"all-targets",
	"__pkg__",
	"__subpackages__",
}

type BaseRuleArgs struct {
	Kind    string
	Name    string
	Package string
}

func makeBaseRule(args BaseRuleArgs, webConfig *extensionConfig.WebConfig) *rule.Rule {
	if extensionUtils.Contains(invalidTargetNames, args.Name) {
		log.Fatalf(extensionUtils.Err("Directory name invalid as Bazel target name: %s", args.Name))
	}

	newRule := rule.NewRule(args.Kind, args.Name)

	// Other custom common attributes can be added here

	return newRule
}
