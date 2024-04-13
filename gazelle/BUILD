# Copyright 2019 The Bazel Authors. All rights reserved.
# Modifications copyright (C) 2021 BenchSci Analytics Inc.
# Modifications copyright (C) 2018 Ecosia GmbH

# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at

# http://www.apache.org/licenses/LICENSE-2.0

# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

load("@bazel_gazelle//:def.bzl", "gazelle_generation_test")
load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "gazelle",
    srcs = [
        "configure.go",
        "generate.go",
        "kinds.go",
        "lang.go",
        "resolve.go",
    ],
    importpath = "github.com/kartones/bazel-gazelle-sample-web-extension/gazelle",
    visibility = ["//visibility:public"],
    deps = [
        "//gazelle/config",
        "//gazelle/constants",
        "//gazelle/file",
        "//gazelle/generate",
        "//gazelle/parse",
        "//gazelle/utils",
        "@bazel_gazelle//config:go_default_library",
        "@bazel_gazelle//label:go_default_library",
        "@bazel_gazelle//language:go_default_library",
        "@bazel_gazelle//merger:go_default_library",
        "@bazel_gazelle//repo:go_default_library",
        "@bazel_gazelle//resolve:go_default_library",
        "@bazel_gazelle//rule:go_default_library",
        "@com_github_bazelbuild_buildtools//build",
    ],
)

[gazelle_generation_test(
    name = "{}_test".format(
        file[len("tests/"):-len("/WORKSPACE")],
    ),
    size = "small",
    gazelle_binary = "//:gazelle_bin",
    test_data = glob(
        include = [file[0:-len("/WORKSPACE")] + "/**"],
    ),
) for file in glob(
    include = ["**/WORKSPACE"],
)]
