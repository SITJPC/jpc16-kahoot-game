package function

import (
	"github.com/xuri/excelize/v2"
	"kahoot/types/payload"
	"log"
)

func OpenAndGetDataFromExcel(filePath string) []*payload.Score {
	// * Open Excel file
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
	mapScores, err := MapTableToScoreStruct(scores)
	if err != nil {
		log.Fatal("Unable to map excel to score struct", err)
	}
	return mapScores
}
