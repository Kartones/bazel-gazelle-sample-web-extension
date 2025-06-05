.DEFAULT_GOAL := run

GAZELLE_ARGS ?=

run:
	bazelisk run //:gazelle $(GAZELLE_ARGS)
.PHONY: run

test:
	bazelisk test //...
.PHONY: test
