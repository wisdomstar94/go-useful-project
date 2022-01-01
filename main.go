package main

import (
	"go-useful-project/bySituation/database"
	"go-useful-project/customLibrarys/common"
	"log"
)

func main() {
	database.MariadbStartAndAfter()
	// testCommand()
	// testSumString()
}

func testCommand() {
	var result = common.Command("netstat", "-nap", "|", "grep", "LISTEN")
	log.Println("result", result)
}

func testSumString() {
	var stringValues = []string{"1", "2"}
	// stringValues = []string{"1", "2"}

	var result = common.SumString(stringValues, ",")
	log.Println(result)
}
