
Example Bazel [Gazelle](https://github.com/bazelbuild/bazel-gazelle) extension tho manage BUILD file contents for Javascript and Typescript code. 

Based on [rules_nodejs_gazelle](https://github.com/benchsci/rules_nodejs_gazelle).

Intended to generate some of the rules used at my [bazel-web-template repository](https://github.com/Kartones/bazel-web-template).


⚠️ The extension is still WIP. It is missing some basic features, like import parsing and resolution ⚠️


## Setup / Requirements

- [Bazelisk](https://github.com/bazelbuild/bazelisk)


## Run

```bash
bazelisk run //:gazelle

# apply only to a specific path:
bazelisk run //:gazelle path/to/folder
```


## Development

### Build

```bash
bazelisk build //...
```


### Recommendations

- [Bazel plugin for Visual Studio Code](https://marketplace.visualstudio.com/items?itemName=BazelBuild.vscode-bazel)
- [buildifier](https://github.com/bazelbuild/buildtools)


## Testing

```bash
bazelisk test //...
```


## Updates

Apart from updating `WORKSPACE` and `MODULE.bazel`, you should run:

```bash
bazelisk run //:update_go_deps
```
