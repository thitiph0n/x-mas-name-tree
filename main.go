package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Generate your X'mas name tree")
	fmt.Println("-----------------------------")

	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	printTree(name)

}

func printTree(name string) {
	fmt.Println("")
	length := len(name)
	space := int(length/2) - 1

	// Leafs
	for i := 0; i < length; i += 2 {
		var sb strings.Builder
		for i := 0; i <= space; i++ {
			sb.WriteString(" ")
		}
		sb.WriteString(name[0 : i+1])
		space--
		fmt.Println(sb.String())
	}

	// Trunk
	var sb strings.Builder
	for i := 0; i <= int(length/2)-2; i++ {
		sb.WriteString(" ")
	}
	sb.WriteString("|||")
	fmt.Println(sb.String())

}
