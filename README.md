# deprec-cli
CLI tool for using deprec

## installation

```bash
wget https://github.com/a-grasso/deprec-cli/releases/latest/download/deprec-cli_Linux_x86_64
```

## usage

```bash
Usage: ./deprec-cli [options] <sbomJson>                                         
Options:                                                                             
  -config string                                                                     
        Evaluation config file (default "config.json")                               
  -env string                                                                        
        Environment variables file (default ".env")                                  
  -output string                                                                     
        Output file (default "deprec-output.txt")                                    
  -runMode string                                                                    
        Run mode - parallel or linear (default "parallel")                           
  -workers int                                                                       
        Number of workers if in parallel mode (default 5)    
```

### caching

A cache can be started to help with processing and performance, though it is not necessary for the functionality.

The mongodb cache can be started with:
```bash
docker-compose up -d
```
!!Adjust the .env accordingly if cache is activated (URI, Username and Password).

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
