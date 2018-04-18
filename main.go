package main 

import (
"fmt"
"github.com/agoda/goNetAgent/model"
)


func main(){

	options := []string{"ping","-c","10","google.com"}

	ping := model.NewPing(options)

	_, err := ping.Run()
	if err != nil {
		fmt.Println(err)
	} else{
		output, err2 := ping.GetResults()

		if err2 != nil {
			fmt.Println(err2)
		} else{
			fmt.Println(string(output))
		}
	}
}