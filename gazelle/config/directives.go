package config

import "github.com/kartones/bazel-gazelle-sample-web-extension/gazelle/constants"

var KnownDirectives = []string{
	constants.DIRECTIVE_EXTENSION_ENABLED,
	constants.DIRECTIVE_EXTENSION_VERBOSE,
	constants.DIRECTIVE_EXTENSION_QUIET,
}
