package main

import (
	"bufio"
	"fmt"
	"os"
)

type Card struct {
	question string
	answer   string
	category string
}

func main() {
	fileNames := []string{"vocab.txt", "grammar.txt"}
	ctg := []string{"v", "g"}
	modes := []string{"[1]Quiz", "[2]Review", "[3]Add Questions"}

	fmt.Println("Lets learn japanese!")
	fmt.Println("Select a mode of operation:")
	fmt.Println(modes)
	if mode := userSelect(len(modes)); mode >= 0 {
		fmt.Println(mode)
	}

	mode := userSelect(len(modes))
	fmt.Println(mode)
	var questions *os.File
	fmt.Println("Select a ctg [v]Vocab | [g]Grammar")
	if toggle := userSelect(len(ctg)); toggle >= 0 && toggle < 3 {
		questions, _ = os.Open(fileNames[toggle])
		fmt.Println("opening " + fileNames[toggle])
	}

	defer questions.Close()
	_ = questions

	z := Card{"What is the meaining of Wa?",
		"The particle “wa” tells us that the word or" +
			"phrase before it is the topic of that sentence.", ctg[1]}
	//println(z.category)
	_ = z
}

//Functions below
func makeCard(q, a, ctg string) Card {
	return Card{q, a, ctg}
}

func userSelect(optCount int) int {
	input := bufio.NewReader(os.Stdin)
	s, _ := input.ReadByte()

	fmt.Println(int(s))

	fmt.Println(byte(optCount))

	if s > byte(optCount) {
		return -1
	}
	return int(s)
}
