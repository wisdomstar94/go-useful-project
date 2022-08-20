package dockerpscheck

import (
	"bytes"
	"fmt"
	"go-useful-project/customLibrarys/common"
	"os/exec"
	"strings"
)

type Container struct {
	Name string
}

// docker container 가 꺼져 있는지 켜져 있는지 체크 할 컨테이명들을 입력합니다.
var targetDockerContainerNames = []Container{
	{Name: "test-container-1"},
}

// 컨테이너가 꺼져 있는지 체크하고 꺼져 있다면 해당 컨테이너를 구동합니다.
func Check() {
	cmd := exec.Command("docker", "ps")
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err2 := cmd.Run()
	if err2 != nil {
		fmt.Println("** cmd", cmd)
		fmt.Println(fmt.Sprint(err2) + ": " + stderr.String())
		return
	}
	var result = out.String()
	for _, value := range targetDockerContainerNames {
		if !strings.Contains(result, common.SumString([]string{"   ", value.Name}, "")) {
			fmt.Println(common.SumString([]string{value.Name, " 컨테이너는 실행상태가 아닙니다. 아래 명령어로 실행시킵니다.."}, ""))
			fmt.Println(common.SumString([]string{"docker", " ", "start", " ", value.Name}, ""))
			cmd2 := exec.Command("docker", "start", value.Name)
			var out2 bytes.Buffer
			var stderr2 bytes.Buffer
			cmd2.Stdout = &out2
			cmd2.Stderr = &stderr2
			err3 := cmd2.Run()
			if err3 != nil {
				fmt.Println("** cmd2", cmd2)
				fmt.Println(fmt.Sprint(err3) + ": " + stderr2.String())
			}
		} else {
			fmt.Println(common.SumString([]string{value.Name, " 컨테이너는 실행 상태입니다."}, ""))
		}
	}
}
