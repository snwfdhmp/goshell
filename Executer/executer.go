package Executer

import (
"fmt"
"io/ioutil"
//"github.com/fatih/color"
"../Pim"
//"strings"
)

func Exec(command string, args []string) int {
	status := 0 // return status
	switch(command) {
	case "print" :
		for _, arg := range args {
			fmt.Print(arg, " ")
		}
		fmt.Print("\n")
		break
	case "create":
		var content []byte
		if(len(args) <= 0) {
			pim.Error("Usage : create path/to/file [content]", "\n")
			break;
		} else if (len(args)>1) {
			content = []byte(args[1])
		} else {
			content = []byte("")
		}
		ioutil.WriteFile(args[0], content, 0600)
		break
	case "ls":
		if(len(args) <= 0) {
			args = []string {"./"}
		}
		for _, arg := range args {
			ListDirectory(arg)
		}
		break
	case "cat":
		if(len(args) <= 0) {
			pim.Error("Usage : cat path/to/file", "\n")
			break;
		} else {
			for _, path := range args {
				DisplayFile(path)
			}
		}
		break
	case "quit":
		pim.Success("See you soon :-)", "\n")
		status= -99
	default :
		pim.Error("Erreur: Commande inconnue : ", command, "\n")
		status = -1
		break
	}

	return status
}

func ListDirectory(path string) {
	files, _ := ioutil.ReadDir(path)
	for _, f := range files {
		pim.Info(f.Name(), "\n")
	}
}

func DisplayFile(path string) {
	file, _ := ioutil.ReadFile(path)
	pim.Info(string(file))
}