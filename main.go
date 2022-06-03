package main

import "fmt"

func main() {
	var name string
	var score int

	fmt.Println("Welcome to the quiz game!")
	fmt.Println("Enter your name to continue: ")
	fmt.Scan(&name)

	fmt.Printf("\nWelcome %v!\n", name)
	fmt.Println("Answer the following questions to win the game!")

	var answer1 string
	fmt.Println("\n1). Which one of these is the oldest language?")
	fmt.Println("C++ or Python or Java?")
	fmt.Println("Answer: ")
	fmt.Scan(&answer1)

	if answer1 == "C++" {
		fmt.Println("+1 Correct answer!")
		score++
	} else {
		fmt.Println("Aw, wrong answer! C++ is the correct answer.")
	}

	var answer2 string
	fmt.Println("\n2). When was C++ developed?")
	fmt.Println("1972 or 1979 or 1985?")
	fmt.Println("Answer: ")
	fmt.Scan(&answer2)

	if answer2 == "1979" {
		fmt.Println("+1 Correct answer!")
		score++
	} else {
		fmt.Println("Aw, nice try! C++ was invented in 1972.")
	}
	
	var answer3 string
	fmt.Println("\n3). Which one of these has the slowest compilation time?")
	fmt.Println("Java or Python or C++")
	fmt.Println("Answer: ")
	fmt.Scan(&answer3)

	if answer3 == "Python" {
		fmt.Println("+1 Correct answer!")
		score++
	} else {
		fmt.Println("Aw, sad reacts only! Python is the slowest because it's an interpreted language!")
	}

	fmt.Printf("\nYour score is %v.\n", score)
	fmt.Println("Goodbye, have a nice day! :)")
}