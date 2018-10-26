// Echo1 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	// initialization; condition; post
	for i := 1; i < len(os.Args); i++ { // i++ -> i += 1 -> i = i + 1
		s += sep + os.Args[i] // s = s + sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
