// https://hackercollider.com/articles/2016/07/06/create-your-own-shell-in-python-part-2/

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var cmd string
	var err error

	for {
		fmt.Print(">> ")
		cmd, err = reader.ReadString('\n')

		if err != nil {
			log.Println(":(")
		}

		cmds := Tokenize(cmd)
		fmt.Print(cmds)
	}
}

func Tokenize(cmd string) []string {
	return strings.Split(cmd, " ")
}
