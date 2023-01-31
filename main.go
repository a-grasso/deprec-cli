package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	cdx "github.com/CycloneDX/cyclonedx-go"
	"github.com/a-grasso/deprec"
	"github.com/a-grasso/deprec/configuration"
	"github.com/a-grasso/deprec/logging"
	"github.com/a-grasso/deprec/model"
	"os"
	"strconv"
	"strings"
)

type CommandLineInput struct {
	sbomPath   string
	configPath string
	numWorkers int
}

func main() {

	logging.Logger.Info("DepRec run started...")

	flag.Usage = func() {
		fmt.Printf("Usage: %s <sbom> <config> <workers>\nOptions:\n none", os.Args[0])
	}

	input, err := getInput()

	if err != nil {
		exitGracefully(err)
	}

	config, err := configuration.Load(input.configPath)
	if err != nil {
		exitGracefully(err)
	}

	cdxBom, err := decodeSBOM(input.sbomPath)
	if err != nil {
		exitGracefully(err)
	}

	deprecClient := deprec.NewClient(config)

	runConfig := deprec.RunConfig{
		Mode: deprec.Linear,
		//Mode:       Parallel,
		NumWorkers: input.numWorkers,
	}

	deprecClient.Run(cdxBom, runConfig)
}

func getInput() (*CommandLineInput, error) {
	if len(os.Args) < 3 {
		return &CommandLineInput{}, errors.New("cli argument error: SBOM file and config path arguments required")
	}

	flag.Parse()

	sbom := flag.Arg(0)
	config := flag.Arg(1)
	workers, err := strconv.Atoi(flag.Arg(2))
	if err != nil {
		workers = 5
	}

	return &CommandLineInput{sbom, config, workers}, nil
}

func exitGracefully(err error) {
	logging.SugaredLogger.Fatalf("exited gracefully : %v\n", err)
}

func decodeSBOM(path string) (*cdx.BOM, error) {

	json, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("could not read sbom file '%s': %s", path, err)
	}
	reader := bytes.NewReader(json)

	bom := new(cdx.BOM)
	decoder := cdx.NewBOMDecoder(reader, cdx.BOMFileFormatJSON)
	if err = decoder.Decode(bom); err != nil {
		return nil, fmt.Errorf("could not decode SBOM: %s", err)
	}

	calcSBOMStats(bom)

	return bom, nil
}

func calcSBOMStats(bom *cdx.BOM) {
	noVCS := 0
	vcsGitHub := 0
	for _, component := range *bom.Components {
		if component.ExternalReferences == nil {
			noVCS += 1
			continue
		}

		externalReference := parseExternalReference(component)
		vcs, exists := externalReference["vcs"]

		if !exists {
			noVCS += 1
			continue
		}

		if strings.Contains(vcs, "github.com") {
			vcsGitHub += 1
		}
	}

	logging.SugaredLogger.Infof("%d/%d/%d github/vcs/total", vcsGitHub, len(*bom.Components)-noVCS, len(*bom.Components))
}

func parseExternalReference(component cdx.Component) map[model.ExternalReference]string {

	references := component.ExternalReferences

	if references == nil {
		logging.SugaredLogger.Infof("SBOM component '%s' has no external references", component.Name)
		return nil
	}

	result := make(map[model.ExternalReference]string)

	for _, reference := range *references {
		result[model.ExternalReference(reference.Type)] = reference.URL
	}

	return result
}
