package Handler

import (
//"fmt"
"strings"
"../Executer"
)

func Parse(str string) []string { //"i am a smart command" => ["i", "am", "a", "smart", "command"]
	//fmt.Println("Parser called")

	return strings.Split(str, " ")
}

func Handle(input string) int{
	//fmt.Println("Handler called")

	request := Parse(input) // split string into array

	return Executer.Exec(request[0], request[1:])
}