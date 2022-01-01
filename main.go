package main

import (
	"fmt"
	"go-useful-project/bySituation/database"
	"go-useful-project/customLibrarys/common"
)

func main() {
	database.MariadbStartAndAfter()
	// testCommand()
	// testSumString()
}

func testCommand() {
	var result = common.Command("netstat", "-nap", "|", "grep", "LISTEN")
	fmt.Println("result", result)
}

func testSumString() {
	var stringValues = []string{"1", "2"}
	// stringValues = []string{"1", "2"}

	var result = common.SumString(stringValues, ",")
	fmt.Println(result)
}
