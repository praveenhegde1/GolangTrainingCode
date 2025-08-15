package main

import (
	"fmt"
)

func main() {
	// Declare and initialize a complex array
	complexArray := [3][2]complex128{
		{1 + 2i, 3 + 4i},
		{5 + 6i, 7 + 8i},
		{9 + 10i, 11 + 12i},
	}

	intarray := [3][2] int{
		{1,2},
		{3,4},
        {5,6},
	}


     for i :=0; i<len(intarray); i++ {
		for j :=0; j<len(intarray[i]); j++ {
			fmt.Printf("  %d", intarray[i][j] )
	 }
	 fmt.Println()
	}

	// Print the complex array
	for i := 0; i < len(complexArray); i++ {
		for j := 0; j < len(complexArray[i]); j++ {
			fmt.Printf("%v ", complexArray[i][j])
		}
		fmt.Println()
	}
}