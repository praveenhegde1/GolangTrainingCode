package main

import( 
	"fmt"
	"os"
)

func main() {

	args := os.Args
	fmt.Println("number of arguments",len(args))

	for i, arg := range args {
        fmt.Printf("Argument %d: %s\n", i+1, arg)
    }

	if len(args) > 1 {
	firstarg := args[1]
	fmt.Println("First argument:", firstarg)
	}


}
