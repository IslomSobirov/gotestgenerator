package main

import (
	"fmt"

	"github.com/testapp/models/test"
)

func main() {
	// test.CreateOption("First option", true, 1, 1)
	fmt.Println(test.GetByQuestionID(1))
	// test.DeleteOption(1)
}
