# Simple quiz game

My answer to this Gophercise: https://github.com/gophercises/quiz

Quiz-game is a simple program that reads a quiz provided via a CSV file and then will give the quiz to the user. 
The quiz is timed and will end either when the timer runs out or the user answers all questions, whichever happens 
first.

## How to play

```
$ cd quiz-game
$ go build quiz.go
$ ./quiz
```

## Flags
The quiz file, time limit, and shuffled questions are customizable via flags.

```
$ ./quiz -h

Usage of ./quiz:
  -csv string
        a csv file in the format of question,answer (default "problems.csv")
  -limit int
        the time limit for the quiz in seconds (default 10)
  -shuffle
        true/false if the questions will be shuffled
```