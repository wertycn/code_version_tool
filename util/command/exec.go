package command

import (
	"fmt"
	"os/exec"
	"runtime"
)

func Run(command string, args string) {
	//command := `./dir_size.sh .`
	sysType := runtime.GOOS
	var shell *exec.Cmd
	//fmt.Println(sysType)
	if sysType == "windows" {

		// windows系统
		shell = exec.Command( command,args)
		fmt.Println(command)

	}else {
		//shell = exec.Command("/bin/bash",  command)

	}

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
	fmt.Printf("Execute Shell:%s finished with output:\n%s", command, string(output))

}
