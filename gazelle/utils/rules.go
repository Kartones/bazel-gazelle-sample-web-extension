package js

import (
	"github.com/bazelbuild/bazel-gazelle/config"
	"github.com/bazelbuild/bazel-gazelle/rule"
)

func GetKind(c *config.Config, kindName string) string {
	// Extract kind_name from KindMap
	if kind, ok := c.KindMap[kindName]; ok {
		return kind.KindName
	}
	return kindName
}

// Version of rule.Rule.AttrStrings that never returns nil (returns an empty slice instead)
func GetRuleAttributeStrings(rule *rule.Rule, attributeName string) []string {
	attributeValue := rule.AttrStrings(attributeName)
	if attributeValue == nil {
		return []string{}
	}
	return attributeValue
}
