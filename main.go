package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

func main() {
	mode := flag.String("mode", "", "hex or dec")
	value := flag.String("value", "", "value")
	flag.Parse()

	if *mode == "" {
		fmt.Fprintln(os.Stderr, "mode is empty")
		os.Exit(1)
	}
	if *value == "" {
		fmt.Fprintln(os.Stderr, "value is empty")
		os.Exit(1)
	}

	switch *mode {
	case "hex":
		pprint(*value, 16, "%d")
	case "dec":
		pprint(*value, 10, "0x%x")
	}
}

func pprint(value string, base int, format string) {
	v, err := strconv.ParseInt(value, base, 64)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Printf(format, v)
}
