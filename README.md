# deprec-cli
CLI tool for using deprec

## installation

```bash
wget https://github.com/a-grasso/deprec-cli/releases/latest/download/deprec-cli_Linux_x86_64
```

## local build

```bash
# first time
sudo apt install goreleaser

# first time
goreleaser init

goreleaser release --snapshot --clean
```
