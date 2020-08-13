package app

import (
	"fmt"
	"os/exec"
)

func GitCommit(commitType, commitSubject string) bool {
	var commond = []string{GIT_SHELL_NAME, "commit", "-m", commitType + ": " + commitSubject}
	shell := exec.Command(GIT_SHELL_NAME, "commit", "-m", commitType+": "+commitSubject)
	output, err := shell.Output()
	outString := string(output)
	if err != nil {
		fmt.Printf("Execute Shell:%s\n failed with error:%s\n failed with output:\n %s", commond, err.Error(), outString)
		return false
	}
	fmt.Println(outString)
	return true
}

func GitAdd(r string) bool {
	var commond = []string{GIT_SHELL_NAME, "add", r}
	shell := exec.Command(GIT_SHELL_NAME, "add", r)
	output, err := shell.Output()
	outString := string(output)
	if err != nil {
		fmt.Printf("Execute Shell:%s\n failed with error:%s\n failed with output:\n %s", commond, err.Error(), outString)
		return false
	}
	fmt.Println(outString)
	return true
}

func GitAddAll() bool {
	return GitAdd("--all")
}

func GitPull(remote, branch string) bool {
	var commond = []string{GIT_SHELL_NAME, "pull", remote, branch}
	shell := exec.Command(GIT_SHELL_NAME, "pull", remote, branch)
	output, err := shell.Output()
	outString := string(output)
	if err != nil {
		fmt.Printf("Execute Shell:%s\n failed with error:%s\n failed with output:\n %s", commond, err.Error(), outString)
		return false
	}
	fmt.Println(outString)
	return true
}

func GitFetch(remote, branch string) bool {
	var commond = []string{GIT_SHELL_NAME, "fetch", remote, branch}
	shell := exec.Command(GIT_SHELL_NAME, "fetch", remote, branch)
	output, err := shell.Output()
	outString := string(output)
	if err != nil {
		fmt.Printf("Execute Shell:%s\n failed with error:%s\n failed with output:\n %s", commond, err.Error(), outString)
		return false
	}
	fmt.Println(outString)
	return true
}
