package common

import "fmt"

/*
	패키지 로드시 호출되는 init 함수
*/
func init() {

}

/*
	폴더 생성하는 함수
*/
func CreateFolder(folderPath string) bool {
	fmt.Println("CreateFolder 함수 호출됨.. folderPath=", folderPath)
	if IsFileOrFolderExist(folderPath) {
		return false
	}

	Command("mkdir", folderPath)
	return true
}
