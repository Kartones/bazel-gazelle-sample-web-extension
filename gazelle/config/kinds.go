package config

import (
	"github.com/bazelbuild/bazel-gazelle/rule"
	"github.com/kartones/bazel-gazelle-sample-web-extension/gazelle/constants"
)

// Map of maps rule names (kinds) and information on how to match and merge attributes that may be found
// in rules of those kinds.
// This extension will only generate rules of the kinds defined here.
var SupportedKinds = map[string]rule.KindInfo{
	constants.RULE_JS_LIBRARY: {
		MatchAny: false,
		NonEmptyAttrs: map[string]bool{
			constants.ATTR_SOURCES: true,
		},
		MergeableAttrs: map[string]bool{
			constants.ATTR_SOURCES: true,
			constants.ATTR_TAGS:    true,
		},
		ResolveAttrs: map[string]bool{
			constants.ATTR_DEPENDENCIES: true,
			constants.ATTR_DATA:         true,
		},
	},
	constants.RULE_JS_TEST: {
		MatchAny: false,
		NonEmptyAttrs: map[string]bool{
			constants.ATTR_ENTRY_POINT: true,
		},
		MergeableAttrs: map[string]bool{
			constants.ATTR_TAGS: true,
		},
		ResolveAttrs: map[string]bool{
			constants.ATTR_DEPENDENCIES: true,
			constants.ATTR_DATA:         true,
		},
	},
	constants.RULE_TS_PROJECT: {
		MatchAny: false,
		NonEmptyAttrs: map[string]bool{
			constants.ATTR_SOURCES: true,
		},
		MergeableAttrs: map[string]bool{
			constants.ATTR_SOURCES: true,
			constants.ATTR_TAGS:    true,
		},
		ResolveAttrs: map[string]bool{
			constants.ATTR_DEPENDENCIES: true,
			constants.ATTR_DATA:         true,
		},
	},
}
