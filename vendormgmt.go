package main

import (
	"errors"
	"fmt"
)

type Vendor struct {
 Name string
 ContactInfo string
 PerformanceRating int

}

func EvaluateVendor(vendors []*Vendor) (*Vendor, error) {
	if len(vendors) == 0 {
		return nil, errors.New("no vendors found")
    }
	
	bestVendor := vendors[0]

	for _, vendor := range vendors {
		if vendor.PerformanceRating > bestVendor.PerformanceRating {
            bestVendor = vendor
        }
	}
	return bestVendor, nil
}
func main() {	
	vendors := []*Vendor{

	/*	{ Name: "Vendor ABC", ContactInfo: "9900992060", PerformanceRating: 0 },
		{ Name: "Vendor BCD", ContactInfo: "9980812060", PerformanceRating: 2 },
		{ Name: "Vendor CBZ", ContactInfo: "8940720060", PerformanceRating: 4 }, */
	}
	 bestVendor, err := EvaluateVendor(vendors)
	 if err!= nil {
         fmt.Println("Error: ", err)
         return
	 } 
	 fmt.Println("Best Vendor: ", bestVendor.Name)

}