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
	파일 생성하는 함수
*/
func CreateFile(filePath string) bool {
	log.Println("CreateFile 함수 호출됨.. filePath=", filePath)

	if IsFileOrFolderExist(filePath) {
		return false
	}

	_, err := os.Create(filePath)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	return true
}
