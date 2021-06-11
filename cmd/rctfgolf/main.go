package main

import (
	"fmt"
	"os"
	"time"

	"github.com/antonmedv/expr"
	rctfgolf "github.com/redpwn/rctf-golf"
	"github.com/spf13/pflag"
)

type globals struct {
	Elapsed time.Duration
}

func main() {
	baseURL := pflag.StringP("base-url", "u", "", "Base URL of rCTF")
	challID := pflag.StringP("chall-id", "c", "", "Challenge ID")
	expression := pflag.StringP("function", "f", "", "Value expression to evaluate")

	pflag.Parse()

	if *baseURL == "" {
		fmt.Println("Must provide base-url\n")
		pflag.Usage()
		os.Exit(2)
	}
	if *challID == "" {
		fmt.Println("Must provide chall-id\n")
		pflag.Usage()
		os.Exit(2)
	}
	if *expression == "" {
		fmt.Println("Must provide function\n")
		pflag.Usage()
		os.Exit(2)
	}

	program, err := expr.Compile(*expression, expr.Env(globals{}))
	if err != nil {
		fmt.Printf("Invalid expression: %v\n", err)
		os.Exit(1)
	}

	elapsed, err := rctfgolf.GetTime(*baseURL, *challID)
	if err != nil {
		panic(err)
	}
	result, err := expr.Run(program, globals{Elapsed: elapsed})
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}
