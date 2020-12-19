package main

import (
	"fmt"

	"github.com/testapp/models/test"
)

func main() {
	// test.CreateQuestion("second question", 2)
	fmt.Println(test.QuestionByTestID(2))
}
