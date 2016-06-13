package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if (len(args) < 3) {
		fmt.Println("########### Expecting two Argumemnts ... Exiting")
		return
	}
	fmt.Println("############### Arguments supplied are :", args[1],args[2])	
}
