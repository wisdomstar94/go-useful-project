package common

import (
	"log"
	"os"
)

/*
	패키지 로드시 호출되는 init 함수
*/
func init() {

}

/*
	폴더 생성하는 함수
*/
func CreateFolder(folderPath string) bool {
	log.Println("CreateFolder 함수 호출됨.. folderPath=", folderPath)
	if IsFileOrFolderExist(folderPath) {
		return false
	}

	// Command("mkdir", folderPath)
	err := os.Mkdir(folderPath, 0755)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}
