load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "utils",
    srcs = [
        "colors.go",
        "rules.go",
        "utils.go",
    ],
    importpath = "github.com/kartones/bazel-gazelle-sample-web-extension/gazelle/utils",
    visibility = ["//visibility:public"],
    deps = [
        "//gazelle/constants",
        "@bazel_gazelle//config:go_default_library",
        "@bazel_gazelle//rule:go_default_library",
    ],
)
