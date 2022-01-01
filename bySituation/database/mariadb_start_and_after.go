package database

import (
	"fmt"
	"go-useful-project/customLibrarys/common"
)

const initCheckFile string = "/golang/init_success.txt"

/*
	커맨드로 mariadb 를 실행시키고 이후 설정해야 할 커맨드 추가 입력
*/
func MariadbStartAndAfter() {
	fmt.Println("MariadbStartAndAfter 함수 시작")

	// mariadb 실행
	var result = common.Command("service", "mariadb", "start")
	fmt.Println("result", result)

	if !isInit() {
		fmt.Println("init 되지 않음.. init 시작!")

		// mariadb 의 root 계정 비밀번호를 112233abc 으로 변경
		var result2 = common.Command("mysql", "-e", "set password=password('112233abc');FLUSH PRIVILEGES;")
		fmt.Println("result2", result2)

		// 파일 생성
		if !common.IsFileOrFolderExist("/golang") {
			// 폴더 생성
			common.CreateFolder("/golang")
		}
		common.CreateFile(initCheckFile)
	} else {
		fmt.Println("init 된 상태!")
	}
}

func isInit() bool {
	var isInit = common.IsFileOrFolderExist(initCheckFile)
	return isInit
}
