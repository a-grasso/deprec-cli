package main

import (
	"bytes"
	"encoding/csv"
	"errors"
	"flag"
	"fmt"
	cdx "github.com/CycloneDX/cyclonedx-go"
	"github.com/a-grasso/deprec"
	"github.com/a-grasso/deprec/configuration"
	"github.com/a-grasso/deprec/logging"
	"github.com/a-grasso/deprec/model"
	"os"
	"strings"
)

type CommandLineInput struct {
	sbomPath   string
	configPath string
	envPath    string
	numWorkers int
	runMode    string
	outputFile string
}

func main() {

	logging.Logger.Info("DepRec run started...")

	flag.Usage = func() {
		fmt.Printf("Usage: %s [options] <sbomJson>\nOptions:\n", os.Args[0])
		flag.PrintDefaults()
	}

	input, err := getInput()

	if err != nil {
		exitGracefully(err)
	}

	config, err := configuration.Load(input.configPath, input.envPath)
	if err != nil {
		exitGracefully(err)
	}

	cdxBom, err := decodeSBOM(input.sbomPath)
	if err != nil {
		exitGracefully(err)
	}

	deprecClient := deprec.NewClient(*config)

	runConfig := deprec.RunConfig{
		Mode:       deprec.RunMode(input.runMode),
		NumWorkers: input.numWorkers,
	}

	result := deprecClient.Run(cdxBom, runConfig)

	writeToOutputFile(input.outputFile, result)
}

func writeToOutputFile(outputFile string, result *deprec.Result) {

	csvFile, err := os.Create(outputFile)
	if err != nil {
		logging.SugaredLogger.Errorf("could not create outputfile %s: %s", outputFile, err)
	}
	defer csvFile.Close()

	csvWriter := csv.NewWriter(csvFile)
	var records = [][]string{
		{
			"Dependency Name",
			"Dependency Version",
			"Dependency Package URL",
			"Extracted Data Sources",
			"Used EOS Factors",
			"Recommendation Distribution",
			"Decision Making",
			"Watchlist",
			"No Immediate Action",
			"No Concerns",
		},
	}

	for _, r := range result.Results {
		record := []string{
			fmt.Sprintf("%s", r.Dependency.Name),
			fmt.Sprintf("%s", r.Dependency.Version),
			fmt.Sprintf("%s", r.Dependency.PackageURL),
			fmt.Sprintf("%s", r.DataSources),
			fmt.Sprintf("%s", r.UsedFirstLevelCores()),
			fmt.Sprintf("%s", r.RecommendationsInsights()),
			fmt.Sprintf("%.4f", r.Recommendations[model.DecisionMaking]),
			fmt.Sprintf("%.4f", r.Recommendations[model.Watchlist]),
			fmt.Sprintf("%.4f", r.Recommendations[model.NoImmediateAction]),
			fmt.Sprintf("%.4f", r.Recommendations[model.NoConcerns]),
		}

		records = append(records, record)
	}

	for _, record := range records {
		err = csvWriter.Write(record)
		if err != nil {
			logging.SugaredLogger.Errorf("could not write to outputfile %s: %s", outputFile, err)
		}
	}

	csvWriter.Flush()
}

func getInput() (*CommandLineInput, error) {
	if len(os.Args) < 2 {
		return &CommandLineInput{}, errors.New("cli argument error: SBOM file argument required")
	}

	config := flag.String("config", "config.json", "Evaluation config file")
	env := flag.String("env", ".env", "Environment variables file")
	output := flag.String("output", "deprec-output.csv", "Output file")
	workers := flag.Int("workers", 5, "Number of workers if in parallel mode")
	runMode := flag.String("runMode", "parallel", "Run mode - parallel or linear")

	flag.Parse()

	sbom := flag.Arg(0)

	return &CommandLineInput{sbom, *config, *env, *workers, *runMode, *output}, nil
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
		return nil
	}

	result := make(map[model.ExternalReference]string)

	for _, reference := range *references {
		result[model.ExternalReference(reference.Type)] = reference.URL
	}

	return result
}
