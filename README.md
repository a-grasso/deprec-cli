# deprec-cli
CLI tool for using [deprec](https://github.com/a-grasso/deprec)

# Table Of Contents

1. [Installation](#installation)
2. [Quick Start](#quick-start)
3. [Usage](#usage)
    1. [Caching](#caching)
4. [Local Build](#local-build)


## installation

```bash
wget -q https://github.com/a-grasso/deprec-cli/releases/latest/download/deprec-cli_Linux_x86_64
chmod +x deprec-cli_Linux_x86_64 # maybe necessary
```

## quick start

Trying out deprec-cli/deprec, you can run deprec-cli on the sbom created from this very project, one of the supplied ones or one of your own. Steps below run deprec-cli on a sbom created from this very project:

```bash
## requires Go 1.18 or newer!!!

## install cyclonedx-gomod to create a sbom from golang projects
go install github.com/CycloneDX/cyclonedx-gomod@v1.0.0

## clone this project
git clone https://github.com/a-grasso/deprec-cli.git
cd deprec-cli

## create a sbom for this project with cyclonedx-gomod
cyclonedx-gomod app > sbom.json

## install deprec-cli
wget -q https://github.com/a-grasso/deprec-cli/releases/latest/download/deprec-cli_Linux_x86_64
chmod +x deprec-cli_Linux_x86_64

## copy the supplied configuration template file
cp config-templates/template.config.json config.json

## copy the supplied env variable template file
## we leave the deprec configuration file as is for now, see further below for deprec configuration possibilities
cp config-templates/template.env .env

##  fill out env variables as good as possible 
##  most important for now is that you at least provide a GitHub token
##  you can ignore all others for now, see further below for more information environment variables/caching

## run deprec-cli
./deprec-cli_Linux_x86_64 sbom.json
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
        Output file (default "deprec-output.csv")                                    
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

The config-templates/template.config.json is prefilled with configuration regarding tweaking of evaluation inside _deprec_.

### configuration - environment variables
The provided config-templates/template.env file has environment variables that would contain sensitive information listed. These configurations should be passed alongside rather than written down in a supplied configuration file.

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
