package pim

import "fmt"

func Success(args...string) {
	green := "\033[32m"
	reset := "\033[39m"
	fmt.Printf("%s", green)
	for _, arg := range args {
		fmt.Printf("%s", arg)
	}
	fmt.Printf("%s", reset)
}

func Error(args...string) {
	red := "\033[31m"
	reset := "\033[39m"
	fmt.Printf("%s", red)
	for _, arg := range args {
		fmt.Printf("%s", arg)
	}
	fmt.Printf("%s", reset)
}

func Info(args...string) {
	cyan := "\033[36m"
	reset := "\033[39m"
	fmt.Printf("%s", cyan)
	for _, arg := range args {
		fmt.Printf("%s", arg)
	}
	fmt.Printf("%s", reset)
}