package main

import (
	"bufio"
	// "fmt"
	log "github.com/Sirupsen/logrus"
	"os"
	"strings"
	"tool/config"
)

type t struct {
	Name  string
	Lv    int64
	Equip []int64
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
		input = string([]byte(input)[0 : len([]byte(input))-2])
		commands := strings.Split(input, " ")
		if len(commands) < 3 {
			for _, c := range commands {
				if c == "exit" {
					os.Exit(1)
				}
			}
		} else {
			if commands[2] == "1" {
				config.Read_config(commands[0], commands[1], s1)
				log.Println(s1)
			}
			if commands[2] == "2" {
				config.Read_config(commands[0], commands[1], &s2)
				log.Println(s2)
			}
			if commands[2] == "3" {
				config.Read_config(commands[0], commands[1], s2)
				log.Println(s2)
			}
		}
	}
}
