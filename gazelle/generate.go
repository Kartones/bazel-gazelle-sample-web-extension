// Copyright 2019 The Bazel Authors. All rights reserved.
// Modifications copyright (C) 2021 BenchSci Analytics Inc.
// Modifications copyright (C) 2018 Ecosia GmbH

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

// http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package web

import (
	"context"
	"log"
	"path"
	"sort"

	extensionConfig "github.com/kartones/bazel-gazelle-sample-web-extension/gazelle/config"
	"github.com/kartones/bazel-gazelle-sample-web-extension/gazelle/constants"
	"github.com/kartones/bazel-gazelle-sample-web-extension/gazelle/file"
	"github.com/kartones/bazel-gazelle-sample-web-extension/gazelle/generate"

	"github.com/bazelbuild/bazel-gazelle/language"
	"github.com/bazelbuild/bazel-gazelle/rule"
)

// Customize to destination repository rule definition sub-paths
var jsRules = rule.LoadInfo{
	Name:    constants.RULES_BASE_PATH + "js:defs.bzl",
	Symbols: []string{constants.RULE_JS_LIBRARY, constants.RULE_JS_TEST},
}
var tsRules = rule.LoadInfo{
	Name:    constants.RULES_BASE_PATH + "ts:defs.bzl",
	Symbols: []string{constants.RULE_TS_PROJECT},
}

// Loads returns .bzl files and symbols they define
func (lang *Web) Loads() []rule.LoadInfo {
	return []rule.LoadInfo{
		jsRules,
		tsRules,
	}
}

// GenerateRules extracts build metadata from source files in a directory.
// Called in each directory where an update is requested in depth-first post-order.
func (lang *Web) GenerateRules(args language.GenerateArgs) language.GenerateResult {
	webConfigs := args.Config.Exts[languageName].(extensionConfig.WebConfigs)
	webConfig := webConfigs[args.Rel]

	if !webConfig.ExtensionEnabled {
		return language.GenerateResult{}
	}

	var ruleImports []*extensionConfig.Imports
	var rules []*rule.Rule

	identifiedSources := generate.IdentifiedSources{}
	identifiedImports := generate.IdentifiedImports{}
	generatedRules := make([]*rule.Rule, 0)
	generatedImports := make([]*extensionConfig.Imports, 0)

	processRegularFiles(args, webConfig, &identifiedSources, &identifiedImports)

	ruleImports, rules = generate.GenerateAllRules(generate.ENABLED_RULES, generate.RuleArgs{
		Sources: identifiedSources,
		Imports: identifiedImports,
		Rel:     args.Rel,
	}, webConfig)

	for index := range rules {
		generatedRules = append(generatedRules, rules[index])
		generatedImports = append(generatedImports, ruleImports[index])
	}

	generatedImportsSlice := make([]interface{}, len(generatedImports))
	for index, importsSet := range generatedImports {
		generatedImportsSlice[index] = importsSet
	}

	return language.GenerateResult{
		Gen:     generatedRules,
		Empty:   nil,
		Imports: generatedImportsSlice,
	}
}

func processRegularFiles(args language.GenerateArgs, webConfig *extensionConfig.WebConfig,
	identifiedSources *generate.IdentifiedSources, identifiedImports *generate.IdentifiedImports) {

	for _, baseName := range args.RegularFiles {
		filePath := path.Join(args.Dir, baseName)

		if len(file.JSTestExtensionsPattern.FindStringSubmatch(baseName)) > 0 {
			identifiedSources.JavascriptTests = append(identifiedSources.JavascriptTests, baseName)
			identifiedImports.JavascriptTests = append(identifiedImports.JavascriptTests, *extensionConfig.ReadFileAndParse(filePath))
			continue
		}

		if len(file.JSExtensionsPattern.FindStringSubmatch(baseName)) > 0 {
			identifiedSources.Javascript = append(identifiedSources.Javascript, baseName)
			identifiedImports.Javascript = append(identifiedImports.Javascript, *extensionConfig.ReadFileAndParse(filePath))
			continue
		}

		if len(file.TSTestExtensionsPattern.FindStringSubmatch(baseName)) > 0 {
			identifiedSources.TypescriptTests = append(identifiedSources.TypescriptTests, baseName)
			identifiedImports.TypescriptTests = append(identifiedImports.TypescriptTests, *extensionConfig.ReadFileAndParse(filePath))
			continue
		}

		if len(file.TSExtensionsPattern.FindStringSubmatch(baseName)) > 0 {
			identifiedSources.Typescript = append(identifiedSources.Typescript, baseName)
			identifiedImports.Typescript = append(identifiedImports.Typescript, *extensionConfig.ReadFileAndParse(filePath))
			continue
		}
	}

	// Ensure determinism
	sort.Strings(identifiedSources.Javascript)
	sort.Strings(identifiedSources.Typescript)
	sort.Strings(identifiedSources.JavascriptTests)
	sort.Strings(identifiedSources.TypescriptTests)
}

// Called after initializing all extensions, but before visiting any package
func (language *Web) Before(ctx context.Context) {
	if language.Verbose {
		log.Println("Done setting up the extension")
	}
}

// Called when finished generating all rules in all packages with all extensions, but before resolving imports
func (language *Web) DoneGeneratingRules() {
	if language.Verbose {
		log.Println("Done generating rules")
	}
}

// Called when finished resolving all dependencies with all extensions
func (language *Web) AfterResolvingDeps(ctx context.Context) {
	if language.Verbose {
		log.Println("Done resolving dependencies")
	}
}
