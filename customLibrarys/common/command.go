package common

import (
	"fmt"
	"os/exec"
)

/*
	패키지 로드시 호출되는 init 함수
*/
func init() {

}

/*
	커맨드 날리고 결과 받는 함수
*/
func Command(command string, arg ...string) string {
	cmd := exec.Command(command, arg...)
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Command Error :", err)
		return "COMMAND_ERROR"
	} else {
		fmt.Println("Command Result", string(output))
		return string(output)
	}
}
