package main

import (
	"flag"
	"fmt"
	"go_playground/mini-project/quiz-app/processor"
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
	problems, err := processor.ProblemPuller(*fName)

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

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
