# deprec-cli
CLI tool for using deprec

## installation

```bash
wget -q https://github.com/a-grasso/deprec-cli/releases/latest/download/deprec-cli_Linux_x86_64
chmod +x deprec-cli_Linux_x86_64 # maybe necessary
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

### configuration - evaluation

The template.config.json is prefilled with configuration regarding tweaking of evaluation inside _deprec_.

### configuration - environment variables
The provided template.env file has environment variables that would contain sensitive information listed. These configurations should be passed alongside rather than written down in a configuration file.

To adjust are following configurations:
- Extraction
  - GitHub -> _GitHub API Token_
  - OSSIndex -> _Username_ and _Token_
- Cache (see above)
  - _URI_
  - _Username_
  - _Password_

## local build

```bash
# first time
sudo apt install goreleaser

# first time
goreleaser init

goreleaser release --snapshot --clean
```
