package golang_util

import "os"

import "fmt"
import "path"

func InitializeEnv(dir string) {
	// cmd: go mod init
	cPath := path.Clean(dir)
	file, err := os.Create(cPath + "/go.mod")
	if err != nil {
		m := fmt.Errorf("Failed to Create File: %s", cPath+"go.mod")
		fmt.Println(m)
	}
	buf := make([]byte, 1024)
	_, err = file.Read(buf)
	if err != nil {
		m := fmt.Errorf("Failed to Read File: %s", cPath+"/go.mod")
		fmt.Println(m)
	}
	fmt.Println(string(buf))
}
