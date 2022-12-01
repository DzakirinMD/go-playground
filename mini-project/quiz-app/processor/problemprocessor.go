package processor

import (
	"encoding/csv"
	"fmt"
	"go_playground/mini-project/quiz-app/quizstructs"
	"os"
)

// read all the problems from the quiz.csv
func ProblemPuller(fileName string) ([]quizstructs.Problem, error) {

	fmt.Println("file name is :", fileName)
	// 1. open the file
	if fileObject, err := os.Open(fileName); err == nil {
		// 2. create a new reader
		csvReader := csv.NewReader(fileObject)

		// 3. read the file
		if csvLines, err := csvReader.ReadAll(); err == nil {
			// 4. call the parseProblem function
			return parseProblem(csvLines), nil
		} else {
			return nil, fmt.Errorf("error in reading data in csv"+"format from %s file: %s", fileName, err.Error())
		}
	} else {
		return nil, fmt.Errorf("error in opening %s file: %s", fileName, err.Error())
	}
}

func parseProblem(lines [][]string) []quizstructs.Problem {
	// go over line in csv and parse them, with problem struct
	sliceOfProblem := make([]quizstructs.Problem, len(lines))

	for i := 0; i < len(lines); i++ {
		// take before comma as question and after comma as answer. the position is determined by before or after the comma
		sliceOfProblem[i] = quizstructs.Problem{Question: lines[i][0], Answer: lines[i][1]}
	}
	return sliceOfProblem
}
