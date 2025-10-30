package main

import (
	"fmt"
	"os"
	"os/exec"
)

var (
	args = os.Args[1:]
)

func main() {
	file := ""
	for i := 0; i < len(args); i++ {
		switch args[i] {
			case "-h":
				help()
			default:
				file = args[i]
		}
	}
	if file == "" {
		eror("No input file", nil)
	} else {
		err, stderr := runFile(file)
		if err != nil {
			eror(stderr, err)
		}
	}
}

func runFile(file string) (error, string) {
	input := fmt.Sprintf("so %s", file)
	cmd := exec.Command("vim", "--clean", "-n", "--not-a-term", "--cmd", input, "--cmd", "qall!")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return err, string(out)
	}

	fmt.Printf(string(out))
	return nil, ""
}

func help() {
		var stuff = []string{
			"VIMc - help",
			"  usage",
		  "    vimc [file path]",
			"      (eg: vimc /foo/bar/baz.vim)",
			"  -h",
			"    show this help message",
		}
		for i := 0; i < len(stuff); i++ {
			fmt.Println(stuff[i])
		}
}

func eror(str string, err error) {
	if err != nil {
		str = fmt.Sprintf("err:  %v\n", err)
	}
	fmt.Fprintln(os.Stderr, str)
	os.Exit(1)
}
