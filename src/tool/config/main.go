package config

import (
	"bufio"
	"os"
	"strings"
)

type t struct {
	Name  string
	lv    int
	equip []int
}

var s1 *t
var s2 map[interface{}]interface{}

func main() {
	s1 = &t{}
	s2 = make(map[interface{}]interface{})
	go console()
	select {}
}

func console() {
	reader := bufio.NewReader(os.Stdin)
	for {
		input, _ := reader.ReadString('\n')
		commands := strings.Split(strings.Split(input, "\n")[0], " ")
		if len(commands) < 3 {
			for _, c := range commands {
				if c == "exit" {
					os.Exit(1)
				}
			}
		} else {
			if commands[2] == "1" {
				Read_config(commands[0], commands[1], s1)
				log.Println(s1)
			}
			if commands[2] == "2" {
				Read_config(commands[0], commands[1], &s2)
				log.Println(s2)
			}
			if commands[2] == "3" {
				Read_config(commands[0], commands[1], s2)
				log.Println(s2)
			}
		}
	}
}
