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

### caching

A cache can be started to help with processing and performance, though it is not necessary for the functionality.

The mongodb cache can be started with:
```bash
docker-compose up -d
```
!!Adjust the config.json accordingly if cache is activated (URI, Username and Password).

### configuration

The config.template.json is prefilled with configuration regarding tweeking of evaluation inside of _deprec_.

To adjust are following configurations:
- Extraction
  - GitHub -> GitHub API Token
  - OSSIndex -> Username and Token

## local build

```bash
# first time
sudo apt install goreleaser

# first time
goreleaser init

goreleaser release --snapshot --clean
```
