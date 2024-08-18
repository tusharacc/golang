/*
Different variant of loop, which iterates over slice
*/
package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

/*
In go one cannot have
1. Unused packages
2. Unused variable will result in compilation error
*/
