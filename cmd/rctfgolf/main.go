package main

import (
	"fmt"
	"math"
	"os"
	"time"

	"github.com/antonmedv/expr"
	rctfgolf "github.com/redpwn/rctf-golf"
	"github.com/spf13/pflag"
)

func run() error {
	baseURL := pflag.StringP("base-url", "u", "", "Base URL of rCTF")
	challID := pflag.StringP("chall-id", "c", "", "Challenge ID")
	expression := pflag.StringP("function", "f", "", "Value expression to evaluate")

	pflag.Parse()

	if *baseURL == "" {
		fmt.Fprint(os.Stderr, "Must provide base-url\n\n")
		pflag.Usage()
		os.Exit(2)
	}
	if *challID == "" {
		fmt.Fprint(os.Stderr, "Must provide chall-id\n\n")
		pflag.Usage()
		os.Exit(2)
	}
	if *expression == "" {
		fmt.Fprint(os.Stderr, "Must provide function\n\n")
		pflag.Usage()
		os.Exit(2)
	}

	globals := map[string]interface{}{
		"elapsed": time.Duration(0),
		"int":     func(v float64) int { return int(v) },
		"round":   math.Round,
		"floor":   math.Floor,
		"ceil":    math.Ceil,
	}

	program, err := expr.Compile(*expression, expr.Env(globals))
	if err != nil {
		return fmt.Errorf("invalid expression: %w", err)
	}

	elapsed, err := rctfgolf.GetTime(*baseURL, *challID)
	if err != nil {
		return err
	}
	globals["elapsed"] = elapsed
	result, err := expr.Run(program, globals)
	if err != nil {
		return err
	}

	fmt.Println(result)
	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
