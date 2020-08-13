package app

import (
	"fmt"
	"os/exec"
)

func CheckCommitType(t string) bool {
	switch t {
	case "feat":
	case "fix":
	case "refactor":
	case "docs":
	case "style":
	case "test":
	case "chore":
	case "pref":
	case "build":
	case "ci":
	case "revert":
		return true
		break
	default:
		return false
	}
	return true
}

func CheckCommitSubjectLength(subject string) bool {
	if len(subject) < 24 {
		fmt.Println(len(subject))
		return false
	}
	return true
}

func SubmitCommit(s string) {
	shell := exec.Command(GIT_SHELL_NAME, "commit", "-m", s)
	output, err := shell.Output()
	if err != nil {
		fmt.Printf("Execute Shell:%s \nfailed with error:%s\n", "", err.Error())
		//return nil, ""
		return
	}
	fmt.Println(string(output))
}
