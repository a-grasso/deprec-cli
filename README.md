# deprec-cli
CLI tool for using deprec

## installation

```bash
wget https://github.com/a-grasso/deprec-cli/releases/latest/download/deprec-cli_Linux_x86_64
```

## usage

```bash
./deprec-cli <sbom-path> <config-path> <output-file-path> <number of workers> <runMode>
```

- sbom-path = path to sbom file
- config-path = path to .json file (see config.template.json)
- output-file-path = path to file to be written to
- OPTIONAL number of workers = number of concurrent dependencies to be processed
- OPTIONAL runMode = "linear" OR "parellel" processing of dependencies

## local build

```bash
# first time
sudo apt install goreleaser

# first time
goreleaser init

goreleaser release --snapshot --clean
```
