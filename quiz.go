package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type problem struct {
	question string
	answer   string
}

func main() {
	fileName, timeLimit, shuffle := readArguments()

	fmt.Println("Welcome to the Quiz Game!")
	fmt.Println("-------------------------")

	questions, err := loadQuiz(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	if shuffle {
		shuffleQuiz(questions)
	}

	fmt.Println("Answer all the questions before the time is up! The timer will start once you press Enter")
	fmt.Scanln()

	correct, err := playQuiz(questions, timeLimit)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("You scored %d out of %d \n", correct, len(questions))
}

func readArguments() (string, int, bool) {
	fileName := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	timeLimit := flag.Int("limit", 10, "the time limit for the quiz in seconds")
	shuffle := flag.Bool("shuffle", false, "true/false if the questions will be shuffled")
	flag.Parse()

	return *fileName, *timeLimit, *shuffle
}

func loadQuiz(fileName string) ([]problem, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		return nil, err
	}

	problems := make([]problem, len(records))
	for i, record := range records {
		problem := problem{
			question: record[0],
			answer:   strings.TrimSpace(record[1]),
		}
		problems[i] = problem
	}

	return problems, err
}

func shuffleQuiz(problems []problem) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(problems), func(i, j int) {
		problems[i], problems[j] = problems[j], problems[i]
	})
}

func playQuiz(problems []problem, timeLimit int) (int, error) {
	timer := time.NewTimer(time.Duration(timeLimit) * time.Second)

	correct := 0
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s ", i+1, p.question)

		responseCh := make(chan string)
		go func() {
			var response string
			fmt.Scanln(&response)
			response = strings.TrimSpace(response)
			responseCh <- response
		}()

		select {
		case <-timer.C:
			fmt.Println("\nTime's up!")
			return correct, nil
		case response := <-responseCh:
			if response == p.answer {
				correct += 1
			} else {
				fmt.Println("Wrong! The correct answer is " + p.answer)
			}
		}
	}

	return correct, nil
}
