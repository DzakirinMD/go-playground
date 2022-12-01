package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"go_playground/mini-project/quiz-app/quizstructs"
	"os"
	"time"
)

// read csv file
// timeout in 30 second

func main() {
	// 1. input the name of the file
	fName := flag.String("filepath", "mini-project/quiz-app/quiz.csv", "path of the csv file")
	// 2. set the duration of the timer
	timer := flag.Int("timer", 120, "timer for the quiz")
	flag.Parse()

	// 3. pull the problems from the file (calling our problem puller func)
	problems, err := problemPuller(*fName)

	// 4. handle error from pulling file
	if err != nil {
		exit(fmt.Sprintf("Something went wrong:%s", err.Error()))

	}
	// 5. create a variable to count our correct answer
	correctAns := 0

	// 6. using the duration of the timer, we want to initialize the timer
	ticker := time.NewTicker(time.Duration(*timer) * time.Second)
	answerChannel := make(chan string)

	// 7. loop thru problems, print the questions, accept the answers

	// label the loop as problemLoop
problemLoop:
	for index, value := range problems {
		var answer string

		fmt.Printf("Problem %d: %s=", index+1, value.Question)

		go func() {
			// linux new line - \n
			// windows new line - \r\n
			fmt.Scanf("%s\r\n", &answer)
			// add answer into channel
			answerChannel <- answer
		}()

		select {
		// when timer run out we break from the loop
		case <-ticker.C:
			fmt.Println("")
			break problemLoop
		// when answer is correct
		case iAns := <-answerChannel:
			if iAns == value.Answer {
				correctAns++
			}
			if index == len(problems)-1 {
				close(answerChannel)
			}
		}
	}
	// 8. we'll calculate and print out the result
	fmt.Printf("Your result is %d out of %d", correctAns, len(problems))
	fmt.Println("\nPress enter to exit")
	<-answerChannel
}

// read all the problems from the quiz.csv
func problemPuller(fileName string) ([]quizstructs.Problem, error) {

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

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
