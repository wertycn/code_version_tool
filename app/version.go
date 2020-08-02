package app

import (
	"fmt"
)
import "gopkg.in/src-d/go-git.v4"
/**
 * 获取当前版本数据
 */
func GetVersionInfo() {
	//split := " -|---|- "
	r, err := git.PlainOpen("./")
	if err != nil{

	}

	ref, err := r.Head()

	//since := time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC)
	//until := time.Date(2019, 7, 30, 0, 0, 0, 0, time.UTC)
	w, err := r.Log(&git.LogOptions{From: ref.Hash()})
	fmt.Println(w)
	//common := `git log ...master  --format="%H` + split + `%ci` + split + `%ce` + split + `%s"`
	//common := "git branch"
	//command.Run(common)

}
