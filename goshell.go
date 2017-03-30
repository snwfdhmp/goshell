package main

import (
"fmt"
"bufio"
"os"
"os/user"
"./Handler"
"./Pim"
//"./Executer"
"strings"
"path/filepath"
//"github.com/fatih/color"
)


func Prompter(prefix ...string) string {
	//color.Red("We have red")
	u,_ := user.Current()
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))

	if strings.Contains(dir, u.HomeDir) {
		replacer := strings.NewReplacer(u.HomeDir, "~")
		dir = replacer.Replace(dir)
	}

	prompt := u.Username+" @ "+dir+" : "

	return prompt
}

func Prompt(reader *bufio.Reader, state int) int {
	if(state == 0) {
		pim.Success("-> ")
	} else {
		pim.Error("-> ")
	}

	fmt.Print(Prompter())
	input, err := reader.ReadString('\n')
	if(err != nil) {
		fmt.Println(err)
	}
	replacer := strings.NewReplacer("\n", "")
	input = replacer.Replace(input)
	return Handler.Handle(input)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var status = 0;
	for status != -99 {
		status = Prompt(reader, status)
	}
}
