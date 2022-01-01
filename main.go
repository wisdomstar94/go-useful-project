package main

import (
	"go-useful-project/customLibrarys/common"
	"log"
)

func main() {
	testUuid()
	// database.MariadbStartAndAfter()
	// testCommand()
	// testSumString()
}

func testUuid() {
	log.Println("testUuid() 함수 호출됨!")
	var result = common.GetUuid()
	log.Println("result", result)
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
