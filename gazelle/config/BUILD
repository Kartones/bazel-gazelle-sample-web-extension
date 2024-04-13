load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "config",
    srcs = [
        "directives.go",
        "imports.go",
        "webConfig.go",
        "kinds.go",
    ],
    importpath = "github.com/kartones/bazel-gazelle-sample-web-extension/gazelle/config",
    visibility = ["//visibility:public"],
    deps = [
        "//gazelle/constants",
        "//gazelle/file",
        "//gazelle/parse",
        "//gazelle/utils",
        "@bazel_gazelle//rule:go_default_library",
        "@com_github_bazelbuild_buildtools//build",
    ],
)
