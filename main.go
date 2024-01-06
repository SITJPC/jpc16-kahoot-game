package main

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"kahoot/utils"
	"log"
	"os"
)

func main() {
	filePath := os.Args[1]

	// * Parse flag
	//fs := flag.NewFlagSet("kahoot", flag.ContinueOnError)

	//var count int
	//fs.IntVar(&count, "n", 5, "number of lines to read from the file")
	//args := os.Args[2:]
	//fmt.Println(args)
	//_ = fs.Parse(args)
	//
	//fmt.Printf("Count: %d\n", count)
	// * Open Excel file

	// * Read list participants

	// * Read Kahoot report
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		log.Fatal("Unable to open excel file", err)
	}
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			log.Fatal("Unable to close file", err)
		}
	}()

	// * Access sheet name "Final Scores"
	scores, err := f.GetRows("Final Scores")
	if err != nil {
		log.Fatal("Unable to open sheet Final-Scores")
	}

	// * Map struct
	mapScores, err := utils.MapTableToScoreStruct(scores)
	if err != nil {
		log.Fatal("Unable to map excel to score struct", err)
	}

	for _, score := range mapScores {
		fmt.Println(*score.TotalScore)
	}

}
