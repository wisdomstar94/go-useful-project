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

	// 최초 초기화 작업이 이루어졌는지 아닌지 체크
	if !isInit() {
		fmt.Println("init 되지 않음.. init 시작!")

		// mariadb 의 root 계정 비밀번호를 112233abc 으로 변경
		var result2 = common.Command("mysql", "-e", "set password=password('112233abc');FLUSH PRIVILEGES;")
		fmt.Println("result2", result2)

		// 폴더 생성 (이미 존재하는 폴더면 생성하지 않음)
		common.CreateFolder("/golang")

		// 파일 생성 (이미 존재하는 파일이면 생성하지 않음)
		common.CreateFile(initCheckFile)
	} else {
		fmt.Println("init 된 상태!")
	}
}

func isInit() bool {
	var isInit = common.IsFileOrFolderExist(initCheckFile)
	return isInit
}
