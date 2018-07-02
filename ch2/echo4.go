
package main

import ("fmt"
		"flag"
		"strings")

var n = flag.Bool("n",false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

func mainEcho2() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))

	if !*n {
		fmt.Println()
	}
}