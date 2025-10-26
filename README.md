# generate-action-readme

`gar` (generate-action-readme) is a simple CLI tool to convert a YAML specification for a GitHub Action into a well-formatted README written in GitHub-flavored markdown.

## Installation

In order to install `gar` there are three ways:

1. Using `go`: if you already have `go` 1.22+ installed in your environment, installing `gar` is effortless

```bash
go install github.com/AstraBert/generate-action-readme
```

2. Using `npm` (recommended):

```bash
npm install @cle-does-things/gar
```

3. Downloading the executable from the [releases page](https://github.com/AstraBert/generate-action-readme/releases): you can download it directly from the GitHub repository or, if you do not want to leave your terminal, you can use `curl`:

```bash
curl -L -o gar https://github.com/AstraBert/generate-action-readme/releases/download/<version>/generate-action-readme_<version>_<OS>_<processor>.tar.gz ## e.g. https://github.com/AstraBert/generate-action-readme/releases/download/0.1.1/generate-action-readme_0.1.1_darwin_amd64.tar.gz

# make sure the downloaded binary is executable (not needed for Windows)
chmod +x gar
```

In this last case, be careful to specify your OS (supported: linux, windows, macos) and your processor type (supported: amd, arm).

## Usage

`gar` has one command: `generate` (aliased also to `g` and `gen`).

You can use it with two optional flags:

- `-a`, `--action`: path to the YAML file with the action specification
- `-r`, `--readme`: path to the README file to write

### Examples

- With a specific path to action and README:

```bash
gar generate --action action/action.yml --readme action/README.md
```

- With default paths for action (`action.yml`) and README (`README.md`)

```bash
gar generate
```

- Using command aliases:

```bash
gar g --action v1/action.yml --readme v1/README.md
gar gen
```

## Contributing

We welcome contributions! Please read our [Contributing Guide](./CONTRIBUTING.md) to get started.

## License

This project is licensed under the [MIT License](./LICENSE)
