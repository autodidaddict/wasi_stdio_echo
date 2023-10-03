package main

import (
	"bufio"
	"fmt"
	"os"
)

func init() {
	fmt.Println("> ")
	reader := bufio.NewReader(os.Stdin)
	for {
		text, _ := reader.ReadString('\n')
		fmt.Println("MODULE: " + text)
	}
}

func main() {
}
