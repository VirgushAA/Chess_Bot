package main

import (
	. "Chess_Bot/core"
	"fmt"
)

var ()

func main() {
	var (
		name string
	)
	fmt.Print("Kto zdes?\n")
	_, err := fmt.Scanf("%s", &name)
	if err != nil {
		fmt.Println("ERROR:", err)
	} else {
		fmt.Println(Greatings(name))
	}
}
