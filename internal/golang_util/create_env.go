package golang_util

import (
	"fmt"
	"os"
	"os/exec"
)

func InitializeEnv(dir string) {
	pathStat, err := os.Stat(dir)
	if err != nil {
		m := fmt.Errorf("%v", err)
		fmt.Println(m)
		os.Exit(1)
	}
	if !pathStat.IsDir() {
		m := fmt.Errorf("%v is not directory", pathStat.Name())
		fmt.Println(m)
		os.Exit(1)
	}
	pwd, _ := os.Getwd()
	os.Chdir(dir)
	cmd, err := exec.Command("go", "mod", "init").Output()
	if err != nil {
		fmt.Printf("execute cmd error: %v", err)
		os.Exit(1)
	}
	fmt.Printf("%v\n", cmd)
	buf := make([]byte, 1024)
	f, err := os.OpenFile(string(cmd)+"go.mod", os.O_RDONLY, 0755)
	f.Read(buf)
	if err != nil {
		fmt.Printf("Failed Read: %v", err)
		os.Exit(1)
	}
	if err := f.Close(); err != nil {
		fmt.Printf("err: %v", err)
		os.Exit(1)
	}
	os.Chdir(pwd)
	fmt.Println(string(buf))
}
