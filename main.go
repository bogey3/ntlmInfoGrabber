package main

import (
	"fmt"
	"github.com/bogey3/NTLM_Info"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Please provide a URL as an argument.")
		return
	}

	input := os.Args[1]
	if target, err := NTLM_Info.NewTarget(input); err == nil {
		if err2 := target.GetChallenge(); err2 == nil {
			target.Print()
		} else {
			panic(err2)
		}
	} else {
		panic(err)
	}
}
