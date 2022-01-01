package main

import (
	"fmt"
	"go-useful-project/customLibrarys/common"
)

func main() {
	var stringValues = []string{"1", "2"}
	// stringValues = []string{"1", "2"}

	var result = common.SumString(stringValues, ",")
	fmt.Println(result)
}
