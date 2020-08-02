package app

import (
	"fmt"
	"os/exec"
)

/**
 * 获取当前版本数据
 */
func GetVersionInfo() {
	split := " -|---|- "
	//common := `git log ...master  --format="%H` + split + `%ci` + split + `%ce` + split + `%s"`
	//common := "git version"
	shell := exec.Command("git","log",`--format=%H` + split + `%ci` + split + `%ce` + split + `%s`)


	//fmt.Println(command)

	output, err := shell.Output()
	if err != nil {
		fmt.Println(err)
		fmt.Println(output)
		//fmt.Printf("Execute Shell:%s failed with error:%s", command, err.Error())
		return
	}
	fmt.Println(err)
	fmt.Println(string(output))


}
