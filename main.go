package main

import (
	"fmt"
	"strings"
	"github.com/cucumber/gherkin-go"
	"path/filepath"
	"log"
	"io/ioutil"
	"flag"
)

var featuresDir = flag.String("features", "./features", "Features directory")

func main() {
	flag.Parse()

	featureDirPattern := filepath.Join(*featuresDir,"*.feature")

	featureFilenames, err := filepath.Glob(featureDirPattern)
	if err != nil {
		log.Println(err)
	}

	log.Println("Detected the following files: ", featureFilenames)

	for _, featureFilename := range (featureFilenames) {
		input, err := ioutil.ReadFile(featureFilename)
		if err != nil {
			log.Fatal("Error: ", err.Error())

		}

		r := strings.NewReader(string(input))

		gherkinDocument, err := gherkin.ParseGherkinDocument(r)
		if err != nil {
			log.Fatal("Error: ", err.Error())
		}

		feature := gherkinDocument.Feature
		fmt.Println()

		for _, featureChild := range (feature.Children) {
			switch featureChild.(type) {
			case *gherkin.Scenario:
				scenario, _ := featureChild.(*gherkin.Scenario)
				fmt.Println(feature.Name, ", ", scenario.Name, ", ", isAutomated(scenario.Tags))
				for _, step := range (scenario.Steps) {
					switch step.Argument.(type) {
					case *gherkin.DataTable:
						dataTable, _ := step.Argument.(*gherkin.DataTable)
						for _, dataTableRow := range (dataTable.Rows) {
							fmt.Print(feature.Name, ", Table: ")
							for _, dataTableCell := range (dataTableRow.Cells) {
								fmt.Print(dataTableCell.Value)
								fmt.Print(" ")
							}
							fmt.Printf(",%s\n", isAutomated(scenario.Tags))
						}
					}
				}
			case *gherkin.ScenarioOutline:
				scenarioOutline, _ := featureChild.(*gherkin.ScenarioOutline)
				fmt.Println()
				for _, examples := range (scenarioOutline.Examples) {
					for count, tableBody := range (examples.TableBody) {
						if count == 0 {
							continue
						}
						fmt.Print(feature.Name, ", ", scenarioOutline.Name, " Example: ")
						for _, cells := range (tableBody.Cells) {
							fmt.Print(cells.Value)
							fmt.Print(" ")
						}
						fmt.Printf(",%s\n", isAutomated(scenarioOutline.Tags))
					}
				}
			}
		}
		fmt.Printf("\n")
	}
}

func isAutomated(tags []*gherkin.Tag) string {
	for _, tag := range tags {
		if tag.Name == "@wip" {
			return "MANUAL"
		}
	}
	return "AUTOMATED"
}
