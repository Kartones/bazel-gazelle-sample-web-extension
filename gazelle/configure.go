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
	"log"
	"strconv"

	extensionConfig "github.com/kartones/bazel-gazelle-sample-web-extension/gazelle/config"
	"github.com/kartones/bazel-gazelle-sample-web-extension/gazelle/constants"
	extensionUtils "github.com/kartones/bazel-gazelle-sample-web-extension/gazelle/utils"

	"github.com/bazelbuild/bazel-gazelle/config"
	"github.com/bazelbuild/bazel-gazelle/rule"
)

// KnownDirectives returns a list of directive keys that this Configurer can
// interpret. Gazelle prints errors for directives that are not recoginized by
// any Configurer.
func (*Web) KnownDirectives() []string {
	return extensionConfig.KnownDirectives
}

// Read configuration from current directory's BUILD file, and apply changes from inherited configuration.
// Called in each directory when first entering it.
// c is the configuration for the current directory. It starts out as a copy of the configuration
// for the parent directory.
//
// rel is the slash-separated relative path from the repository root to the current directory.
// It is "" for the root directory itself.
//
// f is the build file for the current directory or nil if there is no existing build file.
func (language *Web) Configure(config *config.Config, rel string, file *rule.File) {
	if _, exists := config.Exts[languageName]; !exists {
		config.Exts[languageName] = createWithRootConfig()
	}

	webConfigs := config.Exts[languageName].(extensionConfig.WebConfigs)

	webConfig, exists := webConfigs[rel]
	if !exists {
		parent := webConfigs.ParentForPackage(rel)
		webConfig = parent.NewChild()
		webConfigs[rel] = webConfig
	}

	webConfig.Verbose = webConfig.Verbose || language.Verbose

	if file != nil {
		for _, directive := range file.Directives {

			switch directive.Key {

			case constants.DIRECTIVE_EXTENSION_ENABLED:
				webConfig.ExtensionEnabled = readBoolDirective(directive)

			case constants.DIRECTIVE_EXTENSION_QUIET:
				webConfig.Quiet = readBoolDirective(directive) && !language.Verbose
				if webConfig.Quiet {
					webConfig.Verbose = false
				}

			case constants.DIRECTIVE_EXTENSION_VERBOSE:
				webConfig.Verbose = readBoolDirective(directive) || language.Verbose
				if webConfig.Verbose {
					webConfig.Quiet = false
				}

			}
		}
	}

	if webConfig.Verbose {
		log.Printf("Configure() - '%s'", rel)
	}

	if rel == "" && webConfig.ExtensionEnabled {
		if webConfig.Verbose {
			log.Println("Root Configuration initialized")
		}
	}

}

func createWithRootConfig() extensionConfig.WebConfigs {
	rootConfig := extensionConfig.NewWebConfig()
	return extensionConfig.WebConfigs{
		"": rootConfig,
	}
}

func readBoolDirective(directive rule.Directive) bool {
	if directive.Value == "" {
		return true
	} else {
		val, err := strconv.ParseBool(directive.Value)
		if err != nil {
			log.Fatalf(extensionUtils.Err("failed to read directive %s: %v", directive.Key, err))
		}
		return val
	}
}
