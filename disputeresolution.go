package main 

import (
	"fmt"

)


type Dispute struct {
Description string
Status string
Resolution string
}
func EvaluateDispute(disputes []*Dispute)  {

	for _, dispute := range disputes {
		fmt.Println("Description: ", dispute.Description)
		fmt.Println("Status: ", dispute.Status)
		fmt.Print("Resolution: ", dispute.Resolution)
}
}


func main(){
dispute := []*Dispute{
	{Description: "Customer not receiving product", Status: "Open", Resolution: "Contact customer"},
	{Description: "Product not as described", Status: "Closed", Resolution: "Return product to manufacturer"},
    {Description: "Customer complained about poor service", Status: "Open", Resolution: "Investigate and resolve"},	
	
}

 EvaluateDispute(dispute)


//fmt.Print("Dispute Status: ", disputeStatus)



}

