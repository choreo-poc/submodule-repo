package main

import (
	"fmt"
	submodule "submodules/nestedsubmodule"

	
)

func Hello() string {
	return "Hello from main repo"
}


func main() {
	fmt.Println("result from Hello():", Hello())
	fmt.Println("result from nestedsubmodule.HelloMsg():", submodule.HelloMsg())
}
