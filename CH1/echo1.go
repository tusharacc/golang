/*
IMPLEMENTS UNIX ECHO PROGRAM
*/

//Understanding Command Line Arguments
// Command line arguments are slices.
// Slices are dynamically sized sequences. This is a contiguous memory
// s[m:n] -> 0 <= m <=n <= len(s), contains n-m element
// half-open: contains the first index but excludes the last.

package main

import (
	"fmt"
	"os"
)

/*
In GO variables are implicitly declared. In case of integer it is 0
In case of string it is ""

:= -> short variable declaration.Compiler gives a type based on initializer value
*/

func main() {
	var s, sep string //variable s, sep declaration which is of type string.
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

/*

i++ is not an expression but statement in GO
for-loop is the only loop in go

for loop initialization must be a simple statement that is
	1. short variable declaration
	2. an increment or assignment operator
	3. or a function call

for initialization; condition; post { <-- the brace has to be on same line

}

traditional while loop can be written as
for condition {

}

infinite loop
for {

}

*/
