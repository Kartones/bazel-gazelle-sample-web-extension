.DEFAULT_GOAL := run

GAZELLE_ARGS ?=

run:
	bazel run //:gazelle $(GAZELLE_ARGS)
.PHONY: run

test:
	bazel test //...
.PHONY: test
