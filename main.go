package main

import (
	"fmt"

	"github.com/jay13jay/livepeer/create"
)

func main() {
	newStream := create.CreateStream()

	fmt.Printf("newStream:\n%s\n\n", newStream)

}
